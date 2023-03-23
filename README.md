# Go Bible

Gobible is a library for interacting with the protestant bible in Go.

It is designed to be a simple, easy to use API for working with the Bible.

## Bible Formats

Initially, Gobible was created as an effort to support an easy to use JSON format for the Bible for use in other projects,
but through development has grown to instead be a library for working with that format, as well and importing other formats into the Gobible format.
### Supported Formats
- Gobible JSON [Example](https://raw.githubusercontent.com/applehat/gobible/master/data/WEB.json)
- OSIS XML [Example](https://raw.githubusercontent.com/applehat/gobible/master/data/WEB.xml)


Please note that since Gobible works around a data structure defined by the Gobible JSON format, some features of other formats may not be supported.




## Usage

There are a few ways to initialize a bible:

```go

    // Load a bible directly in the gobible json format
    bible := gobible.New("WEB.json") // load from file
    bible := gobible.NewFromString("{...}") // load from json string


    // load an OSIS formatted XML bible
    bible := gobible.NewOSIS("WEB.xml")

    // load internal bibles (KJV or WEB)
    bible := gobible.LoadInternal("WEB") // load from internal WEB
    bible := gobible.LoadInternal("KJV") // load from internal KJV

```

One you have a Bible object, you can load a verse like this:

```go
    // get a verse
    verse := bible.GetBook("Genesis").GetChapter(1).GetVerse(1)

    fmt.Println("Genesis 1:1 - " + verse.Text)
    // Genesis 1:1 - In the beginning God created the heaven and the earth.

```

Or you can lookup by reference:

```go
    // Lookup by reference
    refs, _ := bible.ParseReference("Gen 1:31-2:1")
    for _, r := range refs {
      fmt.Printf("%s %d:%d %s", r.Book, r.Chapter, r.VerseNumber, r.Verse.Text)
    }
    // Genesis 1:31 God saw everything that he had made, and, behold, it was very good. There was evening and there was morning, the sixth day.
    // Genesis 2:1 The heavens and the earth were finished, and all the host of them.

```