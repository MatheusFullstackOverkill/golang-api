package handlers

import (
	"api_study/pkg/bcrypto"
	"api_study/pkg/encrypt"
	jwt_pkg "api_study/pkg/jwt"
	users_repo "api_study/repository/users"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Authtoken string `json:"auth_token"`
}

func Login(c *fiber.Ctx) error {
	loginData := LoginRequest{}

	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	decriptedPassword, err := encrypt.Decrypt(loginData.Password)
	if err != nil {
		return c.JSON(map[string]string{"error": "invalid request"})
	}

	user, _ := users_repo.RetrieveUserByEmail(loginData.Email)
	if user.UserID == 0 {
		return c.JSON(map[string]string{"error": "user not found"})
	}

	check := bcrypto.CheckPasswordHash(decriptedPassword, user.Password)

	if !check {
		return c.JSON(map[string]string{"error": "user not found"})
	}

	jwt, err := jwt_pkg.CreateToken(user.UserID)
	if err != nil {
		return c.JSON(map[string]string{"error": "an unexpected error has ocurred"})
	}

	jwt_pkg.VerifyToken(jwt)

	loginResponse := &LoginResponse{Authtoken: jwt}

	return c.JSON(loginResponse)
}
