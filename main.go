package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/Noyhaaa/QuizzGameTwitch/src/pkg/modules/dbutils"
	"github.com/Noyhaaa/QuizzGameTwitch/src/pkg/modules/quizzapi"
)

func main(){
	dbutils.InitDB()
	router := gin.Default()
	router.LoadHTMLGlob("src/web/templates/*")
	router.Static("/static/css", "src/web/static/css")

	router.GET("/", quizzapi.HomePage)

	router.GET("/myroom", quizzapi.MyRoomPage)

	router.GET("/game", quizzapi.StartGame)

	// Start web server on port 8080
    router.Run(":8080")
	
	fmt.Println(dbutils.CurrentQuestionInfo())
	//create new question info
	/*
	answers := []string{"L'oxygène (O2)", "Le dioxyde de carbone (CO2)", "L'argon (Ar)", "L'azote (N2)"}
	dbutils.AddQuestionAnswers(
		"Quel est l'élément chimique le plus abondant dans l'atmosphère terrestre ?",
		"Moyen",
		"Sciences",
		answers,
		"L'azote (N2)",
	)
	*/
}