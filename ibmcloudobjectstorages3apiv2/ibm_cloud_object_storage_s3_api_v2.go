/**
 * (C) Copyright IBM Corp. 2021.
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
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-ab7e866a-20210428-100631
 */

// Package ibmcloudobjectstorages3apiv2 : Operations and models for the IbmCloudObjectStorageS3ApiV2 service
package ibmcloudobjectstorages3apiv2

import (
	"context"
	"encoding/json"
	"fmt"
	common "github.com/IBM/experimental-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"net/http"
	"reflect"
	"time"
)

// IbmCloudObjectStorageS3ApiV2 : <p/>
//
// Version: 2.5
// See: https://cloud.ibm.com/docs/services/cloud-object-storage/
type IbmCloudObjectStorageS3ApiV2 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_cloud_object_storage_s3_api"

// IbmCloudObjectStorageS3ApiV2Options : Service options
type IbmCloudObjectStorageS3ApiV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig : constructs an instance of IbmCloudObjectStorageS3ApiV2 with passed in options and external configuration.
func NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(options *IbmCloudObjectStorageS3ApiV2Options) (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmCloudObjectStorageS3Api, err = NewIbmCloudObjectStorageS3ApiV2(options)
	if err != nil {
		return
	}

	err = ibmCloudObjectStorageS3Api.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmCloudObjectStorageS3Api.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmCloudObjectStorageS3ApiV2 : constructs an instance of IbmCloudObjectStorageS3ApiV2 with passed in options.
func NewIbmCloudObjectStorageS3ApiV2(options *IbmCloudObjectStorageS3ApiV2Options) (service *IbmCloudObjectStorageS3ApiV2, err error) {
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

	service = &IbmCloudObjectStorageS3ApiV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "ibmCloudObjectStorageS3Api" suitable for processing requests.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) Clone() *IbmCloudObjectStorageS3ApiV2 {
	if core.IsNil(ibmCloudObjectStorageS3Api) {
		return nil
	}
	clone := *ibmCloudObjectStorageS3Api
	clone.Service = ibmCloudObjectStorageS3Api.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) SetServiceURL(url string) error {
	return ibmCloudObjectStorageS3Api.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetServiceURL() string {
	return ibmCloudObjectStorageS3Api.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) SetDefaultHeaders(headers http.Header) {
	ibmCloudObjectStorageS3Api.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) SetEnableGzipCompression(enableGzip bool) {
	ibmCloudObjectStorageS3Api.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetEnableGzipCompression() bool {
	return ibmCloudObjectStorageS3Api.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	ibmCloudObjectStorageS3Api.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DisableRetries() {
	ibmCloudObjectStorageS3Api.Service.DisableRetries()
}

// HeadBucket : Read a bucket's headers
// This request is useful for checking whether a bucket has Key Protect enabled.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) HeadBucket(headBucketOptions *HeadBucketOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.HeadBucketWithContext(context.Background(), headBucketOptions)
}

// HeadBucketWithContext is an alternate form of the HeadBucket method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) HeadBucketWithContext(ctx context.Context, headBucketOptions *HeadBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(headBucketOptions, "headBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(headBucketOptions, "headBucketOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *headBucketOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.HEAD)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range headBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "HeadBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// HeadObject : Read object metadata
// The HEAD operation retrieves metadata from an object without returning the object itself. This operation is useful if
// you're only interested in an object's metadata or it's existence.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) HeadObject(headObjectOptions *HeadObjectOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.HeadObjectWithContext(context.Background(), headObjectOptions)
}

// HeadObjectWithContext is an alternate form of the HeadObject method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) HeadObjectWithContext(ctx context.Context, headObjectOptions *HeadObjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(headObjectOptions, "headObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(headObjectOptions, "headObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *headObjectOptions.Bucket,
		"Key": *headObjectOptions.Key,
	}

	builder := core.NewRequestBuilder(core.HEAD)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range headObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "HeadObject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if headObjectOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*headObjectOptions.IfMatch))
	}
	if headObjectOptions.IfModifiedSince != nil {
		builder.AddHeader("If-Modified-Since", fmt.Sprint(*headObjectOptions.IfModifiedSince))
	}
	if headObjectOptions.IfNoneMatch != nil {
		builder.AddHeader("If-None-Match", fmt.Sprint(*headObjectOptions.IfNoneMatch))
	}
	if headObjectOptions.IfUnmodifiedSince != nil {
		builder.AddHeader("If-Unmodified-Since", fmt.Sprint(*headObjectOptions.IfUnmodifiedSince))
	}
	if headObjectOptions.Range != nil {
		builder.AddHeader("Range", fmt.Sprint(*headObjectOptions.Range))
	}
	if headObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*headObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if headObjectOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*headObjectOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if headObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*headObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}

	if headObjectOptions.PartNumber != nil {
		builder.AddQuery("partNumber", fmt.Sprint(*headObjectOptions.PartNumber))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// ListBuckets : List buckets in a service instance
// This operation lists all buckets within the specified service instance, regardless of location.  Note that while any
// endpoint may be used to list all buckets, any operations that target a specific bucket must use the appropriate
// endpoint for that bucket's location.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListBuckets(listBucketsOptions *ListBucketsOptions) (result *BucketListing, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.ListBucketsWithContext(context.Background(), listBucketsOptions)
}

// ListBucketsWithContext is an alternate form of the ListBuckets method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListBucketsWithContext(ctx context.Context, listBucketsOptions *ListBucketsOptions) (result *BucketListing, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBucketsOptions, "listBucketsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listBucketsOptions, "listBucketsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBucketsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "ListBuckets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	if listBucketsOptions.IbmServiceInstanceID != nil {
		builder.AddHeader("ibm-service-instance-id", fmt.Sprint(*listBucketsOptions.IbmServiceInstanceID))
	}

	if listBucketsOptions.Extended != nil {
		builder.AddQuery("extended", fmt.Sprint(*listBucketsOptions.Extended))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBucketListing)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateBucket : Create a new bucket
// To create a bucket, you must create a COS service instance, and create an API key or HMAC credentials to authenticate
// requests.
//
// Bucket names must be globally unique and DNS-compliant; names between 3 and 63 characters long must be made of
// lowercase letters, numbers, and dashes. Bucket names must begin and end with a lowercase letter or number. Bucket
// names resembling IP addresses are not allowed. Bucket names must be unique because all buckets in the public cloud
// share a global namespace, allowing access to buckets without the need to provide any service instance or account
// information. It is not possible to create a bucket with a name beginning with `cosv1-` or `account-` as these
// prefixes are reserved by the system.
//
// Buckets are created in the location specified in the endpoint used to make the request. Once a bucket is created, it
// can be accessed at that location using any of the three networks: public, private, or direct. Any requests targeting
// an existing bucket using an endpoint with an incorrect location will result in a `404 NoSuchKey` error.
//
// All data in IBM Cloud Object Storage is encrypted at rest. This technology individually encrypts each object by using
// per-object generated keys. These keys are secured and reliably stored by using the same Information Dispersal
// Algorithms that protect object data by using an All-or-Nothing Transform (AONT). Key data is impossible to recover,
// even if individual nodes or hard disks are compromised. If it is necessary to control the encryption keys used, an
// IBM Key Protect or Hyper Protect Crypto Services root key CRN can be provided during bucket creation.
//
// It is possible to create a bucket with a "storage class" that alters the way storage charges are incurred based on
// frequency of access. This can be helpful when dealing with "cool" data that might need to be accessed without the
// delay of restoring from an archive, but is unlikely to be accessed frequently. A provisioning code can be passed in
// the S3 API `LocationConstraint` parameter to specify the storage class of a new bucket. A storage class can not be
// altered after a bucket is created.
//
// The S3 API concept of a "bucket owner" is not an individual user, but instead is considered to be the Service
// Instance associated with the bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) CreateBucket(createBucketOptions *CreateBucketOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.CreateBucketWithContext(context.Background(), createBucketOptions)
}

// CreateBucketWithContext is an alternate form of the CreateBucket method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) CreateBucketWithContext(ctx context.Context, createBucketOptions *CreateBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createBucketOptions, "createBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createBucketOptions, "createBucketOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *createBucketOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "CreateBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "text/xml")
	if createBucketOptions.IbmServiceInstanceID != nil {
		builder.AddHeader("ibm-service-instance-id", fmt.Sprint(*createBucketOptions.IbmServiceInstanceID))
	}
	if createBucketOptions.IbmSseKpEncryptionAlgorithm != nil {
		builder.AddHeader("ibm-sse-kp-encryption-algorithm", fmt.Sprint(*createBucketOptions.IbmSseKpEncryptionAlgorithm))
	}
	if createBucketOptions.IbmSseKpCustomerRootKeyCrn != nil {
		builder.AddHeader("ibm-sse-kp-customer-root-key-crn", fmt.Sprint(*createBucketOptions.IbmSseKpCustomerRootKeyCrn))
	}
	if createBucketOptions.XAmzAcl != nil {
		builder.AddHeader("x-amz-acl", fmt.Sprint(*createBucketOptions.XAmzAcl))
	}

	_, err = builder.SetBodyContent("text/xml", nil, nil, createBucketOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// DeleteBucket : Delete a bucket
// Only empty buckets may be deleted. A bucket name is returned to the available namespace approximately 10 minutes
// after deletion.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucket(deleteBucketOptions *DeleteBucketOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteBucketWithContext(context.Background(), deleteBucketOptions)
}

// DeleteBucketWithContext is an alternate form of the DeleteBucket method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketWithContext(ctx context.Context, deleteBucketOptions *DeleteBucketOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBucketOptions, "deleteBucketOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBucketOptions, "deleteBucketOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteBucketOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBucketOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteBucket")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// ListObjects : List objects in a bucket (v1)
// Returns some or all (up to 1,000) of the objects in a bucket. You can use the request parameters as selection
// criteria to return a subset of the objects in a bucket. A `200 OK` response can contain valid or invalid XML. Make
// sure to design your application to parse the contents of the response and handle it appropriately. This version (v1)
// uses a `marker` parameter to list objects starting with a given object. Version 2 of this API provides a continuation
// token instead, making it a bit more straightforward to chain listing requests for buckets with large numbers of
// objects.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListObjects(listObjectsOptions *ListObjectsOptions) (result *ListObjectsOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.ListObjectsWithContext(context.Background(), listObjectsOptions)
}

// ListObjectsWithContext is an alternate form of the ListObjects method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListObjectsWithContext(ctx context.Context, listObjectsOptions *ListObjectsOptions) (result *ListObjectsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listObjectsOptions, "listObjectsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listObjectsOptions, "listObjectsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *listObjectsOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listObjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "ListObjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	if listObjectsOptions.Delimiter != nil {
		builder.AddQuery("delimiter", fmt.Sprint(*listObjectsOptions.Delimiter))
	}
	if listObjectsOptions.EncodingType != nil {
		builder.AddQuery("encoding-type", fmt.Sprint(*listObjectsOptions.EncodingType))
	}
	if listObjectsOptions.Marker != nil {
		builder.AddQuery("marker", fmt.Sprint(*listObjectsOptions.Marker))
	}
	if listObjectsOptions.MaxKeys != nil {
		builder.AddQuery("max-keys", fmt.Sprint(*listObjectsOptions.MaxKeys))
	}
	if listObjectsOptions.Prefix != nil {
		builder.AddQuery("prefix", fmt.Sprint(*listObjectsOptions.Prefix))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListObjectsOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListObjectsV2 : List objects in a bucket (v2)
// Returns some or all (up to 1,000) of the objects in a bucket. You can use the request parameters as selection
// criteria to return a subset of the objects in a bucket. A `200 OK` response can contain valid or invalid XML. Make
// sure to design your application to parse the contents of the response and handle it appropriately.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListObjectsV2(listObjectsV2Options *ListObjectsV2Options) (result *ListObjectsV2Output, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.ListObjectsV2WithContext(context.Background(), listObjectsV2Options)
}

// ListObjectsV2WithContext is an alternate form of the ListObjectsV2 method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListObjectsV2WithContext(ctx context.Context, listObjectsV2Options *ListObjectsV2Options) (result *ListObjectsV2Output, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listObjectsV2Options, "listObjectsV2Options cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listObjectsV2Options, "listObjectsV2Options")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *listObjectsV2Options.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?list-type=2`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listObjectsV2Options.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "ListObjectsV2")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("list-type", fmt.Sprint(*listObjectsV2Options.ListType))
	if listObjectsV2Options.Delimiter != nil {
		builder.AddQuery("delimiter", fmt.Sprint(*listObjectsV2Options.Delimiter))
	}
	if listObjectsV2Options.EncodingType != nil {
		builder.AddQuery("encoding-type", fmt.Sprint(*listObjectsV2Options.EncodingType))
	}
	if listObjectsV2Options.MaxKeys != nil {
		builder.AddQuery("max-keys", fmt.Sprint(*listObjectsV2Options.MaxKeys))
	}
	if listObjectsV2Options.Prefix != nil {
		builder.AddQuery("prefix", fmt.Sprint(*listObjectsV2Options.Prefix))
	}
	if listObjectsV2Options.ContinuationToken != nil {
		builder.AddQuery("continuation-token", fmt.Sprint(*listObjectsV2Options.ContinuationToken))
	}
	if listObjectsV2Options.FetchOwner != nil {
		builder.AddQuery("fetch-owner", fmt.Sprint(*listObjectsV2Options.FetchOwner))
	}
	if listObjectsV2Options.StartAfter != nil {
		builder.AddQuery("start-after", fmt.Sprint(*listObjectsV2Options.StartAfter))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListObjectsV2Output)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PutObject : Create (upload) an object
// Adds an object to a bucket using a single request. IBM COS never adds partial objects; if you receive a success
// response, IBM COS added the entire object to the bucket. IBM COS is a distributed system. If it receives multiple
// write requests for the same object simultaneously, it overwrites all but the last object written. IBM COS does not
// provide object locking; if you need this, make sure to build it into your application layer.
//
// All objects written to IBM COS are encrypted by default using SecureSlice.  If you require possession of encryption
// keys, you can use Key Protect or SSE-C.
//
// To ensure that data is not corrupted traversing the network, use the `Content-MD5` header. When you use this header,
// IBM COS checks the object against the provided MD5 value and, if they do not match, returns an error. Additionally,
// you can calculate the MD5 while putting an object to IBM COS and compare the returned ETag to the calculated MD5
// value. The `Content-MD5` header is required for any request to upload an object with a retention period configured
// using Immutable Object Storage.
//
// Larger objects (greater than 100 MiB) may benefit from breaking the object into multiple parts and uploading the
// parts in parallel.  For more information, see the **Multipart uploads** methods.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutObject(putObjectOptions *PutObjectOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutObjectWithContext(context.Background(), putObjectOptions)
}

// PutObjectWithContext is an alternate form of the PutObject method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutObjectWithContext(ctx context.Context, putObjectOptions *PutObjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putObjectOptions, "putObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putObjectOptions, "putObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putObjectOptions.Bucket,
		"Key": *putObjectOptions.Key,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutObject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")
	if putObjectOptions.XAmzAcl != nil {
		builder.AddHeader("x-amz-acl", fmt.Sprint(*putObjectOptions.XAmzAcl))
	}
	if putObjectOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*putObjectOptions.IfMatch))
	}
	if putObjectOptions.IfNoneMatch != nil {
		builder.AddHeader("If-None-Match", fmt.Sprint(*putObjectOptions.IfNoneMatch))
	}
	if putObjectOptions.IfUnmodifiedSince != nil {
		builder.AddHeader("If-Unmodified-Since", fmt.Sprint(*putObjectOptions.IfUnmodifiedSince))
	}
	if putObjectOptions.CacheControl != nil {
		builder.AddHeader("Cache-Control", fmt.Sprint(*putObjectOptions.CacheControl))
	}
	if putObjectOptions.ContentDisposition != nil {
		builder.AddHeader("Content-Disposition", fmt.Sprint(*putObjectOptions.ContentDisposition))
	}
	if putObjectOptions.ContentEncoding != nil {
		builder.AddHeader("Content-Encoding", fmt.Sprint(*putObjectOptions.ContentEncoding))
	}
	if putObjectOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*putObjectOptions.ContentLanguage))
	}
	if putObjectOptions.ContentLength != nil {
		builder.AddHeader("Content-Length", fmt.Sprint(*putObjectOptions.ContentLength))
	}
	if putObjectOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*putObjectOptions.ContentMD5))
	}
	if putObjectOptions.Expires != nil {
		builder.AddHeader("Expires", fmt.Sprint(*putObjectOptions.Expires))
	}
	if putObjectOptions.XAmzServerSideEncryption != nil {
		builder.AddHeader("x-amz-server-side-encryption", fmt.Sprint(*putObjectOptions.XAmzServerSideEncryption))
	}
	if putObjectOptions.XAmzWebsiteRedirectLocation != nil {
		builder.AddHeader("x-amz-website-redirect-location", fmt.Sprint(*putObjectOptions.XAmzWebsiteRedirectLocation))
	}
	if putObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*putObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if putObjectOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*putObjectOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if putObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*putObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}
	if putObjectOptions.XAmzTagging != nil {
		builder.AddHeader("x-amz-tagging", fmt.Sprint(*putObjectOptions.XAmzTagging))
	}

	_, err = builder.SetBodyContent("text/xml", nil, nil, putObjectOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetObject : Read (download) an object
// Retrieves objects from IBM COS. If the object you are retrieving has been archived, before you can retrieve the
// object you must first restore a copy. Otherwise, this operation returns an `InvalidObjectStateError` error. If you
// encrypt an object by using server-side encryption with customer-provided encryption keys (SSE-C) when you store the
// object in IBM COS, then when you GET the object, you must use the following headers:
// * `x-amz-server-side-encryption-customer-algorithm`
// * `x-amz-server-side-encryption-customer-key`
// * `x-amz-server-side-encryption-customer-key-MD5`.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetObject(getObjectOptions *GetObjectOptions) (result *GetObjectOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetObjectWithContext(context.Background(), getObjectOptions)
}

// GetObjectWithContext is an alternate form of the GetObject method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetObjectWithContext(ctx context.Context, getObjectOptions *GetObjectOptions) (result *GetObjectOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getObjectOptions, "getObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getObjectOptions, "getObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getObjectOptions.Bucket,
		"Key": *getObjectOptions.Key,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetObject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	if getObjectOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*getObjectOptions.IfMatch))
	}
	if getObjectOptions.IfModifiedSince != nil {
		builder.AddHeader("If-Modified-Since", fmt.Sprint(*getObjectOptions.IfModifiedSince))
	}
	if getObjectOptions.IfNoneMatch != nil {
		builder.AddHeader("If-None-Match", fmt.Sprint(*getObjectOptions.IfNoneMatch))
	}
	if getObjectOptions.IfUnmodifiedSince != nil {
		builder.AddHeader("If-Unmodified-Since", fmt.Sprint(*getObjectOptions.IfUnmodifiedSince))
	}
	if getObjectOptions.Range != nil {
		builder.AddHeader("Range", fmt.Sprint(*getObjectOptions.Range))
	}
	if getObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*getObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if getObjectOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*getObjectOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if getObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*getObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}

	if getObjectOptions.ResponseCacheControl != nil {
		builder.AddQuery("response-cache-control", fmt.Sprint(*getObjectOptions.ResponseCacheControl))
	}
	if getObjectOptions.ResponseContentDisposition != nil {
		builder.AddQuery("response-content-disposition", fmt.Sprint(*getObjectOptions.ResponseContentDisposition))
	}
	if getObjectOptions.ResponseContentEncoding != nil {
		builder.AddQuery("response-content-encoding", fmt.Sprint(*getObjectOptions.ResponseContentEncoding))
	}
	if getObjectOptions.ResponseContentLanguage != nil {
		builder.AddQuery("response-content-language", fmt.Sprint(*getObjectOptions.ResponseContentLanguage))
	}
	if getObjectOptions.ResponseContentType != nil {
		builder.AddQuery("response-content-type", fmt.Sprint(*getObjectOptions.ResponseContentType))
	}
	if getObjectOptions.ResponseExpires != nil {
		builder.AddQuery("response-expires", fmt.Sprint(*getObjectOptions.ResponseExpires))
	}
	if getObjectOptions.PartNumber != nil {
		builder.AddQuery("partNumber", fmt.Sprint(*getObjectOptions.PartNumber))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetObjectOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteObject : Delete an object
// Permantently deletes an object. This operation is final - there is no way to recover a deleted object. Data stored in
// IBM COS is erasure coded and distributed to multiple individual storage devices in multiple data centers. When data
// is deleted, various mechanisms exist which prevent recovery or reconstruction of the deleted objects.
//
// Deletion of an object undergoes various stages. First, the metadata is marked to indicate the object is deleted,
// then, the data is removed. Eventually, deleted metadata is overwritten by a process of compaction and the deleted
// data blocks are overwritten with new data in the course of normal operations. As soon as the metadata is marked
// deleted, it is not possible to read an object remotely. IBM's provider-managed encryption and erasure coding prevents
// data (both before and after deletion) from being accessible from within individual data centers.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteObject(deleteObjectOptions *DeleteObjectOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteObjectWithContext(context.Background(), deleteObjectOptions)
}

// DeleteObjectWithContext is an alternate form of the DeleteObject method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteObjectWithContext(ctx context.Context, deleteObjectOptions *DeleteObjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteObjectOptions, "deleteObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteObjectOptions, "deleteObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteObjectOptions.Bucket,
		"Key": *deleteObjectOptions.Key,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteObject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// CopyObject : Copy an object
// Copies an object.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) CopyObject(copyObjectOptions *CopyObjectOptions) (result *CopyObjectOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.CopyObjectWithContext(context.Background(), copyObjectOptions)
}

// CopyObjectWithContext is an alternate form of the CopyObject method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) CopyObjectWithContext(ctx context.Context, copyObjectOptions *CopyObjectOptions) (result *CopyObjectOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(copyObjectOptions, "copyObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(copyObjectOptions, "copyObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *copyObjectOptions.Bucket,
		"TargetKey": *copyObjectOptions.TargetKey,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{TargetKey}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range copyObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "CopyObject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")
	if copyObjectOptions.XAmzCopySource != nil {
		builder.AddHeader("x-amz-copy-source", fmt.Sprint(*copyObjectOptions.XAmzCopySource))
	}
	if copyObjectOptions.XAmzAcl != nil {
		builder.AddHeader("x-amz-acl", fmt.Sprint(*copyObjectOptions.XAmzAcl))
	}
	if copyObjectOptions.CacheControl != nil {
		builder.AddHeader("Cache-Control", fmt.Sprint(*copyObjectOptions.CacheControl))
	}
	if copyObjectOptions.ContentDisposition != nil {
		builder.AddHeader("Content-Disposition", fmt.Sprint(*copyObjectOptions.ContentDisposition))
	}
	if copyObjectOptions.ContentEncoding != nil {
		builder.AddHeader("Content-Encoding", fmt.Sprint(*copyObjectOptions.ContentEncoding))
	}
	if copyObjectOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*copyObjectOptions.ContentLanguage))
	}
	if copyObjectOptions.XAmzCopySourceIfMatch != nil {
		builder.AddHeader("x-amz-copy-source-if-match", fmt.Sprint(*copyObjectOptions.XAmzCopySourceIfMatch))
	}
	if copyObjectOptions.XAmzCopySourceIfModifiedSince != nil {
		builder.AddHeader("x-amz-copy-source-if-modified-since", fmt.Sprint(*copyObjectOptions.XAmzCopySourceIfModifiedSince))
	}
	if copyObjectOptions.XAmzCopySourceIfNoneMatch != nil {
		builder.AddHeader("x-amz-copy-source-if-none-match", fmt.Sprint(*copyObjectOptions.XAmzCopySourceIfNoneMatch))
	}
	if copyObjectOptions.XAmzCopySourceIfUnmodifiedSince != nil {
		builder.AddHeader("x-amz-copy-source-if-unmodified-since", fmt.Sprint(*copyObjectOptions.XAmzCopySourceIfUnmodifiedSince))
	}
	if copyObjectOptions.Expires != nil {
		builder.AddHeader("Expires", fmt.Sprint(*copyObjectOptions.Expires))
	}
	if copyObjectOptions.XAmzMetadataDirective != nil {
		builder.AddHeader("x-amz-metadata-directive", fmt.Sprint(*copyObjectOptions.XAmzMetadataDirective))
	}
	if copyObjectOptions.XAmzTaggingDirective != nil {
		builder.AddHeader("x-amz-tagging-directive", fmt.Sprint(*copyObjectOptions.XAmzTaggingDirective))
	}
	if copyObjectOptions.XAmzServerSideEncryption != nil {
		builder.AddHeader("x-amz-server-side-encryption", fmt.Sprint(*copyObjectOptions.XAmzServerSideEncryption))
	}
	if copyObjectOptions.XAmzWebsiteRedirectLocation != nil {
		builder.AddHeader("x-amz-website-redirect-location", fmt.Sprint(*copyObjectOptions.XAmzWebsiteRedirectLocation))
	}
	if copyObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*copyObjectOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if copyObjectOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*copyObjectOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if copyObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*copyObjectOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}
	if copyObjectOptions.XAmzCopySourceServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-copy-source-server-side-encryption-customer-algorithm", fmt.Sprint(*copyObjectOptions.XAmzCopySourceServerSideEncryptionCustomerAlgorithm))
	}
	if copyObjectOptions.XAmzCopySourceServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-copy-source-server-side-encryption-customer-key", fmt.Sprint(*copyObjectOptions.XAmzCopySourceServerSideEncryptionCustomerKey))
	}
	if copyObjectOptions.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-copy-source-server-side-encryption-customer-key-MD5", fmt.Sprint(*copyObjectOptions.XAmzCopySourceServerSideEncryptionCustomerKeyMD5))
	}
	if copyObjectOptions.XAmzTagging != nil {
		builder.AddHeader("x-amz-tagging", fmt.Sprint(*copyObjectOptions.XAmzTagging))
	}

	_, err = builder.SetBodyContent("text/xml", nil, nil, copyObjectOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCopyObjectOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteObjects : Delete multiple objects
// This operation enables you to delete multiple objects from a bucket using a single HTTP request. If you know the
// object keys that you want to delete, then this operation provides a suitable alternative to sending individual delete
// requests, reducing per-request overhead. The request contains a list of up to 1000 keys that you want to delete. In
// the XML, you provide the object key names. For each key, IBM COS performs a delete operation and returns the result
// of that delete, success, or failure, in the response. Note that if the object specified in the request is not found,
// IBM COS returns the result as deleted. The operation supports two modes for the response: verbose and quiet. By
// default, the operation uses verbose mode in which the response includes the result of deletion of each key in your
// request. In quiet mode the response includes only keys where the delete operation encountered an error. For a
// successful deletion, the operation does not return any information about the delete in the response body.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteObjects(deleteObjectsOptions *DeleteObjectsOptions) (result *DeleteObjectsOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteObjectsWithContext(context.Background(), deleteObjectsOptions)
}

// DeleteObjectsWithContext is an alternate form of the DeleteObjects method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteObjectsWithContext(ctx context.Context, deleteObjectsOptions *DeleteObjectsOptions) (result *DeleteObjectsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteObjectsOptions, "deleteObjectsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteObjectsOptions, "deleteObjectsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteObjectsOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?delete`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteObjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteObjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")

	builder.AddQuery("delete", fmt.Sprint(*deleteObjectsOptions.Delete))

	_, err = builder.SetBodyContent("text/xml", nil, nil, deleteObjectsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteObjectsOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PutBucketProtectionConfiguration : Create a protection configuration
// Creates a new protection configuration (this term is interchangeable with "retention policy") for the bucket or
// replaces an existing protection configuration. You specify the protection configuration in your request body. The
// protection configuration is specified as XML.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketProtectionConfiguration(putBucketProtectionConfigurationOptions *PutBucketProtectionConfigurationOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutBucketProtectionConfigurationWithContext(context.Background(), putBucketProtectionConfigurationOptions)
}

