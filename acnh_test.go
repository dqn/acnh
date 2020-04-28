package acnh

import (
	"fmt"
	"testing"

	"github.com/dqn/go-nso"
)

func TestACNH(t *testing.T) {
	accessToken, err := nso.New().Auth()
	if err != nil {
		t.Fatal(err)
	}
	a, err := New(accessToken)
	if err != nil {
		t.Fatal(err)
	}

	var (
		userID string
		landID string
	)

	t.Run("Users", func(t *testing.T) {
		if _, err := a.Users(); err != nil {
			t.Fatal(err)
		}
		r, err := a.Users()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.Users[0].Name)
		userID = r.Users[0].ID
		landID = r.Users[0].Land.ID
	})

	var token string

	t.Run("AuthToken", func(t *testing.T) {
		r, err := a.AuthToken(userID)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.Token)
		token = r.Token
	})

	t.Run("UserProfile", func(t *testing.T) {
		r, err := a.UserProfile(token, userID)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.MHandleName)
		fmt.Println(r.MComment)
	})

	t.Run("LandProfile", func(t *testing.T) {
		r, err := a.LandProfile(token, landID)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.MFruit.Name)
		fmt.Println(r.MNormalNpc[0].Name)
	})

	t.Run("Friends", func(t *testing.T) {
		r, err := a.Friends(token)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.Friends[0].Name)
		fmt.Println(r.Friends[0].Land.Name)
	})

	t.Run("PresenceFriends", func(t *testing.T) {
		r, err := a.PresenceFriends(token)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(len(r.Presences))
	})
}
