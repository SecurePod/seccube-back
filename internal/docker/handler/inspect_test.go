package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func TestInspect(t *testing.T) {
	json := `[{"id": "4decdeec567f58"}, {"id": "104783feccf"}]`
	r := httptest.NewRequest("POST", "/api/v1/docker/inspect", strings.NewReader(json))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := echo.New().NewContext(r, w)

	if err := Inspect(c); err != nil {
		t.Error(err)
		return
	}

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal("Failed to read response body:", err)
	}
	log.Debug().Str("responseBody", string(responseBody)).Msg("responseBody")
}
