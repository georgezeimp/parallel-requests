package request

import (
	"fmt"
	"io"
	"net/http"

	"github.com/georgezeimp/parallel-requests/hasher"
	"github.com/georgezeimp/parallel-requests/output"
)

type Service struct {
	hasher          *hasher.Hasher
	outputPresenter *output.Presenter
}

func NewService(hasher *hasher.Hasher, outputPresenter *output.Presenter) *Service {
	return &Service{
		hasher:          hasher,
		outputPresenter: outputPresenter,
	}
}

func (s *Service) Get(address string) []byte {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println(fmt.Errorf("unexpected failure while sending request to %s with error %s", address, err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("unexpected failure while extracting body from the response with error %s", err))
	}

	return body
}
