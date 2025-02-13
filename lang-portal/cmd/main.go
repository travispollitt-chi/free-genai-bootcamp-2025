package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// ...other imports...
	"lang-portal/internal/handlers"
	"lang-portal/internal/database"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("words.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize the database
	err = database.InitDB("words.db")
	if err != nil {
		panic(err)
	}

	// Define routes
	r.GET("/api/dashboard/last_study_session", handlers.GetLastStudySession)
	r.GET("/api/dashboard/study_progress", handlers.GetStudyProgress)
	r.GET("/api/dashboard/quick_stats", handlers.GetQuickStats)
	r.GET("/api/study_activities/:id", handlers.GetStudyActivity)
	r.GET("/api/study_activities/:id/study_sessions", handlers.GetStudyActivitySessions)
	r.POST("/api/study_activities", handlers.CreateStudyActivity)
	r.GET("/api/words", handlers.GetWords)
	r.GET("/api/words/:id", handlers.GetWord)
	r.GET("/api/groups", handlers.GetGroups)
	r.GET("/api/groups/:id", handlers.GetGroup)
	r.GET("/api/groups/:id/words", handlers.GetGroupWords)
	r.GET("/api/groups/:id/study_sessions", handlers.GetGroupStudySessions)
	r.GET("/api/study_sessions/:id", handlers.GetStudySession)
	r.GET("/api/study_sessions/:id/words", handlers.GetStudySessionWords)
	r.POST("/api/reset_history", handlers.ResetHistory)
	r.POST("/api/full_reset", handlers.FullReset)
	r.POST("/api/study_sessions/:id/words/:word_id/review", handlers.ReviewWord)

	r.Run() // listen and serve on 0.0.0.0:8080
}
