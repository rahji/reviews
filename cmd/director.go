/*
2023 Rob Duarte MIT License
github.com/rahji
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

func init() {
	rootCmd.AddCommand(directorCmd)
	// Here you will define your flags and configuration settings.
}

// directorCmd represents the director command
var directorCmd = &cobra.Command{
	Use:   "director",
	Short: "Creates a single markdown file containing all private comments in the csv file.",
	Long: `Creates a single markdown file containing all private comments in the csv file. 
(Assumes the csv contains reviews for a single year/semester/review combination (eg: 2023 Fall Second)
`,
	Run: director,
}

func director(cmd *cobra.Command, args []string) {

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
