/*
Mist API

Testing OrgsRFTemplatesAPIService

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

func Test_mistsdkgo_OrgsRFTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsRFTemplatesAPIService CreateOrgRfTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsRFTemplatesAPI.CreateOrgRfTemplate(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsRFTemplatesAPIService DeleteOrgRfTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var rftemplateId string

		httpRes, err := apiClient.OrgsRFTemplatesAPI.DeleteOrgRfTemplate(context.Background(), orgId, rftemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsRFTemplatesAPIService GetOrgRfTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var rftemplateId string

		resp, httpRes, err := apiClient.OrgsRFTemplatesAPI.GetOrgRfTemplate(context.Background(), orgId, rftemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsRFTemplatesAPIService ListOrgRfTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsRFTemplatesAPI.ListOrgRfTemplates(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsRFTemplatesAPIService UpdateOrgRfTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var rftemplateId string

		resp, httpRes, err := apiClient.OrgsRFTemplatesAPI.UpdateOrgRfTemplate(context.Background(), orgId, rftemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