// PutBucketProtectionConfigurationWithContext is an alternate form of the PutBucketProtectionConfiguration method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketProtectionConfigurationWithContext(ctx context.Context, putBucketProtectionConfigurationOptions *PutBucketProtectionConfigurationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putBucketProtectionConfigurationOptions, "putBucketProtectionConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putBucketProtectionConfigurationOptions, "putBucketProtectionConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putBucketProtectionConfigurationOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?protection`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putBucketProtectionConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutBucketProtectionConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "text/xml")

	builder.AddQuery("protection", fmt.Sprint(*putBucketProtectionConfigurationOptions.Protection))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putBucketProtectionConfigurationOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// PutBucketLifecycleConfiguration : Create a lifecycle configuration
// Creates a new lifecycle configuration for the bucket or replaces an existing lifecycle configuration. You specify the
// lifecycle configuration in your request body. The lifecycle configuration is specified as XML consisting of one or
// more rules. Each rule consists of the following:
//
// * Filter identifying a subset of objects to which the rule applies. The filter can be based on a key name prefix,
// object tags, or a combination of both.
// * Status whether the rule is in effect.
// * One or more lifecycle transition and expiration actions that you want IBM COS to perform on the objects identified
// by the filter. Note that there can only be one `Transistion` rule, and filters are not supported for `Transition`
// rules.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketLifecycleConfiguration(putBucketLifecycleConfigurationOptions *PutBucketLifecycleConfigurationOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutBucketLifecycleConfigurationWithContext(context.Background(), putBucketLifecycleConfigurationOptions)
}

// PutBucketLifecycleConfigurationWithContext is an alternate form of the PutBucketLifecycleConfiguration method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketLifecycleConfigurationWithContext(ctx context.Context, putBucketLifecycleConfigurationOptions *PutBucketLifecycleConfigurationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putBucketLifecycleConfigurationOptions, "putBucketLifecycleConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putBucketLifecycleConfigurationOptions, "putBucketLifecycleConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putBucketLifecycleConfigurationOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?lifecycle`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putBucketLifecycleConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutBucketLifecycleConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "text/xml")

	builder.AddQuery("lifecycle", fmt.Sprint(*putBucketLifecycleConfigurationOptions.Lifecycle))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putBucketLifecycleConfigurationOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetBucketLifecycleConfiguration : Read a lifecycle configuration
// Returns the lifecycle configuration information set on the bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptions *GetBucketLifecycleConfigurationOptions) (result *GetBucketLifecycleConfigurationOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetBucketLifecycleConfigurationWithContext(context.Background(), getBucketLifecycleConfigurationOptions)
}

// GetBucketLifecycleConfigurationWithContext is an alternate form of the GetBucketLifecycleConfiguration method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketLifecycleConfigurationWithContext(ctx context.Context, getBucketLifecycleConfigurationOptions *GetBucketLifecycleConfigurationOptions) (result *GetBucketLifecycleConfigurationOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketLifecycleConfigurationOptions, "getBucketLifecycleConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketLifecycleConfigurationOptions, "getBucketLifecycleConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getBucketLifecycleConfigurationOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?lifecycle`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketLifecycleConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetBucketLifecycleConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("lifecycle", fmt.Sprint(*getBucketLifecycleConfigurationOptions.Lifecycle))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBucketLifecycleConfigurationOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteBucketLifecycle : Delete a lifecycle configuration
// Deletes the lifecycle configuration from the specified bucket. IBM COS removes all the lifecycle configuration rules
// in the lifecycle subresource associated with the bucket. Your objects never expire, and IBM COS no longer
// automatically deletes any objects on the basis of rules contained in the deleted lifecycle configuration. There is
// usually some time lag before lifecycle configuration deletion is fully propagated across the IBM COS system.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketLifecycle(deleteBucketLifecycleOptions *DeleteBucketLifecycleOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteBucketLifecycleWithContext(context.Background(), deleteBucketLifecycleOptions)
}

// DeleteBucketLifecycleWithContext is an alternate form of the DeleteBucketLifecycle method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketLifecycleWithContext(ctx context.Context, deleteBucketLifecycleOptions *DeleteBucketLifecycleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBucketLifecycleOptions, "deleteBucketLifecycleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBucketLifecycleOptions, "deleteBucketLifecycleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteBucketLifecycleOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?lifecycle`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBucketLifecycleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteBucketLifecycle")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("lifecycle", fmt.Sprint(*deleteBucketLifecycleOptions.Lifecycle))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// RestoreObject : Temporarily restore an archived object
// Restores an archived copy of an object back into IBM COS. To access an archived object, you must first initiate a
// restore request. This restores a temporary copy of the archived object. In a restore request, you specify the number
// of days that you want the restored copy to exist. After the specified period, IBM COS deletes the temporary copy but
// the object remains archived.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) RestoreObject(restoreObjectOptions *RestoreObjectOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.RestoreObjectWithContext(context.Background(), restoreObjectOptions)
}

// RestoreObjectWithContext is an alternate form of the RestoreObject method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) RestoreObjectWithContext(ctx context.Context, restoreObjectOptions *RestoreObjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restoreObjectOptions, "restoreObjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(restoreObjectOptions, "restoreObjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *restoreObjectOptions.Bucket,
		"Key": *restoreObjectOptions.Key,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?restore`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range restoreObjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "RestoreObject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")

	builder.AddQuery("restore", fmt.Sprint(*restoreObjectOptions.Restore))

	_, err = builder.SetBodyContent("text/xml", nil, nil, restoreObjectOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// InitiateMultipartUpload : Initiate a multipart upload
// This operation initiates a multipart upload and returns an upload ID. This upload ID is used to associate all of the
// parts in the specific multipart upload. You specify this upload ID in each of your subsequent upload part requests.
// You also include this upload ID in the final request to either complete or abort the multipart upload request.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) InitiateMultipartUpload(initiateMultipartUploadOptions *InitiateMultipartUploadOptions) (result *CreateMultipartUploadOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.InitiateMultipartUploadWithContext(context.Background(), initiateMultipartUploadOptions)
}

// InitiateMultipartUploadWithContext is an alternate form of the InitiateMultipartUpload method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) InitiateMultipartUploadWithContext(ctx context.Context, initiateMultipartUploadOptions *InitiateMultipartUploadOptions) (result *CreateMultipartUploadOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(initiateMultipartUploadOptions, "initiateMultipartUploadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(initiateMultipartUploadOptions, "initiateMultipartUploadOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *initiateMultipartUploadOptions.Bucket,
		"Key": *initiateMultipartUploadOptions.Key,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?uploads`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range initiateMultipartUploadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "InitiateMultipartUpload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")
	if initiateMultipartUploadOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*initiateMultipartUploadOptions.IfMatch))
	}
	if initiateMultipartUploadOptions.IfNoneMatch != nil {
		builder.AddHeader("If-None-Match", fmt.Sprint(*initiateMultipartUploadOptions.IfNoneMatch))
	}
	if initiateMultipartUploadOptions.IfUnmodifiedSince != nil {
		builder.AddHeader("If-Unmodified-Since", fmt.Sprint(*initiateMultipartUploadOptions.IfUnmodifiedSince))
	}
	if initiateMultipartUploadOptions.CacheControl != nil {
		builder.AddHeader("Cache-Control", fmt.Sprint(*initiateMultipartUploadOptions.CacheControl))
	}
	if initiateMultipartUploadOptions.ContentDisposition != nil {
		builder.AddHeader("Content-Disposition", fmt.Sprint(*initiateMultipartUploadOptions.ContentDisposition))
	}
	if initiateMultipartUploadOptions.ContentEncoding != nil {
		builder.AddHeader("Content-Encoding", fmt.Sprint(*initiateMultipartUploadOptions.ContentEncoding))
	}
	if initiateMultipartUploadOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*initiateMultipartUploadOptions.ContentLanguage))
	}
	if initiateMultipartUploadOptions.Expires != nil {
		builder.AddHeader("Expires", fmt.Sprint(*initiateMultipartUploadOptions.Expires))
	}
	if initiateMultipartUploadOptions.XAmzServerSideEncryption != nil {
		builder.AddHeader("x-amz-server-side-encryption", fmt.Sprint(*initiateMultipartUploadOptions.XAmzServerSideEncryption))
	}
	if initiateMultipartUploadOptions.XAmzWebsiteRedirectLocation != nil {
		builder.AddHeader("x-amz-website-redirect-location", fmt.Sprint(*initiateMultipartUploadOptions.XAmzWebsiteRedirectLocation))
	}
	if initiateMultipartUploadOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*initiateMultipartUploadOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if initiateMultipartUploadOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*initiateMultipartUploadOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if initiateMultipartUploadOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*initiateMultipartUploadOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}
	if initiateMultipartUploadOptions.XAmzTagging != nil {
		builder.AddHeader("x-amz-tagging", fmt.Sprint(*initiateMultipartUploadOptions.XAmzTagging))
	}
	if initiateMultipartUploadOptions.XAmzAcl != nil {
		builder.AddHeader("x-amz-acl", fmt.Sprint(*initiateMultipartUploadOptions.XAmzAcl))
	}

	builder.AddQuery("uploads", fmt.Sprint(*initiateMultipartUploadOptions.Uploads))

	_, err = builder.SetBodyContent("text/xml", nil, nil, initiateMultipartUploadOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateMultipartUploadOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CompleteMultipartUpload : Complete a multipart upload
// Completes a multipart upload by assembling previously uploaded parts.
//
// After successfully uploading all parts of an upload, you call this operation to complete the upload. Upon receiving
// this request, IBM COS concatenates all the parts in ascending order by part number to create a new object. In the
// Complete Multipart Upload request, you must provide the parts list. You must ensure that the parts list is complete.
// This operation concatenates the parts that you provide in the list. For each part in the list, you must provide the
// part number and the `ETag` value, returned after that part was uploaded.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) CompleteMultipartUpload(completeMultipartUploadOptions *CompleteMultipartUploadOptions) (result *CompleteMultipartUploadOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.CompleteMultipartUploadWithContext(context.Background(), completeMultipartUploadOptions)
}

// CompleteMultipartUploadWithContext is an alternate form of the CompleteMultipartUpload method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) CompleteMultipartUploadWithContext(ctx context.Context, completeMultipartUploadOptions *CompleteMultipartUploadOptions) (result *CompleteMultipartUploadOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(completeMultipartUploadOptions, "completeMultipartUploadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(completeMultipartUploadOptions, "completeMultipartUploadOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *completeMultipartUploadOptions.Bucket,
		"Key": *completeMultipartUploadOptions.Key,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?uploadId={uploadId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range completeMultipartUploadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "CompleteMultipartUpload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")

	builder.AddQuery("uploadId", fmt.Sprint(*completeMultipartUploadOptions.UploadID))

	_, err = builder.SetBodyContent("text/xml", nil, nil, completeMultipartUploadOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCompleteMultipartUploadOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListParts : List parts of a multipart upload
// Lists the parts that have been uploaded for a specific multipart upload. This operation must include the upload ID,
// which you obtain by sending the initiate multipart upload request. This request by default returns a maximum of 1,000
// uploaded parts. You can restrict the number of parts returned by specifying the `max-parts` request parameter. If
// your multipart upload consists of more than 1,000 parts, the response returns an `IsTruncated` field with the value
// of true, and a `NextPartNumberMarker` element. In subsequent `ListParts` requests you can include the
// `part-number-marker` query string parameter and set its value to the `NextPartNumberMarker` field value from the
// previous response.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListParts(listPartsOptions *ListPartsOptions) (result *ListPartsOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.ListPartsWithContext(context.Background(), listPartsOptions)
}

// ListPartsWithContext is an alternate form of the ListParts method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListPartsWithContext(ctx context.Context, listPartsOptions *ListPartsOptions) (result *ListPartsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPartsOptions, "listPartsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPartsOptions, "listPartsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *listPartsOptions.Bucket,
		"Key": *listPartsOptions.Key,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?uploadId={uploadId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPartsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "ListParts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("uploadId", fmt.Sprint(*listPartsOptions.UploadID))
	if listPartsOptions.MaxParts != nil {
		builder.AddQuery("max-parts", fmt.Sprint(*listPartsOptions.MaxParts))
	}
	if listPartsOptions.PartNumberMarker != nil {
		builder.AddQuery("part-number-marker", fmt.Sprint(*listPartsOptions.PartNumberMarker))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPartsOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AbortMultipartUpload : Abort a multipart upload
// Stops a current multipart upload and removes any parts of an incomplete upload, which would otherwise incur storage
// costs.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) AbortMultipartUpload(abortMultipartUploadOptions *AbortMultipartUploadOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.AbortMultipartUploadWithContext(context.Background(), abortMultipartUploadOptions)
}

// AbortMultipartUploadWithContext is an alternate form of the AbortMultipartUpload method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) AbortMultipartUploadWithContext(ctx context.Context, abortMultipartUploadOptions *AbortMultipartUploadOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(abortMultipartUploadOptions, "abortMultipartUploadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(abortMultipartUploadOptions, "abortMultipartUploadOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *abortMultipartUploadOptions.Bucket,
		"Key": *abortMultipartUploadOptions.Key,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?uploadId={uploadId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range abortMultipartUploadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "AbortMultipartUpload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("uploadId", fmt.Sprint(*abortMultipartUploadOptions.UploadID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// ListMultipartUploads : List active multipart uploads
// This operation lists in-progress multipart uploads. An in-progress multipart upload is a multipart upload that has
// been initiated using the Initiate Multipart Upload request, but has not yet been completed or aborted. This operation
// returns at most 1,000 multipart uploads in the response. 1,000 multipart uploads is the maximum number of uploads a
// response can include, which is also the default value. You can further limit the number of uploads in a response by
// specifying the `max-uploads` parameter in the response. If additional multipart uploads satisfy the list criteria,
// the response will contain an `IsTruncated` element with the value true. To list the additional multipart uploads, use
// the `key-marker` and `upload-id-marker` request parameters. In the response, the uploads are sorted by key. If your
// application has initiated more than one multipart upload using the same object key, then uploads in the response are
// first sorted by key. Additionally, uploads are sorted in ascending order within each key by the upload initiation
// time.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListMultipartUploads(listMultipartUploadsOptions *ListMultipartUploadsOptions) (result *ListMultipartUploadsOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.ListMultipartUploadsWithContext(context.Background(), listMultipartUploadsOptions)
}

// ListMultipartUploadsWithContext is an alternate form of the ListMultipartUploads method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) ListMultipartUploadsWithContext(ctx context.Context, listMultipartUploadsOptions *ListMultipartUploadsOptions) (result *ListMultipartUploadsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listMultipartUploadsOptions, "listMultipartUploadsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listMultipartUploadsOptions, "listMultipartUploadsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *listMultipartUploadsOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?uploads`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listMultipartUploadsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "ListMultipartUploads")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("uploads", fmt.Sprint(*listMultipartUploadsOptions.Uploads))
	if listMultipartUploadsOptions.Delimiter != nil {
		builder.AddQuery("delimiter", fmt.Sprint(*listMultipartUploadsOptions.Delimiter))
	}
	if listMultipartUploadsOptions.EncodingType != nil {
		builder.AddQuery("encoding-type", fmt.Sprint(*listMultipartUploadsOptions.EncodingType))
	}
	if listMultipartUploadsOptions.KeyMarker != nil {
		builder.AddQuery("key-marker", fmt.Sprint(*listMultipartUploadsOptions.KeyMarker))
	}
	if listMultipartUploadsOptions.MaxUploads != nil {
		builder.AddQuery("max-uploads", fmt.Sprint(*listMultipartUploadsOptions.MaxUploads))
	}
	if listMultipartUploadsOptions.Prefix != nil {
		builder.AddQuery("prefix", fmt.Sprint(*listMultipartUploadsOptions.Prefix))
	}
	if listMultipartUploadsOptions.UploadIdMarker != nil {
		builder.AddQuery("upload-id-marker", fmt.Sprint(*listMultipartUploadsOptions.UploadIdMarker))
	}
	if listMultipartUploadsOptions.PaginationLimit != nil {
		builder.AddQuery("PaginationLimit", fmt.Sprint(*listMultipartUploadsOptions.PaginationLimit))
	}
	if listMultipartUploadsOptions.PaginationToken != nil {
		builder.AddQuery("PaginationToken", fmt.Sprint(*listMultipartUploadsOptions.PaginationToken))
	}
	if listMultipartUploadsOptions.PaginationToken != nil {
		builder.AddQuery("PaginationToken", fmt.Sprint(*listMultipartUploadsOptions.PaginationToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListMultipartUploadsOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UploadPart : Upload a part of an object
// Uploads a part in a multipart upload. In this operation, you provide part data in your request. However, you have an
// option to specify your existing IBM COS object as a data source for the part you are uploading. To upload a part from
// an existing object, you use the `UploadPartCopy` operation. You must initiate a multipart upload (see
// `CreateMultipartUpload`) before you can upload any part. In response to your initiate request, IBM COS returns an
// upload ID (a unique identifier) that you must include in your upload part requests.Part numbers can be any number
// from 1 to 10,000, inclusive. A part number uniquely identifies a part and also defines its position within the object
// being created. If you upload a new part using the same part number that was used with a previous part, the previously
// uploaded part is overwritten. Each part must be at least 5 MB in size, except the last part. There is no size limit
// on the last part of your multipart upload.
//
// To ensure that data is not corrupted when traversing the network, specify the `Content-MD5` header in the upload part
// request. IBM COS checks the part data against the provided MD5 value. If they do not match, IBM COS returns an error.
// If the upload request uses HMAC authentication (AWS Signature Version 4), then IBM COS uses the
// `x-amz-content-sha256` header as a checksum instead of `Content-MD5`.
//
// After you initiate multipart upload and upload one or more parts, you must either complete or abort multipart upload
// in order to stop getting charged for storage of the uploaded parts. Only after you either complete or abort multipart
// upload, IBM COS frees up the parts storage and stops charging you for the parts storage.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) UploadPart(uploadPartOptions *UploadPartOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.UploadPartWithContext(context.Background(), uploadPartOptions)
}

// UploadPartWithContext is an alternate form of the UploadPart method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) UploadPartWithContext(ctx context.Context, uploadPartOptions *UploadPartOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uploadPartOptions, "uploadPartOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uploadPartOptions, "uploadPartOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *uploadPartOptions.Bucket,
		"Key": *uploadPartOptions.Key,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?partNumber={partNumber}&uploadId={uploadId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uploadPartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "UploadPart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")
	if uploadPartOptions.ContentLength != nil {
		builder.AddHeader("Content-Length", fmt.Sprint(*uploadPartOptions.ContentLength))
	}
	if uploadPartOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*uploadPartOptions.ContentMD5))
	}
	if uploadPartOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*uploadPartOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if uploadPartOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*uploadPartOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if uploadPartOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*uploadPartOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}
	if uploadPartOptions.XAmzRequestPayer != nil {
		builder.AddHeader("x-amz-request-payer", fmt.Sprint(*uploadPartOptions.XAmzRequestPayer))
	}
	if uploadPartOptions.XAmzExpectedBucketOwner != nil {
		builder.AddHeader("x-amz-expected-bucket-owner", fmt.Sprint(*uploadPartOptions.XAmzExpectedBucketOwner))
	}

	builder.AddQuery("partNumber", fmt.Sprint(*uploadPartOptions.PartNumber))
	builder.AddQuery("uploadId", fmt.Sprint(*uploadPartOptions.UploadID))

	_, err = builder.SetBodyContent("text/xml", nil, nil, uploadPartOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// UploadPartCopy : Use an existing object as a part of a new object
// Uploads a part by copying data from an existing object as data source. You specify the data source by adding the
// request header `x-amz-copy-source` in your request and a byte range by adding the request header
// `x-amz-copy-source-range` in your request. The minimum allowable part size for a multipart upload is 5 MB. Instead of
// using an existing object as part data, you might use the `UploadPart`operation and provide data in your request.
//
// You must initiate a multipart upload before you can upload any part. In response to your initiate request. IBM COS
// returns a unique identifier, the upload ID, that you must include in your upload part request.
//
// Note the following additional considerations about the request headers `x-amz-copy-source-if-match`,
// `x-amz-copy-source-if-none-match`, `x-amz-copy-source-if-unmodified-since`, and
// `x-amz-copy-source-if-modified-since`:
// * If both of the `x-amz-copy-source-if-match` and `x-amz-copy-source-if-unmodified-since` headers are present in the
// request as follows: `x-amz-copy-source-if-match` condition evaluates to `true`, and
// `x-amz-copy-source-if-unmodified-since` condition evaluates to `false`, then IBM COS returns `200 OK` and copies the
// data.
// * If both of the `x-amz-copy-source-if-none-match` and `x-amz-copy-source-if-modified-since` headers are present in
// the request as follows: `x-amz-copy-source-if-none-match` condition evaluates to `false`, and
// `x-amz-copy-source-if-modified-since` condition evaluates to `true` IBM COS returns `412 Precondition Failed`
// response code.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) UploadPartCopy(uploadPartCopyOptions *UploadPartCopyOptions) (result *UploadPartCopyOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.UploadPartCopyWithContext(context.Background(), uploadPartCopyOptions)
}

// UploadPartCopyWithContext is an alternate form of the UploadPartCopy method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) UploadPartCopyWithContext(ctx context.Context, uploadPartCopyOptions *UploadPartCopyOptions) (result *UploadPartCopyOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uploadPartCopyOptions, "uploadPartCopyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uploadPartCopyOptions, "uploadPartCopyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *uploadPartCopyOptions.Bucket,
		"Key": *uploadPartCopyOptions.Key,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{TargetKey}?partNumber={partNumber}&uploadId={uploadId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uploadPartCopyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "UploadPartCopy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	if uploadPartCopyOptions.XAmzCopySource != nil {
		builder.AddHeader("x-amz-copy-source", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySource))
	}
	if uploadPartCopyOptions.XAmzCopySourceIfMatch != nil {
		builder.AddHeader("x-amz-copy-source-if-match", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceIfMatch))
	}
	if uploadPartCopyOptions.XAmzCopySourceIfModifiedSince != nil {
		builder.AddHeader("x-amz-copy-source-if-modified-since", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceIfModifiedSince))
	}
	if uploadPartCopyOptions.XAmzCopySourceIfNoneMatch != nil {
		builder.AddHeader("x-amz-copy-source-if-none-match", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceIfNoneMatch))
	}
	if uploadPartCopyOptions.XAmzCopySourceIfUnmodifiedSince != nil {
		builder.AddHeader("x-amz-copy-source-if-unmodified-since", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceIfUnmodifiedSince))
	}
	if uploadPartCopyOptions.XAmzCopySourceRange != nil {
		builder.AddHeader("x-amz-copy-source-range", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceRange))
	}
	if uploadPartCopyOptions.XAmzServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-algorithm", fmt.Sprint(*uploadPartCopyOptions.XAmzServerSideEncryptionCustomerAlgorithm))
	}
	if uploadPartCopyOptions.XAmzServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key", fmt.Sprint(*uploadPartCopyOptions.XAmzServerSideEncryptionCustomerKey))
	}
	if uploadPartCopyOptions.XAmzServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-server-side-encryption-customer-key-MD5", fmt.Sprint(*uploadPartCopyOptions.XAmzServerSideEncryptionCustomerKeyMD5))
	}
	if uploadPartCopyOptions.XAmzCopySourceServerSideEncryptionCustomerAlgorithm != nil {
		builder.AddHeader("x-amz-copy-source-server-side-encryption-customer-algorithm", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceServerSideEncryptionCustomerAlgorithm))
	}
	if uploadPartCopyOptions.XAmzCopySourceServerSideEncryptionCustomerKey != nil {
		builder.AddHeader("x-amz-copy-source-server-side-encryption-customer-key", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceServerSideEncryptionCustomerKey))
	}
	if uploadPartCopyOptions.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 != nil {
		builder.AddHeader("x-amz-copy-source-server-side-encryption-customer-key-MD5", fmt.Sprint(*uploadPartCopyOptions.XAmzCopySourceServerSideEncryptionCustomerKeyMD5))
	}

	builder.AddQuery("partNumber", fmt.Sprint(*uploadPartCopyOptions.PartNumber))
	builder.AddQuery("uploadId", fmt.Sprint(*uploadPartCopyOptions.UploadID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUploadPartCopyOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PutPublicAccessBlock : Create a public ACL block configuration
// Creates or modifies the `PublicAccessBlock` configuration for an IBM COS bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutPublicAccessBlock(putPublicAccessBlockOptions *PutPublicAccessBlockOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutPublicAccessBlockWithContext(context.Background(), putPublicAccessBlockOptions)
}

// PutPublicAccessBlockWithContext is an alternate form of the PutPublicAccessBlock method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutPublicAccessBlockWithContext(ctx context.Context, putPublicAccessBlockOptions *PutPublicAccessBlockOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putPublicAccessBlockOptions, "putPublicAccessBlockOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putPublicAccessBlockOptions, "putPublicAccessBlockOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putPublicAccessBlockOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?publicAccessBlock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putPublicAccessBlockOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutPublicAccessBlock")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "text/xml")
	if putPublicAccessBlockOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*putPublicAccessBlockOptions.ContentMD5))
	}

	builder.AddQuery("publicAccessBlock", fmt.Sprint(*putPublicAccessBlockOptions.PublicAccessBlock))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putPublicAccessBlockOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetPublicAccessBlock : Read a public ACL block configuration
// Retrieves the `PublicAccessBlock` configuration for an IBM COS bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetPublicAccessBlock(getPublicAccessBlockOptions *GetPublicAccessBlockOptions) (result *GetPublicAccessBlockOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetPublicAccessBlockWithContext(context.Background(), getPublicAccessBlockOptions)
}

// GetPublicAccessBlockWithContext is an alternate form of the GetPublicAccessBlock method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetPublicAccessBlockWithContext(ctx context.Context, getPublicAccessBlockOptions *GetPublicAccessBlockOptions) (result *GetPublicAccessBlockOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPublicAccessBlockOptions, "getPublicAccessBlockOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPublicAccessBlockOptions, "getPublicAccessBlockOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getPublicAccessBlockOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?publicAccessBlock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPublicAccessBlockOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetPublicAccessBlock")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("publicAccessBlock", fmt.Sprint(*getPublicAccessBlockOptions.PublicAccessBlock))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetPublicAccessBlockOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeletePublicAccessBlock : Delete a public ACL block configuration
// Removes the `PublicAccessBlock` configuration for an IBM COS bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeletePublicAccessBlock(deletePublicAccessBlockOptions *DeletePublicAccessBlockOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeletePublicAccessBlockWithContext(context.Background(), deletePublicAccessBlockOptions)
}

// DeletePublicAccessBlockWithContext is an alternate form of the DeletePublicAccessBlock method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeletePublicAccessBlockWithContext(ctx context.Context, deletePublicAccessBlockOptions *DeletePublicAccessBlockOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePublicAccessBlockOptions, "deletePublicAccessBlockOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePublicAccessBlockOptions, "deletePublicAccessBlockOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deletePublicAccessBlockOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?publicAccessBlock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePublicAccessBlockOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeletePublicAccessBlock")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("publicAccessBlock", fmt.Sprint(*deletePublicAccessBlockOptions.PublicAccessBlock))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetBucketAcl : Read a bucket ACL
// This implementation of the `GET` operation uses the `acl` subresource to return the access control list (ACL) of a
// bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketAcl(getBucketAclOptions *GetBucketAclOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetBucketAclWithContext(context.Background(), getBucketAclOptions)
}

// GetBucketAclWithContext is an alternate form of the GetBucketAcl method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketAclWithContext(ctx context.Context, getBucketAclOptions *GetBucketAclOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketAclOptions, "getBucketAclOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketAclOptions, "getBucketAclOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getBucketAclOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?acl`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketAclOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetBucketAcl")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("acl", fmt.Sprint(*getBucketAclOptions.Acl))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// PutBucketAcl : Create a bucket ACL
// This operation should not be used.  Instead, use IAM policies to grant public access.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketAcl(putBucketAclOptions *PutBucketAclOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutBucketAclWithContext(context.Background(), putBucketAclOptions)
}

