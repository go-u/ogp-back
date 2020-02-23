package users

import "testing"

func TestIsValidUsername(t *testing.T) {
	cases := []struct {
		name  string
		valid bool
	}{
		{name: "", valid: false},
		{name: " ", valid: false},
		{name: "1", valid: false},
		{name: "#", valid: false},
		{name: "#a", valid: false},
		{name: "氏名", valid: false}, // 半角アスキー文字以外はエラー
		{name: "__", valid: true},
		{name: "go", valid: true},
		{name: "1234567890", valid: true},
		{name: "12345678901", valid: false}, // 境界値
	}

	for _, c := range cases {
		if isValidUsername(c.name) != c.valid {
			t.Errorf("username validation fail: %s\n", c.name)
		}
	}
}

func TestIsValidDisplayName(t *testing.T) {
	cases := []struct {
		name  string
		valid bool
	}{
		{name: "", valid: false},
		{name: " ", valid: false},
		{name: "first last", valid: false}, // 半角アスキー文字以外はエラー
		{name: "first_last", valid: true},
		{name: "123456789012345", valid: true},
		{name: "1234567890123456", valid: false}, // 境界値
	}

	for _, c := range cases {
		if isValidDisplayName(c.name) != c.valid {
			t.Errorf("displayname validation fail: %s\n", c.name)
		}
	}
}
