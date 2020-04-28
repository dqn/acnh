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
type MBirth struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}
type MTimeStamp struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type UserProfileResponse struct {
	ContentID         string     `json:"contentId"`
	CreatedAt         int        `json:"createdAt"`
	Digest            string     `json:"digest"`
	LandName          string     `json:"landName"`
	MBirth            MBirth     `json:"mBirth"`
	MComment          string     `json:"mComment"`
	MHandleName       string     `json:"mHandleName"`
	MIsLandMaster     bool       `json:"mIsLandMaster"`
	MJpeg             string     `json:"mJpeg"`
	MLanguage         string     `json:"mLanguage"`
	MMyDesignAuthorID string     `json:"mMyDesignAuthorId"`
	MPNm              string     `json:"mPNm"`
	MTimeStamp        MTimeStamp `json:"mTimeStamp"`
	MVer              int        `json:"mVer"`
	Code              code       `json:"code"`
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

type LandProfileResponse struct {
	MFruit     MFruit       `json:"mFruit"`
	MLanguage  string       `json:"mLanguage"`
	MNormalNpc []MNormalNpc `json:"mNormalNpc"`
	MVNm       string       `json:"mVNm"`
	MVRuby     int          `json:"mVRuby"`
	MVer       int          `json:"mVer"`
	MVillager  []MVillager  `json:"mVillager"`
	Code       code         `json:"code"`
}

type Friends struct {
	Image  string `json:"image"`
	Land   Land   `json:"land"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type FriendsResponse struct {
	Friends []Friends `json:"friends"`
	Code    code      `json:"code"`
}

type PresenceFriendsResponse struct {
	Presences []Friends `json:"presences"`
	Code      code      `json:"code"`
}

type SendMessageRequest struct {
	Body   string `json:"body"`
	Type   string `json:"type"`
	UserID string `json:"userId"`
}

type SendMessageResponse struct {
	Status string `json:"status"`
	Code   code   `json:"code"`
}
