/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import _ "yunfei/app/event"

import (
	"github.com/spf13/cobra"
	"yunfei/app/route"
	"yunfei/internal/server"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start serve",
	Long:  `start serve`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func run() {
	http := server.New()
	http.GenRouter(route.New())
	http.Run()
}
