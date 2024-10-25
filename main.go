package main

import (
	"angry_embassies/config"
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

		// TODO: Initialize dependencies in the main function
		deps, err := config.InitDependencies()
		if err != nil {
			fmt.Printf("config.InitDependencies error: %v\n", err)
			return
		}

		//TODO: configs get google service to get embassy details
		//TODO: how can I test it , or test different configs
		emb, err := deps.GglService.GetEmbassyDetails(home, host)
		if err != nil {
			fmt.Printf("deps.GglService.GetEmbassy error: %v\n", err)
			return
		}

		//pass emb to insertDocument
		deps.MgoService.InsertDocument(emb)
		fmt.Printf("Fetching embassies for Home Country: %s and Host Country: %s\n", home, host)
	},
	//	TODO: Make new package for cli that starts on a Run method
}

func init() {
	// TODO: wrap this in a Setup function; no need for using init()
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

	// TODO: How to run the CLI in parallel with the web server?
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello, GoLand! Now with the latest Docker!")
}
