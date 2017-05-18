/* Client
 */

package main

import (
	//"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
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

func getOneFlashcardSet(url *url.URL, client *http.Client) CardSet {
	// Get one set of cards
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	// only accept our media types
	request.Header.Add("Accept", flashcard_xml)
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

	var sets CardSet
	contentType := getContentType(response)
	if contentType == "XML" {

		err = xml.Unmarshal(body, &sets)
		checkError(err)
		fmt.Println("XML: ", sets)
		return sets
	}
	/* else if contentType == "JSON" {
		var sets FlashcardSetsJson
		err = json.Unmarshal(body, &sets)
		checkError(err)
		fmt.Println("JSON: ", sets)
	}
        */
	return sets
}

func getFlashcardSets(url *url.URL, client *http.Client) FlashcardSets {
	// Get the toplevel /
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	// only accept our media types
	request.Header.Add("Accept", flashcard_xml)
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

	var sets FlashcardSets
	contentType := getContentType(response)
	if contentType == "XML" {

		err = xml.Unmarshal(body, &sets)
		checkError(err)
		fmt.Println("XML: ", sets)
		return sets
	}
	return sets
}

func createFlashcardSet(url1 *url.URL, client *http.Client, name string) string {
	data := make(url.Values)
	data[`name`] = []string{name}
	response, err := client.PostForm(url1.String(), data)
	checkError(err)
	if response.StatusCode != http.StatusCreated {
		fmt.Println(`Error: `, response.Status)
		return ``
		//os.Exit(2)
	}
	body, err := ioutil.ReadAll(response.Body)
	content := string(body[:])
	return content
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)

	client := &http.Client{}

	// Step 1: get a list of flashcard sets
	flashcardSets := getFlashcardSets(url, client)
	fmt.Println("Step 1: ", flashcardSets)

	// Step 2: try to create a new flashcard set
	new_url := createFlashcardSet(url, client, `NewSet`)
	fmt.Println("Step 2: New flashcard set has URL: ", new_url)

	// Step 3: using the first flashcard set,
	//         get the list of cards in it
	set_url, _ := url.Parse(os.Args[1] + flashcardSets.CardSet[0].Link)

	fmt.Println("Asking for flashcard set URL: ", set_url.String())
	oneFlashcardSet := getOneFlashcardSet(set_url, client)
	fmt.Println("Step 3:", oneFlashcardSet)

	// Step 4: get the contents of one flashcard
	//         be lazy, just get as text/plain and
	//         don't do anything with it
	card_url, _ :=  url.Parse(os.Args[1] + oneFlashcardSet.Cards[0].Link)
	fmt.Println("Asking for URL: ", card_url.String())
	oneFlashcard := getOneFlashcard(card_url, client)
	fmt.Println("Step 4", oneFlashcard)
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
