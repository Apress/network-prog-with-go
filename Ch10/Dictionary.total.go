package main

import (
	"bufio"
	"fmt"
	"os"
	//"fmt"
	"container/vector"
	"strings"
)

func main() {
	dictionaryPath := "/var/www/go/chinese/cedict_ts.u8"
	d := new(Dictionary)
	d.loadDictionary(dictionaryPath)

	/*
		for n := 1; n < 10; n++ {
			v := d.At(n)
			de := v.(*DictionaryEntry)
			fmt.Println(de.String())
		}
	*/
	//haoD := d.LookupPinyin("hao2")
	//fmt.Println(haoD.String())

	goodD := d.LookupEnglish("good")
	fmt.Println(goodD.String())
}

type DictionaryEntry struct {
	Traditional  string
	Simplified   string
	Pinyin       string
	Translations []string
}

func (de DictionaryEntry) String() string {
	str := de.Traditional + de.Simplified + de.Pinyin
	for _, t := range de.Translations {
		str = str + "\n    " + t
	}
	return str
}

type Dictionary struct {
	vector.Vector
}

func (d *Dictionary) String() string {
	str := ""
	for n := 0; n < d.Len(); n++ {
		de := d.At(n).(*DictionaryEntry)
		str += de.String() + "\n"
	}
	return str
}

func (d *Dictionary) LookupPinyin(py string) *Dictionary {
	newD := new(Dictionary)
	for n := 0; n < d.Len(); n++ {
		de := d.At(n).(*DictionaryEntry)
		if de.Pinyin == py {
			newD.Push(de)
		}
	}
	return newD
}

func (d *Dictionary) LookupEnglish(eng string) *Dictionary {
	newD := new(Dictionary)
	for n := 0; n < d.Len(); n++ {
		de := d.At(n).(*DictionaryEntry)
		for _, e := range de.Translations {
			if e == eng {
				newD.Push(de)
			}
		}
	}
	return newD
}

func (d *Dictionary) loadDictionary(path string) {

	f, err := os.Open(path)
	r := bufio.NewReader(f)
	if err != nil {
		fmt.Println(err.String())
		os.Exit(1)
	}
	for {
		line, err := r.ReadString('\n')
		if line[0] == '#' {
			continue
		}
		if err != nil {
			return
		}

		trad, simp, pinyin, translations := parseDictEntry(line)

		de := DictionaryEntry{
			Traditional:  trad,
			Simplified:   simp,
			Pinyin:       pinyin,
			Translations: translations}

		d.Push(&de)
	}
}

func parseDictEntry(line string) (string, string, string, []string) {
	// format is
	//    trad simp [pinyin] /trans/trans/.../
	/* doesn't work:
	_, err = fmt.Sscanf(line, "%s %s [[]%200c[]] /%200c/",
	&trad, &simp, &pinyin, &meaning)
	if err != nil {
	fmt.Print(err.String(), line)
	}
	*/
	tradEnd := strings.Index(line, " ")
	trad := line[0:tradEnd]
	line = strings.TrimSpace(line[tradEnd:])

	simpEnd := strings.Index(line, " ")
	simp := line[0:simpEnd]
	line = strings.TrimSpace(line[simpEnd:])

	pinyinEnd := strings.Index(line, "]")
	pinyin := line[1:pinyinEnd]
	line = strings.TrimSpace(line[pinyinEnd+1:])

	translations := strings.Split(line, "/")
	// includes empty at start and end, so
	translations = translations[1 : len(translations)-1]

	return trad, simp, pinyin, translations
}

type FlashCard struct {
	English string
	Card    DictionaryEntry
}

type FlashCards struct {
	vector.Vector
}
