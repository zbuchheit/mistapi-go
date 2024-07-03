/*
Mist API

Testing OrgsMxClustersAPIService

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

func Test_mistsdkgo_OrgsMxClustersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsMxClustersAPIService CreateOrgMxEdgeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxClustersAPI.CreateOrgMxEdgeCluster(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxClustersAPIService DeleteOrgMxEdgeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxclusterId string

		httpRes, err := apiClient.OrgsMxClustersAPI.DeleteOrgMxEdgeCluster(context.Background(), orgId, mxclusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxClustersAPIService GetOrgMxEdgeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxclusterId string

		resp, httpRes, err := apiClient.OrgsMxClustersAPI.GetOrgMxEdgeCluster(context.Background(), orgId, mxclusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxClustersAPIService ListOrgMxEdgeClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsMxClustersAPI.ListOrgMxEdgeClusters(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsMxClustersAPIService UpdateOrgMxEdgeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var mxclusterId string

		resp, httpRes, err := apiClient.OrgsMxClustersAPI.UpdateOrgMxEdgeCluster(context.Background(), orgId, mxclusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
