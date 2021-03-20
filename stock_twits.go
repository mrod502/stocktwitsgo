package stocktwits

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mrod502/logger"
)

func processResponseHeaders(h http.Header) (remainingRequests, waitUntil int64, err error) {

	s := h["X-Ratelimit-Remaining"]
	if len(s) > 0 {
		remainingRequests, err = strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return
		}
	}

	s = h["X-Ratelimit-Reset"]
	if len(s) > 0 {
		waitUntil, err = strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			return
		}
	}

	return
}

func getTrending() (msg []Message, remainingCalls, callReset int64, err error) {
	res, err := cli.Get(trendingURL)
	var data Response
	if err != nil {
		return
	}
	remainingCalls, callReset, err = processResponseHeaders(res.Header)

	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.Unmarshal(b, &data)

	if err != nil {
		return
	}
	msg = data.Messages

	if len(data.Errors) > 0 {
		// make the process wait for a while before sending another request
		err = errors.New(data.Errors[0].Message)
		return
	}

	return
}

func getSuggested() (msg []Message, remainingCalls, callReset int64, err error) {
	res, err := cli.Get(suggestedURL)
	var data Response
	if err != nil {
		return
	}
	remainingCalls, callReset, err = processResponseHeaders(res.Header)

	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.Unmarshal(b, &data)

	if err != nil {
		return
	}
	msg = data.Messages

	if len(data.Errors) > 0 {
		// make the process wait for a while before sending another request
		err = errors.New(data.Errors[0].Message)
		return
	}

	return
}

//TrendingStream -- passes stocktwits posts to chan c once every refreshInterval
//set refreshInterval to more than 20 Seconds or will have to wait for limit count to reset.
func TrendingStream(c chan Message, refreshInterval time.Duration) {

	for {
		msg, calls, reset, err := getTrending()

		for _, v := range msg {
			c <- v
		}
		if calls == 0 && reset > time.Now().Unix() {
			time.Sleep(time.Duration((reset-time.Now().Unix())+1) * time.Second)
			continue
		}

		if err != nil {
			logger.Error("StockTwit", err.Error())
		}
		time.Sleep(refreshInterval)
	}
}

//SuggestedStream -- passes stocktwits posts to chan c once every refreshInterval
//set refreshInterval to more than 20 Seconds or will have to wait for limit count to reset.
func SuggestedStream(c chan []Message, refreshInterval time.Duration) {

	for {
		msg, calls, reset, err := getSuggested()

		c <- msg

		if calls == 0 && reset > time.Now().Unix() {
			time.Sleep(time.Duration((reset-time.Now().Unix())+1) * time.Second)
			continue
		}

		if err != nil {
			logger.Error("StockTwit", err.Error())
		}
		time.Sleep(refreshInterval)
	}
}
