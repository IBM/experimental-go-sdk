/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-b03c6166-20201125-093339
 */
 

// Package ibmclouddatabasesv5 : Operations and models for the IbmCloudDatabasesV5 service
package ibmclouddatabasesv5

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	common "github.com/IBM/experimental-go-sdk/common"
	"net/http"
	"reflect"
	"time"
)

// IbmCloudDatabasesV5 : The IBM Cloud Databases API enables interaction between applications and Cloud Databases
// database deployments.
//
// Access to the API requires an IAM Bearer Token or an IAM API Key to be presented through bearer authentication.
//
// Deployment IDs are CRNs on the IBM Cloud Databases v5 API platform. No lookup or translation the Compose style UUIDs
// is needed. The Deployment ID is a traditional UUID on the Compose v5 API platform.
//
// When you use CRNs, remember to URL escape the CRN value as they can include the forward-slash (/) character.
//
// Version: 5.0.0
type IbmCloudDatabasesV5 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_cloud_databases"

// IbmCloudDatabasesV5Options : Service options
type IbmCloudDatabasesV5Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmCloudDatabasesV5UsingExternalConfig : constructs an instance of IbmCloudDatabasesV5 with passed in options and external configuration.
func NewIbmCloudDatabasesV5UsingExternalConfig(options *IbmCloudDatabasesV5Options) (ibmCloudDatabases *IbmCloudDatabasesV5, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmCloudDatabases, err = NewIbmCloudDatabasesV5(options)
	if err != nil {
		return
	}

	err = ibmCloudDatabases.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmCloudDatabases.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmCloudDatabasesV5 : constructs an instance of IbmCloudDatabasesV5 with passed in options.
func NewIbmCloudDatabasesV5(options *IbmCloudDatabasesV5Options) (service *IbmCloudDatabasesV5, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &IbmCloudDatabasesV5{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetServiceURL(url string) error {
	return ibmCloudDatabases.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetServiceURL() string {
	return ibmCloudDatabases.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetDefaultHeaders(headers http.Header) {
	ibmCloudDatabases.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetEnableGzipCompression(enableGzip bool) {
	ibmCloudDatabases.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetEnableGzipCompression() bool {
	return ibmCloudDatabases.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (ibmCloudDatabases *IbmCloudDatabasesV5) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	ibmCloudDatabases.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (ibmCloudDatabases *IbmCloudDatabasesV5) DisableRetries() {
	ibmCloudDatabases.Service.DisableRetries()
}

// GetDeployables : Get all deployable databases
// Returns a list of all the types and associated major versions of database deployments that can be provisioned.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeployables(getDeployablesOptions *GetDeployablesOptions) (result *GetDeployablesResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDeployablesWithContext(context.Background(), getDeployablesOptions)
}

// GetDeployablesWithContext is an alternate form of the GetDeployables method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeployablesWithContext(ctx context.Context, getDeployablesOptions *GetDeployablesOptions) (result *GetDeployablesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getDeployablesOptions, "getDeployablesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployables`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeployablesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDeployables")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDeployablesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRegions : Get all deployable regions
// Returns a list of all the regions that deployments can be provisioned into from the current region. Used to determine
// region availability for read-only replicas.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetRegions(getRegionsOptions *GetRegionsOptions) (result *GetRegionsResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetRegionsWithContext(context.Background(), getRegionsOptions)
}

// GetRegionsWithContext is an alternate form of the GetRegions method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetRegionsWithContext(ctx context.Context, getRegionsOptions *GetRegionsOptions) (result *GetRegionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getRegionsOptions, "getRegionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/regions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRegionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetRegions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetRegionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDeploymentInfo : Get deployment information
// Gets the full data that is associated with a deployment. This data includes the ID, name, database type, and version.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions) (result *GetDeploymentInfoResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDeploymentInfoWithContext(context.Background(), getDeploymentInfoOptions)
}

// GetDeploymentInfoWithContext is an alternate form of the GetDeploymentInfo method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentInfoWithContext(ctx context.Context, getDeploymentInfoOptions *GetDeploymentInfoOptions) (result *GetDeploymentInfoResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDeploymentInfoOptions, "getDeploymentInfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDeploymentInfoOptions, "getDeploymentInfoOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDeploymentInfoOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeploymentInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDeploymentInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDeploymentInfoResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateDatabaseUser : Creates a user based on user type
// Creates a user in the database that can access the database through a connection.
func (ibmCloudDatabases *IbmCloudDatabasesV5) CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions) (result *CreateDatabaseUserResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.CreateDatabaseUserWithContext(context.Background(), createDatabaseUserOptions)
}

// CreateDatabaseUserWithContext is an alternate form of the CreateDatabaseUser method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) CreateDatabaseUserWithContext(ctx context.Context, createDatabaseUserOptions *CreateDatabaseUserOptions) (result *CreateDatabaseUserResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDatabaseUserOptions, "createDatabaseUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDatabaseUserOptions, "createDatabaseUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *createDatabaseUserOptions.ID,
		"user_type": *createDatabaseUserOptions.UserType,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDatabaseUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "CreateDatabaseUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDatabaseUserOptions.User != nil {
		body["user"] = createDatabaseUserOptions.User
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateDatabaseUserResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ChangeUserPassword : Set specified user's password
// Sets the password of a specified user.
func (ibmCloudDatabases *IbmCloudDatabasesV5) ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions) (result *ChangeUserPasswordResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.ChangeUserPasswordWithContext(context.Background(), changeUserPasswordOptions)
}

// ChangeUserPasswordWithContext is an alternate form of the ChangeUserPassword method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) ChangeUserPasswordWithContext(ctx context.Context, changeUserPasswordOptions *ChangeUserPasswordOptions) (result *ChangeUserPasswordResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(changeUserPasswordOptions, "changeUserPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(changeUserPasswordOptions, "changeUserPasswordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *changeUserPasswordOptions.ID,
		"user_type": *changeUserPasswordOptions.UserType,
		"username": *changeUserPasswordOptions.Username,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{username}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range changeUserPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "ChangeUserPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if changeUserPasswordOptions.User != nil {
		body["user"] = changeUserPasswordOptions.User
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalChangeUserPasswordResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteDatabaseUser : Deletes a user based on user type
// Removes a user from the deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions) (result *DeleteDatabaseUserResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.DeleteDatabaseUserWithContext(context.Background(), deleteDatabaseUserOptions)
}

// DeleteDatabaseUserWithContext is an alternate form of the DeleteDatabaseUser method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) DeleteDatabaseUserWithContext(ctx context.Context, deleteDatabaseUserOptions *DeleteDatabaseUserOptions) (result *DeleteDatabaseUserResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDatabaseUserOptions, "deleteDatabaseUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDatabaseUserOptions, "deleteDatabaseUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteDatabaseUserOptions.ID,
		"user_type": *deleteDatabaseUserOptions.UserType,
		"username": *deleteDatabaseUserOptions.Username,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{username}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDatabaseUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "DeleteDatabaseUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteDatabaseUserResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetUser : Discover user name and password information for a deployment for a user with an endpoint type
// Only for Redis v5 and prior: Discover connection information for a deployment for a user with an endpoint type.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetUser(getUserOptions *GetUserOptions) (result *Task, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetUserWithContext(context.Background(), getUserOptions)
}

// GetUserWithContext is an alternate form of the GetUser method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetUserWithContext(ctx context.Context, getUserOptions *GetUserOptions) (result *Task, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserOptions, "getUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserOptions, "getUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getUserOptions.ID,
		"user_id": *getUserOptions.UserID,
		"endpoint_type": *getUserOptions.EndpointType,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_id}/`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTask)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetDatabaseConfiguration : Change your database configuration
// Change your database configuration. Available for PostgreSQL, EnterpriseDB, and Redis ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetDatabaseConfiguration(setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions) (result *SetDatabaseConfigurationResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.SetDatabaseConfigurationWithContext(context.Background(), setDatabaseConfigurationOptions)
}

// SetDatabaseConfigurationWithContext is an alternate form of the SetDatabaseConfiguration method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetDatabaseConfigurationWithContext(ctx context.Context, setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions) (result *SetDatabaseConfigurationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setDatabaseConfigurationOptions, "setDatabaseConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setDatabaseConfigurationOptions, "setDatabaseConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setDatabaseConfigurationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setDatabaseConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "SetDatabaseConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setDatabaseConfigurationOptions.Configuration != nil {
		body["configuration"] = setDatabaseConfigurationOptions.Configuration
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetDatabaseConfigurationResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDatabaseConfigurationSchema : Get the schema of the database configuration
// Get the schema of the database configuration. Available for PostgreSQL and Redis ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions) (result *GetDatabaseConfigurationSchemaResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDatabaseConfigurationSchemaWithContext(context.Background(), getDatabaseConfigurationSchemaOptions)
}

// GetDatabaseConfigurationSchemaWithContext is an alternate form of the GetDatabaseConfigurationSchema method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDatabaseConfigurationSchemaWithContext(ctx context.Context, getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions) (result *GetDatabaseConfigurationSchemaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDatabaseConfigurationSchemaOptions, "getDatabaseConfigurationSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDatabaseConfigurationSchemaOptions, "getDatabaseConfigurationSchemaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDatabaseConfigurationSchemaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/configuration/schema`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDatabaseConfigurationSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDatabaseConfigurationSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetDatabaseConfigurationSchemaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRemotes : Get read-only replica information
// Get the read-only replicas associated with a deployment. Available for PostgreSQL and EnterpriseDB ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetRemotes(getRemotesOptions *GetRemotesOptions) (result *GetRemotesResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetRemotesWithContext(context.Background(), getRemotesOptions)
}

// GetRemotesWithContext is an alternate form of the GetRemotes method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetRemotesWithContext(ctx context.Context, getRemotesOptions *GetRemotesOptions) (result *GetRemotesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRemotesOptions, "getRemotesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRemotesOptions, "getRemotesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getRemotesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/remotes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRemotesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetRemotes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetRemotesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetRemotes : Modify read-only replication on a deployment
// Promote a read-only remote replica to leader by calling with leader set to an empty string. Available for PostgreSQL
// and EnterpriseDB ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetRemotes(setRemotesOptions *SetRemotesOptions) (result *SetRemotesResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.SetRemotesWithContext(context.Background(), setRemotesOptions)
}

// SetRemotesWithContext is an alternate form of the SetRemotes method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetRemotesWithContext(ctx context.Context, setRemotesOptions *SetRemotesOptions) (result *SetRemotesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setRemotesOptions, "setRemotesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setRemotesOptions, "setRemotesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setRemotesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/remotes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setRemotesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "SetRemotes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setRemotesOptions.Remotes != nil {
		body["remotes"] = setRemotesOptions.Remotes
	}
	if setRemotesOptions.SkipInitialBackup != nil {
		body["skip_initial_backup"] = setRemotesOptions.SkipInitialBackup
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetRemotesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRemotesSchema : Resync read-only replica
// Reinitialize a read-only replica. Available for PostgreSQL and EnterpriseDB ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetRemotesSchema(getRemotesSchemaOptions *GetRemotesSchemaOptions) (result *GetRemotesSchemaResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetRemotesSchemaWithContext(context.Background(), getRemotesSchemaOptions)
}

// GetRemotesSchemaWithContext is an alternate form of the GetRemotesSchema method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetRemotesSchemaWithContext(ctx context.Context, getRemotesSchemaOptions *GetRemotesSchemaOptions) (result *GetRemotesSchemaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRemotesSchemaOptions, "getRemotesSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRemotesSchemaOptions, "getRemotesSchemaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getRemotesSchemaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/remotes/resync`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRemotesSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetRemotesSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetRemotesSchemaResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetPromotion : Promote read-only replica to a full deployment
// Promote a read-only replica or upgrade and promote a read-only replica. Available for PostgreSQL and EnterpriseDB
// ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetPromotion(setPromotionOptions *SetPromotionOptions) (result *SetPromotionResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.SetPromotionWithContext(context.Background(), setPromotionOptions)
}

// SetPromotionWithContext is an alternate form of the SetPromotion method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetPromotionWithContext(ctx context.Context, setPromotionOptions *SetPromotionOptions) (result *SetPromotionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setPromotionOptions, "setPromotionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setPromotionOptions, "setPromotionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setPromotionOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/remotes/promotion`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setPromotionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "SetPromotion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setPromotionOptions.Promotion != nil {
		body["Promotion"] = setPromotionOptions.Promotion
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetPromotionResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDeploymentTasks : Get currently running tasks on a deployment
// Obtain a list of tasks currently running or recently run on a deployment. Tasks are ephemeral. Records of successful
// tasks are shown for 24-48 hours, and unsuccessful tasks are shown for 7-8 days.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentTasks(getDeploymentTasksOptions *GetDeploymentTasksOptions) (result *Tasks, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDeploymentTasksWithContext(context.Background(), getDeploymentTasksOptions)
}

// GetDeploymentTasksWithContext is an alternate form of the GetDeploymentTasks method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentTasksWithContext(ctx context.Context, getDeploymentTasksOptions *GetDeploymentTasksOptions) (result *Tasks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDeploymentTasksOptions, "getDeploymentTasksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDeploymentTasksOptions, "getDeploymentTasksOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDeploymentTasksOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/tasks`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeploymentTasksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDeploymentTasks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTasks)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetTasks : Get information about a task
// Get information about a task and its status. Tasks themselves are persistent so old tasks can be consulted as well as
// running tasks.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetTasks(getTasksOptions *GetTasksOptions) (result *GetTasksResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetTasksWithContext(context.Background(), getTasksOptions)
}

// GetTasksWithContext is an alternate form of the GetTasks method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetTasksWithContext(ctx context.Context, getTasksOptions *GetTasksOptions) (result *GetTasksResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTasksOptions, "getTasksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTasksOptions, "getTasksOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getTasksOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/tasks/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTasksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetTasks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTasksResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetBackupInfo : Get information about a backup
// Get information about a backup, such as creation date.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions) (result *GetBackupInfoResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetBackupInfoWithContext(context.Background(), getBackupInfoOptions)
}

// GetBackupInfoWithContext is an alternate form of the GetBackupInfo method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetBackupInfoWithContext(ctx context.Context, getBackupInfoOptions *GetBackupInfoOptions) (result *GetBackupInfoResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBackupInfoOptions, "getBackupInfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBackupInfoOptions, "getBackupInfoOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"backup_id": *getBackupInfoOptions.BackupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/backups/{backup_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBackupInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetBackupInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBackupInfoResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDeploymentBackups : Get currently available backups from a deployment
// Get details of all currently available backups from a deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentBackups(getDeploymentBackupsOptions *GetDeploymentBackupsOptions) (result *Backups, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDeploymentBackupsWithContext(context.Background(), getDeploymentBackupsOptions)
}

// GetDeploymentBackupsWithContext is an alternate form of the GetDeploymentBackups method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentBackupsWithContext(ctx context.Context, getDeploymentBackupsOptions *GetDeploymentBackupsOptions) (result *Backups, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDeploymentBackupsOptions, "getDeploymentBackupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDeploymentBackupsOptions, "getDeploymentBackupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDeploymentBackupsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/backups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeploymentBackupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDeploymentBackups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBackups)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// StartOndemandBackup : Initiate an on-demand backup
// Signal the platform to create an on-demand backup for the specified deployment. The returned task can be polled to
// track progress of the backup as it takes place.
func (ibmCloudDatabases *IbmCloudDatabasesV5) StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions) (result *StartOndemandBackupResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.StartOndemandBackupWithContext(context.Background(), startOndemandBackupOptions)
}

// StartOndemandBackupWithContext is an alternate form of the StartOndemandBackup method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) StartOndemandBackupWithContext(ctx context.Context, startOndemandBackupOptions *StartOndemandBackupOptions) (result *StartOndemandBackupResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(startOndemandBackupOptions, "startOndemandBackupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(startOndemandBackupOptions, "startOndemandBackupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *startOndemandBackupOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/backups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range startOndemandBackupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "StartOndemandBackup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStartOndemandBackupResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetPITRdata : Get earliest point-in-time-recovery timestamp
// Returns the earliest available time for point-in-time-recovery in ISO8601 UTC format. PostgreSQL and EnterpriseDB
// only.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetPITRdata(getPITRdataOptions *GetPITRdataOptions) (result *PointInTimeRecoveryData, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetPITRdataWithContext(context.Background(), getPITRdataOptions)
}

// GetPITRdataWithContext is an alternate form of the GetPITRdata method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetPITRdataWithContext(ctx context.Context, getPITRdataOptions *GetPITRdataOptions) (result *PointInTimeRecoveryData, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPITRdataOptions, "getPITRdataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPITRdataOptions, "getPITRdataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getPITRdataOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/point_in_time_recovery_data`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPITRdataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetPITRdata")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPointInTimeRecoveryData)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetConnection : Discover connection information for a deployment for a user with an endpoint type
// Discover connection information for a deployment for a user with an endpoint type.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetConnection(getConnectionOptions *GetConnectionOptions) (result *Connection, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetConnectionWithContext(context.Background(), getConnectionOptions)
}

// GetConnectionWithContext is an alternate form of the GetConnection method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetConnectionWithContext(ctx context.Context, getConnectionOptions *GetConnectionOptions) (result *Connection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConnectionOptions, "getConnectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConnectionOptions, "getConnectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getConnectionOptions.ID,
		"user_type": *getConnectionOptions.UserType,
		"user_id": *getConnectionOptions.UserID,
		"endpoint_type": *getConnectionOptions.EndpointType,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{user_id}/connections/{endpoint_type}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConnectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetConnection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getConnectionOptions.CertificateRoot != nil {
		builder.AddQuery("certificate_root", fmt.Sprint(*getConnectionOptions.CertificateRoot))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConnection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CompleteConnection : Discover connection information for a deployment for a user with substitutions and an endpoint type
// Discover connection information for a deployment for a user. Behaves the same as the GET method but substitutes the
// provided password parameter into the returned connection information.
func (ibmCloudDatabases *IbmCloudDatabasesV5) CompleteConnection(completeConnectionOptions *CompleteConnectionOptions) (result *Connection, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.CompleteConnectionWithContext(context.Background(), completeConnectionOptions)
}

// CompleteConnectionWithContext is an alternate form of the CompleteConnection method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) CompleteConnectionWithContext(ctx context.Context, completeConnectionOptions *CompleteConnectionOptions) (result *Connection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeConnectionOptions, "completeConnectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(completeConnectionOptions, "completeConnectionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *completeConnectionOptions.ID,
		"user_type": *completeConnectionOptions.UserType,
		"user_id": *completeConnectionOptions.UserID,
		"endpoint_type": *completeConnectionOptions.EndpointType,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{user_id}/connections/{endpoint_type}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range completeConnectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "CompleteConnection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if completeConnectionOptions.Password != nil {
		body["password"] = completeConnectionOptions.Password
	}
	if completeConnectionOptions.CertificateRoot != nil {
		body["certificate_root"] = completeConnectionOptions.CertificateRoot
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConnection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetConnectionDeprecated : Discover connection information for a deployment for a user
// Discover connection information for a deployment for a user.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetConnectionDeprecated(getConnectionDeprecatedOptions *GetConnectionDeprecatedOptions) (result *Connection, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetConnectionDeprecatedWithContext(context.Background(), getConnectionDeprecatedOptions)
}

// GetConnectionDeprecatedWithContext is an alternate form of the GetConnectionDeprecated method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetConnectionDeprecatedWithContext(ctx context.Context, getConnectionDeprecatedOptions *GetConnectionDeprecatedOptions) (result *Connection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConnectionDeprecatedOptions, "getConnectionDeprecatedOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConnectionDeprecatedOptions, "getConnectionDeprecatedOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getConnectionDeprecatedOptions.ID,
		"user_type": *getConnectionDeprecatedOptions.UserType,
		"user_id": *getConnectionDeprecatedOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{user_id}/connections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConnectionDeprecatedOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetConnectionDeprecated")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getConnectionDeprecatedOptions.CertificateRoot != nil {
		builder.AddQuery("certificate_root", fmt.Sprint(*getConnectionDeprecatedOptions.CertificateRoot))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConnection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CompleteConnectionDeprecated : Discover connection information for a deployment for a user with substitutions
