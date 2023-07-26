---
header-includes: \usepackage{longtable}\setlength{\LTleft}{0em}
---

# {{.answers.Year}} {{.answers.Review}}-Year {{.answers.Semester}} Review Evaluation{{if eq .answers.rereview "Yes"}} (RE-REVIEW){{end}}

Student Name: {{.answers.studentname}}  
Faculty Name: {{.answers.facultyname}}  

## {{.answers.Review}}-Year Review Goals

| Goal | Rating |
| ---- | ---- |
| {{.questions.firstyeargoals_1}} | **{{.answers.firstyeargoals_1}}** |
| {{.questions.firstyeargoals_2}} | **{{.answers.firstyeargoals_2}}** |

## {{.answers.Review}}-Year Student Responsibilities

| Goal | Rating |
| --------------------------------------------------------------------- | ---- |
| {{.questions.firstyearprepared_1}} | **{{.answers.firstyearprepared_1}}** |
| {{.questions.firstyearprepared_2}} | **{{.answers.firstyearprepared_2}}** |
| {{.questions.firstyearprepared_3}} | **{{.answers.firstyearprepared_3}}** |
| {{.questions.firstyearprepared_4}} | **{{.answers.firstyearprepared_4}}** |
| {{.questions.firstyearprepared_5}} | **{{.answers.firstyearprepared_5}}** |

## Overall Rating

> **{{.answers.overallevaluation}}**

*Please see the “Formal Review Evaluation” section of the MFA Handbook to understand what this rating means.*

## Comments

*{{.questions.studentcomments}}*

> {{.answers.studentcomments}}
