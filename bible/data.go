package bible

// This file contains the data structures used by the Bible struct

// NumberToEnglishBook is a map of the book number to the English name of the book
var bibleNumberToEnglishBook = map[int]string{
	1:  "Genesis",
	2:  "Exodus",
	3:  "Leviticus",
	4:  "Numbers",
	5:  "Deuteronomy",
	6:  "Joshua",
	7:  "Judges",
	8:  "Ruth",
	9:  "1 Samuel",
	10: "2 Samuel",
	11: "1 Kings",
	12: "2 Kings",
	13: "1 Chronicles",
	14: "2 Chronicles",
	15: "Ezra",
	16: "Nehemiah",
	17: "Esther",
	18: "Job",
	19: "Psalms",
	20: "Proverbs",
	21: "Ecclesiastes",
	22: "Song of Solomon",
	23: "Isaiah",
	24: "Jeremiah",
	25: "Lamentations",
	26: "Ezekiel",
	27: "Daniel",
	28: "Hosea",
	29: "Joel",
	30: "Amos",
	31: "Obadiah",
	32: "Jonah",
	33: "Micah",
	34: "Nahum",
	35: "Habakkuk",
	36: "Zephaniah",
	37: "Haggai",
	38: "Zechariah",
	39: "Malachi",
	40: "Matthew",
	41: "Mark",
	42: "Luke",
	43: "John",
	44: "Acts",
	45: "Romans",
	46: "1 Corinthians",
	47: "2 Corinthians",
	48: "Galatians",
	49: "Ephesians",
	50: "Philippians",
	51: "Colossians",
	52: "1 Thessalonians",
	53: "2 Thessalonians",
	54: "1 Timothy",
	55: "2 Timothy",
	56: "Titus",
	57: "Philemon",
	58: "Hebrews",
	59: "James",
	60: "1 Peter",
	61: "2 Peter",
	62: "1 John",
	63: "2 John",
	64: "3 John",
	65: "Jude",
	66: "Revelation",
}

// bookEnglishShortNames is a map of the English name of the book to a slice of short names
var bookEnglishShortNames = map[string][]string{
	"Genesis":         {"Gen", "Ge", "Genesis"},
	"Exodus":          {"Exod", "Ex", "Exo", "Exodus"},
	"Leviticus":       {"Lev", "Lv", "Leviticus"},
	"Numbers":         {"Num", "Nu", "Nm", "Numbers"},
	"Deuteronomy":     {"Deut", "Dt", "Deuteronomy"},
	"Joshua":          {"Josh", "Jos", "Joshua"},
	"Judges":          {"Judg", "Jdg", "Jg", "Judges"},
	"Ruth":            {"Ruth", "Rth"},
	"1 Samuel":        {"1 Sam", "1Sa", "1S", "1 Samuel"},
	"2 Samuel":        {"2 Sam", "2Sa", "2S", "2 Samuel"},
	"1 Kings":         {"1 Kgs", "1Ki", "1K", "1 Kings"},
	"2 Kings":         {"2 Kgs", "2Ki", "2K", "2 Kings"},
	"1 Chronicles":    {"1 Chron", "1Ch", "1 Chronicles"},
	"2 Chronicles":    {"2 Chron", "2Ch", "2 Chronicles"},
	"Ezra":            {"Ezra", "Ezr"},
	"Nehemiah":        {"Neh", "Ne", "Nehemiah"},
	"Esther":          {"Esth", "Est", "Es", "Esther"},
	"Job":             {"Job", "Jb"},
	"Psalms":          {"Psalm", "Psalms", "Ps"},
	"Proverbs":        {"Prov", "Pr", "Proverbs"},
	"Ecclesiastes":    {"Eccl", "Ecc", "Qohelet", "Ecclesiastes"},
	"Song of Solomon": {"Song", "Song of Sol", "Song of Songs", "SOS"},
	"Isaiah":          {"Isa", "Is", "Isaiah"},
	"Jeremiah":        {"Jer", "Je", "Jeremiah"},
	"Lamentations":    {"Lam", "La", "Lamentations"},
	"Ezekiel":         {"Ezek", "Eze", "Ezk", "Ezekiel"},
	"Daniel":          {"Dan", "Dn", "Daniel"},
	"Hosea":           {"Hos", "Ho", "Hosea"},
	"Joel":            {"Joel", "Jl"},
	"Amos":            {"Amos", "Am"},
	"Obadiah":         {"Obad", "Ob", "Oba", "Obadiah"},
	"Jonah":           {"Jonah", "Jon"},
	"Micah":           {"Micah", "Mic"},
	"Nahum":           {"Nahum", "Nah"},
	"Habakkuk":        {"Hab", "Habakkuk"},
	"Zephaniah":       {"Zeph", "Zep", "Zephaniah"},
	"Haggai":          {"Hag", "Haggai"},
	"Zechariah":       {"Zech", "Zec", "Zechariah"},
	"Malachi":         {"Mal", "Malachi"},
	"Matthew":         {"Matt", "Mt", "Matthew"},
	"Mark":            {"Mk", "Mark"},
	"Luke":            {"Lk", "Luke"},
	"John":            {"Jn", "John"},
	"Acts":            {"Acts", "Ac"},
	"Romans":          {"Rom", "Ro", "Rm", "Romans"},
	"1 Corinthians":   {"1 Cor", "1Co", "1 Corinthians"},
	"2 Corinthians":   {"2 Cor", "2Co", "2 Corinthians"},
	"Galatians":       {"Gal", "Galatians"},
	"Ephesians":       {"Eph", "Ephesians"},
	"Philippians":     {"Phil", "Php", "Philippians"},
	"Colossians":      {"Col", "Colossians"},
	"1 Thessalonians": {"1 Thess", "1Th", "1 Thessalonians"},
	"2 Thessalonians": {"2 Thess", "2Th", "2 Thessalonians"},
	"1 Timothy":       {"1 Tim", "1Ti", "1 Timothy"},
	"2 Timothy":       {"2 Tim", "2Ti", "2 Timothy"},
	"Titus":           {"Titus", "Ti"},
	"Philemon":        {"Philem", "Phile", "Phm", "Philemon"},
	"Hebrews":         {"Heb", "Hebrews"},
	"James":           {"Jam", "Jas", "James"},
	"1 Peter":         {"1 Pet", "1Pe", "1 Peter"},
	"2 Peter":         {"2 Pet", "2Pe", "2 Peter"},
	"1 John":          {"1 Jn", "1Jn", "1 John"},
	"2 John":          {"2 Jn", "2Jn", "2 John"},
	"3 John":          {"3 Jn", "3Jn", "3 John"},
	"Jude":            {"Jude", "Jud"},
	"Revelation":      {"Rev", "Re", "Revelation"},
}

/*
func NormalizeFileFormat(file string) {
	// load the file into the bible struct
	bible := NewBible(file)
	for i, book := range bible.Books {
		// set the book number based on the book name
		for number, name := range bibleNumberToEnglishBook {
			if name == book.Name {
				bible.Books[i].Number = number
			}
		}
	}
	// write to originalname-normalized.json
	b, err := json.MarshalIndent(bible, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("file-normalized.json", b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
*/
