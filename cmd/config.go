/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Token     string `mapstructure:"token" json:"token"`
	ChannelID string `mapstructure:"channel-id" json:"channel-id"`
	Cookie    string `mapstructure:"cookie" json:"cookie"`
	CronExpr  string `mapstructure:"cron-expr" json:"cron-expr"`
	Message   string `mapstructure:"message" json:"message"`
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show all configurations",
	Long:  `Show all configurations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("No config file found, please set config first")
				return nil
			} else {
				return err
			}
		}

		config := Config{}

		err := viper.Unmarshal(&config)
		if err != nil {
			return err
		}

		confJson, err := json.MarshalIndent(config, "", "  ")

		if err != nil {
			return err
		}

		fmt.Printf("%s\n", confJson)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
