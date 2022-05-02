package main

import (
	json2 "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func WordFrequencyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeClientError(w, "Method not allowed")
		return
	}

	// Parse request
	body, err, done := readRequestBody(w, r)
	if done {
		return
	}

	// Sanitize strings from special characters and convert bytes to string
	bodyString := sanitizeWords(body, err)

	// Count words frequency
	wordCount := countWords(bodyString)

	// Find top 10 words by bucket sorting
	top10Words := getTop10Words(wordCount)

	writeSuccessResponse(w, top10Words)
}

func main() {
	http.HandleFunc("/", WordFrequencyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getTop10Words This method uses bucket sorting O(n) instead of full sorting O(n lg n)
func getTop10Words(wordCount map[string]int) map[string]int {
	frequencyBucket := make(map[int][]string) // maps frequency to list of words
	for word, count := range wordCount {
		frequencyBucket[count] = append(frequencyBucket[count], word)
	}

	top10Words := make(map[string]int)

	maxFrequency := 0
	for freq := range frequencyBucket {
		if freq > maxFrequency {
			maxFrequency = freq
		}
	}

top10WordsLoop:
	for i := maxFrequency; i >= 0; i-- {
		for _, word := range frequencyBucket[i] {
			top10Words[word] = i
			if len(top10Words) >= 10 {
				break top10WordsLoop
			}
		}
	}
	return top10Words
}

func countWords(bodyString string) map[string]int {
	words := strings.Fields(bodyString)
	wordCounter := make(map[string]int)
	for _, word := range words {
		wordCounter[strings.ToLower(word)]++
	}
	return wordCounter
}

func sanitizeWords(body []byte, err error) string {
	bodyString := string(body)
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	bodyString = re.ReplaceAllString(bodyString, " ")
	return bodyString
}

func readRequestBody(w http.ResponseWriter, r *http.Request) ([]byte, error, bool) {
	var body []byte
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		writeClientError(w, "Cannot read request body")
		return nil, nil, true
	}
	return body, err, false
}

func writeClientError(w http.ResponseWriter, message string) {
	response := make(map[string]string)
	response["error"] = message
	json, err := json2.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(400)
	_, err = w.Write(json)
	if err != nil {
		panic(err)
	}
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	response := make(map[string]interface{})
	response["data"] = data
	json, err := json2.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write(json)
	if err != nil {
		panic(err)
	}
}
