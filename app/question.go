package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Question struct {
	ID        int
	Question  string
	Options   []string
	Answer    string
	IsCorrect bool
	Feedback  string
}

func (app *GrammarTest) checkAnswer(c *gin.Context) {
	questionIDStr := c.Query("question_id")
	selectedOption := c.Query("selected_option")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	userAnswer := c.Query("user_answer")
	if userAnswer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User answer is required"})
		return
	}

	var question Question
	found := false
	for _, q := range app.questions {
		if q.ID == questionID {
			question = q
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if userAnswer == question.Answer {
		question.IsCorrect = true
	} else {
		question.IsCorrect = false
	}

	c.HTML(http.StatusOK, "question.html", gin.H{
		"ID":             question.ID,
		"Question":       question.Question,
		"Options":        question.Options,
		"UserAnswer":     userAnswer,
		"IsCorrect":      question.IsCorrect,
		"Feedback":       question.Feedback,
		"DisplayCheck":   true,
		"SelectedOption": selectedOption,
		"CorrectAnswers": getCorrectAnswers(app.questions),
	})
}
