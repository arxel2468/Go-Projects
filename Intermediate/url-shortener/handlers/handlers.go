package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "url-shortener/storage"
    "url-shortener/utils"
)

func ShortenURL(c *gin.Context) {
    var request struct {
        URL string `json:"url" binding:"required"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    shortCode := utils.GenerateShortCode(6)
    storage.SaveURL(shortCode, request.URL)

    c.JSON(http.StatusOK, gin.H{
        "short_url": "http://localhost:8080/" + shortCode,
    })
}

func RedirectURL(c *gin.Context) {
    shortCode := c.Param("shortURL")
    originalURL, exists := storage.GetURL(shortCode)

    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    c.Redirect(http.StatusMovedPermanently, originalURL)
}
