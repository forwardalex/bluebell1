package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/isdamir/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{
		"community_id":1,
		"title":"test",
		"content":"test content"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, "服务繁忙", w.Body.String())
	//assert.Equal(t, "pong", w.Body.String())
}
