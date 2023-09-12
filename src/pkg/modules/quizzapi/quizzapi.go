package quizzapi

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/Noyhaaa/QuizzGameTwitch/src/pkg/modules/dbutils"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Bienvenue dans le JiJi Quizz",
	})
}

func MyRoomPage(c *gin.Context) {
	c.HTML(http.StatusOK, "myroom.html", gin.H{
		"message": "Bienvenue dans la room de JiJi",
	})
}

func StartGame(c *gin.Context) {
	var answers []string
	questionInfo := dbutils.CurrentQuestionInfo()

	if _, ok := questionInfo["answers"].([]string); ok{ 
		answers = questionInfo["answers"].([]string)
	} else {
		fmt.Println("Can't retrieve the answers of the question")
	}

	c.HTML(http.StatusOK, "game.html", gin.H{
		"question": questionInfo["question"],
		"answer1": answers[0],
		"answer2": answers[1],
		"answer3": answers[2],
		"answer4": answers[3],
		"correctAnswer": questionInfo["correct"],
	})
}

func NextQuestion(c *gin.Context) {
	var answers []string
	questionInfo := dbutils.CurrentQuestionInfo()

	if _, ok := questionInfo["answers"].([]string); ok{ 
		answers = questionInfo["answers"].([]string)
	} else {
		fmt.Println("Can't retrieve the answers of the question")
	}

	c.JSON(http.StatusOK, gin.H{
		"question": questionInfo["question"],
		"answer1": answers[0],
		"answer2": answers[1],
		"answer3": answers[2],
		"answer4": answers[3],
		"correctAnswer": questionInfo["correct"],
	})
}