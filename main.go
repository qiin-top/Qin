package main

import (
	"qin"
)

func main() {
	r := qin.New()
	r.GET("/", func(c *qin.Context) {
		c.String(200, "URL.Path = %q\n", c.Req.Header)
	})

	r.GET("/hello", func(c *qin.Context) {
		for k, v := range c.Req.Header {
			c.String(200, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":8080")
}
