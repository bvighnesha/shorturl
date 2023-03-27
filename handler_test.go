package main

import (
	"fmt"
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
	resp, err := app.Test(req, 1)
	assert.NoError(t, err)

	resp.Body.Close()

	assert.NoError(t, err)
	url, err := resp.Location()
	assert.NoError(t, err)

	assert.Equal(t, fmt.Sprintf("%s%s", url.Host, url.Path), "short.url/1R9doEiEshDAGoYo")
}
