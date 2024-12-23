<div align="center">
    <picture>
        <img src="https://files.readme.io/7fb504d-small-FireHydrant-logo.png" width="300">
    </picture>
   <p>Manage incidents from ring to retro</p>
	<p>Developer-friendly & type-safe Go SDK specifically catered to leverage <strong>FireHydrant</strong> API.</p>
   <a href=""><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
  * [SDK Installation](#sdk-installation)
  * [Authentication](#authentication)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/firehydrant/go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  |
| -------- | ------ | ------- |
| `APIKey` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.AccountSettings.GetAiPreferences(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AIEntitiesPreferencesEntity != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.AccountSettings.GetAiPreferences(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AIEntitiesPreferencesEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountSettings](docs/sdks/accountsettings/README.md)

* [GetAiPreferences](docs/sdks/accountsettings/README.md#getaipreferences) - Get AI preferences
* [UpdateAiPreferences](docs/sdks/accountsettings/README.md#updateaipreferences) - Update AI preferences
* [VoteOnIncidentSummary](docs/sdks/accountsettings/README.md#voteonincidentsummary) - Vote on an AI-generated incident summary
* [GetBootstrap](docs/sdks/accountsettings/README.md#getbootstrap) - Get initial application configuration and settings
* [ListEntitlements](docs/sdks/accountsettings/README.md#listentitlements) - List entitlements
* [Ping](docs/sdks/accountsettings/README.md#ping) - Check API connectivity
* [GetSavedSearch](docs/sdks/accountsettings/README.md#getsavedsearch) - Get a saved search
* [DeleteSavedSearch](docs/sdks/accountsettings/README.md#deletesavedsearch) - Delete a saved search
* [UpdateSavedSearch](docs/sdks/accountsettings/README.md#updatesavedsearch) - Update a saved search

### [Alerts](docs/sdks/alerts/README.md)

* [List](docs/sdks/alerts/README.md#list) - List alerts
* [Get](docs/sdks/alerts/README.md#get) - Get an alert
* [ListForIncident](docs/sdks/alerts/README.md#listforincident) - List alerts for an incident
* [Create](docs/sdks/alerts/README.md#create) - Attach alerts to an incident
* [ListProcessingLogs](docs/sdks/alerts/README.md#listprocessinglogs) - List alert processing log entries

### [AwsCloudtrailBatchEvents](docs/sdks/awscloudtrailbatchevents/README.md)

* [List](docs/sdks/awscloudtrailbatchevents/README.md#list) - List events for an AWS CloudTrail batch

### [AwsConnections](docs/sdks/awsconnections/README.md)

* [List](docs/sdks/awsconnections/README.md#list) - List AWS integration connections
* [Get](docs/sdks/awsconnections/README.md#get) - Get an AWS connection

### [Changes](docs/sdks/changes/README.md)

* [ListTypes](docs/sdks/changes/README.md#listtypes) - List change types
* [List](docs/sdks/changes/README.md#list) - List changes
* [Create](docs/sdks/changes/README.md#create) - Create a change
* [ListEvents](docs/sdks/changes/README.md#listevents) - List change events
* [CreateEvent](docs/sdks/changes/README.md#createevent) - Create a change event
* [GetEvent](docs/sdks/changes/README.md#getevent) - Get a change event
* [DeleteEvent](docs/sdks/changes/README.md#deleteevent) - Delete a change event
* [UpdateEvent](docs/sdks/changes/README.md#updateevent) - Update a change event
* [Delete](docs/sdks/changes/README.md#delete) - Archive a change
* [Update](docs/sdks/changes/README.md#update) - Update a change
* [ListIdentities](docs/sdks/changes/README.md#listidentities) - List identities for a change
* [CreateIdentity](docs/sdks/changes/README.md#createidentity) - Create an identity for a change
* [DeleteIdentity](docs/sdks/changes/README.md#deleteidentity) - Delete an identity from a change
* [UpdateIdentity](docs/sdks/changes/README.md#updateidentity) - Update an identity for a change
* [Get](docs/sdks/changes/README.md#get) - Get a scheduled maintenance event
* [UpdateScheduledMaintenance](docs/sdks/changes/README.md#updatescheduledmaintenance) - Update a scheduled maintenance event

### [ChecklistTemplates](docs/sdks/checklisttemplates/README.md)

* [Get](docs/sdks/checklisttemplates/README.md#get) - Get a checklist template

### [Communication](docs/sdks/communication/README.md)

* [ListTemplates](docs/sdks/communication/README.md#listtemplates) - List status update templates
* [GetTemplate](docs/sdks/communication/README.md#gettemplate) - Get a status update template
* [DeleteStatusUpdateTemplate](docs/sdks/communication/README.md#deletestatusupdatetemplate) - Delete a status update template
* [UpdateTemplate](docs/sdks/communication/README.md#updatetemplate) - Update a status update template

### [Confluence](docs/sdks/confluence/README.md)

* [ListSpaces](docs/sdks/confluence/README.md#listspaces) - List Confluence spaces

### [Conversations](docs/sdks/conversations/README.md)

* [ListComments](docs/sdks/conversations/README.md#listcomments) - List comments for a conversation
* [CreateComment](docs/sdks/conversations/README.md#createcomment) - Create a comment for a conversation
* [GetComment](docs/sdks/conversations/README.md#getcomment) - Get a conversation comment
* [DeleteComment](docs/sdks/conversations/README.md#deletecomment) - Delete a conversation comment
* [UpdateComment](docs/sdks/conversations/README.md#updatecomment) - Update a conversation comment
* [ListCommentReactions](docs/sdks/conversations/README.md#listcommentreactions) - List reactions for a conversation comment
* [CreateCommentReaction](docs/sdks/conversations/README.md#createcommentreaction) - Create a reaction for a conversation comment
* [DeleteCommentReaction](docs/sdks/conversations/README.md#deletecommentreaction) - Delete a reaction from a conversation comment

### [Environments](docs/sdks/environments/README.md)

* [List](docs/sdks/environments/README.md#list) - List environments
* [Create](docs/sdks/environments/README.md#create) - Create an environment
* [Get](docs/sdks/environments/README.md#get) - Get an environment
* [Delete](docs/sdks/environments/README.md#delete) - Archive an environment
* [Update](docs/sdks/environments/README.md#update) - Update an environment


### [Functionalities](docs/sdks/functionalities/README.md)

* [List](docs/sdks/functionalities/README.md#list) - List functionalities
* [Create](docs/sdks/functionalities/README.md#create) - Create a functionality
* [Get](docs/sdks/functionalities/README.md#get) - Get a functionality
* [Delete](docs/sdks/functionalities/README.md#delete) - Archive a functionality
* [Update](docs/sdks/functionalities/README.md#update) - Update a functionality
* [ListServices](docs/sdks/functionalities/README.md#listservices) - List services for a functionality

### [Incidents](docs/sdks/incidents/README.md)

* [GetAiSummaryVoteStatus](docs/sdks/incidents/README.md#getaisummaryvotestatus) - Get the vote status for an AI-generated incident summary
* [List](docs/sdks/incidents/README.md#list) - List incidents
* [Create](docs/sdks/incidents/README.md#create) - Create an incident
* [Get](docs/sdks/incidents/README.md#get) - Get an incident
* [Archive](docs/sdks/incidents/README.md#archive) - Archive an incident
* [Update](docs/sdks/incidents/README.md#update) - Update an incident
* [DeleteAlert](docs/sdks/incidents/README.md#deletealert) - Delete an alert from an incident
* [SetAlertAsPrimary](docs/sdks/incidents/README.md#setalertasprimary) - Set an alert as primary for an incident
* [ListAttachments](docs/sdks/incidents/README.md#listattachments) - List attachments for an incident
* [CreateAttachment](docs/sdks/incidents/README.md#createattachment) - Create an attachment for an incident
* [GetChannel](docs/sdks/incidents/README.md#getchannel) - Get chat channel information for an incident
* [Close](docs/sdks/incidents/README.md#close) - Close an incident
* [ListEvents](docs/sdks/incidents/README.md#listevents) - List events for an incident
* [GetEvent](docs/sdks/incidents/README.md#getevent) - Get an incident event
* [DeleteEvent](docs/sdks/incidents/README.md#deleteevent) - Delete an incident event
* [UpdateEvent](docs/sdks/incidents/README.md#updateevent) - Update an incident event
* [UpdateEventVotes](docs/sdks/incidents/README.md#updateeventvotes) - Update votes for an incident event
* [GetEventVoteStatus](docs/sdks/incidents/README.md#geteventvotestatus) - Get vote counts for an incident event
* [CreateGenericChatMessage](docs/sdks/incidents/README.md#creategenericchatmessage) - Create a chat message for an incident
* [DeleteChatMessage](docs/sdks/incidents/README.md#deletechatmessage) - Delete a chat message from an incident
* [UpdateChatMessage](docs/sdks/incidents/README.md#updatechatmessage) - Update a chat message in an incident
* [UpdateImpacts](docs/sdks/incidents/README.md#updateimpacts) - Replace all impacts for an incident
* [PartialUpdateImpacts](docs/sdks/incidents/README.md#partialupdateimpacts) - Update impacts for an incident
* [ListImpact](docs/sdks/incidents/README.md#listimpact) - List impacted infrastructure for an incident
* [CreateImpact](docs/sdks/incidents/README.md#createimpact) - Add impacted infrastructure to an incident
* [DeleteImpact](docs/sdks/incidents/README.md#deleteimpact) - Remove impacted infrastructure from an incident
* [ListLinks](docs/sdks/incidents/README.md#listlinks) - List links for an incident
* [CreateLink](docs/sdks/incidents/README.md#createlink) - Create a link for an incident
* [UpdateLink](docs/sdks/incidents/README.md#updatelink) - Update an external link for an incident
* [DeleteLink](docs/sdks/incidents/README.md#deletelink) - Delete an external link from an incident
* [ListMilestones](docs/sdks/incidents/README.md#listmilestones) - List milestones for an incident
* [UpdateMilestonesBulk](docs/sdks/incidents/README.md#updatemilestonesbulk) - Bulk update milestone timestamps for an incident
* [CreateNote](docs/sdks/incidents/README.md#createnote) - Create a note for an incident
* [UpdateNote](docs/sdks/incidents/README.md#updatenote) - Update a note for an incident
* [ListRelatedChangeEvents](docs/sdks/incidents/README.md#listrelatedchangeevents) - List related changes for an incident
* [CreateRelatedChange](docs/sdks/incidents/README.md#createrelatedchange) - Add a related change to an incident
* [UpdateRelatedChangeEvent](docs/sdks/incidents/README.md#updaterelatedchangeevent) - Update a related change event for an incident
* [GetRelationships](docs/sdks/incidents/README.md#getrelationships) - List incident relationships
* [Resolve](docs/sdks/incidents/README.md#resolve) - Resolve an incident
* [ListRoleAssignments](docs/sdks/incidents/README.md#listroleassignments) - List role assignments for an incident
* [CreateRoleAssignment](docs/sdks/incidents/README.md#createroleassignment) - Create a role assignment for an incident
* [DeleteRoleAssignment](docs/sdks/incidents/README.md#deleteroleassignment) - Delete a role assignment from an incident
* [ListSimilar](docs/sdks/incidents/README.md#listsimilar) - List similar incidents
* [ListStatusPages](docs/sdks/incidents/README.md#liststatuspages) - List status pages for an incident
* [AddStatusPage](docs/sdks/incidents/README.md#addstatuspage) - Add a status page to an incident
* [CreateTaskList](docs/sdks/incidents/README.md#createtasklist) - Add tasks from a task list to an incident
* [CreateTeamAssignment](docs/sdks/incidents/README.md#createteamassignment) - Assign a team to an incident
* [DeleteTeamAssignment](docs/sdks/incidents/README.md#deleteteamassignment) - Remove a team assignment from an incident
* [GetTranscript](docs/sdks/incidents/README.md#gettranscript) - List transcript messages for an incident
* [DeleteTranscript](docs/sdks/incidents/README.md#deletetranscript) - Delete a transcript from an incident
* [Unarchive](docs/sdks/incidents/README.md#unarchive) - Unarchive an incident
* [GetUserRole](docs/sdks/incidents/README.md#getuserrole) - Get a user's role in an incident

### [IncidentSettings](docs/sdks/incidentsettings/README.md)

* [ListCustomFieldDefinitions](docs/sdks/incidentsettings/README.md#listcustomfielddefinitions) - List custom field definitions
* [CreateCustomFieldDefinition](docs/sdks/incidentsettings/README.md#createcustomfielddefinition) - Create a custom field definition
* [DeleteCustomFieldDefinition](docs/sdks/incidentsettings/README.md#deletecustomfielddefinition) - Delete a custom field definition
* [UpdateCustomFieldDefinition](docs/sdks/incidentsettings/README.md#updatecustomfielddefinition) - Update a custom field definition
* [ListSelectOptions](docs/sdks/incidentsettings/README.md#listselectoptions) - List select options for a custom field
* [GetFormConfiguration](docs/sdks/incidentsettings/README.md#getformconfiguration) - Get a form configuration
* [ListRoles](docs/sdks/incidentsettings/README.md#listroles) - List incident roles
* [CreateIncidentRole](docs/sdks/incidentsettings/README.md#createincidentrole) - Create an incident role
* [GetIncidentRole](docs/sdks/incidentsettings/README.md#getincidentrole) - Get an incident role
* [DeleteRole](docs/sdks/incidentsettings/README.md#deleterole) - Archive an incident role
* [UpdateIncidentRole](docs/sdks/incidentsettings/README.md#updateincidentrole) - Update an incident role
* [ListIncidentTags](docs/sdks/incidentsettings/README.md#listincidenttags) - List incident tags
* [ValidateIncidentTags](docs/sdks/incidentsettings/README.md#validateincidenttags) - Validate incident tags
* [ListIncidentTypes](docs/sdks/incidentsettings/README.md#listincidenttypes) - List incident types
* [CreateIncidentType](docs/sdks/incidentsettings/README.md#createincidenttype) - Create an incident type
* [GetIncidentType](docs/sdks/incidentsettings/README.md#getincidenttype) - Get an incident type
* [ArchiveIncidentType](docs/sdks/incidentsettings/README.md#archiveincidenttype) - Archive an incident type
* [UpdateType](docs/sdks/incidentsettings/README.md#updatetype) - Update an incident type
* [CreateMilestone](docs/sdks/incidentsettings/README.md#createmilestone) - Create a milestone for an incident lifecycle
* [DeleteLifecycleMilestone](docs/sdks/incidentsettings/README.md#deletelifecyclemilestone) - Delete a lifecycle milestone
* [UpdateLifecycleMilestone](docs/sdks/incidentsettings/README.md#updatelifecyclemilestone) - Update a lifecycle milestone
* [ListLifecyclePhases](docs/sdks/incidentsettings/README.md#listlifecyclephases) - List lifecycle phases and milestones
* [ListPriorities](docs/sdks/incidentsettings/README.md#listpriorities) - List priorities
* [CreatePriority](docs/sdks/incidentsettings/README.md#createpriority) - Create a priority
* [GetPriority](docs/sdks/incidentsettings/README.md#getpriority) - Get a priority
* [DeletePriority](docs/sdks/incidentsettings/README.md#deletepriority) - Delete a priority
* [UpdatePriority](docs/sdks/incidentsettings/README.md#updatepriority) - Update a priority
* [ListSeverities](docs/sdks/incidentsettings/README.md#listseverities) - List severities
* [CreateSeverity](docs/sdks/incidentsettings/README.md#createseverity) - Create a severity
* [GetSeverity](docs/sdks/incidentsettings/README.md#getseverity) - Get a severity
* [DeleteSeverity](docs/sdks/incidentsettings/README.md#deleteseverity) - Delete a severity
* [UpdateSeverity](docs/sdks/incidentsettings/README.md#updateseverity) - Update a severity
* [GetSeverityMatrix](docs/sdks/incidentsettings/README.md#getseveritymatrix) - Get severity matrix
* [UpdateSeverityMatrix](docs/sdks/incidentsettings/README.md#updateseveritymatrix) - Update severity matrix
* [ListSeverityMatrixConditions](docs/sdks/incidentsettings/README.md#listseveritymatrixconditions) - List severity matrix conditions
* [CreateSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#createseveritymatrixcondition) - Create a severity matrix condition
* [GetSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#getseveritymatrixcondition) - Get a severity matrix condition
* [DeleteSeverityMatrixCondition](docs/sdks/incidentsettings/README.md#deleteseveritymatrixcondition) - Delete a severity matrix condition
* [UpdateCondition](docs/sdks/incidentsettings/README.md#updatecondition) - Update a severity matrix condition
* [ListSeverityMatrixImpacts](docs/sdks/incidentsettings/README.md#listseveritymatriximpacts) - List severity matrix impacts
* [CreateImpact](docs/sdks/incidentsettings/README.md#createimpact) - Create a severity matrix impact
* [DeleteSeverityMatrixImpact](docs/sdks/incidentsettings/README.md#deleteseveritymatriximpact) - Delete an impact from the severity matrix
* [UpdateImpact](docs/sdks/incidentsettings/README.md#updateimpact) - Update an impact in the severity matrix
* [ListTicketingPriorities](docs/sdks/incidentsettings/README.md#listticketingpriorities) - List ticketing priorities
* [CreateTicketingPriority](docs/sdks/incidentsettings/README.md#createticketingpriority) - Create a ticketing priority

### [Infrastructures](docs/sdks/infrastructures/README.md)

* [List](docs/sdks/infrastructures/README.md#list) - List catalog entries

### [Integrations](docs/sdks/integrations/README.md)

* [List](docs/sdks/integrations/README.md#list) - List all available integrations
* [ListCloudtrailBatches](docs/sdks/integrations/README.md#listcloudtrailbatches) - List AWS CloudTrail batches
* [UpdateCloudTrailBatch](docs/sdks/integrations/README.md#updatecloudtrailbatch) - Update an AWS CloudTrail batch
* [ListConnections](docs/sdks/integrations/README.md#listconnections) - List integration connections
* [CreateConnection](docs/sdks/integrations/README.md#createconnection) - Create a new integration connection
* [UpdateConnection](docs/sdks/integrations/README.md#updateconnection) - Update an integration connection
* [RefreshConnection](docs/sdks/integrations/README.md#refreshconnection) - Refresh an integration connection
* [UpdateFieldMap](docs/sdks/integrations/README.md#updatefieldmap) - Update a field mapping configuration
* [GetFieldMapAvailableFields](docs/sdks/integrations/README.md#getfieldmapavailablefields) - List available fields for field mapping
* [ListEmojiActions](docs/sdks/integrations/README.md#listemojiactions) - List Slack emoji actions
* [GetStatus](docs/sdks/integrations/README.md#getstatus) - Get an integration status
* [GetStatuspageConnection](docs/sdks/integrations/README.md#getstatuspageconnection) - Get a Statuspage connection
* [DeleteStatuspageConnection](docs/sdks/integrations/README.md#deletestatuspageconnection) - Delete a Statuspage connection
* [Get](docs/sdks/integrations/README.md#get) - Get an integration
* [DeletePriority](docs/sdks/integrations/README.md#deletepriority) - Delete a ticketing priority
* [UpdatePriority](docs/sdks/integrations/README.md#updatepriority) - Update a ticketing priority
* [ListProjects](docs/sdks/integrations/README.md#listprojects) - List ticketing projects
* [GetProjectConfigurationOptions](docs/sdks/integrations/README.md#getprojectconfigurationoptions) - List configuration options for a ticketing project
* [GetProjectFieldOptions](docs/sdks/integrations/README.md#getprojectfieldoptions) - List configuration options for a ticketing project field
* [CreateFieldMap](docs/sdks/integrations/README.md#createfieldmap) - Create a field mapping for a ticketing project
* [GetAvailableFields](docs/sdks/integrations/README.md#getavailablefields) - List available fields for ticket field mapping
* [DeleteFieldMap](docs/sdks/integrations/README.md#deletefieldmap) - Delete a field map for a ticketing project
* [UpdateTicketingFieldMap](docs/sdks/integrations/README.md#updateticketingfieldmap) - Update a field map for a ticketing project
* [DeleteProjectConfig](docs/sdks/integrations/README.md#deleteprojectconfig) - Delete a ticketing project configuration
* [GetTicket](docs/sdks/integrations/README.md#getticket) - Get a ticket

#### [Integrations.Aws](docs/sdks/aws/README.md)

* [GetCloudTrailBatch](docs/sdks/aws/README.md#getcloudtrailbatch) - Get an AWS CloudTrail batch
* [UpdateConnection](docs/sdks/aws/README.md#updateconnection) - Update an AWS connection

#### [Integrations.Slack](docs/sdks/firehydrantslack/README.md)

* [GetEmojiAction](docs/sdks/firehydrantslack/README.md#getemojiaction) - Get a Slack emoji action
* [DeleteEmojiAction](docs/sdks/firehydrantslack/README.md#deleteemojiaction) - Delete a Slack emoji action
* [ListWorkspaces](docs/sdks/firehydrantslack/README.md#listworkspaces) - List Slack workspaces for a connection

#### [Integrations.Statuspage](docs/sdks/firehydrantstatuspage/README.md)

* [ListConnections](docs/sdks/firehydrantstatuspage/README.md#listconnections) - List Statuspage connections
* [ListPages](docs/sdks/firehydrantstatuspage/README.md#listpages) - List StatusPage pages for a connection

#### [Integrations.Ticketing](docs/sdks/firehydrantticketing/README.md)

* [GetProject](docs/sdks/firehydrantticketing/README.md#getproject) - Get a ticketing project
* [ListTags](docs/sdks/firehydrantticketing/README.md#listtags) - List ticket tags

### [Maintenances](docs/sdks/maintenances/README.md)

* [List](docs/sdks/maintenances/README.md#list) - List scheduled maintenance events
* [Create](docs/sdks/maintenances/README.md#create) - Create a scheduled maintenance event
* [Delete](docs/sdks/maintenances/README.md#delete) - Delete a scheduled maintenance event

### [Metrics](docs/sdks/metrics/README.md)

* [GetMttx](docs/sdks/metrics/README.md#getmttx) - Fetch infrastructure metrics based on custom query
* [GetInfrastructure](docs/sdks/metrics/README.md#getinfrastructure) - Get metrics for a specific catalog entry

### [MetricsReporting](docs/sdks/metricsreporting/README.md)

* [ListMeasurementDefinitions](docs/sdks/metricsreporting/README.md#listmeasurementdefinitions) - List measurement definitions
* [CreateMeasurementDefinition](docs/sdks/metricsreporting/README.md#createmeasurementdefinition) - Create a measurement definition
* [GetMeasurementDefinition](docs/sdks/metricsreporting/README.md#getmeasurementdefinition) - Get a measurement definition
* [DeleteMeasurementDefinition](docs/sdks/metricsreporting/README.md#deletemeasurementdefinition) - Archive a measurement definition
* [UpdateMeasurementDefinition](docs/sdks/metricsreporting/README.md#updatemeasurementdefinition) - Update a measurement definition
* [ListIncidentMetrics](docs/sdks/metricsreporting/README.md#listincidentmetrics) - List incident metrics and analytics
* [GetMilestoneFunnelMetrics](docs/sdks/metricsreporting/README.md#getmilestonefunnelmetrics) - List milestone funnel metrics
* [ListRetrospectives](docs/sdks/metricsreporting/README.md#listretrospectives) - List retrospective metrics for a date range
* [GetTicketFunnel](docs/sdks/metricsreporting/README.md#getticketfunnel) - List ticket funnel metrics
* [ListUserInvolvementMetrics](docs/sdks/metricsreporting/README.md#listuserinvolvementmetrics) - List user involvement metrics
* [ListInfrastructureMetrics](docs/sdks/metricsreporting/README.md#listinfrastructuremetrics) - List metrics for all services, environments, functionalities, or customers
* [GetMeanTime](docs/sdks/metricsreporting/README.md#getmeantime) - Get mean time metrics for incidents
* [ListSavedSearches](docs/sdks/metricsreporting/README.md#listsavedsearches) - List saved searches
* [CreateSavedSearch](docs/sdks/metricsreporting/README.md#createsavedsearch) - Create a saved search

### [OnCallSchedules](docs/sdks/oncallschedules/README.md)

* [Delete](docs/sdks/oncallschedules/README.md#delete) - Delete an on-call schedule for a team

### [ProjectConfigurations](docs/sdks/projectconfigurations/README.md)

* [Create](docs/sdks/projectconfigurations/README.md#create) - Create a ticketing project configuration
* [Get](docs/sdks/projectconfigurations/README.md#get) - Get a ticketing project configuration

### [Retrospectives](docs/sdks/retrospectives/README.md)

* [ListQuestions](docs/sdks/retrospectives/README.md#listquestions) - List retrospective questions
* [UpdateQuestions](docs/sdks/retrospectives/README.md#updatequestions) - Update retrospective questions
* [GetQuestion](docs/sdks/retrospectives/README.md#getquestion) - Get a retrospective question
* [ListReports](docs/sdks/retrospectives/README.md#listreports) - List retrospective reports
* [CreateReport](docs/sdks/retrospectives/README.md#createreport) - Create a retrospective report
* [GetReport](docs/sdks/retrospectives/README.md#getreport) - Get a retrospective report
* [UpdateReport](docs/sdks/retrospectives/README.md#updatereport) - Update a retrospective report
* [UpdateField](docs/sdks/retrospectives/README.md#updatefield) - Update a retrospective field
* [PublishReport](docs/sdks/retrospectives/README.md#publishreport) - Publish a retrospective report
* [ListReportReasons](docs/sdks/retrospectives/README.md#listreportreasons) - List contributing factors for a retrospective report
* [CreateReason](docs/sdks/retrospectives/README.md#createreason) - Create a contributing factor for a retrospective report
* [UpdateReportReasonOrder](docs/sdks/retrospectives/README.md#updatereportreasonorder) - Update the order of contributing factors in a retrospective report
* [DeleteReason](docs/sdks/retrospectives/README.md#deletereason) - Delete a contributing factor from a retrospective report
* [UpdateReason](docs/sdks/retrospectives/README.md#updatereason) - Update a contributing factor in a retrospective report

### [Runbooks](docs/sdks/runbooks/README.md)

* [ListAudits](docs/sdks/runbooks/README.md#listaudits) - List runbook audits
* [List](docs/sdks/runbooks/README.md#list) - List runbooks
* [Create](docs/sdks/runbooks/README.md#create) - Create a runbook
* [ListActions](docs/sdks/runbooks/README.md#listactions) - List runbook actions
* [ListExecutions](docs/sdks/runbooks/README.md#listexecutions) - List runbook executions
* [CreateExecution](docs/sdks/runbooks/README.md#createexecution) - Create a runbook execution
* [GetExecution](docs/sdks/runbooks/README.md#getexecution) - Get a runbook execution
* [UpdateExecutionStep](docs/sdks/runbooks/README.md#updateexecutionstep) - Update a runbook execution step
* [GetExecutionStepScript](docs/sdks/runbooks/README.md#getexecutionstepscript) - Get a runbook execution step script
* [UpdateExecutionStepScriptState](docs/sdks/runbooks/README.md#updateexecutionstepscriptstate) - Update the script state for a runbook execution step
* [UpdateExecutionStepVotes](docs/sdks/runbooks/README.md#updateexecutionstepvotes) - Update votes for a runbook execution step
* [GetStepVoteStatus](docs/sdks/runbooks/README.md#getstepvotestatus) - Get vote counts for a runbook step
* [UpdateExecutionVotes](docs/sdks/runbooks/README.md#updateexecutionvotes) - Vote on a runbook execution
* [GetExecutionVoteStatus](docs/sdks/runbooks/README.md#getexecutionvotestatus) - Get vote counts for a runbook execution
* [ListSelectOptions](docs/sdks/runbooks/README.md#listselectoptions) - List select options for a runbook integration action field
* [Get](docs/sdks/runbooks/README.md#get) - Get a runbook
* [Update](docs/sdks/runbooks/README.md#update) - Update a runbook
* [Delete](docs/sdks/runbooks/README.md#delete) - Delete a runbook

#### [Runbooks.Executions](docs/sdks/executions/README.md)

* [Delete](docs/sdks/executions/README.md#delete) - Terminate a runbook execution

### [Scim](docs/sdks/scim/README.md)

* [ListGroups](docs/sdks/scim/README.md#listgroups) - List teams via SCIM
* [Create](docs/sdks/scim/README.md#create) - Create a team via SCIM
* [GetGroup](docs/sdks/scim/README.md#getgroup) - Get a SCIM group
* [UpdateGroup](docs/sdks/scim/README.md#updategroup) - Update a SCIM group
* [DeleteGroup](docs/sdks/scim/README.md#deletegroup) - Delete a SCIM group
* [ListUsers](docs/sdks/scim/README.md#listusers) - List users via SCIM
* [CreateUser](docs/sdks/scim/README.md#createuser) - Create a user via SCIM
* [GetUser](docs/sdks/scim/README.md#getuser) - Get a SCIM user
* [ReplaceUser](docs/sdks/scim/README.md#replaceuser) - Replace a SCIM user
* [DeleteUser](docs/sdks/scim/README.md#deleteuser) - Delete a SCIM user
* [UpdateUser](docs/sdks/scim/README.md#updateuser) - Update a SCIM user

### [Services](docs/sdks/services/README.md)

* [CreateDependency](docs/sdks/services/README.md#createdependency) - Create a dependency relationship between services
* [GetDependency](docs/sdks/services/README.md#getdependency) - Get a service dependency
* [DeleteDependency](docs/sdks/services/README.md#deletedependency) - Delete a service dependency
* [Update](docs/sdks/services/README.md#update) - Update a service dependency
* [List](docs/sdks/services/README.md#list) - List services
* [Create](docs/sdks/services/README.md#create) - Create a service
* [CreateLinks](docs/sdks/services/README.md#createlinks) - Create multiple services and link them to external services
* [Get](docs/sdks/services/README.md#get) - Get a service
* [Delete](docs/sdks/services/README.md#delete) - Delete a service
* [Patch](docs/sdks/services/README.md#patch) - Update a service
* [GetAvailableDownstreamDependencies](docs/sdks/services/README.md#getavailabledownstreamdependencies) - List available downstream service dependencies
* [ListAvailableUpstreamDependencies](docs/sdks/services/README.md#listavailableupstreamdependencies) - List available upstream service dependencies
* [CreateChecklistResponse](docs/sdks/services/README.md#createchecklistresponse) - Create a checklist response for a service
* [ListDependencies](docs/sdks/services/README.md#listdependencies) - List dependencies for a service
* [DeleteLink](docs/sdks/services/README.md#deletelink) - Delete a service link
* [ListForUser](docs/sdks/services/README.md#listforuser) - List services for a user's teams

#### [Services.Catalogs](docs/sdks/catalogs/README.md)

* [Ingest](docs/sdks/catalogs/README.md#ingest) - Ingest service catalog data
* [Refresh](docs/sdks/catalogs/README.md#refresh) - Refresh a service catalog

### [Signals](docs/sdks/signals/README.md)

* [ListGroupedMetrics](docs/sdks/signals/README.md#listgroupedmetrics) - List grouped signal alert metrics
* [GetMttxAnalytics](docs/sdks/signals/README.md#getmttxanalytics) - Get MTTX analytics for signals
* [GetAnalyticsTimeseries](docs/sdks/signals/README.md#getanalyticstimeseries) - List timeseries metrics for signal alerts
* [Debug](docs/sdks/signals/README.md#debug) - Debug a signal
* [ListEmailTargets](docs/sdks/signals/README.md#listemailtargets) - List email targets for signals
* [CreateEmailTarget](docs/sdks/signals/README.md#createemailtarget) - Create an email target for signals
* [GetEmailTarget](docs/sdks/signals/README.md#getemailtarget) - Get a signal email target
* [DeleteEmailTarget](docs/sdks/signals/README.md#deleteemailtarget) - Delete a signal email target
* [UpdateEmailTarget](docs/sdks/signals/README.md#updateemailtarget) - Update a signal email target
* [ListEventSources](docs/sdks/signals/README.md#listeventsources) - List event sources for signals
* [GetIngestURL](docs/sdks/signals/README.md#getingesturl) - Get signal ingestion URL
* [ListTransposers](docs/sdks/signals/README.md#listtransposers) - List signal transposers
* [ListWebhookTargets](docs/sdks/signals/README.md#listwebhooktargets) - List webhook targets for signals
* [CreateWebhookTarget](docs/sdks/signals/README.md#createwebhooktarget) - Create a webhook target for signals
* [GetWebhookTarget](docs/sdks/signals/README.md#getwebhooktarget) - Get a webhook target
* [DeleteWebhookTarget](docs/sdks/signals/README.md#deletewebhooktarget) - Delete a webhook target
* [UpdateWebhookTarget](docs/sdks/signals/README.md#updatewebhooktarget) - Update a webhook target
* [ListOnCall](docs/sdks/signals/README.md#listoncall) - List on-call schedules
* [ListEscalationPolicies](docs/sdks/signals/README.md#listescalationpolicies) - List escalation policies for a team
* [DeleteEscalationPolicy](docs/sdks/signals/README.md#deleteescalationpolicy) - Delete an escalation policy for a team
* [GetOnCallSchedule](docs/sdks/signals/README.md#getoncallschedule) - Get an on-call schedule for a team
* [UpdateOnCallSchedule](docs/sdks/signals/README.md#updateoncallschedule) - Update an on-call schedule for a team
* [ListRules](docs/sdks/signals/README.md#listrules) - List signal rules for a team
* [CreateRule](docs/sdks/signals/README.md#createrule) - Create a signal rule for a team
* [GetRule](docs/sdks/signals/README.md#getrule) - Get a signal rule
* [DeleteRule](docs/sdks/signals/README.md#deleterule) - Delete a signal rule
* [UpdateRule](docs/sdks/signals/README.md#updaterule) - Update a signal rule

#### [Signals.Teams](docs/sdks/firehydrantsignalsteams/README.md)

* [GetEscalationPolicy](docs/sdks/firehydrantsignalsteams/README.md#getescalationpolicy) - Get an escalation policy for a team

### [Slack](docs/sdks/slack/README.md)

* [CreateEmojiAction](docs/sdks/slack/README.md#createemojiaction) - Create a Slack emoji action
* [UpdateEmojiAction](docs/sdks/slack/README.md#updateemojiaction) - Update a Slack emoji action
* [ListUsergroups](docs/sdks/slack/README.md#listusergroups) - List Slack usergroups

### [Statuspage](docs/sdks/statuspage/README.md)

* [UpdateConnection](docs/sdks/statuspage/README.md#updateconnection) - Update a Statuspage connection

### [StatusPages](docs/sdks/statuspages/README.md)

* [DeleteIncident](docs/sdks/statuspages/README.md#deleteincident) - Remove a status page from an incident
* [CreateSubscription](docs/sdks/statuspages/README.md#createsubscription) - Create a status page subscription
* [DeleteSubscription](docs/sdks/statuspages/README.md#deletesubscription) - Unsubscribe from status page notifications
* [List](docs/sdks/statuspages/README.md#list) - List status pages
* [Create](docs/sdks/statuspages/README.md#create) - Create a status page
* [Get](docs/sdks/statuspages/README.md#get) - Get a status page
* [Update](docs/sdks/statuspages/README.md#update) - Update a status page
* [Delete](docs/sdks/statuspages/README.md#delete) - Delete a status page
* [CreateComponentGroup](docs/sdks/statuspages/README.md#createcomponentgroup) - Create a component group for a status page
* [DeleteComponentGroup](docs/sdks/statuspages/README.md#deletecomponentgroup) - Delete a status page component group
* [UpdateComponentGroup](docs/sdks/statuspages/README.md#updatecomponentgroup) - Update a status page component group
* [UpdateImage](docs/sdks/statuspages/README.md#updateimage) - Upload an image for a status page
* [DeleteImage](docs/sdks/statuspages/README.md#deleteimage) - Delete an image from a status page
* [CreateLink](docs/sdks/statuspages/README.md#createlink) - Create a link for a status page
* [DeleteLink](docs/sdks/statuspages/README.md#deletelink) - Delete a status page link
* [UpdateLink](docs/sdks/statuspages/README.md#updatelink) - Update a status page link
* [ListSubscribers](docs/sdks/statuspages/README.md#listsubscribers) - List status page subscribers
* [CreateSubscribers](docs/sdks/statuspages/README.md#createsubscribers) - Add subscribers to a status page
* [DeleteSubscribers](docs/sdks/statuspages/README.md#deletesubscribers) - Remove subscribers from a status page

### [StatusUpdateTemplates](docs/sdks/statusupdatetemplates/README.md)

* [Create](docs/sdks/statusupdatetemplates/README.md#create) - Create a status update template

### [System](docs/sdks/system/README.md)

* [Ping](docs/sdks/system/README.md#ping) - Check API connectivity

### [Tasks](docs/sdks/tasks/README.md)

* [ListChecklistTemplates](docs/sdks/tasks/README.md#listchecklisttemplates) - List checklist templates
* [CreateChecklistTemplate](docs/sdks/tasks/README.md#createchecklisttemplate) - Create a checklist template
* [DeleteChecklistTemplate](docs/sdks/tasks/README.md#deletechecklisttemplate) - Archive a checklist template
* [UpdateChecklistTemplate](docs/sdks/tasks/README.md#updatechecklisttemplate) - Update a checklist template
* [ListForIncident](docs/sdks/tasks/README.md#listforincident) - List tasks for an incident
* [Create](docs/sdks/tasks/README.md#create) - Create a task for an incident
* [GetForIncident](docs/sdks/tasks/README.md#getforincident) - Get a task for an incident
* [Delete](docs/sdks/tasks/README.md#delete) - Delete a task from an incident
* [UpdateTask](docs/sdks/tasks/README.md#updatetask) - Update a task for an incident
* [ConvertToFollowup](docs/sdks/tasks/README.md#converttofollowup) - Convert a task to a follow-up
* [ListTasks](docs/sdks/tasks/README.md#listtasks) - List task lists
* [CreateList](docs/sdks/tasks/README.md#createlist) - Create a task list
* [Get](docs/sdks/tasks/README.md#get) - Get a task list
* [Update](docs/sdks/tasks/README.md#update) - Update a task list

#### [Tasks.List](docs/sdks/list/README.md)

* [Delete](docs/sdks/list/README.md#delete) - Delete a task list

### [Teams](docs/sdks/teams/README.md)

* [ListSchedules](docs/sdks/teams/README.md#listschedules) - List schedules
* [List](docs/sdks/teams/README.md#list) - List teams
* [Create](docs/sdks/teams/README.md#create) - Create a team
* [Get](docs/sdks/teams/README.md#get) - Get a team
* [Archive](docs/sdks/teams/README.md#archive) - Archive a team
* [Update](docs/sdks/teams/README.md#update) - Update a team
* [ListOnCallSchedules](docs/sdks/teams/README.md#listoncallschedules) - List on-call schedules for a team
* [CreateOnCallSchedule](docs/sdks/teams/README.md#createoncallschedule) - Create an on-call schedule for a team
* [GetScheduleShift](docs/sdks/teams/README.md#getscheduleshift) - Get an on-call shift for a team schedule
* [DeleteScheduleShift](docs/sdks/teams/README.md#deletescheduleshift) - Delete an on-call shift from a team schedule
* [UpdateScheduleShift](docs/sdks/teams/README.md#updatescheduleshift) - Update an on-call shift in a team schedule
* [CreateEscalationPolicy](docs/sdks/teams/README.md#createescalationpolicy) - Create an escalation policy for a team
* [UpdateEscalationPolicy](docs/sdks/teams/README.md#updateescalationpolicy) - Update an escalation policy for a team
* [CreateShift](docs/sdks/teams/README.md#createshift) - Create a shift for an on-call schedule

### [Ticketing](docs/sdks/ticketing/README.md)

* [GetFieldMap](docs/sdks/ticketing/README.md#getfieldmap) - Get a field map for a ticketing project
* [UpdateProjectConfig](docs/sdks/ticketing/README.md#updateprojectconfig) - Update a ticketing project configuration

### [TicketingPriorities](docs/sdks/ticketingpriorities/README.md)

* [Get](docs/sdks/ticketingpriorities/README.md#get) - Get a ticketing priority

### [Tickets](docs/sdks/tickets/README.md)

* [List](docs/sdks/tickets/README.md#list) - List tickets
* [Create](docs/sdks/tickets/README.md#create) - Create a ticket
* [Delete](docs/sdks/tickets/README.md#delete) - Delete a ticket
* [Update](docs/sdks/tickets/README.md#update) - Update a ticket

### [Users](docs/sdks/users/README.md)

* [GetCurrent](docs/sdks/users/README.md#getcurrent) - Get the currently authenticated user
* [List](docs/sdks/users/README.md#list) - List users
* [Get](docs/sdks/users/README.md#get) - Get a user

### [Webhooks](docs/sdks/webhooks/README.md)

* [List](docs/sdks/webhooks/README.md#list) - List webhooks
* [Create](docs/sdks/webhooks/README.md#create) - Create a webhook
* [Get](docs/sdks/webhooks/README.md#get) - Get a webhook
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete a webhook
* [Update](docs/sdks/webhooks/README.md#update) - Update a webhook
* [ListDeliveries](docs/sdks/webhooks/README.md#listdeliveries) - List webhook deliveries

### [Zendesk](docs/sdks/zendesk/README.md)

* [SearchTickets](docs/sdks/zendesk/README.md#searchtickets) - Search for Zendesk tickets

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.AccountSettings.GetAiPreferences(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AIEntitiesPreferencesEntity != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.AccountSettings.GetAiPreferences(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AIEntitiesPreferencesEntity != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetAiPreferences` function may return the following errors:

| Error Type                    | Status Code                       | Content Type     |
| ----------------------------- | --------------------------------- | ---------------- |
| sdkerrors.BadRequest          | 400, 413, 414, 415, 422, 431, 510 | application/json |
| sdkerrors.Unauthorized        | 401, 403, 407, 511                | application/json |
| sdkerrors.NotFound            | 404, 501, 505                     | application/json |
| sdkerrors.Timeout             | 408, 504                          | application/json |
| sdkerrors.RateLimited         | 429                               | application/json |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508      | application/json |
| sdkerrors.SDKError            | 4XX, 5XX                          | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"firehydrant"
	"firehydrant/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.AccountSettings.GetAiPreferences(ctx)
	if err != nil {

		var e *sdkerrors.BadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Unauthorized
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFound
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Timeout
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.RateLimited
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.InternalServerError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	ctx := context.Background()

	s := firehydrant.New(
		firehydrant.WithServerURL("https://api.firehydrant.io/"),
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.AccountSettings.GetAiPreferences(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AIEntitiesPreferencesEntity != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go)
