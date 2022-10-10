//go:build mock

package user_test

import (
	"bytes"
	"demo/user"
	"demo/user/test"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostUserRoute(t *testing.T) {
	router := test.NewTestRouterWithMock()

	accountRequest := user.CreateAccount{
		Firstname: "test firstname",
		Lastname: "test lastname",
	}
	body, _ := json.Marshal(accountRequest)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Mock email")
}