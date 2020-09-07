package lorem

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// request
type LoremRequest struct {
	RequestType string
	Min         int
	Max         int
}

// response
type LoremResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

// endpoints wrapper
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

func MakeLoremEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		var (
			txt      string
			min, max int
		)

		min = req.Min
		max = req.Max

		if strings.EqualFold(req.RequestType, "word") {
			txt = svc.Word(min, max)
		} else if strings.EqualFold(req.RequestType, "Sentence") {
			txt = svc.Sentence(min, max)
		} else if strings.EqualFold(req.RequestType, "paragraph") {
			txt = svc.Paragraph(min, max)
		} else {
			return nil, ErrRequestTypeNotFound
		}
		log.Printf("txt: %s", txt)
		return LoremResponse{Message: txt}, nil
	}
}
