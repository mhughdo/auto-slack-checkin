/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	slack "auto-slack-checkin/internal/pkg/slack"

	"github.com/spf13/cobra"
)

// cronCmd represents the runCron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Run a cron job.",
	Long:  `Run a cron job to send the message.`,
	Run: func(cmd *cobra.Command, args []string) {


		err := slack.SendMessage()
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
