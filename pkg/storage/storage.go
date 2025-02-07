package storage

import (
	config "api_study/init"
	"log"
	"os"

	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func StoreFile(c *fiber.Ctx, path string, file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := c.SaveFile(file, fmt.Sprintf("%s/%s", path, file.Filename)); err != nil {
		return "", err
	}

	file_url := fmt.Sprintf("%s/%s/%s", config.GlobaConfig.APIUrl, path, file.Filename)

	return file_url, nil
}

func DeleteFile(path string) (bool, error) {
	if err := os.Remove(path); err != nil {
		return false, err
	}

	return true, nil
}
