package main_test

import (
	"encoding/json"
	"go-demo/main/app/rest/models"
	"os"
	"testing"
)

func TestStuff(t *testing.T) {

	book := models.Book{}

	a, err := os.Open("test.json")

	if err != nil {

		t.Error(err)
	}

	err = json.NewDecoder(a).Decode(&book)

	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", book)

}
