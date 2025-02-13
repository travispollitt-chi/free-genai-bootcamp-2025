package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetLastStudySession(t *testing.T) {
	router := gin.Default()
	router.GET("/last-study-session", GetLastStudySession)

	req, _ := http.NewRequest("GET", "/last-study-session", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on expected response
}

func TestGetStudyProgress(t *testing.T) {
	router := gin.Default()
	router.GET("/study-progress", GetStudyProgress)

	req, _ := http.NewRequest("GET", "/study-progress", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on expected response
}

func TestGetQuickStats(t *testing.T) {
	router := gin.Default()
	router.GET("/quick-stats", GetQuickStats)

	req, _ := http.NewRequest("GET", "/quick-stats", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on expected response
}

func TestGetStudyActivity(t *testing.T) {
	router := gin.Default()
	router.GET("/study-activity", GetStudyActivity)

	req, _ := http.NewRequest("GET", "/study-activity", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on expected response
}
