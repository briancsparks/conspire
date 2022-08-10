/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/briancsparks/conspire/conspire"

	"github.com/spf13/cobra"
)

// servewsCmd represents the servews command
var servewsCmd = &cobra.Command{
	Use:   "servews",
	Short: "Run a WebSocket server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("servews called")
		conspire.WsMain()
	},
}

func init() {
	rootCmd.AddCommand(servewsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// servewsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// servewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
