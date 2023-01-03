package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bufio"
)

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
    }


func main() {

	type Usage struct {
		Prompt_tokens int `json:"prompt_tokens"`
		Completion_tokens int `json:"completion_tokens"`
		Total_tokens int `json:"total_tokens"`
	}

	type Response struct {
		Id string `json:"id"`
	        Object string `json:"object"`
		Created int `json:"created"`
		Model string `json:"model"`
		Choices [] struct {
			Text string `json:"text"`
			Index int `json:"index"`
			Logprobs string `json:"logprobs"`
			Finish_reason string `json:finish_reason"`
		} `json:"choices"`

		Usage struct {
                	Prompt_tokens int `json:"prompt_tokens"`
                	Completion_tokens int `json:"completion_tokens"`
                	Total_tokens int `json:"total_tokens"`
		} `json:"usage"`

	}


	// Set the API endpoint URL
	apiEndpoint := "https://api.openai.com/v1/completions"

	// Set the API key for your ChatGPT account
	apiKey := "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"
	if apiKey == "" {
		fmt.Println ("API key is null, get it from https://beta.openai.com/account/api-keys")
		os.Exit(-1)
	}

	// Set the request body with the input text and model
	Query := map[string]string{
                "model": "text-davinci-002",
	}
	
   for true {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nAsk a question:..")
        text, _ := reader.ReadString('\n')
        Query["prompt"] = text
	

	requestBody, err := json.Marshal(Query)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(requestBody));

	// Set the headers for the request
	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", apiKey)},
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the headers for the request
	req.Header = headers

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response
	//fmt.Println(string(body))
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {   // Parse []byte to go struct pointer
    	fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result.Choices[0].Text)
   }
}

