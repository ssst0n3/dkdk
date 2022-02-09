package huaweicloud

import (
	"context"
	"github.com/containerd/containerd/remotes"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

type Resolver struct {
	Resolver remotes.Resolver
}

func NewResolver(resolver remotes.Resolver) remotes.Resolver {
	return Resolver{Resolver: resolver}
}

func (r Resolver) Resolve(ctx context.Context, ref string) (name string, desc ocispec.Descriptor, err error) {
	name, desc, err = r.Resolver.Resolve(ctx, ref)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	desc.MediaType = ocispec.MediaTypeImageManifest
	return
}

func (r Resolver) Fetcher(ctx context.Context, ref string) (remotes.Fetcher, error) {
	return r.Resolver.Fetcher(ctx, ref)
}

func (r Resolver) Pusher(ctx context.Context, ref string) (remotes.Pusher, error) {
	return r.Resolver.Pusher(ctx, ref)
}
