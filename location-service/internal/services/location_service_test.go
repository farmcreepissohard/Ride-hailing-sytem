package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/dto"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/models/enum"
	"github.com/farmcreepissohard/Ride-hailing-sytem/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLocationRepository struct {
	mock.Mock
}

// ChangeStatus implements [repositories.LocationRepository].
func (m *MockLocationRepository) ChangeStatus(id string, status enum.DriverStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

// MatchingNearbyDrivers implements [repositories.LocationRepository].
func (m *MockLocationRepository) MatchingNearbyDrivers(lng float64, lat float64, radius float64) ([]dto.MatchingResponse, error) {
	args := m.Called(lng, lat, radius)
	var drivers []dto.MatchingResponse
	if args.Get(0) != nil {
		drivers = args.Get(0).([]dto.MatchingResponse)
	}
	return drivers, args.Error(1)
}

// OnBusy implements [repositories.LocationRepository].
func (m *MockLocationRepository) OnBusy(driverId string) error {
	args := m.Called(driverId)
	return args.Error(0)
}

// OnCancel implements [repositories.LocationRepository].
func (m *MockLocationRepository) OnCancel(driverId string) error {
	args := m.Called(driverId)
	return args.Error(0)
}

// OnReady implements [repositories.LocationRepository].
func (m *MockLocationRepository) OnReady(driverId string) error {
	args := m.Called(driverId)
	return args.Error(0)
}

// OutBusy implements [repositories.LocationRepository].
func (m *MockLocationRepository) OutBusy(driverId string) error {
	args := m.Called(driverId)
	return args.Error(0)
}

// UpdateLocation implements [repositories.LocationRepository].
func (m *MockLocationRepository) UpdateLocation(id string, lng float64, lat float64) error {
	args := m.Called(id, lng, lat)
	return args.Error(0)
}

type MockTripGrpcClient struct {
	mock.Mock
}

// Timeout implements [grpc_client.TripGrpcClient].
func (m *MockTripGrpcClient) Timeout(ctx context.Context, tripId string) (bool, error) {
	args := m.Called(ctx, tripId)
	return args.Bool(0), args.Error(1)
}

// CancelTrip implements [grpc_client.TripGrpcClient].
func (m *MockTripGrpcClient) CancelTrip(ctx context.Context, tripId string, cancelledBy string, reason string) (bool, error) {
	args := m.Called(ctx, tripId, cancelledBy, reason)
	return args.Bool(0), args.Error(1)
}

// CompleteTrip implements [grpc_client.TripGrpcClient].
func (m *MockTripGrpcClient) CompleteTrip(ctx context.Context, tripId string, driverId string) (bool, error) {
	args := m.Called(ctx, tripId, driverId)
	return args.Bool(0), args.Error(1)
}

// MatchTrip implements [grpc_client.TripGrpcClient].
func (m *MockTripGrpcClient) MatchTrip(ctx context.Context, tripId string, driverId string) (bool, error) {
	args := m.Called(ctx, tripId, driverId)
	return args.Bool(0), args.Error(1)
}

// StartTrip implements [grpc_client.TripGrpcClient].
func (m *MockTripGrpcClient) StartTrip(ctx context.Context, tripId string, driverId string) (bool, error) {
	args := m.Called(ctx, tripId, driverId)
	return args.Bool(0), args.Error(1)
}

func TestChangeStatus(t *testing.T) {
	testCases := []struct {
		name             string
		driverId         string
		status           enum.DriverStatus
		repoReturnErr    error
		expectedError    string
		expectedRepoCall bool
	}{
		{
			name:             "success when valid input",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			status:           enum.StatusOnline,
			repoReturnErr:    nil,
			expectedError:    "",
			expectedRepoCall: true,
		},
		{
			name:             "failed when repo returns error",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			status:           enum.StatusOnline,
			repoReturnErr:    errors.New("redis connection error"),
			expectedError:    "redis connection error",
			expectedRepoCall: true,
		},
		{
			name:             "failed when null id",
			status:           enum.StatusOnline,
			repoReturnErr:    nil,
			expectedError:    "Driver id is required",
			expectedRepoCall: false,
		},
		{
			name:             "failed when null status",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			repoReturnErr:    nil,
			expectedError:    "Status is required",
			expectedRepoCall: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpcClient := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpcClient, mockRepo)

			if tc.expectedRepoCall {
				mockRepo.On("ChangeStatus", tc.driverId, tc.status).Return(tc.repoReturnErr)
			}
			err := service.ChangeStatus(tc.driverId, tc.status)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateLocation(t *testing.T) {
	testCases := []struct {
		name             string
		driverId         string
		lng              float64
		lat              float64
		expectedRepoCall bool
		repoReturnErr    error
		expectedError    string
	}{
		{
			name:             "successful when valid input",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			lng:              23.69932,
			lat:              13.65963,
			expectedRepoCall: true,
			repoReturnErr:    nil,
			expectedError:    "",
		},
		{
			name:             "failed when repo returns error",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			lng:              23.69932,
			lat:              13.65963,
			expectedRepoCall: true,
			repoReturnErr:    errors.New("redis connection error"),
			expectedError:    "redis connection error",
		},
		{
			name:             "failed when null driver id",
			lng:              23.69932,
			lat:              13.65963,
			expectedRepoCall: false,
			repoReturnErr:    nil,
			expectedError:    "Driver id is required",
		},
		{
			name:             "failed when invalid latitude < -90",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			lng:              23.69932,
			lat:              -93.65963,
			expectedRepoCall: false,
			repoReturnErr:    nil,
			expectedError:    "Invalid latitude value",
		},
		{
			name:             "failed when invalid latitude > 90",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			lng:              23.69932,
			lat:              90.65963,
			expectedRepoCall: false,
			repoReturnErr:    nil,
			expectedError:    "Invalid latitude value",
		},
		{
			name:             "failed when invalid longitude < -180",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			lng:              -182.69932,
			lat:              13.65963,
			expectedRepoCall: false,
			repoReturnErr:    nil,
			expectedError:    "Invalid longitude value",
		},
		{
			name:             "failed when invalid longitude > 180",
			driverId:         "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			lng:              182.69932,
			lat:              13.65963,
			expectedRepoCall: false,
			repoReturnErr:    nil,
			expectedError:    "Invalid longitude value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpc := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpc, mockRepo)

			if tc.expectedRepoCall {
				mockRepo.On("UpdateLocation", tc.driverId, tc.lng, tc.lat).Return(tc.repoReturnErr)
			}
			err := service.UpdateLocation(tc.driverId, tc.lng, tc.lat)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestMatchingNearbyDrivers(t *testing.T) {
	testCases := []struct {
		name               string
		lng                float64
		lat                float64
		radius             float64
		expectedRepoCall   bool
		repoReturnErr      error
		expectedError      string
		repoReturnResponse []dto.MatchingResponse
		expectedResponse   []dto.MatchingResponse
	}{
		{
			name:               "successful when valid input",
			lng:                23.69932,
			lat:                13.65963,
			radius:             3,
			expectedRepoCall:   true,
			repoReturnErr:      nil,
			expectedError:      "",
			repoReturnResponse: []dto.MatchingResponse{{DriverId: "fbd3ed3e-c385-46cd-ad93-975bbe174488"}},
			expectedResponse:   []dto.MatchingResponse{{DriverId: "fbd3ed3e-c385-46cd-ad93-975bbe174488"}},
		},
		{
			name:               "failed when repo returns error",
			lng:                23.69932,
			lat:                13.65963,
			radius:             3,
			expectedRepoCall:   true,
			repoReturnErr:      errors.New("redis connection error"),
			expectedError:      "redis connection error",
			repoReturnResponse: nil,
			expectedResponse:   nil,
		},
		{
			name:               "failed when invalid latitude < -90",
			lng:                23.69932,
			lat:                -93.65963,
			radius:             3,
			expectedRepoCall:   false,
			repoReturnErr:      nil,
			expectedError:      "Invalid latitude value",
			repoReturnResponse: nil,
			expectedResponse:   nil,
		},
		{
			name:               "failed when invalid latitude > 90",
			lng:                23.69932,
			lat:                90.65963,
			radius:             3,
			expectedRepoCall:   false,
			repoReturnErr:      nil,
			expectedError:      "Invalid latitude value",
			repoReturnResponse: nil,
			expectedResponse:   nil,
		},
		{
			name:               "failed when invalid longitude < -180",
			lng:                -182.69932,
			lat:                13.65963,
			radius:             3,
			expectedRepoCall:   false,
			repoReturnErr:      nil,
			expectedError:      "Invalid longitude value",
			repoReturnResponse: nil,
			expectedResponse:   nil,
		},
		{
			name:               "failed when invalid longitude > 180",
			lng:                182.69932,
			lat:                13.65963,
			radius:             3,
			expectedRepoCall:   false,
			repoReturnErr:      nil,
			expectedError:      "Invalid longitude value",
			repoReturnResponse: nil,
			expectedResponse:   nil,
		},
		{
			name:               "failed when invalid raidus",
			lng:                23.69932,
			lat:                13.65963,
			radius:             -0.00001,
			expectedRepoCall:   false,
			repoReturnErr:      nil,
			expectedError:      "Invalid radius value",
			repoReturnResponse: nil,
			expectedResponse:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpcClient := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpcClient, mockRepo)

			if tc.expectedRepoCall {
				mockRepo.On("MatchingNearbyDrivers", tc.lng, tc.lat, tc.radius).Return(tc.repoReturnResponse, tc.repoReturnErr)
			}
			res, err := service.MatchingNearbyDrivers(tc.lng, tc.lat, tc.radius)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
				assert.Equal(t, tc.expectedResponse, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, res)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestAcceptTrip(t *testing.T) {
	testCases := []struct {
		name                string
		ctx                 context.Context
		tripId              string
		driverId            string
		expectedRepoCalled  bool
		expectedGrpcCalled  bool
		repoReturnError     error
		grpcReturnError     error
		grpcReturnIsSuccess bool
		expectedError       string
	}{
		{
			name:                "successful when happy case",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: true,
			expectedError:       "",
		},
		{
			name:               "failed when context is null",
			ctx:                nil,
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Context is required",
		},
		{
			name:               "failed when trip id is empty",
			ctx:                context.Background(),
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Trip id is required",
		},
		{
			name:               "failed when driver id is empty",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Driver id is required",
		},
		{
			name:               "failed when repo return error",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: true,
			expectedGrpcCalled: false,
			repoReturnError:    errors.New("redis connection error"),
			expectedError:      "redis connection error",
		},
		{
			name:                "failed when grpc return error",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     errors.New("grpc called error"),
			grpcReturnIsSuccess: true,
			expectedError:       "grpc called error",
		},
		{
			name:                "failed when grpc called unsuccessfully",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: false,
			expectedError:       "Trip already taken or forbidden",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpcClient := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpcClient, mockRepo)

			if tc.expectedRepoCalled {
				mockRepo.On("OnReady", tc.driverId).Return(tc.repoReturnError)
			}
			if tc.expectedGrpcCalled {
				mockGrpcClient.On("MatchTrip", tc.ctx, tc.tripId, tc.driverId).Return(tc.grpcReturnIsSuccess, tc.grpcReturnError)
			}

			err := service.AcceptTrip(tc.ctx, tc.tripId, tc.driverId)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockGrpcClient.AssertExpectations(t)
		})
	}
}

func TestStartTrip(t *testing.T) {
	testCases := []struct {
		name                string
		ctx                 context.Context
		tripId              string
		driverId            string
		expectedRepoCalled  bool
		expectedGrpcCalled  bool
		repoReturnError     error
		grpcReturnError     error
		grpcReturnIsSuccess bool
		expectedError       string
	}{
		{
			name:                "successful when happy case",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: true,
			expectedError:       "",
		},
		{
			name:               "failed when context is null",
			ctx:                nil,
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Context is required",
		},
		{
			name:               "failed when trip id is empty",
			ctx:                context.Background(),
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Trip id is required",
		},
		{
			name:               "failed when driver id is empty",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Driver id is required",
		},
		{
			name:               "failed when repo return error",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: true,
			expectedGrpcCalled: false,
			repoReturnError:    errors.New("redis connection error"),
			expectedError:      "redis connection error",
		},
		{
			name:                "failed when grpc return error",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     errors.New("grpc called error"),
			grpcReturnIsSuccess: true,
			expectedError:       "grpc called error",
		},
		{
			name:                "failed when grpc called unsuccessfully",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: false,
			expectedError:       "Can not start trip, please try again later",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpcClient := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpcClient, mockRepo)

			if tc.expectedRepoCalled {
				mockRepo.On("OnBusy", tc.driverId).Return(tc.repoReturnError)
			}
			if tc.expectedGrpcCalled {
				mockGrpcClient.On("StartTrip", tc.ctx, tc.tripId, tc.driverId).Return(tc.grpcReturnIsSuccess, tc.grpcReturnError)
			}

			err := service.StartTrip(tc.ctx, tc.tripId, tc.driverId)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockGrpcClient.AssertExpectations(t)
		})
	}
}

func TestCompleteTrip(t *testing.T) {
	testCases := []struct {
		name                string
		ctx                 context.Context
		tripId              string
		driverId            string
		expectedRepoCalled  bool
		expectedGrpcCalled  bool
		repoReturnError     error
		grpcReturnError     error
		grpcReturnIsSuccess bool
		expectedError       string
	}{
		{
			name:                "successful when happy case",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: true,
			expectedError:       "",
		},
		{
			name:               "failed when context is null",
			ctx:                nil,
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Context is required",
		},
		{
			name:               "failed when trip id is empty",
			ctx:                context.Background(),
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Trip id is required",
		},
		{
			name:               "failed when driver id is empty",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Driver id is required",
		},
		{
			name:               "failed when repo return error",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: true,
			expectedGrpcCalled: false,
			repoReturnError:    errors.New("redis connection error"),
			expectedError:      "redis connection error",
		},
		{
			name:                "failed when grpc return error",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     errors.New("grpc called error"),
			grpcReturnIsSuccess: true,
			expectedError:       "grpc called error",
		},
		{
			name:                "failed when grpc called unsuccessfully",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: false,
			expectedError:       "Can not complete trip, please try again later",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpcClient := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpcClient, mockRepo)

			if tc.expectedRepoCalled {
				mockRepo.On("OutBusy", tc.driverId).Return(tc.repoReturnError)
			}
			if tc.expectedGrpcCalled {
				mockGrpcClient.On("CompleteTrip", tc.ctx, tc.tripId, tc.driverId).Return(tc.grpcReturnIsSuccess, tc.grpcReturnError)
			}

			err := service.CompleteTrip(tc.ctx, tc.tripId, tc.driverId)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockGrpcClient.AssertExpectations(t)
		})
	}
}

func TestCancelTrip(t *testing.T) {
	testCases := []struct {
		name                string
		ctx                 context.Context
		tripId              string
		driverId            string
		reason              string
		expectedRepoCalled  bool
		expectedGrpcCalled  bool
		repoReturnError     error
		grpcReturnError     error
		grpcReturnIsSuccess bool
		expectedError       string
	}{
		{
			name:                "successful when happy case",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			reason:              "Traffic problem",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: true,
			expectedError:       "",
		},
		{
			name:               "failed when context is null",
			ctx:                nil,
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			reason:             "Traffic problem",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Context is required",
		},
		{
			name:               "failed when trip id is empty",
			ctx:                context.Background(),
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			reason:             "Traffic problem",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Trip id is required",
		},
		{
			name:               "failed when driver id is empty",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			reason:             "Traffic problem",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Driver id is required",
		},
		{
			name:               "failed when reason is empty",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			expectedRepoCalled: false,
			expectedGrpcCalled: false,
			expectedError:      "Unknown reason",
		},
		{
			name:               "failed when repo return error",
			ctx:                context.Background(),
			tripId:             "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:           "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			reason:             "Traffic problem",
			expectedRepoCalled: true,
			expectedGrpcCalled: false,
			repoReturnError:    errors.New("redis connection error"),
			expectedError:      "redis connection error",
		},
		{
			name:                "failed when grpc return error",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			reason:              "Traffic problem",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     errors.New("grpc called error"),
			grpcReturnIsSuccess: true,
			expectedError:       "grpc called error",
		},
		{
			name:                "failed when grpc called unsuccessfully",
			ctx:                 context.Background(),
			tripId:              "55fe0a54-76e7-4802-9ecd-a9bde5464ea4",
			driverId:            "fbd3ed3e-c385-46cd-ad93-975bbe174488",
			reason:              "Traffic problem",
			expectedRepoCalled:  true,
			expectedGrpcCalled:  true,
			repoReturnError:     nil,
			grpcReturnError:     nil,
			grpcReturnIsSuccess: false,
			expectedError:       "Not allowed to cancel this trip",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockLocationRepository)
			mockGrpcClient := new(MockTripGrpcClient)
			service := services.NewLocationService(mockGrpcClient, mockRepo)

			if tc.expectedRepoCalled {
				mockRepo.On("OnCancel", tc.driverId).Return(tc.repoReturnError)
			}
			if tc.expectedGrpcCalled {
				mockGrpcClient.On("CancelTrip", tc.ctx, tc.tripId, tc.driverId, tc.reason).Return(tc.grpcReturnIsSuccess, tc.grpcReturnError)
			}

			err := service.CancelTrip(tc.ctx, tc.tripId, tc.driverId, tc.reason)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
			mockGrpcClient.AssertExpectations(t)
		})
	}
}
