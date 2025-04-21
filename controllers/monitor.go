package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sumukhj1219/fiber-go/utils"
)

type Payload struct {
	Url    string `json:"url"`
	Wallet string `json:"wallet"`
}

func MonitorUptime(c *fiber.Ctx) error {
	client := http.Client{Timeout: 5 * time.Second}
	var payload Payload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).SendString("Invalid payload")
	}

	resp, err := client.Get(payload.Url)
	if err != nil {
		log.Printf("[%s] Website DOWN: %s\n", time.Now().Format(time.RFC3339), err)
		return c.Status(200).SendString("Website down")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("[%s] Website UP: %d OK\n", time.Now().Format(time.RFC3339), resp.StatusCode)

		go utils.RewardUser(payload.Wallet)
		return c.Status(200).SendString("Website up. Reward triggered.")
	} else {
		log.Printf("[%s] Website DOWN: Status Code %d\n", time.Now().Format(time.RFC3339), resp.StatusCode)
		return c.Status(200).SendString("Website returned error status.")
	}
}
