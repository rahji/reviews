# Review Comments for 

{{range $key, $val := .}}
{{$key}}
{{end}}

{{range $element := .}}
{{$element.Year}} {{$element.Semester}} {{$element.Review}}-Year {{if eq $element.Rereview "Yes"}}**RE-REVIEW**{{else}}Review{{end}}
{{$element.FacultyName}}
{{$element.OverallRating}}
{{$element.PrivateComment}}
{{$element.PublicComment}}
{{end}}
