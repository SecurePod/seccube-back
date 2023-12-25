package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCreate(t *testing.T) {
	r := httptest.NewRequest("POST", "/api/v1/docker/create", nil)
	w := httptest.NewRecorder()
	c := echo.New().NewContext(r, w)

	if err := Create(c); err != nil {
		t.Error(err)
		return
	}
}
