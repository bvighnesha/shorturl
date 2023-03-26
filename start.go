package main

import (
	"github.com/bluele/gcache"
	"github.com/gofiber/fiber/v2"
	"sync"
)

func Start() {
	urls := URLs{
		Cache: gcache.New(1000000).
			ARC().Build(),
		URLPrefix: "http://short.url",
		Locker:    &sync.Mutex{},
		Metrics:   make(map[string]int),
	}
	app := fiber.New()

	app.Get("url/:url", urls.Handle)
	app.Get("metrics", urls.metrics)

	app.Listen(":3000")
}
