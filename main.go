package main

import (
	"bufio"
	"fmt"
	"errors"
	"log"
	"os"
	"strings"

	//"github.com/zmb3/spotify"
)

func readCredentialsFromFile(filename string) (clientID, clientSecret string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Client ID:") {
			clientID = strings.TrimPrefix(line, "Client ID:")
		} else if strings.HasPrefix(line, "Client Secret:") {
			clientSecret = strings.TrimPrefix(line, "Client Secret:")
		}
	}
	if clientID == "" || clientSecret == "" {
		missingInfo := make([]string, 0)
		if clientID == "" {
			missingInfo = append(missingInfo, "Client ID")
		}
		if clientSecret == "" {
			missingInfo = append(missingInfo, "Client Secret")
		}
		return "", "", errors.New("Missing information in the file: " + strings.Join(missingInfo, ", "))
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	return clientID, clientSecret, nil
}

func main() {
	clientID, clientSecret, err := readCredentialsFromFile("Login.txt")
	if err != nil {
		log.Fatal(err)
	}

	//auth := spotify.NewAuthenticator(clientID, spotify.ScopeUserReadPrivate, spotify.ScopeUserReadPlaybackState)
	//auth.SetAuthInfo(clientID, clientSecret)
	fmt.Println("Client ID:",clientID);
	fmt.Println("Client Secret:",clientSecret);

	// Rest of your authentication and API setup code...
}
