package controllers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Visit struct {
	WebsiteId  string `json:"website_id"`
	Url        string `json:"url"`
	Referrer   string `json:"referrer"` 
	UserAgent  string `json:"user_agent"`
	ScreenSize string `json:"screen_size"`
	TimeStamp  string `json:"time_stamp"`
}

var (
	visitorCount = make(map[string]int)
	mutex        sync.Mutex
)

func TrackVisitors(c *fiber.Ctx) error {
	var visit Visit
	if err := c.BodyParser(&visit); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	mutex.Lock()
	visitorCount[visit.WebsiteId]++
	mutex.Unlock()

	log.Printf("New visitor on %s: Total %d", visit.WebsiteId, visitorCount[visit.WebsiteId])

	return c.JSON(fiber.Map{
		"message":   "Visitor tracked",
		"website":   visit.WebsiteId,
		"totalHits": visitorCount[visit.WebsiteId],
	})
}

func GetVisitors(c *fiber.Ctx) error {
	websiteId := c.Query("websiteId")
	if websiteId == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing websiteId"})
	}

	mutex.Lock()
	count, exists := visitorCount[websiteId]
	mutex.Unlock()

	if !exists {
		count = 0
	}

	return c.JSON(fiber.Map{
		"websiteId": websiteId,
		"totalHits": count,
	})
}