// Discover connection information for a deployment for a user. Behaves the same as the GET method but substitutes the
// given password parameter into the returned connection information.
func (ibmCloudDatabases *IbmCloudDatabasesV5) CompleteConnectionDeprecated(completeConnectionDeprecatedOptions *CompleteConnectionDeprecatedOptions) (result *Connection, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.CompleteConnectionDeprecatedWithContext(context.Background(), completeConnectionDeprecatedOptions)
}

// CompleteConnectionDeprecatedWithContext is an alternate form of the CompleteConnectionDeprecated method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) CompleteConnectionDeprecatedWithContext(ctx context.Context, completeConnectionDeprecatedOptions *CompleteConnectionDeprecatedOptions) (result *Connection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeConnectionDeprecatedOptions, "completeConnectionDeprecatedOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(completeConnectionDeprecatedOptions, "completeConnectionDeprecatedOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *completeConnectionDeprecatedOptions.ID,
		"user_type": *completeConnectionDeprecatedOptions.UserType,
		"user_id": *completeConnectionDeprecatedOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/users/{user_type}/{user_id}/connections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range completeConnectionDeprecatedOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "CompleteConnectionDeprecated")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if completeConnectionDeprecatedOptions.Password != nil {
		body["password"] = completeConnectionDeprecatedOptions.Password
	}
	if completeConnectionDeprecatedOptions.CertificateRoot != nil {
		body["certificate_root"] = completeConnectionDeprecatedOptions.CertificateRoot
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConnection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDeploymentScalingGroups : Get currently available scaling groups from a deployment
// Scaling groups represent the various resources that are allocated to a deployment. This command allows for the
// retrieval of all of the groups for a particular deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentScalingGroups(getDeploymentScalingGroupsOptions *GetDeploymentScalingGroupsOptions) (result *Groups, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDeploymentScalingGroupsWithContext(context.Background(), getDeploymentScalingGroupsOptions)
}

// GetDeploymentScalingGroupsWithContext is an alternate form of the GetDeploymentScalingGroups method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDeploymentScalingGroupsWithContext(ctx context.Context, getDeploymentScalingGroupsOptions *GetDeploymentScalingGroupsOptions) (result *Groups, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDeploymentScalingGroupsOptions, "getDeploymentScalingGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDeploymentScalingGroupsOptions, "getDeploymentScalingGroupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getDeploymentScalingGroupsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/groups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDeploymentScalingGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDeploymentScalingGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroups)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDefaultScalingGroups : Get default scaling groups for a new deployment
// Scaling groups represent the various resources allocated to a deployment. When a new deployment is created, there are
// a set of defaults for each database type. This endpoint returns them for a particular database.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) (result *Groups, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetDefaultScalingGroupsWithContext(context.Background(), getDefaultScalingGroupsOptions)
}

// GetDefaultScalingGroupsWithContext is an alternate form of the GetDefaultScalingGroups method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetDefaultScalingGroupsWithContext(ctx context.Context, getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) (result *Groups, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDefaultScalingGroupsOptions, "getDefaultScalingGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDefaultScalingGroupsOptions, "getDefaultScalingGroupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"type": *getDefaultScalingGroupsOptions.Type,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployables/{type}/groups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDefaultScalingGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetDefaultScalingGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroups)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetDeploymentScalingGroup : Set scaling values on a specified group
// Set scaling value on a specified group. Can only be performed on is_adjustable=true groups. Values set are for the
// group as a whole and resources are distributed amongst the group. Values must be greater than or equal to the minimum
// size and must be a multiple of the step size.
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) (result *SetDeploymentScalingGroupResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.SetDeploymentScalingGroupWithContext(context.Background(), setDeploymentScalingGroupOptions)
}

// SetDeploymentScalingGroupWithContext is an alternate form of the SetDeploymentScalingGroup method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetDeploymentScalingGroupWithContext(ctx context.Context, setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) (result *SetDeploymentScalingGroupResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setDeploymentScalingGroupOptions, "setDeploymentScalingGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setDeploymentScalingGroupOptions, "setDeploymentScalingGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setDeploymentScalingGroupOptions.ID,
		"group_id": *setDeploymentScalingGroupOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/groups/{group_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setDeploymentScalingGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "SetDeploymentScalingGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(setDeploymentScalingGroupOptions.SetDeploymentScalingGroupRequest)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetDeploymentScalingGroupResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAutoscalingConditions : Get the autoscaling configuration from a deployment
// The Autoscaling configuration represents the various conditions that control autoscaling for a deployment. This
// command allows for the retrieval of all autoscaling conditions for a particular deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) (result *GetAutoscalingConditionsResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetAutoscalingConditionsWithContext(context.Background(), getAutoscalingConditionsOptions)
}

// GetAutoscalingConditionsWithContext is an alternate form of the GetAutoscalingConditions method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetAutoscalingConditionsWithContext(ctx context.Context, getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) (result *GetAutoscalingConditionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAutoscalingConditionsOptions, "getAutoscalingConditionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAutoscalingConditionsOptions, "getAutoscalingConditionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getAutoscalingConditionsOptions.ID,
		"group_id": *getAutoscalingConditionsOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/groups/{group_id}/autoscaling`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAutoscalingConditionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetAutoscalingConditions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetAutoscalingConditionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetAutoscalingConditions : Set the autoscaling configuration from a deployment
// Enable, disable, or set the conditions for autoscaling on your deployment. Memory, disk, and CPU (if available) can
// be set separately and are not all required.
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) (result *SetAutoscalingConditionsResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.SetAutoscalingConditionsWithContext(context.Background(), setAutoscalingConditionsOptions)
}

// SetAutoscalingConditionsWithContext is an alternate form of the SetAutoscalingConditions method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) SetAutoscalingConditionsWithContext(ctx context.Context, setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) (result *SetAutoscalingConditionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setAutoscalingConditionsOptions, "setAutoscalingConditionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setAutoscalingConditionsOptions, "setAutoscalingConditionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *setAutoscalingConditionsOptions.ID,
		"group_id": *setAutoscalingConditionsOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/groups/{group_id}/autoscaling`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setAutoscalingConditionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "SetAutoscalingConditions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setAutoscalingConditionsOptions.Autoscaling != nil {
		body["autoscaling"] = setAutoscalingConditionsOptions.Autoscaling
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSetAutoscalingConditionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// KillConnections : Kill connections to a PostgreSQL or EnterpriseDB deployment
// Closes all the connections on a deployment. Available for PostgreSQL and EnterpriseDB ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) KillConnections(killConnectionsOptions *KillConnectionsOptions) (result *KillConnectionsResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.KillConnectionsWithContext(context.Background(), killConnectionsOptions)
}

// KillConnectionsWithContext is an alternate form of the KillConnections method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) KillConnectionsWithContext(ctx context.Context, killConnectionsOptions *KillConnectionsOptions) (result *KillConnectionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(killConnectionsOptions, "killConnectionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(killConnectionsOptions, "killConnectionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *killConnectionsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/management/database_connections`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range killConnectionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "KillConnections")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKillConnectionsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// FileSync : Sync files uploaded to Elasticsearch deployment
// Starts a task that writes files to disk. Available for Elasticsearch ONLY.
func (ibmCloudDatabases *IbmCloudDatabasesV5) FileSync(fileSyncOptions *FileSyncOptions) (result *FileSyncResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.FileSyncWithContext(context.Background(), fileSyncOptions)
}

// FileSyncWithContext is an alternate form of the FileSync method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) FileSyncWithContext(ctx context.Context, fileSyncOptions *FileSyncOptions) (result *FileSyncResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(fileSyncOptions, "fileSyncOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(fileSyncOptions, "fileSyncOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *fileSyncOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/elasticsearch/file_syncs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range fileSyncOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "FileSync")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFileSyncResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateLogicalReplicationSlot : Create a new logical replication slot
// Creates a new logical replication slot on the specified database. For use with PostgreSQL, EnterpriseDB, and wal2json
// only.
func (ibmCloudDatabases *IbmCloudDatabasesV5) CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions) (result *CreateLogicalReplicationSlotResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.CreateLogicalReplicationSlotWithContext(context.Background(), createLogicalReplicationSlotOptions)
}

// CreateLogicalReplicationSlotWithContext is an alternate form of the CreateLogicalReplicationSlot method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) CreateLogicalReplicationSlotWithContext(ctx context.Context, createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions) (result *CreateLogicalReplicationSlotResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createLogicalReplicationSlotOptions, "createLogicalReplicationSlotOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createLogicalReplicationSlotOptions, "createLogicalReplicationSlotOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *createLogicalReplicationSlotOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/postgresql/logical_replication_slots`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createLogicalReplicationSlotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "CreateLogicalReplicationSlot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createLogicalReplicationSlotOptions.LogicalReplicationSlot != nil {
		body["logical_replication_slot"] = createLogicalReplicationSlotOptions.LogicalReplicationSlot
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateLogicalReplicationSlotResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteLogicalReplicationSlot : Delete a logical replication slot
// Deletes a logical replication slot from a database. For use with PostgreSQL, EnterpriseDB, and wal2json only.
func (ibmCloudDatabases *IbmCloudDatabasesV5) DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions) (result *DeleteLogicalReplicationSlotResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.DeleteLogicalReplicationSlotWithContext(context.Background(), deleteLogicalReplicationSlotOptions)
}

// DeleteLogicalReplicationSlotWithContext is an alternate form of the DeleteLogicalReplicationSlot method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) DeleteLogicalReplicationSlotWithContext(ctx context.Context, deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions) (result *DeleteLogicalReplicationSlotResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLogicalReplicationSlotOptions, "deleteLogicalReplicationSlotOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLogicalReplicationSlotOptions, "deleteLogicalReplicationSlotOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteLogicalReplicationSlotOptions.ID,
		"name": *deleteLogicalReplicationSlotOptions.Name,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/postgresql/logical_replication_slots/{name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLogicalReplicationSlotOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "DeleteLogicalReplicationSlot")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteLogicalReplicationSlotResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetWhitelist : Retrieve the allowlisted addresses and ranges for a deployment
// Retrieve the allowlisted addresses and ranges for a deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetWhitelist(getWhitelistOptions *GetWhitelistOptions) (result *Whitelist, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.GetWhitelistWithContext(context.Background(), getWhitelistOptions)
}

// GetWhitelistWithContext is an alternate form of the GetWhitelist method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) GetWhitelistWithContext(ctx context.Context, getWhitelistOptions *GetWhitelistOptions) (result *Whitelist, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWhitelistOptions, "getWhitelistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWhitelistOptions, "getWhitelistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getWhitelistOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/whitelists/ip_addresses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWhitelistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "GetWhitelist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWhitelist)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceWhitelist : Replace the allowlist for a deployment
// Replace the allowlist for a deployment. This action overwrites all existing entries, so when you modify the allowlist
// via a GET/update/PUT, provide the GET response's ETag header value in this endpoint's If-Match header to ensure that
// changes that are made by other clients are not accidentally overwritten.
func (ibmCloudDatabases *IbmCloudDatabasesV5) ReplaceWhitelist(replaceWhitelistOptions *ReplaceWhitelistOptions) (result *ReplaceWhitelistResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.ReplaceWhitelistWithContext(context.Background(), replaceWhitelistOptions)
}

// ReplaceWhitelistWithContext is an alternate form of the ReplaceWhitelist method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) ReplaceWhitelistWithContext(ctx context.Context, replaceWhitelistOptions *ReplaceWhitelistOptions) (result *ReplaceWhitelistResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceWhitelistOptions, "replaceWhitelistOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceWhitelistOptions, "replaceWhitelistOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *replaceWhitelistOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/whitelists/ip_addresses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceWhitelistOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "ReplaceWhitelist")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceWhitelistOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*replaceWhitelistOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if replaceWhitelistOptions.IpAddresses != nil {
		body["ip_addresses"] = replaceWhitelistOptions.IpAddresses
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalReplaceWhitelistResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddWhitelistEntry : Add an address or range to the allowlist for a deployment
// Add an address or range to the allowlist for a deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) AddWhitelistEntry(addWhitelistEntryOptions *AddWhitelistEntryOptions) (result *AddWhitelistEntryResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.AddWhitelistEntryWithContext(context.Background(), addWhitelistEntryOptions)
}

// AddWhitelistEntryWithContext is an alternate form of the AddWhitelistEntry method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) AddWhitelistEntryWithContext(ctx context.Context, addWhitelistEntryOptions *AddWhitelistEntryOptions) (result *AddWhitelistEntryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWhitelistEntryOptions, "addWhitelistEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWhitelistEntryOptions, "addWhitelistEntryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *addWhitelistEntryOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/whitelists/ip_addresses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWhitelistEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "AddWhitelistEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWhitelistEntryOptions.IpAddress != nil {
		body["ip_address"] = addWhitelistEntryOptions.IpAddress
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddWhitelistEntryResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteWhitelistEntry : Delete an address or range from the allowlist of a deployment
// Delete an address or range from the allowlist of a deployment.
func (ibmCloudDatabases *IbmCloudDatabasesV5) DeleteWhitelistEntry(deleteWhitelistEntryOptions *DeleteWhitelistEntryOptions) (result *DeleteWhitelistEntryResponse, response *core.DetailedResponse, err error) {
	return ibmCloudDatabases.DeleteWhitelistEntryWithContext(context.Background(), deleteWhitelistEntryOptions)
}

// DeleteWhitelistEntryWithContext is an alternate form of the DeleteWhitelistEntry method which supports a Context parameter
func (ibmCloudDatabases *IbmCloudDatabasesV5) DeleteWhitelistEntryWithContext(ctx context.Context, deleteWhitelistEntryOptions *DeleteWhitelistEntryOptions) (result *DeleteWhitelistEntryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWhitelistEntryOptions, "deleteWhitelistEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWhitelistEntryOptions, "deleteWhitelistEntryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteWhitelistEntryOptions.ID,
		"ipaddress": *deleteWhitelistEntryOptions.Ipaddress,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudDatabases.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudDatabases.Service.Options.URL, `/deployments/{id}/whitelists/ip_addresses/{ipaddress}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteWhitelistEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_databases", "V5", "DeleteWhitelistEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudDatabases.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteWhitelistEntryResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// APasswordSettingUser : APasswordSettingUser struct
type APasswordSettingUser struct {
	Password *string `json:"password,omitempty"`
}


// UnmarshalAPasswordSettingUser unmarshals an instance of APasswordSettingUser from the specified map of raw messages.
func UnmarshalAPasswordSettingUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(APasswordSettingUser)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddWhitelistEntryOptions : The AddWhitelistEntry options.
type AddWhitelistEntryOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	IpAddress *WhitelistEntry `json:"ip_address,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWhitelistEntryOptions : Instantiate AddWhitelistEntryOptions
func (*IbmCloudDatabasesV5) NewAddWhitelistEntryOptions(id string) *AddWhitelistEntryOptions {
	return &AddWhitelistEntryOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *AddWhitelistEntryOptions) SetID(id string) *AddWhitelistEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIpAddress : Allow user to set IpAddress
func (options *AddWhitelistEntryOptions) SetIpAddress(ipAddress *WhitelistEntry) *AddWhitelistEntryOptions {
	options.IpAddress = ipAddress
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddWhitelistEntryOptions) SetHeaders(param map[string]string) *AddWhitelistEntryOptions {
	options.Headers = param
	return options
}

// AddWhitelistEntryResponse : AddWhitelistEntryResponse struct
type AddWhitelistEntryResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalAddWhitelistEntryResponse unmarshals an instance of AddWhitelistEntryResponse from the specified map of raw messages.
func UnmarshalAddWhitelistEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddWhitelistEntryResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingCPUGroup : AutoscalingCPUGroup struct
type AutoscalingCPUGroup struct {
	Cpu *AutoscalingCPUGroup `json:"cpu,omitempty"`
}


// UnmarshalAutoscalingCPUGroup unmarshals an instance of AutoscalingCPUGroup from the specified map of raw messages.
func UnmarshalAutoscalingCPUGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingCPUGroup)
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalAutoscalingCPUGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDisk : AutoscalingDiskGroupDisk struct
type AutoscalingDiskGroupDisk struct {
	Scalers *AutoscalingDiskGroupDiskScalers `json:"scalers,omitempty"`

	Rate *AutoscalingDiskGroupDiskRate `json:"rate,omitempty"`
}


// UnmarshalAutoscalingDiskGroupDisk unmarshals an instance of AutoscalingDiskGroupDisk from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDisk)
	err = core.UnmarshalModel(m, "scalers", &obj.Scalers, UnmarshalAutoscalingDiskGroupDiskScalers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rate", &obj.Rate, UnmarshalAutoscalingDiskGroupDiskRate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskRate : AutoscalingDiskGroupDiskRate struct
type AutoscalingDiskGroupDiskRate struct {
	IncreasePercent *float64 `json:"increase_percent,omitempty"`

	PeriodSeconds *int64 `json:"period_seconds,omitempty"`

	LimitMbPerMember *float64 `json:"limit_mb_per_member,omitempty"`

	Units *string `json:"units,omitempty"`
}


// UnmarshalAutoscalingDiskGroupDiskRate unmarshals an instance of AutoscalingDiskGroupDiskRate from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskRate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskRate)
	err = core.UnmarshalPrimitive(m, "increase_percent", &obj.IncreasePercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "period_seconds", &obj.PeriodSeconds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit_mb_per_member", &obj.LimitMbPerMember)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskScalers : AutoscalingDiskGroupDiskScalers struct
type AutoscalingDiskGroupDiskScalers struct {
	Capacity *AutoscalingDiskGroupDiskScalersCapacity `json:"capacity,omitempty"`

	IoUtilization *AutoscalingDiskGroupDiskScalersIoUtilization `json:"io_utilization,omitempty"`
}


// UnmarshalAutoscalingDiskGroupDiskScalers unmarshals an instance of AutoscalingDiskGroupDiskScalers from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskScalers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskScalers)
	err = core.UnmarshalModel(m, "capacity", &obj.Capacity, UnmarshalAutoscalingDiskGroupDiskScalersCapacity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "io_utilization", &obj.IoUtilization, UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskScalersCapacity : AutoscalingDiskGroupDiskScalersCapacity struct
type AutoscalingDiskGroupDiskScalersCapacity struct {
	Enabled *bool `json:"enabled,omitempty"`

	FreeSpaceRemainingPercent *int64 `json:"free_space_remaining_percent,omitempty"`
}


// UnmarshalAutoscalingDiskGroupDiskScalersCapacity unmarshals an instance of AutoscalingDiskGroupDiskScalersCapacity from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskScalersCapacity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskScalersCapacity)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "free_space_remaining_percent", &obj.FreeSpaceRemainingPercent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingDiskGroupDiskScalersIoUtilization : AutoscalingDiskGroupDiskScalersIoUtilization struct
type AutoscalingDiskGroupDiskScalersIoUtilization struct {
	Enabled *bool `json:"enabled,omitempty"`

	OverPeriod *string `json:"over_period,omitempty"`

	AbovePercent *int64 `json:"above_percent,omitempty"`
}


// UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization unmarshals an instance of AutoscalingDiskGroupDiskScalersIoUtilization from the specified map of raw messages.
func UnmarshalAutoscalingDiskGroupDiskScalersIoUtilization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingDiskGroupDiskScalersIoUtilization)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "over_period", &obj.OverPeriod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "above_percent", &obj.AbovePercent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingGroup : AutoscalingGroup struct
type AutoscalingGroup struct {
	Autoscaling *AutoscalingGroupAutoscaling `json:"autoscaling" validate:"required"`
}


// UnmarshalAutoscalingGroup unmarshals an instance of AutoscalingGroup from the specified map of raw messages.
func UnmarshalAutoscalingGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingGroup)
	err = core.UnmarshalModel(m, "autoscaling", &obj.Autoscaling, UnmarshalAutoscalingGroupAutoscaling)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingGroupAutoscaling : AutoscalingGroupAutoscaling struct
type AutoscalingGroupAutoscaling struct {
	Disk *AutoscalingDiskGroupDisk `json:"disk,omitempty"`

	Memory *AutoscalingMemoryGroupMemory `json:"memory,omitempty"`

	Cpu *AutoscalingCPUGroup `json:"cpu,omitempty"`
}


// UnmarshalAutoscalingGroupAutoscaling unmarshals an instance of AutoscalingGroupAutoscaling from the specified map of raw messages.
func UnmarshalAutoscalingGroupAutoscaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingGroupAutoscaling)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalAutoscalingDiskGroupDisk)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalAutoscalingMemoryGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalAutoscalingCPUGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemory : AutoscalingMemoryGroupMemory struct
type AutoscalingMemoryGroupMemory struct {
	Scalers *AutoscalingMemoryGroupMemoryScalers `json:"scalers,omitempty"`

	Rate *AutoscalingMemoryGroupMemoryRate `json:"rate,omitempty"`
}


