package httpio

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteJSON(t *testing.T) {
	type data struct {
		Name string `json:"name"`
	}

	tcs := map[string]struct {
		givenInput Response[data]
		expStatus  int
		expResult  string
	}{
		"success": {
			givenInput: Response[data]{
				Status: http.StatusOK,
				Body: data{
					Name: "test",
				},
			},
			expStatus: http.StatusOK,
			expResult: "{\"name\":\"test\"}\n",
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			// When
			WriteJSON(w, r, tc.givenInput)

			// Then
			require.Equal(t, tc.expStatus, w.Code)
			require.Equal(t, tc.expResult, w.Body.String())
		})
	}
}

func TestReadJSON(t *testing.T) {
	type data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tcs := map[string]struct {
		givenInput string
		expResult  data
	}{
		"success": {
			givenInput: "{\"name\":\"test\",\"age\":1}",
			expResult: data{
				Name: "test",
				Age:  1,
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tc.givenInput))

			// When
			req, err := BindJSON[data](w, r)
			require.NoError(t, err)

			// Then
			require.Equal(t, tc.expResult, req)
		})
	}
}
