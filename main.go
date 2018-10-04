package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// paramWord picks the word argument.
//
// The parameter should be declared as *word.
//
// (This function is stolen from api.hangulize.org)
//
func paramWord(c *gin.Context) string {
	word := c.Param("word")

	// Remove the initial slash.
	word = word[1:]

	unescaped, err := url.QueryUnescape(word)
	if err == nil {
		word = unescaped
	}

	return word
}

func handler(c *gin.Context) {
	pID := c.Param("phonemizer")
	word := paramWord(c)

	p, ok := phonemizers[pID]

	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	phonemized := p.Phonemize(word)
	c.String(http.StatusOK, phonemized)
}

func register(r gin.IRouter) {
	r.GET("/:phonemizer/*word", handler)
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
