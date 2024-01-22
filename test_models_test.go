package concourse_test

import (
	"context"
	"io"

	"github.com/suhlig/concourse-resource-go"
)

type Source struct {
	URL string `validate:"required,http_url"`
}

type Version struct {
	Number uint64 `validate:"required,max=11"`
}

type Params struct{}

type NullResource[S Source, V Version, P Params] struct{}

func (r NullResource[S, V, P]) Check(ctx context.Context, request concourse.CheckRequest[Source, Version], log io.Writer) (concourse.CheckResponse[Version], error) {
	return concourse.CheckResponse[Version]{}, nil
}

func (r NullResource[S, V, P]) Get(ctx context.Context, request concourse.GetRequest[Source, Version, Params], log io.Writer, destination string) (*concourse.Response[Version], error) {
	return &concourse.Response[Version]{}, nil
}

func (r NullResource[S, V, P]) Put(ctx context.Context, request concourse.PutRequest[Source, Params], log io.Writer, source string) (*concourse.Response[Version], error) {
	return &concourse.Response[Version]{}, nil
}

type Troublemaker[S Source, V Version, P Params] struct{}

func (r Troublemaker[S, V, P]) Check(ctx context.Context, request concourse.CheckRequest[Source, Version], log io.Writer) (concourse.CheckResponse[Version], error) {
	return concourse.CheckResponse[Version]{Version{Number: 77}}, nil
}

func (r Troublemaker[S, V, P]) Get(ctx context.Context, request concourse.GetRequest[Source, Version, Params], log io.Writer, destination string) (*concourse.Response[Version], error) {
	return &concourse.Response[Version]{}, nil
}

func (r Troublemaker[S, V, P]) Put(ctx context.Context, request concourse.PutRequest[Source, Params], log io.Writer, source string) (*concourse.Response[Version], error) {
	return &concourse.Response[Version]{}, nil
}
