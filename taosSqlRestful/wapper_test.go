package taossqlrestful

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestWapper(t *testing.T) {
	http.HandleFunc("/", handler)
	go func() {
		time.Sleep(5 * time.Second)
		w := NewWrapper("123456")
		resp, err := w.Post("http://localhost:8000", strings.NewReader("SELECT CLIENT_VERSION()"))
		if err != nil {
			log.Print(err.Error())
		} else {
			print(resp.Status)
		}

	}()
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func handler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Header.Get(Authorization))
	body, _ := ioutil.ReadAll(r.Body)

	log.Print(string(body))
	fmt.Fprint(w, "{\n    \"status\": \"err\",\n    \"head\": [\"ts\",\"current\",\"voltage\",\"phase\"],\n    \"column_meta\": [[\"ts\",9,8],[\"current\",6,4],[\"voltage\",4,4],[\"phase\",6,4]],\n    \"data\": [\n        [\"2018-10-03 14:38:05.000\",10.3,219,0.31],\n        [\"2018-10-03 14:38:15.000\",12.6,218,0.33]\n    ],\n    \"rows\": 2\n}")
}
