package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	version = "dev" // is set during build process

	// Default values
	defaultDebug                 = os.Getenv("DEBUG") == "1"
	defaultLogProd               = os.Getenv("LOG_PROD") == "1"

	// Flags
	debugPtr                 = flag.Bool("debug", defaultDebug, "print debug output")
	logProdPtr               = flag.Bool("log-prod", defaultLogProd, "log in production mode (json)")
)

func main() {
	flag.Parse()
	fmt.Println("Hello, World!")
	fmt.Println("Version:", version)
	fmt.Println("Debug:", *debugPtr)
	fmt.Println("LogProd:", *logProdPtr)

	// 환경설정값 읽기
	args := os.Args[1:]
	err := godotenv.Load(args[0] + ".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	connectionsClosed := make(chan struct{})
	go func() {
		close(connectionsClosed)
		fmt.Println("Connections closed")
	}()

	<-connectionsClosed
}

func getConfig() {
	
}
