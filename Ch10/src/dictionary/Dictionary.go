package dictionary

type Entry struct {
	Traditional  string
	Simplified   string
	Pinyin       string
	Translations []string
}

type Dictionary struct {
	Entries []*Entry
}
