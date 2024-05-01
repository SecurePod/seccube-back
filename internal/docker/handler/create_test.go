package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreate(t *testing.T) {
	e := echo.New()
	e.POST("/api/v1/docker/create/:tag", Create)

	r := httptest.NewRequest("POST", "/api/v1/docker/create/sshBrute", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	if err := Create(c); err != nil {
		t.Error(err)
		return
	}
}
