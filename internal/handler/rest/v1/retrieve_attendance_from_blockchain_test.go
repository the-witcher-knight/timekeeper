package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/the-witcher-knight/timekeeper/internal/controller/attendance"
	"github.com/the-witcher-knight/timekeeper/internal/model"
	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

func TestHandler_RetrieveAttendanceFromBlockchain(t *testing.T) {
	fromTime, err := time.Parse(time.RFC3339, "2024-06-20T00:00:00+00:00")
	require.NoError(t, err)
	toTime, err := time.Parse(time.RFC3339, "2024-06-20T23:00:00+00:00")
	require.NoError(t, err)

	type mockCtrl struct {
		isCalled  bool
		input     attendance.FilterInput
		outResult []model.Attendance
		outErr    error
	}

	tcs := map[string]struct {
		queryParams   map[string]string
		mockCtrl      mockCtrl
		expStatusCode int
		expRespBody   []attendanceBlockchainResponse
		expErr        error
	}{
		"success": {
			queryParams: map[string]string{
				"employer_id": "1",
				"from_time":   "2024-06-20T00:00:00+00:00",
				"to_time":     "2024-06-20T23:00:00+00:00",
			},
			mockCtrl: mockCtrl{
				isCalled: true,
				input: attendance.FilterInput{
					EmployerID: 1,
					FromTime:   fromTime,
					ToTime:     toTime,
				},
				outResult: []model.Attendance{
					{
						ID:          1,
						EmployerID:  1,
						CheckInTime: time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC),
						Notes:       "Check in",
					},
				},
			},
			expStatusCode: http.StatusOK,
			expRespBody: []attendanceBlockchainResponse{
				{
					ID:          1,
					EmployerID:  1,
					CheckInTime: time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC),
					Notes:       "Check in",
				},
			},
		},
		"error - invalid employer id": {
			queryParams: map[string]string{
				"employer_id": "invalid",
				"from_time":   "2024-06-20T00:00:00+00:00",
				"to_time":     "2024-06-20T23:00:00+00:00",
			},
			expStatusCode: http.StatusBadRequest,
			expErr: httpio.Error{
				Code: "invalid_request",
				Desc: "strconv.ParseInt: parsing \\\"invalid\\\": invalid syntax",
			},
		},
		"error - invalid from time": {
			queryParams: map[string]string{
				"employer_id": "1",
				"from_time":   "invalid",
				"to_time":     "2024-06-20T23:00:00+00:00",
			},
			expStatusCode: http.StatusBadRequest,
			expErr: httpio.Error{
				Code: "invalid_request",
				Desc: "parsing time \\\"invalid\\\" as \\\"2006-01-02T15:04:05Z07:00\\\": cannot parse \\\"invalid\\\" as \\\"2006\\\"",
			},
		},
		"error - invalid to time": {
			queryParams: map[string]string{
				"employer_id": "1",
				"from_time":   "2024-06-20T00:00:00+00:00",
				"to_time":     "invalid",
			},
			expStatusCode: http.StatusBadRequest,
			expErr: httpio.Error{
				Code: "invalid_request",
				Desc: "parsing time \\\"invalid\\\" as \\\"2006-01-02T15:04:05Z07:00\\\": cannot parse \\\"invalid\\\" as \\\"2006\\\"",
			},
		},
		"error - invalid time range": {
			queryParams: map[string]string{
				"employer_id": "1",
				"from_time":   "2024-06-20T23:00:00+00:00",
				"to_time":     "2024-06-20T00:00:00+00:00",
			},
			expStatusCode: http.StatusBadRequest,
			expErr:        errInvalidTimeRange,
		},
		"error - unexpected": {
			queryParams: map[string]string{
				"employer_id": "1",
				"from_time":   "2024-06-20T00:00:00+00:00",
				"to_time":     "2024-06-20T23:00:00+00:00",
			},
			mockCtrl: mockCtrl{
				isCalled: true,
				input: attendance.FilterInput{
					EmployerID: 1,
					FromTime:   fromTime,
					ToTime:     toTime,
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
			r := httptest.NewRequest(http.MethodGet, "/attendances", nil)
			q := r.URL.Query()
			for k, v := range tc.queryParams {
				q.Add(k, v)
			}
			r.URL.RawQuery = q.Encode()

			attCtrlMock := new(attendance.MockController)
			if tc.mockCtrl.isCalled {
				attCtrlMock.On("RetrieveAttendanceFromBlockchain", r.Context(), tc.mockCtrl.input).
					Return(tc.mockCtrl.outResult, tc.mockCtrl.outErr)
			}

			// When
			h := New(nil, attCtrlMock, nil)
			h.RetrieveAttendanceFromBlockchain().ServeHTTP(w, r)

			// Then
			require.Equal(t, tc.expStatusCode, w.Code)
			if tc.expErr != nil {
				require.Equal(t, tc.expErr.Error(), strings.TrimSuffix(w.Body.String(), "\n"))
			} else {
				var respBody []attendanceBlockchainResponse
				require.NoError(t, json.Unmarshal(w.Body.Bytes(), &respBody))
				require.Equal(t, tc.expRespBody, respBody)
			}

			if tc.mockCtrl.isCalled {
				attCtrlMock.AssertExpectations(t)
			}
		})
	}
}
