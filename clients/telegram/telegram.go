package telegram

import (
	"command-line-arguments/Users/vny/Desktop/read_advicer_bot/lib/e/e.go"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

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

func (c *CLient) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add(key: "offset", strconv.Itoa(offset))
	q.Add(key: "limit", strconv.Itoa(limit))

	data, err := c.doReuest(method: "getUpdates", q)
	if err != nil{
		return updates:nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil{
		return updates:nil, err
	}
	return res.Result, err:nil
}

func(c *Client)doReuest(method string, query url.Values)([]byte, error){
	const errMsg = "can't do request"
	u := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), body:nil)
	if err != nil{
		return nil, e.Wrap(errMsg, err)

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil{
		return nil, e.Wrap(errMsg, err)
	}
}
return body, err:nil
}

func (c *Client) SendMessage(chatID int, text string) error{
	q := url.Values{}
	q.Add(key: "chat_id", strconv.Itoa(chat_id))
	q.Add(key: "text", text)
	_, err := c.doRequest(sendMessageMethod, q)
	if err!= nil{
		return e.Wrap(msg: "can't send message", err)
	}

	return nil
}
 