// UnmarshalAutoscalingMemoryGroupMemory unmarshals an instance of AutoscalingMemoryGroupMemory from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemory)
	err = core.UnmarshalModel(m, "scalers", &obj.Scalers, UnmarshalAutoscalingMemoryGroupMemoryScalers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rate", &obj.Rate, UnmarshalAutoscalingMemoryGroupMemoryRate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemoryRate : AutoscalingMemoryGroupMemoryRate struct
type AutoscalingMemoryGroupMemoryRate struct {
	IncreasePercent *float64 `json:"increase_percent,omitempty"`

	PeriodSeconds *int64 `json:"period_seconds,omitempty"`

	LimitMbPerMember *float64 `json:"limit_mb_per_member,omitempty"`

	Units *string `json:"units,omitempty"`
}


// UnmarshalAutoscalingMemoryGroupMemoryRate unmarshals an instance of AutoscalingMemoryGroupMemoryRate from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemoryRate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemoryRate)
	err = core.UnmarshalPrimitive(m, "increase_percent", &obj.IncreasePercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "period_seconds", &obj.PeriodSeconds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit_mb_per_member", &obj.LimitMbPerMember)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemoryScalers : AutoscalingMemoryGroupMemoryScalers struct
type AutoscalingMemoryGroupMemoryScalers struct {
	IoUtilization *AutoscalingMemoryGroupMemoryScalersIoUtilization `json:"io_utilization,omitempty"`
}


// UnmarshalAutoscalingMemoryGroupMemoryScalers unmarshals an instance of AutoscalingMemoryGroupMemoryScalers from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemoryScalers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemoryScalers)
	err = core.UnmarshalModel(m, "io_utilization", &obj.IoUtilization, UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingMemoryGroupMemoryScalersIoUtilization : AutoscalingMemoryGroupMemoryScalersIoUtilization struct
type AutoscalingMemoryGroupMemoryScalersIoUtilization struct {
	Enabled *bool `json:"enabled,omitempty"`

	OverPeriod *string `json:"over_period,omitempty"`

	AbovePercent *int64 `json:"above_percent,omitempty"`
}


// UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization unmarshals an instance of AutoscalingMemoryGroupMemoryScalersIoUtilization from the specified map of raw messages.
func UnmarshalAutoscalingMemoryGroupMemoryScalersIoUtilization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingMemoryGroupMemoryScalersIoUtilization)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "over_period", &obj.OverPeriod)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "above_percent", &obj.AbovePercent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroup : AutoscalingSetGroup struct
type AutoscalingSetGroup struct {
	Autoscaling AutoscalingSetGroupAutoscalingIntf `json:"autoscaling" validate:"required"`
}


// NewAutoscalingSetGroup : Instantiate AutoscalingSetGroup (Generic Model Constructor)
func (*IbmCloudDatabasesV5) NewAutoscalingSetGroup(autoscaling AutoscalingSetGroupAutoscalingIntf) (model *AutoscalingSetGroup, err error) {
	model = &AutoscalingSetGroup{
		Autoscaling: autoscaling,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAutoscalingSetGroup unmarshals an instance of AutoscalingSetGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroup)
	err = core.UnmarshalModel(m, "autoscaling", &obj.Autoscaling, UnmarshalAutoscalingSetGroupAutoscaling)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscaling : AutoscalingSetGroupAutoscaling struct
// Models which "extend" this model:
// - AutoscalingSetGroupAutoscalingAutoscalingDiskGroup
// - AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup
// - AutoscalingSetGroupAutoscalingAutoscalingCPUGroup
type AutoscalingSetGroupAutoscaling struct {
	Disk *AutoscalingDiskGroupDisk `json:"disk,omitempty"`

	Memory *AutoscalingMemoryGroupMemory `json:"memory,omitempty"`

	Cpu *AutoscalingCPUGroup `json:"cpu,omitempty"`
}

func (*AutoscalingSetGroupAutoscaling) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

type AutoscalingSetGroupAutoscalingIntf interface {
	isaAutoscalingSetGroupAutoscaling() bool
}

// UnmarshalAutoscalingSetGroupAutoscaling unmarshals an instance of AutoscalingSetGroupAutoscaling from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscaling)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalAutoscalingDiskGroupDisk)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalAutoscalingMemoryGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalAutoscalingCPUGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Backup : Backup struct
type Backup struct {
	// ID of this backup.
	ID *string `json:"id,omitempty"`

	// ID of the deployment this backup relates to.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// The type of backup.
	Type *string `json:"type,omitempty"`

	// The status of this backup.
	Status *string `json:"status,omitempty"`

	// Is this backup available to download?.
	IsDownloadable *bool `json:"is_downloadable,omitempty"`

	// Can this backup be used to restore an instance?.
	IsRestorable *bool `json:"is_restorable,omitempty"`

	// Date and time when this backup was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// Constants associated with the Backup.Type property.
// The type of backup.
const (
	Backup_Type_OnDemand = "on_demand"
	Backup_Type_Scheduled = "scheduled"
)

// Constants associated with the Backup.Status property.
// The status of this backup.
const (
	Backup_Status_Completed = "completed"
	Backup_Status_Failed = "failed"
	Backup_Status_Running = "running"
)


// UnmarshalBackup unmarshals an instance of Backup from the specified map of raw messages.
func UnmarshalBackup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_downloadable", &obj.IsDownloadable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_restorable", &obj.IsRestorable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Backups : Backups struct
type Backups struct {
	Backups []Backup `json:"backups,omitempty"`
}


// UnmarshalBackups unmarshals an instance of Backups from the specified map of raw messages.
func UnmarshalBackups(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backups)
	err = core.UnmarshalModel(m, "backups", &obj.Backups, UnmarshalBackup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChangeUserPasswordOptions : The ChangeUserPassword options.
type ChangeUserPasswordOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	Username *string `json:"username" validate:"required,ne="`

	User *APasswordSettingUser `json:"user,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewChangeUserPasswordOptions : Instantiate ChangeUserPasswordOptions
func (*IbmCloudDatabasesV5) NewChangeUserPasswordOptions(id string, userType string, username string) *ChangeUserPasswordOptions {
	return &ChangeUserPasswordOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		Username: core.StringPtr(username),
	}
}

// SetID : Allow user to set ID
func (options *ChangeUserPasswordOptions) SetID(id string) *ChangeUserPasswordOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *ChangeUserPasswordOptions) SetUserType(userType string) *ChangeUserPasswordOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUsername : Allow user to set Username
func (options *ChangeUserPasswordOptions) SetUsername(username string) *ChangeUserPasswordOptions {
	options.Username = core.StringPtr(username)
	return options
}

// SetUser : Allow user to set User
func (options *ChangeUserPasswordOptions) SetUser(user *APasswordSettingUser) *ChangeUserPasswordOptions {
	options.User = user
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ChangeUserPasswordOptions) SetHeaders(param map[string]string) *ChangeUserPasswordOptions {
	options.Headers = param
	return options
}

// ChangeUserPasswordResponse : ChangeUserPasswordResponse struct
type ChangeUserPasswordResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalChangeUserPasswordResponse unmarshals an instance of ChangeUserPasswordResponse from the specified map of raw messages.
func UnmarshalChangeUserPasswordResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChangeUserPasswordResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChoicePropertySchema : Choice Property Schema.
type ChoicePropertySchema struct {
	// Whether the setting is customer-configurable.
	CustomerConfigurable *bool `json:"customer_configurable,omitempty"`

	// The default value of the setting.
	Default *int64 `json:"default,omitempty"`

	// The description of the default value.
	DefaultDescription *string `json:"default_description,omitempty"`

	// The description of the setting.
	Description *string `json:"description,omitempty"`

	// The type of this setting (e.g., string, integer).
	Kind *string `json:"kind,omitempty"`

	// Whether or not changing this setting will restart the database.
	RequiresRestart *bool `json:"requires_restart,omitempty"`

	// The valid choices for this setting.
	Choices []string `json:"choices,omitempty"`
}


// UnmarshalChoicePropertySchema unmarshals an instance of ChoicePropertySchema from the specified map of raw messages.
func UnmarshalChoicePropertySchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ChoicePropertySchema)
	err = core.UnmarshalPrimitive(m, "customer_configurable", &obj.CustomerConfigurable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_description", &obj.DefaultDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requires_restart", &obj.RequiresRestart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "choices", &obj.Choices)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CompleteConnectionDeprecatedOptions : The CompleteConnectionDeprecated options.
type CompleteConnectionDeprecatedOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Password to be substituted into the response.
	Password *string `json:"password,omitempty"`

	// Optional certificate root path to prepend certificate names. Certificates would be stored in this directory for use
	// by other commands.
	CertificateRoot *string `json:"certificate_root,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCompleteConnectionDeprecatedOptions : Instantiate CompleteConnectionDeprecatedOptions
func (*IbmCloudDatabasesV5) NewCompleteConnectionDeprecatedOptions(id string, userType string, userID string) *CompleteConnectionDeprecatedOptions {
	return &CompleteConnectionDeprecatedOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		UserID: core.StringPtr(userID),
	}
}

// SetID : Allow user to set ID
func (options *CompleteConnectionDeprecatedOptions) SetID(id string) *CompleteConnectionDeprecatedOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *CompleteConnectionDeprecatedOptions) SetUserType(userType string) *CompleteConnectionDeprecatedOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUserID : Allow user to set UserID
func (options *CompleteConnectionDeprecatedOptions) SetUserID(userID string) *CompleteConnectionDeprecatedOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetPassword : Allow user to set Password
func (options *CompleteConnectionDeprecatedOptions) SetPassword(password string) *CompleteConnectionDeprecatedOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetCertificateRoot : Allow user to set CertificateRoot
func (options *CompleteConnectionDeprecatedOptions) SetCertificateRoot(certificateRoot string) *CompleteConnectionDeprecatedOptions {
	options.CertificateRoot = core.StringPtr(certificateRoot)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CompleteConnectionDeprecatedOptions) SetHeaders(param map[string]string) *CompleteConnectionDeprecatedOptions {
	options.Headers = param
	return options
}

// CompleteConnectionOptions : The CompleteConnection options.
type CompleteConnectionOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type of `database` is the only currently supported value.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Endpoint Type. The select endpoint must be enabled on the deployment before its connection information can be
	// fetched.
	EndpointType *string `json:"endpoint_type" validate:"required,ne="`

	// Password to be substituted into the response.
	Password *string `json:"password,omitempty"`

	// Optional certificate root path to prepend certificate names. Certificates would be stored in this directory for use
	// by other commands.
	CertificateRoot *string `json:"certificate_root,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CompleteConnectionOptions.EndpointType property.
// Endpoint Type. The select endpoint must be enabled on the deployment before its connection information can be
// fetched.
const (
	CompleteConnectionOptions_EndpointType_Private = "private"
	CompleteConnectionOptions_EndpointType_Public = "public"
)

// NewCompleteConnectionOptions : Instantiate CompleteConnectionOptions
func (*IbmCloudDatabasesV5) NewCompleteConnectionOptions(id string, userType string, userID string, endpointType string) *CompleteConnectionOptions {
	return &CompleteConnectionOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		UserID: core.StringPtr(userID),
		EndpointType: core.StringPtr(endpointType),
	}
}

// SetID : Allow user to set ID
func (options *CompleteConnectionOptions) SetID(id string) *CompleteConnectionOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *CompleteConnectionOptions) SetUserType(userType string) *CompleteConnectionOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUserID : Allow user to set UserID
func (options *CompleteConnectionOptions) SetUserID(userID string) *CompleteConnectionOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetEndpointType : Allow user to set EndpointType
func (options *CompleteConnectionOptions) SetEndpointType(endpointType string) *CompleteConnectionOptions {
	options.EndpointType = core.StringPtr(endpointType)
	return options
}

// SetPassword : Allow user to set Password
func (options *CompleteConnectionOptions) SetPassword(password string) *CompleteConnectionOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetCertificateRoot : Allow user to set CertificateRoot
func (options *CompleteConnectionOptions) SetCertificateRoot(certificateRoot string) *CompleteConnectionOptions {
	options.CertificateRoot = core.StringPtr(certificateRoot)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CompleteConnectionOptions) SetHeaders(param map[string]string) *CompleteConnectionOptions {
	options.Headers = param
	return options
}

// ConfigurationSchema : Database Configuration Schema.
type ConfigurationSchema struct {
	Schema ConfigurationSchemaSchemaIntf `json:"schema" validate:"required"`
}


// UnmarshalConfigurationSchema unmarshals an instance of ConfigurationSchema from the specified map of raw messages.
func UnmarshalConfigurationSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationSchema)
	err = core.UnmarshalModel(m, "schema", &obj.Schema, UnmarshalConfigurationSchemaSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationSchemaSchema : ConfigurationSchemaSchema struct
// Models which "extend" this model:
// - ConfigurationSchemaSchemaPGConfigurationSchema
// - ConfigurationSchemaSchemaRedisConfigurationSchema
type ConfigurationSchemaSchema struct {
	// Integer Property Schema.
	MaxConnections *IntegerPropertySchema `json:"max_connections,omitempty"`

	// Integer Property Schema.
	MaxPreparedConnections *IntegerPropertySchema `json:"max_prepared_connections,omitempty"`

	// Integer Property Schema.
	BackupRetentionPeriod *IntegerPropertySchema `json:"backup_retention_period,omitempty"`

	// Integer Property Schema.
	DeadlockTimeout *IntegerPropertySchema `json:"deadlock_timeout,omitempty"`

	// Integer Property Schema.
	EffectiveIoConcurrency *IntegerPropertySchema `json:"effective_io_concurrency,omitempty"`

	// Integer Property Schema.
	MaxReplicationSlots *IntegerPropertySchema `json:"max_replication_slots,omitempty"`

	// Integer Property Schema.
	MaxWalSenders *IntegerPropertySchema `json:"max_wal_senders,omitempty"`

	// Integer Property Schema.
	SharedBuffers *IntegerPropertySchema `json:"shared_buffers,omitempty"`

	// Choice Property Schema.
	SynchronousCommit *ChoicePropertySchema `json:"synchronous_commit,omitempty"`

	// Choice Property Schema.
	WalLevel *ChoicePropertySchema `json:"wal_level,omitempty"`

	// Integer Property Schema.
	ArchiveTimeout *IntegerPropertySchema `json:"archive_timeout,omitempty"`

	// Integer Property Schema.
	LogMinDurationStatement *IntegerPropertySchema `json:"log_min_duration_statement,omitempty"`

	// Integer Property Schema.
	MaxmemoryRedis *IntegerPropertySchema `json:"maxmemory-redis,omitempty"`

	// Choice Property Schema.
	MaxmemoryPolicy *ChoicePropertySchema `json:"maxmemory-policy,omitempty"`

	// Choice Property Schema.
	Appendonly *ChoicePropertySchema `json:"appendonly,omitempty"`

	// Integer Property Schema.
	MaxmemorySamples *IntegerPropertySchema `json:"maxmemory-samples,omitempty"`

	// Choice Property Schema.
	StopWritesOnBgsaveError *ChoicePropertySchema `json:"stop-writes-on-bgsave-error,omitempty"`
}

func (*ConfigurationSchemaSchema) isaConfigurationSchemaSchema() bool {
	return true
}

type ConfigurationSchemaSchemaIntf interface {
	isaConfigurationSchemaSchema() bool
}

// UnmarshalConfigurationSchemaSchema unmarshals an instance of ConfigurationSchemaSchema from the specified map of raw messages.
func UnmarshalConfigurationSchemaSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationSchemaSchema)
	err = core.UnmarshalModel(m, "max_connections", &obj.MaxConnections, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_prepared_connections", &obj.MaxPreparedConnections, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "backup_retention_period", &obj.BackupRetentionPeriod, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deadlock_timeout", &obj.DeadlockTimeout, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "effective_io_concurrency", &obj.EffectiveIoConcurrency, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_replication_slots", &obj.MaxReplicationSlots, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_wal_senders", &obj.MaxWalSenders, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "shared_buffers", &obj.SharedBuffers, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "synchronous_commit", &obj.SynchronousCommit, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "wal_level", &obj.WalLevel, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "archive_timeout", &obj.ArchiveTimeout, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_min_duration_statement", &obj.LogMinDurationStatement, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "maxmemory-redis", &obj.MaxmemoryRedis, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "maxmemory-policy", &obj.MaxmemoryPolicy, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "appendonly", &obj.Appendonly, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "maxmemory-samples", &obj.MaxmemorySamples, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stop-writes-on-bgsave-error", &obj.StopWritesOnBgsaveError, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Connection : Connection struct
type Connection struct {
	Connection ConnectionConnectionIntf `json:"connection" validate:"required"`
}


// UnmarshalConnection unmarshals an instance of Connection from the specified map of raw messages.
func UnmarshalConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Connection)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalConnectionConnection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionCLI : CLI Connection.
type ConnectionCLI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// A map of environment variables for a CLI connection.
	Environment map[string]string `json:"environment,omitempty"`

	// The name of the executable the CLI should run.
	Bin *string `json:"bin,omitempty"`

	// Sets of arguments to call the executable with. The outer array corresponds to a possible way to call the CLI; the
	// inner array is the set of arguments to use with that call.
	Arguments [][]string `json:"arguments,omitempty"`

	Certificate *ConnectionCLICertificate `json:"certificate,omitempty"`
}


// UnmarshalConnectionCLI unmarshals an instance of ConnectionCLI from the specified map of raw messages.
func UnmarshalConnectionCLI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionCLI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bin", &obj.Bin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "arguments", &obj.Arguments)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalConnectionCLICertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionCLICertificate : ConnectionCLICertificate struct
type ConnectionCLICertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalConnectionCLICertificate unmarshals an instance of ConnectionCLICertificate from the specified map of raw messages.
func UnmarshalConnectionCLICertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionCLICertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnection : ConnectionConnection struct
// Models which "extend" this model:
// - ConnectionConnectionPostgreSQLConnection
// - ConnectionConnectionRedisConnection
// - ConnectionConnectionElasticsearchConnection
// - ConnectionConnectionRabbitMQConnection
// - ConnectionConnectionEtcdConnection
// - ConnectionConnectionMongoDBConnection
type ConnectionConnection struct {
	// Connection information for drivers and libraries.
	Postgres *PostgreSQLConnectionURI `json:"postgres,omitempty"`

	// Connection information for psql.
	Cli *ConnectionCLI `json:"cli,omitempty"`

	// Connection information for drivers and libraries.
	Rediss *RedisConnectionURI `json:"rediss,omitempty"`

	// Elasticsearch Connection information for drivers and libraries.
	Https *ElasticsearchConnectionHTTPS `json:"https,omitempty"`

	// RabbitMQ Connection information for AMQPS drivers and libraries.
	Amqps *RabbitMQConnectionAMQPS `json:"amqps,omitempty"`

	// RabbitMQ Connection information for MQTTS drivers and libraries.
	Mqtts *RabbitMQConnectionMQTTS `json:"mqtts,omitempty"`

	// RabbitMQ Connection information for STOMP drivers and libraries.
	StompSsl *RabbitMQConnectionStompSSL `json:"stomp_ssl,omitempty"`

	// GRPC(etcd3) Connection information for drivers and libraries.
	Grpc *GRPCConnectionURI `json:"grpc,omitempty"`

	// MongoDB Connection information for drivers and libraries.
	Mongodb *MongoDBConnectionURI `json:"mongodb,omitempty"`
}

func (*ConnectionConnection) isaConnectionConnection() bool {
	return true
}

type ConnectionConnectionIntf interface {
	isaConnectionConnection() bool
}

// UnmarshalConnectionConnection unmarshals an instance of ConnectionConnection from the specified map of raw messages.
func UnmarshalConnectionConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnection)
	err = core.UnmarshalModel(m, "postgres", &obj.Postgres, UnmarshalPostgreSQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rediss", &obj.Rediss, UnmarshalRedisConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "https", &obj.Https, UnmarshalElasticsearchConnectionHTTPS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "amqps", &obj.Amqps, UnmarshalRabbitMQConnectionAMQPS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mqtts", &obj.Mqtts, UnmarshalRabbitMQConnectionMQTTS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stomp_ssl", &obj.StompSsl, UnmarshalRabbitMQConnectionStompSSL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "grpc", &obj.Grpc, UnmarshalGRPCConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mongodb", &obj.Mongodb, UnmarshalMongoDBConnectionURI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDatabaseUserOptions : The CreateDatabaseUser options.
type CreateDatabaseUserOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	User *CreateDatabaseUserRequestUser `json:"user,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDatabaseUserOptions : Instantiate CreateDatabaseUserOptions
func (*IbmCloudDatabasesV5) NewCreateDatabaseUserOptions(id string, userType string) *CreateDatabaseUserOptions {
	return &CreateDatabaseUserOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
	}
}

// SetID : Allow user to set ID
func (options *CreateDatabaseUserOptions) SetID(id string) *CreateDatabaseUserOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *CreateDatabaseUserOptions) SetUserType(userType string) *CreateDatabaseUserOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUser : Allow user to set User
func (options *CreateDatabaseUserOptions) SetUser(user *CreateDatabaseUserRequestUser) *CreateDatabaseUserOptions {
	options.User = user
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDatabaseUserOptions) SetHeaders(param map[string]string) *CreateDatabaseUserOptions {
	options.Headers = param
	return options
}

