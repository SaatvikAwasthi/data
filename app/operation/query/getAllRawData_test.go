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
	"tester/app/operation/query"
	mockContract "tester/mock/app/operation/contract"
	"tester/persistance/mongo/dao"
)

func TestGetAllRawData_Handle(t *testing.T) {
	// Define test cases
	mockID := bson.NewObjectID()
	mockTime := time.Now().UTC().String()
	ctx := context.Background()

	testCases := []struct {
		name          string
		setupMock     func(provider *mockContract.MockRawDataProvider)
		expectedResp  operation.RawDataResponse
		expectedError error
		assertions    func(result operation.RawDataResponse, err error)
	}{
		{
			name: "success with data",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				mockRawData := []dao.RawData{
					{
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
					},
				}
				provider.EXPECT().GetAll(ctx).Return(mockRawData, nil)
			},
			expectedResp: operation.RawDataResponse{
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
			},
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
			name: "provider returns error",
			setupMock: func(provider *mockContract.MockRawDataProvider) {
				provider.EXPECT().GetAll(ctx).Return(nil, errors.New("provider error"))
			},
			expectedResp:  operation.RawDataResponse{},
			expectedError: errors.New("provider error"),
			assertions: func(result operation.RawDataResponse, err error) {
				assert.Equal(t, "provider error", err.Error())
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
			q := query.NewGetAllRawDataQuery(mockProvider)

			// Act
			result, err := q.Handle(ctx)

			// Assert
			tc.assertions(result, err)
		})
	}
}
