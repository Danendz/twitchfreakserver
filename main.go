package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"twitchfreakserver/entities"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configFile, err := os.Open("./config.json")

	if err != nil {
		if os.IsNotExist(err) {
			if err:=entities.NewConfig(); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		entities.SetConfigFromFile(configFile)
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
		return c.JSON(entities.GlobalConfig)
	})

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
