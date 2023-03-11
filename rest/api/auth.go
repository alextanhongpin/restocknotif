package api

type LoginRequest struct {
	Name string `json:"name"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
