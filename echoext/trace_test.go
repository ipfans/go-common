package echoext

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestExtStdContext(t *testing.T) {
	e := echo.New()
	hw := []byte("Hello, World!")
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(hw))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, string(body))
	}
	assert := assert.New(t)
	if assert.NoError(ExtStdContext(nil)(h)(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.NotEmpty(rec.Header().Get("X-Request-ID"))
		assert.Equal(c.Request().Context().Value("logger_request_id"), rec.Header().Get("X-Request-ID"))
	}
	c = e.NewContext(req, rec)
	ut := func(c echo.Context) interface{} {
		return "111"
	}
	if assert.NoError(ExtStdContext(ut)(h)(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.NotEmpty(rec.Header().Get("X-Request-ID"))
		assert.Equal(c.Request().Context().Value("logger_request_id"), rec.Header().Get("X-Request-ID"))
		assert.Equal(c.Request().Context().Value("logger_unique_id"), "111")
	}
}
