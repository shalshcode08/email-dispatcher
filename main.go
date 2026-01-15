package main

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipentChannel := make(chan Recipient)
	loadRecipients("./emails.csv", recipentChannel)
}
