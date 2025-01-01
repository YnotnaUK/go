package twitch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	authorizationURL string = "https://id.twitch.tv/oauth2/authorize"
	codeExchangeURL  string = "https://id.twitch.tv/oauth2/token"
	httpAgentName    string = "https://github.com/YnotnaUK/go/pkg/twitch - v1.0"
	validationURL    string = "https://id.twitch.tv/oauth2/validate"
)

var (
	ErrInvalidJSONFile error = errors.New("file is not a valid json file")
)

func CreateAccessTokenFromJSONFile(filePath string) (*AccessToken, error) {
	// Ensure the file ends in .json
	if filepath.Ext(filePath) != ".json" {
		return nil, ErrInvalidJSONFile
	}
	// Read file contents
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// Marshall file contents into blank struct
	accessToken := &AccessToken{}
	err = json.Unmarshal(fileContents, accessToken)
	if err != nil {
		return nil, err
	}
	// Return access token
	return accessToken, nil
}

func ExchangeCode(
	clientId string,
	clientSecret string,
	code string,
	redirectUri string,
) (*AccessTokenExternal, error) {
	requestBody := []byte(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code&redirect_uri=%s",
		clientId,
		clientSecret,
		code,
		redirectUri,
	))
	// Create a http request
	httpRequest, err := http.NewRequest("POST", codeExchangeURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	// Update request headers
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpRequest.Header.Add("User-Agent", httpAgentName)
	// Create http client
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	// Process the request
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	// Check response
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", httpResponse.StatusCode, httpResponse.Status)
	}
	// Unmarshal valid JSON response into struct
	accessTokenExternal := &AccessTokenExternal{}
	err = json.NewDecoder(httpResponse.Body).Decode(accessTokenExternal)
	if err != nil {
		return nil, err
	}
	// Return access token
	return accessTokenExternal, nil
}

func GenerateAuthorizationUrl(clientId string, redirect_uri string, scopes []string, state string) (string, error) {
	responseType := "code"
	scopeString := ""
	for scopeIndex, scope := range scopes {
		if scopeIndex == 0 {
			scopeString = scope
		} else {
			scopeString = scopeString + "+" + scope
		}
	}
	// Build URL based on inputs
	completeUrl := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s",
		authorizationURL,
		clientId,
		redirect_uri,
		responseType,
		scopeString,
		state,
	)
	// Return complete URL
	return completeUrl, nil
}

func ValidateAccessToken(accessToken string) (*AccessTokenValidationExternal, error) {
	// Create a http request
	httpRequest, err := http.NewRequest("GET", validationURL, nil)
	if err != nil {
		return nil, err
	}
	// Update request headers
	httpRequest.Header.Add("Authorization", fmt.Sprintf("OAuth %s", accessToken))
	httpRequest.Header.Add("User-Agent", httpAgentName)
	// Create http client
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	// Process the request
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	// Check response
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d - %s", httpResponse.StatusCode, httpResponse.Status)
	}
	// Unmarshal valid JSON response into struct
	accessTokenValidationExternal := &AccessTokenValidationExternal{}
	err = json.NewDecoder(httpResponse.Body).Decode(accessTokenValidationExternal)
	if err != nil {
		return nil, err
	}
	// Return validation
	return accessTokenValidationExternal, nil
}
