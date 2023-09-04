package main

import (
	"fmt"
	"github.com/Noyhaaa/QuizzGameTwitch/src/pkg/modules/dbutils"
)

func main(){
	dbutils.InitDB()
	
	fmt.Println(dbutils.CurrentQuestionInfo())

}