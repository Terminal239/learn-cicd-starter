package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	type input struct {
		input http.Header
		want  string
	}

	tests := map[string]input{
		"correct":              {input: http.Header{"Authorization": {"ApiKey token"}}, want: "token"},
		"incorrect seperation": {input: http.Header{"Authorization": {"ApiKey-token"}}, want: ""},
		"incorrect marker":     {input: http.Header{"Authorization": {"Bearer token"}}, want: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := auth.GetAPIKey(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
