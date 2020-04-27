package acnh

import (
	"net/http"

	"github.com/dqn/go-nso"
)

type ACNH struct {
	client      *http.Client
	accessToken string
}

func New() *ACNH {
	return &ACNH{client: &http.Client{}}
}

func (a *ACNH) Auth() error {
	accessToken, err := nso.New().Auth()
	if err != nil {
		return err
	}
	a.accessToken = accessToken
	return nil
}
