package dictionary

import (
	"bufio"
	//"fmt"
	"os"
	"strings"
)

type Entry struct {
	Traditional  string
	Simplified   string
	Pinyin       string
	Translations []string
}

func (de Entry) String() string {
	str := de.Traditional + ` ` + de.Simplified + ` ` + de.Pinyin
	for _, t := range de.Translations {
		str = str + "\n    " + t
	}
	return str
}

type Dictionary struct {
	Entries []*Entry
}

func (d *Dictionary) String() string {
	str := ""
	for n := 0; n < len(d.Entries); n++ {
		de := d.Entries[n]
		str += de.String() + "\n"
	}
	return str
}

func (d *Dictionary) LookupPinyin(py string) *Dictionary {
	newD := new(Dictionary)
	v := make([]*Entry, 0, 100)
	for n := 0; n < len(d.Entries); n++ {
		de := d.Entries[n]
		if de.Pinyin == py {
			v = append(v, de)
		}
	}
	newD.Entries = v
	return newD
}

func (d *Dictionary) LookupEnglish(eng string) *Dictionary {
	newD := new(Dictionary)
	v := make([]*Entry, 0, 100)
	for n := 0; n < len(d.Entries); n++ {
		de := d.Entries[n]
		for _, e := range de.Translations {
			if e == eng {
				v = append(v, de)
			}
		}
	}
	newD.Entries = v
	return newD
}

func (d *Dictionary) LookupSimplified(simp string) *Dictionary {
	newD := new(Dictionary)
	v := make([]*Entry, 0, 100)

	for n := 0; n < len(d.Entries); n++ {
		de := d.Entries[n]
		if de.Simplified == simp {
			v = append(v, de)
		}
	}
	newD.Entries = v
	return newD
}

func (d *Dictionary) Load(path string) {

	f, err := os.Open(path)
	r := bufio.NewReader(f)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	v := make([]*Entry, 0, 100000)
	numEntries := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if line[0] == '#' {
			continue
		}
		// fmt.Println(line)
		trad, simp, pinyin, translations := parseDictEntry(line)

		de := Entry{
			Traditional:  trad,
			Simplified:   simp,
			Pinyin:       pinyin,
			Translations: translations}

		v = append(v, &de)
		numEntries++
	}
	// fmt.Printf("Num entries %d\n", numEntries)
	d.Entries = v
}

func parseDictEntry(line string) (string, string, string, []string) {
	// format is
	//    trad simp [pinyin] /trans/trans/.../
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
