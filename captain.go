package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// this file is a simple demonstration of how to authenticate githooks sent
// by Github's Webhooks API: https://developer.github.com/webhooks/

// Note that this is designed for Github specifically.
func webhookIsAuthenticated(xHubSignature string, hookBody []byte) bool {
	hmacDigest := hmac.New(sha1.New, []byte(os.Getenv("GITHOOK_SECRET")))
	hmacDigest.Write(hookBody)

	var buffer bytes.Buffer
	// all signatures are preceeded with 'sha1='
	buffer.WriteString("sha1=")
	buffer.WriteString(hex.EncodeToString(hmacDigest.Sum(nil)))

	return hmac.Equal(buffer.Bytes(), []byte(xHubSignature))
}

func githookListener(w http.ResponseWriter, req *http.Request) {
	// all Github-sent webhooks are POST requests, we should reject otherwise.
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// if no header is present for X-Hub-Signature, we should reject the request.
		if len(req.Header["X-Hub-Signature"]) > 0 {
			if webhookIsAuthenticated(req.Header["X-Hub-Signature"][0], body) {
				fmt.Println("Received Webhook is authentic")
				// do whatever here
			} else {
				w.WriteHeader(404)
			}
		} else {
			w.WriteHeader(404)
		}
	} else {
		w.WriteHeader(404)
	}
}

func main() {
	// sets up a simple route for your webhook to send payloads to
	http.HandleFunc("/build/github", githookListener)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
