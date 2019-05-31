package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {
	book := Book{Title: "Cloud native go", Author: "M.-L. Reimer", ISBN: "0123456789" }
	json := book.ToJSON()

	assert.Equal(t,
		`{"title":"Cloud native go","author":"M.-L. Reimer","isbn":"0123456789"}`,
		string(json),
		"Book JSON marshling wrong")
}

func TestJsonToBook(t *testing.T) {
	json := []byte(`{"title":"Cloud native go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)

	assert.Equal(t,
		Book{Title: "Cloud native go", Author: "M.-L. Reimer", ISBN: "0123456789" },
		book,
		"Book JSON unmarshling wrong")
}
