package twitch

import (
	"testing"
)

var ValidStorePath string = "../../examples/twitch/data/token_store"

func TestRefreshingAuthProvider(t *testing.T) {
	store, err := NewRefreshingAuthProviderFilesystemStore(ValidStorePath)
	if err != nil {
		t.Fatal(err)
	}
	_, _ = NewRefreshingAuthProvider(store)
}
