package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/go-fingerprint/fingerprint"
	chromaprint "github.com/go-fingerprint/gochroma"
)

var API_KEY = os.Getenv("ACOUSTID_API_KEY")

type AcoustIDRequest struct {
	fingerprint string
	duration    int
	client      string
}

func getAudioFingerPrint(filePath string) (string, int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	duration := GetDuration(filePath)

	printer := chromaprint.New(chromaprint.AlgorithmDefault)
	defer printer.Close()

	fp, err := printer.Fingerprint(fingerprint.RawInfo{
		Src:        file,
		Channels:   2,
		Rate:       44100,
		MaxSeconds: 120,
	})
	if err != nil {
		log.Fatalf("Failed to fingerprint: %v", err)
	}

	return fp, int(duration)
}

func createAPIRequest(p AcoustIDRequest) *http.Request {
	req, err := http.NewRequest("POST", "https://api.acoustid.org/v2/lookup", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("client", p.client)
	q.Add("duration", fmt.Sprint(p.duration))
	q.Add("fingerprint", p.fingerprint)
	q.Add("meta", "recordings+releasegroups+compress")
	req.URL.RawQuery = q.Encode()

	return req
}

func main() {
	if len(API_KEY) == 0 {
		log.Fatalf((`No api key found. Provide the api key when running:
  $ ACOUSTID_API_KEY=... go run .
		`))
	}

	cwd, _ := os.Getwd()
	filePath := path.Join(cwd, os.Args[1])
	fmt.Printf("Getting fingerprint for %s\n", filePath)

	fp, duration := getAudioFingerPrint(filePath)

	client := &http.Client{}
	res, err := client.Do(createAPIRequest(AcoustIDRequest{
		fingerprint: fp,
		duration:    duration,
		client:      API_KEY,
	}))
	if err != nil {
		log.Fatalf("POST request failed: %v", err)
	}

	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)

	fmt.Printf("JSON response: %s", data)
}
