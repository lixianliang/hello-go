package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// CalculateEndpoint define endpoint
type ArithmeticEndpoints struct {
	CalculateEndpoint   endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
	AuthEndpoint        endpoint.Endpoint
}

type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

// make endpoint
func MakeArithmeticEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)

		var (
			res, a, b int
			calError  error
		)

		a = req.A
		b = req.B
		res, calError = svc.Calculate(ctx, req.RequestType, a, b)
		return ArithmeticResponse{Result: res, Error: calError}, nil
	}
}

type HealthRequest struct{}
type HealthResponse struct {
	Status bool `json:"status"`
}

func MakeHealthCheckEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{status}, nil
	}
}

type AuthRequest struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type AuthResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

func MakeAuthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AuthRequest)
		token, err := svc.Login(req.Name, req.Pwd)
		var resp AuthResponse
		if err != nil {
			resp = AuthResponse{
				Success: false,
				Token:   token,
				Error:   err.Error(),
			}
		} else {
			resp = AuthResponse{
				Success: true,
				Token:   token,
			}
		}
		return resp, nil
	}
}

func (ae ArithmeticEndpoints) Calculate(ctx context.Context, reqType string, a, b int) (res int, err error) {
	resp, err := ae.CalculateEndpoint(ctx, ArithmeticRequest{
		RequestType: reqType,
		A:           a,
		B:           b,
	})
	if err != nil {
		return 0, err
	}
	response := resp.(ArithmeticResponse)
	return response.Result, nil
}

func (ae ArithmeticEndpoints) HealthCheck() bool {
	return false
}

func (ae ArithmeticEndpoints) Login(name, pwd string) (string, error) {
	return "", errors.New("not implemented")
}
