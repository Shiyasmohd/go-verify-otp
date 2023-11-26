package main

import (
	"github.com/shiyasmohd/go-verify-otp/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	app := api.Config{Router: router}

	app.Routes()

	router.Run(":8080")
}
