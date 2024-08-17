package routes

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UploadDocument(w http.ResponseWriter, r *http.Request) {
	// Get params
	contentType := r.URL.Query().Get("content_type")
	sideUrl := r.URL.Query().Get("url")

	// Read body, it is an image
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// New request
	reqUrl := sideUrl
	req, err := http.NewRequest(http.MethodPut, reqUrl, strings.NewReader(string(body)))
	if err != nil {
		http.Error(w, "Error making PUT request", http.StatusInternalServerError)
		return
	}

	// Add content-type
	req.Header.Set("Content-Type", contentType)

	// Do PUT request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending PUT request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Println("request sent!")

	// Read response
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading PUT response", http.StatusInternalServerError)
		return
	}

	// Write same response
	w.WriteHeader(resp.StatusCode)
	w.Write(resBody)
}
