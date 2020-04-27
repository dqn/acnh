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

	var userID string

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
	})

	// var au *ACNHUser

	t.Run("NewACNHUsers", func(t *testing.T) {
		au, err := a.NewACNHUser(userID)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(au.token)
	})
}
