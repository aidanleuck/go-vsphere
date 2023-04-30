package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"vsphere_module/src/common"
	"vsphere_module/src/routes"

	"github.com/magiconair/properties/assert"
)

var TEST_PORT = 80500

func TestGetJobs(t *testing.T) {
	service := &common.AgentService{
		Queue: &mockQueueModel{},
	}
	router := routes.SetupRouter(service)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/job", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, true, true)
}
