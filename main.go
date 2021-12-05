package main

import (
	"advent-of-code-2021/day01"
	"advent-of-code-2021/day02"
	"advent-of-code-2021/day03"
	"advent-of-code-2021/day04"
	"advent-of-code-2021/day05"
	"advent-of-code-2021/templates"
	"html/template"
	"net/http"
	"strings"
)

var days = []Day{
	&day01.Day{},
	&day02.Day{},
	&day03.Day{},
	&day04.Day{},
	&day05.Day{},
}

type Day interface {
	GetDay() string
	GetInput() string
	GetReadme() string
	Part1() string
	Part2() string
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "" || r.RequestURI == "/" {
		// main site
		x := &templates.DaysTemplateData{}
		for i := range days {
			x.Days = append(x.Days, templates.Day{
				Name: days[i].GetDay(),
				Url: "/" + days[i].GetDay(),
			})
		}

		templates.TmplDays.Execute(w, x)
	} else {
		var day string
		splitted := strings.Split(r.RequestURI, "/")
		if strings.HasSuffix(r.RequestURI, "/") {
			day = splitted[len(splitted)-2]
		} else {
			day = splitted[len(splitted)-1]
		}

		for i := range days {
			if days[i].GetDay() == day {
				templates.TmplDay.Execute(w, &templates.DayTemplateData{
					Day: days[i].GetDay(),
					Input: strings.Split(days[i].GetInput(), "\n"),
					Readme: template.HTML(strings.ReplaceAll(template.HTMLEscapeString(days[i].GetReadme()), "\n", "<br/>")),
					Part1: days[i].Part1(),
					Part2: days[i].Part2(),
				})
			}
		}
	}
}