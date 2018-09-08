package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	pID := c.Param("phonemizer")
	word := c.Param("word")

	p, ok := phonemizers[pID]

	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	phonemized := p.Phonemize(word)
	c.String(http.StatusOK, phonemized)
}

func register(r gin.IRouter) {
	r.GET("/:phonemizer/:word", handler)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	register(router)
	router.Run(":" + port)
}
