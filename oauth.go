package stocktwits

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrod502/logger"
)

const (
	consumerKey        = "de91a29094a23527"
	consumerSecret     = "8f1eea304f900077c737f84444a3f882acff256b"
	pw                 = "N3xt@lpha33"
	uname              = "stonksguy80808"
	scope              = "read,watch_lists,publish_messages,follow_users,follow_stocks"
	apiURL             = "https://api.stocktwits.com/api/2"
	oauthURL           = apiURL + "/oauth"
	codeURL            = oauthURL + "/authorize?client_id=%v&response_type=token&redirect_uri=%v&scope=%v"
	tokenURL           = oauthURL + ""
	defaultRedirectURI = "http://www.stocktwits.com"
	accessToken        = "5520998f7eec1a96d2ba707cd5cfca729391a1c2"
)

func requestCode(clientID, redirectURI, scope string) (code string, err error) {
	var res *http.Response

	reqURL := fmt.Sprintf(codeURL, clientID, redirectURI, scope)

	request, err := http.NewRequest("GET", reqURL, nil)
	request.SetBasicAuth(uname, pw)
	if err != nil {
		return
	}

	res, err = cli.Do(request)

	if err != nil {
		logger.Error("Stocktwit", "auth", err.Error())
		return
	}
	b, _ := ioutil.ReadAll(res.Body)

	ioutil.WriteFile("code_response.html", b, 0644)

	return
}

func oauthVerify() (token string, ok bool) {

	return
}
