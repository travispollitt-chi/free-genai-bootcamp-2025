package routes

import (
	"github.com/gin-gonic/gin"
	"your_project/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/last-study-session", handlers.GetLastStudySession)
	r.GET("/study-progress", handlers.GetStudyProgress)
	r.GET("/quick-stats", handlers.GetQuickStats)
	r.GET("/study-activity", handlers.GetStudyActivity)
	// ...add routes for other handler functions...

	return r
}
