package gobible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseReference(t *testing.T) {
	b := NewGoBible()
	b.LoadFormat("data/WEB.xml", "osis")

	// Test a valid reference
	r, err := b.ParseReference("Genesis 1:1-3")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(r))
	assert.Equal(t, "Genesis", r[2].Book)
	assert.Equal(t, 1, r[2].Chapter)
	assert.Equal(t, 3, r[2].Verse)

	// Test a valid reference
	r, err = b.ParseReference("Genesis 1:1-2:3")
	assert.Nil(t, err)
	assert.Equal(t, 34, len(r))
	assert.Equal(t, "Genesis", r[33].Book)
	assert.Equal(t, 2, r[33].Chapter)
	assert.Equal(t, 3, r[33].Verse)
	assert.Equal(t, 1, r[4].Chapter)
	assert.Equal(t, 5, r[4].Verse)

	// Test a valid reference
	r, err = b.ParseReference("1 Corinthians 1:2")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, "1 Corinthians", r[0].Book)
	assert.Equal(t, 1, r[0].Chapter)
	assert.Equal(t, 2, r[0].Verse)

}

func TestParseReferenceVersion(t *testing.T) {
	b := NewGoBible()

	_, err := b.ParseReference("Genesis 1:1")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "no bibles loaded")
	}

	err = b.LoadFormat("data/WEB.xml", "osis")
	assert.Nil(t, err)
	err = b.Load("data/KJV.json")
	assert.Nil(t, err)

	// Test a valid reference
	r, err := b.ParseReference("Genesis 1:1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, "Genesis", r[0].Book)
	assert.Equal(t, 1, r[0].Chapter)
	assert.Equal(t, 1, r[0].Verse)
	assert.Equal(t, "WEB", r[0].Translation)
	assert.Equal(t, "In the beginning God created the heavens and the earth.", r[0].VerseRef.Text)

	// Test a valid reference with a translation
	r, err = b.ParseReference("Genesis 1:1 KJV")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, "Genesis", r[0].Book)
	assert.Equal(t, 1, r[0].Chapter)
	assert.Equal(t, 1, r[0].Verse)
	assert.Equal(t, "KJV", r[0].Translation)
	assert.Equal(t, "In the beginning God created the heaven and the earth.", r[0].VerseRef.Text)

	// Test a reference with an invalid translation
	// This should just return the WEB
	r, err = b.ParseReference("Genesis 1:1 ASV")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, "Genesis", r[0].Book)
	assert.Equal(t, 1, r[0].Chapter)
	assert.Equal(t, 1, r[0].Verse)
	assert.Equal(t, "WEB", r[0].Translation)

	s := r[0].String()
	assert.Equal(t, "In the beginning God created the heavens and the earth. - Genesis 1:1 WEB", s)
}

func TestParseReferenceBad(t *testing.T) {
	b := NewGoBible()
	err := b.Load("data/KJV.json")
	assert.Nil(t, err)

	// Test a bad reference
	_, err = b.ParseReference("Genesis 1:1:1")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid chapter format")
	}

	// Test a bad reference
	_, err = b.ParseReference("Genesis")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid reference format")
	}

	// Test a bad reference
	_, err = b.ParseReference("Genesis 1:B")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid syntax")
	}

	// Test a bad reference
	_, err = b.ParseReference("Genesis A:2")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid chapter format")
	}

	// Test a bad reference
	_, err = b.ParseReference("Genesis 1:2-")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid syntax")
	}

	// Test a bad reference
	_, err = b.ParseReference("Genesis 3.5-2.3")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid chapter format")
	}

	// Test a bad reference
	_, err = b.ParseReference("Genesis B:B")
	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "invalid chapter format")
	}

}
