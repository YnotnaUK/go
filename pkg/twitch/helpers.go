package twitch

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

const httpAgentName string = "TwitchBot v1.0"

/*
	https://id.twitch.tv/oauth2/authorize
    ?response_type=code
    &client_id=hof5gwx0su6owfnys0nyan9c87zr6t
    &redirect_uri=http://localhost:3000
    &scope=channel%3Amanage%3Apolls+channel%3Aread%3Apolls
    &state=c3ab8aa609ea11e793ae92361f002671
*/

func GenerateAuthorizationUrl(clientId string, redirect_uri string, scopes []string) (string, error) {
	authUrl := "https://id.twitch.tv/oauth2/authorize"
	responseType := "code"
	scopeString := ""
	for scopeIndex, scope := range scopes {
		if scopeIndex == 0 {
			scopeString = scope
		} else {
			scopeString = scopeString + "+" + scope
		}
	}
	state := "1234567890"
	completeUrl := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s", authUrl, clientId, redirect_uri, responseType, scopeString, state)
	return completeUrl, nil
}

func Exchange(clientId string, clientSecret string, code string, redirectUri string) (string, error) {
	//exchangeUrl := "https://id.twitch.tv/oauth2/token"
	exchangeUrl := "http://locaohst:8080/test/post"

	requestBody := []byte(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s0&grant_type=authorization_code&redirect_uri=%s",
		clientId,
		clientSecret,
		code,
		redirectUri,
	))

	// Create a http request
	httpRequest, err := http.NewRequest("POST", exchangeUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// Update headers
	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("User-Agent", httpAgentName)

	// Create http client
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", nil
	}

	defer httpResponse.Body.Close()

	return "", nil

	/*
		client_id=hof5gwx0su6owfnys0yan9c87zr6t
		&client_secret=41vpdji4e9gif29md0ouet6fktd2
		&code=gulfwdmys5lsm6qyz4xiz9q32l10
		&grant_type=authorization_code
		&redirect_uri=http://localhost:3000
	*/
}
