package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/auth"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

func TestHandler_RecordAttendanceToBlockchain(t *testing.T) {
	type mockCtrl struct {
		isCalled bool
		input    int64
		outErr   error
	}

	tcs := map[string]struct {
		givenInput    auth.UserProfile
		mockCtrl      mockCtrl
		expStatusCode int
		expRespBody   httpio.Message
		expErr        error
	}{
		"success": {
			givenInput: auth.UserProfile{
				ID: 1,
			},
			mockCtrl: mockCtrl{
				isCalled: true,
				input:    1,
			},
			expStatusCode: http.StatusOK,
			expRespBody: httpio.Message{
				Code: "created",
				Desc: "Attendance created successfully",
			},
		},
		"error - user profile not found": {
			expStatusCode: http.StatusBadRequest,
			expErr:        errUserProfileNotFound,
		},
		"error - unexpected": {
			givenInput: auth.UserProfile{
				ID: 1,
			},
			mockCtrl: mockCtrl{
				isCalled: true,
				input:    1,
				outErr:   errors.New("unexpected error"),
			},
			expStatusCode: http.StatusInternalServerError,
			expErr: httpio.Error{
				Code: "internal_server_error",
				Desc: "internal server error",
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/attendances", nil)

			r = r.WithContext(auth.SetInCtx(r.Context(), tc.givenInput))

			attCtrlMock := new(attendance.MockController)
			if tc.mockCtrl.isCalled {
				attCtrlMock.On("RecordAttendanceToBlockchain", r.Context(), tc.mockCtrl.input).
					Return(tc.mockCtrl.outErr)
			}

			// When
			h := New(nil, attCtrlMock)
			h.RecordAttendanceToBlockchain().ServeHTTP(w, r)

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
				attCtrlMock.AssertExpectations(t)
			}
		})
	}
}
