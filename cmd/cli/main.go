package main

import (
	config "configtest"
	"fmt"
	"github.com/spf13/cobra"
	"models"
)

type cliCommands struct {
	writeEmbassy *cobra.Command
	getEmbassy   *cobra.Command
}

var rootCmd = &cobra.Command{
	Use:   "angry-embassy",
	Short: "Get embassies for a given home and host country",
	Long:  "This application demonstrates passing parameters through the terminal using Cobra.",
}

// TODO: make this return a struct with all commands
func registerCommands(deps *config.Dependencies) (cmds cliCommands) {
	var writeEmbassy = &cobra.Command{
		Use:   "writeembassy",
		Short: "Write embassies to the database",
		Long:  "This command takes two parameters and passes them to the handler function.",
		//	TODO: Make new package for cli that starts on a Run method
	}
	cmds.writeEmbassy = writeEmbassy

	var getEmbassy = &cobra.Command{
		Use:   "getembassy",
		Short: "Retrieve embassy from the database",
		Long:  "This command retrieves an embassy from the database",
	}

	cmds.getEmbassy = getEmbassy

	getEmbassy.Run = func(cmd *cobra.Command, args []string) {
		// Retrieve the values of the flags
		home, _ := cmd.Flags().GetString("home")
		host, _ := cmd.Flags().GetString("host")
		city, _ := cmd.Flags().GetString("city")

		deps.MgoService.GetDocument(models.Embassy{
			HomeCountry: home,
			HostCountry: host,
			City:        city,
		})
	}

	writeEmbassy.Run = func(cmd *cobra.Command, args []string) {

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
	// TODO: decouple initialization from the cli package so that it can be used in the web server as well
	// Initialize dependencies in the main function
	deps, err := config.InitDependencies()
	if err != nil {
		fmt.Printf("config.InitDependencies error: %v\n", err)
		return
	}

	cmds := registerCommands(&deps)

	fmt.Println("Initializing Cobra CLI")
	rootCmd.AddCommand(cmds.writeEmbassy)

	// Register flags
	cmds.writeEmbassy.Flags().String("home", "", "Home country")
	cmds.writeEmbassy.Flags().String("host", "", "Host country")
	cmds.writeEmbassy.Flags().Bool("consulate", false, "Is this a consulate?")
	cmds.writeEmbassy.Flags().String("city", "", "City")

	// Mark flags as required if needed
	cmds.writeEmbassy.MarkFlagRequired("home")
	cmds.writeEmbassy.MarkFlagRequired("host")
	cmds.writeEmbassy.MarkFlagRequired("consulate")
	cmds.writeEmbassy.MarkFlagRequired("city")

	rootCmd.AddCommand(cmds.getEmbassy)
	cmds.getEmbassy.Flags().String("home", "", "Home country")
	cmds.getEmbassy.Flags().String("host", "", "Host country")
	cmds.getEmbassy.Flags().String("city", "", "City")
}

func main() {

	// TODO: How to run the CLI in parallel with the web server?
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello, GoLand! Now with the latest Docker!")
}
