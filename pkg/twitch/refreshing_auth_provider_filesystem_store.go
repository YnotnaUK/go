package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrBlankStoreLocation error = errors.New("storeLocation cannot be blank")
)

type RefreshingAuthProviderFilesystemStore struct {
	storeLocation string
}

func (s *RefreshingAuthProviderFilesystemStore) buildUserTokensPath(userId string) string {
	return fmt.Sprintf("%s/tokens.%s.json", s.storeLocation, userId)
}

func (s *RefreshingAuthProviderFilesystemStore) GetTokensForUserId(userId string) (*UserTokens, error) {
	// build store path
	storeFilePath := s.buildUserTokensPath(userId)
	// Attempt to read the file
	contents, err := os.ReadFile(storeFilePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read token file: %s", storeFilePath)
	}
	// Create struct for auth
	userTokens := &UserTokens{}
	// Write store to struct
	err = json.Unmarshal(contents, userTokens)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal json: %s", storeFilePath)
	}
	return userTokens, nil
}

func (s *RefreshingAuthProviderFilesystemStore) SaveTokensForUserId(userId string, tokens *UserTokens) (*UserTokens, error) {
	// build store path
	storeFilePath := s.buildUserTokensPath(userId)
	// create file contents from tokens
	fileContents, _ := json.MarshalIndent(tokens, "", "  ")
	err := os.WriteFile(storeFilePath, fileContents, 0644)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func NewRefreshingAuthProviderFilesystemStore(storeLocation string) (*RefreshingAuthProviderFilesystemStore, error) {
	// Ensure store location is not blank
	if storeLocation == "" {
		return nil, ErrBlankStoreLocation
	}
	// Ensure the path does not have a trailing slash
	storeLocation = strings.TrimSuffix(storeLocation, "/")
	// Create store
	store := &RefreshingAuthProviderFilesystemStore{
		storeLocation: storeLocation,
	}
	// return store
	return store, nil
}
