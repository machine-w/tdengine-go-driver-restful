package taossqlrestful

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	Authorization = "Authorization"
	Successful    = "succ"
)

//Wrapper a simple wrapper for http request
type Wrapper struct {
	token       string
	reqHandler  []func(req *http.Request)
	respHandler []func(*TaosResp) error
}

func NewWrapper(token string) *Wrapper {
	w := &Wrapper{
		token:       token,
		reqHandler:  make([]func(req *http.Request), 0),
		respHandler: make([]func(*TaosResp) error, 0)}
	w.reqHandler = append(w.reqHandler, func(req *http.Request) {
		req.Header.Add(Authorization, generateToken(token))
	})
	w.respHandler = append(w.respHandler, func(resq *TaosResp) error {
		if resq.Status != Successful {
			return errors.New(fmt.Sprintf("return %s", resq.Status))
		}
		return nil
	})
	return w
}

// Post post wrapper
func (w *Wrapper) Post(url string, body io.Reader) (*TaosResp, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return w.Do(request)
}

func (w *Wrapper) Do(req *http.Request) (*TaosResp, error) {
	for _, reqHandler := range w.reqHandler {
		reqHandler(req)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("taos restful api fail! resp is nill ")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tResp := new(TaosResp)
	err = json.Unmarshal(content, &tResp)
	if err != nil {
		return nil, err
	}

	for _, handler := range w.respHandler {
		err := handler(tResp)
		if err != nil {
			return nil, err
		}
	}
	return tResp, nil
}

func generateToken(token string) string {
	return fmt.Sprintf("Basic %s", token)
}