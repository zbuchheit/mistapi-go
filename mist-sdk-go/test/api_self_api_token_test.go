/*
Mist API

Testing SelfAPITokenAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistsdkgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistsdkgo"
)

func Test_mistsdkgo_SelfAPITokenAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SelfAPITokenAPIService CreateApiToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SelfAPITokenAPI.CreateApiToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SelfAPITokenAPIService DeleteApiToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apitokenId string

		httpRes, err := apiClient.SelfAPITokenAPI.DeleteApiToken(context.Background(), apitokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SelfAPITokenAPIService GetApiToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apitokenId string

		resp, httpRes, err := apiClient.SelfAPITokenAPI.GetApiToken(context.Background(), apitokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SelfAPITokenAPIService ListApiTokens", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SelfAPITokenAPI.ListApiTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SelfAPITokenAPIService UpdateApiToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apitokenId string

		resp, httpRes, err := apiClient.SelfAPITokenAPI.UpdateApiToken(context.Background(), apitokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
