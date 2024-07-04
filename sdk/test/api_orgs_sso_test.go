/*
Mist API

Testing OrgsSSOAPIService

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

func Test_mistapigo_OrgsSSOAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsSSOAPIService CreateOrgSso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.CreateOrgSso(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService DeleteOrgSso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoId string

		httpRes, err := apiClient.OrgsSSOAPI.DeleteOrgSso(context.Background(), orgId, ssoId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService DownloadOrgSsoSamlMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.DownloadOrgSsoSamlMetadata(context.Background(), orgId, ssoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService GetOrgSso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.GetOrgSso(context.Background(), orgId, ssoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService GetOrgSsoSamlMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.GetOrgSsoSamlMetadata(context.Background(), orgId, ssoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService ListOrgSsoLatestFailures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.ListOrgSsoLatestFailures(context.Background(), orgId, ssoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService ListOrgSsos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.ListOrgSsos(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSOAPIService UpdateOrgSso", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoId string

		resp, httpRes, err := apiClient.OrgsSSOAPI.UpdateOrgSso(context.Background(), orgId, ssoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
