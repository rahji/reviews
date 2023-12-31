/*
2023 Rob Duarte MIT License
github.com/rahji
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Reviews",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
