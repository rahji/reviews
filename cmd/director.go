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

// each student has a slice of FacultyReview structs
type FacultyReview struct {
	Semester       string
	Year           string
	Review         string
	Rereview       string
	Student        string
	Faculty        string
	PrivateComment string
	PublicComment  string
	OverallRating  string
}

var outputPrefix string

func init() {
	rootCmd.AddCommand(directorCmd)
	directorCmd.PersistentFlags().StringVar(&outputPrefix, "outputprefix", "", "The prefix for the markdown output filename (required)")
	directorCmd.MarkPersistentFlagRequired("outputprefix")
}

// directorCmd represents the director command
var directorCmd = &cobra.Command{
	Use:   "director",
	Short: "Creates markdown files for the grad director and grad advisor & coordinator.",
	Long: `Creates markdown files for the graduate director and grad advisor & coordinator.
One file per student is created, as opposed to the many files created for each student by the CONVERT command.
(Assumes the csv contains reviews for a single year/semester/review combination (eg: 2023 Fall Second))
`,
	Run: director,
}

func director(cmd *cobra.Command, args []string) {

	csv := readCsvFile(csvFile)
	// line 0 is the column names
	// line 1 is the full text of the question for that column
	// line 2 is not useful for us
	// next lines are the survey results, one per line

	// make a map of colnames->colnum using the first two rows of the csv file
	cols := make(map[string]int)
	for i := 0; i < len(csv[0]); i++ {
		cols[csv[0][i]] = i
	}

	// loop through rest of the csv lines
	// to set up a map of slices (each of which contains a map with the keys shown below):
	// student1 -> [0] -> FacultyReview struct (facultyname, private comment, public comment, overall rating)
	//             [1] -> FacultyReview struct (facultyname, private comment, public comment, overall rating)
	//             etc.
	// * actually, we'll include all review info in the FacultyReview struct, instead of overall metadata about this csv
	// so this command can be run on a giant csv file containing all reviews for multiple years/semesters/reviews!

	studentMap := make(map[string][]FacultyReview)

	for i := 3; i < len(csv); i++ {
		studentMap[csv[i][cols["Student"]]] = append(studentMap[csv[i][cols["Student"]]],
			FacultyReview{
				Semester:       csv[i][cols["Semester"]],
				Year:           csv[i][cols["Year"]],
				Rereview:       csv[i][cols["rereview"]],
				Review:         csv[i][cols["Review"]],
				Student:        csv[i][cols["Student"]],
				Faculty:        csv[i][cols["Faculty"]],
				PrivateComment: csv[i][cols["privatecomments"]],
				PublicComment:  csv[i][cols["studentcomments"]],
				OverallRating:  csv[i][cols["overallevaluation"]],
			})
	}

	// loop through the studentMap and create a markdown file for each student
	for student, reviews := range studentMap {

		// log.Println("student...", student)
		// log.Println("reviews...", reviews)

		// create a filename for this student
		nameswapreg := regexp.MustCompile(`^(\p{L}).*?(\p{L}+)$`)
		studentname := nameswapreg.ReplaceAllString(student, "${2}${1}") // "Rob Duarte" => "DuarteR"

		fn_str := fmt.Sprintf("%s_%s.md", outputPrefix, studentname)
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

		// loop through the reviews for this student
		// and create a markdown file for each review
		template, err := template.ParseFiles(templateFile)
		// Capture any error
		if err != nil {
			log.Fatalln(err)
		}

		template.Execute(fo, reviews)

	}
}
