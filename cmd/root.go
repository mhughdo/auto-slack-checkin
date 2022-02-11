/*
Copyright © 2022 Hugh Do <mhughdo@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto-slack-checkin",
	Short: "Auto checkin to slack",
	Long:  `This is a CLI tool to automatically checkin to slack.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.auto-slack-checkin.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".auto-slack-checkin" (without extension).
		configName := ".auto-slack-checkin"
		configType := "json"
		configPath := filepath.Join(home, configName+"."+configType)
		viper.AddConfigPath(home)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		_, err = os.Stat(configPath)
		if os.IsNotExist(err) {
			if _, err := os.Create(configPath); err != nil { // perm 0666
				fmt.Printf("Failed to create config file: %s\n", err)
				os.Exit(1)
			}
			config := Config{
				Token:     "",
				ChannelID: "",
				Cookie:    "",
				CronExpr:  "0 8 * * *",
				Message:   "",
			}
			err = viper.Unmarshal(&config)
			if err != nil {
				return
			}

			viper.WriteConfig()
			fmt.Printf("Config file created: %s\n", configPath)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
