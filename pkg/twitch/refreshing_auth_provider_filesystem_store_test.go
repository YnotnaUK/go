package twitch

import (
	"strings"
	"testing"
)

func TestNewStoreWithBlankLocation(t *testing.T) {
	_, err := NewRefreshingAuthProviderFilesystemStore("")
	if err != ErrBlankStoreLocation {
		t.Fatal(err)
	}
}

func TestStoreWithInvalidUserId(t *testing.T) {
	userId := "44444"
	store, err := NewRefreshingAuthProviderFilesystemStore("../../examples/twitch/data/token_store")
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetTokensForUserId(userId)
	if err != nil {
		message := err.Error()
		if !strings.HasPrefix(message, "cannot read token file: ") {
			t.Fatal(err)
		}
	}
}

func TestStoreWithInvalidJSON(t *testing.T) {
	userId := "invalid"
	store, err := NewRefreshingAuthProviderFilesystemStore("../../examples/twitch/data/token_store")
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.GetTokensForUserId(userId)
	if err != nil {
		message := err.Error()
		if !strings.HasPrefix(message, "cannot unmarshal json: ") {
			t.Fatal(err)
		}
	}
}
func TestValidStore(t *testing.T) {
	userId := "123456789"
	store, err := NewRefreshingAuthProviderFilesystemStore("../../examples/twitch/data/token_store")
	if err != nil {
		t.Fatal(err)
	}
	tokens, err := store.GetTokensForUserId(userId)
	if err != nil {
		t.Fatal(err)
	}
	store.SaveTokensForUserId(userId, tokens)
}
