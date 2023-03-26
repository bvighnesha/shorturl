package main

import (
	"fmt"
	"github.com/bluele/gcache"
	"github.com/btcsuite/btcutil/base58"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type URLs struct {
	Cache     gcache.Cache
	URLPrefix string
	Locker    *sync.Mutex
	Metrics   map[string]int
}

func (urls *URLs) Handle(c *fiber.Ctx) error {
	url := c.Params("url")
	if urls.Cache.Has(url) {
		value, ok := urls.Metrics[url]
		if ok {
			urls.Metrics[url] = value + 1
		}

		urls.Locker.Lock()
		existedURL, _ := urls.Cache.Get(url)
		urls.Locker.Unlock()
		return c.Redirect(existedURL.(string), 307)
	} else {
		urls.Metrics[url] = 1
		encoded := base58.CheckEncode([]byte(url), 0)
		encodedURL := fmt.Sprintf("%s/%s", urls.URLPrefix, encoded)
		urls.Locker.Lock()
		urls.Cache.Set(url, encodedURL)
		urls.Locker.Unlock()
		return c.Redirect(encodedURL, 307)
	}
}

func (urls *URLs) metrics(c *fiber.Ctx) error {
	metrics := urls.Metrics
	return c.SendString(fmt.Sprintln(metrics))
}
