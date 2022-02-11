/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	slack "auto-slack-checkin/internal/pkg/slack"

	"github.com/go-co-op/gocron"
	"github.com/lnquy/cron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cronCmd represents the runCron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Run a cron job.",
	Long:  `Run a cron job to send the message.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := gocron.NewScheduler(time.FixedZone("Asia/Ho_Chi_Minh", 7*3600))
		cronExpr := viper.GetString("cron-expr")

		if viper.GetString("cron-expr") == "" {
			return fmt.Errorf("cron-expr is not set")
		}

		if viper.GetString("cron-expr") == "" {
			return fmt.Errorf("token is not set")
		}

		if viper.GetString("channel-id") == "" {
			return fmt.Errorf("channel-id is not set")
		}

		if viper.GetString("cookie") == "" {
			return fmt.Errorf("cookie is not set")
		}

		if viper.GetString("message") == "" {
			return fmt.Errorf("message is not set")
		}

		exprDesc, _ := cron.NewDescriptor(cron.Verbose(true))
		desc, err := exprDesc.ToDescription(cronExpr, "")

		if err != nil {
			return fmt.Errorf("cron-expr is invalid: %s", err)
		}

		fmt.Printf("Cron expression description: %s\n", desc)

		s.Cron(cronExpr).Do(func() error {
			err := slack.SendMessage()
			if err != nil {
				fmt.Println(err)
				return err
			}

			return nil
		})
		s.StartBlocking()
		return nil
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
