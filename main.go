package main

import (
	"fmt"
	"os"

	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Routes"
)

func main() {
	/* These two if sections are so important, as much as if they fail ,
	application must exit */
	if err := Config.Getenv(); err != nil {
		fmt.Printf("Runtime Error: Could not load environment variables : %s\n", err.Error())
		os.Exit(2)
	}
	if err := Config.Connect(); err != nil {
		fmt.Printf("Connection Error : Could not connect to the database\n")
		fmt.Printf("%v\n", err.Error())
		os.Exit(2)
	}
	Routes.SetupRoutes(Config.App)
	cmd.Execute()
}
