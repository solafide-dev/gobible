package bible

// This file contains the data structures used by the Bible struct
type BookData struct {
	BookNumber int
	Name       string
	Alts       []string
	Testament  string
	Genre      string
}

var booksTable = []BookData{
	{
		BookNumber: 1,
		Name:       "Genesis",
		Alts:       []string{"Gen", "Ge", "Genesis"},
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 2,
		Name:       "Exodus",
		Alts:       []string{"Exod", "Ex", "Exo", "Exodus"},
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 3,
		Name:       "Leviticus",
		Alts:       []string{"Lev", "Lv", "Leviticus"},
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 4,
		Name:       "Numbers",
		Alts:       []string{"Num", "Nu", "Nm", "Numbers"},
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 5,
		Name:       "Deuteronomy",
		Alts:       []string{"Deut", "Deu", "De", "Deuteronomy"},
		Testament:  "OT",
		Genre:      "Law",
	},
	{
		BookNumber: 6,
		Name:       "Joshua",
		Alts:       []string{"Jos", "Josh", "Joshua"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 7,
		Name:       "Judges",
		Alts:       []string{"Jdg", "Jg", "Judg", "Judges"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 8,
		Name:       "Ruth",
		Alts:       []string{"Ru", "Rut", "Ruth"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 9,
		Name:       "1 Samuel",
		Alts:       []string{"1 Sam", "1Sa", "1S", "1 Samuel"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 10,
		Name:       "2 Samuel",
		Alts:       []string{"2 Sam", "2Sa", "2S", "2 Samuel"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 11,
		Name:       "1 Kings",
		Alts:       []string{"1 Kgs", "1Ki", "1K", "1 Kings"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 12,
		Name:       "2 Kings",
		Alts:       []string{"2 Kgs", "2Ki", "2K", "2 Kings"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 13,
		Name:       "1 Chronicles",
		Alts:       []string{"1Ch", "1Chr", "1Chronicles"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 14,
		Name:       "2 Chronicles",
		Alts:       []string{"2Ch", "2Chr", "2Chronicles"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 15,
		Name:       "Ezra",
		Alts:       []string{"Ezr", "Ezra"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 16,
		Name:       "Nehemiah",
		Alts:       []string{"Ne", "Neh", "Nehemiah"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 17,
		Name:       "Esther",
		Alts:       []string{"Esth", "Est", "Es", "Esther"},
		Testament:  "OT",
		Genre:      "History",
	},
	{
		BookNumber: 18,
		Name:       "Job",
		Alts:       []string{"Jb", "Job"},
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 19,
		Name:       "Psalms",
		Alts:       []string{"Ps", "Psa", "Psalms"},
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 20,
		Name:       "Proverbs",
		Alts:       []string{"Pr", "Prov", "Proverbs"},
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 21,
		Name:       "Ecclesiastes",
		Alts:       []string{"Ec", "Ecc", "Ecclesiastes"},
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 22,
		Name:       "Song of Solomon",
		Alts:       []string{"Song", "Song of Sol", "Song of Songs", "SOS"},
		Testament:  "OT",
		Genre:      "Wisdom",
	},
	{
		BookNumber: 23,
		Name:       "Isaiah",
		Alts:       []string{"Isa", "Is", "Isaiah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 24,
		Name:       "Jeremiah",
		Alts:       []string{"Jer", "Je", "Jeremiah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 25,
		Name:       "Lamentations",
		Alts:       []string{"La", "Lam", "Lamentations"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 26,
		Name:       "Ezekiel",
		Alts:       []string{"Ezek", "Eze", "Ezk", "Ezekiel"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 27,
		Name:       "Daniel",
		Alts:       []string{"Da", "Dan", "Daniel"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 28,
		Name:       "Hosea",
		Alts:       []string{"Ho", "Hos", "Hosea"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 29,
		Name:       "Joel",
		Alts:       []string{"Joe", "Joel", "Jl"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 30,
		Name:       "Amos",
		Alts:       []string{"Am", "Amos"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 31,
		Name:       "Obadiah",
		Alts:       []string{"Obad", "Ob", "Oba", "Obadiah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 32,
		Name:       "Jonah",
		Alts:       []string{"Jon", "Jonah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 33,
		Name:       "Micah",
		Alts:       []string{"Mic", "Micah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 34,
		Name:       "Nahum",
		Alts:       []string{"Na", "Nah", "Nahum"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 35,
		Name:       "Habakkuk",
		Alts:       []string{"Hab", "Habakkuk"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 36,
		Name:       "Zephaniah",
		Alts:       []string{"Zeph", "Zep", "Zephaniah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 37,
		Name:       "Haggai",
		Alts:       []string{"Hag", "Haggai"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 38,
		Name:       "Zechariah",
		Alts:       []string{"Zech", "Zec", "Zechariah"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 39,
		Name:       "Malachi",
		Alts:       []string{"Mal", "Malachi"},
		Testament:  "OT",
		Genre:      "Prophets",
	},
	{
		BookNumber: 40,
		Name:       "Matthew",
		Alts:       []string{"Matt", "Mt", "Mat", "Matthew"},
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 41,
		Name:       "Mark",
		Alts:       []string{"Mrk", "Mr", "Mk", "Mark"},
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 42,
		Name:       "Luke",
		Alts:       []string{"Lk", "Lu", "Luk", "Luke"},
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 43,
		Name:       "John",
		Alts:       []string{"Jn", "Jhn", "Joh", "John"},
		Testament:  "NT",
		Genre:      "Gospels",
	},
	{
		BookNumber: 44,
		Name:       "Acts",
		Alts:       []string{"Ac", "Act", "Acts"},
		Testament:  "NT",
		Genre:      "Acts",
	},
	{
		BookNumber: 45,
		Name:       "Romans",
		Alts:       []string{"Ro", "Rom", "Romans"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 46,
		Name:       "1 Corinthians",
		Alts:       []string{"1 Cor", "1 Co", "1 Cor", "1 Corinthians", "1Co", "1Cor"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 47,
		Name:       "2 Corinthians",
		Alts:       []string{"2 Cor", "2 Co", "2 Cor", "2 Corinthians", "2Co", "2Cor"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 48,
		Name:       "Galatians",
		Alts:       []string{"Ga", "Gal", "Galatians"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 49,
		Name:       "Ephesians",
		Alts:       []string{"Eph", "Ephesians"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 50,
		Name:       "Philippians",
		Alts:       []string{"Php", "Phil", "Philippians"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 51,
		Name:       "Colossians",
		Alts:       []string{"Col", "Colossians"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 52,
		Name:       "1 Thessalonians",
		Alts:       []string{"1 Thess", "1 Th", "1 Thes", "1 Thess", "1 Thessalonians", "1Th", "1Thess"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 53,
		Name:       "2 Thessalonians",
		Alts:       []string{"2 Thess", "2 Th", "2 Thes", "2 Thess", "2 Thessalonians", "2Th", "2Thess"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 54,
		Name:       "1 Timothy",
		Alts:       []string{"1 Tim", "1 Ti", "1 Tim", "1 Timothy", "1Ti", "1Tim"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 55,
		Name:       "2 Timothy",
		Alts:       []string{"2 Tim", "2 Ti", "2 Tim", "2 Timothy", "2Ti", "2Tim"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 56,
		Name:       "Titus",
		Alts:       []string{"Titus", "Ti"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 57,
		Name:       "Philemon",
		Alts:       []string{"Phm", "Philem", "Philemon"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 58,
		Name:       "Hebrews",
		Alts:       []string{"Heb", "Hebrews"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 59,
		Name:       "James",
		Alts:       []string{"Jas", "Jm", "Jas", "James"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 60,
		Name:       "1 Peter",
		Alts:       []string{"1 Pet", "1 Pe", "1 Pet", "1 Peter", "1Pe", "1Pet"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 61,
		Name:       "2 Peter",
		Alts:       []string{"2 Pet", "2 Pe", "2 Pet", "2 Peter", "2Pe", "2Pet"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 62,
		Name:       "1 John",
		Alts:       []string{"1 Jn", "1 Jo", "1 Jn", "1 John", "1Jn", "1John"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 63,
		Name:       "2 John",
		Alts:       []string{"2 Jn", "2 Jo", "2 Jn", "2 John", "2Jn", "2John"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 64,
		Name:       "3 John",
		Alts:       []string{"3 Jn", "3 Jo", "3 Jn", "3 John", "3Jn", "3John"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 65,
		Name:       "Jude",
		Alts:       []string{"Jude", "Jd"},
		Testament:  "NT",
		Genre:      "Epistles",
	},
	{
		BookNumber: 66,
		Name:       "Revelation",
		Alts:       []string{"Rev", "Re", "Revelation"},
		Testament:  "NT",
		Genre:      "Apocalyptic",
	},
}
