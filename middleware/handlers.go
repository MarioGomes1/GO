package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mariogomes1/GO/models"
)

//response
type response struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func CreateConnection() *sql.DB {
	//this loads the .env
	err := godotenv.Load(".env")
	//Fatal is equivalent to Print() followed by a call to os.Exit(1).
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//open connection to the db
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sucessfully connected")
	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//creates an empty user of type models.User
	var user models.User

	// decode json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertUser(user)

	res := response{
		ID:      insertID,
		Message: "User created sucessfully",
	}

	json.NewEncoder(w).Encode(res)

}

// ----------------------------------Handler funcs----------------------------------------

func insertUser(user models.User) int64 {
	db := CreateConnection()

	// close db connection
	defer db.Close()
	//"returning id" returns the id of the inserted user
	sqlStatement := `INSERT INTO users (name, location, age)
					VALUES($1, $2, $3) RETURNING userid`

	var id int64

	//Scan func saves the insert/returned id in the id var
	err := db.QueryRow(sqlStatement,
		user.Name,
		user.Location,
		user.Age).Scan(id)

	if err != nil {
		log.Fatalf("Unable to exec query, %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}
