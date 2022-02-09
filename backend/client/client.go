package client

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_api/example/resource/auth"
	"io/ioutil"
	"net/http"
)

const contentTypeJson = "application/json"

type DkdkClient struct {
	Client   *http.Client
	Domain   string
	Username string
	Password string
	Token    string
	Insecure bool
	protocol string
}

func NewDkdkClient(domain, username, password string, insecure bool) (client *DkdkClient) {
	protocol := "http"
	if !insecure {
		protocol += "s"
	}
	authUrl := fmt.Sprintf("%s://%s%s", protocol, domain, auth.Resource.BaseRelativePath)
	return &DkdkClient{
		Client: &http.Client{
			Timeout:   0,
			Transport: NewTransport(authUrl, username, password),
		},
		Domain:   domain,
		Username: username,
		Password: password,
		Insecure: insecure,
		protocol: protocol,
	}
}

func (c *DkdkClient) RepositoryList() (err error) {
	url := fmt.Sprintf("%s://%s%s", c.protocol, c.Domain, "/api/v1/repository")
	resp, err := c.Client.Get(url)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	//var repositoryConfigResponse []model.RepositoryConfigResponse
	log.Logger.Info(string(content))
	return
}
