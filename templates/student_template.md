---
header-includes: 
  - \usepackage{longtable}
  - \setlength{\LTleft}{0em}
---

# {{.answers.Year}} {{.answers.Review}}-Year {{.answers.Semester}} Review Evaluation{{if eq .answers.rereview "Yes"}} (RE-REVIEW){{end}}

Student Name: {{.answers.Student}}  
Faculty Name: {{.answers.Faculty}}  

## {{.answers.Review}}-Year Review Goals

| Goal | Rating | |
| --------------------------------------------------------------------- | ---- | --- |
{{if .answers.firstyeargoals_1}}| {{.questions.firstyeargoals_1}} | **{{.answers.firstyeargoals_1}}** | |
{{end -}}
{{if .answers.firstyeargoals_2}}| {{.questions.firstyeargoals_2}} | **{{.answers.firstyeargoals_2}}** | |
{{end -}}
{{if .answers.secondyeargoals_1}}| {{.questions.secondyeargoals_1}} | **{{.answers.secondyeargoals_1}}** | |
{{end -}}
{{if .answers.secondyeargoals_2}}| {{.questions.secondyeargoals_2}} | **{{.answers.secondyeargoals_2}}** | |
{{end -}}
{{if .answers.secondyeargoals_3}}| {{.questions.secondyeargoals_3}} | **{{.answers.secondyeargoals_3}}** | |
{{end -}}
{{if .answers.thirdyeargoals_1}}| {{.questions.thirdyeargoals_1}} | **{{.answers.thirdyeargoals_1}}** | |
{{end -}}
{{if .answers.thirdyeargoals_2}}| {{.questions.thirdyeargoals_2}} | **{{.answers.thirdyeargoals_2}}** | |
{{end}}

## {{.answers.Review}}-Year Student Responsibilities

| Goal | Rating | |
| --------------------------------------------------------------------- | ---- | --- |
{{if .answers.firstyearprepared_1}}| {{.questions.firstyearprepared_1}} | **{{.answers.firstyearprepared_1}}** | |
{{end -}}
{{if .answers.firstyearprepared_2}}| {{.questions.firstyearprepared_2}} | **{{.answers.firstyearprepared_2}}** | |
{{end -}}
{{if .answers.firstyearprepared_3}}| {{.questions.firstyearprepared_3}} | **{{.answers.firstyearprepared_3}}** | |
{{end -}}
{{if .answers.firstyearprepared_4}}| {{.questions.firstyearprepared_4}} | **{{.answers.firstyearprepared_4}}** | |
{{end -}}
{{if .answers.firstyearprepared_5}}| {{.questions.firstyearprepared_5}} | **{{.answers.firstyearprepared_5}}** | |
{{end -}}
{{if .answers.secondyearprepared_1}}| {{.questions.secondyearprepared_1}} | **{{.answers.secondyearprepared_1}}** | |
{{end -}}
{{if .answers.secondyearprepared_2}}| {{.questions.secondyearprepared_2}} | **{{.answers.secondyearprepared_2}}** | |
{{end -}}
{{if .answers.secondyearprepared_3}}| {{.questions.secondyearprepared_3}} | **{{.answers.secondyearprepared_3}}** | |
{{end -}}
{{if .answers.secondyearprepared_4}}| {{.questions.secondyearprepared_4}} | **{{.answers.secondyearprepared_4}}** | |
{{end -}}
{{if .answers.thirdyearprepared_1}}| {{.questions.thirdyearprepared_1}} | **{{.answers.thirdyearprepared_1}}** | |
{{end -}}
{{if .answers.thirdyearprepared_2}}| {{.questions.thirdyearprepared_2}} | **{{.answers.thirdyearprepared_2}}** | |
{{end -}}
{{if .answers.thirdyearprepared_3}}| {{.questions.thirdyearprepared_3}} | **{{.answers.thirdyearprepared_3}}** | |
{{end -}}
{{if .answers.thirdyearprepared_4}}| {{.questions.thirdyearprepared_4}} | **{{.answers.thirdyearprepared_4}}** | |
{{end -}}
{{if .answers.thirdyearprepared_5}}| {{.questions.thirdyearprepared_5}} | **{{.answers.thirdyearprepared_5}}** | |
{{end -}}
{{if .answers.thirdyearprepared_6}}| {{.questions.thirdyearprepared_6}} | **{{.answers.thirdyearprepared_6}}** | |
{{end}}

## Overall Rating

> **{{.answers.overallevaluation}}**

Please see the “Formal Review Evaluation” section of the MFA Handbook to understand what this rating means.{{if eq .answers.Review "First"}} Note that the only possible Ratings for a First-Year student are "Satisfactory" and "Provisional."{{end}}

## Comments

*{{.questions.studentcomments}}*

{{.answers.studentcomments}}
