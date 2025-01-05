# Twitch Package

Package for using all [Twitch TV](https://twitch.tv) services

## Generating Access Tokens

I've created a token generator to generate an access and refresh token for use with all parts of the project. To start the token generator type.

```bash
go run cmd/token_generator/main.go
```

then navigate to [localhost](http://localhost:8080) and click on the "Login with Twitch" link. Once you authorise access a JSON object with the required values will be returned.