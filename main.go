package main

import (
	"fmt"
	"os"

	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Routes"
	"github.com/vahidlotfi71/online-store-api.git/cmd"
)

func main() {
	// 1️⃣ load env
	if err := Config.Getenv(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime Error: Could not load environment variables : %s\n", err.Error())
		os.Exit(2)
	}

	// 2️⃣ connect DB
	if err := Config.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "Connection Error : Could not connect to the database\n%v\n", err.Error())
		os.Exit(2)
	}

	// 3️⃣ build Fiber (سراسری در Config آماده است)
	app := Config.App

	// 4️⃣ register routes
	Routes.SetupRoutes(app)

	// 5️⃣ اجرای CLI
	cmd.Execute()
}
