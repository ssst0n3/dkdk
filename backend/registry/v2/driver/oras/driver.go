package oras

import (
	"context"
	"fmt"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/remotes"
	"github.com/davecgh/go-spew/spew"
	"github.com/deislabs/oras/pkg/auth"
	orasDocker "github.com/deislabs/oras/pkg/auth/docker"
	"github.com/deislabs/oras/pkg/content"
	"github.com/deislabs/oras/pkg/oras"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/dkdk/registry/v2/service_provider/huaweicloud"
	"net/http"
	"sync"
)

type Driver struct {
	Hostname        string `json:"hostname"`
	Username        string `json:"username"`
	Secret          string `json:"secret"`
	InSecure        bool   `json:"secure"`
	CustomMediaType string `json:"custom_media_type"`
	Context         context.Context
	client          auth.Client
	Resolver        remotes.Resolver
	getResolverFunc func(resolver remotes.Resolver) remotes.Resolver
}

func NewDriver(hostname, username, secret string, insecure bool, customMediaType string, getResolverFunc func(resolver remotes.Resolver) remotes.Resolver) (driver Driver, err error) {
	client, err := orasDocker.NewClient()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	driver = Driver{
		Hostname:        hostname,
		Username:        username,
		Secret:          secret,
		InSecure:        insecure,
		CustomMediaType: customMediaType,
		Context:         context.Background(),
		client:          client,
		getResolverFunc: getResolverFunc,
	}
	return
}

func (d *Driver) Login() (err error) {
	client, err := orasDocker.NewClient()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = client.Login(d.Context, d.Hostname, d.Username, d.Secret, d.InSecure)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if d.Resolver == nil {
		resolver, err := d.client.Resolver(d.Context, http.DefaultClient, d.InSecure)
		if err != nil {
			awesome_error.CheckErr(err)
			return err
		}
		if d.getResolverFunc == nil {
			d.Resolver = resolver
		} else {
			d.Resolver = huaweicloud.NewResolver(resolver)
		}
	}
	return
}

func (d Driver) List(ref string) (err error) {
	//d.resolver.Fetcher()
	name, desc, err := d.Resolver.Resolve(d.Context, ref)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	spew.Dump(name)
	spew.Dump(desc)
	return
}

func (d Driver) Upload(ref, fileName string, fileContent []byte) (err error) {
	//resolver := docker.NewResolver(docker.ResolverOptions{})
	memoryStore := content.NewMemoryStore()
	desc := memoryStore.Add(fileName, d.CustomMediaType, fileContent)
	pushContents := []ocispec.Descriptor{
		desc,
		//descM,
	}
	log.Logger.Infof("Pushing %s to %s...", fileName, ref)
	desc, err = oras.Push(d.Context, d.Resolver, ref, memoryStore, pushContents) //oras.WithConfigMediaType("application/vnd.docker.container.image.v1+json"),

	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	log.Logger.Infof("Pushed to %s with digest %s\n", ref, desc.Digest)
	return
}

func (d Driver) Download(ref, fileName string) (err error) {
	ctx := context.Background()
	//resolver := docker.NewResolver(docker.ResolverOptions{})

	// Pull file(s) from registry and save to disk
	log.Logger.Infof("Pulling from %s and saving to %s...\n", ref, fileName)
	fileStore := content.NewFileStore("/tmp/")
	fileStore.DisableOverwrite = false
	fileStore.AllowPathTraversalOnWrite = true
	defer fileStore.Close()
	//allowedMediaTypes := []string{d.CustomMediaType}
	desc, descs, err := oras.Pull(ctx, d.Resolver, ref, fileStore, oras.WithPullCallbackHandler(pullStatusTrack()))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	spew.Dump(desc)
	spew.Dump(descs)
	log.Logger.Infof("Pulled from %s with digest %s\n", ref, desc.Digest)
	log.Logger.Infof("Try running 'cat %s'\n", fileName)
	return
}

func pullStatusTrack() images.Handler {
	var printLock sync.Mutex
	return images.HandlerFunc(func(ctx context.Context, desc ocispec.Descriptor) ([]ocispec.Descriptor, error) {
		if name, ok := content.ResolveName(desc); ok {
			digestString := desc.Digest.String()
			if err := desc.Digest.Validate(); err == nil {
				if algo := desc.Digest.Algorithm(); algo == digest.SHA256 {
					digestString = desc.Digest.Encoded()[:12]
				}
			}
			printLock.Lock()
			defer printLock.Unlock()
			fmt.Println("Downloaded", digestString, name)
		}
		return nil, nil
	})
}
