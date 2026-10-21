package users

import "testing"

func TestPrepareValid(t *testing.T) {
	u := User{Name: "Ana", Email: "ana@mail.com", Password: "123"}

	if err := u.Prepare("register"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPrepareInvalid(t *testing.T) {
	cases := map[string]User{
		"empty name": {Email: "ana@mail.com", Password: "123"},
		"name only spaces": {Name: "   ", Email: "ana@mail.com", Password: "123"},
		"empty email": {Name: "Ana", Password: "123"},
		"empty password": {Name: "Ana", Email: "ana@mail.com"},
	}

	for name, u := range cases {
		t.Run(name, func(t *testing.T) {
			if err := u.Prepare("register"); err == nil {
				t.Errorf("expected error")
			}
		})
	}
}

func TestFormatDerivesUsernameFromEmail(t *testing.T) {
	u := User{Name: "  Ana  ", Email: "ana@mail.com"}
	if err := u.format("register"); err != nil {
		t.Error(err)
	}

	if u.Username != "ana" {
		t.Errorf("Username = %q, want %q", u.Username, "ana")
	}

	if u.Name != "Ana" {
		t.Errorf("Name = %q, want %q", u.Name, "Ana")
	}
}
