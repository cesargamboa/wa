package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	baseURL, err := url.Parse("https://api.whatsapp.com/send")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	params := url.Values{}
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("Enter the phone number you want to send a WhatsApp message")
	} else {
		params.Add("phone", args[0])
	}

	if len(args) >= 2 {
		params.Add("text", args[1])
	}

	baseURL.RawQuery = params.Encode()
	endpoint := baseURL.String()

	fmt.Println("Opening", endpoint)
	exec.Command("open", endpoint).Start()
}
