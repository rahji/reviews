/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"text/template"

	"github.com/flytam/filenamify"
	"github.com/spf13/cobra"
)

var templateFile string
var csvFile string
var outputDir string

// make a type for each csv field, which is a struct with a string for the column name
// and a string for the column text (eg: the full text of the question)

type Field struct {
	ColumnName     string
	ColumnQuestion string
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

	convertCmd.Flags().StringVar(&templateFile, "template", "template.md", "template file (default: template.md)")
	convertCmd.Flags().StringVar(&csvFile, "input", "input.csv", "input csv file (default: input.csv)")
	convertCmd.Flags().StringVar(&outputDir, "outputdir", ".", "output folder (default: current directory)")

	// convertCmd.MarkFlagRequired("template")
	// convertCmd.MarkFlagRequired("input")
	// convertCmd.MarkFlagRequired("outputdir")

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

	csv := readCsvFile(csvFile)
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

	for i := 3; i < len(csv); i++ {

		for j := 0; j < len(csv[0]); j++ {
			data["questions"][fields[j].ColumnName] = fields[j].ColumnQuestion
			data["answers"][fields[j].ColumnName] = csv[i][j]
		}

		// make a filename from the review info
		// eg: 2023Spring_First_REREVIEW_studentname_facultyname.md
		// or: 2024Fall_Second_studentname_facultyname.md
		rereview := ""
		if data["answers"]["rereview"] == "Yes" {
			rereview = "_REREVIEW"
		}
		nameswapreg := regexp.MustCompile(`^(\p{L}).*?(\p{L}+)$`)
		studentname := nameswapreg.ReplaceAllString(data["answers"]["studentname"], "${2}${1}") // "Rob Duarte" => "DuarteR"
		facultyname := nameswapreg.ReplaceAllString(data["answers"]["facultyname"], "${2}${1}") // "Rob Duarte" => "DuarteR"

		fn_str := fmt.Sprintf("%s%s_%s%s_%s_%s.md",
			data["answers"]["Year"],
			data["answers"]["Semester"],
			data["answers"]["Review"],
			rereview,
			studentname,
			facultyname)
		fn, err := filenamify.Filenamify(fn_str, filenamify.Options{})
		if err != nil {
			panic(err)
		}

		// create the output file
		fo, err := os.Create(outputDir + "/" + fn)
		if err != nil {
			panic(err)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		// Parse the template file
		template, err := template.ParseFiles(templateFile)
		// Capture any error
		if err != nil {
			log.Fatalln(err)
		}
		// Print out the template to std
		template.Execute(fo, data)

	}

}
