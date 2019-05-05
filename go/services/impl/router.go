package impl

import (
	// Go native packages
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// Dep packages
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	contact := router.Group("/api/v0/contacts")
	{
		contact.GET(":contact-id", GetContact)
		contact.DELETE(":contact-id", DeleteContact)
		contact.GET("", ListContact)
		contact.POST("", CreateContact)
	}

	return router
}

func performRequest(t *testing.T, r http.Handler, method, path string,
	body interface{}) *httptest.ResponseRecorder {
	if body == nil {
		req, _ := http.NewRequest(method, path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w
	} else {
		j, err := json.Marshal(body)
		assert.NoError(t, err, "NO ERROR!")
		req, _ := http.NewRequest(method, path, strings.NewReader(string(j)))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w
	}
}
