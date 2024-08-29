package main

import (
	"angry-embassies/config"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "angry-embassy",
	Short: "Get embassies for a given home and host country",
	Long:  "This application demonstrates passing parameters through the terminal using Cobra.",
}

var getEmbassies = &cobra.Command{
	Use:   "getembassies",
	Short: "Fetch embassies for a given home and host country",
	Long:  "This command takes two parameters and passes them to the handler function.",
	Run: func(cmd *cobra.Command, args []string) {

		// Retrieve the values of the flags
		home, _ := cmd.Flags().GetString("home")
		host, _ := cmd.Flags().GetString("host")

		fmt.Println("Dependencies initialized")

		deps, err := config.InitDependencies()
		if err != nil {
			fmt.Printf("config.InitDependencies error: %v\n", err)
			return
		}

		test := "Hello, MongoDB!"
		scndTest := "Hello, MongoDB! 2"

		deps.MgoService.InsertDocument(*deps.ApiClient, test, scndTest)
		fmt.Printf("Fetching embassies for Home Country: %s and Host Country: %s\n", home, host)
	},
}

func init() {
	fmt.Println("Initializing Cobra CLI")
	rootCmd.AddCommand(getEmbassies)

	// Register flags
	getEmbassies.Flags().String("home", "", "Home country")
	getEmbassies.Flags().String("host", "", "Host country")

	// Mark flags as required if needed
	getEmbassies.MarkFlagRequired("home")
	getEmbassies.MarkFlagRequired("host")
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}

	//deps, err := config.InitDependencies()
	//if err != nil {
	//	fmt.Printf("config.InitDependencies error: %v\n", err)
	//	return
	//}
	//
	//googleID := deps.MakerClient.GetGoogleID("Spain Embassy in Berlin")
	//fmt.Printf("Google ID: %v\n", googleID)
	//
	//document, err := deps.PersistorClient.InsertDocument("Hello, MongoDB!")
	//if err != nil {
	//	fmt.Printf("deps.PersistorClient.InsertDocument error: %v\n", err)
	//	return
	//}
	//
	fmt.Println("Hello, GoLand!")
}
