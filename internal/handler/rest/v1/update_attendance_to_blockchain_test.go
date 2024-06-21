package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

func TestHandler_UpdateAttendanceToBlockchain(t *testing.T) {
	checkInTime, err := time.Parse(time.RFC3339, "2024-06-20T00:00:00+00:00")
	require.NoError(t, err)

	type mockCtrl struct {
		isCalled bool
		input    model.Attendance
		outErr   error
	}

	tcs := map[string]struct {
		givenID       int64
		givenReqBody  string
		mockCtrl      mockCtrl
		expStatusCode int
		expRespBody   httpio.Message
		expErr        error
	}{
		"success": {
			givenID:      1,
			givenReqBody: `{"id":1,"employer_id":1,"check_in_time":"2024-06-20T00:00:00+00:00","notes":"Check In"}`,
			mockCtrl: mockCtrl{
				isCalled: true,
				input: model.Attendance{
					ID:          1,
					EmployerID:  1,
					CheckInTime: checkInTime,
					Notes:       "Check In",
				},
			},
			expStatusCode: http.StatusOK,
			expRespBody: httpio.Message{
				Code: "updated",
				Desc: "Attendance updated successfully",
			},
		},
		"error - invalid id": {
			givenID:       1,
			givenReqBody:  `{"employer_id":1,"check_in_time":"2024-06-20T00:00:00+00:00","notes":"Check In"}`,
			expStatusCode: http.StatusBadRequest,
			expErr:        errInvalidAttendanceID,
		},
		"error - invalid employer id": {
			givenID:       1,
			givenReqBody:  `{"id":1,"check_in_time":"2024-06-20T00:00:00+00:00","notes":"Check In"}`,
			expStatusCode: http.StatusBadRequest,
			expErr:        errInvalidEmployerID,
		},
		"error - id from url and request body not match": {
			givenID:      1111,
			givenReqBody: `{"id":1,"employer_id":1,"check_in_time":"2024-06-20T00:00:00+00:00","notes":"Check In"}`,

			expStatusCode: http.StatusBadRequest,
			expErr:        errInvalidAttendanceID,
		},
		"error - unexpected": {
			givenID:      1,
			givenReqBody: `{"id":1,"employer_id":1,"check_in_time":"2024-06-20T00:00:00+00:00","notes":"Check In"}`,
			mockCtrl: mockCtrl{
				isCalled: true,
				input: model.Attendance{
					ID:          1,
					EmployerID:  1,
					CheckInTime: checkInTime,
					Notes:       "Check In",
				},
				outErr: errors.New("unexpected error"),
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
			r := httptest.NewRequest(http.MethodPost, "/attendances/{id}", strings.NewReader(tc.givenReqBody))

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", fmt.Sprint(tc.givenID))

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			attCtrlMock := new(attendance.MockController)
			if tc.mockCtrl.isCalled {
				attCtrlMock.On("UpdateAttendanceToBlockchain", r.Context(), tc.mockCtrl.input).
					Return(tc.mockCtrl.outErr)
			}

			// When
			h := New(nil, attCtrlMock, nil)
			h.UpdateAttendanceToBlockchain().ServeHTTP(w, r)

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
