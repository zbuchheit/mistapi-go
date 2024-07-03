/*
Mist API

Testing SitesPsksAPIService

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

func Test_mistsdkgo_SitesPsksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesPsksAPIService CreateSitePsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesPsksAPI.CreateSitePsk(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesPsksAPIService DeleteSitePsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var pskId string

		httpRes, err := apiClient.SitesPsksAPI.DeleteSitePsk(context.Background(), siteId, pskId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesPsksAPIService GetSitePsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var pskId string

		resp, httpRes, err := apiClient.SitesPsksAPI.GetSitePsk(context.Background(), siteId, pskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesPsksAPIService ImportSitePsks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesPsksAPI.ImportSitePsks(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesPsksAPIService ListSitePsks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesPsksAPI.ListSitePsks(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesPsksAPIService UpdateSiteMultiplePsks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesPsksAPI.UpdateSiteMultiplePsks(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesPsksAPIService UpdateSitePsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var pskId string

		resp, httpRes, err := apiClient.SitesPsksAPI.UpdateSitePsk(context.Background(), siteId, pskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
