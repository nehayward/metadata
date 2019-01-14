package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpload(t *testing.T) {
	cases := []struct {
		name     string
		input    []byte
		expected int
	}{
		{"Valid",
			[]byte(`title: Valid App 1
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
  Some application content, and description`),
			http.StatusAccepted,
		},
		{"Invalid",
			[]byte(`title: App w/ Invalid maintainer email
version: 1.0.1
maintainers:
- name: Firstname Lastname
  email: apptwohotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |
 ### blob of markdown
 More markdown`),
			http.StatusBadRequest,
		},
	}

	for _, tt := range cases {
		req, _ := http.NewRequest("POST", "/upload", bytes.NewBuffer(tt.input))
		req.Header.Set("Content-Type", "text/x-yaml")

		rr := httptest.NewRecorder()

		router := NewRouter(AllRoutes())

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != tt.expected {
			t.Errorf("Wrong status")
		}
	}
}
