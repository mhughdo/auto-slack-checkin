/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	slack "auto-slack-checkin/internal/pkg/slack"

	"github.com/spf13/cobra"
)

// runCronCmd represents the runCron command
var runCronCmd = &cobra.Command{
	Use:   "runCron",
	Short: "Run a cron job.",
	Long:  `Run a cron job to send the message.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := slack.SendMessage()
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(runCronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
