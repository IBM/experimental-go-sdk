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

package ibmcloudobjectstorages3apiv2_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/experimental-go-sdk/ibmcloudobjectstorages3apiv2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`IbmCloudObjectStorageS3ApiV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`HeadBucket(headBucketOptions *HeadBucketOptions)`, func() {
		headBucketPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headBucketPath))
					Expect(req.Method).To(Equal("HEAD"))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadBucket successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.HeadBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadBucketOptions model
				headBucketOptionsModel := new(ibmcloudobjectstorages3apiv2.HeadBucketOptions)
				headBucketOptionsModel.Bucket = core.StringPtr("testString")
				headBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.HeadBucket(headBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadBucket with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the HeadBucketOptions model
				headBucketOptionsModel := new(ibmcloudobjectstorages3apiv2.HeadBucketOptions)
				headBucketOptionsModel.Bucket = core.StringPtr("testString")
				headBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.HeadBucket(headBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadBucketOptions model with no property values
				headBucketOptionsModelNew := new(ibmcloudobjectstorages3apiv2.HeadBucketOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.HeadBucket(headBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`HeadObject(headObjectOptions *HeadObjectOptions)`, func() {
		headObjectPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(headObjectPath))
					Expect(req.Method).To(Equal("HEAD"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Range"]).ToNot(BeNil())
					Expect(req.Header["Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke HeadObject successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.HeadObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the HeadObjectOptions model
				headObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.HeadObjectOptions)
				headObjectOptionsModel.Bucket = core.StringPtr("testString")
				headObjectOptionsModel.Key = core.StringPtr("testString")
				headObjectOptionsModel.IfMatch = core.StringPtr("testString")
				headObjectOptionsModel.IfModifiedSince = CreateMockDateTime()
				headObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				headObjectOptionsModel.Range = core.StringPtr("testString")
				headObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				headObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				headObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				headObjectOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				headObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.HeadObject(headObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke HeadObject with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the HeadObjectOptions model
				headObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.HeadObjectOptions)
				headObjectOptionsModel.Bucket = core.StringPtr("testString")
				headObjectOptionsModel.Key = core.StringPtr("testString")
				headObjectOptionsModel.IfMatch = core.StringPtr("testString")
				headObjectOptionsModel.IfModifiedSince = CreateMockDateTime()
				headObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				headObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				headObjectOptionsModel.Range = core.StringPtr("testString")
				headObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				headObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				headObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				headObjectOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				headObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.HeadObject(headObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the HeadObjectOptions model with no property values
				headObjectOptionsModelNew := new(ibmcloudobjectstorages3apiv2.HeadObjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.HeadObject(headObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBuckets(listBucketsOptions *ListBucketsOptions) - Operation response error`, func() {
		listBucketsPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Ibm-Service-Instance-Id"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Service-Instance-Id"][0]).To(Equal(fmt.Sprintf("%v", "d6f76k03-6k4f-4a82-n165-697654o63903")))
					// TODO: Add check for extended query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBuckets with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListBucketsOptions model
				listBucketsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListBucketsOptions)
				listBucketsOptionsModel.IbmServiceInstanceID = core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")
				listBucketsOptionsModel.Extended = core.BoolPtr(true)
				listBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListBuckets(listBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListBuckets(listBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListBuckets(listBucketsOptions *ListBucketsOptions)`, func() {
		listBucketsPath := "/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Service-Instance-Id"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Service-Instance-Id"][0]).To(Equal(fmt.Sprintf("%v", "d6f76k03-6k4f-4a82-n165-697654o63903")))
					// TODO: Add check for extended query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"owner": {"id": "d6f76k03-6k4f-4a82-n165-697654o63903", "displayName": "d6f76k03-6k4f-4a82-n165-697654o63903"}, "buckets": [{"name": "myBucket", "creationDate": "2019-04-29T20:49:22.374Z"}]}`)
				}))
			})
			It(`Invoke ListBuckets successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the ListBucketsOptions model
				listBucketsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListBucketsOptions)
				listBucketsOptionsModel.IbmServiceInstanceID = core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")
				listBucketsOptionsModel.Extended = core.BoolPtr(true)
				listBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.ListBucketsWithContext(ctx, listBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListBuckets(listBucketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.ListBucketsWithContext(ctx, listBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBucketsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Ibm-Service-Instance-Id"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Service-Instance-Id"][0]).To(Equal(fmt.Sprintf("%v", "d6f76k03-6k4f-4a82-n165-697654o63903")))
					// TODO: Add check for extended query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"owner": {"id": "d6f76k03-6k4f-4a82-n165-697654o63903", "displayName": "d6f76k03-6k4f-4a82-n165-697654o63903"}, "buckets": [{"name": "myBucket", "creationDate": "2019-04-29T20:49:22.374Z"}]}`)
				}))
			})
			It(`Invoke ListBuckets successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListBuckets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBucketsOptions model
				listBucketsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListBucketsOptions)
				listBucketsOptionsModel.IbmServiceInstanceID = core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")
				listBucketsOptionsModel.Extended = core.BoolPtr(true)
				listBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListBuckets(listBucketsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBuckets with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListBucketsOptions model
				listBucketsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListBucketsOptions)
				listBucketsOptionsModel.IbmServiceInstanceID = core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")
				listBucketsOptionsModel.Extended = core.BoolPtr(true)
				listBucketsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListBuckets(listBucketsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBucketsOptions model with no property values
				listBucketsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.ListBucketsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListBuckets(listBucketsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateBucket(createBucketOptions *CreateBucketOptions)`, func() {
		createBucketPath := "/my-new-bucket"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createBucketPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Ibm-Service-Instance-Id"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Service-Instance-Id"][0]).To(Equal(fmt.Sprintf("%v", "d6f76k03-6k4f-4a82-n165-697654o63903")))
					Expect(req.Header["Ibm-Sse-Kp-Encryption-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Sse-Kp-Encryption-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["Ibm-Sse-Kp-Customer-Root-Key-Crn"]).ToNot(BeNil())
					Expect(req.Header["Ibm-Sse-Kp-Customer-Root-Key-Crn"][0]).To(Equal(fmt.Sprintf("%v", "crn:v1:bluemix:public:kms:us-south:a/f047b55a3362ac06afad8a3f2f5586ea:12e8c9c2-a162-472d-b7d6-8b9a86b815a6:key:02fd6835-6001-4482-a892-13bd2085f75d")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke CreateBucket successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.CreateBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateBucketOptions model
				createBucketOptionsModel := new(ibmcloudobjectstorages3apiv2.CreateBucketOptions)
				createBucketOptionsModel.Bucket = core.StringPtr("my-new-bucket")
				createBucketOptionsModel.IbmServiceInstanceID = core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")
				createBucketOptionsModel.Body = core.StringPtr("testString")
				createBucketOptionsModel.IbmSseKpEncryptionAlgorithm = core.StringPtr("AES256")
				createBucketOptionsModel.IbmSseKpCustomerRootKeyCrn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/f047b55a3362ac06afad8a3f2f5586ea:12e8c9c2-a162-472d-b7d6-8b9a86b815a6:key:02fd6835-6001-4482-a892-13bd2085f75d")
				createBucketOptionsModel.XAmzAcl = core.StringPtr("private")
				createBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.CreateBucket(createBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateBucket with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the CreateBucketOptions model
				createBucketOptionsModel := new(ibmcloudobjectstorages3apiv2.CreateBucketOptions)
				createBucketOptionsModel.Bucket = core.StringPtr("my-new-bucket")
				createBucketOptionsModel.IbmServiceInstanceID = core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")
				createBucketOptionsModel.Body = core.StringPtr("testString")
				createBucketOptionsModel.IbmSseKpEncryptionAlgorithm = core.StringPtr("AES256")
				createBucketOptionsModel.IbmSseKpCustomerRootKeyCrn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/f047b55a3362ac06afad8a3f2f5586ea:12e8c9c2-a162-472d-b7d6-8b9a86b815a6:key:02fd6835-6001-4482-a892-13bd2085f75d")
				createBucketOptionsModel.XAmzAcl = core.StringPtr("private")
				createBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.CreateBucket(createBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CreateBucketOptions model with no property values
				createBucketOptionsModelNew := new(ibmcloudobjectstorages3apiv2.CreateBucketOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.CreateBucket(createBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteBucket(deleteBucketOptions *DeleteBucketOptions)`, func() {
		deleteBucketPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBucketPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBucket successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucket(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBucketOptions model
				deleteBucketOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketOptions)
				deleteBucketOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucket(deleteBucketOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBucket with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteBucketOptions model
				deleteBucketOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketOptions)
				deleteBucketOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucket(deleteBucketOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBucketOptions model with no property values
				deleteBucketOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteBucketOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucket(deleteBucketOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListObjects(listObjectsOptions *ListObjectsOptions) - Operation response error`, func() {
		listObjectsPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-keys"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListObjects with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsOptions)
				listObjectsOptionsModel.Bucket = core.StringPtr("testString")
				listObjectsOptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsOptionsModel.EncodingType = core.StringPtr("url")
				listObjectsOptionsModel.Marker = core.StringPtr("testString")
				listObjectsOptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Prefix = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListObjects(listObjectsOptions *ListObjectsOptions)`, func() {
		listObjectsPath := "/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-keys"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"IsTruncated": false, "Marker": "Marker", "NextMarker": "NextMarker", "Contents": {}, "Name": "Name", "Prefix": "Prefix", "Delimiter": "Delimiter", "MaxKeys": 7, "CommonPrefixes": {}, "EncodingType": "url"}`)
				}))
			})
			It(`Invoke ListObjects successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsOptions)
				listObjectsOptionsModel.Bucket = core.StringPtr("testString")
				listObjectsOptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsOptionsModel.EncodingType = core.StringPtr("url")
				listObjectsOptionsModel.Marker = core.StringPtr("testString")
				listObjectsOptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Prefix = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.ListObjectsWithContext(ctx, listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.ListObjectsWithContext(ctx, listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-keys"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"IsTruncated": false, "Marker": "Marker", "NextMarker": "NextMarker", "Contents": {}, "Name": "Name", "Prefix": "Prefix", "Delimiter": "Delimiter", "MaxKeys": 7, "CommonPrefixes": {}, "EncodingType": "url"}`)
				}))
			})
			It(`Invoke ListObjects successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsOptions)
				listObjectsOptionsModel.Bucket = core.StringPtr("testString")
				listObjectsOptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsOptionsModel.EncodingType = core.StringPtr("url")
				listObjectsOptionsModel.Marker = core.StringPtr("testString")
				listObjectsOptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Prefix = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListObjects with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListObjectsOptions model
				listObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsOptions)
				listObjectsOptionsModel.Bucket = core.StringPtr("testString")
				listObjectsOptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsOptionsModel.EncodingType = core.StringPtr("url")
				listObjectsOptionsModel.Marker = core.StringPtr("testString")
				listObjectsOptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsOptionsModel.Prefix = core.StringPtr("testString")
				listObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjects(listObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListObjectsOptions model with no property values
				listObjectsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.ListObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListObjects(listObjectsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListObjectsV2(listObjectsV2Options *ListObjectsV2Options) - Operation response error`, func() {
		listObjectsV2Path := "/testString?list-type=2"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsV2Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["list-type"]).To(Equal([]string{"2"}))
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["max-keys"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["continuation-token"]).To(Equal([]string{"testString"}))
					// TODO: Add check for fetch-owner query parameter
					Expect(req.URL.Query()["start-after"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListObjectsV2 with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListObjectsV2Options model
				listObjectsV2OptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsV2Options)
				listObjectsV2OptionsModel.Bucket = core.StringPtr("testString")
				listObjectsV2OptionsModel.ListType = core.StringPtr("2")
				listObjectsV2OptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsV2OptionsModel.EncodingType = core.StringPtr("url")
				listObjectsV2OptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsV2OptionsModel.Prefix = core.StringPtr("testString")
				listObjectsV2OptionsModel.ContinuationToken = core.StringPtr("testString")
				listObjectsV2OptionsModel.FetchOwner = core.BoolPtr(true)
				listObjectsV2OptionsModel.StartAfter = core.StringPtr("testString")
				listObjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjectsV2(listObjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListObjectsV2(listObjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListObjectsV2(listObjectsV2Options *ListObjectsV2Options)`, func() {
		listObjectsV2Path := "/testString?list-type=2"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["list-type"]).To(Equal([]string{"2"}))
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["max-keys"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["continuation-token"]).To(Equal([]string{"testString"}))
					// TODO: Add check for fetch-owner query parameter
					Expect(req.URL.Query()["start-after"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"IsTruncated": false, "Contents": {}, "Name": "Name", "Prefix": "Prefix", "Delimiter": "Delimiter", "MaxKeys": 7, "CommonPrefixes": {}, "EncodingType": "url", "KeyCount": 8, "ContinuationToken": "ContinuationToken", "NextContinuationToken": "NextContinuationToken", "StartAfter": "StartAfter"}`)
				}))
			})
			It(`Invoke ListObjectsV2 successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the ListObjectsV2Options model
				listObjectsV2OptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsV2Options)
				listObjectsV2OptionsModel.Bucket = core.StringPtr("testString")
				listObjectsV2OptionsModel.ListType = core.StringPtr("2")
				listObjectsV2OptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsV2OptionsModel.EncodingType = core.StringPtr("url")
				listObjectsV2OptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsV2OptionsModel.Prefix = core.StringPtr("testString")
				listObjectsV2OptionsModel.ContinuationToken = core.StringPtr("testString")
				listObjectsV2OptionsModel.FetchOwner = core.BoolPtr(true)
				listObjectsV2OptionsModel.StartAfter = core.StringPtr("testString")
				listObjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.ListObjectsV2WithContext(ctx, listObjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjectsV2(listObjectsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.ListObjectsV2WithContext(ctx, listObjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listObjectsV2Path))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["list-type"]).To(Equal([]string{"2"}))
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["max-keys"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["continuation-token"]).To(Equal([]string{"testString"}))
					// TODO: Add check for fetch-owner query parameter
					Expect(req.URL.Query()["start-after"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"IsTruncated": false, "Contents": {}, "Name": "Name", "Prefix": "Prefix", "Delimiter": "Delimiter", "MaxKeys": 7, "CommonPrefixes": {}, "EncodingType": "url", "KeyCount": 8, "ContinuationToken": "ContinuationToken", "NextContinuationToken": "NextContinuationToken", "StartAfter": "StartAfter"}`)
				}))
			})
			It(`Invoke ListObjectsV2 successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjectsV2(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListObjectsV2Options model
				listObjectsV2OptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsV2Options)
				listObjectsV2OptionsModel.Bucket = core.StringPtr("testString")
				listObjectsV2OptionsModel.ListType = core.StringPtr("2")
				listObjectsV2OptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsV2OptionsModel.EncodingType = core.StringPtr("url")
				listObjectsV2OptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsV2OptionsModel.Prefix = core.StringPtr("testString")
				listObjectsV2OptionsModel.ContinuationToken = core.StringPtr("testString")
				listObjectsV2OptionsModel.FetchOwner = core.BoolPtr(true)
				listObjectsV2OptionsModel.StartAfter = core.StringPtr("testString")
				listObjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListObjectsV2(listObjectsV2OptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListObjectsV2 with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListObjectsV2Options model
				listObjectsV2OptionsModel := new(ibmcloudobjectstorages3apiv2.ListObjectsV2Options)
				listObjectsV2OptionsModel.Bucket = core.StringPtr("testString")
				listObjectsV2OptionsModel.ListType = core.StringPtr("2")
				listObjectsV2OptionsModel.Delimiter = core.StringPtr("testString")
				listObjectsV2OptionsModel.EncodingType = core.StringPtr("url")
				listObjectsV2OptionsModel.MaxKeys = core.Int64Ptr(int64(38))
				listObjectsV2OptionsModel.Prefix = core.StringPtr("testString")
				listObjectsV2OptionsModel.ContinuationToken = core.StringPtr("testString")
				listObjectsV2OptionsModel.FetchOwner = core.BoolPtr(true)
				listObjectsV2OptionsModel.StartAfter = core.StringPtr("testString")
				listObjectsV2OptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListObjectsV2(listObjectsV2OptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListObjectsV2Options model with no property values
				listObjectsV2OptionsModelNew := new(ibmcloudobjectstorages3apiv2.ListObjectsV2Options)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListObjectsV2(listObjectsV2OptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutObject(putObjectOptions *PutObjectOptions)`, func() {
		putObjectPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putObjectPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Length"]).ToNot(BeNil())
					Expect(req.Header["Content-Length"][0]).To(Equal(fmt.Sprintf("%v", int64(38))))
					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutObject successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutObjectOptions model
				putObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.PutObjectOptions)
				putObjectOptionsModel.Bucket = core.StringPtr("testString")
				putObjectOptionsModel.Key = core.StringPtr("testString")
				putObjectOptionsModel.Body = core.StringPtr("testString")
				putObjectOptionsModel.XAmzAcl = core.StringPtr("private")
				putObjectOptionsModel.IfMatch = core.StringPtr("testString")
				putObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				putObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				putObjectOptionsModel.CacheControl = core.StringPtr("testString")
				putObjectOptionsModel.ContentDisposition = core.StringPtr("testString")
				putObjectOptionsModel.ContentEncoding = core.StringPtr("testString")
				putObjectOptionsModel.ContentLanguage = core.StringPtr("testString")
				putObjectOptionsModel.ContentLength = core.Int64Ptr(int64(38))
				putObjectOptionsModel.ContentMD5 = core.StringPtr("testString")
				putObjectOptionsModel.Expires = CreateMockDateTime()
				putObjectOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				putObjectOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				putObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				putObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				putObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				putObjectOptionsModel.XAmzTagging = core.StringPtr("testString")
				putObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutObject(putObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutObject with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutObjectOptions model
				putObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.PutObjectOptions)
				putObjectOptionsModel.Bucket = core.StringPtr("testString")
				putObjectOptionsModel.Key = core.StringPtr("testString")
				putObjectOptionsModel.Body = core.StringPtr("testString")
				putObjectOptionsModel.XAmzAcl = core.StringPtr("private")
				putObjectOptionsModel.IfMatch = core.StringPtr("testString")
				putObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				putObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				putObjectOptionsModel.CacheControl = core.StringPtr("testString")
				putObjectOptionsModel.ContentDisposition = core.StringPtr("testString")
				putObjectOptionsModel.ContentEncoding = core.StringPtr("testString")
				putObjectOptionsModel.ContentLanguage = core.StringPtr("testString")
				putObjectOptionsModel.ContentLength = core.Int64Ptr(int64(38))
				putObjectOptionsModel.ContentMD5 = core.StringPtr("testString")
				putObjectOptionsModel.Expires = CreateMockDateTime()
				putObjectOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				putObjectOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				putObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				putObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				putObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				putObjectOptionsModel.XAmzTagging = core.StringPtr("testString")
				putObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutObject(putObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutObjectOptions model with no property values
				putObjectOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutObjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutObject(putObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetObject(getObjectOptions *GetObjectOptions) - Operation response error`, func() {
		getObjectPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Range"]).ToNot(BeNil())
					Expect(req.Header["Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["response-cache-control"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-disposition"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-encoding"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-language"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for response-expires query parameter
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetObject with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectOptions)
				getObjectOptionsModel.Bucket = core.StringPtr("testString")
				getObjectOptionsModel.Key = core.StringPtr("testString")
				getObjectOptionsModel.IfMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfModifiedSince = CreateMockDateTime()
				getObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				getObjectOptionsModel.Range = core.StringPtr("testString")
				getObjectOptionsModel.ResponseCacheControl = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentDisposition = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentEncoding = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentLanguage = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentType = core.StringPtr("testString")
				getObjectOptionsModel.ResponseExpires = CreateMockDateTime()
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				getObjectOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObject(getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObject(getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetObject(getObjectOptions *GetObjectOptions)`, func() {
		getObjectPath := "/testString/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Range"]).ToNot(BeNil())
					Expect(req.Header["Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["response-cache-control"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-disposition"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-encoding"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-language"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for response-expires query parameter
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Body": "Body", "Metadata": {}}`)
				}))
			})
			It(`Invoke GetObject successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectOptions)
				getObjectOptionsModel.Bucket = core.StringPtr("testString")
				getObjectOptionsModel.Key = core.StringPtr("testString")
				getObjectOptionsModel.IfMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfModifiedSince = CreateMockDateTime()
				getObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				getObjectOptionsModel.Range = core.StringPtr("testString")
				getObjectOptionsModel.ResponseCacheControl = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentDisposition = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentEncoding = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentLanguage = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentType = core.StringPtr("testString")
				getObjectOptionsModel.ResponseExpires = CreateMockDateTime()
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				getObjectOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectWithContext(ctx, getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObject(getObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectWithContext(ctx, getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Range"]).ToNot(BeNil())
					Expect(req.Header["Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["response-cache-control"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-disposition"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-encoding"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-language"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["response-content-type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for response-expires query parameter
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Body": "Body", "Metadata": {}}`)
				}))
			})
			It(`Invoke GetObject successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectOptions)
				getObjectOptionsModel.Bucket = core.StringPtr("testString")
				getObjectOptionsModel.Key = core.StringPtr("testString")
				getObjectOptionsModel.IfMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfModifiedSince = CreateMockDateTime()
				getObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				getObjectOptionsModel.Range = core.StringPtr("testString")
				getObjectOptionsModel.ResponseCacheControl = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentDisposition = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentEncoding = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentLanguage = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentType = core.StringPtr("testString")
				getObjectOptionsModel.ResponseExpires = CreateMockDateTime()
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				getObjectOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObject(getObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetObject with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetObjectOptions model
				getObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectOptions)
				getObjectOptionsModel.Bucket = core.StringPtr("testString")
				getObjectOptionsModel.Key = core.StringPtr("testString")
				getObjectOptionsModel.IfMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfModifiedSince = CreateMockDateTime()
				getObjectOptionsModel.IfNoneMatch = core.StringPtr("testString")
				getObjectOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				getObjectOptionsModel.Range = core.StringPtr("testString")
				getObjectOptionsModel.ResponseCacheControl = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentDisposition = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentEncoding = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentLanguage = core.StringPtr("testString")
				getObjectOptionsModel.ResponseContentType = core.StringPtr("testString")
				getObjectOptionsModel.ResponseExpires = CreateMockDateTime()
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				getObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				getObjectOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				getObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObject(getObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetObjectOptions model with no property values
				getObjectOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetObjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObject(getObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteObject(deleteObjectOptions *DeleteObjectOptions)`, func() {
		deleteObjectPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteObject successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteObjectOptions model
				deleteObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectOptions)
				deleteObjectOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectOptionsModel.Key = core.StringPtr("testString")
				deleteObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteObject with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteObjectOptions model
				deleteObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectOptions)
				deleteObjectOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectOptionsModel.Key = core.StringPtr("testString")
				deleteObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObject(deleteObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteObjectOptions model with no property values
				deleteObjectOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteObjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObject(deleteObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CopyObject(copyObjectOptions *CopyObjectOptions) - Operation response error`, func() {
		copyObjectPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(copyObjectPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Amz-Copy-Source"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Metadata-Directive"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Metadata-Directive"][0]).To(Equal(fmt.Sprintf("%v", "COPY")))
					Expect(req.Header["X-Amz-Tagging-Directive"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging-Directive"][0]).To(Equal(fmt.Sprintf("%v", "COPY")))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CopyObject with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the CopyObjectOptions model
				copyObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.CopyObjectOptions)
				copyObjectOptionsModel.Bucket = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySource = core.StringPtr("testString")
				copyObjectOptionsModel.TargetKey = core.StringPtr("testString")
				copyObjectOptionsModel.Body = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzAcl = core.StringPtr("private")
				copyObjectOptionsModel.CacheControl = core.StringPtr("testString")
				copyObjectOptionsModel.ContentDisposition = core.StringPtr("testString")
				copyObjectOptionsModel.ContentEncoding = core.StringPtr("testString")
				copyObjectOptionsModel.ContentLanguage = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.Expires = CreateMockDateTime()
				copyObjectOptionsModel.XAmzMetadataDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzTaggingDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				copyObjectOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzTagging = core.StringPtr("testString")
				copyObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CopyObject(copyObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.CopyObject(copyObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CopyObject(copyObjectOptions *CopyObjectOptions)`, func() {
		copyObjectPath := "/testString/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(copyObjectPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Amz-Copy-Source"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Metadata-Directive"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Metadata-Directive"][0]).To(Equal(fmt.Sprintf("%v", "COPY")))
					Expect(req.Header["X-Amz-Tagging-Directive"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging-Directive"][0]).To(Equal(fmt.Sprintf("%v", "COPY")))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"CopyObjectResult": {"ETag": "ETag", "LastModified": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CopyObject successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the CopyObjectOptions model
				copyObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.CopyObjectOptions)
				copyObjectOptionsModel.Bucket = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySource = core.StringPtr("testString")
				copyObjectOptionsModel.TargetKey = core.StringPtr("testString")
				copyObjectOptionsModel.Body = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzAcl = core.StringPtr("private")
				copyObjectOptionsModel.CacheControl = core.StringPtr("testString")
				copyObjectOptionsModel.ContentDisposition = core.StringPtr("testString")
				copyObjectOptionsModel.ContentEncoding = core.StringPtr("testString")
				copyObjectOptionsModel.ContentLanguage = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.Expires = CreateMockDateTime()
				copyObjectOptionsModel.XAmzMetadataDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzTaggingDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				copyObjectOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzTagging = core.StringPtr("testString")
				copyObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.CopyObjectWithContext(ctx, copyObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CopyObject(copyObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.CopyObjectWithContext(ctx, copyObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(copyObjectPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Amz-Copy-Source"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Metadata-Directive"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Metadata-Directive"][0]).To(Equal(fmt.Sprintf("%v", "COPY")))
					Expect(req.Header["X-Amz-Tagging-Directive"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging-Directive"][0]).To(Equal(fmt.Sprintf("%v", "COPY")))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"CopyObjectResult": {"ETag": "ETag", "LastModified": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CopyObject successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CopyObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CopyObjectOptions model
				copyObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.CopyObjectOptions)
				copyObjectOptionsModel.Bucket = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySource = core.StringPtr("testString")
				copyObjectOptionsModel.TargetKey = core.StringPtr("testString")
				copyObjectOptionsModel.Body = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzAcl = core.StringPtr("private")
				copyObjectOptionsModel.CacheControl = core.StringPtr("testString")
				copyObjectOptionsModel.ContentDisposition = core.StringPtr("testString")
				copyObjectOptionsModel.ContentEncoding = core.StringPtr("testString")
				copyObjectOptionsModel.ContentLanguage = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.Expires = CreateMockDateTime()
				copyObjectOptionsModel.XAmzMetadataDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzTaggingDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				copyObjectOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzTagging = core.StringPtr("testString")
				copyObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.CopyObject(copyObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CopyObject with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the CopyObjectOptions model
				copyObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.CopyObjectOptions)
				copyObjectOptionsModel.Bucket = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySource = core.StringPtr("testString")
				copyObjectOptionsModel.TargetKey = core.StringPtr("testString")
				copyObjectOptionsModel.Body = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzAcl = core.StringPtr("private")
				copyObjectOptionsModel.CacheControl = core.StringPtr("testString")
				copyObjectOptionsModel.ContentDisposition = core.StringPtr("testString")
				copyObjectOptionsModel.ContentEncoding = core.StringPtr("testString")
				copyObjectOptionsModel.ContentLanguage = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				copyObjectOptionsModel.Expires = CreateMockDateTime()
				copyObjectOptionsModel.XAmzMetadataDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzTaggingDirective = core.StringPtr("COPY")
				copyObjectOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				copyObjectOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				copyObjectOptionsModel.XAmzTagging = core.StringPtr("testString")
				copyObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CopyObject(copyObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CopyObjectOptions model with no property values
				copyObjectOptionsModelNew := new(ibmcloudobjectstorages3apiv2.CopyObjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.CopyObject(copyObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteObjects(deleteObjectsOptions *DeleteObjectsOptions) - Operation response error`, func() {
		deleteObjectsPath := "/testString?delete"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectsPath))
					Expect(req.Method).To(Equal("POST"))
					// TODO: Add check for delete query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteObjects with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteObjectsOptions model
				deleteObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectsOptions)
				deleteObjectsOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectsOptionsModel.Delete = core.BoolPtr(true)
				deleteObjectsOptionsModel.Body = core.StringPtr("testString")
				deleteObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjects(deleteObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObjects(deleteObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteObjects(deleteObjectsOptions *DeleteObjectsOptions)`, func() {
		deleteObjectsPath := "/testString?delete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for delete query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Errors": {}}`)
				}))
			})
			It(`Invoke DeleteObjects successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteObjectsOptions model
				deleteObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectsOptions)
				deleteObjectsOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectsOptionsModel.Delete = core.BoolPtr(true)
				deleteObjectsOptionsModel.Body = core.StringPtr("testString")
				deleteObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjectsWithContext(ctx, deleteObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjects(deleteObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObjectsWithContext(ctx, deleteObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for delete query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Errors": {}}`)
				}))
			})
			It(`Invoke DeleteObjects successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteObjectsOptions model
				deleteObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectsOptions)
				deleteObjectsOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectsOptionsModel.Delete = core.BoolPtr(true)
				deleteObjectsOptionsModel.Body = core.StringPtr("testString")
				deleteObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObjects(deleteObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteObjects with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteObjectsOptions model
				deleteObjectsOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectsOptions)
				deleteObjectsOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectsOptionsModel.Delete = core.BoolPtr(true)
				deleteObjectsOptionsModel.Body = core.StringPtr("testString")
				deleteObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjects(deleteObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteObjectsOptions model with no property values
				deleteObjectsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObjects(deleteObjectsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`PutBucketProtectionConfiguration(putBucketProtectionConfigurationOptions *PutBucketProtectionConfigurationOptions)`, func() {
		putBucketProtectionConfigurationPath := "/testString?protection"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putBucketProtectionConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for protection query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutBucketProtectionConfiguration successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketProtectionConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutBucketProtectionConfigurationOptions model
				putBucketProtectionConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketProtectionConfigurationOptions)
				putBucketProtectionConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				putBucketProtectionConfigurationOptionsModel.Protection = core.BoolPtr(true)
				putBucketProtectionConfigurationOptionsModel.Body = core.StringPtr("testString")
				putBucketProtectionConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketProtectionConfiguration(putBucketProtectionConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutBucketProtectionConfiguration with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutBucketProtectionConfigurationOptions model
				putBucketProtectionConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketProtectionConfigurationOptions)
				putBucketProtectionConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				putBucketProtectionConfigurationOptionsModel.Protection = core.BoolPtr(true)
				putBucketProtectionConfigurationOptionsModel.Body = core.StringPtr("testString")
				putBucketProtectionConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketProtectionConfiguration(putBucketProtectionConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutBucketProtectionConfigurationOptions model with no property values
				putBucketProtectionConfigurationOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutBucketProtectionConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketProtectionConfiguration(putBucketProtectionConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`PutBucketLifecycleConfiguration(putBucketLifecycleConfigurationOptions *PutBucketLifecycleConfigurationOptions)`, func() {
		putBucketLifecycleConfigurationPath := "/testString?lifecycle"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putBucketLifecycleConfigurationPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for lifecycle query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutBucketLifecycleConfiguration successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketLifecycleConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutBucketLifecycleConfigurationOptions model
				putBucketLifecycleConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketLifecycleConfigurationOptions)
				putBucketLifecycleConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				putBucketLifecycleConfigurationOptionsModel.Lifecycle = core.BoolPtr(true)
				putBucketLifecycleConfigurationOptionsModel.Body = core.StringPtr("testString")
				putBucketLifecycleConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketLifecycleConfiguration(putBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutBucketLifecycleConfiguration with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutBucketLifecycleConfigurationOptions model
				putBucketLifecycleConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketLifecycleConfigurationOptions)
				putBucketLifecycleConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				putBucketLifecycleConfigurationOptionsModel.Lifecycle = core.BoolPtr(true)
				putBucketLifecycleConfigurationOptionsModel.Body = core.StringPtr("testString")
				putBucketLifecycleConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketLifecycleConfiguration(putBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutBucketLifecycleConfigurationOptions model with no property values
				putBucketLifecycleConfigurationOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutBucketLifecycleConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketLifecycleConfiguration(putBucketLifecycleConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptions *GetBucketLifecycleConfigurationOptions) - Operation response error`, func() {
		getBucketLifecycleConfigurationPath := "/testString?lifecycle"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketLifecycleConfigurationPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for lifecycle query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketLifecycleConfiguration with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketLifecycleConfigurationOptions model
				getBucketLifecycleConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketLifecycleConfigurationOptions)
				getBucketLifecycleConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				getBucketLifecycleConfigurationOptionsModel.Lifecycle = core.BoolPtr(true)
				getBucketLifecycleConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptions *GetBucketLifecycleConfigurationOptions)`, func() {
		getBucketLifecycleConfigurationPath := "/testString?lifecycle"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketLifecycleConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for lifecycle query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Rules": {}}`)
				}))
			})
			It(`Invoke GetBucketLifecycleConfiguration successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketLifecycleConfigurationOptions model
				getBucketLifecycleConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketLifecycleConfigurationOptions)
				getBucketLifecycleConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				getBucketLifecycleConfigurationOptionsModel.Lifecycle = core.BoolPtr(true)
				getBucketLifecycleConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfigurationWithContext(ctx, getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfigurationWithContext(ctx, getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketLifecycleConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for lifecycle query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Rules": {}}`)
				}))
			})
			It(`Invoke GetBucketLifecycleConfiguration successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketLifecycleConfigurationOptions model
				getBucketLifecycleConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketLifecycleConfigurationOptions)
				getBucketLifecycleConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				getBucketLifecycleConfigurationOptionsModel.Lifecycle = core.BoolPtr(true)
				getBucketLifecycleConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketLifecycleConfiguration with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketLifecycleConfigurationOptions model
				getBucketLifecycleConfigurationOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketLifecycleConfigurationOptions)
				getBucketLifecycleConfigurationOptionsModel.Bucket = core.StringPtr("testString")
				getBucketLifecycleConfigurationOptionsModel.Lifecycle = core.BoolPtr(true)
				getBucketLifecycleConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketLifecycleConfigurationOptions model with no property values
				getBucketLifecycleConfigurationOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetBucketLifecycleConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteBucketLifecycle(deleteBucketLifecycleOptions *DeleteBucketLifecycleOptions)`, func() {
		deleteBucketLifecyclePath := "/testString?lifecycle"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBucketLifecyclePath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for lifecycle query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBucketLifecycle successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucketLifecycle(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBucketLifecycleOptions model
				deleteBucketLifecycleOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketLifecycleOptions)
				deleteBucketLifecycleOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketLifecycleOptionsModel.Lifecycle = core.BoolPtr(true)
				deleteBucketLifecycleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucketLifecycle(deleteBucketLifecycleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBucketLifecycle with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteBucketLifecycleOptions model
				deleteBucketLifecycleOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketLifecycleOptions)
				deleteBucketLifecycleOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketLifecycleOptionsModel.Lifecycle = core.BoolPtr(true)
				deleteBucketLifecycleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucketLifecycle(deleteBucketLifecycleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBucketLifecycleOptions model with no property values
				deleteBucketLifecycleOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteBucketLifecycleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucketLifecycle(deleteBucketLifecycleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RestoreObject(restoreObjectOptions *RestoreObjectOptions)`, func() {
		restoreObjectPath := "/testString/testString?restore"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreObjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// TODO: Add check for restore query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke RestoreObject successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.RestoreObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RestoreObjectOptions model
				restoreObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.RestoreObjectOptions)
				restoreObjectOptionsModel.Bucket = core.StringPtr("testString")
				restoreObjectOptionsModel.Key = core.StringPtr("testString")
				restoreObjectOptionsModel.Restore = core.BoolPtr(true)
				restoreObjectOptionsModel.Body = core.StringPtr("testString")
				restoreObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.RestoreObject(restoreObjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RestoreObject with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the RestoreObjectOptions model
				restoreObjectOptionsModel := new(ibmcloudobjectstorages3apiv2.RestoreObjectOptions)
				restoreObjectOptionsModel.Bucket = core.StringPtr("testString")
				restoreObjectOptionsModel.Key = core.StringPtr("testString")
				restoreObjectOptionsModel.Restore = core.BoolPtr(true)
				restoreObjectOptionsModel.Body = core.StringPtr("testString")
				restoreObjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.RestoreObject(restoreObjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RestoreObjectOptions model with no property values
				restoreObjectOptionsModelNew := new(ibmcloudobjectstorages3apiv2.RestoreObjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.RestoreObject(restoreObjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`InitiateMultipartUpload(initiateMultipartUploadOptions *InitiateMultipartUploadOptions) - Operation response error`, func() {
		initiateMultipartUploadPath := "/testString/testString?uploads"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initiateMultipartUploadPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					// TODO: Add check for uploads query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke InitiateMultipartUpload with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the InitiateMultipartUploadOptions model
				initiateMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.InitiateMultipartUploadOptions)
				initiateMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Uploads = core.BoolPtr(true)
				initiateMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfNoneMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.CacheControl = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentDisposition = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentEncoding = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentLanguage = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Expires = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				initiateMultipartUploadOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzTagging = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzAcl = core.StringPtr("private")
				initiateMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(initiateMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(initiateMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`InitiateMultipartUpload(initiateMultipartUploadOptions *InitiateMultipartUploadOptions)`, func() {
		initiateMultipartUploadPath := "/testString/testString?uploads"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initiateMultipartUploadPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					// TODO: Add check for uploads query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Bucket": "Bucket", "Key": "Key", "UploadId": "UploadID"}`)
				}))
			})
			It(`Invoke InitiateMultipartUpload successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the InitiateMultipartUploadOptions model
				initiateMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.InitiateMultipartUploadOptions)
				initiateMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Uploads = core.BoolPtr(true)
				initiateMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfNoneMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.CacheControl = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentDisposition = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentEncoding = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentLanguage = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Expires = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				initiateMultipartUploadOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzTagging = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzAcl = core.StringPtr("private")
				initiateMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.InitiateMultipartUploadWithContext(ctx, initiateMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(initiateMultipartUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.InitiateMultipartUploadWithContext(ctx, initiateMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(initiateMultipartUploadPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["Cache-Control"]).ToNot(BeNil())
					Expect(req.Header["Cache-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Disposition"]).ToNot(BeNil())
					Expect(req.Header["Content-Disposition"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Encoding"]).ToNot(BeNil())
					Expect(req.Header["Content-Encoding"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Content-Language"]).ToNot(BeNil())
					Expect(req.Header["Content-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Expires"]).ToNot(BeNil())
					Expect(req.Header["Expires"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Server-Side-Encryption"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption"][0]).To(Equal(fmt.Sprintf("%v", "AES256")))
					Expect(req.Header["X-Amz-Website-Redirect-Location"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Website-Redirect-Location"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Tagging"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Tagging"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					// TODO: Add check for uploads query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Bucket": "Bucket", "Key": "Key", "UploadId": "UploadID"}`)
				}))
			})
			It(`Invoke InitiateMultipartUpload successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InitiateMultipartUploadOptions model
				initiateMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.InitiateMultipartUploadOptions)
				initiateMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Uploads = core.BoolPtr(true)
				initiateMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfNoneMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.CacheControl = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentDisposition = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentEncoding = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentLanguage = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Expires = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				initiateMultipartUploadOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzTagging = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzAcl = core.StringPtr("private")
				initiateMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(initiateMultipartUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke InitiateMultipartUpload with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the InitiateMultipartUploadOptions model
				initiateMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.InitiateMultipartUploadOptions)
				initiateMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Uploads = core.BoolPtr(true)
				initiateMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfNoneMatch = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.IfUnmodifiedSince = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.CacheControl = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentDisposition = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentEncoding = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.ContentLanguage = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.Expires = CreateMockDateTime()
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryption = core.StringPtr("AES256")
				initiateMultipartUploadOptionsModel.XAmzWebsiteRedirectLocation = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzTagging = core.StringPtr("testString")
				initiateMultipartUploadOptionsModel.XAmzAcl = core.StringPtr("private")
				initiateMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(initiateMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the InitiateMultipartUploadOptions model with no property values
				initiateMultipartUploadOptionsModelNew := new(ibmcloudobjectstorages3apiv2.InitiateMultipartUploadOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.InitiateMultipartUpload(initiateMultipartUploadOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteMultipartUpload(completeMultipartUploadOptions *CompleteMultipartUploadOptions) - Operation response error`, func() {
		completeMultipartUploadPath := "/testString/testString?uploadId={uploadId}"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeMultipartUploadPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteMultipartUpload with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the CompleteMultipartUploadOptions model
				completeMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.CompleteMultipartUploadOptions)
				completeMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.UploadID = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(completeMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(completeMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CompleteMultipartUpload(completeMultipartUploadOptions *CompleteMultipartUploadOptions)`, func() {
		completeMultipartUploadPath := "/testString/testString?uploadId={uploadId}"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeMultipartUploadPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Location": "Location", "Bucket": "Bucket", "Key": "Key", "ETag": "ETag"}`)
				}))
			})
			It(`Invoke CompleteMultipartUpload successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the CompleteMultipartUploadOptions model
				completeMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.CompleteMultipartUploadOptions)
				completeMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.UploadID = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.CompleteMultipartUploadWithContext(ctx, completeMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(completeMultipartUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.CompleteMultipartUploadWithContext(ctx, completeMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeMultipartUploadPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Location": "Location", "Bucket": "Bucket", "Key": "Key", "ETag": "ETag"}`)
				}))
			})
			It(`Invoke CompleteMultipartUpload successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteMultipartUploadOptions model
				completeMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.CompleteMultipartUploadOptions)
				completeMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.UploadID = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(completeMultipartUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteMultipartUpload with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the CompleteMultipartUploadOptions model
				completeMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.CompleteMultipartUploadOptions)
				completeMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.UploadID = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Body = core.StringPtr("testString")
				completeMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(completeMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteMultipartUploadOptions model with no property values
				completeMultipartUploadOptionsModelNew := new(ibmcloudobjectstorages3apiv2.CompleteMultipartUploadOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.CompleteMultipartUpload(completeMultipartUploadOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListParts(listPartsOptions *ListPartsOptions) - Operation response error`, func() {
		listPartsPath := "/testString/testString?uploadId={uploadId}"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPartsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-parts"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["part-number-marker"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListParts with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListPartsOptions model
				listPartsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListPartsOptions)
				listPartsOptionsModel.Bucket = core.StringPtr("testString")
				listPartsOptionsModel.Key = core.StringPtr("testString")
				listPartsOptionsModel.UploadID = core.StringPtr("testString")
				listPartsOptionsModel.MaxParts = core.Int64Ptr(int64(38))
				listPartsOptionsModel.PartNumberMarker = core.Int64Ptr(int64(38))
				listPartsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListParts(listPartsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListParts(listPartsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListParts(listPartsOptions *ListPartsOptions)`, func() {
		listPartsPath := "/testString/testString?uploadId={uploadId}"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPartsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-parts"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["part-number-marker"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Bucket": "Bucket", "Key": "Key", "UploadId": "UploadID", "PartNumberMarker": 16, "NextPartNumberMarker": 20, "MaxParts": 8, "IsTruncated": false, "Parts": {}, "Initiator": {"ID": "ID", "DisplayName": "DisplayName"}, "Owner": {"DisplayName": "DisplayName", "ID": "ID"}, "StorageClass": "STANDARD"}`)
				}))
			})
			It(`Invoke ListParts successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the ListPartsOptions model
				listPartsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListPartsOptions)
				listPartsOptionsModel.Bucket = core.StringPtr("testString")
				listPartsOptionsModel.Key = core.StringPtr("testString")
				listPartsOptionsModel.UploadID = core.StringPtr("testString")
				listPartsOptionsModel.MaxParts = core.Int64Ptr(int64(38))
				listPartsOptionsModel.PartNumberMarker = core.Int64Ptr(int64(38))
				listPartsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.ListPartsWithContext(ctx, listPartsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListParts(listPartsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.ListPartsWithContext(ctx, listPartsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPartsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-parts"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["part-number-marker"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Bucket": "Bucket", "Key": "Key", "UploadId": "UploadID", "PartNumberMarker": 16, "NextPartNumberMarker": 20, "MaxParts": 8, "IsTruncated": false, "Parts": {}, "Initiator": {"ID": "ID", "DisplayName": "DisplayName"}, "Owner": {"DisplayName": "DisplayName", "ID": "ID"}, "StorageClass": "STANDARD"}`)
				}))
			})
			It(`Invoke ListParts successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListParts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPartsOptions model
				listPartsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListPartsOptions)
				listPartsOptionsModel.Bucket = core.StringPtr("testString")
				listPartsOptionsModel.Key = core.StringPtr("testString")
				listPartsOptionsModel.UploadID = core.StringPtr("testString")
				listPartsOptionsModel.MaxParts = core.Int64Ptr(int64(38))
				listPartsOptionsModel.PartNumberMarker = core.Int64Ptr(int64(38))
				listPartsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListParts(listPartsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListParts with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListPartsOptions model
				listPartsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListPartsOptions)
				listPartsOptionsModel.Bucket = core.StringPtr("testString")
				listPartsOptionsModel.Key = core.StringPtr("testString")
				listPartsOptionsModel.UploadID = core.StringPtr("testString")
				listPartsOptionsModel.MaxParts = core.Int64Ptr(int64(38))
				listPartsOptionsModel.PartNumberMarker = core.Int64Ptr(int64(38))
				listPartsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListParts(listPartsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPartsOptions model with no property values
				listPartsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.ListPartsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListParts(listPartsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AbortMultipartUpload(abortMultipartUploadOptions *AbortMultipartUploadOptions)`, func() {
		abortMultipartUploadPath := "/testString/testString?uploadId={uploadId}"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(abortMultipartUploadPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke AbortMultipartUpload successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.AbortMultipartUpload(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AbortMultipartUploadOptions model
				abortMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.AbortMultipartUploadOptions)
				abortMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				abortMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				abortMultipartUploadOptionsModel.UploadID = core.StringPtr("testString")
				abortMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.AbortMultipartUpload(abortMultipartUploadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AbortMultipartUpload with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the AbortMultipartUploadOptions model
				abortMultipartUploadOptionsModel := new(ibmcloudobjectstorages3apiv2.AbortMultipartUploadOptions)
				abortMultipartUploadOptionsModel.Bucket = core.StringPtr("testString")
				abortMultipartUploadOptionsModel.Key = core.StringPtr("testString")
				abortMultipartUploadOptionsModel.UploadID = core.StringPtr("testString")
				abortMultipartUploadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.AbortMultipartUpload(abortMultipartUploadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AbortMultipartUploadOptions model with no property values
				abortMultipartUploadOptionsModelNew := new(ibmcloudobjectstorages3apiv2.AbortMultipartUploadOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.AbortMultipartUpload(abortMultipartUploadOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListMultipartUploads(listMultipartUploadsOptions *ListMultipartUploadsOptions) - Operation response error`, func() {
		listMultipartUploadsPath := "/testString?uploads"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMultipartUploadsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for uploads query parameter
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["key-marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-uploads"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["upload-id-marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationLimit"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationToken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationToken"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListMultipartUploads with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListMultipartUploadsOptions model
				listMultipartUploadsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListMultipartUploadsOptions)
				listMultipartUploadsOptionsModel.Bucket = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Uploads = core.BoolPtr(true)
				listMultipartUploadsOptionsModel.Delimiter = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.EncodingType = core.StringPtr("url")
				listMultipartUploadsOptionsModel.KeyMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.MaxUploads = core.Int64Ptr(int64(38))
				listMultipartUploadsOptionsModel.Prefix = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.UploadIdMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationLimit = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListMultipartUploads(listMultipartUploadsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListMultipartUploads(listMultipartUploadsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListMultipartUploads(listMultipartUploadsOptions *ListMultipartUploadsOptions)`, func() {
		listMultipartUploadsPath := "/testString?uploads"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMultipartUploadsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for uploads query parameter
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["key-marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-uploads"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["upload-id-marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationLimit"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationToken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationToken"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Bucket": "Bucket", "KeyMarker": "KeyMarker", "UploadIdMarker": "UploadIdMarker", "NextKeyMarker": "NextKeyMarker", "Prefix": "Prefix", "Delimiter": "Delimiter", "NextUploadIdMarker": "NextUploadIdMarker", "MaxUploads": 10, "IsTruncated": false, "Uploads": {}, "CommonPrefixes": {}, "EncodingType": "url"}`)
				}))
			})
			It(`Invoke ListMultipartUploads successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the ListMultipartUploadsOptions model
				listMultipartUploadsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListMultipartUploadsOptions)
				listMultipartUploadsOptionsModel.Bucket = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Uploads = core.BoolPtr(true)
				listMultipartUploadsOptionsModel.Delimiter = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.EncodingType = core.StringPtr("url")
				listMultipartUploadsOptionsModel.KeyMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.MaxUploads = core.Int64Ptr(int64(38))
				listMultipartUploadsOptionsModel.Prefix = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.UploadIdMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationLimit = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.ListMultipartUploadsWithContext(ctx, listMultipartUploadsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListMultipartUploads(listMultipartUploadsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.ListMultipartUploadsWithContext(ctx, listMultipartUploadsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMultipartUploadsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for uploads query parameter
					Expect(req.URL.Query()["delimiter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["encoding-type"]).To(Equal([]string{"url"}))
					Expect(req.URL.Query()["key-marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["max-uploads"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["prefix"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["upload-id-marker"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationLimit"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationToken"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["PaginationToken"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Bucket": "Bucket", "KeyMarker": "KeyMarker", "UploadIdMarker": "UploadIdMarker", "NextKeyMarker": "NextKeyMarker", "Prefix": "Prefix", "Delimiter": "Delimiter", "NextUploadIdMarker": "NextUploadIdMarker", "MaxUploads": 10, "IsTruncated": false, "Uploads": {}, "CommonPrefixes": {}, "EncodingType": "url"}`)
				}))
			})
			It(`Invoke ListMultipartUploads successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListMultipartUploads(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListMultipartUploadsOptions model
				listMultipartUploadsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListMultipartUploadsOptions)
				listMultipartUploadsOptionsModel.Bucket = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Uploads = core.BoolPtr(true)
				listMultipartUploadsOptionsModel.Delimiter = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.EncodingType = core.StringPtr("url")
				listMultipartUploadsOptionsModel.KeyMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.MaxUploads = core.Int64Ptr(int64(38))
				listMultipartUploadsOptionsModel.Prefix = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.UploadIdMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationLimit = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListMultipartUploads(listMultipartUploadsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListMultipartUploads with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the ListMultipartUploadsOptions model
				listMultipartUploadsOptionsModel := new(ibmcloudobjectstorages3apiv2.ListMultipartUploadsOptions)
				listMultipartUploadsOptionsModel.Bucket = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Uploads = core.BoolPtr(true)
				listMultipartUploadsOptionsModel.Delimiter = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.EncodingType = core.StringPtr("url")
				listMultipartUploadsOptionsModel.KeyMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.MaxUploads = core.Int64Ptr(int64(38))
				listMultipartUploadsOptionsModel.Prefix = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.UploadIdMarker = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationLimit = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.PaginationToken = core.StringPtr("testString")
				listMultipartUploadsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.ListMultipartUploads(listMultipartUploadsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListMultipartUploadsOptions model with no property values
				listMultipartUploadsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.ListMultipartUploadsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.ListMultipartUploads(listMultipartUploadsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadPart(uploadPartOptions *UploadPartOptions)`, func() {
		uploadPartPath := "/testString/testString?partNumber={partNumber}&uploadId={uploadId}"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadPartPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Length"]).ToNot(BeNil())
					Expect(req.Header["Content-Length"][0]).To(Equal(fmt.Sprintf("%v", int64(38))))
					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Request-Payer"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Request-Payer"][0]).To(Equal(fmt.Sprintf("%v", "requester")))
					Expect(req.Header["X-Amz-Expected-Bucket-Owner"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Expected-Bucket-Owner"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UploadPart successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.UploadPart(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UploadPartOptions model
				uploadPartOptionsModel := new(ibmcloudobjectstorages3apiv2.UploadPartOptions)
				uploadPartOptionsModel.Bucket = core.StringPtr("testString")
				uploadPartOptionsModel.Key = core.StringPtr("testString")
				uploadPartOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				uploadPartOptionsModel.UploadID = core.StringPtr("testString")
				uploadPartOptionsModel.Body = core.StringPtr("testString")
				uploadPartOptionsModel.ContentLength = core.Int64Ptr(int64(38))
				uploadPartOptionsModel.ContentMD5 = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzRequestPayer = core.StringPtr("requester")
				uploadPartOptionsModel.XAmzExpectedBucketOwner = core.StringPtr("testString")
				uploadPartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.UploadPart(uploadPartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UploadPart with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the UploadPartOptions model
				uploadPartOptionsModel := new(ibmcloudobjectstorages3apiv2.UploadPartOptions)
				uploadPartOptionsModel.Bucket = core.StringPtr("testString")
				uploadPartOptionsModel.Key = core.StringPtr("testString")
				uploadPartOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				uploadPartOptionsModel.UploadID = core.StringPtr("testString")
				uploadPartOptionsModel.Body = core.StringPtr("testString")
				uploadPartOptionsModel.ContentLength = core.Int64Ptr(int64(38))
				uploadPartOptionsModel.ContentMD5 = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartOptionsModel.XAmzRequestPayer = core.StringPtr("requester")
				uploadPartOptionsModel.XAmzExpectedBucketOwner = core.StringPtr("testString")
				uploadPartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.UploadPart(uploadPartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UploadPartOptions model with no property values
				uploadPartOptionsModelNew := new(ibmcloudobjectstorages3apiv2.UploadPartOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.UploadPart(uploadPartOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadPartCopy(uploadPartCopyOptions *UploadPartCopyOptions) - Operation response error`, func() {
		uploadPartCopyPath := "/testString/{TargetKey}?partNumber={partNumber}&uploadId={uploadId}"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadPartCopyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["X-Amz-Copy-Source"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-Range"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadPartCopy with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the UploadPartCopyOptions model
				uploadPartCopyOptionsModel := new(ibmcloudobjectstorages3apiv2.UploadPartCopyOptions)
				uploadPartCopyOptionsModel.Bucket = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySource = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Key = core.StringPtr("testString")
				uploadPartCopyOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				uploadPartCopyOptionsModel.UploadID = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceRange = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.UploadPartCopy(uploadPartCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.UploadPartCopy(uploadPartCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadPartCopy(uploadPartCopyOptions *UploadPartCopyOptions)`, func() {
		uploadPartCopyPath := "/testString/{TargetKey}?partNumber={partNumber}&uploadId={uploadId}"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadPartCopyPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["X-Amz-Copy-Source"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-Range"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"CopyPartResult": {"ETag": "ETag", "LastModified": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UploadPartCopy successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the UploadPartCopyOptions model
				uploadPartCopyOptionsModel := new(ibmcloudobjectstorages3apiv2.UploadPartCopyOptions)
				uploadPartCopyOptionsModel.Bucket = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySource = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Key = core.StringPtr("testString")
				uploadPartCopyOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				uploadPartCopyOptionsModel.UploadID = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceRange = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.UploadPartCopyWithContext(ctx, uploadPartCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.UploadPartCopy(uploadPartCopyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.UploadPartCopyWithContext(ctx, uploadPartCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadPartCopyPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["X-Amz-Copy-Source"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Modified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-None-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-If-Unmodified-Since"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime())))
					Expect(req.Header["X-Amz-Copy-Source-Range"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Range"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["partNumber"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["uploadId"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"CopyPartResult": {"ETag": "ETag", "LastModified": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UploadPartCopy successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.UploadPartCopy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadPartCopyOptions model
				uploadPartCopyOptionsModel := new(ibmcloudobjectstorages3apiv2.UploadPartCopyOptions)
				uploadPartCopyOptionsModel.Bucket = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySource = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Key = core.StringPtr("testString")
				uploadPartCopyOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				uploadPartCopyOptionsModel.UploadID = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceRange = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.UploadPartCopy(uploadPartCopyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UploadPartCopy with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the UploadPartCopyOptions model
				uploadPartCopyOptionsModel := new(ibmcloudobjectstorages3apiv2.UploadPartCopyOptions)
				uploadPartCopyOptionsModel.Bucket = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySource = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Key = core.StringPtr("testString")
				uploadPartCopyOptionsModel.PartNumber = core.Int64Ptr(int64(38))
				uploadPartCopyOptionsModel.UploadID = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfModifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceIfNoneMatch = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceIfUnmodifiedSince = CreateMockDateTime()
				uploadPartCopyOptionsModel.XAmzCopySourceRange = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey = core.StringPtr("testString")
				uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5 = core.StringPtr("testString")
				uploadPartCopyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.UploadPartCopy(uploadPartCopyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UploadPartCopyOptions model with no property values
				uploadPartCopyOptionsModelNew := new(ibmcloudobjectstorages3apiv2.UploadPartCopyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.UploadPartCopy(uploadPartCopyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`PutPublicAccessBlock(putPublicAccessBlockOptions *PutPublicAccessBlockOptions)`, func() {
		putPublicAccessBlockPath := "/testString?publicAccessBlock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putPublicAccessBlockPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for publicAccessBlock query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutPublicAccessBlock successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutPublicAccessBlock(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutPublicAccessBlockOptions model
				putPublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.PutPublicAccessBlockOptions)
				putPublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				putPublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				putPublicAccessBlockOptionsModel.Body = core.StringPtr("testString")
				putPublicAccessBlockOptionsModel.ContentMD5 = core.StringPtr("testString")
				putPublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutPublicAccessBlock(putPublicAccessBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutPublicAccessBlock with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutPublicAccessBlockOptions model
				putPublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.PutPublicAccessBlockOptions)
				putPublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				putPublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				putPublicAccessBlockOptionsModel.Body = core.StringPtr("testString")
				putPublicAccessBlockOptionsModel.ContentMD5 = core.StringPtr("testString")
				putPublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutPublicAccessBlock(putPublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutPublicAccessBlockOptions model with no property values
				putPublicAccessBlockOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutPublicAccessBlockOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutPublicAccessBlock(putPublicAccessBlockOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPublicAccessBlock(getPublicAccessBlockOptions *GetPublicAccessBlockOptions) - Operation response error`, func() {
		getPublicAccessBlockPath := "/testString?publicAccessBlock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPublicAccessBlockPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for publicAccessBlock query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPublicAccessBlock with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetPublicAccessBlockOptions model
				getPublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.GetPublicAccessBlockOptions)
				getPublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				getPublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				getPublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(getPublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(getPublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPublicAccessBlock(getPublicAccessBlockOptions *GetPublicAccessBlockOptions)`, func() {
		getPublicAccessBlockPath := "/testString?publicAccessBlock"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPublicAccessBlockPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for publicAccessBlock query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"PublicAccessBlockConfiguration": {"BlockPublicAcls": false, "IgnorePublicAcls": true, "BlockPublicPolicy": false, "RestrictPublicBuckets": false}}`)
				}))
			})
			It(`Invoke GetPublicAccessBlock successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetPublicAccessBlockOptions model
				getPublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.GetPublicAccessBlockOptions)
				getPublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				getPublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				getPublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetPublicAccessBlockWithContext(ctx, getPublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(getPublicAccessBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetPublicAccessBlockWithContext(ctx, getPublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPublicAccessBlockPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for publicAccessBlock query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"PublicAccessBlockConfiguration": {"BlockPublicAcls": false, "IgnorePublicAcls": true, "BlockPublicPolicy": false, "RestrictPublicBuckets": false}}`)
				}))
			})
			It(`Invoke GetPublicAccessBlock successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPublicAccessBlockOptions model
				getPublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.GetPublicAccessBlockOptions)
				getPublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				getPublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				getPublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(getPublicAccessBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPublicAccessBlock with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetPublicAccessBlockOptions model
				getPublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.GetPublicAccessBlockOptions)
				getPublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				getPublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				getPublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(getPublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPublicAccessBlockOptions model with no property values
				getPublicAccessBlockOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetPublicAccessBlockOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetPublicAccessBlock(getPublicAccessBlockOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeletePublicAccessBlock(deletePublicAccessBlockOptions *DeletePublicAccessBlockOptions)`, func() {
		deletePublicAccessBlockPath := "/testString?publicAccessBlock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePublicAccessBlockPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for publicAccessBlock query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePublicAccessBlock successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeletePublicAccessBlock(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePublicAccessBlockOptions model
				deletePublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.DeletePublicAccessBlockOptions)
				deletePublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				deletePublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				deletePublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeletePublicAccessBlock(deletePublicAccessBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePublicAccessBlock with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeletePublicAccessBlockOptions model
				deletePublicAccessBlockOptionsModel := new(ibmcloudobjectstorages3apiv2.DeletePublicAccessBlockOptions)
				deletePublicAccessBlockOptionsModel.Bucket = core.StringPtr("testString")
				deletePublicAccessBlockOptionsModel.PublicAccessBlock = core.BoolPtr(true)
				deletePublicAccessBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeletePublicAccessBlock(deletePublicAccessBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePublicAccessBlockOptions model with no property values
				deletePublicAccessBlockOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeletePublicAccessBlockOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeletePublicAccessBlock(deletePublicAccessBlockOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBucketAcl(getBucketAclOptions *GetBucketAclOptions)`, func() {
		getBucketAclPath := "/testString?acl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketAclPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for acl query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBucketAcl successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetBucketAclOptions model
				getBucketAclOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketAclOptions)
				getBucketAclOptionsModel.Bucket = core.StringPtr("testString")
				getBucketAclOptionsModel.Acl = core.BoolPtr(true)
				getBucketAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketAcl(getBucketAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetBucketAcl with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketAclOptions model
				getBucketAclOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketAclOptions)
				getBucketAclOptionsModel.Bucket = core.StringPtr("testString")
				getBucketAclOptionsModel.Acl = core.BoolPtr(true)
				getBucketAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketAcl(getBucketAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetBucketAclOptions model with no property values
				getBucketAclOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetBucketAclOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketAcl(getBucketAclOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutBucketAcl(putBucketAclOptions *PutBucketAclOptions)`, func() {
		putBucketAclPath := "/testString?acl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putBucketAclPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					// TODO: Add check for acl query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutBucketAcl successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutBucketAclOptions model
				putBucketAclOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketAclOptions)
				putBucketAclOptionsModel.Bucket = core.StringPtr("testString")
				putBucketAclOptionsModel.Acl = core.BoolPtr(true)
				putBucketAclOptionsModel.XAmzAcl = core.StringPtr("private")
				putBucketAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketAcl(putBucketAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutBucketAcl with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutBucketAclOptions model
				putBucketAclOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketAclOptions)
				putBucketAclOptionsModel.Bucket = core.StringPtr("testString")
				putBucketAclOptionsModel.Acl = core.BoolPtr(true)
				putBucketAclOptionsModel.XAmzAcl = core.StringPtr("private")
				putBucketAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketAcl(putBucketAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutBucketAclOptions model with no property values
				putBucketAclOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutBucketAclOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketAcl(putBucketAclOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetObjectAcl(getObjectAclOptions *GetObjectAclOptions) - Operation response error`, func() {
		getObjectAclPath := "/testString/testString?acl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectAclPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for acl query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetObjectAcl with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetObjectAclOptions model
				getObjectAclOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectAclOptions)
				getObjectAclOptionsModel.Bucket = core.StringPtr("testString")
				getObjectAclOptionsModel.Key = core.StringPtr("testString")
				getObjectAclOptionsModel.Acl = core.BoolPtr(true)
				getObjectAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectAcl(getObjectAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectAcl(getObjectAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetObjectAcl(getObjectAclOptions *GetObjectAclOptions)`, func() {
		getObjectAclPath := "/testString/testString?acl"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectAclPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for acl query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Owner": {"DisplayName": "DisplayName", "ID": "ID"}, "Grants": {}}`)
				}))
			})
			It(`Invoke GetObjectAcl successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetObjectAclOptions model
				getObjectAclOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectAclOptions)
				getObjectAclOptionsModel.Bucket = core.StringPtr("testString")
				getObjectAclOptionsModel.Key = core.StringPtr("testString")
				getObjectAclOptionsModel.Acl = core.BoolPtr(true)
				getObjectAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectAclWithContext(ctx, getObjectAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectAcl(getObjectAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectAclWithContext(ctx, getObjectAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectAclPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for acl query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"Owner": {"DisplayName": "DisplayName", "ID": "ID"}, "Grants": {}}`)
				}))
			})
			It(`Invoke GetObjectAcl successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetObjectAclOptions model
				getObjectAclOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectAclOptions)
				getObjectAclOptionsModel.Bucket = core.StringPtr("testString")
				getObjectAclOptionsModel.Key = core.StringPtr("testString")
				getObjectAclOptionsModel.Acl = core.BoolPtr(true)
				getObjectAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectAcl(getObjectAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetObjectAcl with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetObjectAclOptions model
				getObjectAclOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectAclOptions)
				getObjectAclOptionsModel.Bucket = core.StringPtr("testString")
				getObjectAclOptionsModel.Key = core.StringPtr("testString")
				getObjectAclOptionsModel.Acl = core.BoolPtr(true)
				getObjectAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectAcl(getObjectAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetObjectAclOptions model with no property values
				getObjectAclOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetObjectAclOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectAcl(getObjectAclOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutObjectAcl(putObjectAclOptions *PutObjectAclOptions)`, func() {
		putObjectAclPath := "/testString/testString?acl"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putObjectAclPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Amz-Acl"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Acl"][0]).To(Equal(fmt.Sprintf("%v", "private")))
					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Grant-Full-Control"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Grant-Full-Control"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Grant-Read"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Grant-Read"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Grant-Read-Acp"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Grant-Read-Acp"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Grant-Write"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Grant-Write"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Amz-Grant-Write-Acp"]).ToNot(BeNil())
					Expect(req.Header["X-Amz-Grant-Write-Acp"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for acl query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutObjectAcl successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutObjectAcl(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutObjectAclOptions model
				putObjectAclOptionsModel := new(ibmcloudobjectstorages3apiv2.PutObjectAclOptions)
				putObjectAclOptionsModel.Bucket = core.StringPtr("testString")
				putObjectAclOptionsModel.Key = core.StringPtr("testString")
				putObjectAclOptionsModel.Acl = core.BoolPtr(true)
				putObjectAclOptionsModel.Body = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzAcl = core.StringPtr("private")
				putObjectAclOptionsModel.ContentMD5 = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantFullControl = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantRead = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantReadAcp = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantWrite = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantWriteAcp = core.StringPtr("testString")
				putObjectAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutObjectAcl(putObjectAclOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutObjectAcl with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutObjectAclOptions model
				putObjectAclOptionsModel := new(ibmcloudobjectstorages3apiv2.PutObjectAclOptions)
				putObjectAclOptionsModel.Bucket = core.StringPtr("testString")
				putObjectAclOptionsModel.Key = core.StringPtr("testString")
				putObjectAclOptionsModel.Acl = core.BoolPtr(true)
				putObjectAclOptionsModel.Body = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzAcl = core.StringPtr("private")
				putObjectAclOptionsModel.ContentMD5 = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantFullControl = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantRead = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantReadAcp = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantWrite = core.StringPtr("testString")
				putObjectAclOptionsModel.XAmzGrantWriteAcp = core.StringPtr("testString")
				putObjectAclOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutObjectAcl(putObjectAclOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutObjectAclOptions model with no property values
				putObjectAclOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutObjectAclOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutObjectAcl(putObjectAclOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`PutObjectTagging(putObjectTaggingOptions *PutObjectTaggingOptions)`, func() {
		putObjectTaggingPath := "/testString/testString?tagging"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putObjectTaggingPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for tagging query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutObjectTagging successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutObjectTagging(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutObjectTaggingOptions model
				putObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.PutObjectTaggingOptions)
				putObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				putObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				putObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				putObjectTaggingOptionsModel.Body = core.StringPtr("testString")
				putObjectTaggingOptionsModel.ContentMD5 = core.StringPtr("testString")
				putObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutObjectTagging(putObjectTaggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutObjectTagging with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutObjectTaggingOptions model
				putObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.PutObjectTaggingOptions)
				putObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				putObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				putObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				putObjectTaggingOptionsModel.Body = core.StringPtr("testString")
				putObjectTaggingOptionsModel.ContentMD5 = core.StringPtr("testString")
				putObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutObjectTagging(putObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutObjectTaggingOptions model with no property values
				putObjectTaggingOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutObjectTaggingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutObjectTagging(putObjectTaggingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetObjectTagging(getObjectTaggingOptions *GetObjectTaggingOptions) - Operation response error`, func() {
		getObjectTaggingPath := "/testString/testString?tagging"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectTaggingPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for tagging query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetObjectTagging with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetObjectTaggingOptions model
				getObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectTaggingOptions)
				getObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				getObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectTagging(getObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectTagging(getObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetObjectTagging(getObjectTaggingOptions *GetObjectTaggingOptions)`, func() {
		getObjectTaggingPath := "/testString/testString?tagging"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectTaggingPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for tagging query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"TagSet": {}}`)
				}))
			})
			It(`Invoke GetObjectTagging successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetObjectTaggingOptions model
				getObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectTaggingOptions)
				getObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				getObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectTaggingWithContext(ctx, getObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectTagging(getObjectTaggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectTaggingWithContext(ctx, getObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getObjectTaggingPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for tagging query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"TagSet": {}}`)
				}))
			})
			It(`Invoke GetObjectTagging successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectTagging(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetObjectTaggingOptions model
				getObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectTaggingOptions)
				getObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				getObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectTagging(getObjectTaggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetObjectTagging with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetObjectTaggingOptions model
				getObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.GetObjectTaggingOptions)
				getObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				getObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				getObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetObjectTagging(getObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetObjectTaggingOptions model with no property values
				getObjectTaggingOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetObjectTaggingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetObjectTagging(getObjectTaggingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteObjectTagging(deleteObjectTaggingOptions *DeleteObjectTaggingOptions)`, func() {
		deleteObjectTaggingPath := "/testString/testString?tagging"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteObjectTaggingPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for tagging query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteObjectTagging successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjectTagging(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteObjectTaggingOptions model
				deleteObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectTaggingOptions)
				deleteObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				deleteObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				deleteObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObjectTagging(deleteObjectTaggingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteObjectTagging with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteObjectTaggingOptions model
				deleteObjectTaggingOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteObjectTaggingOptions)
				deleteObjectTaggingOptionsModel.Bucket = core.StringPtr("testString")
				deleteObjectTaggingOptionsModel.Key = core.StringPtr("testString")
				deleteObjectTaggingOptionsModel.Tagging = core.BoolPtr(true)
				deleteObjectTaggingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteObjectTagging(deleteObjectTaggingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteObjectTaggingOptions model with no property values
				deleteObjectTaggingOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteObjectTaggingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteObjectTagging(deleteObjectTaggingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "https://ibmcloudobjectstorages3apiv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				})
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudObjectStorageS3ApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudObjectStorageS3ApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudObjectStorageS3ApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudObjectStorageS3ApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_URL": "https://ibmcloudobjectstorages3apiv2/api",
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_OBJECT_STORAGE_S3_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2UsingExternalConfig(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudObjectStorageS3ApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmcloudobjectstorages3apiv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})

	Describe(`PutBucketWebsite(putBucketWebsiteOptions *PutBucketWebsiteOptions)`, func() {
		putBucketWebsitePath := "/testString?website"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putBucketWebsitePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for website query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutBucketWebsite successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketWebsite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutBucketWebsiteOptions model
				putBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketWebsiteOptions)
				putBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				putBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				putBucketWebsiteOptionsModel.Body = core.StringPtr("testString")
				putBucketWebsiteOptionsModel.ContentMD5 = core.StringPtr("testString")
				putBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketWebsite(putBucketWebsiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutBucketWebsite with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutBucketWebsiteOptions model
				putBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketWebsiteOptions)
				putBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				putBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				putBucketWebsiteOptionsModel.Body = core.StringPtr("testString")
				putBucketWebsiteOptionsModel.ContentMD5 = core.StringPtr("testString")
				putBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketWebsite(putBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutBucketWebsiteOptions model with no property values
				putBucketWebsiteOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutBucketWebsiteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketWebsite(putBucketWebsiteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketWebsite(getBucketWebsiteOptions *GetBucketWebsiteOptions) - Operation response error`, func() {
		getBucketWebsitePath := "/testString?website"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketWebsitePath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for website query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketWebsite with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketWebsiteOptions model
				getBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketWebsiteOptions)
				getBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				getBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				getBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketWebsite(getBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketWebsite(getBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBucketWebsite(getBucketWebsiteOptions *GetBucketWebsiteOptions)`, func() {
		getBucketWebsitePath := "/testString?website"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketWebsitePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for website query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"RedirectAllRequestsTo": {"HostName": "HostName", "Protocol": "http"}, "IndexDocument": {"Suffix": "Suffix"}, "ErrorDocument": {"Key": "Key"}, "RoutingRules": {}}`)
				}))
			})
			It(`Invoke GetBucketWebsite successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketWebsiteOptions model
				getBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketWebsiteOptions)
				getBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				getBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				getBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketWebsiteWithContext(ctx, getBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketWebsite(getBucketWebsiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketWebsiteWithContext(ctx, getBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketWebsitePath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for website query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"RedirectAllRequestsTo": {"HostName": "HostName", "Protocol": "http"}, "IndexDocument": {"Suffix": "Suffix"}, "ErrorDocument": {"Key": "Key"}, "RoutingRules": {}}`)
				}))
			})
			It(`Invoke GetBucketWebsite successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketWebsite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketWebsiteOptions model
				getBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketWebsiteOptions)
				getBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				getBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				getBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketWebsite(getBucketWebsiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketWebsite with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketWebsiteOptions model
				getBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketWebsiteOptions)
				getBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				getBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				getBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketWebsite(getBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketWebsiteOptions model with no property values
				getBucketWebsiteOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetBucketWebsiteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketWebsite(getBucketWebsiteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteBucketWebsite(deleteBucketWebsiteOptions *DeleteBucketWebsiteOptions)`, func() {
		deleteBucketWebsitePath := "/testString?website"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBucketWebsitePath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for website query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBucketWebsite successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucketWebsite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBucketWebsiteOptions model
				deleteBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketWebsiteOptions)
				deleteBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				deleteBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucketWebsite(deleteBucketWebsiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBucketWebsite with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteBucketWebsiteOptions model
				deleteBucketWebsiteOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketWebsiteOptions)
				deleteBucketWebsiteOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketWebsiteOptionsModel.Website = core.BoolPtr(true)
				deleteBucketWebsiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucketWebsite(deleteBucketWebsiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBucketWebsiteOptions model with no property values
				deleteBucketWebsiteOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteBucketWebsiteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucketWebsite(deleteBucketWebsiteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PutBucketCors(putBucketCorsOptions *PutBucketCorsOptions)`, func() {
		putBucketCorsPath := "/testString?cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(putBucketCorsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Content-Md5"]).ToNot(BeNil())
					Expect(req.Header["Content-Md5"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for cors query parameter
					res.WriteHeader(200)
				}))
			})
			It(`Invoke PutBucketCors successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketCors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PutBucketCorsOptions model
				putBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketCorsOptions)
				putBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				putBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				putBucketCorsOptionsModel.Body = core.StringPtr("testString")
				putBucketCorsOptionsModel.ContentMD5 = core.StringPtr("testString")
				putBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketCors(putBucketCorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke PutBucketCors with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the PutBucketCorsOptions model
				putBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.PutBucketCorsOptions)
				putBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				putBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				putBucketCorsOptionsModel.Body = core.StringPtr("testString")
				putBucketCorsOptionsModel.ContentMD5 = core.StringPtr("testString")
				putBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.PutBucketCors(putBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the PutBucketCorsOptions model with no property values
				putBucketCorsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.PutBucketCorsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.PutBucketCors(putBucketCorsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBucketCors(getBucketCorsOptions *GetBucketCorsOptions) - Operation response error`, func() {
		getBucketCorsPath := "/testString?cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketCorsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for cors query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBucketCors with error: Operation response processing error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketCorsOptions model
				getBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketCorsOptions)
				getBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				getBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				getBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketCors(getBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketCors(getBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBucketCors(getBucketCorsOptions *GetBucketCorsOptions)`, func() {
		getBucketCorsPath := "/testString?cors"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketCorsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for cors query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"CORSRules": {}}`)
				}))
			})
			It(`Invoke GetBucketCors successfully with retries`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())
				ibmCloudObjectStorageS3ApiService.EnableRetries(0, 0)

				// Construct an instance of the GetBucketCorsOptions model
				getBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketCorsOptions)
				getBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				getBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				getBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketCorsWithContext(ctx, getBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudObjectStorageS3ApiService.DisableRetries()
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketCors(getBucketCorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketCorsWithContext(ctx, getBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBucketCorsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for cors query parameter
					// Set mock response
					res.Header().Set("Content-type", "text/xml")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"CORSRules": {}}`)
				}))
			})
			It(`Invoke GetBucketCors successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketCors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBucketCorsOptions model
				getBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketCorsOptions)
				getBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				getBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				getBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketCors(getBucketCorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBucketCors with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the GetBucketCorsOptions model
				getBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.GetBucketCorsOptions)
				getBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				getBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				getBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudObjectStorageS3ApiService.GetBucketCors(getBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBucketCorsOptions model with no property values
				getBucketCorsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.GetBucketCorsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudObjectStorageS3ApiService.GetBucketCors(getBucketCorsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteBucketCors(deleteBucketCorsOptions *DeleteBucketCorsOptions)`, func() {
		deleteBucketCorsPath := "/testString?cors"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteBucketCorsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for cors query parameter
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteBucketCors successfully`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucketCors(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteBucketCorsOptions model
				deleteBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketCorsOptions)
				deleteBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				deleteBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucketCors(deleteBucketCorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteBucketCors with error: Operation validation and request error`, func() {
				ibmCloudObjectStorageS3ApiService, serviceErr := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudObjectStorageS3ApiService).ToNot(BeNil())

				// Construct an instance of the DeleteBucketCorsOptions model
				deleteBucketCorsOptionsModel := new(ibmcloudobjectstorages3apiv2.DeleteBucketCorsOptions)
				deleteBucketCorsOptionsModel.Bucket = core.StringPtr("testString")
				deleteBucketCorsOptionsModel.Cors = core.BoolPtr(true)
				deleteBucketCorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudObjectStorageS3ApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := ibmCloudObjectStorageS3ApiService.DeleteBucketCors(deleteBucketCorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteBucketCorsOptions model with no property values
				deleteBucketCorsOptionsModelNew := new(ibmcloudobjectstorages3apiv2.DeleteBucketCorsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = ibmCloudObjectStorageS3ApiService.DeleteBucketCors(deleteBucketCorsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmCloudObjectStorageS3ApiService, _ := ibmcloudobjectstorages3apiv2.NewIbmCloudObjectStorageS3ApiV2(&ibmcloudobjectstorages3apiv2.IbmCloudObjectStorageS3ApiV2Options{
				URL:           "http://ibmcloudobjectstorages3apiv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAbortMultipartUploadOptions successfully`, func() {
				// Construct an instance of the AbortMultipartUploadOptions model
				bucket := "testString"
				key := "testString"
				uploadID := "testString"
				abortMultipartUploadOptionsModel := ibmCloudObjectStorageS3ApiService.NewAbortMultipartUploadOptions(bucket, key, uploadID)
				abortMultipartUploadOptionsModel.SetBucket("testString")
				abortMultipartUploadOptionsModel.SetKey("testString")
				abortMultipartUploadOptionsModel.SetUploadID("testString")
				abortMultipartUploadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(abortMultipartUploadOptionsModel).ToNot(BeNil())
				Expect(abortMultipartUploadOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(abortMultipartUploadOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(abortMultipartUploadOptionsModel.UploadID).To(Equal(core.StringPtr("testString")))
				Expect(abortMultipartUploadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCompleteMultipartUploadOptions successfully`, func() {
				// Construct an instance of the CompleteMultipartUploadOptions model
				bucket := "testString"
				key := "testString"
				uploadID := "testString"
				body := "testString"
				completeMultipartUploadOptionsModel := ibmCloudObjectStorageS3ApiService.NewCompleteMultipartUploadOptions(bucket, key, uploadID, body)
				completeMultipartUploadOptionsModel.SetBucket("testString")
				completeMultipartUploadOptionsModel.SetKey("testString")
				completeMultipartUploadOptionsModel.SetUploadID("testString")
				completeMultipartUploadOptionsModel.SetBody("testString")
				completeMultipartUploadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeMultipartUploadOptionsModel).ToNot(BeNil())
				Expect(completeMultipartUploadOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(completeMultipartUploadOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(completeMultipartUploadOptionsModel.UploadID).To(Equal(core.StringPtr("testString")))
				Expect(completeMultipartUploadOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(completeMultipartUploadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCopyObjectOptions successfully`, func() {
				// Construct an instance of the CopyObjectOptions model
				bucket := "testString"
				xAmzCopySource := "testString"
				targetKey := "testString"
				body := "testString"
				copyObjectOptionsModel := ibmCloudObjectStorageS3ApiService.NewCopyObjectOptions(bucket, xAmzCopySource, targetKey, body)
				copyObjectOptionsModel.SetBucket("testString")
				copyObjectOptionsModel.SetXAmzCopySource("testString")
				copyObjectOptionsModel.SetTargetKey("testString")
				copyObjectOptionsModel.SetBody("testString")
				copyObjectOptionsModel.SetXAmzAcl("private")
				copyObjectOptionsModel.SetCacheControl("testString")
				copyObjectOptionsModel.SetContentDisposition("testString")
				copyObjectOptionsModel.SetContentEncoding("testString")
				copyObjectOptionsModel.SetContentLanguage("testString")
				copyObjectOptionsModel.SetXAmzCopySourceIfMatch("testString")
				copyObjectOptionsModel.SetXAmzCopySourceIfModifiedSince(CreateMockDateTime())
				copyObjectOptionsModel.SetXAmzCopySourceIfNoneMatch("testString")
				copyObjectOptionsModel.SetXAmzCopySourceIfUnmodifiedSince(CreateMockDateTime())
				copyObjectOptionsModel.SetExpires(CreateMockDateTime())
				copyObjectOptionsModel.SetXAmzMetadataDirective("COPY")
				copyObjectOptionsModel.SetXAmzTaggingDirective("COPY")
				copyObjectOptionsModel.SetXAmzServerSideEncryption("AES256")
				copyObjectOptionsModel.SetXAmzWebsiteRedirectLocation("testString")
				copyObjectOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				copyObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				copyObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				copyObjectOptionsModel.SetXAmzCopySourceServerSideEncryptionCustomerAlgorithm("testString")
				copyObjectOptionsModel.SetXAmzCopySourceServerSideEncryptionCustomerKey("testString")
				copyObjectOptionsModel.SetXAmzCopySourceServerSideEncryptionCustomerKeyMD5("testString")
				copyObjectOptionsModel.SetXAmzTagging("testString")
				copyObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(copyObjectOptionsModel).ToNot(BeNil())
				Expect(copyObjectOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySource).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.TargetKey).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzAcl).To(Equal(core.StringPtr("private")))
				Expect(copyObjectOptionsModel.CacheControl).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.ContentDisposition).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.ContentEncoding).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.ContentLanguage).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySourceIfMatch).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySourceIfModifiedSince).To(Equal(CreateMockDateTime()))
				Expect(copyObjectOptionsModel.XAmzCopySourceIfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySourceIfUnmodifiedSince).To(Equal(CreateMockDateTime()))
				Expect(copyObjectOptionsModel.Expires).To(Equal(CreateMockDateTime()))
				Expect(copyObjectOptionsModel.XAmzMetadataDirective).To(Equal(core.StringPtr("COPY")))
				Expect(copyObjectOptionsModel.XAmzTaggingDirective).To(Equal(core.StringPtr("COPY")))
				Expect(copyObjectOptionsModel.XAmzServerSideEncryption).To(Equal(core.StringPtr("AES256")))
				Expect(copyObjectOptionsModel.XAmzWebsiteRedirectLocation).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.XAmzTagging).To(Equal(core.StringPtr("testString")))
				Expect(copyObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateBucketOptions successfully`, func() {
				// Construct an instance of the CreateBucketOptions model
				bucket := "my-new-bucket"
				ibmServiceInstanceID := "d6f76k03-6k4f-4a82-n165-697654o63903"
				createBucketOptionsModel := ibmCloudObjectStorageS3ApiService.NewCreateBucketOptions(bucket, ibmServiceInstanceID)
				createBucketOptionsModel.SetBucket("my-new-bucket")
				createBucketOptionsModel.SetIbmServiceInstanceID("d6f76k03-6k4f-4a82-n165-697654o63903")
				createBucketOptionsModel.SetBody("testString")
				createBucketOptionsModel.SetIbmSseKpEncryptionAlgorithm("AES256")
				createBucketOptionsModel.SetIbmSseKpCustomerRootKeyCrn("crn:v1:bluemix:public:kms:us-south:a/f047b55a3362ac06afad8a3f2f5586ea:12e8c9c2-a162-472d-b7d6-8b9a86b815a6:key:02fd6835-6001-4482-a892-13bd2085f75d")
				createBucketOptionsModel.SetXAmzAcl("private")
				createBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createBucketOptionsModel).ToNot(BeNil())
				Expect(createBucketOptionsModel.Bucket).To(Equal(core.StringPtr("my-new-bucket")))
				Expect(createBucketOptionsModel.IbmServiceInstanceID).To(Equal(core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")))
				Expect(createBucketOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(createBucketOptionsModel.IbmSseKpEncryptionAlgorithm).To(Equal(core.StringPtr("AES256")))
				Expect(createBucketOptionsModel.IbmSseKpCustomerRootKeyCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/f047b55a3362ac06afad8a3f2f5586ea:12e8c9c2-a162-472d-b7d6-8b9a86b815a6:key:02fd6835-6001-4482-a892-13bd2085f75d")))
				Expect(createBucketOptionsModel.XAmzAcl).To(Equal(core.StringPtr("private")))
				Expect(createBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBucketCorsOptions successfully`, func() {
				// Construct an instance of the DeleteBucketCorsOptions model
				bucket := "testString"
				cors := true
				deleteBucketCorsOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteBucketCorsOptions(bucket, cors)
				deleteBucketCorsOptionsModel.SetBucket("testString")
				deleteBucketCorsOptionsModel.SetCors(true)
				deleteBucketCorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBucketCorsOptionsModel).ToNot(BeNil())
				Expect(deleteBucketCorsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketCorsOptionsModel.Cors).To(Equal(core.BoolPtr(true)))
				Expect(deleteBucketCorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBucketLifecycleOptions successfully`, func() {
				// Construct an instance of the DeleteBucketLifecycleOptions model
				bucket := "testString"
				lifecycle := true
				deleteBucketLifecycleOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteBucketLifecycleOptions(bucket, lifecycle)
				deleteBucketLifecycleOptionsModel.SetBucket("testString")
				deleteBucketLifecycleOptionsModel.SetLifecycle(true)
				deleteBucketLifecycleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBucketLifecycleOptionsModel).ToNot(BeNil())
				Expect(deleteBucketLifecycleOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketLifecycleOptionsModel.Lifecycle).To(Equal(core.BoolPtr(true)))
				Expect(deleteBucketLifecycleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBucketOptions successfully`, func() {
				// Construct an instance of the DeleteBucketOptions model
				bucket := "testString"
				deleteBucketOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteBucketOptions(bucket)
				deleteBucketOptionsModel.SetBucket("testString")
				deleteBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBucketOptionsModel).ToNot(BeNil())
				Expect(deleteBucketOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteBucketWebsiteOptions successfully`, func() {
				// Construct an instance of the DeleteBucketWebsiteOptions model
				bucket := "testString"
				website := true
				deleteBucketWebsiteOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteBucketWebsiteOptions(bucket, website)
				deleteBucketWebsiteOptionsModel.SetBucket("testString")
				deleteBucketWebsiteOptionsModel.SetWebsite(true)
				deleteBucketWebsiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteBucketWebsiteOptionsModel).ToNot(BeNil())
				Expect(deleteBucketWebsiteOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteBucketWebsiteOptionsModel.Website).To(Equal(core.BoolPtr(true)))
				Expect(deleteBucketWebsiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteObjectOptions successfully`, func() {
				// Construct an instance of the DeleteObjectOptions model
				bucket := "testString"
				key := "testString"
				deleteObjectOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteObjectOptions(bucket, key)
				deleteObjectOptionsModel.SetBucket("testString")
				deleteObjectOptionsModel.SetKey("testString")
				deleteObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteObjectOptionsModel).ToNot(BeNil())
				Expect(deleteObjectOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteObjectTaggingOptions successfully`, func() {
				// Construct an instance of the DeleteObjectTaggingOptions model
				bucket := "testString"
				key := "testString"
				tagging := true
				deleteObjectTaggingOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteObjectTaggingOptions(bucket, key, tagging)
				deleteObjectTaggingOptionsModel.SetBucket("testString")
				deleteObjectTaggingOptionsModel.SetKey("testString")
				deleteObjectTaggingOptionsModel.SetTagging(true)
				deleteObjectTaggingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteObjectTaggingOptionsModel).ToNot(BeNil())
				Expect(deleteObjectTaggingOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectTaggingOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectTaggingOptionsModel.Tagging).To(Equal(core.BoolPtr(true)))
				Expect(deleteObjectTaggingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteObjectsOptions successfully`, func() {
				// Construct an instance of the DeleteObjectsOptions model
				bucket := "testString"
				deleteVar := true
				body := "testString"
				deleteObjectsOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeleteObjectsOptions(bucket, deleteVar, body)
				deleteObjectsOptionsModel.SetBucket("testString")
				deleteObjectsOptionsModel.SetDelete(true)
				deleteObjectsOptionsModel.SetBody("testString")
				deleteObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteObjectsOptionsModel).ToNot(BeNil())
				Expect(deleteObjectsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectsOptionsModel.Delete).To(Equal(core.BoolPtr(true)))
				Expect(deleteObjectsOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(deleteObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePublicAccessBlockOptions successfully`, func() {
				// Construct an instance of the DeletePublicAccessBlockOptions model
				bucket := "testString"
				publicAccessBlock := true
				deletePublicAccessBlockOptionsModel := ibmCloudObjectStorageS3ApiService.NewDeletePublicAccessBlockOptions(bucket, publicAccessBlock)
				deletePublicAccessBlockOptionsModel.SetBucket("testString")
				deletePublicAccessBlockOptionsModel.SetPublicAccessBlock(true)
				deletePublicAccessBlockOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePublicAccessBlockOptionsModel).ToNot(BeNil())
				Expect(deletePublicAccessBlockOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(deletePublicAccessBlockOptionsModel.PublicAccessBlock).To(Equal(core.BoolPtr(true)))
				Expect(deletePublicAccessBlockOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketAclOptions successfully`, func() {
				// Construct an instance of the GetBucketAclOptions model
				bucket := "testString"
				acl := true
				getBucketAclOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetBucketAclOptions(bucket, acl)
				getBucketAclOptionsModel.SetBucket("testString")
				getBucketAclOptionsModel.SetAcl(true)
				getBucketAclOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketAclOptionsModel).ToNot(BeNil())
				Expect(getBucketAclOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getBucketAclOptionsModel.Acl).To(Equal(core.BoolPtr(true)))
				Expect(getBucketAclOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketCorsOptions successfully`, func() {
				// Construct an instance of the GetBucketCorsOptions model
				bucket := "testString"
				cors := true
				getBucketCorsOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetBucketCorsOptions(bucket, cors)
				getBucketCorsOptionsModel.SetBucket("testString")
				getBucketCorsOptionsModel.SetCors(true)
				getBucketCorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketCorsOptionsModel).ToNot(BeNil())
				Expect(getBucketCorsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getBucketCorsOptionsModel.Cors).To(Equal(core.BoolPtr(true)))
				Expect(getBucketCorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketLifecycleConfigurationOptions successfully`, func() {
				// Construct an instance of the GetBucketLifecycleConfigurationOptions model
				bucket := "testString"
				lifecycle := true
				getBucketLifecycleConfigurationOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetBucketLifecycleConfigurationOptions(bucket, lifecycle)
				getBucketLifecycleConfigurationOptionsModel.SetBucket("testString")
				getBucketLifecycleConfigurationOptionsModel.SetLifecycle(true)
				getBucketLifecycleConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketLifecycleConfigurationOptionsModel).ToNot(BeNil())
				Expect(getBucketLifecycleConfigurationOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getBucketLifecycleConfigurationOptionsModel.Lifecycle).To(Equal(core.BoolPtr(true)))
				Expect(getBucketLifecycleConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBucketWebsiteOptions successfully`, func() {
				// Construct an instance of the GetBucketWebsiteOptions model
				bucket := "testString"
				website := true
				getBucketWebsiteOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetBucketWebsiteOptions(bucket, website)
				getBucketWebsiteOptionsModel.SetBucket("testString")
				getBucketWebsiteOptionsModel.SetWebsite(true)
				getBucketWebsiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBucketWebsiteOptionsModel).ToNot(BeNil())
				Expect(getBucketWebsiteOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getBucketWebsiteOptionsModel.Website).To(Equal(core.BoolPtr(true)))
				Expect(getBucketWebsiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetObjectAclOptions successfully`, func() {
				// Construct an instance of the GetObjectAclOptions model
				bucket := "testString"
				key := "testString"
				acl := true
				getObjectAclOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetObjectAclOptions(bucket, key, acl)
				getObjectAclOptionsModel.SetBucket("testString")
				getObjectAclOptionsModel.SetKey("testString")
				getObjectAclOptionsModel.SetAcl(true)
				getObjectAclOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getObjectAclOptionsModel).ToNot(BeNil())
				Expect(getObjectAclOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getObjectAclOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(getObjectAclOptionsModel.Acl).To(Equal(core.BoolPtr(true)))
				Expect(getObjectAclOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetObjectOptions successfully`, func() {
				// Construct an instance of the GetObjectOptions model
				bucket := "testString"
				key := "testString"
				getObjectOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetObjectOptions(bucket, key)
				getObjectOptionsModel.SetBucket("testString")
				getObjectOptionsModel.SetKey("testString")
				getObjectOptionsModel.SetIfMatch("testString")
				getObjectOptionsModel.SetIfModifiedSince(CreateMockDateTime())
				getObjectOptionsModel.SetIfNoneMatch("testString")
				getObjectOptionsModel.SetIfUnmodifiedSince(CreateMockDateTime())
				getObjectOptionsModel.SetRange("testString")
				getObjectOptionsModel.SetResponseCacheControl("testString")
				getObjectOptionsModel.SetResponseContentDisposition("testString")
				getObjectOptionsModel.SetResponseContentEncoding("testString")
				getObjectOptionsModel.SetResponseContentLanguage("testString")
				getObjectOptionsModel.SetResponseContentType("testString")
				getObjectOptionsModel.SetResponseExpires(CreateMockDateTime())
				getObjectOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				getObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				getObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				getObjectOptionsModel.SetPartNumber(int64(38))
				getObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getObjectOptionsModel).ToNot(BeNil())
				Expect(getObjectOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.IfModifiedSince).To(Equal(CreateMockDateTime()))
				Expect(getObjectOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.IfUnmodifiedSince).To(Equal(CreateMockDateTime()))
				Expect(getObjectOptionsModel.Range).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ResponseCacheControl).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ResponseContentDisposition).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ResponseContentEncoding).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ResponseContentLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ResponseContentType).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.ResponseExpires).To(Equal(CreateMockDateTime()))
				Expect(getObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(getObjectOptionsModel.PartNumber).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetObjectTaggingOptions successfully`, func() {
				// Construct an instance of the GetObjectTaggingOptions model
				bucket := "testString"
				key := "testString"
				tagging := true
				getObjectTaggingOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetObjectTaggingOptions(bucket, key, tagging)
				getObjectTaggingOptionsModel.SetBucket("testString")
				getObjectTaggingOptionsModel.SetKey("testString")
				getObjectTaggingOptionsModel.SetTagging(true)
				getObjectTaggingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getObjectTaggingOptionsModel).ToNot(BeNil())
				Expect(getObjectTaggingOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getObjectTaggingOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(getObjectTaggingOptionsModel.Tagging).To(Equal(core.BoolPtr(true)))
				Expect(getObjectTaggingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPublicAccessBlockOptions successfully`, func() {
				// Construct an instance of the GetPublicAccessBlockOptions model
				bucket := "testString"
				publicAccessBlock := true
				getPublicAccessBlockOptionsModel := ibmCloudObjectStorageS3ApiService.NewGetPublicAccessBlockOptions(bucket, publicAccessBlock)
				getPublicAccessBlockOptionsModel.SetBucket("testString")
				getPublicAccessBlockOptionsModel.SetPublicAccessBlock(true)
				getPublicAccessBlockOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPublicAccessBlockOptionsModel).ToNot(BeNil())
				Expect(getPublicAccessBlockOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(getPublicAccessBlockOptionsModel.PublicAccessBlock).To(Equal(core.BoolPtr(true)))
				Expect(getPublicAccessBlockOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadBucketOptions successfully`, func() {
				// Construct an instance of the HeadBucketOptions model
				bucket := "testString"
				headBucketOptionsModel := ibmCloudObjectStorageS3ApiService.NewHeadBucketOptions(bucket)
				headBucketOptionsModel.SetBucket("testString")
				headBucketOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headBucketOptionsModel).ToNot(BeNil())
				Expect(headBucketOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(headBucketOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHeadObjectOptions successfully`, func() {
				// Construct an instance of the HeadObjectOptions model
				bucket := "testString"
				key := "testString"
				headObjectOptionsModel := ibmCloudObjectStorageS3ApiService.NewHeadObjectOptions(bucket, key)
				headObjectOptionsModel.SetBucket("testString")
				headObjectOptionsModel.SetKey("testString")
				headObjectOptionsModel.SetIfMatch("testString")
				headObjectOptionsModel.SetIfModifiedSince(CreateMockDateTime())
				headObjectOptionsModel.SetIfNoneMatch("testString")
				headObjectOptionsModel.SetIfUnmodifiedSince(CreateMockDateTime())
				headObjectOptionsModel.SetRange("testString")
				headObjectOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				headObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				headObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				headObjectOptionsModel.SetPartNumber(int64(38))
				headObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(headObjectOptionsModel).ToNot(BeNil())
				Expect(headObjectOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.IfModifiedSince).To(Equal(CreateMockDateTime()))
				Expect(headObjectOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.IfUnmodifiedSince).To(Equal(CreateMockDateTime()))
				Expect(headObjectOptionsModel.Range).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(headObjectOptionsModel.PartNumber).To(Equal(core.Int64Ptr(int64(38))))
				Expect(headObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInitiateMultipartUploadOptions successfully`, func() {
				// Construct an instance of the InitiateMultipartUploadOptions model
				bucket := "testString"
				uploads := true
				key := "testString"
				body := "testString"
				initiateMultipartUploadOptionsModel := ibmCloudObjectStorageS3ApiService.NewInitiateMultipartUploadOptions(bucket, uploads, key, body)
				initiateMultipartUploadOptionsModel.SetBucket("testString")
				initiateMultipartUploadOptionsModel.SetUploads(true)
				initiateMultipartUploadOptionsModel.SetKey("testString")
				initiateMultipartUploadOptionsModel.SetBody("testString")
				initiateMultipartUploadOptionsModel.SetIfMatch("testString")
				initiateMultipartUploadOptionsModel.SetIfNoneMatch("testString")
				initiateMultipartUploadOptionsModel.SetIfUnmodifiedSince(CreateMockDateTime())
				initiateMultipartUploadOptionsModel.SetCacheControl("testString")
				initiateMultipartUploadOptionsModel.SetContentDisposition("testString")
				initiateMultipartUploadOptionsModel.SetContentEncoding("testString")
				initiateMultipartUploadOptionsModel.SetContentLanguage("testString")
				initiateMultipartUploadOptionsModel.SetExpires(CreateMockDateTime())
				initiateMultipartUploadOptionsModel.SetXAmzServerSideEncryption("AES256")
				initiateMultipartUploadOptionsModel.SetXAmzWebsiteRedirectLocation("testString")
				initiateMultipartUploadOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				initiateMultipartUploadOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				initiateMultipartUploadOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				initiateMultipartUploadOptionsModel.SetXAmzTagging("testString")
				initiateMultipartUploadOptionsModel.SetXAmzAcl("private")
				initiateMultipartUploadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(initiateMultipartUploadOptionsModel).ToNot(BeNil())
				Expect(initiateMultipartUploadOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.Uploads).To(Equal(core.BoolPtr(true)))
				Expect(initiateMultipartUploadOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.IfUnmodifiedSince).To(Equal(CreateMockDateTime()))
				Expect(initiateMultipartUploadOptionsModel.CacheControl).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.ContentDisposition).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.ContentEncoding).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.ContentLanguage).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.Expires).To(Equal(CreateMockDateTime()))
				Expect(initiateMultipartUploadOptionsModel.XAmzServerSideEncryption).To(Equal(core.StringPtr("AES256")))
				Expect(initiateMultipartUploadOptionsModel.XAmzWebsiteRedirectLocation).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.XAmzTagging).To(Equal(core.StringPtr("testString")))
				Expect(initiateMultipartUploadOptionsModel.XAmzAcl).To(Equal(core.StringPtr("private")))
				Expect(initiateMultipartUploadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBucketsOptions successfully`, func() {
				// Construct an instance of the ListBucketsOptions model
				ibmServiceInstanceID := "d6f76k03-6k4f-4a82-n165-697654o63903"
				listBucketsOptionsModel := ibmCloudObjectStorageS3ApiService.NewListBucketsOptions(ibmServiceInstanceID)
				listBucketsOptionsModel.SetIbmServiceInstanceID("d6f76k03-6k4f-4a82-n165-697654o63903")
				listBucketsOptionsModel.SetExtended(true)
				listBucketsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBucketsOptionsModel).ToNot(BeNil())
				Expect(listBucketsOptionsModel.IbmServiceInstanceID).To(Equal(core.StringPtr("d6f76k03-6k4f-4a82-n165-697654o63903")))
				Expect(listBucketsOptionsModel.Extended).To(Equal(core.BoolPtr(true)))
				Expect(listBucketsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListMultipartUploadsOptions successfully`, func() {
				// Construct an instance of the ListMultipartUploadsOptions model
				bucket := "testString"
				uploads := true
				listMultipartUploadsOptionsModel := ibmCloudObjectStorageS3ApiService.NewListMultipartUploadsOptions(bucket, uploads)
				listMultipartUploadsOptionsModel.SetBucket("testString")
				listMultipartUploadsOptionsModel.SetUploads(true)
				listMultipartUploadsOptionsModel.SetDelimiter("testString")
				listMultipartUploadsOptionsModel.SetEncodingType("url")
				listMultipartUploadsOptionsModel.SetKeyMarker("testString")
				listMultipartUploadsOptionsModel.SetMaxUploads(int64(38))
				listMultipartUploadsOptionsModel.SetPrefix("testString")
				listMultipartUploadsOptionsModel.SetUploadIdMarker("testString")
				listMultipartUploadsOptionsModel.SetPaginationLimit("testString")
				listMultipartUploadsOptionsModel.SetPaginationToken("testString")
				listMultipartUploadsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listMultipartUploadsOptionsModel).ToNot(BeNil())
				Expect(listMultipartUploadsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.Uploads).To(Equal(core.BoolPtr(true)))
				Expect(listMultipartUploadsOptionsModel.Delimiter).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.EncodingType).To(Equal(core.StringPtr("url")))
				Expect(listMultipartUploadsOptionsModel.KeyMarker).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.MaxUploads).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listMultipartUploadsOptionsModel.Prefix).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.UploadIdMarker).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.PaginationLimit).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.PaginationToken).To(Equal(core.StringPtr("testString")))
				Expect(listMultipartUploadsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListObjectsOptions successfully`, func() {
				// Construct an instance of the ListObjectsOptions model
				bucket := "testString"
				listObjectsOptionsModel := ibmCloudObjectStorageS3ApiService.NewListObjectsOptions(bucket)
				listObjectsOptionsModel.SetBucket("testString")
				listObjectsOptionsModel.SetDelimiter("testString")
				listObjectsOptionsModel.SetEncodingType("url")
				listObjectsOptionsModel.SetMarker("testString")
				listObjectsOptionsModel.SetMaxKeys(int64(38))
				listObjectsOptionsModel.SetPrefix("testString")
				listObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listObjectsOptionsModel).ToNot(BeNil())
				Expect(listObjectsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.Delimiter).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.EncodingType).To(Equal(core.StringPtr("url")))
				Expect(listObjectsOptionsModel.Marker).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.MaxKeys).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listObjectsOptionsModel.Prefix).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListObjectsV2Options successfully`, func() {
				// Construct an instance of the ListObjectsV2Options model
				bucket := "testString"
				listType := "2"
				listObjectsV2OptionsModel := ibmCloudObjectStorageS3ApiService.NewListObjectsV2Options(bucket, listType)
				listObjectsV2OptionsModel.SetBucket("testString")
				listObjectsV2OptionsModel.SetListType("2")
				listObjectsV2OptionsModel.SetDelimiter("testString")
				listObjectsV2OptionsModel.SetEncodingType("url")
				listObjectsV2OptionsModel.SetMaxKeys(int64(38))
				listObjectsV2OptionsModel.SetPrefix("testString")
				listObjectsV2OptionsModel.SetContinuationToken("testString")
				listObjectsV2OptionsModel.SetFetchOwner(true)
				listObjectsV2OptionsModel.SetStartAfter("testString")
				listObjectsV2OptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listObjectsV2OptionsModel).ToNot(BeNil())
				Expect(listObjectsV2OptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsV2OptionsModel.ListType).To(Equal(core.StringPtr("2")))
				Expect(listObjectsV2OptionsModel.Delimiter).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsV2OptionsModel.EncodingType).To(Equal(core.StringPtr("url")))
				Expect(listObjectsV2OptionsModel.MaxKeys).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listObjectsV2OptionsModel.Prefix).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsV2OptionsModel.ContinuationToken).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsV2OptionsModel.FetchOwner).To(Equal(core.BoolPtr(true)))
				Expect(listObjectsV2OptionsModel.StartAfter).To(Equal(core.StringPtr("testString")))
				Expect(listObjectsV2OptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPartsOptions successfully`, func() {
				// Construct an instance of the ListPartsOptions model
				bucket := "testString"
				key := "testString"
				uploadID := "testString"
				listPartsOptionsModel := ibmCloudObjectStorageS3ApiService.NewListPartsOptions(bucket, key, uploadID)
				listPartsOptionsModel.SetBucket("testString")
				listPartsOptionsModel.SetKey("testString")
				listPartsOptionsModel.SetUploadID("testString")
				listPartsOptionsModel.SetMaxParts(int64(38))
				listPartsOptionsModel.SetPartNumberMarker(int64(38))
				listPartsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPartsOptionsModel).ToNot(BeNil())
				Expect(listPartsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(listPartsOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(listPartsOptionsModel.UploadID).To(Equal(core.StringPtr("testString")))
				Expect(listPartsOptionsModel.MaxParts).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listPartsOptionsModel.PartNumberMarker).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listPartsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutBucketAclOptions successfully`, func() {
				// Construct an instance of the PutBucketAclOptions model
				bucket := "testString"
				acl := true
				putBucketAclOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutBucketAclOptions(bucket, acl)
				putBucketAclOptionsModel.SetBucket("testString")
				putBucketAclOptionsModel.SetAcl(true)
				putBucketAclOptionsModel.SetXAmzAcl("private")
				putBucketAclOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putBucketAclOptionsModel).ToNot(BeNil())
				Expect(putBucketAclOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putBucketAclOptionsModel.Acl).To(Equal(core.BoolPtr(true)))
				Expect(putBucketAclOptionsModel.XAmzAcl).To(Equal(core.StringPtr("private")))
				Expect(putBucketAclOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutBucketCorsOptions successfully`, func() {
				// Construct an instance of the PutBucketCorsOptions model
				bucket := "testString"
				cors := true
				body := "testString"
				putBucketCorsOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutBucketCorsOptions(bucket, cors, body)
				putBucketCorsOptionsModel.SetBucket("testString")
				putBucketCorsOptionsModel.SetCors(true)
				putBucketCorsOptionsModel.SetBody("testString")
				putBucketCorsOptionsModel.SetContentMD5("testString")
				putBucketCorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putBucketCorsOptionsModel).ToNot(BeNil())
				Expect(putBucketCorsOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putBucketCorsOptionsModel.Cors).To(Equal(core.BoolPtr(true)))
				Expect(putBucketCorsOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putBucketCorsOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(putBucketCorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutBucketLifecycleConfigurationOptions successfully`, func() {
				// Construct an instance of the PutBucketLifecycleConfigurationOptions model
				bucket := "testString"
				lifecycle := true
				body := "testString"
				putBucketLifecycleConfigurationOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutBucketLifecycleConfigurationOptions(bucket, lifecycle, body)
				putBucketLifecycleConfigurationOptionsModel.SetBucket("testString")
				putBucketLifecycleConfigurationOptionsModel.SetLifecycle(true)
				putBucketLifecycleConfigurationOptionsModel.SetBody("testString")
				putBucketLifecycleConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putBucketLifecycleConfigurationOptionsModel).ToNot(BeNil())
				Expect(putBucketLifecycleConfigurationOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putBucketLifecycleConfigurationOptionsModel.Lifecycle).To(Equal(core.BoolPtr(true)))
				Expect(putBucketLifecycleConfigurationOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putBucketLifecycleConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutBucketProtectionConfigurationOptions successfully`, func() {
				// Construct an instance of the PutBucketProtectionConfigurationOptions model
				bucket := "testString"
				protection := true
				body := "testString"
				putBucketProtectionConfigurationOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutBucketProtectionConfigurationOptions(bucket, protection, body)
				putBucketProtectionConfigurationOptionsModel.SetBucket("testString")
				putBucketProtectionConfigurationOptionsModel.SetProtection(true)
				putBucketProtectionConfigurationOptionsModel.SetBody("testString")
				putBucketProtectionConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putBucketProtectionConfigurationOptionsModel).ToNot(BeNil())
				Expect(putBucketProtectionConfigurationOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putBucketProtectionConfigurationOptionsModel.Protection).To(Equal(core.BoolPtr(true)))
				Expect(putBucketProtectionConfigurationOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putBucketProtectionConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutBucketWebsiteOptions successfully`, func() {
				// Construct an instance of the PutBucketWebsiteOptions model
				bucket := "testString"
				website := true
				body := "testString"
				putBucketWebsiteOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutBucketWebsiteOptions(bucket, website, body)
				putBucketWebsiteOptionsModel.SetBucket("testString")
				putBucketWebsiteOptionsModel.SetWebsite(true)
				putBucketWebsiteOptionsModel.SetBody("testString")
				putBucketWebsiteOptionsModel.SetContentMD5("testString")
				putBucketWebsiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putBucketWebsiteOptionsModel).ToNot(BeNil())
				Expect(putBucketWebsiteOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putBucketWebsiteOptionsModel.Website).To(Equal(core.BoolPtr(true)))
				Expect(putBucketWebsiteOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putBucketWebsiteOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(putBucketWebsiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutObjectAclOptions successfully`, func() {
				// Construct an instance of the PutObjectAclOptions model
				bucket := "testString"
				key := "testString"
				acl := true
				body := "testString"
				putObjectAclOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutObjectAclOptions(bucket, key, acl, body)
				putObjectAclOptionsModel.SetBucket("testString")
				putObjectAclOptionsModel.SetKey("testString")
				putObjectAclOptionsModel.SetAcl(true)
				putObjectAclOptionsModel.SetBody("testString")
				putObjectAclOptionsModel.SetXAmzAcl("private")
				putObjectAclOptionsModel.SetContentMD5("testString")
				putObjectAclOptionsModel.SetXAmzGrantFullControl("testString")
				putObjectAclOptionsModel.SetXAmzGrantRead("testString")
				putObjectAclOptionsModel.SetXAmzGrantReadAcp("testString")
				putObjectAclOptionsModel.SetXAmzGrantWrite("testString")
				putObjectAclOptionsModel.SetXAmzGrantWriteAcp("testString")
				putObjectAclOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putObjectAclOptionsModel).ToNot(BeNil())
				Expect(putObjectAclOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.Acl).To(Equal(core.BoolPtr(true)))
				Expect(putObjectAclOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.XAmzAcl).To(Equal(core.StringPtr("private")))
				Expect(putObjectAclOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.XAmzGrantFullControl).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.XAmzGrantRead).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.XAmzGrantReadAcp).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.XAmzGrantWrite).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.XAmzGrantWriteAcp).To(Equal(core.StringPtr("testString")))
				Expect(putObjectAclOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutObjectOptions successfully`, func() {
				// Construct an instance of the PutObjectOptions model
				bucket := "testString"
				key := "testString"
				body := "testString"
				putObjectOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutObjectOptions(bucket, key, body)
				putObjectOptionsModel.SetBucket("testString")
				putObjectOptionsModel.SetKey("testString")
				putObjectOptionsModel.SetBody("testString")
				putObjectOptionsModel.SetXAmzAcl("private")
				putObjectOptionsModel.SetIfMatch("testString")
				putObjectOptionsModel.SetIfNoneMatch("testString")
				putObjectOptionsModel.SetIfUnmodifiedSince(CreateMockDateTime())
				putObjectOptionsModel.SetCacheControl("testString")
				putObjectOptionsModel.SetContentDisposition("testString")
				putObjectOptionsModel.SetContentEncoding("testString")
				putObjectOptionsModel.SetContentLanguage("testString")
				putObjectOptionsModel.SetContentLength(int64(38))
				putObjectOptionsModel.SetContentMD5("testString")
				putObjectOptionsModel.SetExpires(CreateMockDateTime())
				putObjectOptionsModel.SetXAmzServerSideEncryption("AES256")
				putObjectOptionsModel.SetXAmzWebsiteRedirectLocation("testString")
				putObjectOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				putObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				putObjectOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				putObjectOptionsModel.SetXAmzTagging("testString")
				putObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putObjectOptionsModel).ToNot(BeNil())
				Expect(putObjectOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.XAmzAcl).To(Equal(core.StringPtr("private")))
				Expect(putObjectOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.IfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.IfUnmodifiedSince).To(Equal(CreateMockDateTime()))
				Expect(putObjectOptionsModel.CacheControl).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.ContentDisposition).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.ContentEncoding).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.ContentLanguage).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.ContentLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(putObjectOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.Expires).To(Equal(CreateMockDateTime()))
				Expect(putObjectOptionsModel.XAmzServerSideEncryption).To(Equal(core.StringPtr("AES256")))
				Expect(putObjectOptionsModel.XAmzWebsiteRedirectLocation).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.XAmzTagging).To(Equal(core.StringPtr("testString")))
				Expect(putObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutObjectTaggingOptions successfully`, func() {
				// Construct an instance of the PutObjectTaggingOptions model
				bucket := "testString"
				key := "testString"
				tagging := true
				body := "testString"
				putObjectTaggingOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutObjectTaggingOptions(bucket, key, tagging, body)
				putObjectTaggingOptionsModel.SetBucket("testString")
				putObjectTaggingOptionsModel.SetKey("testString")
				putObjectTaggingOptionsModel.SetTagging(true)
				putObjectTaggingOptionsModel.SetBody("testString")
				putObjectTaggingOptionsModel.SetContentMD5("testString")
				putObjectTaggingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putObjectTaggingOptionsModel).ToNot(BeNil())
				Expect(putObjectTaggingOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putObjectTaggingOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(putObjectTaggingOptionsModel.Tagging).To(Equal(core.BoolPtr(true)))
				Expect(putObjectTaggingOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putObjectTaggingOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(putObjectTaggingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPutPublicAccessBlockOptions successfully`, func() {
				// Construct an instance of the PutPublicAccessBlockOptions model
				bucket := "testString"
				publicAccessBlock := true
				body := "testString"
				putPublicAccessBlockOptionsModel := ibmCloudObjectStorageS3ApiService.NewPutPublicAccessBlockOptions(bucket, publicAccessBlock, body)
				putPublicAccessBlockOptionsModel.SetBucket("testString")
				putPublicAccessBlockOptionsModel.SetPublicAccessBlock(true)
				putPublicAccessBlockOptionsModel.SetBody("testString")
				putPublicAccessBlockOptionsModel.SetContentMD5("testString")
				putPublicAccessBlockOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(putPublicAccessBlockOptionsModel).ToNot(BeNil())
				Expect(putPublicAccessBlockOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(putPublicAccessBlockOptionsModel.PublicAccessBlock).To(Equal(core.BoolPtr(true)))
				Expect(putPublicAccessBlockOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(putPublicAccessBlockOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(putPublicAccessBlockOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreObjectOptions successfully`, func() {
				// Construct an instance of the RestoreObjectOptions model
				bucket := "testString"
				key := "testString"
				restore := true
				body := "testString"
				restoreObjectOptionsModel := ibmCloudObjectStorageS3ApiService.NewRestoreObjectOptions(bucket, key, restore, body)
				restoreObjectOptionsModel.SetBucket("testString")
				restoreObjectOptionsModel.SetKey("testString")
				restoreObjectOptionsModel.SetRestore(true)
				restoreObjectOptionsModel.SetBody("testString")
				restoreObjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreObjectOptionsModel).ToNot(BeNil())
				Expect(restoreObjectOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(restoreObjectOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(restoreObjectOptionsModel.Restore).To(Equal(core.BoolPtr(true)))
				Expect(restoreObjectOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(restoreObjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadPartCopyOptions successfully`, func() {
				// Construct an instance of the UploadPartCopyOptions model
				bucket := "testString"
				xAmzCopySource := "testString"
				key := "testString"
				partNumber := int64(38)
				uploadID := "testString"
				uploadPartCopyOptionsModel := ibmCloudObjectStorageS3ApiService.NewUploadPartCopyOptions(bucket, xAmzCopySource, key, partNumber, uploadID)
				uploadPartCopyOptionsModel.SetBucket("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySource("testString")
				uploadPartCopyOptionsModel.SetKey("testString")
				uploadPartCopyOptionsModel.SetPartNumber(int64(38))
				uploadPartCopyOptionsModel.SetUploadID("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySourceIfMatch("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySourceIfModifiedSince(CreateMockDateTime())
				uploadPartCopyOptionsModel.SetXAmzCopySourceIfNoneMatch("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySourceIfUnmodifiedSince(CreateMockDateTime())
				uploadPartCopyOptionsModel.SetXAmzCopySourceRange("testString")
				uploadPartCopyOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				uploadPartCopyOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				uploadPartCopyOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySourceServerSideEncryptionCustomerAlgorithm("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySourceServerSideEncryptionCustomerKey("testString")
				uploadPartCopyOptionsModel.SetXAmzCopySourceServerSideEncryptionCustomerKeyMD5("testString")
				uploadPartCopyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadPartCopyOptionsModel).ToNot(BeNil())
				Expect(uploadPartCopyOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySource).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.PartNumber).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uploadPartCopyOptionsModel.UploadID).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceIfMatch).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceIfModifiedSince).To(Equal(CreateMockDateTime()))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceIfNoneMatch).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceIfUnmodifiedSince).To(Equal(CreateMockDateTime()))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceRange).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.XAmzCopySourceServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartCopyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadPartOptions successfully`, func() {
				// Construct an instance of the UploadPartOptions model
				bucket := "testString"
				key := "testString"
				partNumber := int64(38)
				uploadID := "testString"
				body := "testString"
				uploadPartOptionsModel := ibmCloudObjectStorageS3ApiService.NewUploadPartOptions(bucket, key, partNumber, uploadID, body)
				uploadPartOptionsModel.SetBucket("testString")
				uploadPartOptionsModel.SetKey("testString")
				uploadPartOptionsModel.SetPartNumber(int64(38))
				uploadPartOptionsModel.SetUploadID("testString")
				uploadPartOptionsModel.SetBody("testString")
				uploadPartOptionsModel.SetContentLength(int64(38))
				uploadPartOptionsModel.SetContentMD5("testString")
				uploadPartOptionsModel.SetXAmzServerSideEncryptionCustomerAlgorithm("testString")
				uploadPartOptionsModel.SetXAmzServerSideEncryptionCustomerKey("testString")
				uploadPartOptionsModel.SetXAmzServerSideEncryptionCustomerKeyMD5("testString")
				uploadPartOptionsModel.SetXAmzRequestPayer("requester")
				uploadPartOptionsModel.SetXAmzExpectedBucketOwner("testString")
				uploadPartOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadPartOptionsModel).ToNot(BeNil())
				Expect(uploadPartOptionsModel.Bucket).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.PartNumber).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uploadPartOptionsModel.UploadID).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.ContentLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uploadPartOptionsModel.ContentMD5).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.XAmzServerSideEncryptionCustomerAlgorithm).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.XAmzServerSideEncryptionCustomerKey).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.XAmzServerSideEncryptionCustomerKeyMD5).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.XAmzRequestPayer).To(Equal(core.StringPtr("requester")))
				Expect(uploadPartOptionsModel.XAmzExpectedBucketOwner).To(Equal(core.StringPtr("testString")))
				Expect(uploadPartOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
