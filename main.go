package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	configFile, err := os.Open("./config.json")

	if err != nil {
		if os.IsNotExist(err) {
			config := map[string]map[string]string{
				"paths": {
					"data": "./data",
				},
			}
			file, _ := json.MarshalIndent(config, "", "")
			if err := os.WriteFile("config.json", file, 0644); err != nil {
				log.Fatal(err)
			}
		}
	}

	defer configFile.Close()

	app.Get("/bot", func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("./settings/bot.json")
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(jsonFile)

		if err != nil {
			return c.SendString("Error reading file")
		}
		byteValue, _ := io.ReadAll(jsonFile)

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(byteValue), &result); err != nil {
			return c.SendString("error")
		}
		return c.JSON(result)
	})

	app.Get("/config", func(c *fiber.Ctx) error {
		byteValue, _ := io.ReadAll(configFile)

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(byteValue), &result); err != nil {
			return c.SendString("error")
		}

		return c.JSON(result)
	})

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
