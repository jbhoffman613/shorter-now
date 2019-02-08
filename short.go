package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
)

// Hello displays `hello world` on the homepage of the site.
func Hello(w http.ResponseWriter, r *http.Request) {
	// Ensure that the root page is only presents on /
	if r.URL.Path != "/" {
		errorf(w, http.StatusNotFound, "page not found: %q", r.URL.Path)
		return
	}
	if r.Method != http.MethodGet {
		errorf(w, http.StatusMethodNotAllowed, "expected method 'GET', recieved %v", r.Method)
		return
	}

	fmt.Fprintf(w, "Hello World: %q", r.URL.Path)
	log.Println("Called Hello")
}

// CreateSha7 creates the shortened version of a string
func CreateSha7(url string) string {
	h := hmac.New(sha256.New, []byte("secret"))
	h.Write([]byte(url))
	return fmt.Sprintf("%x \n", h.Sum(nil)[0:4])
}

// errorf creates an http error by consuming a http error code, a response writer, a formatted string,
// and arguments to insert into the string.
func errorf(w http.ResponseWriter, code int, msg string, args ...interface{}) {
	errorMessage := fmt.Sprintf(msg, args)
	http.Error(w, errorMessage, code)
	log.Printf("WARNING: %s", errorMessage)
}

func main() {
	print(CreateSha7("HELLO!"))
	// http.HandleFunc("/", Hello)
	// localPort := ":8000"
	// log.Printf("Program started on port %s", localPort)
	// log.Fatal(http.ListenAndServe(localPort, nil))
}
