package apigen 

import "github.com/gofiber/fiber/v2"

type AuthFunc func(c *fiber.Ctx, rules ...string) error

func RegisterAuthFunc(app *fiber.App, f AuthFunc) {
	
	app.Post("/api/v1/clusters", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/clusters", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Delete("/api/v1/clusters/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Put("/api/v1/clusters/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/clusters/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/clusters/:id/diagnostics", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Put("/api/v1/clusters/:id/diagnostics/config", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/clusters/:id/diagnostics/config", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/clusters/:id/snapshot-config", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Put("/api/v1/clusters/:id/snapshot-config", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/clusters/:id/snapshots", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Post("/api/v1/clusters/:id/snapshots", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Post("/api/v1/clusters/:id/snapshots/:snapshotId", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Delete("/api/v1/clusters/:id/snapshots/:snapshotId", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/databases", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Post("/api/v1/databases", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Delete("/api/v1/databases/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Put("/api/v1/databases/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
	app.Get("/api/v1/databases/:id", func(c *fiber.Ctx) error { 
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} 
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		
		return c.Next()
	})
}
