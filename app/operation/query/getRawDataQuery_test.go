package query_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/mock/gomock"
	"gotest.tools/assert"

	"tester/app/operation"
	query2 "tester/app/operation/query"
	"tester/crosscutting/util"
	mockContract "tester/mock/app/operation/contract"
	"tester/persistance/mongo/dao"
)

func TestGetRawData_Handle(t *testing.T) {
	// Define test cases
	mockID := bson.NewObjectID()
	mockTime := time.Now().UTC().String()
	ctx := context.Background()
	mockRawData := &dao.RawData{
		ID:        mockID,
		Source:    "test-source",
		CreatedAt: mockTime,
		Posts: []dao.Post{
			{
				UserId: 1,
				Title:  "Test Title",
				Body:   "Test Body",
			},
		},
	}
	expectedData := operation.RawDataResponse{
		Data: []operation.RawData{
			{
				ID:        mockID.String(),
				Source:    "test-source",
				CreatedAt: mockTime,
				Posts: []operation.Post{
					{
						UserId: 1,
						Title:  "Test Title",
						Body:   "Test Body",
					},
				},
			},
		},
	}

	testCases := []struct {
		name          string
		setupMock     func(provider *mockContract.MockRawDataProvider)
		args          query2.GetRawDataQueryRequest
		expectedResp  operation.RawDataResponse
		expectedError error
		assertions    func(result operation.RawDataResponse, err error)
	}{
		{
			name: "success with data by ID",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetByID(ctx, mockID.String()).Return(mockRawData, nil)
			},
			args: query2.GetRawDataQueryRequest{
				Key:   query2.Id,
				Value: mockID.String(),
			},
			expectedResp:  expectedData,
			expectedError: nil,
			assertions: func(result operation.RawDataResponse, err error) {
				assert.NilError(t, err)
				assert.Check(t, len(result.Data) == 1, "Expected one raw data item")
				assert.Equal(t, "test-source", result.Data[0].Source)
				assert.Check(t, len(result.Data[0].Posts) == 1, "Expected one post in raw data")
				assert.Equal(t, uint64(1), result.Data[0].Posts[0].UserId)
				assert.Equal(t, "Test Title", result.Data[0].Posts[0].Title)
				assert.Equal(t, "Test Body", result.Data[0].Posts[0].Body)
			},
		},
		{
			name: "success with data by created at time",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetByTime(ctx, mockTime).Return(mockRawData, nil)
			},
			args: query2.GetRawDataQueryRequest{
				Key:   query2.CreatedAt,
				Value: mockTime,
			},
			expectedResp:  expectedData,
			expectedError: nil,
			assertions: func(result operation.RawDataResponse, err error) {
				assert.NilError(t, err)
				assert.Check(t, len(result.Data) == 1, "Expected one raw data item")
				assert.Equal(t, "test-source", result.Data[0].Source)
				assert.Check(t, len(result.Data[0].Posts) == 1, "Expected one post in raw data")
				assert.Equal(t, uint64(1), result.Data[0].Posts[0].UserId)
				assert.Equal(t, "Test Title", result.Data[0].Posts[0].Title)
				assert.Equal(t, "Test Body", result.Data[0].Posts[0].Body)
			},
		},
		{
			name: "provider returns error for id",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetByID(ctx, mockID.String()).Return(nil, errors.New("provider error"))
			},
			args: query2.GetRawDataQueryRequest{
				Key:   query2.Id,
				Value: mockID.String(),
			},
			expectedResp:  operation.RawDataResponse{},
			expectedError: errors.New("provider error"),
			assertions: func(result operation.RawDataResponse, err error) {
				assert.Equal(t, "failed to get raw data by id: provider error", err.Error())
				assert.DeepEqual(t, operation.RawDataResponse{}, result)
			},
		},
		{
			name: "provider returns error for created at time",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetByTime(ctx, mockTime).Return(nil, errors.New("provider error"))
			},
			args: query2.GetRawDataQueryRequest{
				Key:   query2.CreatedAt,
				Value: mockTime,
			},
			expectedResp:  operation.RawDataResponse{},
			expectedError: errors.New("provider error"),
			assertions: func(result operation.RawDataResponse, err error) {
				assert.Equal(t, "failed to get raw data by time: provider error", err.Error())
				assert.DeepEqual(t, operation.RawDataResponse{}, result)
			},
		},
		{
			name: "provider returns no data for id",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetByID(ctx, mockID.String()).Return(nil, nil)
			},
			args: query2.GetRawDataQueryRequest{
				Key:   query2.Id,
				Value: mockID.String(),
			},
			expectedResp:  operation.RawDataResponse{},
			expectedError: errors.New("provider error"),
			assertions: func(result operation.RawDataResponse, err error) {
				assert.Equal(t, util.Format("raw data not found for id: %s", mockID.String()), err.Error())
				assert.DeepEqual(t, operation.RawDataResponse{}, result)
			},
		},
		{
			name: "provider returns no data for created at time",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetByTime(ctx, mockTime).Return(nil, nil)
			},
			args: query2.GetRawDataQueryRequest{
				Key:   query2.CreatedAt,
				Value: mockTime,
			},
			expectedResp:  operation.RawDataResponse{},
			expectedError: errors.New("provider error"),
			assertions: func(result operation.RawDataResponse, err error) {
				assert.Equal(t, util.Format("raw data not found for time: %s", mockTime), err.Error())
				assert.DeepEqual(t, operation.RawDataResponse{}, result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			mockProvider := mockContract.NewMockRawDataProvider(ctrl)
			tc.setupMock(mockProvider)
			q := query2.NewGetRawDataQuery(mockProvider)

			// Act
			result, err := q.Handle(ctx, tc.args)

			// Assert
			tc.assertions(result, err)
		})
	}
}
