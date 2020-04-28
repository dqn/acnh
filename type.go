package acnh

import "fmt"

type code string

func (c *code) Error() error {
	return fmt.Errorf("nso api error: %s", *c)
}

type Land struct {
	DisplayID int    `json:"displayId"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type Users struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Land  Land   `json:"land"`
	Name  string `json:"name"`
}

type UsersResponse struct {
	Users []Users `json:"users"`
	Code  code    `json:"code"`
}

type AuthTokenRequest struct {
	UserID string `json:"userId"`
}

type AuthTokenResponse struct {
	ExpireAt int    `json:"expireAt"`
	Token    string `json:"token"`
	Code     code   `json:"code"`
}

type MFruit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type MNormalNpc struct {
	BirthDay   int    `json:"birthDay"`
	BirthMonth int    `json:"birthMonth"`
	Image      string `json:"image"`
	Name       string `json:"name"`
}
type MVillager struct {
	MJpeg  string `json:"mJpeg"`
	MPNm   string `json:"mPNm"`
	UserID string `json:"userId"`
}

type LandsProfileResponse struct {
	MFruit     MFruit       `json:"mFruit"`
	MLanguage  string       `json:"mLanguage"`
	MNormalNpc []MNormalNpc `json:"mNormalNpc"`
	MVNm       string       `json:"mVNm"`
	MVRuby     int          `json:"mVRuby"`
	MVer       int          `json:"mVer"`
	MVillager  []MVillager  `json:"mVillager"`
	Code       code         `json:"code"`
}
