/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sanikum3/ORP/internal/configs"
	"github.com/spf13/cobra"
)

// changebaseCmd represents the changebase command
var changebaseCmd = &cobra.Command{
	Use:   "changebase",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			println("url required")
			return
		}

		cfg, err := configs.Load()
		if err != nil {
			panic(err)
		}

		cfg.JitsiBase = args[0]

		err = configs.Save(cfg)
		if err != nil {
			panic(err)
		}

		println("Base URL changed:", cfg.JitsiBase)
	},
}

func init() {
	rootCmd.AddCommand(changebaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changebaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changebaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
