/*
Mist API

Testing MSPsSSORolesAPIService

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

func Test_mistsdkgo_MSPsSSORolesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MSPsSSORolesAPIService CreateMspSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsSSORolesAPI.CreateMspSsoRole(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsSSORolesAPIService DeleteMspSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string
		var ssoroleId string

		httpRes, err := apiClient.MSPsSSORolesAPI.DeleteMspSsoRole(context.Background(), mspId, ssoroleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsSSORolesAPIService ListMspSsoRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsSSORolesAPI.ListMspSsoRoles(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsSSORolesAPIService UpdateMspSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string
		var ssoroleId string

		resp, httpRes, err := apiClient.MSPsSSORolesAPI.UpdateMspSsoRole(context.Background(), mspId, ssoroleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
