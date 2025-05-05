package telegram

import (
	"encoding/json"
	"github.com/AlexLuminare/read_advisor_bot/lib/e"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	errMsg            = "can't do request"
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   *http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   &http.Client{},
	}
}

func (c *Client) SendMessage(chat_id int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chat_id))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)

	return e.WrapIfErr("can't send message", err)
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	//do request <- getUpdates
	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}
	var res UpdatesResponce
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.Result, err
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {

	//формиреуем url
	u := url.URL{
		Scheme: "http",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	// подготовка запроса
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}

	//добавляем query-параметры в запрос
	req.URL.RawQuery = query.Encode()

	//отправка запроса
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap(errMsg, err)
	}
	return body, nil
}

func newBasePath(token string) string {
	return "bot" + token
}
