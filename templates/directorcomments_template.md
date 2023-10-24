# Private Review Comments

*This file is for the Graduate Director and Graduate Advisor & Coordinator only!*

{{range $element := .}}
\newpage
## {{$element.Semester}} {{$element.Year}} {{$element.Review}}-Year {{if eq $element.Rereview "Yes"}}(RE-REVIEW){{else}}Review{{end}} of {{$element.Student}} by {{$element.Faculty}}

**Overall Rating:** {{$element.OverallRating}}  

### Comments to Student:

{{$element.PublicComment}}

### Comments to Grad Director and Grad Advisor & Coordinator:

{{$element.PrivateComment}}

{{end}}