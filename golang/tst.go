package main

import (
  "fmt"
  "os"
  "log"
  "github.com/joho/godotenv"
)

func main () {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }
  fmt.Println(os.Getenv("PSQL_USER"))
}
