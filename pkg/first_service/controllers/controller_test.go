package controllers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-microservices-template/internal/response"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockService struct {
	mock.Mock
}

func NewMockService() MockService {
	return MockService{}
}

func (s *MockService) Ping() (string, error) {
	args := s.Called()
	return args.Get(0).(string), args.Error(1)
}

func TestControllerImpl_Ping(t *testing.T) {
	type vals struct {
		mockVal          string
		mockErr          error
		expectedResponse response.ResponseData
	}
	tests := []struct {
		name string
		vals vals
	}{
		{
			name: "200",
			vals: vals{
				mockVal: "pong",
				mockErr: nil,
				expectedResponse: response.ResponseData{
					Meta: response.Meta{
						Code:    200,
						Message: "ok",
					},
					Data: "pong",
				},
			},
		},
		{
			name: "500",
			vals: vals{
				mockVal: "",
				mockErr: errors.New("nothing"),
				expectedResponse: response.ResponseData{
					Meta: response.Meta{
						Code:    500,
						Message: "nothing",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// set mock return values for function "Ping" of service "Service"
			mockService := MockService{}
			w := httptest.NewRecorder()
			bodyReader := strings.NewReader("hello")
			r := httptest.NewRequest("GET", "/ping?a=x", bodyReader)
			mockService.On("Ping").Return(tt.vals.mockVal, tt.vals.mockErr)

			c := &ControllerImpl{
				service: &mockService,
			}

			c.Ping(w, r)

			// get result from http response
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			var resBody response.ResponseData
			_ = json.Unmarshal(body, &resBody)

			assert.Equal(t, tt.vals.expectedResponse, resBody)
		})
	}
}
