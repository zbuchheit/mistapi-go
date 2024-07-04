/*
Mist API

Testing SitesMapsAutoPlacementAPIService

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

func Test_mistapigo_SitesMapsAutoPlacementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesMapsAutoPlacementAPIService ClearSiteApAutoplacement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesMapsAutoPlacementAPI.ClearSiteApAutoplacement(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMapsAutoPlacementAPIService ConfirmSiteApLocalizationData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesMapsAutoPlacementAPI.ConfirmSiteApLocalizationData(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMapsAutoPlacementAPIService DeleteSiteApAutoplacement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesMapsAutoPlacementAPI.DeleteSiteApAutoplacement(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMapsAutoPlacementAPIService GetSiteApAutoPlacement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mapId string

		resp, httpRes, err := apiClient.SitesMapsAutoPlacementAPI.GetSiteApAutoPlacement(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMapsAutoPlacementAPIService RunSiteApAutoplacement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mapId string

		httpRes, err := apiClient.SitesMapsAutoPlacementAPI.RunSiteApAutoplacement(context.Background(), siteId, mapId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
