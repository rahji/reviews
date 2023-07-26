/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"text/template"

	"github.com/spf13/cobra"
)

// make a type for each csv field, which is a struct with a string for the column name
// and a string for the column text (eg: the full text of the question)

type Field struct {
	ColumnName     string
	ColumnQuestion string
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts csv to markdown using the template file.",
	Long: `Converts csv to markdown using the template file. 
(Assumes the csv contains reviews for a single year/semester/review combination (eg: 2023 Fall Second)
`,
	Run: converter,
}

func converter(cmd *cobra.Command, args []string) {

	csv := readCsvFile("file.csv")
	// line 0 is the column names
	// line 1 is the full text of the question for that column
	// line 2 is not useful for us
	// next lines are the survey results, one per line

	// make a slice of Field structs using the first two rows of the csv file
	fields := []Field{}
	for i := 0; i < len(csv[0]); i++ {
		reg := regexp.MustCompile(`^.*?\? - (.*)`)
		res := reg.ReplaceAllString(csv[1][i], "${1}") // hack to remove redundant question group text
		fields = append(fields, Field{ColumnName: csv[0][i], ColumnQuestion: res})
	}

	// loop through rest of the csv lines,
	// set up a map of maps containing that line's columnquestions and columnanswers (using columname as the key),
	// and output template for that csv line
	data := map[string]map[string]string{}
	data["questions"] = make(map[string]string)
	data["answers"] = make(map[string]string)

	//	for i := 2; i < len(csv[0]); i++ {
	for i := 4; i < 5; i++ {
		for j := 0; j < len(csv[0]); j++ {
			data["questions"][fields[j].ColumnName] = fields[j].ColumnQuestion
			data["answers"][fields[j].ColumnName] = csv[i][j]
		}
	}

	// Parse the template file
	template, err := template.ParseFiles("firstyear_template.md")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}
	// Print out the template to std
	template.Execute(os.Stdout, data)
}
