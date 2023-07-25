/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type Evaluation struct {
	Questions map[string]string
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
	Short: "Converts csv to markdown using the template file",
	Long:  "",
	Run:   converter,
}

func converter(cmd *cobra.Command, args []string) {
	// type Inventory struct {
	// 	Material string
	// 	Count    uint
	// }
	// sweaters := Inventory{"wool", 17}
	// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpl.Execute(os.Stdout, sweaters)
	// if err != nil {
	// 	panic(err)
	// }

	e := &Evaluation{}
	e.Questions = make(map[string]string)
	e.Questions["q_firstyearprepared_1"] = "Was they prepared?"
	e.Questions["q_firstyearprepared_2"] = "How prepared was they?"
	e.Questions["a_firstyearprepared_1"] = "They was prepared."
	e.Questions["a_firstyearprepared_2"] = "They was not prepared."

	template, err := template.ParseFiles("test.md")
	// Capture any error
	if err != nil {
		log.Fatalln(err)
	}
	// Print out the template to std
	log.Println(e.Questions)
	template.Execute(os.Stdout, e.Questions)
}
