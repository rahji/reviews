/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var templateFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "reviews",
	Short: "Converts csv results from Qualtrics MFA reviews to markdown files",
	Long: `Converts csv results from Qualtrics MFA reviews to markdown files.
	
This is specific to the FSU Department of Art MFA reviews, 
but could be the starting point for something more general.
Qualtrics does not have the ability to export survey results
as PDFs in batch.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.reviews.yaml)")
	rootCmd.PersistentFlags().StringVar(&templateFile, "template", "", "template file (default is $HOME/review.md)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}