package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vahidlotfi71/online-store-api.git/Config"
	"github.com/vahidlotfi71/online-store-api.git/Routes"
)

var port string

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "server port")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start development server",
	Run: func(cmd *cobra.Command, args []string) {
		// 1) خواندن env
		if err := Config.Getenv(); err != nil {
			log.Fatalf("env load: %v", err)
		}
		// 2) اتصال DB
		err := Config.Connect()
		if err != nil {
			log.Fatalf("db connect: %v", err)
		}
		// 3) ساخت Fiber
		app := Config.App
		// 4) ثبت مسیرها
		Routes.SetupRoutes(app)
		// 5) پورت
		if port == "" {
			port = os.Getenv("PORT")
		}
		if port == "" {
			port = "8080"
		}
		log.Printf("🚀 Server on :%s", port)
		log.Fatal(app.Listen(":" + port))
	},
}
