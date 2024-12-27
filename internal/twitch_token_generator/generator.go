package twitchtokengenerator

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ynotnauk/go/pkg/twitch"
)

var (
	ErrBlankClientId error = errors.New("clientId cannot be blank")
)

type Generator struct {
	clientId     string
	redirect_uri string
	router       chi.Router
	scopes       []string
	templates    map[string]*template.Template
}

func (g *Generator) loadTemplates() error {
	files := []string{
		"templates/token-generator/index.html",
		"templates/token-generator/callback.html",
	}
	for _, file := range files {
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			return err
		}
		g.templates[tmpl.Name()] = tmpl
		fmt.Println("loaded template", tmpl.Name())
	}
	return nil
}

func (g *Generator) handleIndex(w http.ResponseWriter, r *http.Request) {
	authorizationUrl, err := twitch.GenerateAuthorizationUrl(g.clientId, g.redirect_uri, g.scopes)
	if err != nil {
		fmt.Println(err)
	}
	err = g.templates["index.html"].Execute(w, &IndexTemplateData{
		AuthURL: authorizationUrl,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (g *Generator) twitchCallback(w http.ResponseWriter, r *http.Request) {
	err := g.templates["callback.html"].Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
	queryParts := r.URL.Query()
	fmt.Println(queryParts.Get("code"))
}

func (g *Generator) Start() {
	g.loadTemplates()
	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	http.ListenAndServe(port, g.router)
}

func NewGenerator(clientId string, redirect_uri string, scopes []string) (*Generator, error) {
	// Ensure clientId is not blank
	if clientId == "" {
		return nil, ErrBlankClientId
	}
	// Create Chi Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	generator := &Generator{
		clientId:     clientId,
		redirect_uri: redirect_uri,
		router:       router,
		scopes:       scopes,
		templates:    make(map[string]*template.Template),
	}
	router.Get("/", generator.handleIndex)
	router.Get("/twitch/bot/generate/tokens/callback", generator.twitchCallback)
	return generator, nil
}

/*
	func generateAuthorizeUrl(clientId string, redirectUri string, state string) (string, error) {
	url := "https://id.twitch.tv/oauth2/authorize"
	scopeString := ""
	for i, twitchScope := range twitchScopes {
		if i == 0 {
			scopeString = twitchScope
		} else {
			scopeString = scopeString + "+" + twitchScope
		}
	}
	completeUrl := url + "?client_id=" + clientId + "redirect_uri=" + redirectUri + "&response_type=code" + "&scope=" + scopeString + "&state=" + state
	return completeUrl, nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./templates/token-generator/index.html")
	if err != nil {
		fmt.Println(err)
	}
	// Build template data
	clientId := os.Getenv("TWITCH_CLIENT_ID")
	redirectUri := "http://localhost:8080/twitch/bot/generate/tokens/callback"
	// responseType := "code"
	// tmplData := Config{
	// 	ResponseType: responseType,
	// 	ClientId:     clientId,
	// }
	x, err := generateAuthorizeUrl(clientId, redirectUri, "12345")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("scope string", x)
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
*/
