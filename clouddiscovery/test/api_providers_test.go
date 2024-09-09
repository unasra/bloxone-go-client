/*
Discovery Configuration API V2

Testing ProvidersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package clouddiscovery

import (
	"context"
	"fmt"
	"github.com/infobloxopen/bloxone-go-client/client"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
)

func TestProvidersAPIService(t *testing.T) {

	apiClient := client.NewAPIClient()

	var awsId, azureId, gcpId *string

	t.Run("Test ProvidersAPIService Create AWS Provider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		awsresp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Create(context.Background()).
			Body(clouddiscovery.DiscoveryConfig{
				AccountPreference: clouddiscovery.PtrString("single"),
				CredentialPreference: &clouddiscovery.CredentialPreference{
					AccessIdentifierType: clouddiscovery.PtrString("role_arn"),
					CredentialType:       clouddiscovery.PtrString("dynamic"),
				},
				Description:  clouddiscovery.PtrString("AWS Discovery"),
				Name:         "TestSyncAWS1",
				ProviderType: clouddiscovery.PtrString("Amazon Web Services"),
				SourceConfigs: []clouddiscovery.SourceConfig{
					{
						CredentialConfig: &clouddiscovery.CredentialConfig{
							AccessIdentifier: clouddiscovery.PtrString("arn:aws:iam::111111111111:role/infoblox_discovery"),
						},
						AdditionalProperties: nil,
					},
				},
				SyncInterval: clouddiscovery.PtrString("Auto"),
				Tags:         nil,
			}).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, awsresp)
		assert.Equal(t, 201, httpRes.StatusCode)
		awsId = awsresp.GetResult().Id

	})

	t.Run("Test ProvidersAPIService Create Azure Provider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		azureresp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Create(context.Background()).
			Body(clouddiscovery.DiscoveryConfig{
				AccountPreference: clouddiscovery.PtrString("single"),
				CredentialPreference: &clouddiscovery.CredentialPreference{
					AccessIdentifierType: clouddiscovery.PtrString("tenant_id"),
					CredentialType:       clouddiscovery.PtrString("dynamic"),
					AdditionalProperties: nil,
				},
				Description:  clouddiscovery.PtrString("Azure Discovery"),
				Name:         "TestSyncAzure",
				ProviderType: clouddiscovery.PtrString("Microsoft Azure"),
				SourceConfigs: []clouddiscovery.SourceConfig{
					{
						CredentialConfig: &clouddiscovery.CredentialConfig{
							AccessIdentifier: clouddiscovery.PtrString("1111112111"),
						},
						RestrictedToAccounts: []string{"3333333333"},
						AdditionalProperties: nil,
					},
				},
				SyncInterval: clouddiscovery.PtrString("Auto"),
				Tags:         nil,
			}).
			Execute()

		if err != nil {
			t.Errorf("Error while reading AWS provider: %v", err)
		}
		require.Nil(t, err)
		require.NotNil(t, azureresp)
		assert.Equal(t, 201, httpRes.StatusCode)
		azureId = azureresp.GetResult().Id
		//azureSfc = azureresp.GetResult().SourceConfigs[0].Id

	})

	t.Run("Test ProvidersAPIService Create GCP Provider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		gcpresp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Create(context.Background()).
			Body(clouddiscovery.DiscoveryConfig{
				AccountPreference: clouddiscovery.PtrString("single"),
				CredentialPreference: &clouddiscovery.CredentialPreference{
					AccessIdentifierType: clouddiscovery.PtrString("project_id"),
					CredentialType:       clouddiscovery.PtrString("dynamic"),
					AdditionalProperties: nil,
				},
				Description:  clouddiscovery.PtrString("GCP Discovery"),
				Name:         "TestSyncGCP",
				ProviderType: clouddiscovery.PtrString("Google Cloud Platform"),
				SourceConfigs: []clouddiscovery.SourceConfig{
					{
						CredentialConfig: &clouddiscovery.CredentialConfig{
							AccessIdentifier: clouddiscovery.PtrString("33333333333"),
						},
					},
				},
				SyncInterval: clouddiscovery.PtrString("Auto"),
				Tags:         nil,
			}).
			Execute()

		if err != nil {
			t.Errorf("Error while reading AWS provider: %v", err)
		}
		require.Nil(t, err)
		require.NotNil(t, gcpresp)
		assert.Equal(t, 201, httpRes.StatusCode)
		gcpId = gcpresp.GetResult().Id

	})

	t.Run("Test ProvidersAPIService List", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.List(context.Background()).Execute()

		if err != nil {
			t.Errorf("Error while reading AWS provider: %v", err)
		}
		fmt.Println(resp)
		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProvidersAPIService Read", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Read(context.Background(), *gcpId).Execute()

		// print resp
		fmt.Println(resp)
		require.NoError(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ProvidersAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Update(context.Background(), *awsId).
			Body(clouddiscovery.DiscoveryConfig{
				AccountPreference: clouddiscovery.PtrString("single"),
				CredentialPreference: &clouddiscovery.CredentialPreference{
					AccessIdentifierType: clouddiscovery.PtrString("role_arn"),
					CredentialType:       clouddiscovery.PtrString("dynamic"),
					AdditionalProperties: nil,
				},
				Description:  clouddiscovery.PtrString("AWS Discovery"),
				Name:         "TestSyncAWS",
				ProviderType: clouddiscovery.PtrString("Amazon Web Services"),
				SourceConfigs: []clouddiscovery.SourceConfig{
					{
						//Id: az,
						CredentialConfig: &clouddiscovery.CredentialConfig{
							AccessIdentifier: clouddiscovery.PtrString("arn:aws:iam::111111111111:role/infoblox_discovery"),
						},
						AdditionalProperties: nil,
					},
				},
				SyncInterval: clouddiscovery.PtrString("Auto"),
				Tags:         nil,
			}).
			Execute()

		if err != nil {
			t.Errorf("Error while updating AWS provider: %v", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProvidersAPIService Azure Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Update(context.Background(), *azureId).
			Body(clouddiscovery.DiscoveryConfig{
				AccountPreference: clouddiscovery.PtrString("single"),
				Description:       clouddiscovery.PtrString("Azure Discovery meow"),
				Name:              "TestSyncAzure",
				ProviderType:      clouddiscovery.PtrString("Microsoft Azure"),
				SyncInterval:      clouddiscovery.PtrString("Auto"),
				Tags:              nil,
			}).
			Execute()

		if err != nil {
			t.Errorf("Error while updating AWS provider: %v", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProvidersAPIService Update GCP Provider", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		gcpresp, httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Update(context.Background(), *gcpId).
			Body(clouddiscovery.DiscoveryConfig{
				AccountPreference: clouddiscovery.PtrString("single"),
				Description:       clouddiscovery.PtrString("GCP Discovery update"),
				Name:              "TestSyncGCP",
				ProviderType:      clouddiscovery.PtrString("Google Cloud Platform"),
			}).
			Execute()

		if err != nil {
			t.Errorf("Error while reading AWS provider: %v", err)
		}
		require.Nil(t, err)
		require.NotNil(t, gcpresp)
		assert.Equal(t, 200, httpRes.StatusCode)
		//gcpId = gcpresp.GetResult().Id

	})

	t.Run("Test ProvidersAPIService AWS Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Delete(context.Background(), *awsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)

	})

	t.Run("Test ProvidersAPIService Azure Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Delete(context.Background(), *azureId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)

	})

	t.Run("Test ProvidersAPIService GCP Delete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DiscoveryConfigurationAPIV2.ProvidersAPI.Delete(context.Background(), *gcpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)

	})

}