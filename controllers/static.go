package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What's your LinkedIn?",
			Answer:   `<a href="https://www.linkedin.com/in/alexandru-badragan/">linkedin.com/in/alexandru-badragan</a>`,
		},
		{
			Question: "How can I reach you?",
			Answer:   "Drop me an email at alex@alextldr.com",
		},
		{
			Question: "Why does the frontend look so bad?",
			Answer:   "I'm a backend developer, I don't know how to make things look good",
		},
		{
			Question: "What is your github handle?",
			Answer:   "alextldr",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
