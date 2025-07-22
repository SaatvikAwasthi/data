package command_test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
	"gotest.tools/assert"

	"tester/app/operation"
	"tester/app/operation/command"
	"tester/domain"
	mockContract "tester/mock/app/operation/contract"
)

func TestAddRawDataCommand_Handle(t *testing.T) {
	// Define test cases
	ctx := context.Background()
	posts := domain.Posts{
		{
			UserId: 1,
			Title:  "Test Title",
			Body:   "Test Body",
		},
	}

	testCases := []struct {
		name          string
		setupMock     func(repository *mockContract.MockRawDataRepository, serviceProvider *mockContract.MockPostsServiceProvider)
		expectedResp  operation.AddRawDataResponse
		expectedError error
		assertions    func(result operation.AddRawDataResponse, err error)
	}{
		{
			name: "success fetch and store posts",
			setupMock: func(repository *mockContract.MockRawDataRepository, serviceProvider *mockContract.MockPostsServiceProvider) {
				serviceProvider.EXPECT().Fetch(ctx).Return(posts, nil)
				repository.EXPECT().Add(ctx, gomock.Any()).Return(nil)
			},
			expectedResp: operation.AddRawDataResponse{
				Message: "Successfully fetched and stored posts in db",
			},
			expectedError: nil,
			assertions: func(result operation.AddRawDataResponse, err error) {
				assert.NilError(t, err)
				assert.DeepEqual(t, operation.AddRawDataResponse{
					Message: "Successfully fetched and stored posts in db",
				}, result)
			},
		},
		{
			name: "failed to fetch posts",
			setupMock: func(repository *mockContract.MockRawDataRepository, serviceProvider *mockContract.MockPostsServiceProvider) {
				serviceProvider.EXPECT().Fetch(ctx).Return(nil, errors.New("failed to fetch posts"))
			},
			expectedResp:  operation.AddRawDataResponse{},
			expectedError: nil,
			assertions: func(result operation.AddRawDataResponse, err error) {
				assert.Equal(t, err.Error(), "failed to fetch posts")
				assert.DeepEqual(t, operation.AddRawDataResponse{}, result)
			},
		},
		{
			name: "failed to store posts",
			setupMock: func(repository *mockContract.MockRawDataRepository, serviceProvider *mockContract.MockPostsServiceProvider) {
				serviceProvider.EXPECT().Fetch(ctx).Return(posts, nil)
				repository.EXPECT().Add(ctx, gomock.Any()).Return(errors.New("failed to store posts"))
			},
			expectedResp:  operation.AddRawDataResponse{},
			expectedError: nil,
			assertions: func(result operation.AddRawDataResponse, err error) {
				assert.Equal(t, err.Error(), "failed to store posts")
				assert.DeepEqual(t, operation.AddRawDataResponse{}, result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			mockRepository := mockContract.NewMockRawDataRepository(ctrl)
			mockServiceProvider := mockContract.NewMockPostsServiceProvider(ctrl)
			tc.setupMock(mockRepository, mockServiceProvider)
			c := command.NewAddRawDataCommand(mockRepository, mockServiceProvider)

			// Act
			result, err := c.Handle(ctx)

			// Assert
			tc.assertions(result, err)
		})
	}
}
