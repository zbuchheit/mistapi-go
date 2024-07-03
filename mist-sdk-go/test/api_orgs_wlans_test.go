/*
Mist API

Testing OrgsWlansAPIService

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

func Test_mistsdkgo_OrgsWlansAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsWlansAPIService CreateOrgWlan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsWlansAPI.CreateOrgWlan(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService DeleteOrgWlan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var wlanId string

		httpRes, err := apiClient.OrgsWlansAPI.DeleteOrgWlan(context.Background(), orgId, wlanId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService DeleteOrgWlanPortalImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var wlanId string

		httpRes, err := apiClient.OrgsWlansAPI.DeleteOrgWlanPortalImage(context.Background(), orgId, wlanId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService GetOrgWLAN", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var wlanId string

		resp, httpRes, err := apiClient.OrgsWlansAPI.GetOrgWLAN(context.Background(), orgId, wlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService ListOrgWlans", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsWlansAPI.ListOrgWlans(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService UpdateOrgWlan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var wlanId string

		resp, httpRes, err := apiClient.OrgsWlansAPI.UpdateOrgWlan(context.Background(), orgId, wlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService UpdateOrgWlanPortalTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var wlanId string

		resp, httpRes, err := apiClient.OrgsWlansAPI.UpdateOrgWlanPortalTemplate(context.Background(), orgId, wlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsWlansAPIService UploadOrgWlanPortalImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var wlanId string

		httpRes, err := apiClient.OrgsWlansAPI.UploadOrgWlanPortalImage(context.Background(), orgId, wlanId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
