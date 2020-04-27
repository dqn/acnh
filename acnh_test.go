package acnh

import (
	"fmt"
	"testing"
)

func TestACNH(t *testing.T) {
	a := New()
	if err := a.Auth(); err != nil {
		t.Fatal(err)
	}

	t.Run("users", func(t *testing.T) {
		if _, err := a.Users(); err != nil {
			t.Fatal(err)
		}
		r, err := a.Users()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(r.Users[0].Name)
	})
}
