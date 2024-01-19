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

func (r NullResource[S, V, P]) Check(ctx context.Context, request concourse.CheckRequest[Source, Version], response *concourse.CheckResponse[Version], log io.Writer) error {
	return nil
}

func (r NullResource[S, V, P]) Get(ctx context.Context, request concourse.GetRequest[Source, Version, Params], response *concourse.Response[Version], log io.Writer, destination string) error {
	return nil
}

func (r NullResource[S, V, P]) Put(ctx context.Context, request concourse.PutRequest[Source, Params], response *concourse.Response[Version], log io.Writer, source string) error {
	return nil
}

type Troublemaker[S Source, V Version, P Params] struct{}

func (r Troublemaker[S, V, P]) Check(ctx context.Context, request concourse.CheckRequest[Source, Version], response *concourse.CheckResponse[Version], log io.Writer) error {
	*response = append(*response, Version{Number: 77})
	return nil
}

func (r Troublemaker[S, V, P]) Get(ctx context.Context, request concourse.GetRequest[Source, Version, Params], response *concourse.Response[Version], log io.Writer, destination string) error {
	return nil
}

func (r Troublemaker[S, V, P]) Put(ctx context.Context, request concourse.PutRequest[Source, Params], response *concourse.Response[Version], log io.Writer, source string) error {
	return nil
}
