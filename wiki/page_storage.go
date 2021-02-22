package wiki

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

const filenameFormat = "%s.txt"

func Save(page *Page)  {
	filename := fmt.Sprintf(filenameFormat, page.Title)
	err := ioutil.WriteFile(filename, page.Body, 0600)
	if err != nil {
		log.SetPrefix("page save")
		log.Fatal(err)
	}
	fmt.Printf("Saved: %s\n", page.Title)
}

func Load(title string) (*Page, error) {
	filename := fmt.Sprintf(filenameFormat, title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No file with title: %s", title))
	}
	return &Page{Title: title, Body: body}, err
}