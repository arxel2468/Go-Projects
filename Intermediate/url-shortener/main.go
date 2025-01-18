package main

import (
    "github.com/gin-gonic/gin"
    "url-shortener/handlers"
)

func main() {
    router := gin.Default()

    router.POST("/shorten", handlers.ShortenURL)
    router.GET("/:shortURL", handlers.RedirectURL)

    router.Run(":8080")
}
