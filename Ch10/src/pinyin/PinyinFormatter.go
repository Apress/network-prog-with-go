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
			// no tone, lose '5'
			words[n] = string(chars[0 : len(chars)-1])
			println("lost accent on", words[n])
			continue
		}
		if tone < '1' || tone > '4' {
			// no tone mark
			continue
		}
		words[n] = addAccent(word, int(tone))
	}
	line = strings.Join(words, ` `)
	w.Write([]byte(line))
}

var (
	// maps 'a1' to 'ā' ('\u0101') etc
	aAccent = map[int]rune{
		'1': 'ā', // '\u0101',
		'2': 'á', // '\u00e1',
		'3': 'ǎ', // '\u01ce',
		'4': 'à'} // '\u00e0'}
	eAccent = map[int]rune{
		'1': 'ē', // '\u0113',
		'2': 'é', // '\u00e9',
		'3': 'ě', // '\u011b',
		'4': 'è'} // '\u00e8'}
	iAccent = map[int]rune{
		'1': 'ī', // '\u012b',
		'2': 'í', // '\u00ed',
		'3': 'ǐ', // '\u01d0',
		'4': 'ì'} //'\u00ec'}
	oAccent = map[int]rune{
		'1': 'ō', // '\u014d',
		'2': 'ó', // '\u00f3',
		'3': 'ǒ', // '\u01d2',
		'4': 'ò'} // '\u00f2'}
	uAccent = map[int]rune{
		'1': 'ū', // '\u016b',
		'2': 'ú', // '\u00fa',
		'3': 'ǔ', // '\u01d4',
		'4': 'ù'} // '\u00f9'}
	üAccent = map[int]rune{
		'1': 'ǖ',
		'2': 'ǘ',
		'3': 'ǚ',
		'4': 'ǜ'}
)

func addAccent(word string, tone int) string {
	/*
	 * Based on "Where do the tone marks go?"
	 * at http://www.pinyin.info/rules/where.html:
         *
         *    A and e trump all other vowels and always take the tone mark.
         *    In the combination ou, o takes the mark.
         *    In all other cases, the final vowel takes the mark.
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
