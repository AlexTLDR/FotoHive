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
			Question: "Is there a free version?",
			Answer:   "Yes, but it's not as good as this one",
		},
		{
			Question: "What are your support hours?",
			Answer:   "Nonstop, we're here to help",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@example.com">support@example.com</a>`,
		},
		{
			Question: "Is this how Dynamic templates work?",
			Answer:   "Yes, this is how they work",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
