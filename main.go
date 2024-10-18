package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	fmt.Println("Do you want to create or delete a tweet? (create/delete):")
	var operation string
	fmt.Scanln(&operation)

	if operation == "create" {
		createTweet()
	} else if operation == "delete" {
		deleteTweet()
	} else {
		fmt.Println("Invalid operation")
	}
}

func createTweet() {

	fmt.Println("Enter tweet content:")
	var tweetContent string
	fmt.Scanln(&tweetContent)

	payload := fmt.Sprintf(`{"text": "%s"}`, tweetContent)

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf(`OAuth oauth_consumer_key="%s",oauth_token="%s",oauth_signature_method="%s",oauth_timestamp="%s",oauth_nonce="%s",oauth_version="%s",oauth_signature="%s"`,
		os.Getenv("OAUTH_CONSUMER_KEY"), os.Getenv("OAUTH_TOKEN"), os.Getenv("OAUTH_SIGNATURE_METHOD"), os.Getenv("OAUTH_TIMESTAMP"), os.Getenv("OAUTH_NONCE"), os.Getenv("OAUTH_VERSION"), os.Getenv("OAUTH_SIGNATURE")))
	req.Header.Set("Cookie", os.Getenv("COOKIE"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

func deleteTweet() {

	fmt.Println("Enter tweet ID to delete:")
	var tweetID string
	fmt.Scanln(&tweetID)

	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf(`OAuth oauth_consumer_key="%s",oauth_token="%s",oauth_signature_method="%s",oauth_timestamp="%s",oauth_nonce="%s",oauth_version="%s",oauth_signature="%s"`,
		os.Getenv("OAUTH_CONSUMER_KEY"), os.Getenv("OAUTH_TOKEN"), os.Getenv("OAUTH_SIGNATURE_METHOD"), os.Getenv("OAUTH_TIMESTAMP"), os.Getenv("OAUTH_NONCE"), os.Getenv("OAUTH_VERSION"), os.Getenv("OAUTH_SIGNATURE")))
	req.Header.Set("Cookie", os.Getenv("COOKIE"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
