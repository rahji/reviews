/*
2023 Rob Duarte MIT License
github.com/rahji
*/

package cmd

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var templateFile string
var csvFile string
var outputDir string

// make a type for each csv field, which is a struct with a string for the column name
// and a string for the column text (eg: the full text of the question)

type Field struct {
	ColumnNumber   int
	ColumnName     string
	ColumnQuestion string
}

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

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.reviews.yaml)")
	rootCmd.PersistentFlags().StringVar(&templateFile, "template", "template.md", "template file (default: template.md)")
	rootCmd.PersistentFlags().StringVar(&csvFile, "input", "input.csv", "input csv file (default: input.csv)")
	rootCmd.PersistentFlags().StringVar(&outputDir, "outputdir", ".", "output folder (default: current directory)")

}
