package stocktwits

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestResponseStruct(t *testing.T) {
	b, _ := ioutil.ReadFile("trending.json")

	var res Response

	err := json.Unmarshal(b, &res)

	if err != nil {
		t.Fatal(err)
	}

	b, _ = json.Marshal(res)

	ioutil.WriteFile("trending_marshaled.json", b, 0644)
}

func TestTrendingStream(t *testing.T) {

	messages, remaining, _, err := getTrending()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(messages)
	fmt.Println(remaining)
	ioutil.WriteFile("trending.json", b, 0600)
}

func TestExample(t *testing.T) {

	var refreshInterval = 60 * time.Second
	var messageChan = make(chan Message, 64)
	//storage for messages - if there are duplicates, they should be overwritten so we dont store dups
	var messageMap = make(map[int]Message)

	go TrendingStream(messageChan, refreshInterval)

	time.Sleep(time.Second * 3)
	for len(messageChan) > 0 {
		msg := <-messageChan
		messageMap[msg.ID] = msg
		b, _ := json.Marshal(msg)
		fmt.Println(string(b))
	}

}

func TestRemainingCount(t *testing.T) {

	_, rem1, res1, _ := getTrending()
	_, rem2, res2, _ := getSuggested()

	fmt.Println(rem1, rem2)
	fmt.Println(res1, res2)

}
