package stocktwits

import (
	"fmt"
	"net/http"
	"time"
)

const (
	dtFormatString = "2006-01-02T15:04:05Z"
)

var (
	cli          = http.DefaultClient
	trendingURL  = fmt.Sprintf("https://api.stocktwits.com/api/2/streams/trending.json?access_token=%v", accessToken)
	suggestedURL = fmt.Sprintf("https://api.stocktwits.com/api/2/streams/suggested.json?access_token=%v", accessToken)
)

type oauthRequestBody struct {
	ClientID     string `json:"client_id"`
	ResponseType string `json:"response_type"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	Prompt       string `json:"prompt"`
}

type oauthCert struct {
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	Username    string `json:"username"`
}

//Response - generalized object to unmarshal responses from stocktwits API calls
type Response struct {
	Response struct {
		Status int `json:"status"`
	} `json:"response"`
	Errors []ResError `json:"errors"`

	Cursor struct {
		More  bool `json:"more"`
		Since int  `json:"since"`
		Max   int  `json:"max"`
	}
	Messages []Message `json:"messages"`
}

//ResError - error messages in API call response
type ResError struct {
	Message string `json:"message"`
}

//Message - stocktwits post
type Message struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
	Source    struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		URL   string `json:"url"`
	}
	Symbols  []Symbol `json:"symbols"`
	Entities struct {
		Sentiment `json:"sentiment"`
	} `json:"entities"`
	Links []Link `json:"links"`
}

//User - user info
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Identity string `json:"identity"`
}

//Symbol - symbol info
type Symbol struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Title  string `json:"title"`
}

//Sentiment - bullish/bearish/none
type Sentiment struct {
	Basic string `json:"basic"`
}

//Link - links contained in post
type Link struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

//CreatedUnix - Unix timestamp of creation of message
func (m Message) CreatedUnix() (t int64) {

	tt, _ := time.Parse(dtFormatString, m.CreatedAt)
	t = tt.Unix()
	return
}

//GetSymbols - return []string of ticker symbols tagged in post
func (m Message) GetSymbols() (s []string) {
	for _, v := range m.Symbols {
		s = append(s, v.Symbol)
	}
	return s
}

//GetSentiment - returns "Bullish", "Bearish", or ""
func (m Message) GetSentiment() string {
	return m.Entities.Sentiment.Basic
}
