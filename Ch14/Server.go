/* Server
 */

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FlashcardSet struct {
	Name string
	Link string
}

type Flashcard struct {
	Name string
	Link string
}

const flashcard_xml string = "application/x.flashcards+xml"
const flashcard_json string = "application/x.flashcards+json"

type ValueQuality struct {
	Value   string
	Quality float64
}

/* Based on https://siongui.github.io/2015/02/22/go-parse-accept-language/ */
func parseValueQuality(s string) []ValueQuality {
	var vqs []ValueQuality

	strs := strings.Split(s, `,`)
	for _, str := range strs {
		trimmedStr := strings.Trim(str, ` `)
		valQ := strings.Split(trimmedStr, `;`)
		if len(valQ) == 1 {
			vq := ValueQuality{valQ[0], 1}
			vqs = append(vqs, vq)
		} else {
			qp := strings.Split(valQ[1], `=`)
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				q = 0
			}
			vq := ValueQuality{valQ[0], q}
			vqs = append(vqs, vq)
		}
	}
	return vqs
}

func qualityOfValue(value string, vqs []ValueQuality) float64 {
	for _, vq := range vqs {
		if value == vq.Value {
			return vq.Quality
		}

	}
	return 0
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], ":port\n")
		os.Exit(1)
	}
	port := os.Args[1]

	http.HandleFunc(`/`, handleFlashCardSets)
	files, err := ioutil.ReadDir(`flashcardSets`)
	checkError(err)
	for _, file := range files {
		fmt.Println(file.Name())
		cardset_url := `/flashcardSets/` + url.QueryEscape(file.Name())
		fmt.Println("Adding handlers for ", cardset_url)
		http.HandleFunc(cardset_url, handleOneFlashCardSet)
		http.HandleFunc(cardset_url + `/`, handleOneFlashCard)
	}
	
	
	// deliver requests to the handlers
	err = http.ListenAndServe(port, nil)
	checkError(err)
	// That's it!
}

func hasIllegalChars(s string) bool {
	// check against chars to break out of current dir
	b, err := regexp.Match("[/$~]", []byte(s))
	if err != nil {
		fmt.Println(err)
		return true
	}
	if b {
		return true
	}
	return false
}

func handleOneFlashCard(rw http.ResponseWriter, req *http.Request) {
	// should be PathUnescape
	path, _ := url.QueryUnescape(req.URL.String())
	// lose intial '/'
	path = path[1:]
	if req.Method == "GET" {
		fmt.Println("Handling card: ", path)
		json_contents, err := ioutil.ReadFile(path)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(`Resource not found`))
			return
		}
		// Be lazy here, just return the content as text/plain
		rw.Write(json_contents)
		return
	} else if req.Method == "DELETE" {
		rw.WriteHeader(http.StatusNotImplemented)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

func handleFlashCardSets(rw http.ResponseWriter, req *http.Request) {
	if req.URL.String() != `/` {
		// this function only handles '/'
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Resource not found\n"))
		return
	}
	if req.Method == "GET" {
		acceptTypes := parseValueQuality(req.Header.Get("Accept"))
		fmt.Println(acceptTypes)

		q_xml := qualityOfValue(flashcard_xml, acceptTypes)
		q_json := qualityOfValue(flashcard_json, acceptTypes)
		if q_xml == 0 && q_json == 0 {
			// can't find XML or JSON in Accept header
			rw.Header().Set("Content-Type", "flashcards+xml, flashcards+json")
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}
		
		files, err := ioutil.ReadDir(`flashcardSets`)
		checkError(err)
		numfiles := len(files)
		cardSets := make([]FlashcardSet, numfiles, numfiles)
		for n, file := range files {
			fmt.Println(file.Name())
			cardSets[n].Name = file.Name()
			// should be PathEscape, not in go 1.6
			cardSets[n].Link = `/flashcardSets/` + url.QueryEscape(file.Name())
		}

		if q_xml >= q_json {
			// XML preferred
			t, err := template.ParseFiles("xml/ListFlashcardSets.xml")
			if err != nil {
				fmt.Println("Template error")
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			rw.Header().Set("Content-Type", flashcard_xml)
			t.Execute(rw, cardSets)
		} else {
			// JSON preferred
			t, err := template.ParseFiles("json/ListFlashcardSets.json")
			if err != nil {
				fmt.Println("Template error")
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			rw.Header().Set("Content-Type", flashcard_json)
			t.Execute(rw, cardSets)

		}
	} else if req.Method == "POST" {
		name := req.FormValue(`name`)
		if hasIllegalChars(name) {
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		// lose all spaces as they are a nuisance
		name = strings.Replace(name, ` `, ``, -1)
		err := os.Mkdir(`flashcardSets/`+name,
			(os.ModeDir | os.ModePerm))
		if err != nil {
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		base_url := req.URL.String()
		new_url := base_url + `flashcardSets/` + name + `/`
		rw.Write([]byte(new_url))
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

func handleOneFlashCardSet(rw http.ResponseWriter, req *http.Request) {
	cooked_url, _ := url.QueryUnescape(req.URL.String())
	fmt.Println("Handling one set for: ", cooked_url)

	if req.Method == "GET" {
		acceptTypes := parseValueQuality(req.Header.Get("Accept"))
		fmt.Println(acceptTypes)

		q_xml := qualityOfValue(flashcard_xml, acceptTypes)
		q_json := qualityOfValue(flashcard_json, acceptTypes)
		if q_xml == 0 && q_json == 0 {
			// can't find XML or JSON in Accept header
			rw.Header().Set("Content-Type", "flashcards+xml, flashcards+json")
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}

		path := req.URL.String()
		// lose leading /
		relative_path := path[1:]
		files, err := ioutil.ReadDir(relative_path)
		checkError(err)
		numfiles := len(files)
		cards := make([]Flashcard, numfiles, numfiles)
		for n, file := range files {
			fmt.Println(file.Name())
			cards[n].Name = file.Name()
			// should be PathEscape, not in go 1.6
			cards[n].Link = path + `/` + url.QueryEscape(file.Name())
		}

		if q_xml >= q_json {
			// XML preferred
			t, err := template.ParseFiles("xml/ListOneFlashcardSet.xml")
			if err != nil {
				fmt.Println("Template error")
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			rw.Header().Set("Content-Type", flashcard_xml)
			t.Execute(os.Stdout, cards)
			t.Execute(rw, cards)
		} else {
			// JSON preferred
			t, err := template.ParseFiles("json/ListOneFlashcardSet.json")
			if err != nil {
				fmt.Println("Template error", err)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			rw.Header().Set("Content-Type", flashcard_json)
			t.Execute(rw, cards)

		}
	} else if req.Method == "POST" {
		name := req.FormValue(`name`)
		if hasIllegalChars(name) {
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		err := os.Mkdir(`flashcardSets/`+name,
			(os.ModeDir | os.ModePerm))
		if err != nil {
			rw.WriteHeader(http.StatusForbidden)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		base_url := req.URL.String()
		new_url := base_url + `flashcardSets/` + name
		_, _ = rw.Write([]byte(new_url))
	} else if req.Method == "DELETE" {
		rw.WriteHeader(http.StatusNotImplemented)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
