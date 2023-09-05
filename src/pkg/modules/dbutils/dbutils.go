package dbutils

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Questions struct {
	Questionid int64
	Question   string
	Difficulty string
	Theme      string
	Remanence  int64
}

type Answers struct {
	Answerid        int64
	Questionid      int64
	Question_answer string
	Valid           string
}

func InitDB() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Jiji_quizz",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Print("go error when connecting to the DB")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func DataBaseRowsQuery(cmd string, args ...interface{}) *sql.Rows {
	/*
		Allow the user to query with no or several parameters

		Parameters
		---------------
		cmd: string
			Query to use

		args: interface
			all of the other parameters given to the DataBaseRowsQuery call.

		Returns
		---------------
		*sql.Rows:
			rows pointer get from the database result
	*/
	if len(args) == 0 {
		rows, err := db.Query(cmd)
		if err != nil {
			log.Printf("an occured with the command: %s", cmd)
		}
		return rows
	} else {
		rows, err := db.Query(cmd, args...)
		if err != nil {
			log.Printf("an error occured with the command: %s with followign argument %v", cmd, args)
		}
		return rows
	}
}

func CurrentQuestionInfo() map[string]interface{}{
	/*
		Get all the information usefull to display on a screen

		Returns
		---------------
		map[string]interface
			return map with seevral information around the same ID.
			Informtion into the map:
				- question
				- answers of the question
				- theme of the question
				- difficulty of the question
				- correct answer of the question
	*/
	var questionInfo = make(map[string]interface{})
	// Get a ramdom ID
	questionID := RandQuestion()

	// Get all the information linked to this ID
	questionInfo["question"] = GetQuestion(questionID)
	questionInfo["answers"] = GetAnswers(questionID)
	questionInfo["theme"] = GetQuestionTheme(questionID)
	questionInfo["difficulty"] = GetQuestionDifficulty(questionID)
	questionInfo["correct"] = GetQuestionCorrectAnswer(questionID)

	return questionInfo
}

func RandQuestion() int{
	/*
		Get a random int correspong to an question ID

		Returns
		---------------
		int
			Random ID among the available questions list
	*/
	var highestID int
	cmd := "SELECT questionid FROM questions ORDER BY questionid DESC LIMIT 1;"
	row := db.QueryRow(cmd)
	if err := row.Scan(&highestID); err != nil{
		log.Printf("an occured with the command: %s", cmd)
	}
	return rand.Intn(highestID) + 1
}

func GetQuestion(id int) string{
	/*
		Get a question from the database

		Parameters
		---------------
		id: int
			id corresponding to a question

		Returns
		---------------
		string
			Return the question from the ID associated
	*/
	var questions Questions
	cmd := fmt.Sprintf("SELECT question FROM questions WHERE questionid = %d;", id)
	row := db.QueryRow(cmd)
	if err := row.Scan(&questions.Question); err != nil{
		log.Printf("an occured with the command: %s", cmd)
	}
	return questions.Question
}

func GetQuestionTheme(id int) string{
	/*
		Get the theme of a questionid from the database

		Parameters
		---------------
		id: int
			id corresponding to a question

		Returns
		---------------
		string
			Return the theme of the question from the ID associated
	*/
	var questions Questions
	cmd := fmt.Sprintf("SELECT Theme FROM questions WHERE questionid = %d;", id)
	row := db.QueryRow(cmd)
	if err := row.Scan(&questions.Theme); err != nil{
		log.Printf("an occured with the command: %s", cmd)
	}
	return questions.Theme
}

func GetQuestionDifficulty(id int) string{
	/*
		Get the difficulty of a questionid from the database

		Parameters
		---------------
		id: int
			id corresponding to a question

		Returns
		---------------
		string
			Return the difficulty of the question from the ID associated
	*/
	var questions Questions
	cmd := fmt.Sprintf("SELECT difficulty FROM questions WHERE questionid = %d;", id)
	row := db.QueryRow(cmd)
	if err := row.Scan(&questions.Difficulty); err != nil{
		log.Printf("an occured with the command: %s", cmd)
	}
	return questions.Difficulty
}

func GetAnswers(id int) []string{
	/*
		Get an answer from the database

		Parameters
		---------------
		id: int
			id corresponding to an answer

		Returns
		---------------
		[]string
			Return the answers from the ID associated
	*/
	var questions Questions
	var answers []string

	cmd := "SELECT question_answer FROM answers WHERE questionid = ?;"
	rows := DataBaseRowsQuery(cmd, id)
	defer rows.Close()

	for rows.Next(){
		if err := rows.Scan(&questions.Question); err != nil{
			log.Printf("an occured with the command: %s", cmd)
		}
		answers = append(answers, questions.Question)
	}
	return answers
}

func GetQuestionCorrectAnswer(id int) string{
	/*
		Get the correct answer of a question from the database

		Parameters
		---------------
		id: int
			id corresponding to a question

		Returns
		---------------
		string
			Return the correct answer of the question from the ID associated
	*/
	var answers Answers
	cmd := fmt.Sprintf("SELECT question_answer FROM answers WHERE questionid = %d AND Valid = 'Vrai'; ", id)
	row := db.QueryRow(cmd)
	if err := row.Scan(&answers.Question_answer); err != nil{
		log.Printf("an occured with the command: %s", cmd)
	}
	return answers.Question_answer
}

func AddQuestionAnswers(question string, difficulty string, theme string, answers []string, correctAnswer string){
	//Insert new Question
	cmd := `
		INSERT INTO questions 
		(question, difficulty, theme, remanence) 
		VALUES (?, ?, ?, 1);
	`

	_, err := db.Exec(cmd, question, difficulty, theme)
	if err != nil {
		log.Fatal(err)
	}

	// Last ID created is the reference ID
	cmd = "SET @question_id = LAST_INSERT_ID();"
	_, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	var validAnswer string
	for _, answer := range answers{
		if correctAnswer != answer{
			validAnswer = "Faux"
		} else{
			validAnswer = "Vrai"
		}
		cmd := `
			INSERT INTO answers 
			(questionid, question_answer, Valid) 
			VAlUES (@question_id, ?, ?);
		`

		_, err := db.Exec(cmd, answer, validAnswer)
		if err != nil {
			log.Fatal(err)
		}
	}
}