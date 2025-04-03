package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Url string `json:"url"`
}

var payload Payload

func MonitorUptime(c *fiber.Ctx) error {
	client := http.Client{Timeout: 5 * time.Second}

	if err := c.BodyParser(&payload); err != nil {
		log.Fatal("Error in getting request")
	}

	resp, err := client.Get(payload.Url)
	if err != nil {
		fmt.Printf("[%s] Website DOWN: %s\n", time.Now().Format(time.RFC3339), err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[%s] Website UP: %d OK, SERVER NAME %s\n", time.Now().Format(time.RFC3339), resp.StatusCode, resp.TLS.SignedCertificateTimestamps[0])
	} else {
		fmt.Printf("[%s] Website DOWN: Status Code %d\n", time.Now().Format(time.RFC3339), resp.StatusCode)
	}
	return nil
}