// CreateDatabaseUserRequestUser : CreateDatabaseUserRequestUser struct
type CreateDatabaseUserRequestUser struct {
	// User type for new user.
	UserType *string `json:"user_type,omitempty"`

	// Username for new user.
	Username *string `json:"username,omitempty"`

	// Password for new user.
	Password *string `json:"password,omitempty"`
}


// UnmarshalCreateDatabaseUserRequestUser unmarshals an instance of CreateDatabaseUserRequestUser from the specified map of raw messages.
func UnmarshalCreateDatabaseUserRequestUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateDatabaseUserRequestUser)
	err = core.UnmarshalPrimitive(m, "user_type", &obj.UserType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDatabaseUserResponse : CreateDatabaseUserResponse struct
type CreateDatabaseUserResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalCreateDatabaseUserResponse unmarshals an instance of CreateDatabaseUserResponse from the specified map of raw messages.
func UnmarshalCreateDatabaseUserResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateDatabaseUserResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateLogicalReplicationSlotOptions : The CreateLogicalReplicationSlot options.
type CreateLogicalReplicationSlotOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	LogicalReplicationSlot *LogicalReplicationSlotLogicalReplicationSlot `json:"logical_replication_slot,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateLogicalReplicationSlotOptions : Instantiate CreateLogicalReplicationSlotOptions
func (*IbmCloudDatabasesV5) NewCreateLogicalReplicationSlotOptions(id string) *CreateLogicalReplicationSlotOptions {
	return &CreateLogicalReplicationSlotOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *CreateLogicalReplicationSlotOptions) SetID(id string) *CreateLogicalReplicationSlotOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetLogicalReplicationSlot : Allow user to set LogicalReplicationSlot
func (options *CreateLogicalReplicationSlotOptions) SetLogicalReplicationSlot(logicalReplicationSlot *LogicalReplicationSlotLogicalReplicationSlot) *CreateLogicalReplicationSlotOptions {
	options.LogicalReplicationSlot = logicalReplicationSlot
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLogicalReplicationSlotOptions) SetHeaders(param map[string]string) *CreateLogicalReplicationSlotOptions {
	options.Headers = param
	return options
}

// CreateLogicalReplicationSlotResponse : CreateLogicalReplicationSlotResponse struct
type CreateLogicalReplicationSlotResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalCreateLogicalReplicationSlotResponse unmarshals an instance of CreateLogicalReplicationSlotResponse from the specified map of raw messages.
func UnmarshalCreateLogicalReplicationSlotResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateLogicalReplicationSlotResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDatabaseUserOptions : The DeleteDatabaseUser options.
type DeleteDatabaseUserOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// Username.
	Username *string `json:"username" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDatabaseUserOptions : Instantiate DeleteDatabaseUserOptions
func (*IbmCloudDatabasesV5) NewDeleteDatabaseUserOptions(id string, userType string, username string) *DeleteDatabaseUserOptions {
	return &DeleteDatabaseUserOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		Username: core.StringPtr(username),
	}
}

// SetID : Allow user to set ID
func (options *DeleteDatabaseUserOptions) SetID(id string) *DeleteDatabaseUserOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *DeleteDatabaseUserOptions) SetUserType(userType string) *DeleteDatabaseUserOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUsername : Allow user to set Username
func (options *DeleteDatabaseUserOptions) SetUsername(username string) *DeleteDatabaseUserOptions {
	options.Username = core.StringPtr(username)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDatabaseUserOptions) SetHeaders(param map[string]string) *DeleteDatabaseUserOptions {
	options.Headers = param
	return options
}

// DeleteDatabaseUserResponse : DeleteDatabaseUserResponse struct
type DeleteDatabaseUserResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalDeleteDatabaseUserResponse unmarshals an instance of DeleteDatabaseUserResponse from the specified map of raw messages.
func UnmarshalDeleteDatabaseUserResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteDatabaseUserResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteLogicalReplicationSlotOptions : The DeleteLogicalReplicationSlot options.
type DeleteLogicalReplicationSlotOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Name of the logical replication slot.
	Name *string `json:"name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLogicalReplicationSlotOptions : Instantiate DeleteLogicalReplicationSlotOptions
func (*IbmCloudDatabasesV5) NewDeleteLogicalReplicationSlotOptions(id string, name string) *DeleteLogicalReplicationSlotOptions {
	return &DeleteLogicalReplicationSlotOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
	}
}

// SetID : Allow user to set ID
func (options *DeleteLogicalReplicationSlotOptions) SetID(id string) *DeleteLogicalReplicationSlotOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *DeleteLogicalReplicationSlotOptions) SetName(name string) *DeleteLogicalReplicationSlotOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLogicalReplicationSlotOptions) SetHeaders(param map[string]string) *DeleteLogicalReplicationSlotOptions {
	options.Headers = param
	return options
}

// DeleteLogicalReplicationSlotResponse : DeleteLogicalReplicationSlotResponse struct
type DeleteLogicalReplicationSlotResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalDeleteLogicalReplicationSlotResponse unmarshals an instance of DeleteLogicalReplicationSlotResponse from the specified map of raw messages.
func UnmarshalDeleteLogicalReplicationSlotResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteLogicalReplicationSlotResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteWhitelistEntryOptions : The DeleteWhitelistEntry options.
type DeleteWhitelistEntryOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// An IPv4 address or a CIDR range (netmasked IPv4 address).
	Ipaddress *string `json:"ipaddress" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteWhitelistEntryOptions : Instantiate DeleteWhitelistEntryOptions
func (*IbmCloudDatabasesV5) NewDeleteWhitelistEntryOptions(id string, ipaddress string) *DeleteWhitelistEntryOptions {
	return &DeleteWhitelistEntryOptions{
		ID: core.StringPtr(id),
		Ipaddress: core.StringPtr(ipaddress),
	}
}

// SetID : Allow user to set ID
func (options *DeleteWhitelistEntryOptions) SetID(id string) *DeleteWhitelistEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIpaddress : Allow user to set Ipaddress
func (options *DeleteWhitelistEntryOptions) SetIpaddress(ipaddress string) *DeleteWhitelistEntryOptions {
	options.Ipaddress = core.StringPtr(ipaddress)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWhitelistEntryOptions) SetHeaders(param map[string]string) *DeleteWhitelistEntryOptions {
	options.Headers = param
	return options
}

// DeleteWhitelistEntryResponse : DeleteWhitelistEntryResponse struct
type DeleteWhitelistEntryResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalDeleteWhitelistEntryResponse unmarshals an instance of DeleteWhitelistEntryResponse from the specified map of raw messages.
func UnmarshalDeleteWhitelistEntryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteWhitelistEntryResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Deployables : Deployable databases with their version information.
type Deployables struct {
	// Deployment type - typically the name of the database.
	Type *string `json:"type,omitempty"`

	// An array of versions of the database, their status, preferedness, and transitions.
	Versions []DeployablesVersionsItem `json:"versions,omitempty"`
}


// UnmarshalDeployables unmarshals an instance of Deployables from the specified map of raw messages.
func UnmarshalDeployables(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Deployables)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "versions", &obj.Versions, UnmarshalDeployablesVersionsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeployablesVersionsItem : DeployablesVersionsItem struct
type DeployablesVersionsItem struct {
	// The version number.
	Version *string `json:"version,omitempty"`

	// The status of this version: To be finalized.
	Status *string `json:"status,omitempty"`

	// Should this version be preferred over others?.
	IsPreferred *bool `json:"is_preferred,omitempty"`

	// versions that this version can be upgraded to.
	Transitions []DeployablesVersionsItemTransitionsItem `json:"transitions,omitempty"`
}

// Constants associated with the DeployablesVersionsItem.Status property.
// The status of this version: To be finalized.
const (
	DeployablesVersionsItem_Status_Beta = "beta"
	DeployablesVersionsItem_Status_Deprecated = "deprecated"
	DeployablesVersionsItem_Status_Stable = "stable"
)


// UnmarshalDeployablesVersionsItem unmarshals an instance of DeployablesVersionsItem from the specified map of raw messages.
func UnmarshalDeployablesVersionsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeployablesVersionsItem)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_preferred", &obj.IsPreferred)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "transitions", &obj.Transitions, UnmarshalDeployablesVersionsItemTransitionsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeployablesVersionsItemTransitionsItem : DeployablesVersionsItemTransitionsItem struct
type DeployablesVersionsItemTransitionsItem struct {
	// The database type.
	Application *string `json:"application,omitempty"`

	// method of going from from_version to to_version.
	Method *string `json:"method,omitempty"`

	// The version the transition in from.
	FromVersion *string `json:"from_version,omitempty"`

	// The version the transition is to.
	ToVersion *string `json:"to_version,omitempty"`
}


// UnmarshalDeployablesVersionsItemTransitionsItem unmarshals an instance of DeployablesVersionsItemTransitionsItem from the specified map of raw messages.
func UnmarshalDeployablesVersionsItemTransitionsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeployablesVersionsItemTransitionsItem)
	err = core.UnmarshalPrimitive(m, "application", &obj.Application)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from_version", &obj.FromVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "to_version", &obj.ToVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Deployment : Deployment struct
type Deployment struct {
	// ID of this deployment.
	ID *string `json:"id,omitempty"`

	// Readable name of this deployment.
	Name *string `json:"name,omitempty"`

	// Database type within this deployment.
	Type *string `json:"type,omitempty"`

	// Platform-specific options for this deployment.
	PlatformOptions interface{} `json:"platform_options,omitempty"`

	// Version number of the database.
	Version *string `json:"version,omitempty"`

	// Login name of administration level user.
	AdminUsernames *string `json:"admin_usernames,omitempty"`

	// Whether access to this deployment is enabled from the public internet. This property can be modified by updating
	// this service instance through the Resource Controller API.
	EnablePublicEndpoints *bool `json:"enable_public_endpoints,omitempty"`

	// Whether access to this deployment is enabled from IBM Cloud via the IBM Cloud backbone network. This property can be
	// modified by updating this service instance through the Resource Controller API.
	EnablePrivateEndpoints *bool `json:"enable_private_endpoints,omitempty"`
}


// UnmarshalDeployment unmarshals an instance of Deployment from the specified map of raw messages.
func UnmarshalDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Deployment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "platform_options", &obj.PlatformOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "admin_usernames", &obj.AdminUsernames)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable_public_endpoints", &obj.EnablePublicEndpoints)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enable_private_endpoints", &obj.EnablePrivateEndpoints)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ElasticsearchConnectionHTTPS : ElasticsearchConnectionHTTPS struct
type ElasticsearchConnectionHTTPS struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []ElasticsearchConnectionHTTPSHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *ElasticsearchConnectionHTTPSAuthentication `json:"authentication,omitempty"`

	Certificate *ElasticsearchConnectionHTTPSCertificate `json:"certificate,omitempty"`
}


// UnmarshalElasticsearchConnectionHTTPS unmarshals an instance of ElasticsearchConnectionHTTPS from the specified map of raw messages.
func UnmarshalElasticsearchConnectionHTTPS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ElasticsearchConnectionHTTPS)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalElasticsearchConnectionHTTPSHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalElasticsearchConnectionHTTPSAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalElasticsearchConnectionHTTPSCertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ElasticsearchConnectionHTTPSAuthentication : ElasticsearchConnectionHTTPSAuthentication struct
type ElasticsearchConnectionHTTPSAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalElasticsearchConnectionHTTPSAuthentication unmarshals an instance of ElasticsearchConnectionHTTPSAuthentication from the specified map of raw messages.
func UnmarshalElasticsearchConnectionHTTPSAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ElasticsearchConnectionHTTPSAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ElasticsearchConnectionHTTPSCertificate : ElasticsearchConnectionHTTPSCertificate struct
type ElasticsearchConnectionHTTPSCertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalElasticsearchConnectionHTTPSCertificate unmarshals an instance of ElasticsearchConnectionHTTPSCertificate from the specified map of raw messages.
func UnmarshalElasticsearchConnectionHTTPSCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ElasticsearchConnectionHTTPSCertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ElasticsearchConnectionHTTPSHostsItem : ElasticsearchConnectionHTTPSHostsItem struct
type ElasticsearchConnectionHTTPSHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalElasticsearchConnectionHTTPSHostsItem unmarshals an instance of ElasticsearchConnectionHTTPSHostsItem from the specified map of raw messages.
func UnmarshalElasticsearchConnectionHTTPSHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ElasticsearchConnectionHTTPSHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileSyncOptions : The FileSync options.
type FileSyncOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFileSyncOptions : Instantiate FileSyncOptions
func (*IbmCloudDatabasesV5) NewFileSyncOptions(id string) *FileSyncOptions {
	return &FileSyncOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *FileSyncOptions) SetID(id string) *FileSyncOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *FileSyncOptions) SetHeaders(param map[string]string) *FileSyncOptions {
	options.Headers = param
	return options
}

// FileSyncResponse : FileSyncResponse struct
type FileSyncResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalFileSyncResponse unmarshals an instance of FileSyncResponse from the specified map of raw messages.
func UnmarshalFileSyncResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileSyncResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GRPCConnectionURI : GRPCConnectionURI struct
type GRPCConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []GRPCConnectionURIHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *GRPCConnectionURIAuthentication `json:"authentication,omitempty"`

	Certificate *GRPCConnectionURICertificate `json:"certificate,omitempty"`
}


// UnmarshalGRPCConnectionURI unmarshals an instance of GRPCConnectionURI from the specified map of raw messages.
func UnmarshalGRPCConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GRPCConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalGRPCConnectionURIHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalGRPCConnectionURIAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalGRPCConnectionURICertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GRPCConnectionURIAuthentication : GRPCConnectionURIAuthentication struct
type GRPCConnectionURIAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalGRPCConnectionURIAuthentication unmarshals an instance of GRPCConnectionURIAuthentication from the specified map of raw messages.
func UnmarshalGRPCConnectionURIAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GRPCConnectionURIAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GRPCConnectionURICertificate : GRPCConnectionURICertificate struct
type GRPCConnectionURICertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalGRPCConnectionURICertificate unmarshals an instance of GRPCConnectionURICertificate from the specified map of raw messages.
func UnmarshalGRPCConnectionURICertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GRPCConnectionURICertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GRPCConnectionURIHostsItem : GRPCConnectionURIHostsItem struct
type GRPCConnectionURIHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalGRPCConnectionURIHostsItem unmarshals an instance of GRPCConnectionURIHostsItem from the specified map of raw messages.
func UnmarshalGRPCConnectionURIHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GRPCConnectionURIHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAutoscalingConditionsOptions : The GetAutoscalingConditions options.
type GetAutoscalingConditionsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Group ID.
	GroupID *string `json:"group_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAutoscalingConditionsOptions : Instantiate GetAutoscalingConditionsOptions
func (*IbmCloudDatabasesV5) NewGetAutoscalingConditionsOptions(id string, groupID string) *GetAutoscalingConditionsOptions {
	return &GetAutoscalingConditionsOptions{
		ID: core.StringPtr(id),
		GroupID: core.StringPtr(groupID),
	}
}

// SetID : Allow user to set ID
func (options *GetAutoscalingConditionsOptions) SetID(id string) *GetAutoscalingConditionsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetGroupID : Allow user to set GroupID
func (options *GetAutoscalingConditionsOptions) SetGroupID(groupID string) *GetAutoscalingConditionsOptions {
	options.GroupID = core.StringPtr(groupID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAutoscalingConditionsOptions) SetHeaders(param map[string]string) *GetAutoscalingConditionsOptions {
	options.Headers = param
	return options
}

// GetAutoscalingConditionsResponse : GetAutoscalingConditionsResponse struct
type GetAutoscalingConditionsResponse struct {
	Autoscaling *AutoscalingGroup `json:"autoscaling,omitempty"`
}


// UnmarshalGetAutoscalingConditionsResponse unmarshals an instance of GetAutoscalingConditionsResponse from the specified map of raw messages.
func UnmarshalGetAutoscalingConditionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetAutoscalingConditionsResponse)
	err = core.UnmarshalModel(m, "autoscaling", &obj.Autoscaling, UnmarshalAutoscalingGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBackupInfoOptions : The GetBackupInfo options.
type GetBackupInfoOptions struct {
	// Backup ID.
	BackupID *string `json:"backup_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBackupInfoOptions : Instantiate GetBackupInfoOptions
func (*IbmCloudDatabasesV5) NewGetBackupInfoOptions(backupID string) *GetBackupInfoOptions {
	return &GetBackupInfoOptions{
		BackupID: core.StringPtr(backupID),
	}
}

