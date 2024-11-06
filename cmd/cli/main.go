package main

import (
	config "configtest"
	"fmt"
	"github.com/spf13/cobra"
	"models"
)

type cliCommands struct {
	getEmbassies *cobra.Command
	test         *cobra.Command
}

var rootCmd = &cobra.Command{
	Use:   "angry-embassy",
	Short: "Get embassies for a given home and host country",
	Long:  "This application demonstrates passing parameters through the terminal using Cobra.",
}

// TODO: make this return a struct with all commands
func registerCommands(deps *config.Dependencies) (cmds cliCommands) {
	var getEmbassies = &cobra.Command{
		Use:   "getembassies",
		Short: "Fetch embassies for a given home and host country",
		Long:  "This command takes two parameters and passes them to the handler function.",
		//	TODO: Make new package for cli that starts on a Run method
	}
	cmds.getEmbassies = getEmbassies
	cmds.test = &cobra.Command{
		Use: "test",
	}

	getEmbassies.Run = func(cmd *cobra.Command, args []string) {

		// Retrieve the values of the flags
		home, _ := cmd.Flags().GetString("home")
		host, _ := cmd.Flags().GetString("host")
		consulate, _ := cmd.Flags().GetBool("consulate")
		city, _ := cmd.Flags().GetString("city")

		embassy := *models.NewEmbassy(home, host, consulate, "", city, "", models.PlaceDetails{})

		fmt.Println("Dependencies initialized")
		//TODO: configs get google service to get embassy details
		//TODO: how can I test it , or test different configs
		emb, err := deps.GglService.GetEmbassyDetails(embassy)
		if err != nil {
			fmt.Printf("deps.GglService.GetEmbassy error: %v\n", err)
			return
		}

		//pass emb to insertDocument
		deps.MgoService.InsertDocument(emb)
		fmt.Printf("Fetching embassies for Home Country: %s and Host Country: %s\n", home, host)

	}
	return cmds
}

func init() {
	// TODO: wrap this in a Setup function; no need for using init()
	// Initialize dependencies in the main function
	deps, err := config.InitDependencies()
	if err != nil {
		fmt.Printf("config.InitDependencies error: %v\n", err)
		return
	}

	cmds := registerCommands(&deps)

	fmt.Println("Initializing Cobra CLI")
	rootCmd.AddCommand(cmds.getEmbassies)

	// Register flags
	cmds.getEmbassies.Flags().String("home", "", "Home country")
	cmds.getEmbassies.Flags().String("host", "", "Host country")
	cmds.getEmbassies.Flags().Bool("consulate", false, "Is this a consulate?")
	cmds.getEmbassies.Flags().String("city", "", "City")

	// Mark flags as required if needed
	cmds.getEmbassies.MarkFlagRequired("home")
	cmds.getEmbassies.MarkFlagRequired("host")
	cmds.getEmbassies.MarkFlagRequired("consulate")
	cmds.getEmbassies.MarkFlagRequired("city")
}

func main() {

	// TODO: How to run the CLI in parallel with the web server?
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello, GoLand! Now with the latest Docker!")
}
