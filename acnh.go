package acnh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (a *ACNH) Users() (*UsersResponse, error) {
	url := "https://web.sd.lp1.acbaa.srv.nintendo.net/api/sd/v1/users"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", fmt.Sprintf("_gtoken=%s", a.accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r UsersResponse
	json.Unmarshal(b, &r)

	return &r, nil
}
