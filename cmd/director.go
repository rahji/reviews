/*
2023 Rob Duarte MIT License
github.com/rahji
*/
package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// each student has a slice of FacultyReview structs
type FacultyReview struct {
	FacultyName    string
	PrivateComment string
	PublicComment  string
	OverallRating  string
}

func init() {
	rootCmd.AddCommand(directorCmd)
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

	/*
		loop through the whole file while creating a map of maps
		student --> faculty member -> private comment, public comment, overall rating
		iterate over that map
		for each student, feed its faculty member map to the template (which will have looping logic in it)
	*/

	// loop through rest of the csv lines
	// to set up a map of slices (each of which contains a map with the keys shown below):
	// student1 -> [0] -> FacultyReview struct (facultyname, private comment, public comment, overall rating)
	//             [1] -> FacultyReview struct (facultyname, private comment, public comment, overall rating)
	//             etc.
	studentMap := make(map[string][]FacultyReview)

	for i := 3; i < len(csv); i++ {
		studentMap[csv[i][cols["studentname"]]] = append(studentMap[csv[i][cols["studentname"]]],
			FacultyReview{
				FacultyName:    csv[i][cols["facultyname"]],
				PrivateComment: csv[i][cols["privatecomments"]],
				PublicComment:  csv[i][cols["studentcomments"]],
				OverallRating:  csv[i][cols["overallevaluation"]],
			})
	}

	// data := struct {
	//     Year string
	//     Semester string
	// 	Review string
	// 	Rereview string
	//     Rows   []FacultyReview
	// }{
	//     fields[],
	//     table,
	//     studentMap["jerry"],
	// }

	// data := map[string]map[string]string{}
	// data["questions"] = make(map[string]string)
	// data["answers"] = make(map[string]string)

	// for i := 3; i < len(csv); i++ {

	// 	// for each student, create a map of maps
	// 	// faculty1 -> private comment, public comment, overall rating
	// 	// faculty2 -> private comment, public comment, overall rating
	// 	// etc.
	// 	studentMap[csv[i][0]] = make(map[string]map[string]string)

	// 	for j := 0; j < len(csv[0]); j++ {
	// 		data["questions"][fields[j].ColumnName] = fields[j].ColumnQuestion
	// 		data["answers"][fields[j].ColumnName] = csv[i][j]
	// 	}

	// // make a filename from the review info
	// // eg: 2023Spring_First_REREVIEW_studentname_facultyname.md
	// // or: 2024Fall_Second_studentname_facultyname.md
	// rereview := ""
	// if data["answers"]["rereview"] == "Yes" {
	// 	rereview = "_REREVIEW"
	// }
	// nameswapreg := regexp.MustCompile(`^(\p{L}).*?(\p{L}+)$`)
	// studentname := nameswapreg.ReplaceAllString(data["answers"]["studentname"], "${2}${1}") // "Rob Duarte" => "DuarteR"
	// facultyname := nameswapreg.ReplaceAllString(data["answers"]["facultyname"], "${2}${1}") // "Rob Duarte" => "DuarteR"

	// fn_str := fmt.Sprintf("%s%s_%s%s_%s_%s.md",
	// 	data["answers"]["Year"],
	// 	data["answers"]["Semester"],
	// 	data["answers"]["Review"],
	// 	rereview,
	// 	studentname,
	// 	facultyname)
	// fn, err := filenamify.Filenamify(fn_str, filenamify.Options{})
	// if err != nil {
	// 	panic(err)
	// }

	// // create the output file
	// fo, err := os.Create(outputDir + "/" + fn)
	// if err != nil {
	// 	panic(err)
	// }
	// // close fo on exit and check for its returned error
	// defer func() {
	// 	if err := fo.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// // Parse the template file
	// template, err := template.ParseFiles(templateFile)
	// // Capture any error
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // Print out the template to std
	// template.Execute(fo, data)

	//spew.Dump(studentMap)

	template, err := template.ParseFiles("private.md")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}

	template.Execute(os.Stdout, studentMap["Rich Bott"])

}
