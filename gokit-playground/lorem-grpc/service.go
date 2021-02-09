package lorem_grpc

import (
	"context"
	"errors"
	"log"
	"strings"

	gl "github.com/drhodes/golorem"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// Define service interface
type Service interface {
	Lorem(ctx context.Context, requestType string, min, max int) (string, error)
}

// Implement service with empty struct
type LoremService struct {
}

// Implement service functions
func (LoremService) Lorem(_ context.Context, requestType string, min, max int) (string, error) {
	log.Println("server: %s %d %d", requestType, min, max)
	var result string
	var err error
	if strings.EqualFold(requestType, "Word") {
		result = gl.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = gl.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = gl.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}
	return result, err
}
