/*
IP Address Management API

Testing ConfigProfileAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ipam

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

func TestConfigProfileAPIService(t *testing.T) {

	apiClient := ipam.NewAPIClient()

	t.Run("Test ConfigProfileAPIService AssociateConfigProfileToObjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConfigProfileAPI.AssociateConfigProfileToObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigProfileAPIService AssociateObjectToConfigProfiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConfigProfileAPI.AssociateObjectToConfigProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigProfileAPIService DisassociateConfigProfileFromObjects", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConfigProfileAPI.DisassociateConfigProfileFromObjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigProfileAPIService DisassociateObjectFromConfigProfiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConfigProfileAPI.DisassociateObjectFromConfigProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigProfileAPIService ListConfigProfiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConfigProfileAPI.ListConfigProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigProfileAPIService ListSubnets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConfigProfileAPI.ListSubnets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
