package main

import (
	"fmt"
	"github.com/Noyhaaa/QuizzGameTwitch/src/pkg/modules/dbutils"
)

func main(){
	dbutils.InitDB()
	
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