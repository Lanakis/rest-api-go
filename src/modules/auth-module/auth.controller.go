package auth_module

import (
	"authorization/src/modules/auth-module/dto"
	"encoding/json"
	"net/http"
)

func Handler(c *AuthController) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			c.Login(w, r)

		}
	}
}

type AuthController struct {
	AuthService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	data := dto.SignAuthDto{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := c.AuthService.SignIn(r.Context(), data)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (c *AuthController) Profile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*User)
	json.NewEncoder(w).Encode(user)
}
