// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package firehydrant

import (
	"context"
	"firehydrant/internal/hooks"
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"firehydrant/retry"
	"fmt"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.firehydrant.io/",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

type FireHydrant struct {
	AccountSettings *AccountSettings
	// Operations related to Incidents
	Incidents *Incidents
	// Operations related to Alerts
	Alerts *Alerts
	// Operations related to Services
	Services *Services
	// Operations related to Changes
	Changes *Changes
	// Operations related to Tasks
	Tasks              *Tasks
	ChecklistTemplates *ChecklistTemplates
	// Operations related to Conversations
	Conversations *Conversations
	// Operations related to Users
	Users            *Users
	IncidentSettings *IncidentSettings
	// Operations related to Environments
	Environments *Environments
	// Operations related to Functionalities
	Functionalities *Functionalities
	StatusPages     *StatusPages
	Infrastructures *Infrastructures
	// Operations related to Integrations
	Integrations             *Integrations
	AwsCloudtrailBatchEvents *AwsCloudtrailBatchEvents
	AwsConnections           *AwsConnections
	Confluence               *Confluence
	Slack                    *Slack
	Statuspage               *Statuspage
	Zendesk                  *Zendesk
	MetricsReporting         *MetricsReporting
	Metrics                  *Metrics
	System                   *System
	// Operations related to Retrospectives
	Retrospectives *Retrospectives
	// Operations related to Runbooks
	Runbooks     *Runbooks
	Maintenances *Maintenances
	Teams        *Teams
	Scim         *Scim
	// Operations related to Signals
	Signals *Signals
	// Operations related to Communication
	Communication         *Communication
	StatusUpdateTemplates *StatusUpdateTemplates
	OnCallSchedules       *OnCallSchedules
	TicketingPriorities   *TicketingPriorities
	Ticketing             *Ticketing
	ProjectConfigurations *ProjectConfigurations
	Tickets               *Tickets
	// Operations related to Webhooks
	Webhooks *Webhooks

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*FireHydrant)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *FireHydrant) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *FireHydrant) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *FireHydrant) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *FireHydrant) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKey string) SDKOption {
	return func(sdk *FireHydrant) {
		security := components.Security{APIKey: apiKey}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *FireHydrant) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *FireHydrant) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *FireHydrant) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *FireHydrant {
	sdk := &FireHydrant{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "0.4.5",
			GenVersion:        "2.455.2",
			UserAgent:         "speakeasy-sdk/go 0.4.5 2.455.2 0.0.1 firehydrant",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.AccountSettings = newAccountSettings(sdk.sdkConfiguration)

	sdk.Incidents = newIncidents(sdk.sdkConfiguration)

	sdk.Alerts = newAlerts(sdk.sdkConfiguration)

	sdk.Services = newServices(sdk.sdkConfiguration)

	sdk.Changes = newChanges(sdk.sdkConfiguration)

	sdk.Tasks = newTasks(sdk.sdkConfiguration)

	sdk.ChecklistTemplates = newChecklistTemplates(sdk.sdkConfiguration)

	sdk.Conversations = newConversations(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.IncidentSettings = newIncidentSettings(sdk.sdkConfiguration)

	sdk.Environments = newEnvironments(sdk.sdkConfiguration)

	sdk.Functionalities = newFunctionalities(sdk.sdkConfiguration)

	sdk.StatusPages = newStatusPages(sdk.sdkConfiguration)

	sdk.Infrastructures = newInfrastructures(sdk.sdkConfiguration)

	sdk.Integrations = newIntegrations(sdk.sdkConfiguration)

	sdk.AwsCloudtrailBatchEvents = newAwsCloudtrailBatchEvents(sdk.sdkConfiguration)

	sdk.AwsConnections = newAwsConnections(sdk.sdkConfiguration)

	sdk.Confluence = newConfluence(sdk.sdkConfiguration)

	sdk.Slack = newSlack(sdk.sdkConfiguration)

	sdk.Statuspage = newStatuspage(sdk.sdkConfiguration)

	sdk.Zendesk = newZendesk(sdk.sdkConfiguration)

	sdk.MetricsReporting = newMetricsReporting(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.System = newSystem(sdk.sdkConfiguration)

	sdk.Retrospectives = newRetrospectives(sdk.sdkConfiguration)

	sdk.Runbooks = newRunbooks(sdk.sdkConfiguration)

	sdk.Maintenances = newMaintenances(sdk.sdkConfiguration)

	sdk.Teams = newTeams(sdk.sdkConfiguration)

	sdk.Scim = newScim(sdk.sdkConfiguration)

	sdk.Signals = newSignals(sdk.sdkConfiguration)

	sdk.Communication = newCommunication(sdk.sdkConfiguration)

	sdk.StatusUpdateTemplates = newStatusUpdateTemplates(sdk.sdkConfiguration)

	sdk.OnCallSchedules = newOnCallSchedules(sdk.sdkConfiguration)

	sdk.TicketingPriorities = newTicketingPriorities(sdk.sdkConfiguration)

	sdk.Ticketing = newTicketing(sdk.sdkConfiguration)

	sdk.ProjectConfigurations = newProjectConfigurations(sdk.sdkConfiguration)

	sdk.Tickets = newTickets(sdk.sdkConfiguration)

	sdk.Webhooks = newWebhooks(sdk.sdkConfiguration)

	return sdk
}
