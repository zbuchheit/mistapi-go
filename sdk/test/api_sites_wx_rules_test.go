/*
Mist API

Testing SitesWxRulesAPIService

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

func Test_mistapigo_SitesWxRulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesWxRulesAPIService CreateSiteWxRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxRulesAPI.CreateSiteWxRule(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxRulesAPIService DeleteSiteWxRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxruleId string

		httpRes, err := apiClient.SitesWxRulesAPI.DeleteSiteWxRule(context.Background(), siteId, wxruleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxRulesAPIService GetSiteWxRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxruleId string

		resp, httpRes, err := apiClient.SitesWxRulesAPI.GetSiteWxRule(context.Background(), siteId, wxruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxRulesAPIService GetSiteWxRulesDerived", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxRulesAPI.GetSiteWxRulesDerived(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxRulesAPIService GetSiteWxRulesUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxRulesAPI.GetSiteWxRulesUsage(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxRulesAPIService ListSiteWxRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string

		resp, httpRes, err := apiClient.SitesWxRulesAPI.ListSiteWxRules(context.Background(), siteId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesWxRulesAPIService UpdateSiteWxRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var siteId string
		var wxruleId string

		resp, httpRes, err := apiClient.SitesWxRulesAPI.UpdateSiteWxRule(context.Background(), siteId, wxruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
