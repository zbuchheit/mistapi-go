/*
Mist API

Testing OrgsGatewayTemplatesAPIService

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

func Test_mistsdkgo_OrgsGatewayTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsGatewayTemplatesAPIService CreateOrgGatewayTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsGatewayTemplatesAPI.CreateOrgGatewayTemplate(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGatewayTemplatesAPIService DeleteOrgGatewayTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var gatewaytemplateId string

		httpRes, err := apiClient.OrgsGatewayTemplatesAPI.DeleteOrgGatewayTemplate(context.Background(), orgId, gatewaytemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGatewayTemplatesAPIService GetOrgGatewayTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var gatewaytemplateId string

		resp, httpRes, err := apiClient.OrgsGatewayTemplatesAPI.GetOrgGatewayTemplate(context.Background(), orgId, gatewaytemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGatewayTemplatesAPIService ListOrgGatewayTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsGatewayTemplatesAPI.ListOrgGatewayTemplates(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsGatewayTemplatesAPIService UpdateOrgGatewayTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var gatewaytemplateId string

		resp, httpRes, err := apiClient.OrgsGatewayTemplatesAPI.UpdateOrgGatewayTemplate(context.Background(), orgId, gatewaytemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
