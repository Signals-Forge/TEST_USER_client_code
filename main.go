package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RandomData struct {
	Value1 int    `json:"value1"`
	Value2 string `json:"value2"`
}

type Payload struct {
	URL    string `json:"url"`
	APIKey string `json:"apikey"`
}

func main() {
    http.HandleFunc("/start-client", startClientHandler)
    port := os.Getenv("PORT") 
    if port == "" {
        port = "8080"
    }
	http.ListenAndServe(fmt.Sprintf(":%s", port) , nil)
}

func startClientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Process the extracted values
	fmt.Println("Received URL:", payload.URL)
	fmt.Println("Received API Key:", payload.APIKey)

	// Send a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Client start request received")

	//START CONNECTION
	fmt.Println("Start connection")
	req, err := http.NewRequest("GET", payload.URL, nil)
    if err != nil {
		fmt.Println("START CLIENT HANDLER 3")
        return
    }

	fmt.Println("Set Authorization")
    // Set the Authorization header
    req.Header.Set("Authorization", payload.APIKey)

    // Make the request
    client := &http.Client{}
	fmt.Println("DO REQUEST")
    resp, err := client.Do(req)
    if err != nil {
		fmt.Println("I FAILED")		
		fmt.Println("START CLIENT HANDLER 4", err.Error())
        return
    }	
    defer resp.Body.Close()

	fmt.Println("Should I start?")
    //c.Status(http.StatusOK)
	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF { // Handle errors (but not EOF)
			fmt.Println("START CLIENT HANDLER 7")
			fmt.Println(err)
			return
		}

		// Process only lines starting with "data: "
		if len(line) > 6 && line[:6] == "data: " {
			var data RandomData
			err = json.Unmarshal([]byte(line[6:]), &data)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Received1:", data)
			}
		}

		if err == io.EOF { // End of stream reached
			break
		}
	}
}
