package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achmadsy/go-project/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	main()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/all", nil)
	routes.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	if assert.NotNil(t, w.Body.String()) {
		assert.NotEqual(t, w.Body.String(), "Failed to fetch users", "Error API GetAllUser")
	}

}
