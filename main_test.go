package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	router = gin.New()
	register(router)
}

func GET(path string, accept string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Set("Accept", accept)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	return w
}

func TestFurigana(t *testing.T) {
	w := GET("/furigana/自由ヶ丘", "text/plain")

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "text/plain")

	assert.Equal(t, "ジユーガオカ", w.Body.String())
}
