package app

import (
	"github.com/gin-gonic/gin"
)

type GrammarTest struct {
	questions    []Question
	userAnswers  []string
	totalPoints  int
	currentIndex int
}

func (app *GrammarTest) Initialize() {
	app.LoadQuestionsFromCSV("questions.csv")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", app.showQuestion)
	r.GET("/check-answer", app.checkAnswer)
	r.POST("/submit", app.submitAnswer)
}

func (app *GrammarTest) Run(address string) {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", app.showQuestion)
	r.GET("/check-answer", app.checkAnswer)
	r.POST("/submit", app.submitAnswer)

	err := r.Run(address)
	if err != nil {
		return
	}
}
