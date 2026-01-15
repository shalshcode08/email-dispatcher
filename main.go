package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipentChannel := make(chan Recipient)

	go func() {
		loadRecipients("./emails.csv", recipentChannel)
	}()

	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipentChannel, &wg)
	}

	wg.Wait()
}

func executeTempelate(r Recipient) (string, error) {
	t, err := template.ParseFiles("./email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
