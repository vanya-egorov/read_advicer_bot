package telegram

import "net/http"

type CLient struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) Client {
	return CLient{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {

	return "bot" + token
}

func (c *CLient) Updates() {

}

func (c *Client) SendMessage() {

}
