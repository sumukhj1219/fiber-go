package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sumukhj1219/fiber-go/controllers"
)

func TrackerRoutes(app *fiber.App) {
	app.Post("/track", controllers.TrackVisitors)
	app.Post("/visitors", controllers.GetVisitors)

	app.Get("/tracker.js", func(c *fiber.Ctx) error {
		script := ` (function () {
			const websiteId = document.currentScript.getAttribute("data-website-id");
			if (!websiteId) return;
		
			const trackingData = {
				website_id: websiteId,
				url: window.location.href,
				referrer: document.referrer,
				user_agent: navigator.userAgent,
				screen_size: window.screen.width + "x" + window.screen.height,
				time_stamp: new Date().toISOString(),
			};
		
			fetch("http://localhost:5000/track", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(trackingData)
			}).then(res => res.json()).then(console.log).catch(console.error);
		})();`

		c.Set("Content-Type", "application/javascript")
		return c.SendString(script)
	})
}
