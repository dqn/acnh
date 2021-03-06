package acnh

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (a *ACNH) processRequest(req *http.Request) ([]byte, error) {
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// println(string(b))

	return b, nil
}

func (a *ACNH) get(path string, values *url.Values, token string) ([]byte, error) {
	u := fmt.Sprintf("https://%s/%s", host, path)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	if values != nil {
		req.URL.RawQuery = values.Encode()
	}
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	b, err := a.processRequest(req)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *ACNH) post(path string, body interface{}, token string) ([]byte, error) {
	rawJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	u := fmt.Sprintf("https://%s/%s", host, path)
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(rawJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	b, err := a.processRequest(req)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *ACNH) Users() (*UsersResponse, error) {
	b, err := a.get("api/sd/v1/users", nil, "")
	if err != nil {
		return nil, err
	}

	var r UsersResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}

	return &r, nil
}

func (a *ACNH) AuthToken(userID string) (*AuthTokenResponse, error) {
	b, err := a.post("api/sd/v1/auth_token", &AuthTokenRequest{userID}, "")
	if err != nil {
		return nil, err
	}

	var r AuthTokenResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}
	if r.Token == "" {
		return nil, fmt.Errorf("failed to fetch token")
	}

	return &r, nil
}

func (a *ACNH) UserProfile(token, userID string) (*UserProfileResponse, error) {
	path := fmt.Sprintf("/api/sd/v1/users/%s/profile", userID)
	values := &url.Values{
		"language": {"ja-JP"},
	}
	b, err := a.get(path, values, token)
	if err != nil {
		return nil, err
	}

	var r UserProfileResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}

	return &r, nil
}

func (a *ACNH) LandProfile(token, landID string) (*LandProfileResponse, error) {
	path := fmt.Sprintf("/api/sd/v1/lands/%s/profile", landID)
	values := &url.Values{
		"language": {"ja-JP"},
	}
	b, err := a.get(path, values, token)
	if err != nil {
		return nil, err
	}

	var r LandProfileResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}

	return &r, nil
}

func (a *ACNH) Friends(token string) (*FriendsResponse, error) {
	b, err := a.get("/api/sd/v1/friends", nil, token)
	if err != nil {
		return nil, err
	}

	var r FriendsResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}

	return &r, nil
}

func (a *ACNH) PresenceFriends(token string) (*PresenceFriendsResponse, error) {
	b, err := a.get("/api/sd/v1/friends/presences", nil, token)
	if err != nil {
		return nil, err
	}

	var r PresenceFriendsResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}

	return &r, nil
}

func (a *ACNH) sendMessage(token string, body *SendMessageRequest) (*SendMessageResponse, error) {
	b, err := a.post("/api/sd/v1/messages", body, token)
	if err != nil {
		return nil, err
	}

	var r SendMessageResponse
	if err = json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if r.Code != "" {
		return nil, r.Code
	}

	return &r, nil
}

func (a *ACNH) SendMessageAll(token, message string) (*SendMessageResponse, error) {
	body := &SendMessageRequest{
		Body: message,
		Type: "all_friend",
	}
	r, err := a.sendMessage(token, body)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (a *ACNH) SendMessageFriend(token, userID, message string) (*SendMessageResponse, error) {
	body := &SendMessageRequest{
		Body:   message,
		Type:   "friend",
		UserID: userID,
	}
	r, err := a.sendMessage(token, body)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (a *ACNH) SendMessageLocal(token, message string) (*SendMessageResponse, error) {
	body := &SendMessageRequest{
		Body: message,
		Type: "keyboard",
	}
	r, err := a.sendMessage(token, body)
	if err != nil {
		return nil, err
	}

	return r, nil
}
