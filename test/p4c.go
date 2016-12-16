package main

import (
	"fmt"

	"github.com/insionng/macross"
	"github.com/insionng/macross/logger"
	"github.com/insionng/macross/pongor"
	"github.com/insionng/macross/recover"
	"github.com/macross-contrib/cache"
	"github.com/macross-contrib/captcha"
)

func main() {
	v := macross.New()
	v.Use(logger.Logger())
	v.Use(recover.Recover())
	v.Use(cache.Cacher(cache.Options{Adapter: "memory"}))
	v.Use(captcha.Captchaer())
	v.SetRenderer(pongor.Renderor())

	v.Get("/", func(self *macross.Context) error {
		if cpt := self.Get("Captcha"); cpt != nil {
			fmt.Println("Got:", cpt)
		} else {
			fmt.Println("Captcha is nil!")
		}

		self.Set("title", "你好，世界")
		return self.Render("index")
	})

	v.Listen(":7891")
}
