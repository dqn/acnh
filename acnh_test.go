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

	t.Run("LandsProfile", func(t *testing.T) {
		r, err := a.LandsProfile(token, landID)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.MFruit.Name)
		fmt.Println(r.MNormalNpc[0].Name)
	})
}
