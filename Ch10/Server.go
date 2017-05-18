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
	//"dictionary"
	"flashcards"
	"templatefuncs"
)

//var d *dictionary.Dictionary

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], ":port\n")
		os.Exit(1)
	}
	port := os.Args[1]

	//dictionaryPath := "cedict_ts.u8"
	//d = new(dictionary.Dictionary)
	//d.Load(dictionaryPath)
	//fmt.Println("Loaded dict", len(d.Entries))

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

	flashCardsNames := flashcards.ListFlashCardsNames()
	t, err := template.ParseFiles("html/ListFlashcards.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(rw, flashCardsNames)
}

/*
 * Called from ListFlashcards.html on form submission
 */
func manageFlashCards(rw http.ResponseWriter, req *http.Request) {

	set := req.FormValue("flashcardSets")
	order := req.FormValue("order")
	action := req.FormValue("submit")
	half := req.FormValue("half")
	fmt.Println("set chosen is", set)
	fmt.Println("order is", order)
	fmt.Println("action is", action)
	fmt.Println("half is", half)

	cardname := "flashcardSets/" + set

	fmt.Println("cardname", cardname, "action", action)
	if action == "Show cards in set" {
		showFlashCards(rw, cardname, order, half)
	} else if action == "List words in set" {
		listWords(rw, cardname)
	}
}

func showFlashCards(rw http.ResponseWriter, cardname, order, half string) {
	fmt.Println("Loading card name", cardname)
	cards := new(flashcards.FlashCards)
	flashcards.LoadJSON(cardname, &cards)
	if order == "Sequential" {
		cards.CardOrder = "SEQUENTIAL"
	} else {
		cards.CardOrder = "RANDOM"
	}
	fmt.Println("half is", half)
	if half == "Random" {
		cards.ShowHalf = "RANDOM_HALF"
	} else if half == "English" {
		cards.ShowHalf = "ENGLISH_HALF"
	} else {
		cards.ShowHalf = "CHINESE_HALF"
	}
	fmt.Println("loaded cards", len(cards.Cards))
	fmt.Println("Card name", cards.Name)

	t := template.New("ShowFlashcards.html")
	t = t.Funcs(template.FuncMap{"pinyin": templatefuncs.PinyinFormatter})
	t, err := t.ParseFiles("html/ShowFlashcards.html")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listWords(rw http.ResponseWriter, cardname string) {
	fmt.Println("Loading card name", cardname)
	cards := new(flashcards.FlashCards)
	flashcards.LoadJSON(cardname, cards)
	fmt.Println("loaded cards", len(cards.Cards))
	fmt.Println("Card name", cards.Name)

	t := template.New("ListWords.html")

	t = t.Funcs(template.FuncMap{"pinyin": templatefuncs.PinyinFormatter})
	t, err := t.ParseFiles("html/ListWords.html")

	if err != nil {
		fmt.Println("Parse error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("No error ")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
