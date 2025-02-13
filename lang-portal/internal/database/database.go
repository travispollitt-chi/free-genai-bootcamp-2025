package database

import (
	"database/sql"
	// ...other imports...
)

var DB *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func GetLastStudySession() (*models.StudySession, error) {
	// ...implementation...
}

func GetStudyProgress() (*models.StudyProgress, error) {
	// ...implementation...
}

func GetQuickStats() (*models.QuickStats, error) {
	// ...implementation...
}

// ...other database functions...
