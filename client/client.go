package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	v1 "github.com/ssst0n3/dkdk/api/v1"
	"github.com/ssst0n3/dkdk/api/v1/repository"
	"github.com/ssst0n3/dkdk/model"
	"github.com/ssst0n3/lightweight_api/example/resource/auth"
	"io/ioutil"
	"net/http"
	"net/url"
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
	u := fmt.Sprintf("%s://%s%s", c.protocol, c.Domain, repository.Resource.BaseRelativePath)
	resp, err := c.Client.Get(u)
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

func (c *DkdkClient) BatchTaskCreate(tasks []model.TaskCore) (err error) {
	u := fmt.Sprintf("%s://%s%s/batch/create", c.protocol, c.Domain, v1.TaskResource.BaseRelativePath)
	content, err := json.Marshal(tasks)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	body := bytes.NewReader(content)
	_, err = c.Client.Post(u, "text/json", body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func (c *DkdkClient) CheckFileAlreadyExists(filename string) (exists bool, err error) {
	route := fmt.Sprintf("%s://%s%s", c.protocol, c.Domain, v1.FileResource.BaseRelativePath)
	u, err := url.Parse(route)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	query := u.Query()
	query.Add("filename", filename)
	u.RawQuery = query.Encode()
	resp, err := c.Client.Get(u.String())
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	var body model.ResponseCheckFilenameAlreadyExists
	err = json.Unmarshal(content, &body)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	exists = body.Exists
	return
}
