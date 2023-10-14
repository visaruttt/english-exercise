package app

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *GrammarTest) showQuestion(c *gin.Context) {
	if app.currentIndex >= len(app.questions) {
		c.HTML(http.StatusOK, "completed.html", gin.H{
			"TotalQuestions": app.currentIndex,
			"TotalPoints":    app.totalPoints,
			"Percentage":     float64(app.totalPoints) / float64(len(app.questions)) * 100,
			"UserAnswers":    app.userAnswers,
			"CorrectAnswers": getCorrectAnswers(app.questions),
			"Feedback":       getFeedback(app.questions),
		})
		return
	}
	question := app.questions[app.currentIndex]

	c.HTML(http.StatusOK, "question.html", gin.H{
		"ID":           question.ID,
		"Question":     question.Question,
		"Options":      question.Options,
		"UserAnswer":   "", // Empty string for user's answer before submission
		"IsCorrect":    false,
		"Feedback":     "",
		"DisplayCheck": false,
	})
}

func (app *GrammarTest) submitAnswer(c *gin.Context) {
	userAnswer := c.PostForm("answer")
	questionIDStr := c.PostForm("question_id")
	fmt.Println("questionID", questionIDStr)
	questionID, _ := strconv.ParseInt(questionIDStr, 10, 0)

	question := &app.questions[questionID-1]
	app.currentIndex = question.ID
	if userAnswer == question.Answer {
		app.totalPoints++
		question.IsCorrect = true
	} else {
		question.IsCorrect = false
	}

	app.userAnswers = append(app.userAnswers, userAnswer)

	if app.currentIndex < len(app.questions) {
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		app.showResults(c)
	}
}

func (app *GrammarTest) showResults(c *gin.Context) {
	totalQuestions := len(app.questions)
	percentage := float64(app.totalPoints) / float64(totalQuestions) * 100

	correctAnswers := getCorrectAnswers(app.questions)
	feedback := getFeedback(app.questions)

	var resultData []struct {
		Question      string
		UserAnswer    string
		CorrectAnswer string
		Feedback      string
		IsCorrect     bool
	}

	for i, userAnswer := range app.userAnswers {
		resultData = append(resultData, struct {
			Question      string
			UserAnswer    string
			CorrectAnswer string
			Feedback      string
			IsCorrect     bool
		}{
			Question:      app.questions[i].Question,
			UserAnswer:    userAnswer,
			CorrectAnswer: correctAnswers[i],
			Feedback:      feedback[i],
			IsCorrect:     userAnswer == correctAnswers[i],
		})
	}

	data := struct {
		TotalQuestions int
		TotalPoints    int
		Percentage     float64
		Results        []struct {
			Question      string
			UserAnswer    string
			CorrectAnswer string
			Feedback      string
			IsCorrect     bool
		}
	}{
		TotalQuestions: totalQuestions,
		TotalPoints:    app.totalPoints,
		Percentage:     percentage,
		Results:        resultData,
	}

	tmpl := template.Must(template.ParseFiles("templates/completed.html"))
	err := tmpl.Execute(c.Writer, data)
	if err != nil {
		return
	}
}

func getCorrectAnswers(questions []Question) []string {
	correctAnswers := make([]string, len(questions))
	for i, question := range questions {
		correctAnswers[i] = question.Answer
	}
	return correctAnswers
}

func getFeedback(questions []Question) []string {
	feedback := make([]string, len(questions))
	for i, question := range questions {
		feedback[i] = question.Feedback
	}
	return feedback
}
