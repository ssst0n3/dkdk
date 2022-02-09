package client

import (
	"bytes"
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api/example/resource/auth"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"io/ioutil"
	"net/http"
)

const headerToken = "token"

type Transport struct {
	Transport http.RoundTripper
	URL       string
	Username  string
	Password  string
	Token     string
}

func NewTransport(url, username, password string) *Transport {
	return &Transport{
		Transport: http.DefaultTransport,
		URL:       url,
		Username:  username,
		Password:  password,
	}
}

func Login(url, username, password string) (token string, err error) {
	loginModel := model.LoginModel{
		Username: username,
		Password: password,
	}
	data, err := json.Marshal(loginModel)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	body := bytes.NewReader(data)
	resp, err := http.Post(url, contentTypeJson, body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	var loginSuccessResponse auth.LoginSuccessResponse
	err = json.Unmarshal(content, &loginSuccessResponse)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	token = loginSuccessResponse.Token
	return
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	if req.Header.Get(headerToken) == "" {
		if t.Username != "" || t.Password != "" {
			if t.Token == "" {
				t.Token, err = Login(t.URL, t.Username, t.Password)
				if err != nil {
					return
				}
			}
			req.Header.Set(headerToken, t.Token)
		}
	}
	resp, err = t.Transport.RoundTrip(req)
	return
}
