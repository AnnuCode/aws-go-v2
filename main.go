package main

import (
	"log"
	"net/http"

	"github.com/AnnuCode/devops-practice/utils"
)

var (
	dynamodbClient = utils.CreateLocalClient()
	tableName      = "Movies"
	tableBasics    = utils.TableBasics{TableName: tableName,
		DynamoDbClient: dynamodbClient}
)

func main() {

	exists, err := tableBasics.TableExists()
	if err != nil {
		panic(err)
	}

	if !exists {
		log.Printf("Creating table %v...\n", tableName)
		_, err = tableBasics.CreateMovieTable()
		if err != nil {
			panic(err)
		} else {
			log.Printf("Created table %v.\n", tableName)
		}
		written, _ := tableBasics.AddMovieBatch(utils.Movies, 3)
		log.Printf("%d movies were added to the database", written)
	} else {
		log.Printf("Table %v already exists with movies.\n", tableName)
	}

	movie, err := tableBasics.GetMovie("First movie", 2001)
	if err == nil {
		log.Println(movie)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/secret", secretHandler)
	log.Println("Starting server on :5000")
	err = http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
