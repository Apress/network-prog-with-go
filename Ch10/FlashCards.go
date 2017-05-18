package flashcard

import (
	"container/vector"
	"dictionary"
)

type FlashCard struct {
	English string
	Card    DictionaryEntry
}

type FlashCards struct {
	vector.Vector
}