// PutBucketAclWithContext is an alternate form of the PutBucketAcl method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketAclWithContext(ctx context.Context, putBucketAclOptions *PutBucketAclOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putBucketAclOptions, "putBucketAclOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putBucketAclOptions, "putBucketAclOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putBucketAclOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?acl`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putBucketAclOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutBucketAcl")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if putBucketAclOptions.XAmzAcl != nil {
		builder.AddHeader("x-amz-acl", fmt.Sprint(*putBucketAclOptions.XAmzAcl))
	}

	builder.AddQuery("acl", fmt.Sprint(*putBucketAclOptions.Acl))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetObjectAcl : Read an object ACL
// This operation should not be used.  Instead, use IAM policies to grant public access.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetObjectAcl(getObjectAclOptions *GetObjectAclOptions) (result *GetObjectAclOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetObjectAclWithContext(context.Background(), getObjectAclOptions)
}

// GetObjectAclWithContext is an alternate form of the GetObjectAcl method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetObjectAclWithContext(ctx context.Context, getObjectAclOptions *GetObjectAclOptions) (result *GetObjectAclOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getObjectAclOptions, "getObjectAclOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getObjectAclOptions, "getObjectAclOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getObjectAclOptions.Bucket,
		"Key": *getObjectAclOptions.Key,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?acl`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getObjectAclOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetObjectAcl")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("acl", fmt.Sprint(*getObjectAclOptions.Acl))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetObjectAclOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// PutObjectAcl : Make an object publicly accessible
// This operation should not be used.  Instead, use IAM policies to grant public access. This operation can be used to
// make a single object in a bucket publicly accessible, but it is discouraged.  Instead, create a new bucket with a
// Public Access policy and copy the object to that bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutObjectAcl(putObjectAclOptions *PutObjectAclOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutObjectAclWithContext(context.Background(), putObjectAclOptions)
}

// PutObjectAclWithContext is an alternate form of the PutObjectAcl method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutObjectAclWithContext(ctx context.Context, putObjectAclOptions *PutObjectAclOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putObjectAclOptions, "putObjectAclOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putObjectAclOptions, "putObjectAclOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putObjectAclOptions.Bucket,
		"Key": *putObjectAclOptions.Key,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?acl`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putObjectAclOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutObjectAcl")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")
	if putObjectAclOptions.XAmzAcl != nil {
		builder.AddHeader("x-amz-acl", fmt.Sprint(*putObjectAclOptions.XAmzAcl))
	}
	if putObjectAclOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*putObjectAclOptions.ContentMD5))
	}
	if putObjectAclOptions.XAmzGrantFullControl != nil {
		builder.AddHeader("x-amz-grant-full-control", fmt.Sprint(*putObjectAclOptions.XAmzGrantFullControl))
	}
	if putObjectAclOptions.XAmzGrantRead != nil {
		builder.AddHeader("x-amz-grant-read", fmt.Sprint(*putObjectAclOptions.XAmzGrantRead))
	}
	if putObjectAclOptions.XAmzGrantReadAcp != nil {
		builder.AddHeader("x-amz-grant-read-acp", fmt.Sprint(*putObjectAclOptions.XAmzGrantReadAcp))
	}
	if putObjectAclOptions.XAmzGrantWrite != nil {
		builder.AddHeader("x-amz-grant-write", fmt.Sprint(*putObjectAclOptions.XAmzGrantWrite))
	}
	if putObjectAclOptions.XAmzGrantWriteAcp != nil {
		builder.AddHeader("x-amz-grant-write-acp", fmt.Sprint(*putObjectAclOptions.XAmzGrantWriteAcp))
	}

	builder.AddQuery("acl", fmt.Sprint(*putObjectAclOptions.Acl))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putObjectAclOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// PutObjectTagging : Add a set of tags to an object
// Sets tags on an object that already exists. A tag is a key-value pair.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutObjectTagging(putObjectTaggingOptions *PutObjectTaggingOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutObjectTaggingWithContext(context.Background(), putObjectTaggingOptions)
}

// PutObjectTaggingWithContext is an alternate form of the PutObjectTagging method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutObjectTaggingWithContext(ctx context.Context, putObjectTaggingOptions *PutObjectTaggingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putObjectTaggingOptions, "putObjectTaggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putObjectTaggingOptions, "putObjectTaggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putObjectTaggingOptions.Bucket,
		"Key": *putObjectTaggingOptions.Key,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?tagging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putObjectTaggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutObjectTagging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")
	builder.AddHeader("Content-Type", "text/xml")
	if putObjectTaggingOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*putObjectTaggingOptions.ContentMD5))
	}

	builder.AddQuery("tagging", fmt.Sprint(*putObjectTaggingOptions.Tagging))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putObjectTaggingOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetObjectTagging : Read a set of object tags
// Returns the tags of an object.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetObjectTagging(getObjectTaggingOptions *GetObjectTaggingOptions) (result *GetObjectTaggingOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetObjectTaggingWithContext(context.Background(), getObjectTaggingOptions)
}

// GetObjectTaggingWithContext is an alternate form of the GetObjectTagging method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetObjectTaggingWithContext(ctx context.Context, getObjectTaggingOptions *GetObjectTaggingOptions) (result *GetObjectTaggingOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getObjectTaggingOptions, "getObjectTaggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getObjectTaggingOptions, "getObjectTaggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getObjectTaggingOptions.Bucket,
		"Key": *getObjectTaggingOptions.Key,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?tagging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getObjectTaggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetObjectTagging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("tagging", fmt.Sprint(*getObjectTaggingOptions.Tagging))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetObjectTaggingOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteObjectTagging : Delete a set of object tags
// Removes the entire tag set from the specified object.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteObjectTagging(deleteObjectTaggingOptions *DeleteObjectTaggingOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteObjectTaggingWithContext(context.Background(), deleteObjectTaggingOptions)
}

// DeleteObjectTaggingWithContext is an alternate form of the DeleteObjectTagging method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteObjectTaggingWithContext(ctx context.Context, deleteObjectTaggingOptions *DeleteObjectTaggingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteObjectTaggingOptions, "deleteObjectTaggingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteObjectTaggingOptions, "deleteObjectTaggingOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteObjectTaggingOptions.Bucket,
		"Key": *deleteObjectTaggingOptions.Key,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}/{Key}?tagging`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteObjectTaggingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteObjectTagging")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("tagging", fmt.Sprint(*deleteObjectTaggingOptions.Tagging))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// PutBucketWebsite : Create a website configuration
// Sets the configuration of the website that is specified in the `website` subresource. To configure a bucket as a
// website, you can add this subresource on the bucket with website configuration information such as the file name of
// the index document and any redirect rules.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketWebsite(putBucketWebsiteOptions *PutBucketWebsiteOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutBucketWebsiteWithContext(context.Background(), putBucketWebsiteOptions)
}

// PutBucketWebsiteWithContext is an alternate form of the PutBucketWebsite method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketWebsiteWithContext(ctx context.Context, putBucketWebsiteOptions *PutBucketWebsiteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putBucketWebsiteOptions, "putBucketWebsiteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putBucketWebsiteOptions, "putBucketWebsiteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putBucketWebsiteOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?website`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putBucketWebsiteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutBucketWebsite")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "text/xml")
	if putBucketWebsiteOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*putBucketWebsiteOptions.ContentMD5))
	}

	builder.AddQuery("website", fmt.Sprint(*putBucketWebsiteOptions.Website))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putBucketWebsiteOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetBucketWebsite : Read a website configuration
// Returns the website configuration for a bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketWebsite(getBucketWebsiteOptions *GetBucketWebsiteOptions) (result *GetBucketWebsiteOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetBucketWebsiteWithContext(context.Background(), getBucketWebsiteOptions)
}

// GetBucketWebsiteWithContext is an alternate form of the GetBucketWebsite method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketWebsiteWithContext(ctx context.Context, getBucketWebsiteOptions *GetBucketWebsiteOptions) (result *GetBucketWebsiteOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketWebsiteOptions, "getBucketWebsiteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketWebsiteOptions, "getBucketWebsiteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getBucketWebsiteOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?website`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketWebsiteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetBucketWebsite")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("website", fmt.Sprint(*getBucketWebsiteOptions.Website))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBucketWebsiteOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteBucketWebsite : Delete a website configuration
// This operation removes the website configuration for a bucket. IBM COS returns a `200 OK` response upon successfully
// deleting a website configuration on the specified bucket. You will get a `200 OK` response if the website
// configuration you are trying to delete does not exist on the bucket. IBM COS returns a `404` response if the bucket
// specified in the request does not exist.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketWebsite(deleteBucketWebsiteOptions *DeleteBucketWebsiteOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteBucketWebsiteWithContext(context.Background(), deleteBucketWebsiteOptions)
}

// DeleteBucketWebsiteWithContext is an alternate form of the DeleteBucketWebsite method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketWebsiteWithContext(ctx context.Context, deleteBucketWebsiteOptions *DeleteBucketWebsiteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBucketWebsiteOptions, "deleteBucketWebsiteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBucketWebsiteOptions, "deleteBucketWebsiteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteBucketWebsiteOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?website`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBucketWebsiteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteBucketWebsite")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("website", fmt.Sprint(*deleteBucketWebsiteOptions.Website))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// PutBucketCors : Configure CORS
// Sets the CORS configuration for your bucket. If the configuration exists, it will be overwritten and replaced.
//
// You set this configuration on a bucket so that the bucket can service cross-origin requests. For example, you might
// want to enable a request whose origin is `http://www.example.com` to access your bucket at `my.example.bucket.com` by
// using `XMLHttpRequest` in a browser.
//
// To enable cross-origin resource sharing (CORS) on a bucket, you create a XML configuration in which you configure
// rules that identify origins and the HTTP methods that can be executed on your bucket. The document is limited to 64
// KB in size. When IBM COS receives a cross-origin request (or a pre-flight OPTIONS request) against a bucket, it
// evaluates the CORS configuration on the bucket and uses the first `CORSRule` rule that matches the incoming browser
// request to enable a cross-origin request. For a rule to match, the following conditions must be met:
//
// * The request's `Origin` header must match `AllowedOrigin` elements.
// * The request method (for example, GET, PUT, HEAD, and so on) or the `Access-Control-Request-Method` header in case
// of a pre-flight `OPTIONS` request must be one of the `AllowedMethod` elements.
// * Every header specified in the `Access-Control-Request-Headers` request header of a pre-flight request must match an
// `AllowedHeader` element.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketCors(putBucketCorsOptions *PutBucketCorsOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.PutBucketCorsWithContext(context.Background(), putBucketCorsOptions)
}

// PutBucketCorsWithContext is an alternate form of the PutBucketCors method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) PutBucketCorsWithContext(ctx context.Context, putBucketCorsOptions *PutBucketCorsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putBucketCorsOptions, "putBucketCorsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putBucketCorsOptions, "putBucketCorsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *putBucketCorsOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?cors`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putBucketCorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "PutBucketCors")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "text/xml")
	if putBucketCorsOptions.ContentMD5 != nil {
		builder.AddHeader("Content-MD5", fmt.Sprint(*putBucketCorsOptions.ContentMD5))
	}

	builder.AddQuery("cors", fmt.Sprint(*putBucketCorsOptions.Cors))

	_, err = builder.SetBodyContent("text/xml", nil, nil, putBucketCorsOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// GetBucketCors : Read a CORS configuration
// Returns the CORS configuration information set for the bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketCors(getBucketCorsOptions *GetBucketCorsOptions) (result *GetBucketCorsOutput, response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.GetBucketCorsWithContext(context.Background(), getBucketCorsOptions)
}

// GetBucketCorsWithContext is an alternate form of the GetBucketCors method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) GetBucketCorsWithContext(ctx context.Context, getBucketCorsOptions *GetBucketCorsOptions) (result *GetBucketCorsOutput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBucketCorsOptions, "getBucketCorsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBucketCorsOptions, "getBucketCorsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *getBucketCorsOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?cors`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBucketCorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "GetBucketCors")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/xml")

	builder.AddQuery("cors", fmt.Sprint(*getBucketCorsOptions.Cors))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBucketCorsOutput)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteBucketCors : Delete a CORS configuration
// Deletes the CORS configuration information set for the bucket.
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketCors(deleteBucketCorsOptions *DeleteBucketCorsOptions) (response *core.DetailedResponse, err error) {
	return ibmCloudObjectStorageS3Api.DeleteBucketCorsWithContext(context.Background(), deleteBucketCorsOptions)
}

// DeleteBucketCorsWithContext is an alternate form of the DeleteBucketCors method which supports a Context parameter
func (ibmCloudObjectStorageS3Api *IbmCloudObjectStorageS3ApiV2) DeleteBucketCorsWithContext(ctx context.Context, deleteBucketCorsOptions *DeleteBucketCorsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteBucketCorsOptions, "deleteBucketCorsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteBucketCorsOptions, "deleteBucketCorsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"Bucket": *deleteBucketCorsOptions.Bucket,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudObjectStorageS3Api.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudObjectStorageS3Api.Service.Options.URL, `/{Bucket}?cors`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteBucketCorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_object_storage_s3_api", "V2", "DeleteBucketCors")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("cors", fmt.Sprint(*deleteBucketCorsOptions.Cors))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = ibmCloudObjectStorageS3Api.Service.Request(request, nil)

	return
}

// AbortMultipartUploadOptions : The AbortMultipartUpload options.
type AbortMultipartUploadOptions struct {
	// The destination bucket for the upload.
	Bucket *string `validate:"required,ne="`

	// Key of the object for which the multipart upload was initiated.
	Key *string `validate:"required,ne="`

	// Upload ID that identifies the multipart upload.
	UploadID *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAbortMultipartUploadOptions : Instantiate AbortMultipartUploadOptions
func (*IbmCloudObjectStorageS3ApiV2) NewAbortMultipartUploadOptions(bucket string, key string, uploadID string) *AbortMultipartUploadOptions {
	return &AbortMultipartUploadOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		UploadID: core.StringPtr(uploadID),
	}
}

// SetBucket : Allow user to set Bucket
func (options *AbortMultipartUploadOptions) SetBucket(bucket string) *AbortMultipartUploadOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *AbortMultipartUploadOptions) SetKey(key string) *AbortMultipartUploadOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetUploadID : Allow user to set UploadID
func (options *AbortMultipartUploadOptions) SetUploadID(uploadID string) *AbortMultipartUploadOptions {
	options.UploadID = core.StringPtr(uploadID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AbortMultipartUploadOptions) SetHeaders(param map[string]string) *AbortMultipartUploadOptions {
	options.Headers = param
	return options
}

// BucketListing : This operation returns a list of all buckets within a service instance.
type BucketListing struct {
	Owner *BucketListingOwner `json:"owner,omitempty"`

	Buckets []BucketListingBucketsItem `json:"buckets,omitempty"`
}

// UnmarshalBucketListing unmarshals an instance of BucketListing from the specified map of raw messages.
func UnmarshalBucketListing(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketListing)
	err = core.UnmarshalModel(m, "owner", &obj.Owner, UnmarshalBucketListingOwner)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "buckets", &obj.Buckets, UnmarshalBucketListingBucketsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketListingBucketsItem : BucketListingBucketsItem struct
type BucketListingBucketsItem struct {
	// Bucket name.
	Name *string `json:"name,omitempty"`

	// Timestamp of bucket creation.
	CreationDate *string `json:"creationDate,omitempty"`
}

// UnmarshalBucketListingBucketsItem unmarshals an instance of BucketListingBucketsItem from the specified map of raw messages.
func UnmarshalBucketListingBucketsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketListingBucketsItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creationDate", &obj.CreationDate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BucketListingOwner : BucketListingOwner struct
type BucketListingOwner struct {
	// Service instance ID.
	ID *string `json:"id,omitempty"`

	// Service instance ID.
	DisplayName *string `json:"displayName,omitempty"`
}

// UnmarshalBucketListingOwner unmarshals an instance of BucketListingOwner from the specified map of raw messages.
func UnmarshalBucketListingOwner(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BucketListingOwner)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "displayName", &obj.DisplayName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CORSRule : Specifies a cross-origin access rule for an IBM COS bucket.
type CORSRule struct {
	// Headers that are specified in the `Access-Control-Request-Headers` header. These headers are allowed in a preflight
	// OPTIONS request. In response to any preflight OPTIONS request, IBM COS returns any requested headers that are
	// allowed.
	AllowedHeaders *CORSRuleAllowedHeaders `json:"AllowedHeaders,omitempty"`

	// An HTTP method that you allow the origin to execute. Valid values are `GET`, `PUT`, `HEAD`, `POST`, and `DELETE`.
	AllowedMethods *CORSRuleAllowedMethods `json:"AllowedMethods" validate:"required"`

	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins *CORSRuleAllowedOrigins `json:"AllowedOrigins" validate:"required"`

	// One or more headers in the response that you want customers to be able to access from their applications (for
	// example, from a JavaScript `XMLHttpRequest` object).
	ExposeHeaders *CORSRuleExposeHeaders `json:"ExposeHeaders,omitempty"`

	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAgeSeconds *int64 `json:"MaxAgeSeconds,omitempty"`
}

// UnmarshalCORSRule unmarshals an instance of CORSRule from the specified map of raw messages.
func UnmarshalCORSRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CORSRule)
	err = core.UnmarshalModel(m, "AllowedHeaders", &obj.AllowedHeaders, UnmarshalCORSRuleAllowedHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "AllowedMethods", &obj.AllowedMethods, UnmarshalCORSRuleAllowedMethods)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "AllowedOrigins", &obj.AllowedOrigins, UnmarshalCORSRuleAllowedOrigins)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ExposeHeaders", &obj.ExposeHeaders, UnmarshalCORSRuleExposeHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MaxAgeSeconds", &obj.MaxAgeSeconds)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CommonPrefix : Container for all (if there are any) keys between Prefix and the next occurrence of the string specified by a
// delimiter. CommonPrefixes lists keys that act like subdirectories in the directory specified by Prefix. For example,
// if the prefix is notes/ and the delimiter is a slash (/) as in notes/summer/july, the common prefix is notes/summer/.
type CommonPrefix struct {
	// Container for the specified common prefix.
	Prefix *string `json:"Prefix,omitempty"`
}

// UnmarshalCommonPrefix unmarshals an instance of CommonPrefix from the specified map of raw messages.
func UnmarshalCommonPrefix(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CommonPrefix)
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CompleteMultipartUploadOptions : The CompleteMultipartUpload options.
type CompleteMultipartUploadOptions struct {
	// Name of the bucket to which the multipart upload was initiated.
	Bucket *string `validate:"required,ne="`

	// Object key for which the multipart upload was initiated.
	Key *string `validate:"required,ne="`

	// ID for the initiated multipart upload.
	UploadID *string `validate:"required"`

	Body *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCompleteMultipartUploadOptions : Instantiate CompleteMultipartUploadOptions
func (*IbmCloudObjectStorageS3ApiV2) NewCompleteMultipartUploadOptions(bucket string, key string, uploadID string, body string) *CompleteMultipartUploadOptions {
	return &CompleteMultipartUploadOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		UploadID: core.StringPtr(uploadID),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *CompleteMultipartUploadOptions) SetBucket(bucket string) *CompleteMultipartUploadOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *CompleteMultipartUploadOptions) SetKey(key string) *CompleteMultipartUploadOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetUploadID : Allow user to set UploadID
func (options *CompleteMultipartUploadOptions) SetUploadID(uploadID string) *CompleteMultipartUploadOptions {
	options.UploadID = core.StringPtr(uploadID)
	return options
}

// SetBody : Allow user to set Body
func (options *CompleteMultipartUploadOptions) SetBody(body string) *CompleteMultipartUploadOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CompleteMultipartUploadOptions) SetHeaders(param map[string]string) *CompleteMultipartUploadOptions {
	options.Headers = param
	return options
}

// CompleteMultipartUploadOutput : CompleteMultipartUploadOutput struct
type CompleteMultipartUploadOutput struct {
	// The URI that identifies the newly created object.
	Location *string `json:"Location,omitempty"`

	// The name of the bucket that contains the newly created object.
	Bucket *string `json:"Bucket,omitempty"`

	// The object key of the newly created object.
	Key *string `json:"Key,omitempty"`

	// Entity tag that identifies the newly created object data. Objects with different object data will have different
	// entity tags. The entity tag is an opaque string. The entity tag is an MD5 digest of the object data, unless it was
	// uploaded using SSE-C (and is randomly generated).
	ETag *string `json:"ETag,omitempty"`
}

// UnmarshalCompleteMultipartUploadOutput unmarshals an instance of CompleteMultipartUploadOutput from the specified map of raw messages.
func UnmarshalCompleteMultipartUploadOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CompleteMultipartUploadOutput)
	err = core.UnmarshalPrimitive(m, "Location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ETag", &obj.ETag)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Condition : A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If
// request is for pages in the `/docs` folder, redirect to the `/documents` folder. 2. If request results in HTTP error
// 4xx, redirect request to another host where you might process the error.
type Condition struct {
	// The HTTP error code when the redirect is applied. In the event of an error, if the error code equals this value,
	// then the specified redirect is applied. Required when parent element `Condition` is specified and sibling
	// `KeyPrefixEquals` is not specified. If both are specified, then both must be true for the redirect to be applied.
	HttpErrorCodeReturnedEquals *string `json:"HttpErrorCodeReturnedEquals,omitempty"`

	// The object key name prefix when the redirect is applied. For example, to redirect requests for `ExamplePage.html`,
	// the key prefix will be `ExamplePage.html`. To redirect request for all pages with the prefix `docs/`, the key prefix
	// will be `/docs`, which identifies all objects in the `docs/` folder. Required when the parent element `Condition` is
	// specified and sibling `HttpErrorCodeReturnedEquals` is not specified. If both conditions are specified, both must be
	// true for the redirect to be applied.
	KeyPrefixEquals *string `json:"KeyPrefixEquals,omitempty"`
}

// UnmarshalCondition unmarshals an instance of Condition from the specified map of raw messages.
func UnmarshalCondition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Condition)
	err = core.UnmarshalPrimitive(m, "HttpErrorCodeReturnedEquals", &obj.HttpErrorCodeReturnedEquals)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "KeyPrefixEquals", &obj.KeyPrefixEquals)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CopyObjectOptions : The CopyObject options.
