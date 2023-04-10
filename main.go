package main

import (
	"fmt"
	"log"
	"os"

	"github.com/katsu105/clean-todo/driver"
)

func main() {
	// 環境変数を読み込む
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(".env not found")
	// }

	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
