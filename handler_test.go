package main

import (
	"github.com/bluele/gcache"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestURLs_Handle(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()

	urls := URLs{
		Cache: gcache.New(1000000).
			ARC().Build(),
		URLPrefix: "http://short.url",
		Locker:    &sync.Mutex{},
		Metrics:   make(map[string]int),
	}

	app.Get("url/:url", urls.Handle)
	app.Get("metrics", urls.metrics)

	req := httptest.NewRequest("GET", "/url/abc.com", nil)
	resp, _ := app.Test(req, 1)

	assert.Equal(t, resp, assert.Equal(t, resp, nil))

}
