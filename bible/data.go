package bible

// This file contains the data structures used by the Bible struct
type BookData struct {
	BookNumber int
	Name       string
	Testament  string
	Genre      string
}

var booksTable = []BookData{
	{
		BookNumber: 1,
		Name:       "Genesis",
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 2,
		Name:       "Exodus",
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 3,
		Name:       "Leviticus",
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 4,
		Name:       "Numbers",
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 5,
		Name:       "Deuteronomy",
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 6,
		Name:       "Joshua",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 7,
		Name:       "Judges",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 8,
		Name:       "Ruth",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 9,
		Name:       "1 Samuel",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 10,
		Name:       "2 Samuel",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 11,
		Name:       "1 Kings",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 12,
		Name:       "2 Kings",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 13,
		Name:       "1 Chronicles",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 14,
		Name:       "2 Chronicles",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 15,
		Name:       "Ezra",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 16,
		Name:       "Nehemiah",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 17,
		Name:       "Esther",
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 18,
		Name:       "Job",
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 19,
		Name:       "Psalms",
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 20,
		Name:       "Proverbs",
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 21,
		Name:       "Ecclesiastes",
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 22,
		Name:       "Song of Solomon",
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 23,
		Name:       "Isaiah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 24,
		Name:       "Jeremiah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 25,
		Name:       "Lamentations",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 26,
		Name:       "Ezekiel",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 27,
		Name:       "Daniel",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 28,
		Name:       "Hosea",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 29,
		Name:       "Joel",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 30,
		Name:       "Amos",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 31,
		Name:       "Obadiah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 32,
		Name:       "Jonah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 33,
		Name:       "Micah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 34,
		Name:       "Nahum",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 35,
		Name:       "Habakkuk",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 36,
		Name:       "Zephaniah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 37,
		Name:       "Haggai",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 38,
		Name:       "Zechariah",
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 39,

		Name:      "Malachi",
		Testament: "OT",
		Genre:     "Prophets",
	},
	{
		BookNumber: 40,
		Name:       "Matthew",
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 41,
		Name:       "Mark",
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 42,
		Name:       "Luke",
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 43,
		Name:       "John",
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 44,
		Name:       "Acts",
		Testament:  "NT",
		Genre:      "Acts",
	},
	{
		BookNumber: 45,
		Name:       "Romans",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 46,
		Name:       "1 Corinthians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 47,
		Name:       "2 Corinthians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 48,
		Name:       "Galatians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 49,
		Name:       "Ephesians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 50,
		Name:       "Philippians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 51,
		Name:       "Colossians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 52,
		Name:       "1 Thessalonians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 53,
		Name:       "2 Thessalonians",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 54,
		Name:       "1 Timothy",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 55,
		Name:       "2 Timothy",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 56,
		Name:       "Titus",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 57,
		Name:       "Philemon",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 58,
		Name:       "Hebrews",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 59,
		Name:       "James",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 60,
		Name:       "1 Peter",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 61,
		Name:       "2 Peter",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 62,
		Name:       "1 John",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 63,
		Name:       "2 John",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 64,
		Name:       "3 John",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 65,
		Name:       "Jude",
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 66,
		Name:       "Revelation",
		Testament:  "NT",
		Genre:      "Apocalyptic",
	},
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
