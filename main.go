package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, please send an email to <a href=\"mailto:webdev@gmail.com\">webdev@gmail.com</a>.</p>")
}

func faqtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Is there a free version?</h1><p>Yes!We offer free trial for 30 days.</p>",
		"<h1>What are your support hours?</h1><p>Non-Stop.</p>",
		"<h1>How do I cancel?</h1><p>You can cancel anytime in your account settings.</p>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprintf(w, "<h1>404 Page Not Found</h1>")
// 		http.Error(w, "404 Page Not Found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqtHandler(w, r)
	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "<h1>404 Page Not Found</h1>")
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", router)
}
