package acnh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const host = "web.sd.lp1.acbaa.srv.nintendo.net"

type ACNH struct {
	client      *http.Client
	accessToken string
}

type ACNHUser struct {
	client *http.Client
	token  string
}

func New(accessToken string) (*ACNH, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	cookies := []*http.Cookie{
		{Name: "_gtoken", Value: accessToken},
	}

	u, err := url.Parse(fmt.Sprintf("https://%s", host))
	if err != nil {
		return nil, err
	}
	jar.SetCookies(u, cookies)
	client := &http.Client{Jar: jar}

	return &ACNH{client, accessToken}, nil
}

func (a *ACNH) request(method, path string, body interface{}) ([]byte, error) {
	var reader io.Reader
	if body != nil {
		rawJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(rawJSON)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("https://%s/%s", host, path), reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// println(string(b))

	return b, nil
}

func (a *ACNH) Users() (*UsersResponse, error) {
	b, err := a.request("GET", "api/sd/v1/users", nil)
	if err != nil {
		return nil, err
	}

	var r UsersResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *ACNH) NewACNHUser(userID string) (*ACNHUser, error) {
	b, err := a.request("POST", "api/sd/v1/auth_token", &AuthTokenRequest{userID})
	if err != nil {
		return nil, err
	}

	var r AuthTokenResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Token == "" {
		return nil, fmt.Errorf("failed to fetch token")
	}

	return &ACNHUser{&http.Client{}, r.Token}, nil
}
