// main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vahidlotfi71/online-store-api/Config"
	"github.com/vahidlotfi71/online-store-api/Routes"
	"github.com/vahidlotfi71/online-store-api/cmd"
)

func main() {
	// اگر کاربر دستور CLI داده، فقط CLI اجرا شود
	if len(os.Args) > 1 && isCLICommand(os.Args[1]) {
		executeCLI()
		return
	}

	// در غیر این صورت سرور اجرا شود
	startServer()
}

// تشخیص دستورات CLI
func isCLICommand(arg string) bool {
	cliCommands := []string{"serve", "migrate", "seed", "routes", "version", "help", "--help", "-h"}
	for _, cmd := range cliCommands {
		if strings.ToLower(arg) == cmd {
			return true
		}
	}
	return false
}

// اجرای CLI
func executeCLI() {
	if err := Config.Getenv(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading env: %s\n", err.Error())
		os.Exit(1)
	}

	if err := Config.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to DB: %s\n", err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}

// راه‌اندازی سرور
func startServer() {
	// load env
	if err := Config.Getenv(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime Error: Could not load environment variables : %s\n", err.Error())
		os.Exit(2)
	}

	//connect DB
	if err := Config.Connect(); err != nil {
		fmt.Fprintf(os.Stderr, "Connection Error : Could not connect to the database\n%v\n", err.Error())
		os.Exit(2)
	}

	//build Fiber
	app := Config.App

	//register routes
	Routes.SetupRoutes(app)

	//start server
	port := getPort()
	fmt.Printf("Server starting on :%s\n", port)
	if err := app.Listen(":" + port); err != nil {
		fmt.Fprintf(os.Stderr, "Server Error: %s\n", err.Error())
		os.Exit(2)
	}
}

// دریافت پورت از env یا پیش‌فرض
func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}
