/*
DFP API

Testing DfpAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dfp

import (
	"context"
	"github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/dfp"
)

func Test_dfp_DfpAPIService(t *testing.T) {

	configuration := client.Configuration{
		ClientName: "acceptance-test",
	}
	c, err := client.NewAPIClient(configuration)
	require.Nil(t, err)
	apiClient := c.DNSForwardingProxyAPI

	t.Run("Test DfpAPIService DfpCreateOrUpdateDfp", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		//var id int32

		resp, httpRes, err := apiClient.DfpAPI.DfpCreateOrUpdateDfp(context.Background(), 1310452).Body(openapiclient.AtcdfpDfpCreateOrUpdatePayload{
			DefaultResolvers:    nil,
			ForwardingPolicy:    nil,
			Host:                nil,
			Id:                  ipam.PtrInt32(1310452),
			InternalDomainLists: nil,
			Name:                nil,
			PopRegionId:         ipam.PtrInt32(1),
			Resolvers:           nil,
			ResolversAll:        nil,
			ServiceId:           nil,
			ServiceName:         nil,
			SiteId:              nil,
		}).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

	})

	t.Run("Test DfpAPIService DfpListDfp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DfpAPI.DfpListDfp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DfpAPIService DfpReadDfp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.DfpAPI.DfpReadDfp(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
