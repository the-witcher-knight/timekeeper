package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/controller/bcauth"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

func TestHandler_AuthorizeAccount(t *testing.T) {
	const testAddr = "0xE59b798c3eb36825fEc1f7aB26236Ce73C750d11"

	type mockCtrl struct {
		isCalled bool
		input    common.Address
		outErr   error
	}

	tcs := map[string]struct {
		reqBody       string
		mockCtrl      mockCtrl
		expStatusCode int
		expRespBody   httpio.Message
		expErr        error
	}{
		"success": {
			reqBody: fmt.Sprintf(`{"account":"%s"}`, testAddr),
			mockCtrl: mockCtrl{
				isCalled: true,
				input:    common.HexToAddress(testAddr),
			},
			expStatusCode: http.StatusOK,
			expRespBody: httpio.Message{
				Code: "authorized",
				Desc: "Account authorized",
			},
		},
		"error - account blank": {
			reqBody:       `{"account":""}`,
			expStatusCode: http.StatusBadRequest,
			expErr:        errAccountIsRequired,
		},
		"error - account is invalid": {
			reqBody:       `{"account":"INVALID"}`,
			expStatusCode: http.StatusBadRequest,
			expErr:        errAccountAddressIsInvalid,
		},
		"error - unexpected error": {
			reqBody: fmt.Sprintf(`{"account":"%s"}`, testAddr),
			mockCtrl: mockCtrl{
				isCalled: true,
				input:    common.HexToAddress(testAddr),
				outErr:   errors.New("unexpected error"),
			},
			expStatusCode: http.StatusInternalServerError,
			expRespBody: httpio.Message{
				Code: "internal_server_error",
				Desc: "internal server error",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/authorize", strings.NewReader(tc.reqBody))

			bcMock := new(bcauth.MockController)
			if tc.mockCtrl.isCalled {
				bcMock.On("AuthorizeAccount", r.Context(), tc.mockCtrl.input).
					Return(tc.mockCtrl.outErr)
			}

			// When
			h := New(nil, nil, bcMock)
			h.AuthorizeAccount().ServeHTTP(w, r)

			// Then
			require.Equal(t, tc.expStatusCode, w.Code)
			if tc.expErr != nil {
				require.Equal(t, tc.expErr.Error(), strings.TrimSuffix(w.Body.String(), "\n"))
			} else {
				var respBody httpio.Message
				require.NoError(t, json.Unmarshal(w.Body.Bytes(), &respBody))
				require.Equal(t, tc.expRespBody, respBody)
			}

			if tc.mockCtrl.isCalled {
				bcMock.AssertExpectations(t)
			}
		})
	}
}
