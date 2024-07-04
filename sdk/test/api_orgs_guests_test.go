/*
Mist API

Testing OrgsGuestsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistapi-go/sdk"
)

func Test_mistapigo_OrgsGuestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsGuestsAPIService CountOrgGuestAuthorizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsGuestsAPI.CountOrgGuestAuthorizations(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGuestsAPIService DeleteOrgGuestAuthorization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var guestMac string

		httpRes, err := apiClient.OrgsGuestsAPI.DeleteOrgGuestAuthorization(context.Background(), orgId, guestMac).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGuestsAPIService GetOrgGuestAuthorization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var guestMac string

		resp, httpRes, err := apiClient.OrgsGuestsAPI.GetOrgGuestAuthorization(context.Background(), orgId, guestMac).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGuestsAPIService ListOrgGuestAuthorizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsGuestsAPI.ListOrgGuestAuthorizations(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGuestsAPIService SearchOrgGuestAuthorization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsGuestsAPI.SearchOrgGuestAuthorization(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGuestsAPIService UpdateOrgGuestAuthorization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var guestMac string

		resp, httpRes, err := apiClient.OrgsGuestsAPI.UpdateOrgGuestAuthorization(context.Background(), orgId, guestMac).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