type CopyObjectOptions struct {
	// The name of the destination bucket.
	Bucket *string `validate:"required,ne="`

	// Specifies the source object for the copy operation.
	XAmzCopySource *string `validate:"required"`

	// The key of the destination object.
	TargetKey *string `validate:"required,ne="`

	Body *string `validate:"required"`

	// The canned ACL to apply to the object.
	XAmzAcl *string

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied
	// to obtain the media-type referenced by the Content-Type header field. For more information, see [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11).
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// Copies the object if its entity tag (ETag) matches the specified tag.
	XAmzCopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time.
	XAmzCopySourceIfModifiedSince *strfmt.DateTime

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	XAmzCopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time.
	XAmzCopySourceIfUnmodifiedSince *strfmt.DateTime

	// The date and time at which the object is no longer cacheable. For more information, [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21).
	Expires *strfmt.DateTime

	// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request.
	XAmzMetadataDirective *string

	// Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the
	// request.
	XAmzTaggingDirective *string

	// The server-side encryption algorithm used when storing this object.
	XAmzServerSideEncryption *string

	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or
	// to an external URL.
	XAmzWebsiteRedirectLocation *string

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for use in encrypting data.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Used as a message integrity check to
	// ensure that the encryption key was transmitted without error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// Specifies the algorithm to use when decrypting the source object (for example, AES256).
	XAmzCopySourceServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use to decrypt the source object. The encryption key
	// provided in this header must be one that was used when the source object was created.
	XAmzCopySourceServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Used as a message integrity check to
	// ensure that the encryption key was transmitted without error.
	XAmzCopySourceServerSideEncryptionCustomerKeyMD5 *string

	// The tag-set for the object destination object this value must be used in conjunction with the `TaggingDirective`.
	// The tag-set must be encoded as URL Query parameters.
	XAmzTagging *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CopyObjectOptions.XAmzAcl property.
// The canned ACL to apply to the object.
const (
	CopyObjectOptions_XAmzAcl_Private = "private"
	CopyObjectOptions_XAmzAcl_PublicRead = "public-read"
)

// Constants associated with the CopyObjectOptions.XAmzMetadataDirective property.
// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request.
const (
	CopyObjectOptions_XAmzMetadataDirective_Copy = "COPY"
	CopyObjectOptions_XAmzMetadataDirective_Replace = "REPLACE"
)

// Constants associated with the CopyObjectOptions.XAmzTaggingDirective property.
// Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the
// request.
const (
	CopyObjectOptions_XAmzTaggingDirective_Copy = "COPY"
	CopyObjectOptions_XAmzTaggingDirective_Replace = "REPLACE"
)

// Constants associated with the CopyObjectOptions.XAmzServerSideEncryption property.
// The server-side encryption algorithm used when storing this object.
const (
	CopyObjectOptions_XAmzServerSideEncryption_Aes256 = "AES256"
)

// NewCopyObjectOptions : Instantiate CopyObjectOptions
func (*IbmCloudObjectStorageS3ApiV2) NewCopyObjectOptions(bucket string, xAmzCopySource string, targetKey string, body string) *CopyObjectOptions {
	return &CopyObjectOptions{
		Bucket: core.StringPtr(bucket),
		XAmzCopySource: core.StringPtr(xAmzCopySource),
		TargetKey: core.StringPtr(targetKey),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *CopyObjectOptions) SetBucket(bucket string) *CopyObjectOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetXAmzCopySource : Allow user to set XAmzCopySource
func (options *CopyObjectOptions) SetXAmzCopySource(xAmzCopySource string) *CopyObjectOptions {
	options.XAmzCopySource = core.StringPtr(xAmzCopySource)
	return options
}

// SetTargetKey : Allow user to set TargetKey
func (options *CopyObjectOptions) SetTargetKey(targetKey string) *CopyObjectOptions {
	options.TargetKey = core.StringPtr(targetKey)
	return options
}

// SetBody : Allow user to set Body
func (options *CopyObjectOptions) SetBody(body string) *CopyObjectOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetXAmzAcl : Allow user to set XAmzAcl
func (options *CopyObjectOptions) SetXAmzAcl(xAmzAcl string) *CopyObjectOptions {
	options.XAmzAcl = core.StringPtr(xAmzAcl)
	return options
}

// SetCacheControl : Allow user to set CacheControl
func (options *CopyObjectOptions) SetCacheControl(cacheControl string) *CopyObjectOptions {
	options.CacheControl = core.StringPtr(cacheControl)
	return options
}

// SetContentDisposition : Allow user to set ContentDisposition
func (options *CopyObjectOptions) SetContentDisposition(contentDisposition string) *CopyObjectOptions {
	options.ContentDisposition = core.StringPtr(contentDisposition)
	return options
}

// SetContentEncoding : Allow user to set ContentEncoding
func (options *CopyObjectOptions) SetContentEncoding(contentEncoding string) *CopyObjectOptions {
	options.ContentEncoding = core.StringPtr(contentEncoding)
	return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *CopyObjectOptions) SetContentLanguage(contentLanguage string) *CopyObjectOptions {
	options.ContentLanguage = core.StringPtr(contentLanguage)
	return options
}

// SetXAmzCopySourceIfMatch : Allow user to set XAmzCopySourceIfMatch
func (options *CopyObjectOptions) SetXAmzCopySourceIfMatch(xAmzCopySourceIfMatch string) *CopyObjectOptions {
	options.XAmzCopySourceIfMatch = core.StringPtr(xAmzCopySourceIfMatch)
	return options
}

// SetXAmzCopySourceIfModifiedSince : Allow user to set XAmzCopySourceIfModifiedSince
func (options *CopyObjectOptions) SetXAmzCopySourceIfModifiedSince(xAmzCopySourceIfModifiedSince *strfmt.DateTime) *CopyObjectOptions {
	options.XAmzCopySourceIfModifiedSince = xAmzCopySourceIfModifiedSince
	return options
}

// SetXAmzCopySourceIfNoneMatch : Allow user to set XAmzCopySourceIfNoneMatch
func (options *CopyObjectOptions) SetXAmzCopySourceIfNoneMatch(xAmzCopySourceIfNoneMatch string) *CopyObjectOptions {
	options.XAmzCopySourceIfNoneMatch = core.StringPtr(xAmzCopySourceIfNoneMatch)
	return options
}

// SetXAmzCopySourceIfUnmodifiedSince : Allow user to set XAmzCopySourceIfUnmodifiedSince
func (options *CopyObjectOptions) SetXAmzCopySourceIfUnmodifiedSince(xAmzCopySourceIfUnmodifiedSince *strfmt.DateTime) *CopyObjectOptions {
	options.XAmzCopySourceIfUnmodifiedSince = xAmzCopySourceIfUnmodifiedSince
	return options
}

// SetExpires : Allow user to set Expires
func (options *CopyObjectOptions) SetExpires(expires *strfmt.DateTime) *CopyObjectOptions {
	options.Expires = expires
	return options
}

// SetXAmzMetadataDirective : Allow user to set XAmzMetadataDirective
func (options *CopyObjectOptions) SetXAmzMetadataDirective(xAmzMetadataDirective string) *CopyObjectOptions {
	options.XAmzMetadataDirective = core.StringPtr(xAmzMetadataDirective)
	return options
}

// SetXAmzTaggingDirective : Allow user to set XAmzTaggingDirective
func (options *CopyObjectOptions) SetXAmzTaggingDirective(xAmzTaggingDirective string) *CopyObjectOptions {
	options.XAmzTaggingDirective = core.StringPtr(xAmzTaggingDirective)
	return options
}

// SetXAmzServerSideEncryption : Allow user to set XAmzServerSideEncryption
func (options *CopyObjectOptions) SetXAmzServerSideEncryption(xAmzServerSideEncryption string) *CopyObjectOptions {
	options.XAmzServerSideEncryption = core.StringPtr(xAmzServerSideEncryption)
	return options
}

// SetXAmzWebsiteRedirectLocation : Allow user to set XAmzWebsiteRedirectLocation
func (options *CopyObjectOptions) SetXAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation string) *CopyObjectOptions {
	options.XAmzWebsiteRedirectLocation = core.StringPtr(xAmzWebsiteRedirectLocation)
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *CopyObjectOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *CopyObjectOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *CopyObjectOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *CopyObjectOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *CopyObjectOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *CopyObjectOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetXAmzCopySourceServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzCopySourceServerSideEncryptionCustomerAlgorithm
func (options *CopyObjectOptions) SetXAmzCopySourceServerSideEncryptionCustomerAlgorithm(xAmzCopySourceServerSideEncryptionCustomerAlgorithm string) *CopyObjectOptions {
	options.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzCopySourceServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzCopySourceServerSideEncryptionCustomerKey : Allow user to set XAmzCopySourceServerSideEncryptionCustomerKey
func (options *CopyObjectOptions) SetXAmzCopySourceServerSideEncryptionCustomerKey(xAmzCopySourceServerSideEncryptionCustomerKey string) *CopyObjectOptions {
	options.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr(xAmzCopySourceServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzCopySourceServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzCopySourceServerSideEncryptionCustomerKeyMD5
func (options *CopyObjectOptions) SetXAmzCopySourceServerSideEncryptionCustomerKeyMD5(xAmzCopySourceServerSideEncryptionCustomerKeyMD5 string) *CopyObjectOptions {
	options.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzCopySourceServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetXAmzTagging : Allow user to set XAmzTagging
func (options *CopyObjectOptions) SetXAmzTagging(xAmzTagging string) *CopyObjectOptions {
	options.XAmzTagging = core.StringPtr(xAmzTagging)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CopyObjectOptions) SetHeaders(param map[string]string) *CopyObjectOptions {
	options.Headers = param
	return options
}

// CopyObjectOutput : CopyObjectOutput struct
type CopyObjectOutput struct {
	// Container for all response elements.
	CopyObjectResult *CopyObjectResult `json:"CopyObjectResult,omitempty"`
}

// UnmarshalCopyObjectOutput unmarshals an instance of CopyObjectOutput from the specified map of raw messages.
func UnmarshalCopyObjectOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CopyObjectOutput)
	err = core.UnmarshalModel(m, "CopyObjectResult", &obj.CopyObjectResult, UnmarshalCopyObjectResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CopyObjectResult : Container for all response elements.
type CopyObjectResult struct {
	// Returns the ETag of the new object. The ETag reflects only changes to the contents of an object, not its metadata.
	// The source and destination ETag is identical for a successfully copied object.
	ETag *string `json:"ETag,omitempty"`

	// Returns the date that the object was last modified.
	LastModified *strfmt.DateTime `json:"LastModified,omitempty"`
}

// UnmarshalCopyObjectResult unmarshals an instance of CopyObjectResult from the specified map of raw messages.
func UnmarshalCopyObjectResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CopyObjectResult)
	err = core.UnmarshalPrimitive(m, "ETag", &obj.ETag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LastModified", &obj.LastModified)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CopyPartResult : Container for all response elements.
type CopyPartResult struct {
	// Entity tag of the object.
	ETag *string `json:"ETag,omitempty"`

	// Date and time at which the object was uploaded.
	LastModified *strfmt.DateTime `json:"LastModified,omitempty"`
}

// UnmarshalCopyPartResult unmarshals an instance of CopyPartResult from the specified map of raw messages.
func UnmarshalCopyPartResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CopyPartResult)
	err = core.UnmarshalPrimitive(m, "ETag", &obj.ETag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LastModified", &obj.LastModified)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateBucketOptions : The CreateBucket options.
type CreateBucketOptions struct {
	// The name of the bucket to create.
	Bucket *string `validate:"required,ne="`

	// This header references the service instance where the bucket will be created and to which data usage will be billed.
	// This value can be either the full Cloud Resource Name (CRN) or just the GUID segment that identifies the service
	// instance.
	IbmServiceInstanceID *string `validate:"required"`

	Body *string

	// The algorithm and key size used to for the managed encryption root key. Required if
	// `ibm-sse-kp-customer-root-key-crn` is also present.
	IbmSseKpEncryptionAlgorithm *string

	// The CRN of the root key used to encrypt the bucket. Required if `ibm-sse-kp-encryption-algorithm` is also present.
	IbmSseKpCustomerRootKeyCrn *string

	// The canned ACL to apply to the bucket. This header should not be used - instead create an IAM policy to grant public
	// access to a bucket.
	XAmzAcl *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateBucketOptions.IbmSseKpEncryptionAlgorithm property.
// The algorithm and key size used to for the managed encryption root key. Required if
// `ibm-sse-kp-customer-root-key-crn` is also present.
const (
	CreateBucketOptions_IbmSseKpEncryptionAlgorithm_Aes256 = "AES256"
)

// Constants associated with the CreateBucketOptions.XAmzAcl property.
// The canned ACL to apply to the bucket. This header should not be used - instead create an IAM policy to grant public
// access to a bucket.
const (
	CreateBucketOptions_XAmzAcl_Private = "private"
	CreateBucketOptions_XAmzAcl_PublicRead = "public-read"
)

// NewCreateBucketOptions : Instantiate CreateBucketOptions
func (*IbmCloudObjectStorageS3ApiV2) NewCreateBucketOptions(bucket string, ibmServiceInstanceID string) *CreateBucketOptions {
	return &CreateBucketOptions{
		Bucket: core.StringPtr(bucket),
		IbmServiceInstanceID: core.StringPtr(ibmServiceInstanceID),
	}
}

// SetBucket : Allow user to set Bucket
func (options *CreateBucketOptions) SetBucket(bucket string) *CreateBucketOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetIbmServiceInstanceID : Allow user to set IbmServiceInstanceID
func (options *CreateBucketOptions) SetIbmServiceInstanceID(ibmServiceInstanceID string) *CreateBucketOptions {
	options.IbmServiceInstanceID = core.StringPtr(ibmServiceInstanceID)
	return options
}

// SetBody : Allow user to set Body
func (options *CreateBucketOptions) SetBody(body string) *CreateBucketOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetIbmSseKpEncryptionAlgorithm : Allow user to set IbmSseKpEncryptionAlgorithm
func (options *CreateBucketOptions) SetIbmSseKpEncryptionAlgorithm(ibmSseKpEncryptionAlgorithm string) *CreateBucketOptions {
	options.IbmSseKpEncryptionAlgorithm = core.StringPtr(ibmSseKpEncryptionAlgorithm)
	return options
}

// SetIbmSseKpCustomerRootKeyCrn : Allow user to set IbmSseKpCustomerRootKeyCrn
func (options *CreateBucketOptions) SetIbmSseKpCustomerRootKeyCrn(ibmSseKpCustomerRootKeyCrn string) *CreateBucketOptions {
	options.IbmSseKpCustomerRootKeyCrn = core.StringPtr(ibmSseKpCustomerRootKeyCrn)
	return options
}

// SetXAmzAcl : Allow user to set XAmzAcl
func (options *CreateBucketOptions) SetXAmzAcl(xAmzAcl string) *CreateBucketOptions {
	options.XAmzAcl = core.StringPtr(xAmzAcl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateBucketOptions) SetHeaders(param map[string]string) *CreateBucketOptions {
	options.Headers = param
	return options
}

// CreateMultipartUploadOutput : CreateMultipartUploadOutput struct
type CreateMultipartUploadOutput struct {
	// <p>The name of the bucket to which the multipart upload was initiated. </p> <p>When using this API with an access
	// point, you must direct requests to the access point hostname. The access point hostname takes the form
	// <i>AccessPointName</i>-<i>AccountId</i>.s3-accesspoint.<i>Region</i>.amazonaws.com. When using this operation with
	// an access point through the AWS SDKs, you provide the access point ARN in place of the bucket name. For more
	// information about access point ARNs, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html">Using Access Points</a> in the
	// <i>Amazon Simple Storage Service Developer Guide</i>.</p> <p>When using this API with IBM COS on Outposts, you must
	// direct requests to the S3 on Outposts hostname. The S3 on Outposts hostname takes the form
	// <i>AccessPointName</i>-<i>AccountId</i>.<i>outpostID</i>.s3-outposts.<i>Region</i>.amazonaws.com. When using this
	// operation using S3 on Outposts through the AWS SDKs, you provide the Outposts bucket ARN in place of the bucket
	// name. For more information about S3 on Outposts ARNs, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html">Using S3 on Outposts</a> in the <i>Amazon
	// Simple Storage Service Developer Guide</i>.</p>.
	Bucket *string `json:"Bucket,omitempty"`

	// Object key for which the multipart upload was initiated.
	Key *string `json:"Key,omitempty"`

	// ID for the initiated multipart upload.
	UploadID *string `json:"UploadId,omitempty"`
}

// UnmarshalCreateMultipartUploadOutput unmarshals an instance of CreateMultipartUploadOutput from the specified map of raw messages.
func UnmarshalCreateMultipartUploadOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateMultipartUploadOutput)
	err = core.UnmarshalPrimitive(m, "Bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "UploadId", &obj.UploadID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteBucketCorsOptions : The DeleteBucketCors options.
type DeleteBucketCorsOptions struct {
	// Specifies the bucket whose CORS configuration is being deleted.
	Bucket *string `validate:"required,ne="`

	Cors *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBucketCorsOptions : Instantiate DeleteBucketCorsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteBucketCorsOptions(bucket string, cors bool) *DeleteBucketCorsOptions {
	return &DeleteBucketCorsOptions{
		Bucket: core.StringPtr(bucket),
		Cors: core.BoolPtr(cors),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteBucketCorsOptions) SetBucket(bucket string) *DeleteBucketCorsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetCors : Allow user to set Cors
func (options *DeleteBucketCorsOptions) SetCors(cors bool) *DeleteBucketCorsOptions {
	options.Cors = core.BoolPtr(cors)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBucketCorsOptions) SetHeaders(param map[string]string) *DeleteBucketCorsOptions {
	options.Headers = param
	return options
}

// DeleteBucketLifecycleOptions : The DeleteBucketLifecycle options.
type DeleteBucketLifecycleOptions struct {
	// The bucket name of the lifecycle to delete.
	Bucket *string `validate:"required,ne="`

	Lifecycle *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBucketLifecycleOptions : Instantiate DeleteBucketLifecycleOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteBucketLifecycleOptions(bucket string, lifecycle bool) *DeleteBucketLifecycleOptions {
	return &DeleteBucketLifecycleOptions{
		Bucket: core.StringPtr(bucket),
		Lifecycle: core.BoolPtr(lifecycle),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteBucketLifecycleOptions) SetBucket(bucket string) *DeleteBucketLifecycleOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetLifecycle : Allow user to set Lifecycle
func (options *DeleteBucketLifecycleOptions) SetLifecycle(lifecycle bool) *DeleteBucketLifecycleOptions {
	options.Lifecycle = core.BoolPtr(lifecycle)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBucketLifecycleOptions) SetHeaders(param map[string]string) *DeleteBucketLifecycleOptions {
	options.Headers = param
	return options
}

// DeleteBucketOptions : The DeleteBucket options.
type DeleteBucketOptions struct {
	// The name of the bucket to delete.
	Bucket *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBucketOptions : Instantiate DeleteBucketOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteBucketOptions(bucket string) *DeleteBucketOptions {
	return &DeleteBucketOptions{
		Bucket: core.StringPtr(bucket),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteBucketOptions) SetBucket(bucket string) *DeleteBucketOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBucketOptions) SetHeaders(param map[string]string) *DeleteBucketOptions {
	options.Headers = param
	return options
}

// DeleteBucketWebsiteOptions : The DeleteBucketWebsite options.
type DeleteBucketWebsiteOptions struct {
	// The bucket for which you want to remove the website configuration.
	Bucket *string `validate:"required,ne="`

	Website *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteBucketWebsiteOptions : Instantiate DeleteBucketWebsiteOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteBucketWebsiteOptions(bucket string, website bool) *DeleteBucketWebsiteOptions {
	return &DeleteBucketWebsiteOptions{
		Bucket: core.StringPtr(bucket),
		Website: core.BoolPtr(website),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteBucketWebsiteOptions) SetBucket(bucket string) *DeleteBucketWebsiteOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetWebsite : Allow user to set Website
func (options *DeleteBucketWebsiteOptions) SetWebsite(website bool) *DeleteBucketWebsiteOptions {
	options.Website = core.BoolPtr(website)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteBucketWebsiteOptions) SetHeaders(param map[string]string) *DeleteBucketWebsiteOptions {
	options.Headers = param
	return options
}

// DeleteObjectOptions : The DeleteObject options.
type DeleteObjectOptions struct {
	// The bucket containing the object.
	Bucket *string `validate:"required,ne="`

	// Key name of the object to delete.
	Key *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteObjectOptions : Instantiate DeleteObjectOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteObjectOptions(bucket string, key string) *DeleteObjectOptions {
	return &DeleteObjectOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteObjectOptions) SetBucket(bucket string) *DeleteObjectOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *DeleteObjectOptions) SetKey(key string) *DeleteObjectOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteObjectOptions) SetHeaders(param map[string]string) *DeleteObjectOptions {
	options.Headers = param
	return options
}

// DeleteObjectTaggingOptions : The DeleteObjectTagging options.
type DeleteObjectTaggingOptions struct {
	// The bucket that contains the object.
	Bucket *string `validate:"required,ne="`

	// Name of the object key.
	Key *string `validate:"required,ne="`

	Tagging *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteObjectTaggingOptions : Instantiate DeleteObjectTaggingOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteObjectTaggingOptions(bucket string, key string, tagging bool) *DeleteObjectTaggingOptions {
	return &DeleteObjectTaggingOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Tagging: core.BoolPtr(tagging),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteObjectTaggingOptions) SetBucket(bucket string) *DeleteObjectTaggingOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *DeleteObjectTaggingOptions) SetKey(key string) *DeleteObjectTaggingOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetTagging : Allow user to set Tagging
func (options *DeleteObjectTaggingOptions) SetTagging(tagging bool) *DeleteObjectTaggingOptions {
	options.Tagging = core.BoolPtr(tagging)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteObjectTaggingOptions) SetHeaders(param map[string]string) *DeleteObjectTaggingOptions {
	options.Headers = param
	return options
}

// DeleteObjectsOptions : The DeleteObjects options.
type DeleteObjectsOptions struct {
	// The bucket name containing the objects to delete.
	Bucket *string `validate:"required,ne="`

	Delete *bool `validate:"required"`

	Body *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteObjectsOptions : Instantiate DeleteObjectsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeleteObjectsOptions(bucket string, deleteVar bool, body string) *DeleteObjectsOptions {
	return &DeleteObjectsOptions{
		Bucket: core.StringPtr(bucket),
		Delete: core.BoolPtr(deleteVar),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeleteObjectsOptions) SetBucket(bucket string) *DeleteObjectsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetDelete : Allow user to set Delete
func (options *DeleteObjectsOptions) SetDelete(deleteVar bool) *DeleteObjectsOptions {
	options.Delete = core.BoolPtr(deleteVar)
	return options
}

// SetBody : Allow user to set Body
func (options *DeleteObjectsOptions) SetBody(body string) *DeleteObjectsOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteObjectsOptions) SetHeaders(param map[string]string) *DeleteObjectsOptions {
	options.Headers = param
	return options
}

// DeleteObjectsOutput : DeleteObjectsOutput struct
type DeleteObjectsOutput struct {
	// Container for a failed delete operation that describes the object that IBM COS attempted to delete and the error it
	// encountered.
	Errors *DeleteObjectsOutputErrors `json:"Errors,omitempty"`
}

// UnmarshalDeleteObjectsOutput unmarshals an instance of DeleteObjectsOutput from the specified map of raw messages.
func UnmarshalDeleteObjectsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteObjectsOutput)
	err = core.UnmarshalModel(m, "Errors", &obj.Errors, UnmarshalDeleteObjectsOutputErrors)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeletePublicAccessBlockOptions : The DeletePublicAccessBlock options.
type DeletePublicAccessBlockOptions struct {
	// The IBM COS bucket whose `PublicAccessBlock` configuration you want to delete.
	Bucket *string `validate:"required,ne="`

	PublicAccessBlock *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePublicAccessBlockOptions : Instantiate DeletePublicAccessBlockOptions
func (*IbmCloudObjectStorageS3ApiV2) NewDeletePublicAccessBlockOptions(bucket string, publicAccessBlock bool) *DeletePublicAccessBlockOptions {
	return &DeletePublicAccessBlockOptions{
		Bucket: core.StringPtr(bucket),
		PublicAccessBlock: core.BoolPtr(publicAccessBlock),
	}
}

// SetBucket : Allow user to set Bucket
func (options *DeletePublicAccessBlockOptions) SetBucket(bucket string) *DeletePublicAccessBlockOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetPublicAccessBlock : Allow user to set PublicAccessBlock
func (options *DeletePublicAccessBlockOptions) SetPublicAccessBlock(publicAccessBlock bool) *DeletePublicAccessBlockOptions {
	options.PublicAccessBlock = core.BoolPtr(publicAccessBlock)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePublicAccessBlockOptions) SetHeaders(param map[string]string) *DeletePublicAccessBlockOptions {
	options.Headers = param
	return options
}

// Error : Container for all error elements.
type Error struct {
	// The error key.
	Key *string `json:"Key,omitempty"`

	Code *string `json:"Code,omitempty"`

	// The error message contains a generic description of the error condition in English. It is intended for a human
	// audience. Simple programs display the message directly to the end user if they encounter an error condition they
	// don't know how or don't care to handle. Sophisticated programs with more exhaustive error handling and proper
	// internationalization are more likely to ignore the error message.
	Message *string `json:"Message,omitempty"`
}

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ErrorDocument : The error information.
type ErrorDocument struct {
	// The object key name to use when a 4XX class error occurs.
	Key *string `json:"Key" validate:"required"`
}

// UnmarshalErrorDocument unmarshals an instance of ErrorDocument from the specified map of raw messages.
func UnmarshalErrorDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorDocument)
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketAclOptions : The GetBucketAcl options.
type GetBucketAclOptions struct {
	// Specifies the bucket whose ACL is being requested.
	Bucket *string `validate:"required,ne="`

	Acl *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketAclOptions : Instantiate GetBucketAclOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetBucketAclOptions(bucket string, acl bool) *GetBucketAclOptions {
	return &GetBucketAclOptions{
		Bucket: core.StringPtr(bucket),
		Acl: core.BoolPtr(acl),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetBucketAclOptions) SetBucket(bucket string) *GetBucketAclOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetAcl : Allow user to set Acl
func (options *GetBucketAclOptions) SetAcl(acl bool) *GetBucketAclOptions {
	options.Acl = core.BoolPtr(acl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketAclOptions) SetHeaders(param map[string]string) *GetBucketAclOptions {
	options.Headers = param
	return options
}

// GetBucketCorsOptions : The GetBucketCors options.
type GetBucketCorsOptions struct {
	// The bucket name for which to get the CORS configuration.
	Bucket *string `validate:"required,ne="`

	Cors *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketCorsOptions : Instantiate GetBucketCorsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetBucketCorsOptions(bucket string, cors bool) *GetBucketCorsOptions {
	return &GetBucketCorsOptions{
		Bucket: core.StringPtr(bucket),
		Cors: core.BoolPtr(cors),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetBucketCorsOptions) SetBucket(bucket string) *GetBucketCorsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetCors : Allow user to set Cors
func (options *GetBucketCorsOptions) SetCors(cors bool) *GetBucketCorsOptions {
	options.Cors = core.BoolPtr(cors)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketCorsOptions) SetHeaders(param map[string]string) *GetBucketCorsOptions {
	options.Headers = param
	return options
}

// GetBucketCorsOutput : GetBucketCorsOutput struct
type GetBucketCorsOutput struct {
	// A set of origins and methods (cross-origin access that you want to allow). You can add up to 100 rules to the
	// configuration.
	CORSRules *GetBucketCorsOutputCORSRules `json:"CORSRules,omitempty"`
}

// UnmarshalGetBucketCorsOutput unmarshals an instance of GetBucketCorsOutput from the specified map of raw messages.
func UnmarshalGetBucketCorsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketCorsOutput)
	err = core.UnmarshalModel(m, "CORSRules", &obj.CORSRules, UnmarshalGetBucketCorsOutputCORSRules)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketLifecycleConfigurationOptions : The GetBucketLifecycleConfiguration options.
type GetBucketLifecycleConfigurationOptions struct {
	// The name of the bucket for which to get the lifecycle information.
	Bucket *string `validate:"required,ne="`

	Lifecycle *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketLifecycleConfigurationOptions : Instantiate GetBucketLifecycleConfigurationOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetBucketLifecycleConfigurationOptions(bucket string, lifecycle bool) *GetBucketLifecycleConfigurationOptions {
	return &GetBucketLifecycleConfigurationOptions{
		Bucket: core.StringPtr(bucket),
		Lifecycle: core.BoolPtr(lifecycle),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetBucketLifecycleConfigurationOptions) SetBucket(bucket string) *GetBucketLifecycleConfigurationOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetLifecycle : Allow user to set Lifecycle
func (options *GetBucketLifecycleConfigurationOptions) SetLifecycle(lifecycle bool) *GetBucketLifecycleConfigurationOptions {
	options.Lifecycle = core.BoolPtr(lifecycle)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketLifecycleConfigurationOptions) SetHeaders(param map[string]string) *GetBucketLifecycleConfigurationOptions {
	options.Headers = param
	return options
}

// GetBucketLifecycleConfigurationOutput : GetBucketLifecycleConfigurationOutput struct
type GetBucketLifecycleConfigurationOutput struct {
	// Container for a lifecycle rule.
	Rules *GetBucketLifecycleConfigurationOutputRules `json:"Rules,omitempty"`
}

// UnmarshalGetBucketLifecycleConfigurationOutput unmarshals an instance of GetBucketLifecycleConfigurationOutput from the specified map of raw messages.
func UnmarshalGetBucketLifecycleConfigurationOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketLifecycleConfigurationOutput)
	err = core.UnmarshalModel(m, "Rules", &obj.Rules, UnmarshalGetBucketLifecycleConfigurationOutputRules)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketWebsiteOptions : The GetBucketWebsite options.
type GetBucketWebsiteOptions struct {
	// The bucket name for which to get the website configuration.
	Bucket *string `validate:"required,ne="`

	Website *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBucketWebsiteOptions : Instantiate GetBucketWebsiteOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetBucketWebsiteOptions(bucket string, website bool) *GetBucketWebsiteOptions {
	return &GetBucketWebsiteOptions{
		Bucket: core.StringPtr(bucket),
		Website: core.BoolPtr(website),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetBucketWebsiteOptions) SetBucket(bucket string) *GetBucketWebsiteOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetWebsite : Allow user to set Website
func (options *GetBucketWebsiteOptions) SetWebsite(website bool) *GetBucketWebsiteOptions {
	options.Website = core.BoolPtr(website)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBucketWebsiteOptions) SetHeaders(param map[string]string) *GetBucketWebsiteOptions {
	options.Headers = param
	return options
}

// GetBucketWebsiteOutput : GetBucketWebsiteOutput struct
type GetBucketWebsiteOutput struct {
	// Specifies the redirect behavior of all requests to a website endpoint of an IBM COS bucket.
	RedirectAllRequestsTo *RedirectAllRequestsTo `json:"RedirectAllRequestsTo,omitempty"`

	// The name of the index document for the website (for example `index.html`).
	IndexDocument *IndexDocument `json:"IndexDocument,omitempty"`

	// The object key name of the website error document to use for 4XX class errors.
	ErrorDocument *ErrorDocument `json:"ErrorDocument,omitempty"`

	// Rules that define when a redirect is applied and the redirect behavior.
	RoutingRules *RoutingRules `json:"RoutingRules,omitempty"`
}

// UnmarshalGetBucketWebsiteOutput unmarshals an instance of GetBucketWebsiteOutput from the specified map of raw messages.
func UnmarshalGetBucketWebsiteOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketWebsiteOutput)
	err = core.UnmarshalModel(m, "RedirectAllRequestsTo", &obj.RedirectAllRequestsTo, UnmarshalRedirectAllRequestsTo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "IndexDocument", &obj.IndexDocument, UnmarshalIndexDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ErrorDocument", &obj.ErrorDocument, UnmarshalErrorDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "RoutingRules", &obj.RoutingRules, UnmarshalRoutingRules)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetObjectAclOptions : The GetObjectAcl options.
type GetObjectAclOptions struct {
	// The name of the bucket containing the object being requested.
	Bucket *string `validate:"required,ne="`

	// The key of the object for which to get the ACL information.
	Key *string `validate:"required,ne="`

	Acl *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetObjectAclOptions : Instantiate GetObjectAclOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetObjectAclOptions(bucket string, key string, acl bool) *GetObjectAclOptions {
	return &GetObjectAclOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Acl: core.BoolPtr(acl),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetObjectAclOptions) SetBucket(bucket string) *GetObjectAclOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *GetObjectAclOptions) SetKey(key string) *GetObjectAclOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetAcl : Allow user to set Acl
func (options *GetObjectAclOptions) SetAcl(acl bool) *GetObjectAclOptions {
	options.Acl = core.BoolPtr(acl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetObjectAclOptions) SetHeaders(param map[string]string) *GetObjectAclOptions {
	options.Headers = param
	return options
}

// GetObjectAclOutput : GetObjectAclOutput struct
type GetObjectAclOutput struct {
	// Container for the bucket owner's display name and ID.
	Owner *Owner `json:"Owner,omitempty"`

	// A list of grants.
	Grants *GetObjectAclOutputGrants `json:"Grants,omitempty"`
}

// UnmarshalGetObjectAclOutput unmarshals an instance of GetObjectAclOutput from the specified map of raw messages.
func UnmarshalGetObjectAclOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetObjectAclOutput)
	err = core.UnmarshalModel(m, "Owner", &obj.Owner, UnmarshalOwner)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Grants", &obj.Grants, UnmarshalGetObjectAclOutputGrants)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetObjectOptions : The GetObject options.
type GetObjectOptions struct {
	// The bucket containing the object.
	Bucket *string `validate:"required,ne="`

	// Key of the object to get.
	Key *string `validate:"required,ne="`

	// Returns the object only if its entity tag (ETag) is the same as the one specified, otherwise returns a 412
	// (precondition failed).
	IfMatch *string

	// Returns the object only if it has been modified since the specified time, otherwise returns a 304 (not modified).
	IfModifiedSince *strfmt.DateTime

	// Returns the object only if its entity tag (ETag) is different from the one specified, otherwise returns a 304 (not
	// modified).
	IfNoneMatch *string

	// Returns the object only if it has not been modified since the specified time, otherwise return a 412 (precondition
	// failed).'.
	IfUnmodifiedSince *strfmt.DateTime

	// Downloads the specified range bytes of an object. For more information about the HTTP Range header, see [RFC
	// 2516](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35).
	Range *string

	// Sets the `Cache-Control` header of the response.
	ResponseCacheControl *string

	// Sets the `Content-Disposition` header of the response.
	ResponseContentDisposition *string

	// Sets the `Content-Encoding` header of the response.
	ResponseContentEncoding *string

	// Sets the `Content-Language` header of the response.
	ResponseContentLanguage *string

	// Sets the `Content-Type` header of the response.
	ResponseContentType *string

	// Sets the `Expires` header of the response.
	ResponseExpires *strfmt.DateTime

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use in encrypting data. This value is used to store
	// the object and then it is discarded; IBM COS does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the `x-amz-server-side-encryption-customer-algorithm` header.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IBM COS uses this header for a message
	// integrity check to ensure that the encryption key was transmitted without error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// Part number of the object being read. This is a positive integer between 1 and 10,000. Effectively performs a
	// "ranged" GET request for the part specified. Useful for downloading just a part of an object.
	PartNumber *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetObjectOptions : Instantiate GetObjectOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetObjectOptions(bucket string, key string) *GetObjectOptions {
	return &GetObjectOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetObjectOptions) SetBucket(bucket string) *GetObjectOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *GetObjectOptions) SetKey(key string) *GetObjectOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *GetObjectOptions) SetIfMatch(ifMatch string) *GetObjectOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetIfModifiedSince : Allow user to set IfModifiedSince
func (options *GetObjectOptions) SetIfModifiedSince(ifModifiedSince *strfmt.DateTime) *GetObjectOptions {
	options.IfModifiedSince = ifModifiedSince
	return options
}

// SetIfNoneMatch : Allow user to set IfNoneMatch
func (options *GetObjectOptions) SetIfNoneMatch(ifNoneMatch string) *GetObjectOptions {
	options.IfNoneMatch = core.StringPtr(ifNoneMatch)
	return options
}

// SetIfUnmodifiedSince : Allow user to set IfUnmodifiedSince
func (options *GetObjectOptions) SetIfUnmodifiedSince(ifUnmodifiedSince *strfmt.DateTime) *GetObjectOptions {
	options.IfUnmodifiedSince = ifUnmodifiedSince
	return options
}

// SetRange : Allow user to set Range
func (options *GetObjectOptions) SetRange(rangeVar string) *GetObjectOptions {
	options.Range = core.StringPtr(rangeVar)
	return options
}

// SetResponseCacheControl : Allow user to set ResponseCacheControl
func (options *GetObjectOptions) SetResponseCacheControl(responseCacheControl string) *GetObjectOptions {
	options.ResponseCacheControl = core.StringPtr(responseCacheControl)
	return options
}

// SetResponseContentDisposition : Allow user to set ResponseContentDisposition
func (options *GetObjectOptions) SetResponseContentDisposition(responseContentDisposition string) *GetObjectOptions {
	options.ResponseContentDisposition = core.StringPtr(responseContentDisposition)
	return options
}

// SetResponseContentEncoding : Allow user to set ResponseContentEncoding
func (options *GetObjectOptions) SetResponseContentEncoding(responseContentEncoding string) *GetObjectOptions {
	options.ResponseContentEncoding = core.StringPtr(responseContentEncoding)
	return options
}

// SetResponseContentLanguage : Allow user to set ResponseContentLanguage
func (options *GetObjectOptions) SetResponseContentLanguage(responseContentLanguage string) *GetObjectOptions {
	options.ResponseContentLanguage = core.StringPtr(responseContentLanguage)
	return options
}

// SetResponseContentType : Allow user to set ResponseContentType
func (options *GetObjectOptions) SetResponseContentType(responseContentType string) *GetObjectOptions {
	options.ResponseContentType = core.StringPtr(responseContentType)
	return options
}

// SetResponseExpires : Allow user to set ResponseExpires
func (options *GetObjectOptions) SetResponseExpires(responseExpires *strfmt.DateTime) *GetObjectOptions {
	options.ResponseExpires = responseExpires
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *GetObjectOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *GetObjectOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *GetObjectOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *GetObjectOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *GetObjectOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *GetObjectOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetPartNumber : Allow user to set PartNumber
func (options *GetObjectOptions) SetPartNumber(partNumber int64) *GetObjectOptions {
	options.PartNumber = core.Int64Ptr(partNumber)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetObjectOptions) SetHeaders(param map[string]string) *GetObjectOptions {
	options.Headers = param
	return options
}

// GetObjectOutput : GetObjectOutput struct
type GetObjectOutput struct {
	// Object data.
	Body *string `json:"Body,omitempty"`

	// A map of metadata to store with the object in S3.
	Metadata *GetObjectOutputMetadata `json:"Metadata,omitempty"`
}

// UnmarshalGetObjectOutput unmarshals an instance of GetObjectOutput from the specified map of raw messages.
func UnmarshalGetObjectOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetObjectOutput)
	err = core.UnmarshalPrimitive(m, "Body", &obj.Body)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Metadata", &obj.Metadata, UnmarshalGetObjectOutputMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetObjectOutputMetadata : A map of metadata to store with the object in S3.
type GetObjectOutputMetadata struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]*string
}

// SetProperty allows the user to set an arbitrary property on an instance of GetObjectOutputMetadata
func (o *GetObjectOutputMetadata) SetProperty(key string, value *string) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]*string)
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of GetObjectOutputMetadata
func (o *GetObjectOutputMetadata) GetProperty(key string) *string {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of GetObjectOutputMetadata
func (o *GetObjectOutputMetadata) GetProperties() map[string]*string {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of GetObjectOutputMetadata
func (o *GetObjectOutputMetadata) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalGetObjectOutputMetadata unmarshals an instance of GetObjectOutputMetadata from the specified map of raw messages.
func UnmarshalGetObjectOutputMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetObjectOutputMetadata)
	for k := range m {
		var v *string
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetObjectTaggingOptions : The GetObjectTagging options.
type GetObjectTaggingOptions struct {
	// The bucket containing the object.
	Bucket *string `validate:"required,ne="`

	// Object key for which to get the tagging information.
	Key *string `validate:"required,ne="`

	Tagging *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetObjectTaggingOptions : Instantiate GetObjectTaggingOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetObjectTaggingOptions(bucket string, key string, tagging bool) *GetObjectTaggingOptions {
	return &GetObjectTaggingOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Tagging: core.BoolPtr(tagging),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetObjectTaggingOptions) SetBucket(bucket string) *GetObjectTaggingOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *GetObjectTaggingOptions) SetKey(key string) *GetObjectTaggingOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetTagging : Allow user to set Tagging
func (options *GetObjectTaggingOptions) SetTagging(tagging bool) *GetObjectTaggingOptions {
	options.Tagging = core.BoolPtr(tagging)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetObjectTaggingOptions) SetHeaders(param map[string]string) *GetObjectTaggingOptions {
	options.Headers = param
	return options
}

// GetObjectTaggingOutput : GetObjectTaggingOutput struct
type GetObjectTaggingOutput struct {
	// Contains the tag set.
	TagSet *TagSet `json:"TagSet" validate:"required"`
}

// UnmarshalGetObjectTaggingOutput unmarshals an instance of GetObjectTaggingOutput from the specified map of raw messages.
func UnmarshalGetObjectTaggingOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetObjectTaggingOutput)
	err = core.UnmarshalModel(m, "TagSet", &obj.TagSet, UnmarshalTagSet)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetPublicAccessBlockOptions : The GetPublicAccessBlock options.
type GetPublicAccessBlockOptions struct {
	// The name of the IBM COS bucket whose `PublicAccessBlock` configuration you want to retrieve.
	Bucket *string `validate:"required,ne="`

	PublicAccessBlock *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPublicAccessBlockOptions : Instantiate GetPublicAccessBlockOptions
func (*IbmCloudObjectStorageS3ApiV2) NewGetPublicAccessBlockOptions(bucket string, publicAccessBlock bool) *GetPublicAccessBlockOptions {
	return &GetPublicAccessBlockOptions{
		Bucket: core.StringPtr(bucket),
		PublicAccessBlock: core.BoolPtr(publicAccessBlock),
	}
}

// SetBucket : Allow user to set Bucket
func (options *GetPublicAccessBlockOptions) SetBucket(bucket string) *GetPublicAccessBlockOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetPublicAccessBlock : Allow user to set PublicAccessBlock
func (options *GetPublicAccessBlockOptions) SetPublicAccessBlock(publicAccessBlock bool) *GetPublicAccessBlockOptions {
	options.PublicAccessBlock = core.BoolPtr(publicAccessBlock)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPublicAccessBlockOptions) SetHeaders(param map[string]string) *GetPublicAccessBlockOptions {
	options.Headers = param
	return options
}

// GetPublicAccessBlockOutput : GetPublicAccessBlockOutput struct
type GetPublicAccessBlockOutput struct {
	// The `PublicAccessBlock` configuration currently in effect for this IBM COS bucket.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty"`
}

// UnmarshalGetPublicAccessBlockOutput unmarshals an instance of GetPublicAccessBlockOutput from the specified map of raw messages.
func UnmarshalGetPublicAccessBlockOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetPublicAccessBlockOutput)
	err = core.UnmarshalModel(m, "PublicAccessBlockConfiguration", &obj.PublicAccessBlockConfiguration, UnmarshalPublicAccessBlockConfiguration)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Grantee : Container for the person being granted permissions.
type Grantee struct {
	// Screen name of the grantee.
	DisplayName *string `json:"DisplayName,omitempty"`

	// <p>Email address of the grantee.</p> <note> <p>Using email addresses to specify a grantee is only supported in the
	// following AWS Regions: </p> <ul> <li> <p>US East (N. Virginia)</p> </li> <li> <p>US West (N. California)</p> </li>
	// <li> <p> US West (Oregon)</p> </li> <li> <p> Asia Pacific (Singapore)</p> </li> <li> <p>Asia Pacific (Sydney)</p>
	// </li> <li> <p>Asia Pacific (Tokyo)</p> </li> <li> <p>Europe (Ireland)</p> </li> <li> <p>South America (São
	// Paulo)</p> </li> </ul> <p>For a list of all the IBM COS supported Regions and endpoints, see <a
	// href="https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region">Regions and Endpoints</a> in the AWS
	// General Reference.</p> </note>.
	EmailAddress *string `json:"EmailAddress,omitempty"`

	// The canonical user ID of the grantee.
	ID *string `json:"ID,omitempty"`

	// Type of grantee.
	Type *string `json:"Type" validate:"required"`

	// URI of the grantee group.
	URI *string `json:"URI,omitempty"`
}

// Constants associated with the Grantee.Type property.
// Type of grantee.
const (
	Grantee_Type_Amazoncustomerbyemail = "AmazonCustomerByEmail"
	Grantee_Type_Canonicaluser = "CanonicalUser"
	Grantee_Type_Group = "Group"
)

// UnmarshalGrantee unmarshals an instance of Grantee from the specified map of raw messages.
func UnmarshalGrantee(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Grantee)
	err = core.UnmarshalPrimitive(m, "DisplayName", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "EmailAddress", &obj.EmailAddress)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ID", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "URI", &obj.URI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GrantsItem : Container for grant information.
type GrantsItem struct {
	// The person being granted permissions.
	Grantee *Grantee `json:"Grantee,omitempty"`

	// Specifies the permission given to the grantee.
	Permission *string `json:"Permission,omitempty"`
}

// Constants associated with the GrantsItem.Permission property.
// Specifies the permission given to the grantee.
const (
	GrantsItem_Permission_FullControl = "FULL_CONTROL"
	GrantsItem_Permission_Read = "READ"
	GrantsItem_Permission_ReadAcp = "READ_ACP"
	GrantsItem_Permission_Write = "WRITE"
	GrantsItem_Permission_WriteAcp = "WRITE_ACP"
)

// UnmarshalGrantsItem unmarshals an instance of GrantsItem from the specified map of raw messages.
func UnmarshalGrantsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GrantsItem)
	err = core.UnmarshalModel(m, "Grantee", &obj.Grantee, UnmarshalGrantee)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Permission", &obj.Permission)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HeadBucketOptions : The HeadBucket options.
type HeadBucketOptions struct {
	// The name of the bucket being checked. The bucket must exist in the location specified by the endpoint for the
	// request to succeed.
	Bucket *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewHeadBucketOptions : Instantiate HeadBucketOptions
func (*IbmCloudObjectStorageS3ApiV2) NewHeadBucketOptions(bucket string) *HeadBucketOptions {
	return &HeadBucketOptions{
		Bucket: core.StringPtr(bucket),
	}
}

// SetBucket : Allow user to set Bucket
func (options *HeadBucketOptions) SetBucket(bucket string) *HeadBucketOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *HeadBucketOptions) SetHeaders(param map[string]string) *HeadBucketOptions {
	options.Headers = param
	return options
}

// HeadObjectOptions : The HeadObject options.
type HeadObjectOptions struct {
	// The bucket containing the object.
	Bucket *string `validate:"required,ne="`

	// The object key.
	Key *string `validate:"required,ne="`

	// Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412
	// (precondition failed).
	IfMatch *string

	// Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified).
	IfModifiedSince *strfmt.DateTime

	// Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not
	// modified).
	IfNoneMatch *string

	// Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition
	// failed).
	IfUnmodifiedSince *strfmt.DateTime

	// Downloads the specified range bytes of an object. For more information about the HTTP Range header, see [RFC
	// 2516](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35).
	Range *string

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use in encrypting data. This value is used to store
	// the object and then it is discarded; IBM COS does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the `x-amz-server-side-encryption-customer-algorithm` header.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IBM COS uses this header for a message
	// integrity check to ensure that the encryption key was transmitted without error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// Part number of the object being read. This is a positive integer between 1 and 10,000. Effectively performs a
	// "ranged" HEAD request for the part specified. Useful querying about the size of the part and the number of parts in
	// this object.
	PartNumber *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewHeadObjectOptions : Instantiate HeadObjectOptions
func (*IbmCloudObjectStorageS3ApiV2) NewHeadObjectOptions(bucket string, key string) *HeadObjectOptions {
	return &HeadObjectOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
	}
}

// SetBucket : Allow user to set Bucket
func (options *HeadObjectOptions) SetBucket(bucket string) *HeadObjectOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *HeadObjectOptions) SetKey(key string) *HeadObjectOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *HeadObjectOptions) SetIfMatch(ifMatch string) *HeadObjectOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetIfModifiedSince : Allow user to set IfModifiedSince
func (options *HeadObjectOptions) SetIfModifiedSince(ifModifiedSince *strfmt.DateTime) *HeadObjectOptions {
	options.IfModifiedSince = ifModifiedSince
	return options
}

// SetIfNoneMatch : Allow user to set IfNoneMatch
func (options *HeadObjectOptions) SetIfNoneMatch(ifNoneMatch string) *HeadObjectOptions {
	options.IfNoneMatch = core.StringPtr(ifNoneMatch)
	return options
}

// SetIfUnmodifiedSince : Allow user to set IfUnmodifiedSince
func (options *HeadObjectOptions) SetIfUnmodifiedSince(ifUnmodifiedSince *strfmt.DateTime) *HeadObjectOptions {
	options.IfUnmodifiedSince = ifUnmodifiedSince
	return options
}

// SetRange : Allow user to set Range
func (options *HeadObjectOptions) SetRange(rangeVar string) *HeadObjectOptions {
	options.Range = core.StringPtr(rangeVar)
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *HeadObjectOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *HeadObjectOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *HeadObjectOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *HeadObjectOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *HeadObjectOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *HeadObjectOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetPartNumber : Allow user to set PartNumber
func (options *HeadObjectOptions) SetPartNumber(partNumber int64) *HeadObjectOptions {
	options.PartNumber = core.Int64Ptr(partNumber)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *HeadObjectOptions) SetHeaders(param map[string]string) *HeadObjectOptions {
	options.Headers = param
	return options
}

// IndexDocument : Container for the `Suffix` element.
type IndexDocument struct {
	// A suffix that is appended to a request that is for a directory on the website endpoint (for example,if the suffix is
	// index.html and you make a request to samplebucket/images/ the data that is returned will be for the object with the
	// key name images/index.html) The suffix must not be empty and must not include a slash character.
	Suffix *string `json:"Suffix" validate:"required"`
}

// UnmarshalIndexDocument unmarshals an instance of IndexDocument from the specified map of raw messages.
func UnmarshalIndexDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IndexDocument)
	err = core.UnmarshalPrimitive(m, "Suffix", &obj.Suffix)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InitiateMultipartUploadOptions : The InitiateMultipartUpload options.
type InitiateMultipartUploadOptions struct {
	// The name of the bucket to which to initiate the upload.
	Bucket *string `validate:"required,ne="`

	Uploads *bool `validate:"required"`

	// Object key for which the multipart upload is to be initiated.
	Key *string `validate:"required,ne="`

	Body *string `validate:"required"`

	// Upload the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412
	// (precondition failed).
	IfMatch *string

	// Upload the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not
	// modified).
	IfNoneMatch *string

	// Upload the object only if it has not been modified since the specified time, otherwise return a 412 (precondition
	// failed).
	IfUnmodifiedSince *strfmt.DateTime

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied
	// to obtain the media-type referenced by the Content-Type header field. For more information, see [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11).
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// The date and time at which the object is no longer cacheable. For more information, [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21).
	Expires *strfmt.DateTime

	// The server-side encryption algorithm used when storing this object (for example, AES256).
	XAmzServerSideEncryption *string

	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or
	// to an external URL.
	XAmzWebsiteRedirectLocation *string

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for encrypting data. IBM COS does not store the encryption key - it
	// is discarded use. The key must be appropriate for use with the algorithm specified in the
	// `x-amz-server-side-encryption-customer-algorithm` header.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321, ensuring the encryption key is
	// transmitted without error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// A set of key-value pairs, encoded as URL query parameters.
	XAmzTagging *string

	// The canned ACL to apply to the object.
	XAmzAcl *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the InitiateMultipartUploadOptions.XAmzServerSideEncryption property.
// The server-side encryption algorithm used when storing this object (for example, AES256).
const (
	InitiateMultipartUploadOptions_XAmzServerSideEncryption_Aes256 = "AES256"
)

// Constants associated with the InitiateMultipartUploadOptions.XAmzAcl property.
// The canned ACL to apply to the object.
const (
	InitiateMultipartUploadOptions_XAmzAcl_Private = "private"
	InitiateMultipartUploadOptions_XAmzAcl_PublicRead = "public-read"
)

// NewInitiateMultipartUploadOptions : Instantiate InitiateMultipartUploadOptions
func (*IbmCloudObjectStorageS3ApiV2) NewInitiateMultipartUploadOptions(bucket string, uploads bool, key string, body string) *InitiateMultipartUploadOptions {
	return &InitiateMultipartUploadOptions{
		Bucket: core.StringPtr(bucket),
		Uploads: core.BoolPtr(uploads),
		Key: core.StringPtr(key),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *InitiateMultipartUploadOptions) SetBucket(bucket string) *InitiateMultipartUploadOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetUploads : Allow user to set Uploads
func (options *InitiateMultipartUploadOptions) SetUploads(uploads bool) *InitiateMultipartUploadOptions {
	options.Uploads = core.BoolPtr(uploads)
	return options
}

// SetKey : Allow user to set Key
func (options *InitiateMultipartUploadOptions) SetKey(key string) *InitiateMultipartUploadOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetBody : Allow user to set Body
func (options *InitiateMultipartUploadOptions) SetBody(body string) *InitiateMultipartUploadOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *InitiateMultipartUploadOptions) SetIfMatch(ifMatch string) *InitiateMultipartUploadOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetIfNoneMatch : Allow user to set IfNoneMatch
func (options *InitiateMultipartUploadOptions) SetIfNoneMatch(ifNoneMatch string) *InitiateMultipartUploadOptions {
	options.IfNoneMatch = core.StringPtr(ifNoneMatch)
	return options
}

// SetIfUnmodifiedSince : Allow user to set IfUnmodifiedSince
func (options *InitiateMultipartUploadOptions) SetIfUnmodifiedSince(ifUnmodifiedSince *strfmt.DateTime) *InitiateMultipartUploadOptions {
	options.IfUnmodifiedSince = ifUnmodifiedSince
	return options
}

// SetCacheControl : Allow user to set CacheControl
func (options *InitiateMultipartUploadOptions) SetCacheControl(cacheControl string) *InitiateMultipartUploadOptions {
	options.CacheControl = core.StringPtr(cacheControl)
	return options
}

// SetContentDisposition : Allow user to set ContentDisposition
func (options *InitiateMultipartUploadOptions) SetContentDisposition(contentDisposition string) *InitiateMultipartUploadOptions {
	options.ContentDisposition = core.StringPtr(contentDisposition)
	return options
}

// SetContentEncoding : Allow user to set ContentEncoding
func (options *InitiateMultipartUploadOptions) SetContentEncoding(contentEncoding string) *InitiateMultipartUploadOptions {
	options.ContentEncoding = core.StringPtr(contentEncoding)
	return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *InitiateMultipartUploadOptions) SetContentLanguage(contentLanguage string) *InitiateMultipartUploadOptions {
	options.ContentLanguage = core.StringPtr(contentLanguage)
	return options
}

// SetExpires : Allow user to set Expires
func (options *InitiateMultipartUploadOptions) SetExpires(expires *strfmt.DateTime) *InitiateMultipartUploadOptions {
	options.Expires = expires
	return options
}

// SetXAmzServerSideEncryption : Allow user to set XAmzServerSideEncryption
func (options *InitiateMultipartUploadOptions) SetXAmzServerSideEncryption(xAmzServerSideEncryption string) *InitiateMultipartUploadOptions {
	options.XAmzServerSideEncryption = core.StringPtr(xAmzServerSideEncryption)
	return options
}

// SetXAmzWebsiteRedirectLocation : Allow user to set XAmzWebsiteRedirectLocation
func (options *InitiateMultipartUploadOptions) SetXAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation string) *InitiateMultipartUploadOptions {
	options.XAmzWebsiteRedirectLocation = core.StringPtr(xAmzWebsiteRedirectLocation)
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *InitiateMultipartUploadOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *InitiateMultipartUploadOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *InitiateMultipartUploadOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *InitiateMultipartUploadOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *InitiateMultipartUploadOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *InitiateMultipartUploadOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetXAmzTagging : Allow user to set XAmzTagging
func (options *InitiateMultipartUploadOptions) SetXAmzTagging(xAmzTagging string) *InitiateMultipartUploadOptions {
	options.XAmzTagging = core.StringPtr(xAmzTagging)
	return options
}

// SetXAmzAcl : Allow user to set XAmzAcl
func (options *InitiateMultipartUploadOptions) SetXAmzAcl(xAmzAcl string) *InitiateMultipartUploadOptions {
	options.XAmzAcl = core.StringPtr(xAmzAcl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *InitiateMultipartUploadOptions) SetHeaders(param map[string]string) *InitiateMultipartUploadOptions {
	options.Headers = param
	return options
}

// Initiator : Container element that identifies who initiated the multipart upload.
type Initiator struct {
	// If the principal is an AWS account, it provides the Canonical User ID. If the principal is an IAM User, it provides
	// a user ARN value.
	ID *string `json:"ID,omitempty"`

	// Name of the Principal.
	DisplayName *string `json:"DisplayName,omitempty"`
}

// UnmarshalInitiator unmarshals an instance of Initiator from the specified map of raw messages.
func UnmarshalInitiator(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Initiator)
	err = core.UnmarshalPrimitive(m, "ID", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "DisplayName", &obj.DisplayName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LifecycleExpiration : Container for the expiration for the lifecycle of the object.
type LifecycleExpiration struct {
	// Indicates at what date the object is to be moved or deleted. Should be in GMT ISO 8601 Format.
	Date *strfmt.DateTime `json:"Date,omitempty"`

	// Indicates the lifetime, in days, of the objects that are subject to the rule. The value must be a non-zero positive
	// integer.
	Days *int64 `json:"Days,omitempty"`

	// Indicates whether IBM COS will remove a delete marker with no noncurrent versions. If set to true, the delete marker
	// will be expired; if set to false the policy takes no action. This cannot be specified with Days or Date in a
	// Lifecycle Expiration Policy.
	ExpiredObjectDeleteMarker *bool `json:"ExpiredObjectDeleteMarker,omitempty"`
}

// UnmarshalLifecycleExpiration unmarshals an instance of LifecycleExpiration from the specified map of raw messages.
func UnmarshalLifecycleExpiration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LifecycleExpiration)
	err = core.UnmarshalPrimitive(m, "Date", &obj.Date)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Days", &obj.Days)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ExpiredObjectDeleteMarker", &obj.ExpiredObjectDeleteMarker)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LifecycleRule : A lifecycle rule for individual objects in an IBM COS bucket.
type LifecycleRule struct {
	// Specifies the expiration for the lifecycle of the object in the form of date, days and, whether the object has a
	// delete marker.
	Expiration *LifecycleExpiration `json:"Expiration,omitempty"`

	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"ID,omitempty"`

	// Prefix identifying one or more objects to which the rule applies. This is No longer used; use `Filter` instead.
	Prefix *string `json:"Prefix,omitempty"`

	// The `Filter` is used to identify objects that a Lifecycle Rule applies to. A `Filter` must have exactly one of
	// `Prefix`, `Tag`, or `And` specified.
	Filter *LifecycleRuleFilter `json:"Filter,omitempty"`

	// If 'Enabled', the rule is currently being applied. If 'Disabled', the rule is not currently being applied.
	Status *string `json:"Status" validate:"required"`

	// Specifies when an IBM COS object transitions to a specified storage class.
	Transitions *LifecycleRuleTransitions `json:"Transitions,omitempty"`
}

// Constants associated with the LifecycleRule.Status property.
// If 'Enabled', the rule is currently being applied. If 'Disabled', the rule is not currently being applied.
const (
	LifecycleRule_Status_Disabled = "Disabled"
	LifecycleRule_Status_Enabled = "Enabled"
)

// UnmarshalLifecycleRule unmarshals an instance of LifecycleRule from the specified map of raw messages.
func UnmarshalLifecycleRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LifecycleRule)
	err = core.UnmarshalModel(m, "Expiration", &obj.Expiration, UnmarshalLifecycleExpiration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ID", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Filter", &obj.Filter, UnmarshalLifecycleRuleFilter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Transitions", &obj.Transitions, UnmarshalLifecycleRuleTransitions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LifecycleRuleAndOperator : This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will
// apply to any object matching all of the predicates configured inside the And operator.
type LifecycleRuleAndOperator struct {
	// Prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"Prefix,omitempty"`

	// All of these tags must exist in the object's tag set in order for the rule to apply.
	Tags *LifecycleRuleAndOperatorTags `json:"Tags,omitempty"`
}

// UnmarshalLifecycleRuleAndOperator unmarshals an instance of LifecycleRuleAndOperator from the specified map of raw messages.
func UnmarshalLifecycleRuleAndOperator(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LifecycleRuleAndOperator)
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Tags", &obj.Tags, UnmarshalLifecycleRuleAndOperatorTags)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LifecycleRuleFilter : The `Filter` is used to identify objects that a Lifecycle Rule applies to. A `Filter` must have exactly one of
// `Prefix`, `Tag`, or `And` specified.
type LifecycleRuleFilter struct {
	// Prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"Prefix,omitempty"`

	// This tag must exist in the object's tag set in order for the rule to apply.
	Tag *Tag `json:"Tag,omitempty"`

	// This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will
	// apply to any object matching all of the predicates configured inside the And operator.
	And *LifecycleRuleAndOperator `json:"And,omitempty"`
}

// UnmarshalLifecycleRuleFilter unmarshals an instance of LifecycleRuleFilter from the specified map of raw messages.
func UnmarshalLifecycleRuleFilter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LifecycleRuleFilter)
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Tag", &obj.Tag, UnmarshalTag)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "And", &obj.And, UnmarshalLifecycleRuleAndOperator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBucketsOptions : The ListBuckets options.
type ListBucketsOptions struct {
	// This header references the service instance that contains buckets to list. This value can be either the full Cloud
	// Resource Name (CRN) or just the GUID segment that identifies the service instance.
	IbmServiceInstanceID *string `validate:"required"`

	// If supplied, the returned listing will also include the provisioning code (`LocationConstraint`) for each bucket.
	// This allows for inferring a bucket's location and associated endpoint.
	Extended *bool

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBucketsOptions : Instantiate ListBucketsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewListBucketsOptions(ibmServiceInstanceID string) *ListBucketsOptions {
	return &ListBucketsOptions{
		IbmServiceInstanceID: core.StringPtr(ibmServiceInstanceID),
	}
}

// SetIbmServiceInstanceID : Allow user to set IbmServiceInstanceID
func (options *ListBucketsOptions) SetIbmServiceInstanceID(ibmServiceInstanceID string) *ListBucketsOptions {
	options.IbmServiceInstanceID = core.StringPtr(ibmServiceInstanceID)
	return options
}

// SetExtended : Allow user to set Extended
func (options *ListBucketsOptions) SetExtended(extended bool) *ListBucketsOptions {
	options.Extended = core.BoolPtr(extended)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListBucketsOptions) SetHeaders(param map[string]string) *ListBucketsOptions {
	options.Headers = param
	return options
}

// ListMultipartUploadsOptions : The ListMultipartUploads options.
type ListMultipartUploadsOptions struct {
	// The name of the bucket to which the multipart upload was initiated.
	Bucket *string `validate:"required,ne="`

	Uploads *bool `validate:"required"`

	// Character you use to group keys. All keys that contain the same string between the prefix, if specified, and the
	// first occurrence of the delimiter after the prefix are grouped under a single result element, `CommonPrefixes`. If
	// you don''t specify the prefix parameter, then the substring starts at the beginning of the key. The keys that are
	// grouped under `CommonPrefixes` result element are not returned elsewhere in the response.
	Delimiter *string

	// Requests IBM COS to encode the object keys in the response and specifies the encoding method to use. An object key
	// may contain any Unicode character; however, XML 1.0 parser cannot parse some characters, such as characters with an
	// ASCII value from 0 to 10. For characters that are not supported in XML 1.0, you can add this parameter to request
	// that IBM COS encode the keys in the response.
	EncodingType *string

	// Together with upload-id-marker, this parameter specifies the multipart upload after which listing should begin. If
	// `upload-id-marker` is not specified, only the keys lexicographically greater than the specified `key-marker` will be
	// included in the list. If `upload-id-marker` is specified, any multipart uploads for a key equal to the `key-marker`
	// might also be included, provided those multipart uploads have upload IDs lexicographically greater than the
	// specified `upload-id-marker`.
	KeyMarker *string

	// Sets the maximum number of multipart uploads, from 1 to 1,000, to return in the response body. 1,000 is the maximum
	// number of uploads that can be returned in a response.
	MaxUploads *int64

	// Lists in-progress uploads only for those keys that begin with the specified prefix. You can use prefixes to separate
	// a bucket into different grouping of keys. (You can think of using prefix to make groups in the same way you'd use a
	// folder in a file system.).
	Prefix *string

	// Together with key-marker, specifies the multipart upload after which listing should begin. If key-marker is not
	// specified, the upload-id-marker parameter is ignored. Otherwise, any multipart uploads for a key equal to the
	// key-marker might be included in the list only if they have an upload ID lexicographically greater than the specified
	// `upload-id-marker`.
	UploadIdMarker *string

	// Pagination limit //PD changed MaxUploads to PaginationLimit.
	PaginationLimit *string

	// Pagination token // PD UploadIdMarker ->.
	PaginationToken *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListMultipartUploadsOptions.EncodingType property.
// Requests IBM COS to encode the object keys in the response and specifies the encoding method to use. An object key
// may contain any Unicode character; however, XML 1.0 parser cannot parse some characters, such as characters with an
// ASCII value from 0 to 10. For characters that are not supported in XML 1.0, you can add this parameter to request
// that IBM COS encode the keys in the response.
const (
	ListMultipartUploadsOptions_EncodingType_URL = "url"
)

// NewListMultipartUploadsOptions : Instantiate ListMultipartUploadsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewListMultipartUploadsOptions(bucket string, uploads bool) *ListMultipartUploadsOptions {
	return &ListMultipartUploadsOptions{
		Bucket: core.StringPtr(bucket),
		Uploads: core.BoolPtr(uploads),
	}
}

// SetBucket : Allow user to set Bucket
func (options *ListMultipartUploadsOptions) SetBucket(bucket string) *ListMultipartUploadsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetUploads : Allow user to set Uploads
func (options *ListMultipartUploadsOptions) SetUploads(uploads bool) *ListMultipartUploadsOptions {
	options.Uploads = core.BoolPtr(uploads)
	return options
}

// SetDelimiter : Allow user to set Delimiter
func (options *ListMultipartUploadsOptions) SetDelimiter(delimiter string) *ListMultipartUploadsOptions {
	options.Delimiter = core.StringPtr(delimiter)
	return options
}

// SetEncodingType : Allow user to set EncodingType
func (options *ListMultipartUploadsOptions) SetEncodingType(encodingType string) *ListMultipartUploadsOptions {
	options.EncodingType = core.StringPtr(encodingType)
	return options
}

// SetKeyMarker : Allow user to set KeyMarker
func (options *ListMultipartUploadsOptions) SetKeyMarker(keyMarker string) *ListMultipartUploadsOptions {
	options.KeyMarker = core.StringPtr(keyMarker)
	return options
}

// SetMaxUploads : Allow user to set MaxUploads
func (options *ListMultipartUploadsOptions) SetMaxUploads(maxUploads int64) *ListMultipartUploadsOptions {
	options.MaxUploads = core.Int64Ptr(maxUploads)
	return options
}

// SetPrefix : Allow user to set Prefix
func (options *ListMultipartUploadsOptions) SetPrefix(prefix string) *ListMultipartUploadsOptions {
	options.Prefix = core.StringPtr(prefix)
	return options
}

// SetUploadIdMarker : Allow user to set UploadIdMarker
func (options *ListMultipartUploadsOptions) SetUploadIdMarker(uploadIdMarker string) *ListMultipartUploadsOptions {
	options.UploadIdMarker = core.StringPtr(uploadIdMarker)
	return options
}

// SetPaginationLimit : Allow user to set PaginationLimit
func (options *ListMultipartUploadsOptions) SetPaginationLimit(paginationLimit string) *ListMultipartUploadsOptions {
	options.PaginationLimit = core.StringPtr(paginationLimit)
	return options
}

// SetPaginationToken : Allow user to set PaginationToken
func (options *ListMultipartUploadsOptions) SetPaginationToken(paginationToken string) *ListMultipartUploadsOptions {
	options.PaginationToken = core.StringPtr(paginationToken)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListMultipartUploadsOptions) SetHeaders(param map[string]string) *ListMultipartUploadsOptions {
	options.Headers = param
	return options
}

// ListMultipartUploadsOutput : ListMultipartUploadsOutput struct
type ListMultipartUploadsOutput struct {
	// The name of the bucket to which the multipart upload was initiated.
	Bucket *string `json:"Bucket,omitempty"`

	// The key at or after which the listing began.
	KeyMarker *string `json:"KeyMarker,omitempty"`

	// Upload ID after which listing began.
	UploadIdMarker *string `json:"UploadIdMarker,omitempty"`

	// When a list is truncated, this element specifies the value that should be used for the key-marker request parameter
	// in a subsequent request.
	NextKeyMarker *string `json:"NextKeyMarker,omitempty"`

	// When a prefix is provided in the request, this field contains the specified prefix. The result contains only keys
	// starting with the specified prefix.
	Prefix *string `json:"Prefix,omitempty"`

	// Contains the delimiter you specified in the request. If you don't specify a delimiter in your request, this element
	// is absent from the response.
	Delimiter *string `json:"Delimiter,omitempty"`

	// When a list is truncated, this element specifies the value that should be used for the `upload-id-marker` request
	// parameter in a subsequent request.
	NextUploadIdMarker *string `json:"NextUploadIdMarker,omitempty"`

	// Maximum number of multipart uploads that could have been included in the response.
	MaxUploads *int64 `json:"MaxUploads,omitempty"`

	// Indicates whether the returned list of multipart uploads is truncated. A value of true indicates that the list was
	// truncated. The list can be truncated if the number of multipart uploads exceeds the limit allowed or specified by
	// max uploads.
	IsTruncated *bool `json:"IsTruncated,omitempty"`

	// Container for elements related to a particular multipart upload. A response can contain zero or more `Upload`
	// elements.
	Uploads *ListMultipartUploadsOutputUploads `json:"Uploads,omitempty"`

	// If you specify a delimiter in the request, then the result returns each distinct key prefix containing the delimiter
	// in a `CommonPrefixes` element. The distinct key prefixes are returned in the `Prefix` child element.
	CommonPrefixes *CommonPrefixList `json:"CommonPrefixes,omitempty"`

	// Encoding type used by IBM COS to encode object keys in the response. If you specify `encoding-type` request
	// parameter, IBM COS includes this element in the response, and returns encoded key name values in the following
	// response elements: `Delimiter`, `KeyMarker`, `Prefix`, `NextKeyMarker`, `Key`.
	EncodingType *string `json:"EncodingType,omitempty"`
}

// Constants associated with the ListMultipartUploadsOutput.EncodingType property.
// Encoding type used by IBM COS to encode object keys in the response. If you specify `encoding-type` request
// parameter, IBM COS includes this element in the response, and returns encoded key name values in the following
// response elements: `Delimiter`, `KeyMarker`, `Prefix`, `NextKeyMarker`, `Key`.
const (
	ListMultipartUploadsOutput_EncodingType_URL = "url"
)

// UnmarshalListMultipartUploadsOutput unmarshals an instance of ListMultipartUploadsOutput from the specified map of raw messages.
func UnmarshalListMultipartUploadsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListMultipartUploadsOutput)
	err = core.UnmarshalPrimitive(m, "Bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "KeyMarker", &obj.KeyMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "UploadIdMarker", &obj.UploadIdMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "NextKeyMarker", &obj.NextKeyMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Delimiter", &obj.Delimiter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "NextUploadIdMarker", &obj.NextUploadIdMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MaxUploads", &obj.MaxUploads)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IsTruncated", &obj.IsTruncated)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Uploads", &obj.Uploads, UnmarshalListMultipartUploadsOutputUploads)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "CommonPrefixes", &obj.CommonPrefixes, UnmarshalCommonPrefixList)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "EncodingType", &obj.EncodingType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListObjectsOptions : The ListObjects options.
type ListObjectsOptions struct {
	// The name of the bucket to be listed.
	Bucket *string `validate:"required,ne="`

	// A delimiter is a character you use to group keys.
	Delimiter *string

	// Requests COS to url-encode the object keys in the response. Object keys may contain any Unicode character; however,
	// XML 1.0 parser cannot parse some characters, such as characters with an ASCII value from 0 to 10. For characters
	// that are not supported in XML 1.0, you can add this parameter to request that COS encodes the keys in the response.
	EncodingType *string

	// Specifies the key to start with when listing objects in a bucket.
	Marker *string

	// Sets the maximum number of keys returned in the response. By default the API returns up to 1,000 key names. The
	// response might contain fewer keys but will never contain more.
	MaxKeys *int64

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListObjectsOptions.EncodingType property.
// Requests COS to url-encode the object keys in the response. Object keys may contain any Unicode character; however,
// XML 1.0 parser cannot parse some characters, such as characters with an ASCII value from 0 to 10. For characters that
// are not supported in XML 1.0, you can add this parameter to request that COS encodes the keys in the response.
const (
	ListObjectsOptions_EncodingType_URL = "url"
)

// NewListObjectsOptions : Instantiate ListObjectsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewListObjectsOptions(bucket string) *ListObjectsOptions {
	return &ListObjectsOptions{
		Bucket: core.StringPtr(bucket),
	}
}

// SetBucket : Allow user to set Bucket
func (options *ListObjectsOptions) SetBucket(bucket string) *ListObjectsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetDelimiter : Allow user to set Delimiter
func (options *ListObjectsOptions) SetDelimiter(delimiter string) *ListObjectsOptions {
	options.Delimiter = core.StringPtr(delimiter)
	return options
}

// SetEncodingType : Allow user to set EncodingType
func (options *ListObjectsOptions) SetEncodingType(encodingType string) *ListObjectsOptions {
	options.EncodingType = core.StringPtr(encodingType)
	return options
}

// SetMarker : Allow user to set Marker
func (options *ListObjectsOptions) SetMarker(marker string) *ListObjectsOptions {
	options.Marker = core.StringPtr(marker)
	return options
}

// SetMaxKeys : Allow user to set MaxKeys
func (options *ListObjectsOptions) SetMaxKeys(maxKeys int64) *ListObjectsOptions {
	options.MaxKeys = core.Int64Ptr(maxKeys)
	return options
}

// SetPrefix : Allow user to set Prefix
func (options *ListObjectsOptions) SetPrefix(prefix string) *ListObjectsOptions {
	options.Prefix = core.StringPtr(prefix)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListObjectsOptions) SetHeaders(param map[string]string) *ListObjectsOptions {
	options.Headers = param
	return options
}

// ListObjectsOutput : ListObjectsOutput struct
type ListObjectsOutput struct {
	// A flag that indicates whether IBM COS returned all of the results that satisfied the search criteria.
	IsTruncated *bool `json:"IsTruncated,omitempty"`

	// Indicates where in the bucket listing begins. Marker is included in the response if it was sent with the request.
	Marker *string `json:"Marker,omitempty"`

	// When response is truncated (the `IsTruncated` element value in the response is true), you can use the key name in
	// this field as marker in the subsequent request to get next set of objects. IBM COS lists objects in alphabetical
	// order.
	//
	// **Note:** This element is returned only if you have delimiter request parameter specified. If response does not
	// include the `NextMarker` and it is truncated, you can use the value of the last key in the response as the marker in
	// the subsequent request to get the next set of object keys.
	NextMarker *string `json:"NextMarker,omitempty"`

	// Metadata about each object returned.
	Contents *ObjectList `json:"Contents,omitempty"`

	// The bucket name.
	Name *string `json:"Name,omitempty"`

	// Keys that begin with the indicated prefix.
	Prefix *string `json:"Prefix,omitempty"`

	// Causes keys that contain the same string between the prefix and the first occurrence of the delimiter to be rolled
	// up into a single result element in the `CommonPrefixes` collection. These rolled-up keys are not returned elsewhere
	// in the response. Each rolled-up result counts as only one return against the `MaxKeys` value.
	Delimiter *string `json:"Delimiter,omitempty"`

	// The maximum number of keys returned in the response body.
	MaxKeys *int64 `json:"MaxKeys,omitempty"`

	// <p>All of the keys rolled up in a common prefix count as a single return when calculating the number of returns.
	// </p> <p>A response can contain CommonPrefixes only if you specify a delimiter.</p> <p>CommonPrefixes contains all
	// (if there are any) keys between Prefix and the next occurrence of the string specified by the delimiter.</p> <p>
	// CommonPrefixes lists keys that act like subdirectories in the directory specified by Prefix.</p> <p>For example, if
	// the prefix is notes/ and the delimiter is a slash (/) as in notes/summer/july, the common prefix is notes/summer/.
	// All of the keys that roll up into a common prefix count as a single return when calculating the number of
	// returns.</p>.
	CommonPrefixes *CommonPrefixList `json:"CommonPrefixes,omitempty"`

	// Encoding type used by IBM COS to encode object keys in the response.
	EncodingType *string `json:"EncodingType,omitempty"`
}

// Constants associated with the ListObjectsOutput.EncodingType property.
// Encoding type used by IBM COS to encode object keys in the response.
const (
	ListObjectsOutput_EncodingType_URL = "url"
)

// UnmarshalListObjectsOutput unmarshals an instance of ListObjectsOutput from the specified map of raw messages.
func UnmarshalListObjectsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListObjectsOutput)
	err = core.UnmarshalPrimitive(m, "IsTruncated", &obj.IsTruncated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Marker", &obj.Marker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "NextMarker", &obj.NextMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Contents", &obj.Contents, UnmarshalObjectList)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Delimiter", &obj.Delimiter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MaxKeys", &obj.MaxKeys)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "CommonPrefixes", &obj.CommonPrefixes, UnmarshalCommonPrefixList)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "EncodingType", &obj.EncodingType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListObjectsV2Options : The ListObjectsV2 options.
type ListObjectsV2Options struct {
	// Bucket name to list.
	Bucket *string `validate:"required,ne="`

	ListType *string `validate:"required"`

	// A delimiter is a character you use to group keys.
	Delimiter *string

	// Encoding type used by IBM COS to encode object keys in the response.
	EncodingType *string

	// Sets the maximum number of keys returned in the response. By default the API returns up to 1,000 key names. The
	// response might contain fewer keys but will never contain more.
	MaxKeys *int64

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string

	// `ContinuationToken` indicates IBM COS that the list is being continued on this bucket with a token.
	// ContinuationToken is obfuscated and is not a real key.
	ContinuationToken *string

	// The owner field (Service Instance ID) is not present in listV2 by default, if you want to return the Service
	// Instance ID with each key in the result, then set the fetch owner field to true.
	FetchOwner *bool

	// `StartAfter` is where you want IBM COS to start listing from. IBM COS starts listing after this specified key.
	// `StartAfter` can be any key in the bucket.
	StartAfter *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListObjectsV2Options.EncodingType property.
// Encoding type used by IBM COS to encode object keys in the response.
const (
	ListObjectsV2Options_EncodingType_URL = "url"
)

// NewListObjectsV2Options : Instantiate ListObjectsV2Options
func (*IbmCloudObjectStorageS3ApiV2) NewListObjectsV2Options(bucket string, listType string) *ListObjectsV2Options {
	return &ListObjectsV2Options{
		Bucket: core.StringPtr(bucket),
		ListType: core.StringPtr(listType),
	}
}

// SetBucket : Allow user to set Bucket
func (options *ListObjectsV2Options) SetBucket(bucket string) *ListObjectsV2Options {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetListType : Allow user to set ListType
func (options *ListObjectsV2Options) SetListType(listType string) *ListObjectsV2Options {
	options.ListType = core.StringPtr(listType)
	return options
}

// SetDelimiter : Allow user to set Delimiter
func (options *ListObjectsV2Options) SetDelimiter(delimiter string) *ListObjectsV2Options {
	options.Delimiter = core.StringPtr(delimiter)
	return options
}

// SetEncodingType : Allow user to set EncodingType
func (options *ListObjectsV2Options) SetEncodingType(encodingType string) *ListObjectsV2Options {
	options.EncodingType = core.StringPtr(encodingType)
	return options
}

// SetMaxKeys : Allow user to set MaxKeys
func (options *ListObjectsV2Options) SetMaxKeys(maxKeys int64) *ListObjectsV2Options {
	options.MaxKeys = core.Int64Ptr(maxKeys)
	return options
}

// SetPrefix : Allow user to set Prefix
func (options *ListObjectsV2Options) SetPrefix(prefix string) *ListObjectsV2Options {
	options.Prefix = core.StringPtr(prefix)
	return options
}

// SetContinuationToken : Allow user to set ContinuationToken
func (options *ListObjectsV2Options) SetContinuationToken(continuationToken string) *ListObjectsV2Options {
	options.ContinuationToken = core.StringPtr(continuationToken)
	return options
}

// SetFetchOwner : Allow user to set FetchOwner
func (options *ListObjectsV2Options) SetFetchOwner(fetchOwner bool) *ListObjectsV2Options {
	options.FetchOwner = core.BoolPtr(fetchOwner)
	return options
}

// SetStartAfter : Allow user to set StartAfter
func (options *ListObjectsV2Options) SetStartAfter(startAfter string) *ListObjectsV2Options {
	options.StartAfter = core.StringPtr(startAfter)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListObjectsV2Options) SetHeaders(param map[string]string) *ListObjectsV2Options {
	options.Headers = param
	return options
}

// ListObjectsV2Output : ListObjectsV2Output struct
type ListObjectsV2Output struct {
	// Set to false if all of the results were returned. Set to true if more keys are available to return. If the number of
	// results exceeds that specified by MaxKeys, all of the results might not be returned.
	IsTruncated *bool `json:"IsTruncated,omitempty"`

	// Metadata about each object returned.
	Contents *ObjectList `json:"Contents,omitempty"`

	// <p>The bucket name.</p> <p>When using this API with an access point, you must direct requests to the access point
	// hostname. The access point hostname takes the form
	// <i>AccessPointName</i>-<i>AccountId</i>.s3-accesspoint.<i>Region</i>.amazonaws.com. When using this operation with
	// an access point through the AWS SDKs, you provide the access point ARN in place of the bucket name. For more
	// information about access point ARNs, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html">Using Access Points</a> in the
	// <i>Amazon Simple Storage Service Developer Guide</i>.</p> <p>When using this API with IBM COS on Outposts, you must
	// direct requests to the S3 on Outposts hostname. The S3 on Outposts hostname takes the form
	// <i>AccessPointName</i>-<i>AccountId</i>.<i>outpostID</i>.s3-outposts.<i>Region</i>.amazonaws.com. When using this
	// operation using S3 on Outposts through the AWS SDKs, you provide the Outposts bucket ARN in place of the bucket
	// name. For more information about S3 on Outposts ARNs, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html">Using S3 on Outposts</a> in the <i>Amazon
	// Simple Storage Service Developer Guide</i>.</p>.
	Name *string `json:"Name,omitempty"`

	// Keys that begin with the indicated prefix.
	Prefix *string `json:"Prefix,omitempty"`

	// Causes keys that contain the same string between the prefix and the first occurrence of the delimiter to be rolled
	// up into a single result element in the CommonPrefixes collection. These rolled-up keys are not returned elsewhere in
	// the response. Each rolled-up result counts as only one return against the `MaxKeys` value.
	Delimiter *string `json:"Delimiter,omitempty"`

	// Sets the maximum number of keys returned in the response. By default the API returns up to 1,000 key names. The
	// response might contain fewer keys but will never contain more.
	MaxKeys *int64 `json:"MaxKeys,omitempty"`

	// <p>All of the keys rolled up into a common prefix count as a single return when calculating the number of
	// returns.</p> <p>A response can contain `CommonPrefixes` only if you specify a delimiter.</p> <p> `CommonPrefixes`
	// contains all (if there are any) keys between `Prefix` and the next occurrence of the string specified by a
	// delimiter.</p> <p> `CommonPrefixes` lists keys that act like subdirectories in the directory specified by
	// `Prefix`.</p> <p>For example, if the prefix is `notes/` and the delimiter is a slash (`/`) as in
	// `notes/summer/july`, the common prefix is `notes/summer/`. All of the keys that roll up into a common prefix count
	// as a single return when calculating the number of returns. </p>.
	CommonPrefixes *CommonPrefixList `json:"CommonPrefixes,omitempty"`

	// Encoding type used by IBM COS to encode object keys in the response. If you specify `encoding-type` request
	// parameter, IBM COS includes this element in the response, and returns encoded key name values in the following
	// response elements: `Delimiter`, `KeyMarker`, `Prefix`, `NextKeyMarker`, `Key`.
	EncodingType *string `json:"EncodingType,omitempty"`

	// KeyCount is the number of keys returned with this request. KeyCount will always be less than equals to MaxKeys
	// field. Say you ask for 50 keys, your result will include less than equals 50 keys.
	KeyCount *int64 `json:"KeyCount,omitempty"`

	// If ContinuationToken was sent with the request, it is included in the response.
	ContinuationToken *string `json:"ContinuationToken,omitempty"`

	// `NextContinuationToken` is sent when `isTruncated` is true, which means there are more keys in the bucket that can
	// be listed. The next list requests to IBM COS can be continued with this `NextContinuationToken`.
	// `NextContinuationToken` is obfuscated and is not a real key.
	NextContinuationToken *string `json:"NextContinuationToken,omitempty"`

	// If StartAfter was sent with the request, it is included in the response.
	StartAfter *string `json:"StartAfter,omitempty"`
}

// Constants associated with the ListObjectsV2Output.EncodingType property.
// Encoding type used by IBM COS to encode object keys in the response. If you specify `encoding-type` request
// parameter, IBM COS includes this element in the response, and returns encoded key name values in the following
// response elements: `Delimiter`, `KeyMarker`, `Prefix`, `NextKeyMarker`, `Key`.
const (
	ListObjectsV2Output_EncodingType_URL = "url"
)

// UnmarshalListObjectsV2Output unmarshals an instance of ListObjectsV2Output from the specified map of raw messages.
func UnmarshalListObjectsV2Output(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListObjectsV2Output)
	err = core.UnmarshalPrimitive(m, "IsTruncated", &obj.IsTruncated)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Contents", &obj.Contents, UnmarshalObjectList)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Prefix", &obj.Prefix)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Delimiter", &obj.Delimiter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MaxKeys", &obj.MaxKeys)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "CommonPrefixes", &obj.CommonPrefixes, UnmarshalCommonPrefixList)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "EncodingType", &obj.EncodingType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "KeyCount", &obj.KeyCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ContinuationToken", &obj.ContinuationToken)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "NextContinuationToken", &obj.NextContinuationToken)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "StartAfter", &obj.StartAfter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPartsOptions : The ListParts options.
type ListPartsOptions struct {
	// The bucket where the parts have been uploaded.
	Bucket *string `validate:"required,ne="`

	// Object key for which the multipart upload was initiated.
	Key *string `validate:"required,ne="`

	// Upload ID identifying the multipart upload whose parts are being listed.
	UploadID *string `validate:"required"`

	// Sets the maximum number of parts to return.
	MaxParts *int64

	// Specifies the part after which listing should begin. Only parts with higher part numbers will be listed.
	PartNumberMarker *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPartsOptions : Instantiate ListPartsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewListPartsOptions(bucket string, key string, uploadID string) *ListPartsOptions {
	return &ListPartsOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		UploadID: core.StringPtr(uploadID),
	}
}

// SetBucket : Allow user to set Bucket
func (options *ListPartsOptions) SetBucket(bucket string) *ListPartsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *ListPartsOptions) SetKey(key string) *ListPartsOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetUploadID : Allow user to set UploadID
func (options *ListPartsOptions) SetUploadID(uploadID string) *ListPartsOptions {
	options.UploadID = core.StringPtr(uploadID)
	return options
}

// SetMaxParts : Allow user to set MaxParts
func (options *ListPartsOptions) SetMaxParts(maxParts int64) *ListPartsOptions {
	options.MaxParts = core.Int64Ptr(maxParts)
	return options
}

// SetPartNumberMarker : Allow user to set PartNumberMarker
func (options *ListPartsOptions) SetPartNumberMarker(partNumberMarker int64) *ListPartsOptions {
	options.PartNumberMarker = core.Int64Ptr(partNumberMarker)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPartsOptions) SetHeaders(param map[string]string) *ListPartsOptions {
	options.Headers = param
	return options
}

// ListPartsOutput : ListPartsOutput struct
type ListPartsOutput struct {
	// The name of the bucket to which the multipart upload was initiated.
	Bucket *string `json:"Bucket,omitempty"`

	// Object key for which the multipart upload was initiated.
	Key *string `json:"Key,omitempty"`

	// Upload ID identifying the multipart upload whose parts are being listed.
	UploadID *string `json:"UploadId,omitempty"`

	// When a list is truncated, this element specifies the last part in the list, as well as the value to use for the
	// part-number-marker request parameter in a subsequent request.
	PartNumberMarker *int64 `json:"PartNumberMarker,omitempty"`

	// When a list is truncated, this element specifies the last part in the list, as well as the value to use for the
	// part-number-marker request parameter in a subsequent request.
	NextPartNumberMarker *int64 `json:"NextPartNumberMarker,omitempty"`

	// Maximum number of parts that were allowed in the response.
	MaxParts *int64 `json:"MaxParts,omitempty"`

	// Indicates whether the returned list of parts is truncated. A true value indicates that the list was truncated. A
	// list can be truncated if the number of parts exceeds the limit returned in the MaxParts element.
	IsTruncated *bool `json:"IsTruncated,omitempty"`

	// Container for elements related to a particular part. A response can contain zero or more `Part` elements.
	Parts *ListPartsOutputParts `json:"Parts,omitempty"`

	// Container element that identifies who initiated the multipart upload. If the initiator is an AWS account, this
	// element provides the same information as the `Owner` element. If the initiator is an IAM User, this element provides
	// the user ARN and display name.
	Initiator *Initiator `json:"Initiator,omitempty"`

	// Container element that identifies the object owner, after the object is created. If multipart upload is initiated by
	// an IAM user, this element provides the parent account ID and display name.
	Owner *Owner `json:"Owner,omitempty"`

	// Class of storage (STANDARD or REDUCED_REDUNDANCY) used to store the uploaded object.
	StorageClass *string `json:"StorageClass,omitempty"`
}

// Constants associated with the ListPartsOutput.StorageClass property.
// Class of storage (STANDARD or REDUCED_REDUNDANCY) used to store the uploaded object.
const (
	ListPartsOutput_StorageClass_Accelerated = "ACCELERATED"
	ListPartsOutput_StorageClass_Glacier = "GLACIER"
	ListPartsOutput_StorageClass_Standard = "STANDARD"
)

// UnmarshalListPartsOutput unmarshals an instance of ListPartsOutput from the specified map of raw messages.
func UnmarshalListPartsOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPartsOutput)
	err = core.UnmarshalPrimitive(m, "Bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "UploadId", &obj.UploadID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PartNumberMarker", &obj.PartNumberMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "NextPartNumberMarker", &obj.NextPartNumberMarker)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MaxParts", &obj.MaxParts)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IsTruncated", &obj.IsTruncated)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Parts", &obj.Parts, UnmarshalListPartsOutputParts)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Initiator", &obj.Initiator, UnmarshalInitiator)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Owner", &obj.Owner, UnmarshalOwner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "StorageClass", &obj.StorageClass)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MultipartUpload : Container for the `MultipartUpload` for the IBM COS object.
type MultipartUpload struct {
	// Upload ID that identifies the multipart upload.
	UploadID *string `json:"UploadId,omitempty"`

	// Key of the object for which the multipart upload was initiated.
	Key *string `json:"Key,omitempty"`

	// Date and time at which the multipart upload was initiated.
	Initiated *strfmt.DateTime `json:"Initiated,omitempty"`

	// The class of storage used to store the object.
	StorageClass *string `json:"StorageClass,omitempty"`

	// Specifies the owner of the object that is part of the multipart upload.
	Owner *Owner `json:"Owner,omitempty"`

	// Identifies who initiated the multipart upload.
	Initiator *Initiator `json:"Initiator,omitempty"`
}

// Constants associated with the MultipartUpload.StorageClass property.
// The class of storage used to store the object.
const (
	MultipartUpload_StorageClass_Accelerated = "ACCELERATED"
	MultipartUpload_StorageClass_Glacier = "GLACIER"
	MultipartUpload_StorageClass_Standard = "STANDARD"
)

// UnmarshalMultipartUpload unmarshals an instance of MultipartUpload from the specified map of raw messages.
func UnmarshalMultipartUpload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MultipartUpload)
	err = core.UnmarshalPrimitive(m, "UploadId", &obj.UploadID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Initiated", &obj.Initiated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "StorageClass", &obj.StorageClass)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Owner", &obj.Owner, UnmarshalOwner)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Initiator", &obj.Initiator, UnmarshalInitiator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Object : An object consists of data and its descriptive metadata.
type Object struct {
	// The name that you assign to an object. You use the object key to retrieve the object.
	Key *string `json:"Key,omitempty"`

	// The date the Object was Last Modified.
	LastModified *strfmt.DateTime `json:"LastModified,omitempty"`

	// <p>The entity tag is a hash of the object. The ETag reflects changes only to the contents of an object, not its
	// metadata. The ETag may or may not be an MD5 digest of the object data. Whether or not it is depends on how the
	// object was created and how it is encrypted as described below:</p> <ul> <li> <p>Objects created by the PUT Object,
	// POST Object, or Copy operation, or through the AWS Management Console, and are encrypted by SSE-S3 or plaintext,
	// have ETags that are an MD5 digest of their object data.</p> </li> <li> <p>Objects created by the PUT Object, POST
	// Object, or Copy operation, or through the AWS Management Console, and are encrypted by SSE-C or SSE-KMS, have ETags
	// that are not an MD5 digest of their object data.</p> </li> <li> <p>If an object is created by either the Multipart
	// Upload or Part Copy operation, the ETag is not an MD5 digest, regardless of the method of encryption.</p> </li>
	// </ul>.
	ETag *string `json:"ETag,omitempty"`

	// Size in bytes of the object.
	Size *int64 `json:"Size,omitempty"`

	// The class of storage used to store the object.
	StorageClass *string `json:"StorageClass,omitempty"`

	// The owner of the object.
	Owner *Owner `json:"Owner,omitempty"`
}

// Constants associated with the Object.StorageClass property.
// The class of storage used to store the object.
const (
	Object_StorageClass_Accelerated = "ACCELERATED"
	Object_StorageClass_Glacier = "GLACIER"
	Object_StorageClass_Standard = "STANDARD"
)

// UnmarshalObject unmarshals an instance of Object from the specified map of raw messages.
func UnmarshalObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Object)
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LastModified", &obj.LastModified)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ETag", &obj.ETag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "StorageClass", &obj.StorageClass)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Owner", &obj.Owner, UnmarshalOwner)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Owner : Container for the owner's display name and ID.
type Owner struct {
	// Container for the display name of the owner.
	DisplayName *string `json:"DisplayName,omitempty"`

	// Container for the ID of the owner.
	ID *string `json:"ID,omitempty"`
}

// UnmarshalOwner unmarshals an instance of Owner from the specified map of raw messages.
func UnmarshalOwner(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Owner)
	err = core.UnmarshalPrimitive(m, "DisplayName", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ID", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Part : Container for elements related to a part.
type Part struct {
	// Part number identifying the part. This is a positive integer between 1 and 10,000.
	PartNumber *int64 `json:"PartNumber,omitempty"`

	// Date and time at which the part was uploaded.
	LastModified *strfmt.DateTime `json:"LastModified,omitempty"`

	// Entity tag returned when the part was uploaded.
	ETag *string `json:"ETag,omitempty"`

	// Size in bytes of the uploaded part data.
	Size *int64 `json:"Size,omitempty"`
}

// UnmarshalPart unmarshals an instance of Part from the specified map of raw messages.
func UnmarshalPart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Part)
	err = core.UnmarshalPrimitive(m, "PartNumber", &obj.PartNumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "LastModified", &obj.LastModified)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ETag", &obj.ETag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Size", &obj.Size)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PublicAccessBlockConfiguration : The PublicAccessBlock configuration that you want to apply to this IBM COS bucket. You can enable the configuration
// options in any combination. For more information about when IBM COS considers a bucket or object public, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status">The
// Meaning of "Public"</a> in the <i>Amazon Simple Storage Service Developer Guide</i>.
type PublicAccessBlockConfiguration struct {
	// <p>Specifies whether IBM COS should block public access control lists (ACLs) for this bucket and objects in this
	// bucket. Setting this element to `TRUE` causes the following behavior:</p> <ul> <li> <p>PUT Bucket acl and PUT Object
	// acl calls fail if the specified ACL is public.</p> </li> <li> <p>PUT Object calls fail if the request includes a
	// public ACL.</p> </li> <li> <p>PUT Bucket calls fail if the request includes a public ACL.</p> </li> </ul>
	// <p>Enabling this setting doesn't affect existing policies or ACLs.</p>.
	BlockPublicAcls *bool `json:"BlockPublicAcls,omitempty"`

	// <p>Specifies whether IBM COS should ignore public ACLs for this bucket and objects in this bucket. Setting this
	// element to `TRUE` causes IBM COS to ignore all public ACLs on this bucket and objects in this bucket.</p>
	// <p>Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs
	// from being set.</p>.
	IgnorePublicAcls *bool `json:"IgnorePublicAcls,omitempty"`

	// <p>Specifies whether IBM COS should block public bucket policies for this bucket. Setting this element to `TRUE`
	// causes IBM COS to reject calls to PUT Bucket policy if the specified bucket policy allows public access. </p>
	// <p>Enabling this setting doesn't affect existing bucket policies.</p>.
	BlockPublicPolicy *bool `json:"BlockPublicPolicy,omitempty"`

	// <p>Specifies whether IBM COS should restrict public bucket policies for this bucket. Setting this element to `TRUE`
	// restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a
	// public policy.</p> <p>Enabling this setting doesn't affect previously stored bucket policies, except that public and
	// cross-account access within any public bucket policy, including non-public delegation to specific accounts, is
	// blocked.</p>.
	RestrictPublicBuckets *bool `json:"RestrictPublicBuckets,omitempty"`
}

// UnmarshalPublicAccessBlockConfiguration unmarshals an instance of PublicAccessBlockConfiguration from the specified map of raw messages.
func UnmarshalPublicAccessBlockConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PublicAccessBlockConfiguration)
	err = core.UnmarshalPrimitive(m, "BlockPublicAcls", &obj.BlockPublicAcls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "IgnorePublicAcls", &obj.IgnorePublicAcls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "BlockPublicPolicy", &obj.BlockPublicPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "RestrictPublicBuckets", &obj.RestrictPublicBuckets)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PutBucketAclOptions : The PutBucketAcl options.
type PutBucketAclOptions struct {
	// The bucket to which to apply the ACL.
	Bucket *string `validate:"required,ne="`

	Acl *bool `validate:"required"`

	// The canned ACL to apply to the bucket.
	XAmzAcl *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the PutBucketAclOptions.XAmzAcl property.
// The canned ACL to apply to the bucket.
const (
	PutBucketAclOptions_XAmzAcl_Private = "private"
	PutBucketAclOptions_XAmzAcl_PublicRead = "public-read"
)

// NewPutBucketAclOptions : Instantiate PutBucketAclOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutBucketAclOptions(bucket string, acl bool) *PutBucketAclOptions {
	return &PutBucketAclOptions{
		Bucket: core.StringPtr(bucket),
		Acl: core.BoolPtr(acl),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutBucketAclOptions) SetBucket(bucket string) *PutBucketAclOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetAcl : Allow user to set Acl
func (options *PutBucketAclOptions) SetAcl(acl bool) *PutBucketAclOptions {
	options.Acl = core.BoolPtr(acl)
	return options
}

// SetXAmzAcl : Allow user to set XAmzAcl
func (options *PutBucketAclOptions) SetXAmzAcl(xAmzAcl string) *PutBucketAclOptions {
	options.XAmzAcl = core.StringPtr(xAmzAcl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutBucketAclOptions) SetHeaders(param map[string]string) *PutBucketAclOptions {
	options.Headers = param
	return options
}

// PutBucketCorsOptions : The PutBucketCors options.
type PutBucketCorsOptions struct {
	// Specifies the bucket impacted by the CORS configuration.
	Bucket *string `validate:"required,ne="`

	Cors *bool `validate:"required"`

	Body *string `validate:"required"`

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutBucketCorsOptions : Instantiate PutBucketCorsOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutBucketCorsOptions(bucket string, cors bool, body string) *PutBucketCorsOptions {
	return &PutBucketCorsOptions{
		Bucket: core.StringPtr(bucket),
		Cors: core.BoolPtr(cors),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutBucketCorsOptions) SetBucket(bucket string) *PutBucketCorsOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetCors : Allow user to set Cors
func (options *PutBucketCorsOptions) SetCors(cors bool) *PutBucketCorsOptions {
	options.Cors = core.BoolPtr(cors)
	return options
}

// SetBody : Allow user to set Body
func (options *PutBucketCorsOptions) SetBody(body string) *PutBucketCorsOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *PutBucketCorsOptions) SetContentMD5(contentMD5 string) *PutBucketCorsOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutBucketCorsOptions) SetHeaders(param map[string]string) *PutBucketCorsOptions {
	options.Headers = param
	return options
}

// PutBucketLifecycleConfigurationOptions : The PutBucketLifecycleConfiguration options.
type PutBucketLifecycleConfigurationOptions struct {
	// The name of the bucket for which to set the configuration.
	Bucket *string `validate:"required,ne="`

	Lifecycle *bool `validate:"required"`

	Body *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutBucketLifecycleConfigurationOptions : Instantiate PutBucketLifecycleConfigurationOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutBucketLifecycleConfigurationOptions(bucket string, lifecycle bool, body string) *PutBucketLifecycleConfigurationOptions {
	return &PutBucketLifecycleConfigurationOptions{
		Bucket: core.StringPtr(bucket),
		Lifecycle: core.BoolPtr(lifecycle),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutBucketLifecycleConfigurationOptions) SetBucket(bucket string) *PutBucketLifecycleConfigurationOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetLifecycle : Allow user to set Lifecycle
func (options *PutBucketLifecycleConfigurationOptions) SetLifecycle(lifecycle bool) *PutBucketLifecycleConfigurationOptions {
	options.Lifecycle = core.BoolPtr(lifecycle)
	return options
}

// SetBody : Allow user to set Body
func (options *PutBucketLifecycleConfigurationOptions) SetBody(body string) *PutBucketLifecycleConfigurationOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutBucketLifecycleConfigurationOptions) SetHeaders(param map[string]string) *PutBucketLifecycleConfigurationOptions {
	options.Headers = param
	return options
}

// PutBucketProtectionConfigurationOptions : The PutBucketProtectionConfiguration options.
type PutBucketProtectionConfigurationOptions struct {
	// The name of the bucket for which to set the configuration.
	Bucket *string `validate:"required,ne="`

	Protection *bool `validate:"required"`

	Body *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutBucketProtectionConfigurationOptions : Instantiate PutBucketProtectionConfigurationOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutBucketProtectionConfigurationOptions(bucket string, protection bool, body string) *PutBucketProtectionConfigurationOptions {
	return &PutBucketProtectionConfigurationOptions{
		Bucket: core.StringPtr(bucket),
		Protection: core.BoolPtr(protection),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutBucketProtectionConfigurationOptions) SetBucket(bucket string) *PutBucketProtectionConfigurationOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetProtection : Allow user to set Protection
func (options *PutBucketProtectionConfigurationOptions) SetProtection(protection bool) *PutBucketProtectionConfigurationOptions {
	options.Protection = core.BoolPtr(protection)
	return options
}

// SetBody : Allow user to set Body
func (options *PutBucketProtectionConfigurationOptions) SetBody(body string) *PutBucketProtectionConfigurationOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutBucketProtectionConfigurationOptions) SetHeaders(param map[string]string) *PutBucketProtectionConfigurationOptions {
	options.Headers = param
	return options
}

// PutBucketWebsiteOptions : The PutBucketWebsite options.
type PutBucketWebsiteOptions struct {
	// The bucket that will serve a static website.
	Bucket *string `validate:"required,ne="`

	Website *bool `validate:"required"`

	Body *string `validate:"required"`

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutBucketWebsiteOptions : Instantiate PutBucketWebsiteOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutBucketWebsiteOptions(bucket string, website bool, body string) *PutBucketWebsiteOptions {
	return &PutBucketWebsiteOptions{
		Bucket: core.StringPtr(bucket),
		Website: core.BoolPtr(website),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutBucketWebsiteOptions) SetBucket(bucket string) *PutBucketWebsiteOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetWebsite : Allow user to set Website
func (options *PutBucketWebsiteOptions) SetWebsite(website bool) *PutBucketWebsiteOptions {
	options.Website = core.BoolPtr(website)
	return options
}

// SetBody : Allow user to set Body
func (options *PutBucketWebsiteOptions) SetBody(body string) *PutBucketWebsiteOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *PutBucketWebsiteOptions) SetContentMD5(contentMD5 string) *PutBucketWebsiteOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutBucketWebsiteOptions) SetHeaders(param map[string]string) *PutBucketWebsiteOptions {
	options.Headers = param
	return options
}

// PutObjectAclOptions : The PutObjectAcl options.
type PutObjectAclOptions struct {
	// The bucket name that contains the object to which you want to attach the ACL.
	Bucket *string `validate:"required,ne="`

	// Key for which the PUT operation was initiated.
	Key *string `validate:"required,ne="`

	Acl *bool `validate:"required"`

	Body *string `validate:"required"`

	// The canned ACL to apply to the object.
	XAmzAcl *string

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// <p>Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.</p> <p>This action is not
	// supported by IBM COS on Outposts.</p>.
	XAmzGrantFullControl *string

	// <p>Allows grantee to list the objects in the bucket.</p> <p>This action is not supported by IBM COS on
	// Outposts.</p>.
	XAmzGrantRead *string

	// <p>Allows grantee to read the bucket ACL.</p> <p>This action is not supported by IBM COS on Outposts.</p>.
	XAmzGrantReadAcp *string

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	XAmzGrantWrite *string

	// <p>Allows grantee to write the ACL for the applicable bucket.</p> <p>This action is not supported by IBM COS on
	// Outposts.</p>.
	XAmzGrantWriteAcp *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the PutObjectAclOptions.XAmzAcl property.
// The canned ACL to apply to the object.
const (
	PutObjectAclOptions_XAmzAcl_Private = "private"
	PutObjectAclOptions_XAmzAcl_PublicRead = "public-read"
)

// NewPutObjectAclOptions : Instantiate PutObjectAclOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutObjectAclOptions(bucket string, key string, acl bool, body string) *PutObjectAclOptions {
	return &PutObjectAclOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Acl: core.BoolPtr(acl),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutObjectAclOptions) SetBucket(bucket string) *PutObjectAclOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *PutObjectAclOptions) SetKey(key string) *PutObjectAclOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetAcl : Allow user to set Acl
func (options *PutObjectAclOptions) SetAcl(acl bool) *PutObjectAclOptions {
	options.Acl = core.BoolPtr(acl)
	return options
}

// SetBody : Allow user to set Body
func (options *PutObjectAclOptions) SetBody(body string) *PutObjectAclOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetXAmzAcl : Allow user to set XAmzAcl
func (options *PutObjectAclOptions) SetXAmzAcl(xAmzAcl string) *PutObjectAclOptions {
	options.XAmzAcl = core.StringPtr(xAmzAcl)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *PutObjectAclOptions) SetContentMD5(contentMD5 string) *PutObjectAclOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetXAmzGrantFullControl : Allow user to set XAmzGrantFullControl
func (options *PutObjectAclOptions) SetXAmzGrantFullControl(xAmzGrantFullControl string) *PutObjectAclOptions {
	options.XAmzGrantFullControl = core.StringPtr(xAmzGrantFullControl)
	return options
}

// SetXAmzGrantRead : Allow user to set XAmzGrantRead
func (options *PutObjectAclOptions) SetXAmzGrantRead(xAmzGrantRead string) *PutObjectAclOptions {
	options.XAmzGrantRead = core.StringPtr(xAmzGrantRead)
	return options
}

// SetXAmzGrantReadAcp : Allow user to set XAmzGrantReadAcp
func (options *PutObjectAclOptions) SetXAmzGrantReadAcp(xAmzGrantReadAcp string) *PutObjectAclOptions {
	options.XAmzGrantReadAcp = core.StringPtr(xAmzGrantReadAcp)
	return options
}

// SetXAmzGrantWrite : Allow user to set XAmzGrantWrite
func (options *PutObjectAclOptions) SetXAmzGrantWrite(xAmzGrantWrite string) *PutObjectAclOptions {
	options.XAmzGrantWrite = core.StringPtr(xAmzGrantWrite)
	return options
}

// SetXAmzGrantWriteAcp : Allow user to set XAmzGrantWriteAcp
func (options *PutObjectAclOptions) SetXAmzGrantWriteAcp(xAmzGrantWriteAcp string) *PutObjectAclOptions {
	options.XAmzGrantWriteAcp = core.StringPtr(xAmzGrantWriteAcp)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutObjectAclOptions) SetHeaders(param map[string]string) *PutObjectAclOptions {
	options.Headers = param
	return options
}

// PutObjectOptions : The PutObject options.
type PutObjectOptions struct {
	// The bucket that will hold the object.
	Bucket *string `validate:"required,ne="`

	// Object key for which the PUT operation was initiated.
	Key *string `validate:"required,ne="`

	Body *string `validate:"required"`

	// The canned ACL to apply to the object.
	XAmzAcl *string

	// Upload the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412
	// (precondition failed).
	IfMatch *string

	// Upload the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not
	// modified).
	IfNoneMatch *string

	// Upload the object only if it has not been modified since the specified time, otherwise return a 412 (precondition
	// failed).
	IfUnmodifiedSince *strfmt.DateTime

	// Can be used to specify caching behavior along the request/reply chain. For more information, see [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	CacheControl *string

	// Specifies presentational information for the object. For more information, see [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1).
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied
	// to obtain the media-type referenced by the Content-Type header field. For more information, see [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11).
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// Size of the body in bytes. This parameter is useful when the size of the body cannot be determined automatically.
	// For more information, see [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13).
	ContentLength *int64

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// The date and time at which the object is no longer cacheable. For more information, [RFC
	// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21).
	Expires *strfmt.DateTime

	// The server-side encryption algorithm used when storing this object in IBM COS (`AES256`).
	XAmzServerSideEncryption *string

	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or
	// to an external URL.
	XAmzWebsiteRedirectLocation *string

	// Specifies the algorithm to use to when encrypting the object (for example, `AES256`).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use in encrypting data. This value is used to store
	// the object and then it is discarded; IBM COS does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the `x-amz-server-side-encryption-customer-algorithm` header.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IBM COS uses this header for a message
	// integrity check to ensure that the encryption key was transmitted without error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// A set of tags for the object. The tags must be encoded as URL duery parameters (For example, `SomeKey=SomeValue`).
	XAmzTagging *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the PutObjectOptions.XAmzAcl property.
// The canned ACL to apply to the object.
const (
	PutObjectOptions_XAmzAcl_Private = "private"
	PutObjectOptions_XAmzAcl_PublicRead = "public-read"
)

// Constants associated with the PutObjectOptions.XAmzServerSideEncryption property.
// The server-side encryption algorithm used when storing this object in IBM COS (`AES256`).
const (
	PutObjectOptions_XAmzServerSideEncryption_Aes256 = "AES256"
)

// NewPutObjectOptions : Instantiate PutObjectOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutObjectOptions(bucket string, key string, body string) *PutObjectOptions {
	return &PutObjectOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutObjectOptions) SetBucket(bucket string) *PutObjectOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *PutObjectOptions) SetKey(key string) *PutObjectOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetBody : Allow user to set Body
func (options *PutObjectOptions) SetBody(body string) *PutObjectOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetXAmzAcl : Allow user to set XAmzAcl
func (options *PutObjectOptions) SetXAmzAcl(xAmzAcl string) *PutObjectOptions {
	options.XAmzAcl = core.StringPtr(xAmzAcl)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *PutObjectOptions) SetIfMatch(ifMatch string) *PutObjectOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetIfNoneMatch : Allow user to set IfNoneMatch
func (options *PutObjectOptions) SetIfNoneMatch(ifNoneMatch string) *PutObjectOptions {
	options.IfNoneMatch = core.StringPtr(ifNoneMatch)
	return options
}

// SetIfUnmodifiedSince : Allow user to set IfUnmodifiedSince
func (options *PutObjectOptions) SetIfUnmodifiedSince(ifUnmodifiedSince *strfmt.DateTime) *PutObjectOptions {
	options.IfUnmodifiedSince = ifUnmodifiedSince
	return options
}

// SetCacheControl : Allow user to set CacheControl
func (options *PutObjectOptions) SetCacheControl(cacheControl string) *PutObjectOptions {
	options.CacheControl = core.StringPtr(cacheControl)
	return options
}

// SetContentDisposition : Allow user to set ContentDisposition
func (options *PutObjectOptions) SetContentDisposition(contentDisposition string) *PutObjectOptions {
	options.ContentDisposition = core.StringPtr(contentDisposition)
	return options
}

// SetContentEncoding : Allow user to set ContentEncoding
func (options *PutObjectOptions) SetContentEncoding(contentEncoding string) *PutObjectOptions {
	options.ContentEncoding = core.StringPtr(contentEncoding)
	return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *PutObjectOptions) SetContentLanguage(contentLanguage string) *PutObjectOptions {
	options.ContentLanguage = core.StringPtr(contentLanguage)
	return options
}

// SetContentLength : Allow user to set ContentLength
func (options *PutObjectOptions) SetContentLength(contentLength int64) *PutObjectOptions {
	options.ContentLength = core.Int64Ptr(contentLength)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *PutObjectOptions) SetContentMD5(contentMD5 string) *PutObjectOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetExpires : Allow user to set Expires
func (options *PutObjectOptions) SetExpires(expires *strfmt.DateTime) *PutObjectOptions {
	options.Expires = expires
	return options
}

// SetXAmzServerSideEncryption : Allow user to set XAmzServerSideEncryption
func (options *PutObjectOptions) SetXAmzServerSideEncryption(xAmzServerSideEncryption string) *PutObjectOptions {
	options.XAmzServerSideEncryption = core.StringPtr(xAmzServerSideEncryption)
	return options
}

// SetXAmzWebsiteRedirectLocation : Allow user to set XAmzWebsiteRedirectLocation
func (options *PutObjectOptions) SetXAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation string) *PutObjectOptions {
	options.XAmzWebsiteRedirectLocation = core.StringPtr(xAmzWebsiteRedirectLocation)
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *PutObjectOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *PutObjectOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *PutObjectOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *PutObjectOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *PutObjectOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *PutObjectOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetXAmzTagging : Allow user to set XAmzTagging
func (options *PutObjectOptions) SetXAmzTagging(xAmzTagging string) *PutObjectOptions {
	options.XAmzTagging = core.StringPtr(xAmzTagging)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutObjectOptions) SetHeaders(param map[string]string) *PutObjectOptions {
	options.Headers = param
	return options
}

// PutObjectTaggingOptions : The PutObjectTagging options.
type PutObjectTaggingOptions struct {
	// The bucket containing the object.
	Bucket *string `validate:"required,ne="`

	// Name of the object key.
	Key *string `validate:"required,ne="`

	Tagging *bool `validate:"required"`

	Body *string `validate:"required"`

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutObjectTaggingOptions : Instantiate PutObjectTaggingOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutObjectTaggingOptions(bucket string, key string, tagging bool, body string) *PutObjectTaggingOptions {
	return &PutObjectTaggingOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Tagging: core.BoolPtr(tagging),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutObjectTaggingOptions) SetBucket(bucket string) *PutObjectTaggingOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *PutObjectTaggingOptions) SetKey(key string) *PutObjectTaggingOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetTagging : Allow user to set Tagging
func (options *PutObjectTaggingOptions) SetTagging(tagging bool) *PutObjectTaggingOptions {
	options.Tagging = core.BoolPtr(tagging)
	return options
}

// SetBody : Allow user to set Body
func (options *PutObjectTaggingOptions) SetBody(body string) *PutObjectTaggingOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *PutObjectTaggingOptions) SetContentMD5(contentMD5 string) *PutObjectTaggingOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutObjectTaggingOptions) SetHeaders(param map[string]string) *PutObjectTaggingOptions {
	options.Headers = param
	return options
}

// PutPublicAccessBlockOptions : The PutPublicAccessBlock options.
type PutPublicAccessBlockOptions struct {
	// The name of the IBM COS bucket whose `PublicAccessBlock` configuration you want to set.
	Bucket *string `validate:"required,ne="`

	PublicAccessBlock *bool `validate:"required"`

	Body *string `validate:"required"`

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutPublicAccessBlockOptions : Instantiate PutPublicAccessBlockOptions
func (*IbmCloudObjectStorageS3ApiV2) NewPutPublicAccessBlockOptions(bucket string, publicAccessBlock bool, body string) *PutPublicAccessBlockOptions {
	return &PutPublicAccessBlockOptions{
		Bucket: core.StringPtr(bucket),
		PublicAccessBlock: core.BoolPtr(publicAccessBlock),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *PutPublicAccessBlockOptions) SetBucket(bucket string) *PutPublicAccessBlockOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetPublicAccessBlock : Allow user to set PublicAccessBlock
func (options *PutPublicAccessBlockOptions) SetPublicAccessBlock(publicAccessBlock bool) *PutPublicAccessBlockOptions {
	options.PublicAccessBlock = core.BoolPtr(publicAccessBlock)
	return options
}

// SetBody : Allow user to set Body
func (options *PutPublicAccessBlockOptions) SetBody(body string) *PutPublicAccessBlockOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *PutPublicAccessBlockOptions) SetContentMD5(contentMD5 string) *PutPublicAccessBlockOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutPublicAccessBlockOptions) SetHeaders(param map[string]string) *PutPublicAccessBlockOptions {
	options.Headers = param
	return options
}

// Redirect : Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.
type Redirect struct {
	// The host name to use in the redirect request.
	HostName *string `json:"HostName,omitempty"`

	// The HTTP redirect code to use on the response. Not required if one of the siblings is present.
	HttpRedirectCode *string `json:"HttpRedirectCode,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	Protocol *string `json:"Protocol,omitempty"`

	// The object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix
	// `docs/` (objects in the `docs/` folder) to `documents/`, you can set a condition block with `KeyPrefixEquals` set to
	// `docs/` and in the Redirect set `ReplaceKeyPrefixWith` to `/documents`. Not required if one of the siblings is
	// present. Can be present only if `ReplaceKeyWith` is not provided.
	ReplaceKeyPrefixWith *string `json:"ReplaceKeyPrefixWith,omitempty"`

	// The specific object key to use in the redirect request. For example, redirect request to `error.html`. Not required
	// if one of the siblings is present. Can be present only if `ReplaceKeyPrefixWith` is not provided.
	ReplaceKeyWith *string `json:"ReplaceKeyWith,omitempty"`
}

// Constants associated with the Redirect.Protocol property.
// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
const (
	Redirect_Protocol_Http = "http"
	Redirect_Protocol_Https = "https"
)

// UnmarshalRedirect unmarshals an instance of Redirect from the specified map of raw messages.
func UnmarshalRedirect(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Redirect)
	err = core.UnmarshalPrimitive(m, "HostName", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "HttpRedirectCode", &obj.HttpRedirectCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Protocol", &obj.Protocol)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ReplaceKeyPrefixWith", &obj.ReplaceKeyPrefixWith)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ReplaceKeyWith", &obj.ReplaceKeyWith)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedirectAllRequestsTo : Specifies the redirect behavior of all requests to a website endpoint of an IBM COS bucket.
type RedirectAllRequestsTo struct {
	// Name of the host where requests are redirected.
	HostName *string `json:"HostName" validate:"required"`

	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	Protocol *string `json:"Protocol,omitempty"`
}

// Constants associated with the RedirectAllRequestsTo.Protocol property.
// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
const (
	RedirectAllRequestsTo_Protocol_Http = "http"
	RedirectAllRequestsTo_Protocol_Https = "https"
)

// UnmarshalRedirectAllRequestsTo unmarshals an instance of RedirectAllRequestsTo from the specified map of raw messages.
func UnmarshalRedirectAllRequestsTo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedirectAllRequestsTo)
	err = core.UnmarshalPrimitive(m, "HostName", &obj.HostName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Protocol", &obj.Protocol)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RestoreObjectOptions : The RestoreObject options.
type RestoreObjectOptions struct {
	// The bucket name or containing the object to restore.
	Bucket *string `validate:"required,ne="`

	// Object key for which the operation was initiated.
	Key *string `validate:"required,ne="`

	Restore *bool `validate:"required"`

	Body *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRestoreObjectOptions : Instantiate RestoreObjectOptions
func (*IbmCloudObjectStorageS3ApiV2) NewRestoreObjectOptions(bucket string, key string, restore bool, body string) *RestoreObjectOptions {
	return &RestoreObjectOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		Restore: core.BoolPtr(restore),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *RestoreObjectOptions) SetBucket(bucket string) *RestoreObjectOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *RestoreObjectOptions) SetKey(key string) *RestoreObjectOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetRestore : Allow user to set Restore
func (options *RestoreObjectOptions) SetRestore(restore bool) *RestoreObjectOptions {
	options.Restore = core.BoolPtr(restore)
	return options
}

// SetBody : Allow user to set Body
func (options *RestoreObjectOptions) SetBody(body string) *RestoreObjectOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RestoreObjectOptions) SetHeaders(param map[string]string) *RestoreObjectOptions {
	options.Headers = param
	return options
}

// RoutingRulesItem : Specifies the redirect behavior and when a redirect is applied. For more information about routing rules, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects">Configuring
// advanced conditional redirects</a> in the <i>Amazon Simple Storage Service Developer Guide</i>.
type RoutingRulesItem struct {
	// A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If
	// request is for pages in the `/docs` folder, redirect to the `/documents` folder. 2. If request results in HTTP error
	// 4xx, redirect request to another host where you might process the error.
	Condition *Condition `json:"Condition,omitempty"`

	// Container for redirect information. You can redirect requests to another host, to another page, or with another
	// protocol. In the event of an error, you can specify a different error code to return.
	Redirect *Redirect `json:"Redirect" validate:"required"`
}

// UnmarshalRoutingRulesItem unmarshals an instance of RoutingRulesItem from the specified map of raw messages.
func UnmarshalRoutingRulesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoutingRulesItem)
	err = core.UnmarshalModel(m, "Condition", &obj.Condition, UnmarshalCondition)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "Redirect", &obj.Redirect, UnmarshalRedirect)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tag : A container of a key value name pair.
type Tag struct {
	// Name of the object key.
	Key *string `json:"Key" validate:"required"`

	// Value of the tag.
	Value *string `json:"Value" validate:"required"`
}

// UnmarshalTag unmarshals an instance of Tag from the specified map of raw messages.
func UnmarshalTag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tag)
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TagSetItem : A container of a key value name pair.
type TagSetItem struct {
	// Name of the object key.
	Key *string `json:"Key" validate:"required"`

	// Value of the tag.
	Value *string `json:"Value" validate:"required"`
}

// UnmarshalTagSetItem unmarshals an instance of TagSetItem from the specified map of raw messages.
func UnmarshalTagSetItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TagSetItem)
	err = core.UnmarshalPrimitive(m, "Key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Transition : Specifies when an object transitions to a specified storage class.
type Transition struct {
	// Indicates when objects are transitioned to the specified storage class. The date value must be in ISO 8601 format.
	// The time is always midnight UTC.
	Date *strfmt.DateTime `json:"Date,omitempty"`

	// Indicates the number of days after creation when objects are transitioned to the specified storage class. The value
	// must be a positive integer.
	Days *int64 `json:"Days,omitempty"`

	// The storage class to which you want the object to transition.
	StorageClass *string `json:"StorageClass,omitempty"`
}

// Constants associated with the Transition.StorageClass property.
// The storage class to which you want the object to transition.
const (
	Transition_StorageClass_Accelerated = "ACCELERATED"
	Transition_StorageClass_Glacier = "GLACIER"
)

// UnmarshalTransition unmarshals an instance of Transition from the specified map of raw messages.
func UnmarshalTransition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Transition)
	err = core.UnmarshalPrimitive(m, "Date", &obj.Date)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Days", &obj.Days)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "StorageClass", &obj.StorageClass)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UploadPartCopyOptions : The UploadPartCopy options.
type UploadPartCopyOptions struct {
	// The bucket name.
	Bucket *string `validate:"required,ne="`

	// Specifies the source object to use as a part in the multipart upload.
	XAmzCopySource *string `validate:"required"`

	// Object key for which the multipart upload was initiated.
	Key *string `validate:"required,ne="`

	// Part number of part being copied. This is a positive integer between 1 and 10,000.
	PartNumber *int64 `validate:"required"`

	// Upload ID identifying the multipart upload whose part is being copied.
	UploadID *string `validate:"required"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	XAmzCopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time.
	XAmzCopySourceIfModifiedSince *strfmt.DateTime

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	XAmzCopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time.
	XAmzCopySourceIfUnmodifiedSince *strfmt.DateTime

	// The range of bytes to copy from the source object. The range value must use the form bytes=first-last, where the
	// first and last are the zero-based byte offsets to copy. For example, bytes=0-9 indicates that you want to copy the
	// first 10 bytes of the source. You can copy a range only if the source object is greater than 5 MB.
	XAmzCopySourceRange *string

	// Specifies the algorithm to use to when encrypting the object (`AES256`).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use in encrypting data. This value is used to store
	// the object and then it is discarded; IBM COS does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the `x-amz-server-side-encryption-customer-algorithm` header. This must be the same
	// encryption key specified in the initiate multipart upload request.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to [RFC 1321](https://tools.ietf.org/html/rfc1321).
	// IBM COS uses this header for a message integrity check to ensure that the encryption key was transmitted without
	// error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// Specifies the algorithm to use when decrypting the source object (`AES256`).
	XAmzCopySourceServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use to decrypt the source object. The encryption key
	// provided in this header must be one that was used when the source object was created.
	XAmzCopySourceServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to [RFC 1321](https://tools.ietf.org/html/rfc1321).
	// IBM COS uses this header for a message integrity check to ensure that the encryption key was transmitted without
	// error.
	XAmzCopySourceServerSideEncryptionCustomerKeyMD5 *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUploadPartCopyOptions : Instantiate UploadPartCopyOptions
func (*IbmCloudObjectStorageS3ApiV2) NewUploadPartCopyOptions(bucket string, xAmzCopySource string, key string, partNumber int64, uploadID string) *UploadPartCopyOptions {
	return &UploadPartCopyOptions{
		Bucket: core.StringPtr(bucket),
		XAmzCopySource: core.StringPtr(xAmzCopySource),
		Key: core.StringPtr(key),
		PartNumber: core.Int64Ptr(partNumber),
		UploadID: core.StringPtr(uploadID),
	}
}

// SetBucket : Allow user to set Bucket
func (options *UploadPartCopyOptions) SetBucket(bucket string) *UploadPartCopyOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetXAmzCopySource : Allow user to set XAmzCopySource
func (options *UploadPartCopyOptions) SetXAmzCopySource(xAmzCopySource string) *UploadPartCopyOptions {
	options.XAmzCopySource = core.StringPtr(xAmzCopySource)
	return options
}

// SetKey : Allow user to set Key
func (options *UploadPartCopyOptions) SetKey(key string) *UploadPartCopyOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetPartNumber : Allow user to set PartNumber
func (options *UploadPartCopyOptions) SetPartNumber(partNumber int64) *UploadPartCopyOptions {
	options.PartNumber = core.Int64Ptr(partNumber)
	return options
}

// SetUploadID : Allow user to set UploadID
func (options *UploadPartCopyOptions) SetUploadID(uploadID string) *UploadPartCopyOptions {
	options.UploadID = core.StringPtr(uploadID)
	return options
}

// SetXAmzCopySourceIfMatch : Allow user to set XAmzCopySourceIfMatch
func (options *UploadPartCopyOptions) SetXAmzCopySourceIfMatch(xAmzCopySourceIfMatch string) *UploadPartCopyOptions {
	options.XAmzCopySourceIfMatch = core.StringPtr(xAmzCopySourceIfMatch)
	return options
}

// SetXAmzCopySourceIfModifiedSince : Allow user to set XAmzCopySourceIfModifiedSince
func (options *UploadPartCopyOptions) SetXAmzCopySourceIfModifiedSince(xAmzCopySourceIfModifiedSince *strfmt.DateTime) *UploadPartCopyOptions {
	options.XAmzCopySourceIfModifiedSince = xAmzCopySourceIfModifiedSince
	return options
}

// SetXAmzCopySourceIfNoneMatch : Allow user to set XAmzCopySourceIfNoneMatch
func (options *UploadPartCopyOptions) SetXAmzCopySourceIfNoneMatch(xAmzCopySourceIfNoneMatch string) *UploadPartCopyOptions {
	options.XAmzCopySourceIfNoneMatch = core.StringPtr(xAmzCopySourceIfNoneMatch)
	return options
}

// SetXAmzCopySourceIfUnmodifiedSince : Allow user to set XAmzCopySourceIfUnmodifiedSince
func (options *UploadPartCopyOptions) SetXAmzCopySourceIfUnmodifiedSince(xAmzCopySourceIfUnmodifiedSince *strfmt.DateTime) *UploadPartCopyOptions {
	options.XAmzCopySourceIfUnmodifiedSince = xAmzCopySourceIfUnmodifiedSince
	return options
}

// SetXAmzCopySourceRange : Allow user to set XAmzCopySourceRange
func (options *UploadPartCopyOptions) SetXAmzCopySourceRange(xAmzCopySourceRange string) *UploadPartCopyOptions {
	options.XAmzCopySourceRange = core.StringPtr(xAmzCopySourceRange)
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *UploadPartCopyOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *UploadPartCopyOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *UploadPartCopyOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *UploadPartCopyOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *UploadPartCopyOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *UploadPartCopyOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetXAmzCopySourceServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzCopySourceServerSideEncryptionCustomerAlgorithm
func (options *UploadPartCopyOptions) SetXAmzCopySourceServerSideEncryptionCustomerAlgorithm(xAmzCopySourceServerSideEncryptionCustomerAlgorithm string) *UploadPartCopyOptions {
	options.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzCopySourceServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzCopySourceServerSideEncryptionCustomerKey : Allow user to set XAmzCopySourceServerSideEncryptionCustomerKey
func (options *UploadPartCopyOptions) SetXAmzCopySourceServerSideEncryptionCustomerKey(xAmzCopySourceServerSideEncryptionCustomerKey string) *UploadPartCopyOptions {
	options.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr(xAmzCopySourceServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzCopySourceServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzCopySourceServerSideEncryptionCustomerKeyMD5
func (options *UploadPartCopyOptions) SetXAmzCopySourceServerSideEncryptionCustomerKeyMD5(xAmzCopySourceServerSideEncryptionCustomerKeyMD5 string) *UploadPartCopyOptions {
	options.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzCopySourceServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UploadPartCopyOptions) SetHeaders(param map[string]string) *UploadPartCopyOptions {
	options.Headers = param
	return options
}

// UploadPartCopyOutput : UploadPartCopyOutput struct
type UploadPartCopyOutput struct {
	// Container for all response elements.
	CopyPartResult *CopyPartResult `json:"CopyPartResult,omitempty"`
}

// UnmarshalUploadPartCopyOutput unmarshals an instance of UploadPartCopyOutput from the specified map of raw messages.
func UnmarshalUploadPartCopyOutput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UploadPartCopyOutput)
	err = core.UnmarshalModel(m, "CopyPartResult", &obj.CopyPartResult, UnmarshalCopyPartResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UploadPartOptions : The UploadPart options.
type UploadPartOptions struct {
	// The name of the bucket to which the multipart upload was initiated.
	Bucket *string `validate:"required,ne="`

	// Object key for which the multipart upload was initiated.
	Key *string `validate:"required,ne="`

	// Part number of part being uploaded. This is a positive integer between 1 and 10,000.
	PartNumber *int64 `validate:"required"`

	// Upload ID identifying the multipart upload whose part is being uploaded.
	UploadID *string `validate:"required"`

	Body *string `validate:"required"`

	// Size of the body in bytes. This parameter is useful when the size of the body cannot be determined automatically.
	// For more information, see [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13).
	ContentLength *int64

	// The base64-encoded 128-bit MD5 digest of the payload (just the request body without the headers) according to [RFC
	// 1864](http://www.ietf.org/rfc/rfc1864.txt). This header can be used as a message integrity check to verify that the
	// data is the same data that was originally sent. Although it is optional, it is recommended to use the Content-MD5
	// mechanism as an end-to-end integrity check.
	ContentMD5 *string

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	XAmzServerSideEncryptionCustomerAlgorithm *string

	// Specifies the customer-provided encryption key for IBM COS to use in encrypting data. This value is used to store
	// the object and then it is discarded; IBM COS does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the `x-amz-server-side-encryption-customer-algorithm header`. This must be the same
	// encryption key specified in the initiate multipart upload request.
	XAmzServerSideEncryptionCustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IBM COS uses this header for a message
	// integrity check to ensure that the encryption key was transmitted without error.
	XAmzServerSideEncryptionCustomerKeyMD5 *string

	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this
	// parameter in their requests. For information about downloading objects from requester pays buckets, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html">Downloading Objects in
	// Requestor Pays Buckets</a> in the <i>IBM COS Developer Guide</i>.
	XAmzRequestPayer *string

	// The account id of the expected bucket owner. If the bucket is owned by a different account, the request will fail
	// with an HTTP `403 (Access Denied)` error.
	XAmzExpectedBucketOwner *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UploadPartOptions.XAmzRequestPayer property.
// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this
// parameter in their requests. For information about downloading objects from requester pays buckets, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html">Downloading Objects in
// Requestor Pays Buckets</a> in the <i>IBM COS Developer Guide</i>.
const (
	UploadPartOptions_XAmzRequestPayer_Requester = "requester"
)

// NewUploadPartOptions : Instantiate UploadPartOptions
func (*IbmCloudObjectStorageS3ApiV2) NewUploadPartOptions(bucket string, key string, partNumber int64, uploadID string, body string) *UploadPartOptions {
	return &UploadPartOptions{
		Bucket: core.StringPtr(bucket),
		Key: core.StringPtr(key),
		PartNumber: core.Int64Ptr(partNumber),
		UploadID: core.StringPtr(uploadID),
		Body: core.StringPtr(body),
	}
}

// SetBucket : Allow user to set Bucket
func (options *UploadPartOptions) SetBucket(bucket string) *UploadPartOptions {
	options.Bucket = core.StringPtr(bucket)
	return options
}

// SetKey : Allow user to set Key
func (options *UploadPartOptions) SetKey(key string) *UploadPartOptions {
	options.Key = core.StringPtr(key)
	return options
}

// SetPartNumber : Allow user to set PartNumber
func (options *UploadPartOptions) SetPartNumber(partNumber int64) *UploadPartOptions {
	options.PartNumber = core.Int64Ptr(partNumber)
	return options
}

// SetUploadID : Allow user to set UploadID
func (options *UploadPartOptions) SetUploadID(uploadID string) *UploadPartOptions {
	options.UploadID = core.StringPtr(uploadID)
	return options
}

// SetBody : Allow user to set Body
func (options *UploadPartOptions) SetBody(body string) *UploadPartOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentLength : Allow user to set ContentLength
func (options *UploadPartOptions) SetContentLength(contentLength int64) *UploadPartOptions {
	options.ContentLength = core.Int64Ptr(contentLength)
	return options
}

// SetContentMD5 : Allow user to set ContentMD5
func (options *UploadPartOptions) SetContentMD5(contentMD5 string) *UploadPartOptions {
	options.ContentMD5 = core.StringPtr(contentMD5)
	return options
}

// SetXAmzServerSideEncryptionCustomerAlgorithm : Allow user to set XAmzServerSideEncryptionCustomerAlgorithm
func (options *UploadPartOptions) SetXAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm string) *UploadPartOptions {
	options.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr(xAmzServerSideEncryptionCustomerAlgorithm)
	return options
}

// SetXAmzServerSideEncryptionCustomerKey : Allow user to set XAmzServerSideEncryptionCustomerKey
func (options *UploadPartOptions) SetXAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey string) *UploadPartOptions {
	options.XAmzServerSideEncryptionCustomerKey = core.StringPtr(xAmzServerSideEncryptionCustomerKey)
	return options
}

// SetXAmzServerSideEncryptionCustomerKeyMD5 : Allow user to set XAmzServerSideEncryptionCustomerKeyMD5
func (options *UploadPartOptions) SetXAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5 string) *UploadPartOptions {
	options.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr(xAmzServerSideEncryptionCustomerKeyMD5)
	return options
}

// SetXAmzRequestPayer : Allow user to set XAmzRequestPayer
func (options *UploadPartOptions) SetXAmzRequestPayer(xAmzRequestPayer string) *UploadPartOptions {
	options.XAmzRequestPayer = core.StringPtr(xAmzRequestPayer)
	return options
}

// SetXAmzExpectedBucketOwner : Allow user to set XAmzExpectedBucketOwner
func (options *UploadPartOptions) SetXAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) *UploadPartOptions {
	options.XAmzExpectedBucketOwner = core.StringPtr(xAmzExpectedBucketOwner)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UploadPartOptions) SetHeaders(param map[string]string) *UploadPartOptions {
	options.Headers = param
	return options
}

// CORSRuleAllowedHeaders : Headers that are specified in the `Access-Control-Request-Headers` header. These headers are allowed in a preflight
// OPTIONS request. In response to any preflight OPTIONS request, IBM COS returns any requested headers that are
// allowed.
// This model "extends" String
type CORSRuleAllowedHeaders struct {
}

// UnmarshalCORSRuleAllowedHeaders unmarshals an instance of CORSRuleAllowedHeaders from the specified map of raw messages.
func UnmarshalCORSRuleAllowedHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CORSRuleAllowedHeaders)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CORSRuleAllowedMethods : An HTTP method that you allow the origin to execute. Valid values are `GET`, `PUT`, `HEAD`, `POST`, and `DELETE`.
// This model "extends" String
type CORSRuleAllowedMethods struct {
}

// UnmarshalCORSRuleAllowedMethods unmarshals an instance of CORSRuleAllowedMethods from the specified map of raw messages.
func UnmarshalCORSRuleAllowedMethods(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CORSRuleAllowedMethods)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CORSRuleAllowedOrigins : One or more origins you want customers to be able to access the bucket from.
// This model "extends" String
type CORSRuleAllowedOrigins struct {
}

// UnmarshalCORSRuleAllowedOrigins unmarshals an instance of CORSRuleAllowedOrigins from the specified map of raw messages.
func UnmarshalCORSRuleAllowedOrigins(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CORSRuleAllowedOrigins)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CORSRuleExposeHeaders : One or more headers in the response that you want customers to be able to access from their applications (for
// example, from a JavaScript `XMLHttpRequest` object).
// This model "extends" String
type CORSRuleExposeHeaders struct {
}

// UnmarshalCORSRuleExposeHeaders unmarshals an instance of CORSRuleExposeHeaders from the specified map of raw messages.
func UnmarshalCORSRuleExposeHeaders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CORSRuleExposeHeaders)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CommonPrefixList : CommonPrefixList struct
// This model "extends" CommonPrefix
type CommonPrefixList struct {
}

// UnmarshalCommonPrefixList unmarshals an instance of CommonPrefixList from the specified map of raw messages.
func UnmarshalCommonPrefixList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CommonPrefixList)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteObjectsOutputErrors : Container for a failed delete operation that describes the object that IBM COS attempted to delete and the error it
// encountered.
// This model "extends" Error
type DeleteObjectsOutputErrors struct {
}

// UnmarshalDeleteObjectsOutputErrors unmarshals an instance of DeleteObjectsOutputErrors from the specified map of raw messages.
func UnmarshalDeleteObjectsOutputErrors(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteObjectsOutputErrors)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketCorsOutputCORSRules : A set of origins and methods (cross-origin access that you want to allow). You can add up to 100 rules to the
// configuration.
// This model "extends" CORSRule
type GetBucketCorsOutputCORSRules struct {
}

// UnmarshalGetBucketCorsOutputCORSRules unmarshals an instance of GetBucketCorsOutputCORSRules from the specified map of raw messages.
func UnmarshalGetBucketCorsOutputCORSRules(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketCorsOutputCORSRules)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBucketLifecycleConfigurationOutputRules : Container for a lifecycle rule.
// This model "extends" LifecycleRule
type GetBucketLifecycleConfigurationOutputRules struct {
}

// UnmarshalGetBucketLifecycleConfigurationOutputRules unmarshals an instance of GetBucketLifecycleConfigurationOutputRules from the specified map of raw messages.
func UnmarshalGetBucketLifecycleConfigurationOutputRules(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBucketLifecycleConfigurationOutputRules)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetObjectAclOutputGrants : A list of grants.
// This model "extends" GrantsItem
type GetObjectAclOutputGrants struct {
}

// UnmarshalGetObjectAclOutputGrants unmarshals an instance of GetObjectAclOutputGrants from the specified map of raw messages.
func UnmarshalGetObjectAclOutputGrants(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetObjectAclOutputGrants)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LifecycleRuleAndOperatorTags : All of these tags must exist in the object's tag set in order for the rule to apply.
// This model "extends" TagSetItem
type LifecycleRuleAndOperatorTags struct {
}

// UnmarshalLifecycleRuleAndOperatorTags unmarshals an instance of LifecycleRuleAndOperatorTags from the specified map of raw messages.
func UnmarshalLifecycleRuleAndOperatorTags(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LifecycleRuleAndOperatorTags)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LifecycleRuleTransitions : Specifies when an IBM COS object transitions to a specified storage class.
// This model "extends" Transition
type LifecycleRuleTransitions struct {
}

// UnmarshalLifecycleRuleTransitions unmarshals an instance of LifecycleRuleTransitions from the specified map of raw messages.
func UnmarshalLifecycleRuleTransitions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LifecycleRuleTransitions)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListMultipartUploadsOutputUploads : Container for elements related to a particular multipart upload. A response can contain zero or more `Upload`
// elements.
// This model "extends" MultipartUpload
type ListMultipartUploadsOutputUploads struct {
}

// UnmarshalListMultipartUploadsOutputUploads unmarshals an instance of ListMultipartUploadsOutputUploads from the specified map of raw messages.
func UnmarshalListMultipartUploadsOutputUploads(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListMultipartUploadsOutputUploads)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPartsOutputParts : Container for elements related to a particular part. A response can contain zero or more `Part` elements.
// This model "extends" Part
type ListPartsOutputParts struct {
}

// UnmarshalListPartsOutputParts unmarshals an instance of ListPartsOutputParts from the specified map of raw messages.
func UnmarshalListPartsOutputParts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPartsOutputParts)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectList : ObjectList struct
// This model "extends" Object
type ObjectList struct {
}

// UnmarshalObjectList unmarshals an instance of ObjectList from the specified map of raw messages.
func UnmarshalObjectList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectList)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RoutingRules : RoutingRules struct
// This model "extends" RoutingRulesItem
type RoutingRules struct {
}

// UnmarshalRoutingRules unmarshals an instance of RoutingRules from the specified map of raw messages.
func UnmarshalRoutingRules(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoutingRules)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TagSet : TagSet struct
// This model "extends" TagSetItem
type TagSet struct {
}

// UnmarshalTagSet unmarshals an instance of TagSet from the specified map of raw messages.
func UnmarshalTagSet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TagSet)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
