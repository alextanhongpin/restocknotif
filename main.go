package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/alextanhongpin/restocknotif/rest/api"
	"github.com/alextanhongpin/restocknotif/rest/middleware"
	"github.com/alextanhongpin/restocknotif/usecase/authn"
	"github.com/alextanhongpin/restocknotif/usecase/subscription"
	chi "github.com/go-chi/chi/v5"
)

//go:generate weaver generate ./...
func main() {
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("restock_notif", opts)
	if err != nil {
		log.Fatal(err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not provided")
	}

	authenticator := middleware.NewAuthenticator([]byte(jwtSecret), 1*time.Hour)

	authnUseCase, err := weaver.Get[authn.T](root)
	if err != nil {
		log.Fatal(err)
	}

	subscriptionUseCase, err := weaver.Get[subscription.T](root)
	if err != nil {
		log.Fatal(err)
	}

	subscriptionAPI := api.NewSubscriptionAPI(subscriptionUseCase)
	authAPI := api.NewAuthAPI(authnUseCase, authenticator)

	r := chi.NewRouter()

	// Protected routes.
	r.Group(func(r chi.Router) {
		r.Use(authenticator.Verifier)
		r.Use(authenticator.RequireAuth)

		r.Route("/subscriptions", func(r chi.Router) {
			r.Get("/", subscriptionAPI.GetAll)
			r.Post("/", subscriptionAPI.Post)

			r.Route("/{subscription_id}", func(r chi.Router) {
				r.Patch("/", subscriptionAPI.Patch)
				r.Delete("/", subscriptionAPI.Delete)
			})
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/login", authAPI.Login)
	})

	http.Serve(lis, r)
}
