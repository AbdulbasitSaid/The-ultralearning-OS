package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Prototype:
func demo() {
	resp, err := http.Get("https://api.github.com/users/abdulbasitsaid")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)

	var reply struct {
		Name                string
		NumberOfPublicRepos int `json:"public_repos"`
	}

	// io.Copy(os.Stdout, resp.Body)

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(reply.Name, reply.NumberOfPublicRepos)
	//exercise: try to change the request header.
}

// / Given a github user login, return name and number of public repos
func main() {

	fmt.Println(UserInfo("abdulbasitsaid"))

}

/// clear example

/// User info return name and number of public repositories

func UserInfo(login string) (name string, repo_count int, err error) {
	url := "https://api.github.com/users/" + login

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)

}

func parseResponse(readCloser io.ReadCloser) (string, int, error) {
	var responseData struct {
		Name                string
		NumberOfPublicRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(readCloser)

	if err := dec.Decode(&responseData); err != nil {
		return "", 0, fmt.Errorf("%q - error decoding json %s", "", err)
	}
	return responseData.Name, responseData.NumberOfPublicRepos, nil
}
