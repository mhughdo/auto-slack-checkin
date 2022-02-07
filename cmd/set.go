/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config.",
	Long:  `Set configs. Available flags: token, cookie, channel-id.`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config := Config{}

		err := viper.Unmarshal(&config)
		if err != nil {
			return
		}

		viper.WriteConfig()
	},
}

func init() {
	configCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("token", "t", "", "Your slack token which starts with \"xoxc\"")
	setCmd.Flags().StringP("cookie", "c", "", "Your slack cookie. Only cookie with key \"d\" is required. For example: d=HDGHJSGHJDSGJ723672GJHSGR;")
	setCmd.Flags().StringP("channel-id", "i", "", "Your slack channel id.")
	viper.BindPFlag("token", setCmd.Flags().Lookup("token"))
	viper.BindPFlag("cookie", setCmd.Flags().Lookup("cookie"))
	viper.BindPFlag("channel-id", setCmd.Flags().Lookup("channel-id"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
