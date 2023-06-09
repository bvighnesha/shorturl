package main

import (
	"github.com/bluele/gcache"
	"github.com/gofiber/fiber/v2"
	"sync"
)

func Start(app *fiber.App) {
	urls := URLs{
		Cache: gcache.New(1000000).
			ARC().Build(),
		URLPrefix: "http://short.url",
		Locker:    &sync.Mutex{},
		Metrics:   make(map[string]int),
	}

	app.Get("url/:url", urls.Handle)
	app.Get("metrics", urls.metrics)

	app.Listen(":3000")
}