// SetBackupID : Allow user to set BackupID
func (options *GetBackupInfoOptions) SetBackupID(backupID string) *GetBackupInfoOptions {
	options.BackupID = core.StringPtr(backupID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBackupInfoOptions) SetHeaders(param map[string]string) *GetBackupInfoOptions {
	options.Headers = param
	return options
}

// GetBackupInfoResponse : GetBackupInfoResponse struct
type GetBackupInfoResponse struct {
	Backup *Backup `json:"backup,omitempty"`
}


// UnmarshalGetBackupInfoResponse unmarshals an instance of GetBackupInfoResponse from the specified map of raw messages.
func UnmarshalGetBackupInfoResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBackupInfoResponse)
	err = core.UnmarshalModel(m, "backup", &obj.Backup, UnmarshalBackup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetConnectionDeprecatedOptions : The GetConnectionDeprecated options.
type GetConnectionDeprecatedOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Optional certificate root path to prepend certificate names. Certificates would be stored in this directory for use
	// by other commands.
	CertificateRoot *string `json:"certificate_root,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConnectionDeprecatedOptions : Instantiate GetConnectionDeprecatedOptions
func (*IbmCloudDatabasesV5) NewGetConnectionDeprecatedOptions(id string, userType string, userID string) *GetConnectionDeprecatedOptions {
	return &GetConnectionDeprecatedOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		UserID: core.StringPtr(userID),
	}
}

// SetID : Allow user to set ID
func (options *GetConnectionDeprecatedOptions) SetID(id string) *GetConnectionDeprecatedOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *GetConnectionDeprecatedOptions) SetUserType(userType string) *GetConnectionDeprecatedOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUserID : Allow user to set UserID
func (options *GetConnectionDeprecatedOptions) SetUserID(userID string) *GetConnectionDeprecatedOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetCertificateRoot : Allow user to set CertificateRoot
func (options *GetConnectionDeprecatedOptions) SetCertificateRoot(certificateRoot string) *GetConnectionDeprecatedOptions {
	options.CertificateRoot = core.StringPtr(certificateRoot)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetConnectionDeprecatedOptions) SetHeaders(param map[string]string) *GetConnectionDeprecatedOptions {
	options.Headers = param
	return options
}

// GetConnectionOptions : The GetConnection options.
type GetConnectionOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User type.
	UserType *string `json:"user_type" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
	EndpointType *string `json:"endpoint_type" validate:"required,ne="`

	// Optional certificate root path to prepend certificate names. Certificates would be stored in this directory for use
	// by other commands.
	CertificateRoot *string `json:"certificate_root,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetConnectionOptions.EndpointType property.
// Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
const (
	GetConnectionOptions_EndpointType_Private = "private"
	GetConnectionOptions_EndpointType_Public = "public"
)

// NewGetConnectionOptions : Instantiate GetConnectionOptions
func (*IbmCloudDatabasesV5) NewGetConnectionOptions(id string, userType string, userID string, endpointType string) *GetConnectionOptions {
	return &GetConnectionOptions{
		ID: core.StringPtr(id),
		UserType: core.StringPtr(userType),
		UserID: core.StringPtr(userID),
		EndpointType: core.StringPtr(endpointType),
	}
}

// SetID : Allow user to set ID
func (options *GetConnectionOptions) SetID(id string) *GetConnectionOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserType : Allow user to set UserType
func (options *GetConnectionOptions) SetUserType(userType string) *GetConnectionOptions {
	options.UserType = core.StringPtr(userType)
	return options
}

// SetUserID : Allow user to set UserID
func (options *GetConnectionOptions) SetUserID(userID string) *GetConnectionOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetEndpointType : Allow user to set EndpointType
func (options *GetConnectionOptions) SetEndpointType(endpointType string) *GetConnectionOptions {
	options.EndpointType = core.StringPtr(endpointType)
	return options
}

// SetCertificateRoot : Allow user to set CertificateRoot
func (options *GetConnectionOptions) SetCertificateRoot(certificateRoot string) *GetConnectionOptions {
	options.CertificateRoot = core.StringPtr(certificateRoot)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetConnectionOptions) SetHeaders(param map[string]string) *GetConnectionOptions {
	options.Headers = param
	return options
}

// GetDatabaseConfigurationSchemaOptions : The GetDatabaseConfigurationSchema options.
type GetDatabaseConfigurationSchemaOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDatabaseConfigurationSchemaOptions : Instantiate GetDatabaseConfigurationSchemaOptions
func (*IbmCloudDatabasesV5) NewGetDatabaseConfigurationSchemaOptions(id string) *GetDatabaseConfigurationSchemaOptions {
	return &GetDatabaseConfigurationSchemaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetDatabaseConfigurationSchemaOptions) SetID(id string) *GetDatabaseConfigurationSchemaOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDatabaseConfigurationSchemaOptions) SetHeaders(param map[string]string) *GetDatabaseConfigurationSchemaOptions {
	options.Headers = param
	return options
}

// GetDatabaseConfigurationSchemaResponse : GetDatabaseConfigurationSchemaResponse struct
type GetDatabaseConfigurationSchemaResponse struct {
	// Database Configuration Schema.
	Schema *ConfigurationSchema `json:"schema,omitempty"`
}


// UnmarshalGetDatabaseConfigurationSchemaResponse unmarshals an instance of GetDatabaseConfigurationSchemaResponse from the specified map of raw messages.
func UnmarshalGetDatabaseConfigurationSchemaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDatabaseConfigurationSchemaResponse)
	err = core.UnmarshalModel(m, "schema", &obj.Schema, UnmarshalConfigurationSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDefaultScalingGroupsOptions : The GetDefaultScalingGroups options.
type GetDefaultScalingGroupsOptions struct {
	// Database type name.
	Type *string `json:"type" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetDefaultScalingGroupsOptions.Type property.
// Database type name.
const (
	GetDefaultScalingGroupsOptions_Type_Etcd = "etcd"
	GetDefaultScalingGroupsOptions_Type_Postgresql = "postgresql"
)

// NewGetDefaultScalingGroupsOptions : Instantiate GetDefaultScalingGroupsOptions
func (*IbmCloudDatabasesV5) NewGetDefaultScalingGroupsOptions(typeVar string) *GetDefaultScalingGroupsOptions {
	return &GetDefaultScalingGroupsOptions{
		Type: core.StringPtr(typeVar),
	}
}

// SetType : Allow user to set Type
func (options *GetDefaultScalingGroupsOptions) SetType(typeVar string) *GetDefaultScalingGroupsOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDefaultScalingGroupsOptions) SetHeaders(param map[string]string) *GetDefaultScalingGroupsOptions {
	options.Headers = param
	return options
}

// GetDeployablesOptions : The GetDeployables options.
type GetDeployablesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeployablesOptions : Instantiate GetDeployablesOptions
func (*IbmCloudDatabasesV5) NewGetDeployablesOptions() *GetDeployablesOptions {
	return &GetDeployablesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetDeployablesOptions) SetHeaders(param map[string]string) *GetDeployablesOptions {
	options.Headers = param
	return options
}

// GetDeployablesResponse : GetDeployablesResponse struct
type GetDeployablesResponse struct {
	Deployables []Deployables `json:"deployables,omitempty"`
}


// UnmarshalGetDeployablesResponse unmarshals an instance of GetDeployablesResponse from the specified map of raw messages.
func UnmarshalGetDeployablesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDeployablesResponse)
	err = core.UnmarshalModel(m, "deployables", &obj.Deployables, UnmarshalDeployables)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDeploymentBackupsOptions : The GetDeploymentBackups options.
type GetDeploymentBackupsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentBackupsOptions : Instantiate GetDeploymentBackupsOptions
func (*IbmCloudDatabasesV5) NewGetDeploymentBackupsOptions(id string) *GetDeploymentBackupsOptions {
	return &GetDeploymentBackupsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetDeploymentBackupsOptions) SetID(id string) *GetDeploymentBackupsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDeploymentBackupsOptions) SetHeaders(param map[string]string) *GetDeploymentBackupsOptions {
	options.Headers = param
	return options
}

// GetDeploymentInfoOptions : The GetDeploymentInfo options.
type GetDeploymentInfoOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentInfoOptions : Instantiate GetDeploymentInfoOptions
func (*IbmCloudDatabasesV5) NewGetDeploymentInfoOptions(id string) *GetDeploymentInfoOptions {
	return &GetDeploymentInfoOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetDeploymentInfoOptions) SetID(id string) *GetDeploymentInfoOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDeploymentInfoOptions) SetHeaders(param map[string]string) *GetDeploymentInfoOptions {
	options.Headers = param
	return options
}

// GetDeploymentInfoResponse : GetDeploymentInfoResponse struct
type GetDeploymentInfoResponse struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}


// UnmarshalGetDeploymentInfoResponse unmarshals an instance of GetDeploymentInfoResponse from the specified map of raw messages.
func UnmarshalGetDeploymentInfoResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDeploymentInfoResponse)
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalDeployment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDeploymentScalingGroupsOptions : The GetDeploymentScalingGroups options.
type GetDeploymentScalingGroupsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentScalingGroupsOptions : Instantiate GetDeploymentScalingGroupsOptions
func (*IbmCloudDatabasesV5) NewGetDeploymentScalingGroupsOptions(id string) *GetDeploymentScalingGroupsOptions {
	return &GetDeploymentScalingGroupsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetDeploymentScalingGroupsOptions) SetID(id string) *GetDeploymentScalingGroupsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDeploymentScalingGroupsOptions) SetHeaders(param map[string]string) *GetDeploymentScalingGroupsOptions {
	options.Headers = param
	return options
}

// GetDeploymentTasksOptions : The GetDeploymentTasks options.
type GetDeploymentTasksOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDeploymentTasksOptions : Instantiate GetDeploymentTasksOptions
func (*IbmCloudDatabasesV5) NewGetDeploymentTasksOptions(id string) *GetDeploymentTasksOptions {
	return &GetDeploymentTasksOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetDeploymentTasksOptions) SetID(id string) *GetDeploymentTasksOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDeploymentTasksOptions) SetHeaders(param map[string]string) *GetDeploymentTasksOptions {
	options.Headers = param
	return options
}

// GetPITRdataOptions : The GetPITRdata options.
type GetPITRdataOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPITRdataOptions : Instantiate GetPITRdataOptions
func (*IbmCloudDatabasesV5) NewGetPITRdataOptions(id string) *GetPITRdataOptions {
	return &GetPITRdataOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetPITRdataOptions) SetID(id string) *GetPITRdataOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPITRdataOptions) SetHeaders(param map[string]string) *GetPITRdataOptions {
	options.Headers = param
	return options
}

// GetRegionsOptions : The GetRegions options.
type GetRegionsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRegionsOptions : Instantiate GetRegionsOptions
func (*IbmCloudDatabasesV5) NewGetRegionsOptions() *GetRegionsOptions {
	return &GetRegionsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetRegionsOptions) SetHeaders(param map[string]string) *GetRegionsOptions {
	options.Headers = param
	return options
}

// GetRegionsResponse : GetRegionsResponse struct
type GetRegionsResponse struct {
	// An array of region ids.
	Regions []string `json:"regions,omitempty"`
}


// UnmarshalGetRegionsResponse unmarshals an instance of GetRegionsResponse from the specified map of raw messages.
func UnmarshalGetRegionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRegionsResponse)
	err = core.UnmarshalPrimitive(m, "regions", &obj.Regions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRemotesOptions : The GetRemotes options.
type GetRemotesOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRemotesOptions : Instantiate GetRemotesOptions
func (*IbmCloudDatabasesV5) NewGetRemotesOptions(id string) *GetRemotesOptions {
	return &GetRemotesOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetRemotesOptions) SetID(id string) *GetRemotesOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRemotesOptions) SetHeaders(param map[string]string) *GetRemotesOptions {
	options.Headers = param
	return options
}

// GetRemotesResponse : GetRemotesResponse struct
type GetRemotesResponse struct {
	// Remotes.
	Remotes *Remotes `json:"remotes,omitempty"`
}


// UnmarshalGetRemotesResponse unmarshals an instance of GetRemotesResponse from the specified map of raw messages.
func UnmarshalGetRemotesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRemotesResponse)
	err = core.UnmarshalModel(m, "remotes", &obj.Remotes, UnmarshalRemotes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRemotesSchemaOptions : The GetRemotesSchema options.
type GetRemotesSchemaOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRemotesSchemaOptions : Instantiate GetRemotesSchemaOptions
func (*IbmCloudDatabasesV5) NewGetRemotesSchemaOptions(id string) *GetRemotesSchemaOptions {
	return &GetRemotesSchemaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetRemotesSchemaOptions) SetID(id string) *GetRemotesSchemaOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRemotesSchemaOptions) SetHeaders(param map[string]string) *GetRemotesSchemaOptions {
	options.Headers = param
	return options
}

// GetRemotesSchemaResponse : GetRemotesSchemaResponse struct
type GetRemotesSchemaResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalGetRemotesSchemaResponse unmarshals an instance of GetRemotesSchemaResponse from the specified map of raw messages.
func UnmarshalGetRemotesSchemaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRemotesSchemaResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTasksOptions : The GetTasks options.
type GetTasksOptions struct {
	// Task ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTasksOptions : Instantiate GetTasksOptions
func (*IbmCloudDatabasesV5) NewGetTasksOptions(id string) *GetTasksOptions {
	return &GetTasksOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetTasksOptions) SetID(id string) *GetTasksOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTasksOptions) SetHeaders(param map[string]string) *GetTasksOptions {
	options.Headers = param
	return options
}

// GetTasksResponse : GetTasksResponse struct
type GetTasksResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalGetTasksResponse unmarshals an instance of GetTasksResponse from the specified map of raw messages.
func UnmarshalGetTasksResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTasksResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetUserOptions : The GetUser options.
type GetUserOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// User ID.
	UserID *string `json:"user_id" validate:"required,ne="`

	// Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
	EndpointType *string `json:"endpoint_type" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetUserOptions.EndpointType property.
// Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
const (
	GetUserOptions_EndpointType_Private = "private"
	GetUserOptions_EndpointType_Public = "public"
)

// NewGetUserOptions : Instantiate GetUserOptions
func (*IbmCloudDatabasesV5) NewGetUserOptions(id string, userID string, endpointType string) *GetUserOptions {
	return &GetUserOptions{
		ID: core.StringPtr(id),
		UserID: core.StringPtr(userID),
		EndpointType: core.StringPtr(endpointType),
	}
}

// SetID : Allow user to set ID
func (options *GetUserOptions) SetID(id string) *GetUserOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetUserID : Allow user to set UserID
func (options *GetUserOptions) SetUserID(userID string) *GetUserOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetEndpointType : Allow user to set EndpointType
func (options *GetUserOptions) SetEndpointType(endpointType string) *GetUserOptions {
	options.EndpointType = core.StringPtr(endpointType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserOptions) SetHeaders(param map[string]string) *GetUserOptions {
	options.Headers = param
	return options
}

// GetWhitelistOptions : The GetWhitelist options.
type GetWhitelistOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetWhitelistOptions : Instantiate GetWhitelistOptions
func (*IbmCloudDatabasesV5) NewGetWhitelistOptions(id string) *GetWhitelistOptions {
	return &GetWhitelistOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetWhitelistOptions) SetID(id string) *GetWhitelistOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetWhitelistOptions) SetHeaders(param map[string]string) *GetWhitelistOptions {
	options.Headers = param
	return options
}

// Group : Group struct
type Group struct {
	// Id/name for group.
	ID *string `json:"id,omitempty"`

	// Number of entities in the group.
	Count *int64 `json:"count,omitempty"`

	Members *GroupMembers `json:"members,omitempty"`

	Memory *GroupMemory `json:"memory,omitempty"`

	Cpu *GroupCpu `json:"cpu,omitempty"`

	Disk *GroupDisk `json:"disk,omitempty"`
}


// UnmarshalGroup unmarshals an instance of Group from the specified map of raw messages.
func UnmarshalGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Group)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalGroupMembers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalGroupCpu)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalGroupDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupCpu : GroupCpu struct
type GroupCpu struct {
	// Units used for scaling cpu - count means the value is the number of the unit(s) available.
	Units *string `json:"units,omitempty"`

	// Number of allocated CPUs.
	AllocationCount *int64 `json:"allocation_count,omitempty"`

	// Minimum number of CPUs.
	MinimumCount *int64 `json:"minimum_count,omitempty"`

	// Maximum number of CPUs.
	MaximumCount *int64 `json:"maximum_count,omitempty"`

	// Step size CPUs can be adjusted.
	StepSizeCount *int64 `json:"step_size_count,omitempty"`

	// Is this group's CPU count adjustable.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this group's CPU optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this group's CPU scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}


// UnmarshalGroupCpu unmarshals an instance of GroupCpu from the specified map of raw messages.
func UnmarshalGroupCpu(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupCpu)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_count", &obj.MinimumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_count", &obj.MaximumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_count", &obj.StepSizeCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupDisk : GroupDisk struct
type GroupDisk struct {
	// Units used for scaling storage.
	Units *string `json:"units,omitempty"`

	// Allocated storage in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`

	// Minimum allocated storage.
	MinimumMb *int64 `json:"minimum_mb,omitempty"`

	// Maximum allocated storage.
	MaximumMb *int64 `json:"maximum_mb,omitempty"`

	// Step size storage can be adjusted.
	StepSizeMb *int64 `json:"step_size_mb,omitempty"`

	// Is this group's storage adjustable?.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this group's storage optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this group's storage scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}


// UnmarshalGroupDisk unmarshals an instance of GroupDisk from the specified map of raw messages.
func UnmarshalGroupDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupDisk)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_mb", &obj.MinimumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_mb", &obj.MaximumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_mb", &obj.StepSizeMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupMembers : GroupMembers struct
type GroupMembers struct {
	// Units used for scaling number of members.
	Units *string `json:"units,omitempty"`

	// Allocated number of members.
	AllocationCount *int64 `json:"allocation_count,omitempty"`

	// Minimum number of members.
	MinimumCount *int64 `json:"minimum_count,omitempty"`

	// Maximum number of members.
	MaximumCount *int64 `json:"maximum_count,omitempty"`

	// Step size for number of members.
	StepSizeCount *int64 `json:"step_size_count,omitempty"`

	// Is this deployment's number of members adjustable?.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this deployments's number of members optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this deployment's number of members scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}


// UnmarshalGroupMembers unmarshals an instance of GroupMembers from the specified map of raw messages.
func UnmarshalGroupMembers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupMembers)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_count", &obj.MinimumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_count", &obj.MaximumCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_count", &obj.StepSizeCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupMemory : GroupMemory struct
type GroupMemory struct {
	// Units used for scaling memory.
	Units *string `json:"units,omitempty"`

	// Allocated memory in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`

	// Minimum memory in MB.
	MinimumMb *int64 `json:"minimum_mb,omitempty"`

	// Maximum memory in MB.
	MaximumMb *int64 `json:"maximum_mb,omitempty"`

	// Step size memory can be adjusted by in MB.
	StepSizeMb *int64 `json:"step_size_mb,omitempty"`

	// Is this group's memory adjustable?.
	IsAdjustable *bool `json:"is_adjustable,omitempty"`

	// Is this group's memory optional?.
	IsOptional *bool `json:"is_optional,omitempty"`

	// Can this group's memory scale down?.
	CanScaleDown *bool `json:"can_scale_down,omitempty"`
}


// UnmarshalGroupMemory unmarshals an instance of GroupMemory from the specified map of raw messages.
func UnmarshalGroupMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupMemory)
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_mb", &obj.MinimumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_mb", &obj.MaximumMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step_size_mb", &obj.StepSizeMb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_adjustable", &obj.IsAdjustable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_optional", &obj.IsOptional)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_scale_down", &obj.CanScaleDown)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Groups : Groups struct
type Groups struct {
	Groups []Group `json:"groups,omitempty"`
}


// UnmarshalGroups unmarshals an instance of Groups from the specified map of raw messages.
func UnmarshalGroups(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Groups)
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IntegerPropertySchema : Integer Property Schema.
type IntegerPropertySchema struct {
	// Whether the setting is customer-configurable.
	CustomerConfigurable *bool `json:"customer_configurable,omitempty"`

	// The default value of the setting.
	Default *int64 `json:"default,omitempty"`

	// The description of the default value.
	DefaultDescription *string `json:"default_description,omitempty"`

	// The description of the setting.
	Description *string `json:"description,omitempty"`

	// The type of this setting (e.g., string, integer).
	Kind *string `json:"kind,omitempty"`

	// Whether or not changing this setting will restart the database.
	RequiresRestart *bool `json:"requires_restart,omitempty"`

	// The minimum value that this setting accepts.
	Min *int64 `json:"min,omitempty"`

	// The maximum value that this setting accepts.
	Max *int64 `json:"max,omitempty"`

	// The number that should be skipped between each step of a slider rendered for this setting.
	Step *int64 `json:"step,omitempty"`
}


// UnmarshalIntegerPropertySchema unmarshals an instance of IntegerPropertySchema from the specified map of raw messages.
func UnmarshalIntegerPropertySchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntegerPropertySchema)
	err = core.UnmarshalPrimitive(m, "customer_configurable", &obj.CustomerConfigurable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_description", &obj.DefaultDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind", &obj.Kind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requires_restart", &obj.RequiresRestart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min", &obj.Min)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max", &obj.Max)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "step", &obj.Step)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KillConnectionsOptions : The KillConnections options.
type KillConnectionsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewKillConnectionsOptions : Instantiate KillConnectionsOptions
func (*IbmCloudDatabasesV5) NewKillConnectionsOptions(id string) *KillConnectionsOptions {
	return &KillConnectionsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *KillConnectionsOptions) SetID(id string) *KillConnectionsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *KillConnectionsOptions) SetHeaders(param map[string]string) *KillConnectionsOptions {
	options.Headers = param
	return options
}

// KillConnectionsResponse : KillConnectionsResponse struct
type KillConnectionsResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalKillConnectionsResponse unmarshals an instance of KillConnectionsResponse from the specified map of raw messages.
func UnmarshalKillConnectionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KillConnectionsResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogicalReplicationSlotLogicalReplicationSlot : LogicalReplicationSlotLogicalReplicationSlot struct
type LogicalReplicationSlotLogicalReplicationSlot struct {
	// name of the replication slot.
	Name *string `json:"name,omitempty"`

	// name of the database the replication slot is created on.
	DatabaseName *string `json:"database_name,omitempty"`

	// creating a replication slot is only supported for use with wal2json.
	PluginType *string `json:"plugin_type,omitempty"`
}


// UnmarshalLogicalReplicationSlotLogicalReplicationSlot unmarshals an instance of LogicalReplicationSlotLogicalReplicationSlot from the specified map of raw messages.
func UnmarshalLogicalReplicationSlotLogicalReplicationSlot(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogicalReplicationSlotLogicalReplicationSlot)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plugin_type", &obj.PluginType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MongoDBConnectionURI : MongoDBConnectionURI struct
type MongoDBConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []MongoDBConnectionURIHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *MongoDBConnectionURIAuthentication `json:"authentication,omitempty"`

	Certificate *MongoDBConnectionURICertificate `json:"certificate,omitempty"`

	// Name of the database to use in the URI connection.
	Database *string `json:"database,omitempty"`

	// Name of the replica set to use in the URI connection.
	ReplicaSet *string `json:"replica_set,omitempty"`
}


