package handlers

import (
	"api_study/interfaces"
	"api_study/pkg/bcrypto"
	"api_study/pkg/encrypt"
	"api_study/pkg/storage"
	users_repo "api_study/repository/users"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var err error
	userRequest := new(interfaces.UserRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return err
	}

	currentUser, _ := users_repo.RetrieveUserByEmail(userRequest.Email)
	if currentUser.UserID != 0 {
		return c.Status(400).JSON(map[string]string{"error": "user already exists"})
	}

	userRequest.Password, err = encrypt.Decrypt(userRequest.Password)
	if err != nil {
		return err
	}

	userRequest.Password, err = bcrypto.HashPassword(userRequest.Password)
	if err != nil {
		return err
	}

	userPicture, err := c.FormFile("user_picture")

	user, _ := users_repo.CreateUser(userRequest)
	if err == nil {
		destination := fmt.Sprintf("uploads/users/%s", strconv.Itoa(user.UserID))

		path, err := storage.StoreFile(c, destination, userPicture)
		if err != nil {
			return err
		}

		user.UserPicture = path
	}

	user, _ = users_repo.UpdateUser(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("user_id"))

	userRequest := new(interfaces.UserRequest)

	err := c.BodyParser(userRequest)

	if err != nil {
		return err
	}

	currentUser, _ := users_repo.RetrieveUserByID(userId)
	if currentUser.UserID == 0 {
		return c.Status(400).JSON(map[string]string{"error": "user not found"})
	}

	userRequest.UserID = userId

	user, _ := users_repo.UpdateUser(userRequest)

	userPicture, err := c.FormFile("user_picture")
	if err == nil {
		destination := fmt.Sprintf("uploads/users/%s", string(userId))

		path, err := storage.StoreFile(c, destination, userPicture)
		if err != nil {
			return err
		}

		user.UserPicture = path
	}

	return c.JSON(user)
}

func RetriveUser(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("user_id"))

	user, _ := users_repo.RetrieveUserByID(userId)

	return c.JSON(user)
}

func ListUsers(c *fiber.Ctx) error {
	params := interfaces.ListParams{}

	err := c.QueryParser(&params)

	if err != nil {
		return err
	}

	users := users_repo.ListUsers()

	return c.JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("user_id"))

	user, _ := users_repo.RetrieveUserByID(userId)
	if user.UserID == 0 {
		return c.Status(400).JSON(map[string]string{"error": "user not found"})
	}

	users_repo.DeleteUser(userId)

	return c.JSON(map[string]string{"message": "User successfully deleted!"})
}
