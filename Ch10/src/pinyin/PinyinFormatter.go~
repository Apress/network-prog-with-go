package pinyin

import (
	"io"
	"strings"
)

func PinyinFormatter(w io.Writer, format string, value ...interface{}) {
	line := value[0].(string)
	words := strings.Fields(line)
	for n, word := range words {
		// convert "u:" to "ü" if present
		uColon := strings.Index(word, "u:")
		if uColon != -1 {
			parts := strings.SplitN(word, "u:", 2)
			word = parts[0] + "ü" + parts[1]
		}
		println(word)
		// get last character, will be the tone if present
		chars := []rune(word)
		tone := chars[len(chars)-1]
		if tone == '5' {
			words[n] = string(chars[0 : len(chars)-1])
			println("lost accent on", words[n])
			continue
		}
		if tone < '1' || tone > '4' {
			continue
		}
		words[n] = addAccent(word, int(tone))
	}
	line = strings.Join(words, ` `)
	w.Write([]byte(line))
}

var (
	// maps 'a1' to '\u0101' etc
	aAccent = map[int]rune{
		'1': '\u0101',
		'2': '\u00e1',
		'3': '\u01ce', // '\u0103',
		'4': '\u00e0'}
	eAccent = map[int]rune{
		'1': '\u0113',
		'2': '\u00e9',
		'3': '\u011b', // '\u0115',
		'4': '\u00e8'}
	iAccent = map[int]rune{
		'1': '\u012b',
		'2': '\u00ed',
		'3': '\u01d0', // '\u012d',
		'4': '\u00ec'}
	oAccent = map[int]rune{
		'1': '\u014d',
		'2': '\u00f3',
		'3': '\u01d2', // '\u014f',
		'4': '\u00f2'}
	uAccent = map[int]rune{
		'1': '\u016b',
		'2': '\u00fa',
		'3': '\u01d4', // '\u016d',
		'4': '\u00f9'}
	üAccent = map[int]rune{
		'1': 'ǖ',
		'2': 'ǘ',
		'3': 'ǚ',
		'4': 'ǜ'}
)

func addAccent(word string, tone int) string {
	/*
	 * Based on "Where do the tone marks go?"
	 * at http://www.pinyin.info/rules/where.html
	 */

	n := strings.Index(word, "a")
	if n != -1 {
		aAcc := aAccent[tone]
		// replace 'a' with its tone version
		word = word[0:n] + string(aAcc) + word[(n+1):len(word)-1]
	} else {
		n := strings.Index(word, "e")
		if n != -1 {
			eAcc := eAccent[tone]
			word = word[0:n] + string(eAcc) +
				word[(n+1):len(word)-1]
		} else {
			n = strings.Index(word, "ou")
			if n != -1 {
				oAcc := oAccent[tone]
				word = word[0:n] + string(oAcc) + "u" +
					word[(n+2):len(word)-1]
			} else {
				chars := []rune(word)
				length := len(chars)
				// put tone onthe last vowel
			L:
				for n, _ := range chars {
					m := length - n - 1
					switch chars[m] {
					case 'i':
						chars[m] = iAccent[tone]
						break L
					case 'o':
						chars[m] = oAccent[tone]
						break L
					case 'u':
						chars[m] = uAccent[tone]
						break L
					case 'ü':
						chars[m] = üAccent[tone]
						break L
					default:
					}
				}
				word = string(chars[0 : len(chars)-1])
			}
		}
	}

	return word
}
