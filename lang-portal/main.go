package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "lang-portal/models"   "lang-portal/models"
    "lang-portal/services"    "gorm.io/driver/sqlite"
    "gorm.io/driver/sqlite"gorm"
    "gorm.io/gorm"
)

var db *gorm.DB
var wordService *services.WordServicenit() {
    var err error
func init() {l.db"), &gorm.Config{})
    var err error
    db, err = gorm.Open(sqlite.Open("lang-portal.db"), &gorm.Config{})  panic("failed to connect database")
    if err != nil {    }
        panic("failed to connect database")Migrate(&models.Word{})
    }
    db.AutoMigrate(&models.Word{})


































}    c.JSON(http.StatusOK, word)    }        return        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})    if err := wordService.CreateWord(&word); err != nil {    }        return        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})    if err := c.ShouldBindJSON(&word); err != nil {    var word models.Wordfunc createWord(c *gin.Context) {}    c.JSON(http.StatusOK, words)    }        return        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})    if err != nil {    words, err := wordService.GetAllWords()func getWords(c *gin.Context) {}    r.Run()    r.POST("/api/words", createWord)    r.GET("/api/words", getWords)    r := gin.Default()func main() {}    wordService = &services.WordService{DB: db}func main() {
    r := gin.Default()

    r.GET("/api/words", getWords)
    r.POST("/api/words", createWord)

    r.Run()
}

func getWords(c *gin.Context) {
    var words []models.Word
    db.Find(&words)
    c.JSON(http.StatusOK, words)
}

func createWord(c *gin.Context) {
    var word models.Word
    if err := c.ShouldBindJSON(&word); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&word)
    c.JSON(http.StatusOK, word)
}
