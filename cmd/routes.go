package cmd

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/vahidlotfi71/online-store-api/Routes"
)

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Shows defined routes list",
	Long:  `Displays a clean list of all registered routes.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()
		Routes.SetupRoutes(app)

		type routeDTO struct{ Method, Path, Name string }
		list := []routeDTO{}
		for _, r := range app.GetRoutes() {
			if r.Path == "/*" || (r.Method == "HEAD" && strings.Contains(r.Name, "index")) {
				continue
			}
			list = append(list, routeDTO{r.Method, r.Path, r.Name})
		}
		fmt.Println("Method | Path                         | Name")
		for _, r := range list {
			fmt.Printf("%-6s | %-28s | %s\n", r.Method, r.Path, r.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(routesCmd)
}
