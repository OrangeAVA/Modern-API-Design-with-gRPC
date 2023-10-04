package restbooksclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (a *App) AddBook(bookJson []byte) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	bookReq := bytes.NewBuffer(bookJson)

	url := fmt.Sprintf("%s/books", a.serverAddr)

	resp, err := c.Post(url, "application/json", bookReq)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body : %s\n", body)
}

func (a *App) UpdateBook(bookJson []byte) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	bookReq := bytes.NewBuffer(bookJson)

	url := fmt.Sprintf("%s/books", a.serverAddr)

	req, err := http.NewRequest("PUT", url, bookReq)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body : %s\n", body)
}

func (a *App) ListBooks() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	url := fmt.Sprintf("%s/books/all", a.serverAddr)

	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body : %s\n", body)
}

func (a *App) FetchBook(isbn int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	url := fmt.Sprintf("%s/books/%d", a.serverAddr, isbn)

	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body : %s\n", body)
}

func (a *App) RemoveBook(isbn int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	url := fmt.Sprintf("%s/books/%d", a.serverAddr, isbn)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body : %s\n", body)
}
