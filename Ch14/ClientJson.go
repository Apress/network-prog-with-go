/* ClientGet
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	//"reflect"
)

const flashcard_xml string = "application/x.flashcards+xml"
const flashcard_json string = "application/x.flashcards+json"

type FlashcardSets struct {
        XMLName string `xml:"cardsets"`
        CardSet    []CardSet `xml:"cardset"`
}

type CardSet struct {
        XMLName string `xml:"cardset"`
        Name string `xml:"name"`
        Link string `xml:"href,attr"`
        Cards []Card `xml:"card"`
}

type Card  struct {
        Name string `xml:"name"`
        Link string `xml:"href,attr"`
}


type FlashcardSetsJson struct {
	CardSet []CardSetJson `json:"cardsets"`
}
type CardSetJson struct {
	Name string `json:"name"`
	Link string `json:"@id"`
	Cards []CardJson `json:"cardset,omitempty"`
}
type CardJson struct {
	Name string `json:"name"`
	Link string `json:"@id"`
}

func getOneFlashcard(url *url.URL, client *http.Client) string {
	// Get the card as a string, don't do anything with it
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	response, err := client.Do(request)
	checkError(err)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		fmt.Println(response.Header)

		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	body, err := ioutil.ReadAll(response.Body)
	content := string(body[:])
	//fmt.Printf("Body is %s", content)

	return content
}

func getOneFlashcardSet(url *url.URL, client *http.Client) CardSetJson {
	// Get one set of cards
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	// only accept our media types
	request.Header.Add("Accept", flashcard_json)
	response, err := client.Do(request)
	checkError(err)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		fmt.Println(response.Header)

		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	body, err := ioutil.ReadAll(response.Body)
	content := string(body[:])
	fmt.Printf("Body is %s", content)

	var sets CardSetJson
	contentType := getContentType(response)
	if contentType == "JSON" {
		err = json.Unmarshal(body, &sets)
		checkError(err)
		fmt.Println("JSON: ", sets)
	}
       
	return sets
}

func getFlashcardSets(url *url.URL, client *http.Client) FlashcardSetsJson {
	// Get the toplevel /
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	// only accept our media types
	request.Header.Add("Accept", flashcard_json)
	response, err := client.Do(request)
	checkError(err)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		fmt.Println(response.Header)

		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	body, err := ioutil.ReadAll(response.Body)
	content := string(body[:])
	fmt.Printf("Body is %s", content)

	var sets FlashcardSetsJson
	contentType := getContentType(response)
	if contentType == "JSON" {
		err = json.Unmarshal(body, &sets)
		checkError(err)
		fmt.Println("JSON: ", sets)
		//fmt.Println(reflect.TypeOf(sets))
	}
	return sets
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)

	client := &http.Client{}

	flashcardSets := FlashcardSetsJson{}
	flashcardSets.CardSet = []CardSetJson{CardSetJson{}}
	flashcardSets.CardSet[0].Name = `n1`
	flashcardSets.CardSet[0].Link = `l1`
	//flashcardSets.CardSet[0].Cards = []CardJson{CardJson{Name: `n`, Link: `l`}}
	bytes, _ := json.Marshal(flashcardSets)
	fmt.Println(string(bytes[:]))
		
	url = url
	client = client


	// Step 1: get a list of flashcard sets
	flashcardSets = getFlashcardSets(url, client)
	fmt.Println("Step1, Cardsets are: ", flashcardSets)

	// Step 2: using the first flashcard set,
	//         get the list of cards in it
	set_url, _ := url.Parse(os.Args[1] + flashcardSets.CardSet[0].Link)
	oneFlashcardSet := getOneFlashcardSet(set_url, client)
	fmt.Println("Step 2: ", oneFlashcardSet)

	// Step 3: get the contents of one flashcard
	//         be lazy, just get as text/plain and
	//         don't do anything with it
	card_url, _ :=  url.Parse(os.Args[1] + oneFlashcardSet.Cards[0].Link)
	fmt.Println("Asking for URL: ", card_url.String())
	oneFlashcard := getOneFlashcard(card_url, client)
	fmt.Println("Step 3", oneFlashcard)
	
	os.Exit(0)
}

func getContentType(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if strings.Contains(contentType, flashcard_xml) {
		return "XML"
	}
	if strings.Contains(contentType, flashcard_json) {
		return "JSON"
	}
	return ""
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
