// +build mage

package main

import (
    "fmt"
    "os/exec"
)

func InitDB() error {
    cmd := exec.Command("sqlite3", "words.db", "<", "db/migrations/0001_create_words_table.sql")
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to initialize database: %w", err)
    }
    return nil
}

func MigrateDB() error {
    // Migration logic
    return nil
}

func SeedData() error {
    // Seeding logic
    return nil
}

func Migrate() error {
    cmd := exec.Command("sqlite3", "lang-portal.db", "<", "db/migrations/0005_create_study_activities_table.sql")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func Seed() error {
    cmd := exec.Command("sqlite3", "lang-portal.db", "<", "db/seeds/words.json")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func Build() error {
    fmt.Println("Building the project...")
    cmd := exec.Command("go", "build", "-o", "lang-portal")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
