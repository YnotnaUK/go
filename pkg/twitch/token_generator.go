package twitch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

var (
	ErrBlankClientId            error = errors.New("client Id cannot be blank")
	ErrBlankClientSecret        error = errors.New("client Secret cannot be blank")
	ErrBlankPort                error = errors.New("port cannot be 0")
	ErrBlankRedirectURI         error = errors.New("redirect uri cannot be blank")
	ErrHttpServerAlreadyRunning error = errors.New("http server has already been created")
	ErrNoScopes                 error = errors.New("scopes cannot be empty")
)

type TokenGeneratorConfig struct {
	ClientId     string
	ClientSecret string
	Port         int
	RedirectURI  string
	Scopes       []string
}

type TokenGenerator struct {
	clientId     string
	clientSecret string
	httpServer   *http.Server
	port         int
	redirectUri  string
	scopes       []string
	state        string
	wg           *sync.WaitGroup
}

func (tg *TokenGenerator) handleCallback(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	code := urlValues.Get("code")
	errorV := urlValues.Get("error")
	state := urlValues.Get("state")
	// Ensure state matches
	fmt.Printf("Checking state %s vs %s\n", state, tg.state)
	if state != tg.state {
		fmt.Fprintf(w, "Invalid state value: %s", state)
		return
	}
	// Check for error
	fmt.Printf("Checking error %s\n", errorV)
	if errorV != "" {
		fmt.Fprintf(w, "Error: %s", errorV)
		return
	}
	// Ensure code is not blank
	if code == "" {
		fmt.Fprintf(w, "Missing code: %s", code)
		return
	}
	// Exchange code for access token
	accessTokenExternal, err := ExchangeCode(tg.clientId, tg.clientSecret, code, tg.redirectUri)
	if err != nil {
		fmt.Fprintf(w, "Code Exchange Error: %s", err)
		return
	}
	// Validate token to get additonal information
	accessTokenValidation, err := ValidateAccessToken(accessTokenExternal.AccessToken)
	if err != nil {
		fmt.Fprintf(w, "Token Validation Error: %s", err)
	}
	// Create access token
	accessToken := &AccessToken{
		AccessToken:  accessTokenExternal.AccessToken,
		ExpiresIn:    accessTokenValidation.ExpiresIn,
		Login:        accessTokenValidation.Login,
		RefreshToken: accessTokenExternal.RefreshToken,
		Scopes:       accessTokenValidation.Scopes,
		UserId:       accessTokenValidation.UserId,
	}
	// Create pretty JSON object
	output, err := json.MarshalIndent(accessToken, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "Marshall Error: %s", err)
		return
	}
	// Respond
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(output)
	// Shutdown the server after 5 seconds
	go func() {
		time.Sleep(time.Second * 5)
		tg.httpServer.Shutdown(context.TODO())
	}()
}

func (tg *TokenGenerator) handleIndex(w http.ResponseWriter, r *http.Request) {
	loginUrl, err := GenerateAuthorizationUrl(
		tg.clientId,
		tg.redirectUri,
		tg.scopes,
		tg.state,
	)
	if err != nil {
		fmt.Fprintf(w, "Cannot render page: %s", err.Error())
	}
	fmt.Fprintf(w, "<a href=\"%s\">Login with Twitch</a>\n", loginUrl)
}

func (tg *TokenGenerator) Run() {
	fmt.Println("Starting Twitch Token Generator Service...")
	fmt.Printf("Using CliendId: %s\n", tg.clientId)
	fmt.Printf("Using RedirectURI: %s\n", tg.redirectUri)
	fmt.Printf("Using State: %s\n", tg.state)
	tg.startHTTPServer()
	fmt.Printf("Service started on port %d\n", tg.port)
	tg.wg.Wait()
	fmt.Print("Service stopped\n")
}

func (tg *TokenGenerator) startHTTPServer() error {
	if tg.httpServer != nil {
		return ErrHttpServerAlreadyRunning
	}
	// Create chi router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", tg.handleIndex)
	router.Get("/twitch/bot/generate/tokens/callback", tg.handleCallback)
	// Create a http server
	tg.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", tg.port),
		Handler: router,
	}
	// Start server in go routine
	tg.wg.Add(1)
	go func() {
		defer tg.wg.Done()
		if err := tg.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	return nil
}

func NewTokenGenerator(config *TokenGeneratorConfig) (*TokenGenerator, error) {
	// Check config
	if config.ClientId == "" {
		return nil, ErrBlankClientId
	}
	if config.ClientSecret == "" {
		return nil, ErrBlankClientSecret
	}
	if config.Port == 0 {
		return nil, ErrBlankPort
	}
	if config.RedirectURI == "" {
		return nil, ErrBlankRedirectURI
	}
	if len(config.Scopes) <= 0 {
		return nil, ErrNoScopes
	}
	// Create Token Generator
	tokenGenerator := &TokenGenerator{
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		port:         config.Port,
		redirectUri:  config.RedirectURI,
		scopes:       config.Scopes,
		state:        uuid.New().String(),
		wg:           &sync.WaitGroup{},
	}
	// Return generator
	return tokenGenerator, nil
}
