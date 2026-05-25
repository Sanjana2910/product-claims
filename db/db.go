package db

import (
    "database/sql" // lib for packaging to communicate with the sql database
    "fmt" // used for string building
    "log" //print
    "os" // reads the env variables

    _ "github.com/lib/pq"// translates go to postgres
)

var Conn *sql.DB //Conn- global variable, database connection

func Connect() {
	// read variables of env file
    host     := getEnv("DB_HOST", "localhost")
    port     := getEnv("DB_PORT", "5432")
    user     := getEnv("DB_USER", "postgres")
    password := getEnv("DB_PASSWORD", "postgres")
    name     := getEnv("DB_NAME", "product_claims")

    dsn := fmt.Sprintf( //string address of the database
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, name,
    )

    var err error // error variable 
    Conn, err = sql.Open("postgres", dsn) // prepapes the database connection, but does not actually connect
    if err != nil {
        log.Fatalf("could not open db: %v", err)
    }

    if err = Conn.Ping(); err != nil {
        log.Fatalf("could not reach db: %v", err)
    }

    log.Println("postgres connected!")
}
// if env var is missing, use the fallback value
func getEnv(key, fallback string) string {
    val := os.Getenv(key)// reads env var
    if val == "" {
        return fallback
    }
    return val
}