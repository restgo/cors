package main

import (
	"github.com/restgo/restgo"
	"github.com/restgo/csrf"
)

func main() {
	app := restgo.App()

	app.Use(cors.New(cors.Options{
		Debug:true,
		AllowedOrigins: []string{"*"},
	}))

	app.GET("/users", func(ctx *restgo.Context, next restgo.Next) {
		ctx.ServeText(200, "users")
	})

	app.Run()
}