// UnmarshalMongoDBConnectionURI unmarshals an instance of MongoDBConnectionURI from the specified map of raw messages.
func UnmarshalMongoDBConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MongoDBConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalMongoDBConnectionURIHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalMongoDBConnectionURIAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalMongoDBConnectionURICertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replica_set", &obj.ReplicaSet)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MongoDBConnectionURIAuthentication : MongoDBConnectionURIAuthentication struct
type MongoDBConnectionURIAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalMongoDBConnectionURIAuthentication unmarshals an instance of MongoDBConnectionURIAuthentication from the specified map of raw messages.
func UnmarshalMongoDBConnectionURIAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MongoDBConnectionURIAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MongoDBConnectionURICertificate : MongoDBConnectionURICertificate struct
type MongoDBConnectionURICertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalMongoDBConnectionURICertificate unmarshals an instance of MongoDBConnectionURICertificate from the specified map of raw messages.
func UnmarshalMongoDBConnectionURICertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MongoDBConnectionURICertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MongoDBConnectionURIHostsItem : MongoDBConnectionURIHostsItem struct
type MongoDBConnectionURIHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalMongoDBConnectionURIHostsItem unmarshals an instance of MongoDBConnectionURIHostsItem from the specified map of raw messages.
func UnmarshalMongoDBConnectionURIHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MongoDBConnectionURIHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PointInTimeRecoveryData : PointInTimeRecoveryData struct
type PointInTimeRecoveryData struct {
	EarliestPointInTimeRecoveryTime *string `json:"earliest_point_in_time_recovery_time,omitempty"`
}


// UnmarshalPointInTimeRecoveryData unmarshals an instance of PointInTimeRecoveryData from the specified map of raw messages.
func UnmarshalPointInTimeRecoveryData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PointInTimeRecoveryData)
	err = core.UnmarshalPrimitive(m, "earliest_point_in_time_recovery_time", &obj.EarliestPointInTimeRecoveryTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostgreSQLConnectionURI : PostgreSQLConnectionURI struct
type PostgreSQLConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []PostgreSQLConnectionURIHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *PostgreSQLConnectionURIAuthentication `json:"authentication,omitempty"`

	Certificate *PostgreSQLConnectionURICertificate `json:"certificate,omitempty"`

	// Name of the database to use in the URI connection.
	Database *string `json:"database,omitempty"`
}


// UnmarshalPostgreSQLConnectionURI unmarshals an instance of PostgreSQLConnectionURI from the specified map of raw messages.
func UnmarshalPostgreSQLConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostgreSQLConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalPostgreSQLConnectionURIHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalPostgreSQLConnectionURIAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalPostgreSQLConnectionURICertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostgreSQLConnectionURIAuthentication : PostgreSQLConnectionURIAuthentication struct
type PostgreSQLConnectionURIAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalPostgreSQLConnectionURIAuthentication unmarshals an instance of PostgreSQLConnectionURIAuthentication from the specified map of raw messages.
func UnmarshalPostgreSQLConnectionURIAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostgreSQLConnectionURIAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostgreSQLConnectionURICertificate : PostgreSQLConnectionURICertificate struct
type PostgreSQLConnectionURICertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalPostgreSQLConnectionURICertificate unmarshals an instance of PostgreSQLConnectionURICertificate from the specified map of raw messages.
func UnmarshalPostgreSQLConnectionURICertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostgreSQLConnectionURICertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostgreSQLConnectionURIHostsItem : PostgreSQLConnectionURIHostsItem struct
type PostgreSQLConnectionURIHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalPostgreSQLConnectionURIHostsItem unmarshals an instance of PostgreSQLConnectionURIHostsItem from the specified map of raw messages.
func UnmarshalPostgreSQLConnectionURIHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostgreSQLConnectionURIHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionAMQPS : RabbitMQConnectionAMQPS struct
type RabbitMQConnectionAMQPS struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []RabbitMQConnectionAMQPSHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *RabbitMQConnectionAMQPSAuthentication `json:"authentication,omitempty"`

	Certificate *RabbitMQConnectionAMQPSCertificate `json:"certificate,omitempty"`
}


// UnmarshalRabbitMQConnectionAMQPS unmarshals an instance of RabbitMQConnectionAMQPS from the specified map of raw messages.
func UnmarshalRabbitMQConnectionAMQPS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionAMQPS)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalRabbitMQConnectionAMQPSHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalRabbitMQConnectionAMQPSAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalRabbitMQConnectionAMQPSCertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionAMQPSAuthentication : RabbitMQConnectionAMQPSAuthentication struct
type RabbitMQConnectionAMQPSAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalRabbitMQConnectionAMQPSAuthentication unmarshals an instance of RabbitMQConnectionAMQPSAuthentication from the specified map of raw messages.
func UnmarshalRabbitMQConnectionAMQPSAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionAMQPSAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionAMQPSCertificate : RabbitMQConnectionAMQPSCertificate struct
type RabbitMQConnectionAMQPSCertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalRabbitMQConnectionAMQPSCertificate unmarshals an instance of RabbitMQConnectionAMQPSCertificate from the specified map of raw messages.
func UnmarshalRabbitMQConnectionAMQPSCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionAMQPSCertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionAMQPSHostsItem : RabbitMQConnectionAMQPSHostsItem struct
type RabbitMQConnectionAMQPSHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalRabbitMQConnectionAMQPSHostsItem unmarshals an instance of RabbitMQConnectionAMQPSHostsItem from the specified map of raw messages.
func UnmarshalRabbitMQConnectionAMQPSHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionAMQPSHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionHTTPS : RabbitMQConnectionHTTPS struct
type RabbitMQConnectionHTTPS struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []RabbitMQConnectionHTTPSHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *RabbitMQConnectionHTTPSAuthentication `json:"authentication,omitempty"`

	Certificate *RabbitMQConnectionHTTPSCertificate `json:"certificate,omitempty"`

	// Indicates the address is accessible by browser, for the RabbitMQ Management UI.
	BrowserAccessible *bool `json:"browser_accessible,omitempty"`
}


// UnmarshalRabbitMQConnectionHTTPS unmarshals an instance of RabbitMQConnectionHTTPS from the specified map of raw messages.
func UnmarshalRabbitMQConnectionHTTPS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionHTTPS)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalRabbitMQConnectionHTTPSHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalRabbitMQConnectionHTTPSAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalRabbitMQConnectionHTTPSCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "browser_accessible", &obj.BrowserAccessible)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionHTTPSAuthentication : RabbitMQConnectionHTTPSAuthentication struct
type RabbitMQConnectionHTTPSAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalRabbitMQConnectionHTTPSAuthentication unmarshals an instance of RabbitMQConnectionHTTPSAuthentication from the specified map of raw messages.
func UnmarshalRabbitMQConnectionHTTPSAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionHTTPSAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionHTTPSCertificate : RabbitMQConnectionHTTPSCertificate struct
type RabbitMQConnectionHTTPSCertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalRabbitMQConnectionHTTPSCertificate unmarshals an instance of RabbitMQConnectionHTTPSCertificate from the specified map of raw messages.
func UnmarshalRabbitMQConnectionHTTPSCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionHTTPSCertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionHTTPSHostsItem : RabbitMQConnectionHTTPSHostsItem struct
type RabbitMQConnectionHTTPSHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalRabbitMQConnectionHTTPSHostsItem unmarshals an instance of RabbitMQConnectionHTTPSHostsItem from the specified map of raw messages.
func UnmarshalRabbitMQConnectionHTTPSHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionHTTPSHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionMQTTS : RabbitMQConnectionMQTTS struct
type RabbitMQConnectionMQTTS struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []RabbitMQConnectionMQTTSHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *RabbitMQConnectionMQTTSAuthentication `json:"authentication,omitempty"`

	Certificate *RabbitMQConnectionMQTTSCertificate `json:"certificate,omitempty"`
}


// UnmarshalRabbitMQConnectionMQTTS unmarshals an instance of RabbitMQConnectionMQTTS from the specified map of raw messages.
func UnmarshalRabbitMQConnectionMQTTS(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionMQTTS)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalRabbitMQConnectionMQTTSHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalRabbitMQConnectionMQTTSAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalRabbitMQConnectionMQTTSCertificate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionMQTTSAuthentication : RabbitMQConnectionMQTTSAuthentication struct
type RabbitMQConnectionMQTTSAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalRabbitMQConnectionMQTTSAuthentication unmarshals an instance of RabbitMQConnectionMQTTSAuthentication from the specified map of raw messages.
func UnmarshalRabbitMQConnectionMQTTSAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionMQTTSAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionMQTTSCertificate : RabbitMQConnectionMQTTSCertificate struct
type RabbitMQConnectionMQTTSCertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalRabbitMQConnectionMQTTSCertificate unmarshals an instance of RabbitMQConnectionMQTTSCertificate from the specified map of raw messages.
func UnmarshalRabbitMQConnectionMQTTSCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionMQTTSCertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionMQTTSHostsItem : RabbitMQConnectionMQTTSHostsItem struct
type RabbitMQConnectionMQTTSHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalRabbitMQConnectionMQTTSHostsItem unmarshals an instance of RabbitMQConnectionMQTTSHostsItem from the specified map of raw messages.
func UnmarshalRabbitMQConnectionMQTTSHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionMQTTSHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionStompSSL : RabbitMQConnectionStompSSL struct
type RabbitMQConnectionStompSSL struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	Hosts []RabbitMQConnectionStompSSLHostsItem `json:"hosts,omitempty"`

	Authentication *RabbitMQConnectionStompSSLAuthentication `json:"authentication,omitempty"`

	Certificate *RabbitMQConnectionStompSSLCertificate `json:"certificate,omitempty"`

	// Indicates ssl is required for the connection.
	Ssl *bool `json:"ssl,omitempty"`
}


// UnmarshalRabbitMQConnectionStompSSL unmarshals an instance of RabbitMQConnectionStompSSL from the specified map of raw messages.
func UnmarshalRabbitMQConnectionStompSSL(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionStompSSL)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalRabbitMQConnectionStompSSLHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalRabbitMQConnectionStompSSLAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalRabbitMQConnectionStompSSLCertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionStompSSLAuthentication : RabbitMQConnectionStompSSLAuthentication struct
type RabbitMQConnectionStompSSLAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalRabbitMQConnectionStompSSLAuthentication unmarshals an instance of RabbitMQConnectionStompSSLAuthentication from the specified map of raw messages.
func UnmarshalRabbitMQConnectionStompSSLAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionStompSSLAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionStompSSLCertificate : RabbitMQConnectionStompSSLCertificate struct
type RabbitMQConnectionStompSSLCertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalRabbitMQConnectionStompSSLCertificate unmarshals an instance of RabbitMQConnectionStompSSLCertificate from the specified map of raw messages.
func UnmarshalRabbitMQConnectionStompSSLCertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionStompSSLCertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RabbitMQConnectionStompSSLHostsItem : RabbitMQConnectionStompSSLHostsItem struct
type RabbitMQConnectionStompSSLHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalRabbitMQConnectionStompSSLHostsItem unmarshals an instance of RabbitMQConnectionStompSSLHostsItem from the specified map of raw messages.
func UnmarshalRabbitMQConnectionStompSSLHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RabbitMQConnectionStompSSLHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedisConnectionURI : RedisConnectionURI struct
type RedisConnectionURI struct {
	// Type of connection being described.
	Type *string `json:"type,omitempty"`

	Composed []string `json:"composed,omitempty"`

	// Scheme/protocol for URI connection.
	Scheme *string `json:"scheme,omitempty"`

	Hosts []RedisConnectionURIHostsItem `json:"hosts,omitempty"`

	// Path for URI connection.
	Path *string `json:"path,omitempty"`

	// Query options to add to the URI connection.
	QueryOptions interface{} `json:"query_options,omitempty"`

	Authentication *RedisConnectionURIAuthentication `json:"authentication,omitempty"`

	Certificate *RedisConnectionURICertificate `json:"certificate,omitempty"`

	// Number of the database to use in the URI connection.
	Database *int64 `json:"database,omitempty"`
}


// UnmarshalRedisConnectionURI unmarshals an instance of RedisConnectionURI from the specified map of raw messages.
func UnmarshalRedisConnectionURI(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedisConnectionURI)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composed", &obj.Composed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scheme", &obj.Scheme)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hosts", &obj.Hosts, UnmarshalRedisConnectionURIHostsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "query_options", &obj.QueryOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authentication", &obj.Authentication, UnmarshalRedisConnectionURIAuthentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "certificate", &obj.Certificate, UnmarshalRedisConnectionURICertificate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database", &obj.Database)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedisConnectionURIAuthentication : RedisConnectionURIAuthentication struct
type RedisConnectionURIAuthentication struct {
	// Authentication method for this credential.
	Method *string `json:"method,omitempty"`

	// Username part of credential.
	Username *string `json:"username,omitempty"`

	// Password part of credential.
	Password *string `json:"password,omitempty"`
}


// UnmarshalRedisConnectionURIAuthentication unmarshals an instance of RedisConnectionURIAuthentication from the specified map of raw messages.
func UnmarshalRedisConnectionURIAuthentication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedisConnectionURIAuthentication)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedisConnectionURICertificate : RedisConnectionURICertificate struct
type RedisConnectionURICertificate struct {
	// Name associated with the certificate.
	Name *string `json:"name,omitempty"`

	// Base64 encoded version of the certificate.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`
}


// UnmarshalRedisConnectionURICertificate unmarshals an instance of RedisConnectionURICertificate from the specified map of raw messages.
func UnmarshalRedisConnectionURICertificate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedisConnectionURICertificate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedisConnectionURIHostsItem : RedisConnectionURIHostsItem struct
type RedisConnectionURIHostsItem struct {
	// Hostname for connection.
	Hostname *string `json:"hostname,omitempty"`

	// Port number for connection.
	Port *int64 `json:"port,omitempty"`
}


// UnmarshalRedisConnectionURIHostsItem unmarshals an instance of RedisConnectionURIHostsItem from the specified map of raw messages.
func UnmarshalRedisConnectionURIHostsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedisConnectionURIHostsItem)
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Remotes : Remotes.
type Remotes struct {
	// Leader ID, if applicable.
	Leader *string `json:"leader,omitempty"`

	// Replica IDs, if applicable.
	Replicas []string `json:"replicas,omitempty"`
}


// UnmarshalRemotes unmarshals an instance of Remotes from the specified map of raw messages.
func UnmarshalRemotes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Remotes)
	err = core.UnmarshalPrimitive(m, "leader", &obj.Leader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replicas", &obj.Replicas)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceWhitelistOptions : The ReplaceWhitelist options.
type ReplaceWhitelistOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// An array of allowlist entries.
	IpAddresses []WhitelistEntry `json:"ip_addresses,omitempty"`

	// Verify that the current allowlist matches a provided ETag value. Use in conjunction with the GET operation's ETag
	// header to ensure synchronicity between clients.
	IfMatch *string `json:"If-Match,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceWhitelistOptions : Instantiate ReplaceWhitelistOptions
func (*IbmCloudDatabasesV5) NewReplaceWhitelistOptions(id string) *ReplaceWhitelistOptions {
	return &ReplaceWhitelistOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *ReplaceWhitelistOptions) SetID(id string) *ReplaceWhitelistOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIpAddresses : Allow user to set IpAddresses
func (options *ReplaceWhitelistOptions) SetIpAddresses(ipAddresses []WhitelistEntry) *ReplaceWhitelistOptions {
	options.IpAddresses = ipAddresses
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *ReplaceWhitelistOptions) SetIfMatch(ifMatch string) *ReplaceWhitelistOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceWhitelistOptions) SetHeaders(param map[string]string) *ReplaceWhitelistOptions {
	options.Headers = param
	return options
}

// ReplaceWhitelistResponse : ReplaceWhitelistResponse struct
type ReplaceWhitelistResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalReplaceWhitelistResponse unmarshals an instance of ReplaceWhitelistResponse from the specified map of raw messages.
func UnmarshalReplaceWhitelistResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplaceWhitelistResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetAutoscalingConditionsOptions : The SetAutoscalingConditions options.
type SetAutoscalingConditionsOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Group ID.
	GroupID *string `json:"group_id" validate:"required,ne="`

	Autoscaling *AutoscalingSetGroup `json:"autoscaling,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetAutoscalingConditionsOptions : Instantiate SetAutoscalingConditionsOptions
func (*IbmCloudDatabasesV5) NewSetAutoscalingConditionsOptions(id string, groupID string) *SetAutoscalingConditionsOptions {
	return &SetAutoscalingConditionsOptions{
		ID: core.StringPtr(id),
		GroupID: core.StringPtr(groupID),
	}
}

// SetID : Allow user to set ID
func (options *SetAutoscalingConditionsOptions) SetID(id string) *SetAutoscalingConditionsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetGroupID : Allow user to set GroupID
func (options *SetAutoscalingConditionsOptions) SetGroupID(groupID string) *SetAutoscalingConditionsOptions {
	options.GroupID = core.StringPtr(groupID)
	return options
}

// SetAutoscaling : Allow user to set Autoscaling
func (options *SetAutoscalingConditionsOptions) SetAutoscaling(autoscaling *AutoscalingSetGroup) *SetAutoscalingConditionsOptions {
	options.Autoscaling = autoscaling
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetAutoscalingConditionsOptions) SetHeaders(param map[string]string) *SetAutoscalingConditionsOptions {
	options.Headers = param
	return options
}

// SetAutoscalingConditionsResponse : SetAutoscalingConditionsResponse struct
type SetAutoscalingConditionsResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalSetAutoscalingConditionsResponse unmarshals an instance of SetAutoscalingConditionsResponse from the specified map of raw messages.
func UnmarshalSetAutoscalingConditionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetAutoscalingConditionsResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetCPUGroupCPU : SetCPUGroupCPU struct
type SetCPUGroupCPU struct {
	// Number of allocated CPUs.
	AllocationCount *int64 `json:"allocation_count,omitempty"`
}


// UnmarshalSetCPUGroupCPU unmarshals an instance of SetCPUGroupCPU from the specified map of raw messages.
func UnmarshalSetCPUGroupCPU(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetCPUGroupCPU)
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetConfigurationConfiguration : SetConfigurationConfiguration struct
// Models which "extend" this model:
// - SetConfigurationConfigurationPGConfiguration
// - SetConfigurationConfigurationRedisConfiguration
type SetConfigurationConfiguration struct {
	// Maximum connections allowed.
	MaxConnections *int64 `json:"max_connections,omitempty"`

	// Max number of transactions that can be in the "prepared" state simultaneously.
	MaxPreparedTransactions *int64 `json:"max_prepared_transactions,omitempty"`

	// Deadlock timeout in ms. The time to wait on a lock before checking for deadlock.  Also the duration where lock waits
	// will be logged.
	DeadlockTimeout *int64 `json:"deadlock_timeout,omitempty"`

	// Number of simultaneous requests that can be handled efficiently by the disk subsystem.
	EffectiveIoConcurrency *int64 `json:"effective_io_concurrency,omitempty"`

	// Maximum number of simultaneously defined replication slots.
	MaxReplicationSlots *int64 `json:"max_replication_slots,omitempty"`

	// Maximum number of simultaneously running WAL sender processes.
	MaxWalSenders *int64 `json:"max_wal_senders,omitempty"`

	// The number of 8kB shared memory buffers used by the server.  Set to 1/4 of memory.  Setting too high will cause
	// crashes or prevent the database from starting.
	SharedBuffers *int64 `json:"shared_buffers,omitempty"`

	// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
	// synchronous replication which will impact performance and availabilty.
	SynchronousCommit *string `json:"synchronous_commit,omitempty"`

	// WAL level.  Set to logical to use logical decoding or logical replication.
	WalLevel *string `json:"wal_level,omitempty"`

	// The number of seconds to wait before forces a switch to the next WAL file if a new file has not been started.
	ArchiveTimeout *int64 `json:"archive_timeout,omitempty"`

	// The minimum number of milliseconds for execution time above which statements will be logged.
	LogMinDurationStatement *int64 `json:"log_min_duration_statement,omitempty"`

	// The maximum memory Redis should use, as bytes.
	MaxmemoryRedis *int64 `json:"maxmemory-redis,omitempty"`

	// The policy with which Redis evicts keys when maximum memory is reached.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// If set to yes this will enable AOF persistence.
	Appendonly *string `json:"appendonly,omitempty"`

	// The maximum memory Redis should use, as bytes.
	MaxmemorySamples *int64 `json:"maxmemory-samples,omitempty"`

	// Whether or not to stop accepting writes when background persistence actions fail.
	StopWritesOnBgsaveError *string `json:"stop-writes-on-bgsave-error,omitempty"`
}

// Constants associated with the SetConfigurationConfiguration.SynchronousCommit property.
// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
// synchronous replication which will impact performance and availabilty.
const (
	SetConfigurationConfiguration_SynchronousCommit_Local = "local"
	SetConfigurationConfiguration_SynchronousCommit_Off = "off"
)

// Constants associated with the SetConfigurationConfiguration.WalLevel property.
// WAL level.  Set to logical to use logical decoding or logical replication.
const (
	SetConfigurationConfiguration_WalLevel_HotStandby = "hot_standby"
	SetConfigurationConfiguration_WalLevel_Logical = "logical"
)

// Constants associated with the SetConfigurationConfiguration.MaxmemoryPolicy property.
// The policy with which Redis evicts keys when maximum memory is reached.
const (
	SetConfigurationConfiguration_MaxmemoryPolicy_AllkeysLru = "allkeys-lru"
	SetConfigurationConfiguration_MaxmemoryPolicy_AllkeysRandom = "allkeys-random"
	SetConfigurationConfiguration_MaxmemoryPolicy_Noeviction = "noeviction"
	SetConfigurationConfiguration_MaxmemoryPolicy_VolatileLru = "volatile-lru"
	SetConfigurationConfiguration_MaxmemoryPolicy_VolatileRandom = "volatile-random"
	SetConfigurationConfiguration_MaxmemoryPolicy_VolatileTTL = "volatile-ttl"
)

// Constants associated with the SetConfigurationConfiguration.Appendonly property.
// If set to yes this will enable AOF persistence.
const (
	SetConfigurationConfiguration_Appendonly_No = "no"
	SetConfigurationConfiguration_Appendonly_Yes = "yes"
)

// Constants associated with the SetConfigurationConfiguration.StopWritesOnBgsaveError property.
// Whether or not to stop accepting writes when background persistence actions fail.
const (
	SetConfigurationConfiguration_StopWritesOnBgsaveError_No = "no"
	SetConfigurationConfiguration_StopWritesOnBgsaveError_Yes = "yes"
)

func (*SetConfigurationConfiguration) isaSetConfigurationConfiguration() bool {
	return true
}

type SetConfigurationConfigurationIntf interface {
	isaSetConfigurationConfiguration() bool
}

// UnmarshalSetConfigurationConfiguration unmarshals an instance of SetConfigurationConfiguration from the specified map of raw messages.
func UnmarshalSetConfigurationConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetConfigurationConfiguration)
	err = core.UnmarshalPrimitive(m, "max_connections", &obj.MaxConnections)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_prepared_transactions", &obj.MaxPreparedTransactions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deadlock_timeout", &obj.DeadlockTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_io_concurrency", &obj.EffectiveIoConcurrency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_replication_slots", &obj.MaxReplicationSlots)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_wal_senders", &obj.MaxWalSenders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shared_buffers", &obj.SharedBuffers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "synchronous_commit", &obj.SynchronousCommit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wal_level", &obj.WalLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "archive_timeout", &obj.ArchiveTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_min_duration_statement", &obj.LogMinDurationStatement)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-redis", &obj.MaxmemoryRedis)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-policy", &obj.MaxmemoryPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "appendonly", &obj.Appendonly)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-samples", &obj.MaxmemorySamples)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stop-writes-on-bgsave-error", &obj.StopWritesOnBgsaveError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDatabaseConfigurationOptions : The SetDatabaseConfiguration options.
type SetDatabaseConfigurationOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Configuration SetConfigurationConfigurationIntf `json:"configuration" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetDatabaseConfigurationOptions : Instantiate SetDatabaseConfigurationOptions
func (*IbmCloudDatabasesV5) NewSetDatabaseConfigurationOptions(id string, configuration SetConfigurationConfigurationIntf) *SetDatabaseConfigurationOptions {
	return &SetDatabaseConfigurationOptions{
		ID: core.StringPtr(id),
		Configuration: configuration,
	}
}

// SetID : Allow user to set ID
func (options *SetDatabaseConfigurationOptions) SetID(id string) *SetDatabaseConfigurationOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetConfiguration : Allow user to set Configuration
func (options *SetDatabaseConfigurationOptions) SetConfiguration(configuration SetConfigurationConfigurationIntf) *SetDatabaseConfigurationOptions {
	options.Configuration = configuration
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetDatabaseConfigurationOptions) SetHeaders(param map[string]string) *SetDatabaseConfigurationOptions {
	options.Headers = param
	return options
}

// SetDatabaseConfigurationResponse : SetDatabaseConfigurationResponse struct
type SetDatabaseConfigurationResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalSetDatabaseConfigurationResponse unmarshals an instance of SetDatabaseConfigurationResponse from the specified map of raw messages.
func UnmarshalSetDatabaseConfigurationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDatabaseConfigurationResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupOptions : The SetDeploymentScalingGroup options.
type SetDeploymentScalingGroupOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Group Id.
	GroupID *string `json:"group_id" validate:"required,ne="`

	// Scaling group settings.
	SetDeploymentScalingGroupRequest SetDeploymentScalingGroupRequestIntf `json:"SetDeploymentScalingGroupRequest" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetDeploymentScalingGroupOptions : Instantiate SetDeploymentScalingGroupOptions
func (*IbmCloudDatabasesV5) NewSetDeploymentScalingGroupOptions(id string, groupID string, setDeploymentScalingGroupRequest SetDeploymentScalingGroupRequestIntf) *SetDeploymentScalingGroupOptions {
	return &SetDeploymentScalingGroupOptions{
		ID: core.StringPtr(id),
		GroupID: core.StringPtr(groupID),
		SetDeploymentScalingGroupRequest: setDeploymentScalingGroupRequest,
	}
}

// SetID : Allow user to set ID
func (options *SetDeploymentScalingGroupOptions) SetID(id string) *SetDeploymentScalingGroupOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetGroupID : Allow user to set GroupID
func (options *SetDeploymentScalingGroupOptions) SetGroupID(groupID string) *SetDeploymentScalingGroupOptions {
	options.GroupID = core.StringPtr(groupID)
	return options
}

// SetSetDeploymentScalingGroupRequest : Allow user to set SetDeploymentScalingGroupRequest
func (options *SetDeploymentScalingGroupOptions) SetSetDeploymentScalingGroupRequest(setDeploymentScalingGroupRequest SetDeploymentScalingGroupRequestIntf) *SetDeploymentScalingGroupOptions {
	options.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequest
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetDeploymentScalingGroupOptions) SetHeaders(param map[string]string) *SetDeploymentScalingGroupOptions {
	options.Headers = param
	return options
}

// SetDeploymentScalingGroupRequest : SetDeploymentScalingGroupRequest struct
// Models which "extend" this model:
// - SetDeploymentScalingGroupRequestSetMembersGroup
// - SetDeploymentScalingGroupRequestSetMemoryGroup
// - SetDeploymentScalingGroupRequestSetCPUGroup
// - SetDeploymentScalingGroupRequestSetDiskGroup
type SetDeploymentScalingGroupRequest struct {
	Members *SetMembersGroupMembers `json:"members,omitempty"`

	Memory *SetMemoryGroupMemory `json:"memory,omitempty"`

	Cpu *SetCPUGroupCPU `json:"cpu,omitempty"`

	Disk *SetDiskGroupDisk `json:"disk,omitempty"`
}

func (*SetDeploymentScalingGroupRequest) isaSetDeploymentScalingGroupRequest() bool {
	return true
}

type SetDeploymentScalingGroupRequestIntf interface {
	isaSetDeploymentScalingGroupRequest() bool
}

// UnmarshalSetDeploymentScalingGroupRequest unmarshals an instance of SetDeploymentScalingGroupRequest from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupRequest)
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalSetMembersGroupMembers)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalSetMemoryGroupMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalSetCPUGroupCPU)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalSetDiskGroupDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupResponse : SetDeploymentScalingGroupResponse struct
type SetDeploymentScalingGroupResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalSetDeploymentScalingGroupResponse unmarshals an instance of SetDeploymentScalingGroupResponse from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDiskGroupDisk : SetDiskGroupDisk struct
type SetDiskGroupDisk struct {
	// Allocated storage in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`
}


// UnmarshalSetDiskGroupDisk unmarshals an instance of SetDiskGroupDisk from the specified map of raw messages.
func UnmarshalSetDiskGroupDisk(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDiskGroupDisk)
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetMembersGroupMembers : SetMembersGroupMembers struct
type SetMembersGroupMembers struct {
	// Allocated number of members.
	AllocationCount *int64 `json:"allocation_count,omitempty"`
}


// UnmarshalSetMembersGroupMembers unmarshals an instance of SetMembersGroupMembers from the specified map of raw messages.
func UnmarshalSetMembersGroupMembers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetMembersGroupMembers)
	err = core.UnmarshalPrimitive(m, "allocation_count", &obj.AllocationCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetMemoryGroupMemory : SetMemoryGroupMemory struct
type SetMemoryGroupMemory struct {
	// Allocated memory in MB.
	AllocationMb *int64 `json:"allocation_mb,omitempty"`
}


// UnmarshalSetMemoryGroupMemory unmarshals an instance of SetMemoryGroupMemory from the specified map of raw messages.
func UnmarshalSetMemoryGroupMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetMemoryGroupMemory)
	err = core.UnmarshalPrimitive(m, "allocation_mb", &obj.AllocationMb)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetPromotionOptions : The SetPromotion options.
type SetPromotionOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Promotion SetPromotionPromotionIntf `json:"Promotion" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetPromotionOptions : Instantiate SetPromotionOptions
func (*IbmCloudDatabasesV5) NewSetPromotionOptions(id string, promotion SetPromotionPromotionIntf) *SetPromotionOptions {
	return &SetPromotionOptions{
		ID: core.StringPtr(id),
		Promotion: promotion,
	}
}

// SetID : Allow user to set ID
func (options *SetPromotionOptions) SetID(id string) *SetPromotionOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetPromotion : Allow user to set Promotion
func (options *SetPromotionOptions) SetPromotion(promotion SetPromotionPromotionIntf) *SetPromotionOptions {
	options.Promotion = promotion
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetPromotionOptions) SetHeaders(param map[string]string) *SetPromotionOptions {
	options.Headers = param
	return options
}

// SetPromotionPromotion : SetPromotionPromotion struct
// Models which "extend" this model:
// - SetPromotionPromotionPromote
// - SetPromotionPromotionUpgradePromote
type SetPromotionPromotion struct {
	// Promotion options.
	Promotion map[string]interface{} `json:"promotion,omitempty"`
}

func (*SetPromotionPromotion) isaSetPromotionPromotion() bool {
	return true
}

type SetPromotionPromotionIntf interface {
	isaSetPromotionPromotion() bool
}

// UnmarshalSetPromotionPromotion unmarshals an instance of SetPromotionPromotion from the specified map of raw messages.
func UnmarshalSetPromotionPromotion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetPromotionPromotion)
	err = core.UnmarshalPrimitive(m, "promotion", &obj.Promotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetPromotionResponse : SetPromotionResponse struct
type SetPromotionResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalSetPromotionResponse unmarshals an instance of SetPromotionResponse from the specified map of raw messages.
func UnmarshalSetPromotionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetPromotionResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetRemotesOptions : The SetRemotes options.
type SetRemotesOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	Remotes *SetRemotesRequestRemotes `json:"remotes,omitempty"`

	// Option to restore instance without taking a backup once data is restored. Allows restored deployment to be available
	// sooner.
	SkipInitialBackup *bool `json:"skip_initial_backup,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetRemotesOptions : Instantiate SetRemotesOptions
func (*IbmCloudDatabasesV5) NewSetRemotesOptions(id string) *SetRemotesOptions {
	return &SetRemotesOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *SetRemotesOptions) SetID(id string) *SetRemotesOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetRemotes : Allow user to set Remotes
func (options *SetRemotesOptions) SetRemotes(remotes *SetRemotesRequestRemotes) *SetRemotesOptions {
	options.Remotes = remotes
	return options
}

// SetSkipInitialBackup : Allow user to set SkipInitialBackup
func (options *SetRemotesOptions) SetSkipInitialBackup(skipInitialBackup bool) *SetRemotesOptions {
	options.SkipInitialBackup = core.BoolPtr(skipInitialBackup)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetRemotesOptions) SetHeaders(param map[string]string) *SetRemotesOptions {
	options.Headers = param
	return options
}

// SetRemotesRequestRemotes : SetRemotesRequestRemotes struct
type SetRemotesRequestRemotes struct {
	// Leader should be an empty string.
	Leader *string `json:"leader,omitempty"`
}


// UnmarshalSetRemotesRequestRemotes unmarshals an instance of SetRemotesRequestRemotes from the specified map of raw messages.
func UnmarshalSetRemotesRequestRemotes(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetRemotesRequestRemotes)
	err = core.UnmarshalPrimitive(m, "leader", &obj.Leader)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetRemotesResponse : SetRemotesResponse struct
type SetRemotesResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalSetRemotesResponse unmarshals an instance of SetRemotesResponse from the specified map of raw messages.
func UnmarshalSetRemotesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetRemotesResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StartOndemandBackupOptions : The StartOndemandBackup options.
type StartOndemandBackupOptions struct {
	// Deployment ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewStartOndemandBackupOptions : Instantiate StartOndemandBackupOptions
func (*IbmCloudDatabasesV5) NewStartOndemandBackupOptions(id string) *StartOndemandBackupOptions {
	return &StartOndemandBackupOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *StartOndemandBackupOptions) SetID(id string) *StartOndemandBackupOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *StartOndemandBackupOptions) SetHeaders(param map[string]string) *StartOndemandBackupOptions {
	options.Headers = param
	return options
}

// StartOndemandBackupResponse : StartOndemandBackupResponse struct
type StartOndemandBackupResponse struct {
	Task *Task `json:"task,omitempty"`
}


// UnmarshalStartOndemandBackupResponse unmarshals an instance of StartOndemandBackupResponse from the specified map of raw messages.
func UnmarshalStartOndemandBackupResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StartOndemandBackupResponse)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Task : Task struct
type Task struct {
	// ID of the task.
	ID *string `json:"id,omitempty"`

	// Human-readable description of the task.
	Description *string `json:"description,omitempty"`

	// The status of the task.
	Status *string `json:"status,omitempty"`

	// ID of the deployment the task is being performed on.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// Indicator as percentage of progress of the task.
	ProgressPercent *int64 `json:"progress_percent,omitempty"`

	// Date and time when the task was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// Constants associated with the Task.Status property.
// The status of the task.
const (
	Task_Status_Completed = "completed"
	Task_Status_Failed = "failed"
	Task_Status_Running = "running"
)


// UnmarshalTask unmarshals an instance of Task from the specified map of raw messages.
func UnmarshalTask(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Task)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "progress_percent", &obj.ProgressPercent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tasks : Tasks struct
type Tasks struct {
	Tasks []Task `json:"tasks,omitempty"`
}


// UnmarshalTasks unmarshals an instance of Tasks from the specified map of raw messages.
func UnmarshalTasks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tasks)
	err = core.UnmarshalModel(m, "tasks", &obj.Tasks, UnmarshalTask)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Whitelist : Whitelist struct
type Whitelist struct {
	// An array of allowlist entries.
	IpAddresses []WhitelistEntry `json:"ip_addresses,omitempty"`
}


// UnmarshalWhitelist unmarshals an instance of Whitelist from the specified map of raw messages.
func UnmarshalWhitelist(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Whitelist)
	err = core.UnmarshalModel(m, "ip_addresses", &obj.IpAddresses, UnmarshalWhitelistEntry)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WhitelistEntry : WhitelistEntry struct
type WhitelistEntry struct {
	// An IPv4 address or a CIDR range (netmasked IPv4 address).
	Address *string `json:"address,omitempty"`

	// A human readable description of the address or range for identification purposes.
	Description *string `json:"description,omitempty"`
}


// UnmarshalWhitelistEntry unmarshals an instance of WhitelistEntry from the specified map of raw messages.
func UnmarshalWhitelistEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WhitelistEntry)
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscalingAutoscalingCPUGroup : AutoscalingSetGroupAutoscalingAutoscalingCPUGroup struct
// This model "extends" AutoscalingSetGroupAutoscaling
type AutoscalingSetGroupAutoscalingAutoscalingCPUGroup struct {
	Cpu *AutoscalingCPUGroup `json:"cpu,omitempty"`
}


func (*AutoscalingSetGroupAutoscalingAutoscalingCPUGroup) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

// UnmarshalAutoscalingSetGroupAutoscalingAutoscalingCPUGroup unmarshals an instance of AutoscalingSetGroupAutoscalingAutoscalingCPUGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscalingAutoscalingCPUGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscalingAutoscalingCPUGroup)
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalAutoscalingCPUGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscalingAutoscalingDiskGroup : AutoscalingSetGroupAutoscalingAutoscalingDiskGroup struct
// This model "extends" AutoscalingSetGroupAutoscaling
type AutoscalingSetGroupAutoscalingAutoscalingDiskGroup struct {
	Disk *AutoscalingDiskGroupDisk `json:"disk,omitempty"`
}


func (*AutoscalingSetGroupAutoscalingAutoscalingDiskGroup) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

// UnmarshalAutoscalingSetGroupAutoscalingAutoscalingDiskGroup unmarshals an instance of AutoscalingSetGroupAutoscalingAutoscalingDiskGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscalingAutoscalingDiskGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscalingAutoscalingDiskGroup)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalAutoscalingDiskGroupDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup : AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup struct
// This model "extends" AutoscalingSetGroupAutoscaling
type AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup struct {
	Memory *AutoscalingMemoryGroupMemory `json:"memory,omitempty"`
}


func (*AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup) isaAutoscalingSetGroupAutoscaling() bool {
	return true
}

// UnmarshalAutoscalingSetGroupAutoscalingAutoscalingMemoryGroup unmarshals an instance of AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup from the specified map of raw messages.
func UnmarshalAutoscalingSetGroupAutoscalingAutoscalingMemoryGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AutoscalingSetGroupAutoscalingAutoscalingMemoryGroup)
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalAutoscalingMemoryGroupMemory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationSchemaSchemaPGConfigurationSchema : PostgreSQL and EnterpriseDB Configuration Schema.
// This model "extends" ConfigurationSchemaSchema
type ConfigurationSchemaSchemaPGConfigurationSchema struct {
	// Integer Property Schema.
	MaxConnections *IntegerPropertySchema `json:"max_connections" validate:"required"`

	// Integer Property Schema.
	MaxPreparedConnections *IntegerPropertySchema `json:"max_prepared_connections" validate:"required"`

	// Integer Property Schema.
	BackupRetentionPeriod *IntegerPropertySchema `json:"backup_retention_period" validate:"required"`

	// Integer Property Schema.
	DeadlockTimeout *IntegerPropertySchema `json:"deadlock_timeout" validate:"required"`

	// Integer Property Schema.
	EffectiveIoConcurrency *IntegerPropertySchema `json:"effective_io_concurrency" validate:"required"`

	// Integer Property Schema.
	MaxReplicationSlots *IntegerPropertySchema `json:"max_replication_slots" validate:"required"`

	// Integer Property Schema.
	MaxWalSenders *IntegerPropertySchema `json:"max_wal_senders" validate:"required"`

	// Integer Property Schema.
	SharedBuffers *IntegerPropertySchema `json:"shared_buffers" validate:"required"`

	// Choice Property Schema.
	SynchronousCommit *ChoicePropertySchema `json:"synchronous_commit" validate:"required"`

	// Choice Property Schema.
	WalLevel *ChoicePropertySchema `json:"wal_level" validate:"required"`

	// Integer Property Schema.
	ArchiveTimeout *IntegerPropertySchema `json:"archive_timeout" validate:"required"`

	// Integer Property Schema.
	LogMinDurationStatement *IntegerPropertySchema `json:"log_min_duration_statement" validate:"required"`
}


