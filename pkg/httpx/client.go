package httpx

type httpClient struct{}

func NewClient() (*httpClient, error) {
	client := &httpClient{}
	return client, nil
}
