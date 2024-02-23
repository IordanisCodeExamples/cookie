package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	servicemock "cookie/internal/mocks/service"
	"cookie/internal/service"
)

type serviceMocks struct {
	input *servicemock.MockInput
}

func getMocks(t *testing.T) serviceMocks {
	t.Helper()
	ctrl := gomock.NewController(t)

	return serviceMocks{
		input: servicemock.NewMockInput(ctrl),
	}
}

func getService(t *testing.T, mocks serviceMocks) *service.Service {
	t.Helper()
	srv, err := service.New(mocks.input)
	assert.NoError(t, err)

	return srv
}

// Test_New tests the New function of the service package.
func Test_New(t *testing.T) {
	mocks := getMocks(t)

	tests := []struct {
		name          string
		input         service.Input
		expectError   bool
		errorContains string
	}{
		{
			name:          "Creates successfully a Service object",
			input:         mocks.input,
			expectError:   false,
			errorContains: "",
		},
		{
			name:          "Returns error when input is nil",
			input:         nil,
			expectError:   true,
			errorContains: "input_is_nil",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv, err := service.New(tc.input)

			if tc.expectError {
				assert.NotNil(t, err)
				assert.Nil(t, srv)
				if tc.errorContains != "" {
					assert.Contains(t, err.Error(), tc.errorContains)
				}
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, srv)
			}
		})
	}
}
