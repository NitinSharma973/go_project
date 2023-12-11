package main

import "gofr.dev/pkg/gofr"

func main() {
    app := gofr.New()

    app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
        return "Hello nitnFR", nil
    })

    app.Start()
}