// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMetricsReporting_ListMeasurementDefinitions(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.ListMeasurementDefinitions(ctx, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_CreateMeasurementDefinition(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.CreateMeasurementDefinition(ctx, operations.CreateMeasurementDefinitionRequestBody{
		Name:                "<value>",
		StartsAtMilestoneID: "<id>",
		EndsAtMilestoneID:   "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_GetMeasurementDefinition(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.GetMeasurementDefinition(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_DeleteMeasurementDefinition(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.DeleteMeasurementDefinition(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_UpdateMeasurementDefinition(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.UpdateMeasurementDefinition(ctx, "<id>", nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_ListIncidentMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.ListIncidentMetrics(ctx, operations.ListIncidentMetricsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_GetMilestoneFunnelMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.GetMilestoneFunnelMetrics(ctx, operations.GetMilestoneFunnelMetricsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_ListRetrospectiveMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.ListRetrospectives(ctx, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_GetTicketFunnelMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.GetTicketFunnel(ctx, operations.GetTicketFunnelMetricsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_ListUserInvolvementMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.ListUserInvolvementMetrics(ctx, operations.ListUserInvolvementMetricsRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_ListInfrastructureMetrics(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.ListInfrastructureMetrics(ctx, operations.InfraTypeServices, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_GetMeanTimeReport(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.GetMeanTime(ctx, operations.GetMeanTimeReportRequest{})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_ListSavedSearches(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.ListSavedSearches(ctx, operations.ListSavedSearchesRequest{
		ResourceType: operations.ResourceTypeIncidentEvents,
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMetricsReporting_CreateSavedSearch(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.MetricsReporting.CreateSavedSearch(ctx, operations.PathParamResourceTypeTicketFollowUps, components.PostV1SavedSearchesResourceType{
		Name:         "<value>",
		FilterValues: map[string]any{},
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}
