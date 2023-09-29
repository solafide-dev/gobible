# Go Bible

Gobible is a library for interacting with the Bible in Go.

It is designed to be a simple, easy to use API for working with the Bible.

## Bible Formats

Initially, Gobible was created as an effort to support an easy to use JSON format for the Bible for use in other projects,
but through development has grown to instead be a library for working with that format, as well and importing other formats into the Gobible format.
### Supported Formats
- Gobible JSON [Example](https://raw.githubusercontent.com/gobible/gobible/master/data/WEB.json)
- OSIS XML [Example](https://raw.githubusercontent.com/gobible/gobible/master/data/WEB.xml)


Please note that since Gobible works around a data structure defined by the Gobible JSON format, some features of other formats may not be supported.

## Usage

### Initialization

There are a few ways to initialize a bible:

```go
// Initialize an empty GoBible object
b := NewGoBible()

// Load a GoBible formatted JSON file
b.Load("KJV.json")

// Load an alternate support format
b.LoadFormat("WEB.xml", "osis")
```

One you have a Bible object with some translations loaded into it, you can use this object in a few ways.

### Interacting with Individual Translations

You can load an individual translation, and use it similarly to this:

```go
// Load the KJV translation
kjv := bible.GetTranslation("KJV")

// get a verse
verse := kjv.GetBook("Genesis").GetChapter(1).GetVerse(1)

fmt.Println("Genesis 1:1 - " + verse.Text)
// Genesis 1:1 - In the beginning God created the heaven and the earth.
```

### Interacting by Reference

Once you have 1 or multiple translations loaded, you can interact with them by reference notation.

(If you have multiple translations loaded, the first translation loaded will be used as the default)


```go

// Examples of some valid references:
//   John 3:16
//   John 3:16-18
//   1 Cor 1:2-5
//   1 Cor 1:2-2:5 
// You can also tag the translation at the end of the reference
//   John 3:16 WEB 
//   John 3:16 KJV

// Lookup by reference
refs, _ := b.ParseReference("Gen 1:31-2:1")
for _, r := range refs {
    fmt.Printf("%s %d:%d %s", r.Book, r.Chapter, r.VerseNumber, r.Verse.Text)
}
// Genesis 1:31 God saw everything that he had made, and, behold, it was very good. There was evening and there was morning, the sixth day.
// Genesis 2:1 The heavens and the earth were finished, and all the host of them.
```
