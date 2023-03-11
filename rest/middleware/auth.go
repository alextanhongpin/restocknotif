package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alextanhongpin/restocknotif/rest/response"
	"github.com/alextanhongpin/restocknotif/rest/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	jwtErrorCtxKey types.ContextKey[error]     = "jwt_error"
	jwtTokenCtxKey types.ContextKey[uuid.UUID] = "jwt_token"
)

var (
	ErrInvalidSigningMethod = errors.New("jwt: invalid signing method")
	ErrInvalidToken         = errors.New("jwt: invalid token")
	ErrNoToken              = errors.New("jwt: no token")
)

type Authenticator struct {
	secret        []byte
	validDuration time.Duration
}

func NewAuthenticator(secret []byte, validDuration time.Duration) *Authenticator {
	return &Authenticator{
		secret:        secret,
		validDuration: validDuration,
	}
}

func (a *Authenticator) Verifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := strings.Split(r.Header.Get("authorization"), "Bearer")
		if len(authHeader) == 2 {
			userID, err := a.parse(authHeader[1])
			if err != nil {
				ctx = jwtErrorCtxKey.WithValue(ctx, err)
			} else {
				ctx = jwtTokenCtxKey.WithValue(ctx, userID)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Authenticator) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err, _ := jwtErrorCtxKey.Value(ctx)
		if err != nil {
			response.Failure(w, err, http.StatusUnauthorized)
			return
		}

		_, ok := jwtTokenCtxKey.Value(ctx)
		if !ok {
			response.Failure(w, ErrNoToken, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Authenticator) parse(token string) (uuid.UUID, error) {
	claims, err := a.VerifyToken(token)
	if err != nil {
		return uuid.Nil, err
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return uuid.Nil, errors.Join(err, ErrInvalidToken)
	}

	userID, err := uuid.Parse(subject)
	if err != nil {
		return uuid.Nil, errors.Join(err, ErrInvalidToken)
	}

	return userID, nil
}

func (a *Authenticator) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		return nil, ErrNoToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Join(
				fmt.Errorf("%w: %v", ErrInvalidSigningMethod, token.Header["alg"]),
				ErrInvalidToken,
			)
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return a.secret, nil
	})
	if err != nil {
		return nil, errors.Join(err, ErrInvalidToken)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func (a *Authenticator) CreateToken(userID uuid.UUID) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.validDuration)),
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(a.secret)
	if err != nil {
		return "", errors.Join(err, ErrInvalidToken)
	}

	return ss, nil
}

func MustUserIDFromContext(ctx context.Context) uuid.UUID {
	return jwtTokenCtxKey.MustValue(ctx)
}
