// package main

// import (
//     "log"
//     "product-claims/db"

//     "github.com/joho/godotenv"// to read the .env file and load the env variables into the system environment
// )

// func main() {
// 	// load the .env file 1st so that the env variables are available for the db connection

//     err := godotenv.Load()
//     if err != nil {
//         log.Println("no .env file, using system env")
//     }
// 	// connect to the postgres db

//     db.Connect()

//     log.Println("server starting on port 8080...")
// }
package main

import (
	"log"
	"net/http"
	"product-claims/db"
	"product-claims/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found, using system environment")
	}

	// connect to database
	db.Connect()

	// register our two API routes
	http.HandleFunc("/claims/identify", handlers.IdentifyClaims)
	http.HandleFunc("/claims/status/", handlers.GetClaimStatus)

	log.Println("server running on http://localhost:8080")

	// start the server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}