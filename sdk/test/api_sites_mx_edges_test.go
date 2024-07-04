/*
Mist API

Testing SitesMxEdgesAPIService

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

func Test_mistapigo_SitesMxEdgesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesMxEdgesAPIService CountSiteMxEdgeEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesMxEdgesAPI.CountSiteMxEdgeEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService CreateSiteMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesMxEdgesAPI.CreateSiteMxEdge(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService DeleteSiteMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mxedgeId string

		httpRes, err := apiClient.SitesMxEdgesAPI.DeleteSiteMxEdge(context.Background(), siteId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService GetSiteMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mxedgeId string

		httpRes, err := apiClient.SitesMxEdgesAPI.GetSiteMxEdge(context.Background(), siteId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService ListSiteMxEdges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesMxEdgesAPI.ListSiteMxEdges(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService SearchSiteMistEdgeEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesMxEdgesAPI.SearchSiteMistEdgeEvents(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService UpdateSiteMxEdge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mxedgeId string

		resp, httpRes, err := apiClient.SitesMxEdgesAPI.UpdateSiteMxEdge(context.Background(), siteId, mxedgeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesMxEdgesAPIService UploadSiteMxEdgeSupportFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var mxedgeId string

		httpRes, err := apiClient.SitesMxEdgesAPI.UploadSiteMxEdgeSupportFiles(context.Background(), siteId, mxedgeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
