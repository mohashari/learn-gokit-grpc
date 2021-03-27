package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mohashari/learn-gokit-grpc/service"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

type MathReq struct {
	NumA float32
	NumB float32
}

type MathResp struct {
	Result float32
}

func MakeEndpoint(s service.Service) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}
