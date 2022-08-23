/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/agrimel-0/rio-client/client"
	"github.com/spf13/cobra"
)

// byOffsetCmd represents the byOffset command
var byOffsetCmd = &cobra.Command{
	Use:   "byOffset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		lineOffset, err := cmd.Flags().GetInt("lineOffset")
		if err != nil {
			panic(err)
		}

		state, err := cmd.Flags().GetInt("state")
		if err != nil {
			panic(err)
		}

		client, err := client.Start(&ServerAddress)
		if err != nil {
			panic(err)
		}

		client.SetByOffset(lineOffset, state)

	},
}

func init() {
	setCmd.AddCommand(byOffsetCmd)

	byOffsetCmd.Flags().IntP("lineOffset", "l", 0, "line offset")
	byOffsetCmd.Flags().IntP("state", "s", 0, "state, 1 is high, 0 is low ideally")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byOffsetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byOffsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
