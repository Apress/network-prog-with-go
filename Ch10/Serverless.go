/* Server
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
)

import (
	"dictionary"
	"flashcards"
	"templatefuncs"
)

var d *dictionary.Dictionary

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], ":port\n")
		os.Exit(1)
	}
	port := os.Args[1]

	dictionaryPath := "cedict_ts.u8"
	d = new(dictionary.Dictionary)
	d.Load(dictionaryPath)
	fmt.Println("Loaded dict", len(d.Entries))

	http.HandleFunc("/", listFlashCards)
	fileServer := http.StripPrefix("/jscript/", http.FileServer(http.Dir("jscript")))
	http.Handle("/jscript/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
	http.Handle("/html/", fileServer)

	http.HandleFunc("/flashcards.html", listFlashCards)
	http.HandleFunc("/flashcardSets", manageFlashCards)

	// deliver requests to the handlers
	err := http.ListenAndServe(port, nil)
	checkError(err)
	// That's it!
}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {
	...
}

/*
 * Called from ListFlashcards.html on form submission
 */
func manageFlashCards(rw http.ResponseWriter, req *http.Request) {
	...
}

func showFlashCards(rw http.ResponseWriter, cardname, order, half string) {
	...
}

func listWords(rw http.ResponseWriter, cardname string) {
	...
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
