package templates

import "html/template"

var TmplDays = template.Must(template.ParseFiles("templates/days.html"))
var TmplDay = template.Must(template.ParseFiles("templates/day.html"))

type DaysTemplateData struct {
	Days []Day
}

type Day struct {
	Name string
	Url string
}

type DayTemplateData struct {
	Day string
	Readme template.HTML
	Input []string
	Part1 string
	Part2 string
}