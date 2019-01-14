package app

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected error
	}{
		{"Valid Email", "test@email.com", nil},
	}

	for _, tt := range cases {
		actual := isValidEmail(tt.input)
		if actual != tt.expected {
			t.Errorf("%s: expected %d, actual %d", tt.name, tt.expected, actual)
		}
	}
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected error
	}{
		{"Missing Email",
			`title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description`,
			nil},
	}

	for _, tt := range cases {
		_, err := IsValid([]byte(tt.input))
		if err != tt.expected {
			t.Errorf("%s: expected %d, actual %d", tt.name, tt.expected, err)
		}
	}
}
