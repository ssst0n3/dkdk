package model

import (
	"dkdk/model/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceProviderType uint

const (
	HuaweicloudSWR ServiceProviderType = iota
	Aliyun
	DockerRegistry
)

type RepositoryBasic struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

type Cred struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

type Registry struct {
	ServiceAddress string `json:"service_address"`
	Insecure       bool   `json:"insecure"`
}

type Repository struct {
	gorm.Model
	UserId  uint                `json:"user_id"`
	Type    ServiceProviderType `json:"type"`
	Private bool                `json:"private"`
	Registry
	Cred
	RepositoryBasic
}

type RepositoryConfigUpdateCommonBody struct {
	Private  bool   `json:"private"`
	Username string `json:"username"`
	Registry
	RepositoryBasic
}

type RepositoryConfigUpdateSecretBody struct {
	Secret string `json:"secret"`
}

type RepositoryConfigResponse struct {
	Id uint `json:"id"`
	RepositoryConfigUpdateCommonBody
}

var (
	SchemaRepository schema.Schema
)

func init() {
	awesome_error.CheckFatal(util.InitSchema(&SchemaRepository, &Repository{}))
}