func (*ConfigurationSchemaSchemaPGConfigurationSchema) isaConfigurationSchemaSchema() bool {
	return true
}

// UnmarshalConfigurationSchemaSchemaPGConfigurationSchema unmarshals an instance of ConfigurationSchemaSchemaPGConfigurationSchema from the specified map of raw messages.
func UnmarshalConfigurationSchemaSchemaPGConfigurationSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationSchemaSchemaPGConfigurationSchema)
	err = core.UnmarshalModel(m, "max_connections", &obj.MaxConnections, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_prepared_connections", &obj.MaxPreparedConnections, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "backup_retention_period", &obj.BackupRetentionPeriod, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deadlock_timeout", &obj.DeadlockTimeout, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "effective_io_concurrency", &obj.EffectiveIoConcurrency, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_replication_slots", &obj.MaxReplicationSlots, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_wal_senders", &obj.MaxWalSenders, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "shared_buffers", &obj.SharedBuffers, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "synchronous_commit", &obj.SynchronousCommit, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "wal_level", &obj.WalLevel, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "archive_timeout", &obj.ArchiveTimeout, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "log_min_duration_statement", &obj.LogMinDurationStatement, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationSchemaSchemaRedisConfigurationSchema : Redis Configuration Schema.
// This model "extends" ConfigurationSchemaSchema
type ConfigurationSchemaSchemaRedisConfigurationSchema struct {
	// Integer Property Schema.
	MaxmemoryRedis *IntegerPropertySchema `json:"maxmemory-redis" validate:"required"`

	// Choice Property Schema.
	MaxmemoryPolicy *ChoicePropertySchema `json:"maxmemory-policy" validate:"required"`

	// Choice Property Schema.
	Appendonly *ChoicePropertySchema `json:"appendonly" validate:"required"`

	// Integer Property Schema.
	MaxmemorySamples *IntegerPropertySchema `json:"maxmemory-samples" validate:"required"`

	// Choice Property Schema.
	StopWritesOnBgsaveError *ChoicePropertySchema `json:"stop-writes-on-bgsave-error" validate:"required"`
}


func (*ConfigurationSchemaSchemaRedisConfigurationSchema) isaConfigurationSchemaSchema() bool {
	return true
}

// UnmarshalConfigurationSchemaSchemaRedisConfigurationSchema unmarshals an instance of ConfigurationSchemaSchemaRedisConfigurationSchema from the specified map of raw messages.
func UnmarshalConfigurationSchemaSchemaRedisConfigurationSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationSchemaSchemaRedisConfigurationSchema)
	err = core.UnmarshalModel(m, "maxmemory-redis", &obj.MaxmemoryRedis, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "maxmemory-policy", &obj.MaxmemoryPolicy, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "appendonly", &obj.Appendonly, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "maxmemory-samples", &obj.MaxmemorySamples, UnmarshalIntegerPropertySchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stop-writes-on-bgsave-error", &obj.StopWritesOnBgsaveError, UnmarshalChoicePropertySchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnectionElasticsearchConnection : Elasticsearch Connection Strings.
// This model "extends" ConnectionConnection
type ConnectionConnectionElasticsearchConnection struct {
	// Elasticsearch Connection information for drivers and libraries.
	Https *ElasticsearchConnectionHTTPS `json:"https" validate:"required"`

	// Connection information for cURL.
	Cli *ConnectionCLI `json:"cli" validate:"required"`
}


func (*ConnectionConnectionElasticsearchConnection) isaConnectionConnection() bool {
	return true
}

// UnmarshalConnectionConnectionElasticsearchConnection unmarshals an instance of ConnectionConnectionElasticsearchConnection from the specified map of raw messages.
func UnmarshalConnectionConnectionElasticsearchConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnectionElasticsearchConnection)
	err = core.UnmarshalModel(m, "https", &obj.Https, UnmarshalElasticsearchConnectionHTTPS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnectionEtcdConnection : etcd3 Connection Strings.
// This model "extends" ConnectionConnection
type ConnectionConnectionEtcdConnection struct {
	// GRPC(etcd3) Connection information for drivers and libraries.
	Grpc *GRPCConnectionURI `json:"grpc" validate:"required"`

	// Connection information for etcdctl.
	Cli *ConnectionCLI `json:"cli" validate:"required"`
}


func (*ConnectionConnectionEtcdConnection) isaConnectionConnection() bool {
	return true
}

// UnmarshalConnectionConnectionEtcdConnection unmarshals an instance of ConnectionConnectionEtcdConnection from the specified map of raw messages.
func UnmarshalConnectionConnectionEtcdConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnectionEtcdConnection)
	err = core.UnmarshalModel(m, "grpc", &obj.Grpc, UnmarshalGRPCConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnectionMongoDBConnection : MongoDB Connection Strings.
// This model "extends" ConnectionConnection
type ConnectionConnectionMongoDBConnection struct {
	// MongoDB Connection information for drivers and libraries.
	Mongodb *MongoDBConnectionURI `json:"mongodb" validate:"required"`

	// Connection information for mongo shell.
	Cli *ConnectionCLI `json:"cli" validate:"required"`
}


func (*ConnectionConnectionMongoDBConnection) isaConnectionConnection() bool {
	return true
}

// UnmarshalConnectionConnectionMongoDBConnection unmarshals an instance of ConnectionConnectionMongoDBConnection from the specified map of raw messages.
func UnmarshalConnectionConnectionMongoDBConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnectionMongoDBConnection)
	err = core.UnmarshalModel(m, "mongodb", &obj.Mongodb, UnmarshalMongoDBConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnectionPostgreSQLConnection : PostgreSQL and EnterpriseDB Connection Strings.
// This model "extends" ConnectionConnection
type ConnectionConnectionPostgreSQLConnection struct {
	// Connection information for drivers and libraries.
	Postgres *PostgreSQLConnectionURI `json:"postgres" validate:"required"`

	// Connection information for psql.
	Cli *ConnectionCLI `json:"cli" validate:"required"`
}


func (*ConnectionConnectionPostgreSQLConnection) isaConnectionConnection() bool {
	return true
}

// UnmarshalConnectionConnectionPostgreSQLConnection unmarshals an instance of ConnectionConnectionPostgreSQLConnection from the specified map of raw messages.
func UnmarshalConnectionConnectionPostgreSQLConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnectionPostgreSQLConnection)
	err = core.UnmarshalModel(m, "postgres", &obj.Postgres, UnmarshalPostgreSQLConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnectionRabbitMQConnection : RabbitMQ Connection Strings.
// This model "extends" ConnectionConnection
type ConnectionConnectionRabbitMQConnection struct {
	// RabbitMQ Connection information for AMQPS drivers and libraries.
	Amqps *RabbitMQConnectionAMQPS `json:"amqps" validate:"required"`

	// RabbitMQ Connection information for MQTTS drivers and libraries.
	Mqtts *RabbitMQConnectionMQTTS `json:"mqtts" validate:"required"`

	// RabbitMQ Connection information for STOMP drivers and libraries.
	StompSsl *RabbitMQConnectionStompSSL `json:"stomp_ssl" validate:"required"`

	// RabbitMQ Connection information for HTTPS.
	Https *RabbitMQConnectionHTTPS `json:"https" validate:"required"`

	// Connection information for rabbitmqadmin.
	Cli *ConnectionCLI `json:"cli" validate:"required"`
}


func (*ConnectionConnectionRabbitMQConnection) isaConnectionConnection() bool {
	return true
}

// UnmarshalConnectionConnectionRabbitMQConnection unmarshals an instance of ConnectionConnectionRabbitMQConnection from the specified map of raw messages.
func UnmarshalConnectionConnectionRabbitMQConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnectionRabbitMQConnection)
	err = core.UnmarshalModel(m, "amqps", &obj.Amqps, UnmarshalRabbitMQConnectionAMQPS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mqtts", &obj.Mqtts, UnmarshalRabbitMQConnectionMQTTS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stomp_ssl", &obj.StompSsl, UnmarshalRabbitMQConnectionStompSSL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "https", &obj.Https, UnmarshalRabbitMQConnectionHTTPS)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConnectionConnectionRedisConnection : Redis Connection Strings.
// This model "extends" ConnectionConnection
type ConnectionConnectionRedisConnection struct {
	// Connection information for drivers and libraries.
	Rediss *RedisConnectionURI `json:"rediss" validate:"required"`

	// Connection information for a Redis CLI client.
	Cli *ConnectionCLI `json:"cli" validate:"required"`
}


func (*ConnectionConnectionRedisConnection) isaConnectionConnection() bool {
	return true
}

// UnmarshalConnectionConnectionRedisConnection unmarshals an instance of ConnectionConnectionRedisConnection from the specified map of raw messages.
func UnmarshalConnectionConnectionRedisConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConnectionConnectionRedisConnection)
	err = core.UnmarshalModel(m, "rediss", &obj.Rediss, UnmarshalRedisConnectionURI)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cli", &obj.Cli, UnmarshalConnectionCLI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetConfigurationConfigurationPGConfiguration : PostgreSQL and EnterpriseDB Configuration.
// This model "extends" SetConfigurationConfiguration
type SetConfigurationConfigurationPGConfiguration struct {
	// Maximum connections allowed.
	MaxConnections *int64 `json:"max_connections,omitempty"`

	// Max number of transactions that can be in the "prepared" state simultaneously.
	MaxPreparedTransactions *int64 `json:"max_prepared_transactions,omitempty"`

	// Deadlock timeout in ms. The time to wait on a lock before checking for deadlock.  Also the duration where lock waits
	// will be logged.
	DeadlockTimeout *int64 `json:"deadlock_timeout,omitempty"`

	// Number of simultaneous requests that can be handled efficiently by the disk subsystem.
	EffectiveIoConcurrency *int64 `json:"effective_io_concurrency,omitempty"`

	// Maximum number of simultaneously defined replication slots.
	MaxReplicationSlots *int64 `json:"max_replication_slots,omitempty"`

	// Maximum number of simultaneously running WAL sender processes.
	MaxWalSenders *int64 `json:"max_wal_senders,omitempty"`

	// The number of 8kB shared memory buffers used by the server.  Set to 1/4 of memory.  Setting too high will cause
	// crashes or prevent the database from starting.
	SharedBuffers *int64 `json:"shared_buffers,omitempty"`

	// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
	// synchronous replication which will impact performance and availabilty.
	SynchronousCommit *string `json:"synchronous_commit,omitempty"`

	// WAL level.  Set to logical to use logical decoding or logical replication.
	WalLevel *string `json:"wal_level,omitempty"`

	// The number of seconds to wait before forces a switch to the next WAL file if a new file has not been started.
	ArchiveTimeout *int64 `json:"archive_timeout,omitempty"`

	// The minimum number of milliseconds for execution time above which statements will be logged.
	LogMinDurationStatement *int64 `json:"log_min_duration_statement,omitempty"`
}

// Constants associated with the SetConfigurationConfigurationPGConfiguration.SynchronousCommit property.
// Sets the current transaction's synchronization level.  Off can result in data loss.  remote_write with enable
// synchronous replication which will impact performance and availabilty.
const (
	SetConfigurationConfigurationPGConfiguration_SynchronousCommit_Local = "local"
	SetConfigurationConfigurationPGConfiguration_SynchronousCommit_Off = "off"
)

// Constants associated with the SetConfigurationConfigurationPGConfiguration.WalLevel property.
// WAL level.  Set to logical to use logical decoding or logical replication.
const (
	SetConfigurationConfigurationPGConfiguration_WalLevel_HotStandby = "hot_standby"
	SetConfigurationConfigurationPGConfiguration_WalLevel_Logical = "logical"
)


func (*SetConfigurationConfigurationPGConfiguration) isaSetConfigurationConfiguration() bool {
	return true
}

// UnmarshalSetConfigurationConfigurationPGConfiguration unmarshals an instance of SetConfigurationConfigurationPGConfiguration from the specified map of raw messages.
func UnmarshalSetConfigurationConfigurationPGConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetConfigurationConfigurationPGConfiguration)
	err = core.UnmarshalPrimitive(m, "max_connections", &obj.MaxConnections)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_prepared_transactions", &obj.MaxPreparedTransactions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deadlock_timeout", &obj.DeadlockTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_io_concurrency", &obj.EffectiveIoConcurrency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_replication_slots", &obj.MaxReplicationSlots)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_wal_senders", &obj.MaxWalSenders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shared_buffers", &obj.SharedBuffers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "synchronous_commit", &obj.SynchronousCommit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "wal_level", &obj.WalLevel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "archive_timeout", &obj.ArchiveTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_min_duration_statement", &obj.LogMinDurationStatement)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetConfigurationConfigurationRedisConfiguration : Redis Configuration.
// This model "extends" SetConfigurationConfiguration
type SetConfigurationConfigurationRedisConfiguration struct {
	// The maximum memory Redis should use, as bytes.
	MaxmemoryRedis *int64 `json:"maxmemory-redis,omitempty"`

	// The policy with which Redis evicts keys when maximum memory is reached.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// If set to yes this will enable AOF persistence.
	Appendonly *string `json:"appendonly,omitempty"`

	// The maximum memory Redis should use, as bytes.
	MaxmemorySamples *int64 `json:"maxmemory-samples,omitempty"`

	// Whether or not to stop accepting writes when background persistence actions fail.
	StopWritesOnBgsaveError *string `json:"stop-writes-on-bgsave-error,omitempty"`
}

// Constants associated with the SetConfigurationConfigurationRedisConfiguration.MaxmemoryPolicy property.
// The policy with which Redis evicts keys when maximum memory is reached.
const (
	SetConfigurationConfigurationRedisConfiguration_MaxmemoryPolicy_AllkeysLru = "allkeys-lru"
	SetConfigurationConfigurationRedisConfiguration_MaxmemoryPolicy_AllkeysRandom = "allkeys-random"
	SetConfigurationConfigurationRedisConfiguration_MaxmemoryPolicy_Noeviction = "noeviction"
	SetConfigurationConfigurationRedisConfiguration_MaxmemoryPolicy_VolatileLru = "volatile-lru"
	SetConfigurationConfigurationRedisConfiguration_MaxmemoryPolicy_VolatileRandom = "volatile-random"
	SetConfigurationConfigurationRedisConfiguration_MaxmemoryPolicy_VolatileTTL = "volatile-ttl"
)

// Constants associated with the SetConfigurationConfigurationRedisConfiguration.Appendonly property.
// If set to yes this will enable AOF persistence.
const (
	SetConfigurationConfigurationRedisConfiguration_Appendonly_No = "no"
	SetConfigurationConfigurationRedisConfiguration_Appendonly_Yes = "yes"
)

// Constants associated with the SetConfigurationConfigurationRedisConfiguration.StopWritesOnBgsaveError property.
// Whether or not to stop accepting writes when background persistence actions fail.
const (
	SetConfigurationConfigurationRedisConfiguration_StopWritesOnBgsaveError_No = "no"
	SetConfigurationConfigurationRedisConfiguration_StopWritesOnBgsaveError_Yes = "yes"
)


func (*SetConfigurationConfigurationRedisConfiguration) isaSetConfigurationConfiguration() bool {
	return true
}

// UnmarshalSetConfigurationConfigurationRedisConfiguration unmarshals an instance of SetConfigurationConfigurationRedisConfiguration from the specified map of raw messages.
func UnmarshalSetConfigurationConfigurationRedisConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetConfigurationConfigurationRedisConfiguration)
	err = core.UnmarshalPrimitive(m, "maxmemory-redis", &obj.MaxmemoryRedis)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-policy", &obj.MaxmemoryPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "appendonly", &obj.Appendonly)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxmemory-samples", &obj.MaxmemorySamples)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stop-writes-on-bgsave-error", &obj.StopWritesOnBgsaveError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupRequestSetCPUGroup : SetDeploymentScalingGroupRequestSetCPUGroup struct
// This model "extends" SetDeploymentScalingGroupRequest
type SetDeploymentScalingGroupRequestSetCPUGroup struct {
	Cpu *SetCPUGroupCPU `json:"cpu,omitempty"`
}


func (*SetDeploymentScalingGroupRequestSetCPUGroup) isaSetDeploymentScalingGroupRequest() bool {
	return true
}

// UnmarshalSetDeploymentScalingGroupRequestSetCPUGroup unmarshals an instance of SetDeploymentScalingGroupRequestSetCPUGroup from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupRequestSetCPUGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupRequestSetCPUGroup)
	err = core.UnmarshalModel(m, "cpu", &obj.Cpu, UnmarshalSetCPUGroupCPU)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupRequestSetDiskGroup : SetDeploymentScalingGroupRequestSetDiskGroup struct
// This model "extends" SetDeploymentScalingGroupRequest
type SetDeploymentScalingGroupRequestSetDiskGroup struct {
	Disk *SetDiskGroupDisk `json:"disk,omitempty"`
}


func (*SetDeploymentScalingGroupRequestSetDiskGroup) isaSetDeploymentScalingGroupRequest() bool {
	return true
}

// UnmarshalSetDeploymentScalingGroupRequestSetDiskGroup unmarshals an instance of SetDeploymentScalingGroupRequestSetDiskGroup from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupRequestSetDiskGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupRequestSetDiskGroup)
	err = core.UnmarshalModel(m, "disk", &obj.Disk, UnmarshalSetDiskGroupDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupRequestSetMembersGroup : SetDeploymentScalingGroupRequestSetMembersGroup struct
// This model "extends" SetDeploymentScalingGroupRequest
type SetDeploymentScalingGroupRequestSetMembersGroup struct {
	Members *SetMembersGroupMembers `json:"members,omitempty"`
}


func (*SetDeploymentScalingGroupRequestSetMembersGroup) isaSetDeploymentScalingGroupRequest() bool {
	return true
}

// UnmarshalSetDeploymentScalingGroupRequestSetMembersGroup unmarshals an instance of SetDeploymentScalingGroupRequestSetMembersGroup from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupRequestSetMembersGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupRequestSetMembersGroup)
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalSetMembersGroupMembers)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetDeploymentScalingGroupRequestSetMemoryGroup : SetDeploymentScalingGroupRequestSetMemoryGroup struct
// This model "extends" SetDeploymentScalingGroupRequest
type SetDeploymentScalingGroupRequestSetMemoryGroup struct {
	Memory *SetMemoryGroupMemory `json:"memory,omitempty"`
}


func (*SetDeploymentScalingGroupRequestSetMemoryGroup) isaSetDeploymentScalingGroupRequest() bool {
	return true
}

// UnmarshalSetDeploymentScalingGroupRequestSetMemoryGroup unmarshals an instance of SetDeploymentScalingGroupRequestSetMemoryGroup from the specified map of raw messages.
func UnmarshalSetDeploymentScalingGroupRequestSetMemoryGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetDeploymentScalingGroupRequestSetMemoryGroup)
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalSetMemoryGroupMemory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetPromotionPromotionPromote : Promotes a read-only replica to a full deployment.
// This model "extends" SetPromotionPromotion
type SetPromotionPromotionPromote struct {
	// Promotion options.
	Promotion map[string]interface{} `json:"promotion,omitempty"`
}


func (*SetPromotionPromotionPromote) isaSetPromotionPromotion() bool {
	return true
}

// UnmarshalSetPromotionPromotionPromote unmarshals an instance of SetPromotionPromotionPromote from the specified map of raw messages.
func UnmarshalSetPromotionPromotionPromote(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetPromotionPromotionPromote)
	err = core.UnmarshalPrimitive(m, "promotion", &obj.Promotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetPromotionPromotionUpgradePromote : Promotes a read-only replica to a full deployment running a new database version.
// This model "extends" SetPromotionPromotion
type SetPromotionPromotionUpgradePromote struct {
	// Promotion and Upgrade options.
	Promotion map[string]interface{} `json:"promotion,omitempty"`
}


func (*SetPromotionPromotionUpgradePromote) isaSetPromotionPromotion() bool {
	return true
}

// UnmarshalSetPromotionPromotionUpgradePromote unmarshals an instance of SetPromotionPromotionUpgradePromote from the specified map of raw messages.
func UnmarshalSetPromotionPromotionUpgradePromote(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SetPromotionPromotionUpgradePromote)
	err = core.UnmarshalPrimitive(m, "promotion", &obj.Promotion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
