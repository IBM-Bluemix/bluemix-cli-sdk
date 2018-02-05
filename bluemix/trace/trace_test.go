package trace_test

import (
	"testing"

	"github.com/IBM-Bluemix/bluemix-cli-sdk/bluemix/trace"
	"github.com/stretchr/testify/assert"
)

type sanitizeTest struct {
	input, expected string
}

func TestSanitize(t *testing.T) {
	tests := []sanitizeTest{
		{
			input:    "Authorization: Bearer eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.rewyrifvoi4243efjlwejf1pZC0yNzAwMDc2UURLIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwNzZdgyidkgfdkgkfdhgkfdhjldfgklfdkiOiJ3d3dlaXdlaUBjbi5pYm0uY29tIiwic3ViIjoid3d3ZWl3ZWlAY24uaWJtLmNvbSIsImlhdCI6MTUxNzUzMzYyMCwiZXhwIjoxNTE3NTM3MjIwLCJpc3MiOiJodHRwczovL2lhbS5ibHVlbWl4Lm5ldC9pZGVudGl0eSIsImdyYW50X3R5cGUiOiJwYXNzd29yZCIsInNjb3BlIjoib3BlbmlkIiwiY2xpZW50X2lkIjoiYngifQ.EoevrKa7BmxQiY7tew98turjldjgns11_jQ3E8Ay1iPcDGqofZ-YwpUPbEwpZn1lERentiv8Dd0939c3qZlXCalLAlryaX99RgyigtrekjgflwrtprewiteroWD7TIouqIaYar-Me72Uug2obW3Nd9G41NHH_5WnvlnyKrbyCl6G__Jyn8CVbo6TaiorKtQa-1_g4kOOA6tVbWiq93LklIi_-0nSUI2-wgWP4IRE4kwOge92NBzTPgeAvQ",
			expected: "Authorization: [PRIVATE DATA HIDDEN]",
		},
		{
			input:    "X-Auth-Token: the-auth-token",
			expected: "X-Auth-Token: [PRIVATE DATA HIDDEN]",
		},
		{
			input:    "grant_type=password&password=my-password",
			expected: "grant_type=password&password=[PRIVATE DATA HIDDEN]",
		},
		{
			input:    "grant_type=refresh_token&refresh_token=eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.rewyrifvoi4243efjlwejf1pZC0yNzAwMDc2UURLIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwNzZdgyidkgfdkgkfdhgkfdhjldfgklfdkiOiJ3d3dlaXdlaUBjbi5pYm0uY29tIiwic3ViIjoid3d3ZWl3ZWlAY24uaWJtLmNvbSIsImlhdCI6MTUxNzUzMzYyMCwiZXhwIjoxNTE3NTM3MjIwLCJpc3MiOiJodHRwczovL2lhbS5ibHVlbWl4Lm5ldC9pZGVudGl0eSIsImdyYW50X3R5cGUiOiJwYXNzd29yZCIsInNjb3BlIjoib3BlbmlkIiwiY2xpZW50X2lkIjoiYngifQ.EoevrKa7BmxQiY7tew98turjldjgns11_jQ3E8Ay1iPcDGqofZ-YwpUPbEwpZn1lERentiv8Dd0939c3qZlXCalLAlryaX99RgyigtrekjgflwrtprewiteroWD7TIouqIaYar-Me72Uug2obW3Nd9G41NHH_5WnvlnyKrbyCl6G__Jyn8CVbo6TaiorKtQa-1_g4kOOA6tVbWiq93LklIi_-0nSUI2-wgWP4IRE4kwOge92NBzTPgeAvQ",
			expected: "grant_type=refresh_token&refresh_token=[PRIVATE DATA HIDDEN]",
		},
		{
			input:    "apikey=my-api-key&grant_type=urn:ibm:params:oauth:grant-type:apikey",
			expected: "apikey=[PRIVATE DATA HIDDEN]&grant_type=urn:ibm:params:oauth:grant-type:apikey",
		},
		{
			input:    "passcode=the-one-time-code",
			expected: "passcode=[PRIVATE DATA HIDDEN]",
		},
		{
			input:    "PASSCODE=the-one-time-code",
			expected: "PASSCODE=[PRIVATE DATA HIDDEN]",
		},
		{
			input:    `{"access_token":"the-access-token","refresh_token":"the-refresh-token"}`,
			expected: `{"access_token":"[PRIVATE DATA HIDDEN]","refresh_token":"[PRIVATE DATA HIDDEN]"}`,
		},
		{
			input:    `{"token_endpoint":"https://the-token-endpoint"}`,
			expected: `{"token_endpoint":"https://the-token-endpoint"}`,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, trace.Sanitize(test.input))
	}
}
