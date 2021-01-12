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

package ibmclouddatabasesv5_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/experimental-go-sdk/ibmclouddatabasesv5"
	"github.com/IBM/go-sdk-core/v4/core"
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

var _ = Describe(`IbmCloudDatabasesV5`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetDeployables(getDeployablesOptions *GetDeployablesOptions) - Operation response error`, func() {
		getDeployablesPath := "/deployables"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeployablesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeployables with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeployablesOptions model
				getDeployablesOptionsModel := new(ibmclouddatabasesv5.GetDeployablesOptions)
				getDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDeployables(getDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDeployables(getDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDeployables(getDeployablesOptions *GetDeployablesOptions)`, func() {
		getDeployablesPath := "/deployables"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeployablesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployables": [{"type": "elasticsearch", "versions": [{"version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "from_version": "5.6", "to_version": "6.7"}]}]}]}`)
				}))
			})
			It(`Invoke GetDeployables successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeployablesOptions model
				getDeployablesOptionsModel := new(ibmclouddatabasesv5.GetDeployablesOptions)
				getDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDeployablesWithContext(ctx, getDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDeployables(getDeployablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDeployablesWithContext(ctx, getDeployablesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDeployablesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployables": [{"type": "elasticsearch", "versions": [{"version": "5.6", "status": "stable", "is_preferred": true, "transitions": [{"application": "elasticsearch", "method": "restore", "from_version": "5.6", "to_version": "6.7"}]}]}]}`)
				}))
			})
			It(`Invoke GetDeployables successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDeployables(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeployablesOptions model
				getDeployablesOptionsModel := new(ibmclouddatabasesv5.GetDeployablesOptions)
				getDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeployables(getDeployablesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeployables with error: Operation request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeployablesOptions model
				getDeployablesOptionsModel := new(ibmclouddatabasesv5.GetDeployablesOptions)
				getDeployablesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDeployables(getDeployablesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRegions(getRegionsOptions *GetRegionsOptions) - Operation response error`, func() {
		getRegionsPath := "/regions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRegions with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(ibmclouddatabasesv5.GetRegionsOptions)
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRegions(getRegionsOptions *GetRegionsOptions)`, func() {
		getRegionsPath := "/regions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regions": ["Regions"]}`)
				}))
			})
			It(`Invoke GetRegions successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(ibmclouddatabasesv5.GetRegionsOptions)
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetRegionsWithContext(ctx, getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetRegionsWithContext(ctx, getRegionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"regions": ["Regions"]}`)
				}))
			})
			It(`Invoke GetRegions successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetRegions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(ibmclouddatabasesv5.GetRegionsOptions)
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRegions with error: Operation request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := new(ibmclouddatabasesv5.GetRegionsOptions)
				getRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetRegions(getRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions) - Operation response error`, func() {
		getDeploymentInfoPath := "/deployments/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeploymentInfo with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(ibmclouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDeploymentInfo(getDeploymentInfoOptions *GetDeploymentInfoOptions)`, func() {
		getDeploymentInfoPath := "/deployments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment": {"id": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "name": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "type": "redis", "platform_options": {"anyKey": "anyValue"}, "version": "4", "admin_usernames": "database: admin", "enable_public_endpoints": true, "enable_private_endpoints": false}}`)
				}))
			})
			It(`Invoke GetDeploymentInfo successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(ibmclouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDeploymentInfoWithContext(ctx, getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDeploymentInfoWithContext(ctx, getDeploymentInfoOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deployment": {"id": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "name": "crn:v1:bluemix:public:databases-for-redis:us-south:a/274074dce64e9c423ffc238516c755e1:29caf0e7-120f-4da8-9551-3abf57ebcfc7::", "type": "redis", "platform_options": {"anyKey": "anyValue"}, "version": "4", "admin_usernames": "database: admin", "enable_public_endpoints": true, "enable_private_endpoints": false}}`)
				}))
			})
			It(`Invoke GetDeploymentInfo successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(ibmclouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeploymentInfo with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentInfoOptions model
				getDeploymentInfoOptionsModel := new(ibmclouddatabasesv5.GetDeploymentInfoOptions)
				getDeploymentInfoOptionsModel.ID = core.StringPtr("testString")
				getDeploymentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDeploymentInfoOptions model with no property values
				getDeploymentInfoOptionsModelNew := new(ibmclouddatabasesv5.GetDeploymentInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentInfo(getDeploymentInfoOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions) - Operation response error`, func() {
		createDatabaseUserPath := "/deployments/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDatabaseUser with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(ibmclouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(ibmclouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateDatabaseUser(createDatabaseUserOptions *CreateDatabaseUserOptions)`, func() {
		createDatabaseUserPath := "/deployments/testString/users/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke CreateDatabaseUser successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(ibmclouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(ibmclouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.CreateDatabaseUserWithContext(ctx, createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.CreateDatabaseUserWithContext(ctx, createDatabaseUserOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDatabaseUserPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke CreateDatabaseUser successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.CreateDatabaseUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(ibmclouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(ibmclouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDatabaseUser with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(ibmclouddatabasesv5.CreateDatabaseUserRequestUser)
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")

				// Construct an instance of the CreateDatabaseUserOptions model
				createDatabaseUserOptionsModel := new(ibmclouddatabasesv5.CreateDatabaseUserOptions)
				createDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				createDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				createDatabaseUserOptionsModel.User = createDatabaseUserRequestUserModel
				createDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDatabaseUserOptions model with no property values
				createDatabaseUserOptionsModelNew := new(ibmclouddatabasesv5.CreateDatabaseUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.CreateDatabaseUser(createDatabaseUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions) - Operation response error`, func() {
		changeUserPasswordPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(changeUserPasswordPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangeUserPassword with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(ibmclouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(ibmclouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeUserPassword(changeUserPasswordOptions *ChangeUserPasswordOptions)`, func() {
		changeUserPasswordPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(changeUserPasswordPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke ChangeUserPassword successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(ibmclouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(ibmclouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.ChangeUserPasswordWithContext(ctx, changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.ChangeUserPasswordWithContext(ctx, changeUserPasswordOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(changeUserPasswordPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke ChangeUserPassword successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.ChangeUserPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(ibmclouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(ibmclouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ChangeUserPassword with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(ibmclouddatabasesv5.APasswordSettingUser)
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")

				// Construct an instance of the ChangeUserPasswordOptions model
				changeUserPasswordOptionsModel := new(ibmclouddatabasesv5.ChangeUserPasswordOptions)
				changeUserPasswordOptionsModel.ID = core.StringPtr("testString")
				changeUserPasswordOptionsModel.UserType = core.StringPtr("testString")
				changeUserPasswordOptionsModel.Username = core.StringPtr("testString")
				changeUserPasswordOptionsModel.User = aPasswordSettingUserModel
				changeUserPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ChangeUserPasswordOptions model with no property values
				changeUserPasswordOptionsModelNew := new(ibmclouddatabasesv5.ChangeUserPasswordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.ChangeUserPassword(changeUserPasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions) - Operation response error`, func() {
		deleteDatabaseUserPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDatabaseUser with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(ibmclouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteDatabaseUser(deleteDatabaseUserOptions *DeleteDatabaseUserOptions)`, func() {
		deleteDatabaseUserPath := "/deployments/testString/users/testString/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteDatabaseUser successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(ibmclouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.DeleteDatabaseUserWithContext(ctx, deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.DeleteDatabaseUserWithContext(ctx, deleteDatabaseUserOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteDatabaseUserPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteDatabaseUser successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.DeleteDatabaseUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(ibmclouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDatabaseUser with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteDatabaseUserOptions model
				deleteDatabaseUserOptionsModel := new(ibmclouddatabasesv5.DeleteDatabaseUserOptions)
				deleteDatabaseUserOptionsModel.ID = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.UserType = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Username = core.StringPtr("testString")
				deleteDatabaseUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDatabaseUserOptions model with no property values
				deleteDatabaseUserOptionsModelNew := new(ibmclouddatabasesv5.DeleteDatabaseUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.DeleteDatabaseUser(deleteDatabaseUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUser(getUserOptions *GetUserOptions) - Operation response error`, func() {
		getUserPath := "/deployments/testString/users/testString/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUser with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(ibmclouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.EndpointType = core.StringPtr("public")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUser(getUserOptions *GetUserOptions)`, func() {
		getUserPath := "/deployments/testString/users/testString/"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetUser successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(ibmclouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.EndpointType = core.StringPtr("public")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetUserWithContext(ctx, getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetUserWithContext(ctx, getUserOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetUser successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(ibmclouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.EndpointType = core.StringPtr("public")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUser with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(ibmclouddatabasesv5.GetUserOptions)
				getUserOptionsModel.ID = core.StringPtr("testString")
				getUserOptionsModel.UserID = core.StringPtr("testString")
				getUserOptionsModel.EndpointType = core.StringPtr("public")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserOptions model with no property values
				getUserOptionsModelNew := new(ibmclouddatabasesv5.GetUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetUser(getUserOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`SetDatabaseConfiguration(setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions) - Operation response error`, func() {
		setDatabaseConfigurationPath := "/deployments/testString/configuration"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetDatabaseConfiguration with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetConfigurationConfigurationPGConfiguration model
				setConfigurationConfigurationModel := new(ibmclouddatabasesv5.SetConfigurationConfigurationPGConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(ibmclouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetDatabaseConfiguration(setDatabaseConfigurationOptions *SetDatabaseConfigurationOptions)`, func() {
		setDatabaseConfigurationPath := "/deployments/testString/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetDatabaseConfiguration successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetConfigurationConfigurationPGConfiguration model
				setConfigurationConfigurationModel := new(ibmclouddatabasesv5.SetConfigurationConfigurationPGConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(ibmclouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.SetDatabaseConfigurationWithContext(ctx, setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.SetDatabaseConfigurationWithContext(ctx, setDatabaseConfigurationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setDatabaseConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetDatabaseConfiguration successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.SetDatabaseConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetConfigurationConfigurationPGConfiguration model
				setConfigurationConfigurationModel := new(ibmclouddatabasesv5.SetConfigurationConfigurationPGConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(ibmclouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetDatabaseConfiguration with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetConfigurationConfigurationPGConfiguration model
				setConfigurationConfigurationModel := new(ibmclouddatabasesv5.SetConfigurationConfigurationPGConfiguration)
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				setDatabaseConfigurationOptionsModel := new(ibmclouddatabasesv5.SetDatabaseConfigurationOptions)
				setDatabaseConfigurationOptionsModel.ID = core.StringPtr("testString")
				setDatabaseConfigurationOptionsModel.Configuration = setConfigurationConfigurationModel
				setDatabaseConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetDatabaseConfigurationOptions model with no property values
				setDatabaseConfigurationOptionsModelNew := new(ibmclouddatabasesv5.SetDatabaseConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.SetDatabaseConfiguration(setDatabaseConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions) - Operation response error`, func() {
		getDatabaseConfigurationSchemaPath := "/deployments/testString/configuration/schema"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseConfigurationSchemaPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDatabaseConfigurationSchema with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(ibmclouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptions *GetDatabaseConfigurationSchemaOptions)`, func() {
		getDatabaseConfigurationSchemaPath := "/deployments/testString/configuration/schema"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseConfigurationSchemaPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"schema": {"max_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_prepared_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "backup_retention_period": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "deadlock_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "effective_io_concurrency": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_replication_slots": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_wal_senders": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "shared_buffers": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "synchronous_commit": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "wal_level": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "archive_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "log_min_duration_statement": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}}}}`)
				}))
			})
			It(`Invoke GetDatabaseConfigurationSchema successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(ibmclouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDatabaseConfigurationSchemaWithContext(ctx, getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDatabaseConfigurationSchemaWithContext(ctx, getDatabaseConfigurationSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDatabaseConfigurationSchemaPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"schema": {"max_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_prepared_connections": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "backup_retention_period": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "deadlock_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "effective_io_concurrency": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_replication_slots": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "max_wal_senders": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "shared_buffers": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "synchronous_commit": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "wal_level": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "choices": ["Choices"]}, "archive_timeout": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}, "log_min_duration_statement": {"customer_configurable": true, "default": 7, "default_description": "DefaultDescription", "description": "Description", "kind": "Kind", "requires_restart": false, "min": 3, "max": 3, "step": 4}}}}`)
				}))
			})
			It(`Invoke GetDatabaseConfigurationSchema successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDatabaseConfigurationSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(ibmclouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDatabaseConfigurationSchema with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				getDatabaseConfigurationSchemaOptionsModel := new(ibmclouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				getDatabaseConfigurationSchemaOptionsModel.ID = core.StringPtr("testString")
				getDatabaseConfigurationSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDatabaseConfigurationSchemaOptions model with no property values
				getDatabaseConfigurationSchemaOptionsModelNew := new(ibmclouddatabasesv5.GetDatabaseConfigurationSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetDatabaseConfigurationSchema(getDatabaseConfigurationSchemaOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetRemotes(getRemotesOptions *GetRemotesOptions) - Operation response error`, func() {
		getRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRemotes with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRemotesOptions model
				getRemotesOptionsModel := new(ibmclouddatabasesv5.GetRemotesOptions)
				getRemotesOptionsModel.ID = core.StringPtr("testString")
				getRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetRemotes(getRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetRemotes(getRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRemotes(getRemotesOptions *GetRemotesOptions)`, func() {
		getRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"remotes": {"leader": "01f30581-54f8-41a4-8193-4a04cc022e9b-h", "replicas": ["Replicas"]}}`)
				}))
			})
			It(`Invoke GetRemotes successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetRemotesOptions model
				getRemotesOptionsModel := new(ibmclouddatabasesv5.GetRemotesOptions)
				getRemotesOptionsModel.ID = core.StringPtr("testString")
				getRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetRemotesWithContext(ctx, getRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetRemotes(getRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetRemotesWithContext(ctx, getRemotesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"remotes": {"leader": "01f30581-54f8-41a4-8193-4a04cc022e9b-h", "replicas": ["Replicas"]}}`)
				}))
			})
			It(`Invoke GetRemotes successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetRemotes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRemotesOptions model
				getRemotesOptionsModel := new(ibmclouddatabasesv5.GetRemotesOptions)
				getRemotesOptionsModel.ID = core.StringPtr("testString")
				getRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetRemotes(getRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRemotes with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRemotesOptions model
				getRemotesOptionsModel := new(ibmclouddatabasesv5.GetRemotesOptions)
				getRemotesOptionsModel.ID = core.StringPtr("testString")
				getRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetRemotes(getRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRemotesOptions model with no property values
				getRemotesOptionsModelNew := new(ibmclouddatabasesv5.GetRemotesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetRemotes(getRemotesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetRemotes(setRemotesOptions *SetRemotesOptions) - Operation response error`, func() {
		setRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setRemotesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetRemotes with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetRemotesRequestRemotes model
				setRemotesRequestRemotesModel := new(ibmclouddatabasesv5.SetRemotesRequestRemotes)
				setRemotesRequestRemotesModel.Leader = core.StringPtr("testString")

				// Construct an instance of the SetRemotesOptions model
				setRemotesOptionsModel := new(ibmclouddatabasesv5.SetRemotesOptions)
				setRemotesOptionsModel.ID = core.StringPtr("testString")
				setRemotesOptionsModel.Remotes = setRemotesRequestRemotesModel
				setRemotesOptionsModel.SkipInitialBackup = core.BoolPtr(true)
				setRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.SetRemotes(setRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.SetRemotes(setRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetRemotes(setRemotesOptions *SetRemotesOptions)`, func() {
		setRemotesPath := "/deployments/testString/remotes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setRemotesPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetRemotes successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetRemotesRequestRemotes model
				setRemotesRequestRemotesModel := new(ibmclouddatabasesv5.SetRemotesRequestRemotes)
				setRemotesRequestRemotesModel.Leader = core.StringPtr("testString")

				// Construct an instance of the SetRemotesOptions model
				setRemotesOptionsModel := new(ibmclouddatabasesv5.SetRemotesOptions)
				setRemotesOptionsModel.ID = core.StringPtr("testString")
				setRemotesOptionsModel.Remotes = setRemotesRequestRemotesModel
				setRemotesOptionsModel.SkipInitialBackup = core.BoolPtr(true)
				setRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.SetRemotesWithContext(ctx, setRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.SetRemotes(setRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.SetRemotesWithContext(ctx, setRemotesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setRemotesPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetRemotes successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.SetRemotes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetRemotesRequestRemotes model
				setRemotesRequestRemotesModel := new(ibmclouddatabasesv5.SetRemotesRequestRemotes)
				setRemotesRequestRemotesModel.Leader = core.StringPtr("testString")

				// Construct an instance of the SetRemotesOptions model
				setRemotesOptionsModel := new(ibmclouddatabasesv5.SetRemotesOptions)
				setRemotesOptionsModel.ID = core.StringPtr("testString")
				setRemotesOptionsModel.Remotes = setRemotesRequestRemotesModel
				setRemotesOptionsModel.SkipInitialBackup = core.BoolPtr(true)
				setRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.SetRemotes(setRemotesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetRemotes with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetRemotesRequestRemotes model
				setRemotesRequestRemotesModel := new(ibmclouddatabasesv5.SetRemotesRequestRemotes)
				setRemotesRequestRemotesModel.Leader = core.StringPtr("testString")

				// Construct an instance of the SetRemotesOptions model
				setRemotesOptionsModel := new(ibmclouddatabasesv5.SetRemotesOptions)
				setRemotesOptionsModel.ID = core.StringPtr("testString")
				setRemotesOptionsModel.Remotes = setRemotesRequestRemotesModel
				setRemotesOptionsModel.SkipInitialBackup = core.BoolPtr(true)
				setRemotesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.SetRemotes(setRemotesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetRemotesOptions model with no property values
				setRemotesOptionsModelNew := new(ibmclouddatabasesv5.SetRemotesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.SetRemotes(setRemotesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRemotesSchema(getRemotesSchemaOptions *GetRemotesSchemaOptions) - Operation response error`, func() {
		getRemotesSchemaPath := "/deployments/testString/remotes/resync"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRemotesSchema with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(ibmclouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRemotesSchema(getRemotesSchemaOptions *GetRemotesSchemaOptions)`, func() {
		getRemotesSchemaPath := "/deployments/testString/remotes/resync"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesSchemaPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetRemotesSchema successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(ibmclouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetRemotesSchemaWithContext(ctx, getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetRemotesSchemaWithContext(ctx, getRemotesSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRemotesSchemaPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetRemotesSchema successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetRemotesSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(ibmclouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRemotesSchema with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetRemotesSchemaOptions model
				getRemotesSchemaOptionsModel := new(ibmclouddatabasesv5.GetRemotesSchemaOptions)
				getRemotesSchemaOptionsModel.ID = core.StringPtr("testString")
				getRemotesSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRemotesSchemaOptions model with no property values
				getRemotesSchemaOptionsModelNew := new(ibmclouddatabasesv5.GetRemotesSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetRemotesSchema(getRemotesSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetPromotion(setPromotionOptions *SetPromotionOptions) - Operation response error`, func() {
		setPromotionPath := "/deployments/testString/remotes/promotion"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setPromotionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetPromotion with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(ibmclouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(ibmclouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetPromotion(setPromotionOptions *SetPromotionOptions)`, func() {
		setPromotionPath := "/deployments/testString/remotes/promotion"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setPromotionPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetPromotion successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(ibmclouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(ibmclouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.SetPromotionWithContext(ctx, setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.SetPromotionWithContext(ctx, setPromotionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setPromotionPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetPromotion successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.SetPromotion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(ibmclouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(ibmclouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetPromotion with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(ibmclouddatabasesv5.SetPromotionPromotionPromote)
				setPromotionPromotionModel.Promotion = make(map[string]interface{})

				// Construct an instance of the SetPromotionOptions model
				setPromotionOptionsModel := new(ibmclouddatabasesv5.SetPromotionOptions)
				setPromotionOptionsModel.ID = core.StringPtr("testString")
				setPromotionOptionsModel.Promotion = setPromotionPromotionModel
				setPromotionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.SetPromotion(setPromotionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetPromotionOptions model with no property values
				setPromotionOptionsModelNew := new(ibmclouddatabasesv5.SetPromotionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.SetPromotion(setPromotionOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetDeploymentTasks(getDeploymentTasksOptions *GetDeploymentTasksOptions) - Operation response error`, func() {
		getDeploymentTasksPath := "/deployments/testString/tasks"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeploymentTasks with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentTasksOptions model
				getDeploymentTasksOptionsModel := new(ibmclouddatabasesv5.GetDeploymentTasksOptions)
				getDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				getDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentTasks(getDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentTasks(getDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDeploymentTasks(getDeploymentTasksOptions *GetDeploymentTasksOptions)`, func() {
		getDeploymentTasksPath := "/deployments/testString/tasks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke GetDeploymentTasks successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentTasksOptions model
				getDeploymentTasksOptionsModel := new(ibmclouddatabasesv5.GetDeploymentTasksOptions)
				getDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				getDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDeploymentTasksWithContext(ctx, getDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentTasks(getDeploymentTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDeploymentTasksWithContext(ctx, getDeploymentTasksOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke GetDeploymentTasks successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentTasks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentTasksOptions model
				getDeploymentTasksOptionsModel := new(ibmclouddatabasesv5.GetDeploymentTasksOptions)
				getDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				getDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentTasks(getDeploymentTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeploymentTasks with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentTasksOptions model
				getDeploymentTasksOptionsModel := new(ibmclouddatabasesv5.GetDeploymentTasksOptions)
				getDeploymentTasksOptionsModel.ID = core.StringPtr("testString")
				getDeploymentTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentTasks(getDeploymentTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDeploymentTasksOptions model with no property values
				getDeploymentTasksOptionsModelNew := new(ibmclouddatabasesv5.GetDeploymentTasksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentTasks(getDeploymentTasksOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTasks(getTasksOptions *GetTasksOptions) - Operation response error`, func() {
		getTasksPath := "/tasks/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTasks with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetTasksOptions model
				getTasksOptionsModel := new(ibmclouddatabasesv5.GetTasksOptions)
				getTasksOptionsModel.ID = core.StringPtr("testString")
				getTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetTasks(getTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetTasks(getTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTasks(getTasksOptions *GetTasksOptions)`, func() {
		getTasksPath := "/tasks/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetTasks successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetTasksOptions model
				getTasksOptionsModel := new(ibmclouddatabasesv5.GetTasksOptions)
				getTasksOptionsModel.ID = core.StringPtr("testString")
				getTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetTasksWithContext(ctx, getTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetTasks(getTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetTasksWithContext(ctx, getTasksOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetTasks successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetTasks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTasksOptions model
				getTasksOptionsModel := new(ibmclouddatabasesv5.GetTasksOptions)
				getTasksOptionsModel.ID = core.StringPtr("testString")
				getTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetTasks(getTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTasks with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetTasksOptions model
				getTasksOptionsModel := new(ibmclouddatabasesv5.GetTasksOptions)
				getTasksOptionsModel.ID = core.StringPtr("testString")
				getTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetTasks(getTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTasksOptions model with no property values
				getTasksOptionsModelNew := new(ibmclouddatabasesv5.GetTasksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetTasks(getTasksOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions) - Operation response error`, func() {
		getBackupInfoPath := "/backups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBackupInfo with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(ibmclouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetBackupInfo(getBackupInfoOptions *GetBackupInfoOptions)`, func() {
		getBackupInfoPath := "/backups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backup": {"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetBackupInfo successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(ibmclouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetBackupInfoWithContext(ctx, getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetBackupInfoWithContext(ctx, getBackupInfoOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBackupInfoPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backup": {"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetBackupInfo successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetBackupInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(ibmclouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBackupInfo with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetBackupInfoOptions model
				getBackupInfoOptionsModel := new(ibmclouddatabasesv5.GetBackupInfoOptions)
				getBackupInfoOptionsModel.BackupID = core.StringPtr("testString")
				getBackupInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBackupInfoOptions model with no property values
				getBackupInfoOptionsModelNew := new(ibmclouddatabasesv5.GetBackupInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetBackupInfo(getBackupInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDeploymentBackups(getDeploymentBackupsOptions *GetDeploymentBackupsOptions) - Operation response error`, func() {
		getDeploymentBackupsPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeploymentBackups with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentBackupsOptions model
				getDeploymentBackupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentBackupsOptions)
				getDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentBackups(getDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentBackups(getDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDeploymentBackups(getDeploymentBackupsOptions *GetDeploymentBackupsOptions)`, func() {
		getDeploymentBackupsPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke GetDeploymentBackups successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentBackupsOptions model
				getDeploymentBackupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentBackupsOptions)
				getDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDeploymentBackupsWithContext(ctx, getDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentBackups(getDeploymentBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDeploymentBackupsWithContext(ctx, getDeploymentBackupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentBackupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "5a970218cb7544000671c094", "deployment_id": "595eada310b7ac00116dd48b", "type": "scheduled", "status": "running", "is_downloadable": true, "is_restorable": true, "created_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke GetDeploymentBackups successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentBackups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentBackupsOptions model
				getDeploymentBackupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentBackupsOptions)
				getDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentBackups(getDeploymentBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeploymentBackups with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentBackupsOptions model
				getDeploymentBackupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentBackupsOptions)
				getDeploymentBackupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentBackups(getDeploymentBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDeploymentBackupsOptions model with no property values
				getDeploymentBackupsOptionsModelNew := new(ibmclouddatabasesv5.GetDeploymentBackupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentBackups(getDeploymentBackupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions) - Operation response error`, func() {
		startOndemandBackupPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke StartOndemandBackup with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(ibmclouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`StartOndemandBackup(startOndemandBackupOptions *StartOndemandBackupOptions)`, func() {
		startOndemandBackupPath := "/deployments/testString/backups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke StartOndemandBackup successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(ibmclouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.StartOndemandBackupWithContext(ctx, startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.StartOndemandBackupWithContext(ctx, startOndemandBackupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(startOndemandBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke StartOndemandBackup successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.StartOndemandBackup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(ibmclouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke StartOndemandBackup with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the StartOndemandBackupOptions model
				startOndemandBackupOptionsModel := new(ibmclouddatabasesv5.StartOndemandBackupOptions)
				startOndemandBackupOptionsModel.ID = core.StringPtr("testString")
				startOndemandBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the StartOndemandBackupOptions model with no property values
				startOndemandBackupOptionsModelNew := new(ibmclouddatabasesv5.StartOndemandBackupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.StartOndemandBackup(startOndemandBackupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPITRdata(getPITRdataOptions *GetPITRdataOptions) - Operation response error`, func() {
		getPitRdataPath := "/deployments/testString/point_in_time_recovery_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitRdataPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPITRdata with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPITRdataOptions model
				getPitRdataOptionsModel := new(ibmclouddatabasesv5.GetPITRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetPITRdata(getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetPITRdata(getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPITRdata(getPITRdataOptions *GetPITRdataOptions)`, func() {
		getPitRdataPath := "/deployments/testString/point_in_time_recovery_data"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPitRdataPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"earliest_point_in_time_recovery_time": "EarliestPointInTimeRecoveryTime"}`)
				}))
			})
			It(`Invoke GetPITRdata successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetPITRdataOptions model
				getPitRdataOptionsModel := new(ibmclouddatabasesv5.GetPITRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetPITRdataWithContext(ctx, getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetPITRdata(getPitRdataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetPITRdataWithContext(ctx, getPitRdataOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPitRdataPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"earliest_point_in_time_recovery_time": "EarliestPointInTimeRecoveryTime"}`)
				}))
			})
			It(`Invoke GetPITRdata successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetPITRdata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPITRdataOptions model
				getPitRdataOptionsModel := new(ibmclouddatabasesv5.GetPITRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetPITRdata(getPitRdataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPITRdata with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetPITRdataOptions model
				getPitRdataOptionsModel := new(ibmclouddatabasesv5.GetPITRdataOptions)
				getPitRdataOptionsModel.ID = core.StringPtr("testString")
				getPitRdataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetPITRdata(getPitRdataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPITRdataOptions model with no property values
				getPitRdataOptionsModelNew := new(ibmclouddatabasesv5.GetPITRdataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetPITRdata(getPitRdataOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetConnection(getConnectionOptions *GetConnectionOptions) - Operation response error`, func() {
		getConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConnection with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(ibmclouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetConnection(getConnectionOptions *GetConnectionOptions)`, func() {
		getConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke GetConnection successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(ibmclouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetConnectionWithContext(ctx, getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetConnectionWithContext(ctx, getConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke GetConnection successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(ibmclouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConnection with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionOptions model
				getConnectionOptionsModel := new(ibmclouddatabasesv5.GetConnectionOptions)
				getConnectionOptionsModel.ID = core.StringPtr("testString")
				getConnectionOptionsModel.UserType = core.StringPtr("testString")
				getConnectionOptionsModel.UserID = core.StringPtr("testString")
				getConnectionOptionsModel.EndpointType = core.StringPtr("public")
				getConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetConnection(getConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConnectionOptions model with no property values
				getConnectionOptionsModelNew := new(ibmclouddatabasesv5.GetConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetConnection(getConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions) - Operation response error`, func() {
		completeConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteConnection with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CompleteConnection(completeConnectionOptions *CompleteConnectionOptions)`, func() {
		completeConnectionPath := "/deployments/testString/users/testString/testString/connections/public"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke CompleteConnection successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.CompleteConnectionWithContext(ctx, completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.CompleteConnectionWithContext(ctx, completeConnectionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke CompleteConnection successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteConnection with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionOptions model
				completeConnectionOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionOptions)
				completeConnectionOptionsModel.ID = core.StringPtr("testString")
				completeConnectionOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionOptionsModel.EndpointType = core.StringPtr("public")
				completeConnectionOptionsModel.Password = core.StringPtr("testString")
				completeConnectionOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnection(completeConnectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteConnectionOptions model with no property values
				completeConnectionOptionsModelNew := new(ibmclouddatabasesv5.CompleteConnectionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.CompleteConnection(completeConnectionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConnectionDeprecated(getConnectionDeprecatedOptions *GetConnectionDeprecatedOptions) - Operation response error`, func() {
		getConnectionDeprecatedPath := "/deployments/testString/users/testString/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionDeprecatedPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConnectionDeprecated with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionDeprecatedOptions model
				getConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.GetConnectionDeprecatedOptions)
				getConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetConnectionDeprecated(getConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetConnectionDeprecated(getConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetConnectionDeprecated(getConnectionDeprecatedOptions *GetConnectionDeprecatedOptions)`, func() {
		getConnectionDeprecatedPath := "/deployments/testString/users/testString/testString/connections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionDeprecatedPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke GetConnectionDeprecated successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetConnectionDeprecatedOptions model
				getConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.GetConnectionDeprecatedOptions)
				getConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetConnectionDeprecatedWithContext(ctx, getConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetConnectionDeprecated(getConnectionDeprecatedOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetConnectionDeprecatedWithContext(ctx, getConnectionDeprecatedOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getConnectionDeprecatedPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["certificate_root"]).To(Equal([]string{"testString"}))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke GetConnectionDeprecated successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetConnectionDeprecated(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConnectionDeprecatedOptions model
				getConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.GetConnectionDeprecatedOptions)
				getConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetConnectionDeprecated(getConnectionDeprecatedOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConnectionDeprecated with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetConnectionDeprecatedOptions model
				getConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.GetConnectionDeprecatedOptions)
				getConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				getConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetConnectionDeprecated(getConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConnectionDeprecatedOptions model with no property values
				getConnectionDeprecatedOptionsModelNew := new(ibmclouddatabasesv5.GetConnectionDeprecatedOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetConnectionDeprecated(getConnectionDeprecatedOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CompleteConnectionDeprecated(completeConnectionDeprecatedOptions *CompleteConnectionDeprecatedOptions) - Operation response error`, func() {
		completeConnectionDeprecatedPath := "/deployments/testString/users/testString/testString/connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionDeprecatedPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CompleteConnectionDeprecated with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionDeprecatedOptions model
				completeConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionDeprecatedOptions)
				completeConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Password = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnectionDeprecated(completeConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.CompleteConnectionDeprecated(completeConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CompleteConnectionDeprecated(completeConnectionDeprecatedOptions *CompleteConnectionDeprecatedOptions)`, func() {
		completeConnectionDeprecatedPath := "/deployments/testString/users/testString/testString/connections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionDeprecatedPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke CompleteConnectionDeprecated successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the CompleteConnectionDeprecatedOptions model
				completeConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionDeprecatedOptions)
				completeConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Password = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.CompleteConnectionDeprecatedWithContext(ctx, completeConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnectionDeprecated(completeConnectionDeprecatedOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.CompleteConnectionDeprecatedWithContext(ctx, completeConnectionDeprecatedOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(completeConnectionDeprecatedPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"connection": {"postgres": {"type": "uri", "composed": ["Composed"], "scheme": "Scheme", "hosts": [{"hostname": "Hostname", "port": 4}], "path": "/ibmclouddb", "query_options": {"anyKey": "anyValue"}, "authentication": {"method": "Method", "username": "Username", "password": "Password"}, "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}, "database": "Database"}, "cli": {"type": "cli", "composed": ["Composed"], "environment": {"mapKey": "Inner"}, "bin": "Bin", "arguments": [["Arguments"]], "certificate": {"name": "Name", "certificate_base64": "CertificateBase64"}}}}`)
				}))
			})
			It(`Invoke CompleteConnectionDeprecated successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnectionDeprecated(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CompleteConnectionDeprecatedOptions model
				completeConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionDeprecatedOptions)
				completeConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Password = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.CompleteConnectionDeprecated(completeConnectionDeprecatedOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CompleteConnectionDeprecated with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the CompleteConnectionDeprecatedOptions model
				completeConnectionDeprecatedOptionsModel := new(ibmclouddatabasesv5.CompleteConnectionDeprecatedOptions)
				completeConnectionDeprecatedOptionsModel.ID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserType = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.UserID = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Password = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.CertificateRoot = core.StringPtr("testString")
				completeConnectionDeprecatedOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.CompleteConnectionDeprecated(completeConnectionDeprecatedOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CompleteConnectionDeprecatedOptions model with no property values
				completeConnectionDeprecatedOptionsModelNew := new(ibmclouddatabasesv5.CompleteConnectionDeprecatedOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.CompleteConnectionDeprecated(completeConnectionDeprecatedOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetDeploymentScalingGroups(getDeploymentScalingGroupsOptions *GetDeploymentScalingGroupsOptions) - Operation response error`, func() {
		getDeploymentScalingGroupsPath := "/deployments/testString/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDeploymentScalingGroups with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentScalingGroupsOptions model
				getDeploymentScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentScalingGroupsOptions)
				getDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentScalingGroups(getDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentScalingGroups(getDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDeploymentScalingGroups(getDeploymentScalingGroupsOptions *GetDeploymentScalingGroupsOptions)`, func() {
		getDeploymentScalingGroupsPath := "/deployments/testString/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "2", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
				}))
			})
			It(`Invoke GetDeploymentScalingGroups successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDeploymentScalingGroupsOptions model
				getDeploymentScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentScalingGroupsOptions)
				getDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDeploymentScalingGroupsWithContext(ctx, getDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentScalingGroups(getDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDeploymentScalingGroupsWithContext(ctx, getDeploymentScalingGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDeploymentScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "2", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
				}))
			})
			It(`Invoke GetDeploymentScalingGroups successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentScalingGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDeploymentScalingGroupsOptions model
				getDeploymentScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentScalingGroupsOptions)
				getDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentScalingGroups(getDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDeploymentScalingGroups with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDeploymentScalingGroupsOptions model
				getDeploymentScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDeploymentScalingGroupsOptions)
				getDeploymentScalingGroupsOptionsModel.ID = core.StringPtr("testString")
				getDeploymentScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDeploymentScalingGroups(getDeploymentScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDeploymentScalingGroupsOptions model with no property values
				getDeploymentScalingGroupsOptionsModelNew := new(ibmclouddatabasesv5.GetDeploymentScalingGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetDeploymentScalingGroups(getDeploymentScalingGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions) - Operation response error`, func() {
		getDefaultScalingGroupsPath := "/deployables/postgresql/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDefaultScalingGroups with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDefaultScalingGroups(getDefaultScalingGroupsOptions *GetDefaultScalingGroupsOptions)`, func() {
		getDefaultScalingGroupsPath := "/deployables/postgresql/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "2", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
				}))
			})
			It(`Invoke GetDefaultScalingGroups successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetDefaultScalingGroupsWithContext(ctx, getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetDefaultScalingGroupsWithContext(ctx, getDefaultScalingGroupsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDefaultScalingGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"groups": [{"id": "member", "count": 2, "members": {"units": "count", "allocation_count": 2, "minimum_count": 2, "maximum_count": 20, "step_size_count": 1, "is_adjustable": true, "is_optional": false, "can_scale_down": false}, "memory": {"units": "mb", "allocation_mb": 12288, "minimum_mb": 1024, "maximum_mb": 114688, "step_size_mb": 1024, "is_adjustable": true, "is_optional": false, "can_scale_down": true}, "cpu": {"units": "2", "allocation_count": 2, "minimum_count": 2, "maximum_count": 32, "step_size_count": 2, "is_adjustable": false, "is_optional": false, "can_scale_down": true}, "disk": {"units": "mb", "allocation_mb": 10240, "minimum_mb": 2048, "maximum_mb": 4194304, "step_size_mb": 2048, "is_adjustable": true, "is_optional": false, "can_scale_down": false}}]}`)
				}))
			})
			It(`Invoke GetDefaultScalingGroups successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetDefaultScalingGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDefaultScalingGroups with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetDefaultScalingGroupsOptions model
				getDefaultScalingGroupsOptionsModel := new(ibmclouddatabasesv5.GetDefaultScalingGroupsOptions)
				getDefaultScalingGroupsOptionsModel.Type = core.StringPtr("postgresql")
				getDefaultScalingGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDefaultScalingGroupsOptions model with no property values
				getDefaultScalingGroupsOptionsModelNew := new(ibmclouddatabasesv5.GetDefaultScalingGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetDefaultScalingGroups(getDefaultScalingGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions) - Operation response error`, func() {
		setDeploymentScalingGroupPath := "/deployments/testString/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(ibmclouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetDeploymentScalingGroup(setDeploymentScalingGroupOptions *SetDeploymentScalingGroupOptions)`, func() {
		setDeploymentScalingGroupPath := "/deployments/testString/groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(ibmclouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.SetDeploymentScalingGroupWithContext(ctx, setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.SetDeploymentScalingGroupWithContext(ctx, setDeploymentScalingGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setDeploymentScalingGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetDeploymentScalingGroup successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.SetDeploymentScalingGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(ibmclouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetDeploymentScalingGroup with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(ibmclouddatabasesv5.SetMembersGroupMembers)
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				setDeploymentScalingGroupOptionsModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupOptions)
				setDeploymentScalingGroupOptionsModel.ID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.GroupID = core.StringPtr("testString")
				setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest = setDeploymentScalingGroupRequestModel
				setDeploymentScalingGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetDeploymentScalingGroupOptions model with no property values
				setDeploymentScalingGroupOptionsModelNew := new(ibmclouddatabasesv5.SetDeploymentScalingGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.SetDeploymentScalingGroup(setDeploymentScalingGroupOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions) - Operation response error`, func() {
		getAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAutoscalingConditions with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAutoscalingConditions(getAutoscalingConditionsOptions *GetAutoscalingConditionsOptions)`, func() {
		getAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"autoscaling": {"disk": {"scalers": {"capacity": {"enabled": true, "free_space_less_than_percent": 10}, "io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 20, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "memory": {"scalers": {"io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "cpu": {"scalers": {"anyKey": "anyValue"}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_count_per_member": 10, "units": "count"}}}}`)
				}))
			})
			It(`Invoke GetAutoscalingConditions successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetAutoscalingConditionsWithContext(ctx, getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetAutoscalingConditionsWithContext(ctx, getAutoscalingConditionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"autoscaling": {"disk": {"scalers": {"capacity": {"enabled": true, "free_space_less_than_percent": 10}, "io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 20, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "memory": {"scalers": {"io_utilization": {"enabled": true, "over_period": "30m", "above_percent": 45}}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_mb_per_member": 3670016, "units": "mb"}}, "cpu": {"scalers": {"anyKey": "anyValue"}, "rate": {"increase_percent": 10, "period_seconds": 900, "limit_count_per_member": 10, "units": "count"}}}}`)
				}))
			})
			It(`Invoke GetAutoscalingConditions successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetAutoscalingConditions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAutoscalingConditions with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetAutoscalingConditionsOptions model
				getAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.GetAutoscalingConditionsOptions)
				getAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				getAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAutoscalingConditionsOptions model with no property values
				getAutoscalingConditionsOptionsModelNew := new(ibmclouddatabasesv5.GetAutoscalingConditionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetAutoscalingConditions(getAutoscalingConditionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions) - Operation response error`, func() {
		setAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetAutoscalingConditions with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingDiskGroup model
				autoscalingSetGroupModel := new(ibmclouddatabasesv5.AutoscalingSetGroupAutoscalingDiskGroup)
				autoscalingSetGroupModel.Disk = autoscalingDiskGroupDiskModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetAutoscalingConditions(setAutoscalingConditionsOptions *SetAutoscalingConditionsOptions)`, func() {
		setAutoscalingConditionsPath := "/deployments/testString/groups/testString/autoscaling"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetAutoscalingConditions successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingDiskGroup model
				autoscalingSetGroupModel := new(ibmclouddatabasesv5.AutoscalingSetGroupAutoscalingDiskGroup)
				autoscalingSetGroupModel.Disk = autoscalingDiskGroupDiskModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.SetAutoscalingConditionsWithContext(ctx, setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.SetAutoscalingConditionsWithContext(ctx, setAutoscalingConditionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(setAutoscalingConditionsPath))
					Expect(req.Method).To(Equal("PATCH"))

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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke SetAutoscalingConditions successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.SetAutoscalingConditions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingDiskGroup model
				autoscalingSetGroupModel := new(ibmclouddatabasesv5.AutoscalingSetGroupAutoscalingDiskGroup)
				autoscalingSetGroupModel.Disk = autoscalingDiskGroupDiskModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SetAutoscalingConditions with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskRate)
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDisk)
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel

				// Construct an instance of the AutoscalingSetGroupAutoscalingDiskGroup model
				autoscalingSetGroupModel := new(ibmclouddatabasesv5.AutoscalingSetGroupAutoscalingDiskGroup)
				autoscalingSetGroupModel.Disk = autoscalingDiskGroupDiskModel

				// Construct an instance of the SetAutoscalingConditionsOptions model
				setAutoscalingConditionsOptionsModel := new(ibmclouddatabasesv5.SetAutoscalingConditionsOptions)
				setAutoscalingConditionsOptionsModel.ID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.GroupID = core.StringPtr("testString")
				setAutoscalingConditionsOptionsModel.Autoscaling = autoscalingSetGroupModel
				setAutoscalingConditionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SetAutoscalingConditionsOptions model with no property values
				setAutoscalingConditionsOptionsModelNew := new(ibmclouddatabasesv5.SetAutoscalingConditionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.SetAutoscalingConditions(setAutoscalingConditionsOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`KillConnections(killConnectionsOptions *KillConnectionsOptions) - Operation response error`, func() {
		killConnectionsPath := "/deployments/testString/management/database_connections"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke KillConnections with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(ibmclouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`KillConnections(killConnectionsOptions *KillConnectionsOptions)`, func() {
		killConnectionsPath := "/deployments/testString/management/database_connections"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke KillConnections successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(ibmclouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.KillConnectionsWithContext(ctx, killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.KillConnectionsWithContext(ctx, killConnectionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(killConnectionsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke KillConnections successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.KillConnections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(ibmclouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke KillConnections with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the KillConnectionsOptions model
				killConnectionsOptionsModel := new(ibmclouddatabasesv5.KillConnectionsOptions)
				killConnectionsOptionsModel.ID = core.StringPtr("testString")
				killConnectionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.KillConnections(killConnectionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the KillConnectionsOptions model with no property values
				killConnectionsOptionsModelNew := new(ibmclouddatabasesv5.KillConnectionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.KillConnections(killConnectionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`FileSync(fileSyncOptions *FileSyncOptions) - Operation response error`, func() {
		fileSyncPath := "/deployments/testString/elasticsearch/file_syncs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(fileSyncPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke FileSync with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(ibmclouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`FileSync(fileSyncOptions *FileSyncOptions)`, func() {
		fileSyncPath := "/deployments/testString/elasticsearch/file_syncs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(fileSyncPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke FileSync successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(ibmclouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.FileSyncWithContext(ctx, fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.FileSyncWithContext(ctx, fileSyncOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(fileSyncPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke FileSync successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.FileSync(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(ibmclouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke FileSync with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the FileSyncOptions model
				fileSyncOptionsModel := new(ibmclouddatabasesv5.FileSyncOptions)
				fileSyncOptionsModel.ID = core.StringPtr("testString")
				fileSyncOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.FileSync(fileSyncOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the FileSyncOptions model with no property values
				fileSyncOptionsModelNew := new(ibmclouddatabasesv5.FileSyncOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.FileSync(fileSyncOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions) - Operation response error`, func() {
		createLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(ibmclouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLogicalReplicationSlot(createLogicalReplicationSlotOptions *CreateLogicalReplicationSlotOptions)`, func() {
		createLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(ibmclouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.CreateLogicalReplicationSlotWithContext(ctx, createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.CreateLogicalReplicationSlotWithContext(ctx, createLogicalReplicationSlotOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createLogicalReplicationSlotPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke CreateLogicalReplicationSlot successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.CreateLogicalReplicationSlot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(ibmclouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLogicalReplicationSlot with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(ibmclouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				createLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.CreateLogicalReplicationSlotOptions)
				createLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot = logicalReplicationSlotLogicalReplicationSlotModel
				createLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateLogicalReplicationSlotOptions model with no property values
				createLogicalReplicationSlotOptionsModelNew := new(ibmclouddatabasesv5.CreateLogicalReplicationSlotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.CreateLogicalReplicationSlot(createLogicalReplicationSlotOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions) - Operation response error`, func() {
		deleteLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptions *DeleteLogicalReplicationSlotOptions)`, func() {
		deleteLogicalReplicationSlotPath := "/deployments/testString/postgresql/logical_replication_slots/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.DeleteLogicalReplicationSlotWithContext(ctx, deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.DeleteLogicalReplicationSlotWithContext(ctx, deleteLogicalReplicationSlotOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteLogicalReplicationSlotPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteLogicalReplicationSlot successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.DeleteLogicalReplicationSlot(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteLogicalReplicationSlot with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				deleteLogicalReplicationSlotOptionsModel := new(ibmclouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				deleteLogicalReplicationSlotOptionsModel.ID = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Name = core.StringPtr("testString")
				deleteLogicalReplicationSlotOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLogicalReplicationSlotOptions model with no property values
				deleteLogicalReplicationSlotOptionsModelNew := new(ibmclouddatabasesv5.DeleteLogicalReplicationSlotOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.DeleteLogicalReplicationSlot(deleteLogicalReplicationSlotOptionsModelNew)
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
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(ibmCloudDatabasesService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "https://ibmclouddatabasesv5/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(ibmCloudDatabasesService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL: "https://testService/api",
				})
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				})
				err := ibmCloudDatabasesService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := ibmCloudDatabasesService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != ibmCloudDatabasesService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(ibmCloudDatabasesService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(ibmCloudDatabasesService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_URL": "https://ibmclouddatabasesv5/api",
				"IBM_CLOUD_DATABASES_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IBM_CLOUD_DATABASES_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5UsingExternalConfig(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(ibmCloudDatabasesService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = ibmclouddatabasesv5.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetWhitelist(getWhitelistOptions *GetWhitelistOptions) - Operation response error`, func() {
		getWhitelistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWhitelistPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWhitelist with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetWhitelistOptions model
				getWhitelistOptionsModel := new(ibmclouddatabasesv5.GetWhitelistOptions)
				getWhitelistOptionsModel.ID = core.StringPtr("testString")
				getWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.GetWhitelist(getWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.GetWhitelist(getWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWhitelist(getWhitelistOptions *GetWhitelistOptions)`, func() {
		getWhitelistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWhitelistPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "Address", "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetWhitelist successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the GetWhitelistOptions model
				getWhitelistOptionsModel := new(ibmclouddatabasesv5.GetWhitelistOptions)
				getWhitelistOptionsModel.ID = core.StringPtr("testString")
				getWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.GetWhitelistWithContext(ctx, getWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.GetWhitelist(getWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.GetWhitelistWithContext(ctx, getWhitelistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getWhitelistPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ip_addresses": [{"address": "Address", "description": "Description"}]}`)
				}))
			})
			It(`Invoke GetWhitelist successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.GetWhitelist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWhitelistOptions model
				getWhitelistOptionsModel := new(ibmclouddatabasesv5.GetWhitelistOptions)
				getWhitelistOptionsModel.ID = core.StringPtr("testString")
				getWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.GetWhitelist(getWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetWhitelist with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the GetWhitelistOptions model
				getWhitelistOptionsModel := new(ibmclouddatabasesv5.GetWhitelistOptions)
				getWhitelistOptionsModel.ID = core.StringPtr("testString")
				getWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.GetWhitelist(getWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWhitelistOptions model with no property values
				getWhitelistOptionsModelNew := new(ibmclouddatabasesv5.GetWhitelistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.GetWhitelist(getWhitelistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceWhitelist(replaceWhitelistOptions *ReplaceWhitelistOptions) - Operation response error`, func() {
		replaceWhitelistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceWhitelistPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceWhitelist with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceWhitelistOptions model
				replaceWhitelistOptionsModel := new(ibmclouddatabasesv5.ReplaceWhitelistOptions)
				replaceWhitelistOptionsModel.ID = core.StringPtr("testString")
				replaceWhitelistOptionsModel.IpAddresses = []ibmclouddatabasesv5.WhitelistEntry{*whitelistEntryModel}
				replaceWhitelistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.ReplaceWhitelist(replaceWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.ReplaceWhitelist(replaceWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceWhitelist(replaceWhitelistOptions *ReplaceWhitelistOptions)`, func() {
		replaceWhitelistPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceWhitelistPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke ReplaceWhitelist successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceWhitelistOptions model
				replaceWhitelistOptionsModel := new(ibmclouddatabasesv5.ReplaceWhitelistOptions)
				replaceWhitelistOptionsModel.ID = core.StringPtr("testString")
				replaceWhitelistOptionsModel.IpAddresses = []ibmclouddatabasesv5.WhitelistEntry{*whitelistEntryModel}
				replaceWhitelistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.ReplaceWhitelistWithContext(ctx, replaceWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.ReplaceWhitelist(replaceWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.ReplaceWhitelistWithContext(ctx, replaceWhitelistOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceWhitelistPath))
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

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke ReplaceWhitelist successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.ReplaceWhitelist(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceWhitelistOptions model
				replaceWhitelistOptionsModel := new(ibmclouddatabasesv5.ReplaceWhitelistOptions)
				replaceWhitelistOptionsModel.ID = core.StringPtr("testString")
				replaceWhitelistOptionsModel.IpAddresses = []ibmclouddatabasesv5.WhitelistEntry{*whitelistEntryModel}
				replaceWhitelistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.ReplaceWhitelist(replaceWhitelistOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceWhitelist with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the ReplaceWhitelistOptions model
				replaceWhitelistOptionsModel := new(ibmclouddatabasesv5.ReplaceWhitelistOptions)
				replaceWhitelistOptionsModel.ID = core.StringPtr("testString")
				replaceWhitelistOptionsModel.IpAddresses = []ibmclouddatabasesv5.WhitelistEntry{*whitelistEntryModel}
				replaceWhitelistOptionsModel.IfMatch = core.StringPtr("testString")
				replaceWhitelistOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.ReplaceWhitelist(replaceWhitelistOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceWhitelistOptions model with no property values
				replaceWhitelistOptionsModelNew := new(ibmclouddatabasesv5.ReplaceWhitelistOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.ReplaceWhitelist(replaceWhitelistOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddWhitelistEntry(addWhitelistEntryOptions *AddWhitelistEntryOptions) - Operation response error`, func() {
		addWhitelistEntryPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWhitelistEntryPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddWhitelistEntry with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the AddWhitelistEntryOptions model
				addWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.AddWhitelistEntryOptions)
				addWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				addWhitelistEntryOptionsModel.IpAddress = whitelistEntryModel
				addWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.AddWhitelistEntry(addWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.AddWhitelistEntry(addWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddWhitelistEntry(addWhitelistEntryOptions *AddWhitelistEntryOptions)`, func() {
		addWhitelistEntryPath := "/deployments/testString/whitelists/ip_addresses"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWhitelistEntryPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke AddWhitelistEntry successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the AddWhitelistEntryOptions model
				addWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.AddWhitelistEntryOptions)
				addWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				addWhitelistEntryOptionsModel.IpAddress = whitelistEntryModel
				addWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.AddWhitelistEntryWithContext(ctx, addWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.AddWhitelistEntry(addWhitelistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.AddWhitelistEntryWithContext(ctx, addWhitelistEntryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(addWhitelistEntryPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke AddWhitelistEntry successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.AddWhitelistEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the AddWhitelistEntryOptions model
				addWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.AddWhitelistEntryOptions)
				addWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				addWhitelistEntryOptionsModel.IpAddress = whitelistEntryModel
				addWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.AddWhitelistEntry(addWhitelistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddWhitelistEntry with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")

				// Construct an instance of the AddWhitelistEntryOptions model
				addWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.AddWhitelistEntryOptions)
				addWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				addWhitelistEntryOptionsModel.IpAddress = whitelistEntryModel
				addWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.AddWhitelistEntry(addWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddWhitelistEntryOptions model with no property values
				addWhitelistEntryOptionsModelNew := new(ibmclouddatabasesv5.AddWhitelistEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.AddWhitelistEntry(addWhitelistEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWhitelistEntry(deleteWhitelistEntryOptions *DeleteWhitelistEntryOptions) - Operation response error`, func() {
		deleteWhitelistEntryPath := "/deployments/testString/whitelists/ip_addresses/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWhitelistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteWhitelistEntry with error: Operation response processing error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteWhitelistEntryOptions model
				deleteWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.DeleteWhitelistEntryOptions)
				deleteWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := ibmCloudDatabasesService.DeleteWhitelistEntry(deleteWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				ibmCloudDatabasesService.EnableRetries(0, 0)
				result, response, operationErr = ibmCloudDatabasesService.DeleteWhitelistEntry(deleteWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteWhitelistEntry(deleteWhitelistEntryOptions *DeleteWhitelistEntryOptions)`, func() {
		deleteWhitelistEntryPath := "/deployments/testString/whitelists/ip_addresses/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWhitelistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteWhitelistEntry successfully with retries`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())
				ibmCloudDatabasesService.EnableRetries(0, 0)

				// Construct an instance of the DeleteWhitelistEntryOptions model
				deleteWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.DeleteWhitelistEntryOptions)
				deleteWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := ibmCloudDatabasesService.DeleteWhitelistEntryWithContext(ctx, deleteWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				ibmCloudDatabasesService.DisableRetries()
				result, response, operationErr := ibmCloudDatabasesService.DeleteWhitelistEntry(deleteWhitelistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = ibmCloudDatabasesService.DeleteWhitelistEntryWithContext(ctx, deleteWhitelistEntryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteWhitelistEntryPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task": {"id": "ID", "description": "Description", "status": "running", "deployment_id": "DeploymentID", "progress_percent": 15, "created_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke DeleteWhitelistEntry successfully`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := ibmCloudDatabasesService.DeleteWhitelistEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteWhitelistEntryOptions model
				deleteWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.DeleteWhitelistEntryOptions)
				deleteWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = ibmCloudDatabasesService.DeleteWhitelistEntry(deleteWhitelistEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteWhitelistEntry with error: Operation validation and request error`, func() {
				ibmCloudDatabasesService, serviceErr := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(ibmCloudDatabasesService).ToNot(BeNil())

				// Construct an instance of the DeleteWhitelistEntryOptions model
				deleteWhitelistEntryOptionsModel := new(ibmclouddatabasesv5.DeleteWhitelistEntryOptions)
				deleteWhitelistEntryOptionsModel.ID = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Ipaddress = core.StringPtr("testString")
				deleteWhitelistEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := ibmCloudDatabasesService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := ibmCloudDatabasesService.DeleteWhitelistEntry(deleteWhitelistEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteWhitelistEntryOptions model with no property values
				deleteWhitelistEntryOptionsModelNew := new(ibmclouddatabasesv5.DeleteWhitelistEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = ibmCloudDatabasesService.DeleteWhitelistEntry(deleteWhitelistEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			ibmCloudDatabasesService, _ := ibmclouddatabasesv5.NewIbmCloudDatabasesV5(&ibmclouddatabasesv5.IbmCloudDatabasesV5Options{
				URL:           "http://ibmclouddatabasesv5modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddWhitelistEntryOptions successfully`, func() {
				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				Expect(whitelistEntryModel).ToNot(BeNil())
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")
				Expect(whitelistEntryModel.Address).To(Equal(core.StringPtr("testString")))
				Expect(whitelistEntryModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddWhitelistEntryOptions model
				id := "testString"
				addWhitelistEntryOptionsModel := ibmCloudDatabasesService.NewAddWhitelistEntryOptions(id)
				addWhitelistEntryOptionsModel.SetID("testString")
				addWhitelistEntryOptionsModel.SetIpAddress(whitelistEntryModel)
				addWhitelistEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWhitelistEntryOptionsModel).ToNot(BeNil())
				Expect(addWhitelistEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(addWhitelistEntryOptionsModel.IpAddress).To(Equal(whitelistEntryModel))
				Expect(addWhitelistEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewChangeUserPasswordOptions successfully`, func() {
				// Construct an instance of the APasswordSettingUser model
				aPasswordSettingUserModel := new(ibmclouddatabasesv5.APasswordSettingUser)
				Expect(aPasswordSettingUserModel).ToNot(BeNil())
				aPasswordSettingUserModel.Password = core.StringPtr("xyzzy")
				Expect(aPasswordSettingUserModel.Password).To(Equal(core.StringPtr("xyzzy")))

				// Construct an instance of the ChangeUserPasswordOptions model
				id := "testString"
				userType := "testString"
				username := "testString"
				changeUserPasswordOptionsModel := ibmCloudDatabasesService.NewChangeUserPasswordOptions(id, userType, username)
				changeUserPasswordOptionsModel.SetID("testString")
				changeUserPasswordOptionsModel.SetUserType("testString")
				changeUserPasswordOptionsModel.SetUsername("testString")
				changeUserPasswordOptionsModel.SetUser(aPasswordSettingUserModel)
				changeUserPasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeUserPasswordOptionsModel).ToNot(BeNil())
				Expect(changeUserPasswordOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(changeUserPasswordOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(changeUserPasswordOptionsModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(changeUserPasswordOptionsModel.User).To(Equal(aPasswordSettingUserModel))
				Expect(changeUserPasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCompleteConnectionDeprecatedOptions successfully`, func() {
				// Construct an instance of the CompleteConnectionDeprecatedOptions model
				id := "testString"
				userType := "testString"
				userID := "testString"
				completeConnectionDeprecatedOptionsModel := ibmCloudDatabasesService.NewCompleteConnectionDeprecatedOptions(id, userType, userID)
				completeConnectionDeprecatedOptionsModel.SetID("testString")
				completeConnectionDeprecatedOptionsModel.SetUserType("testString")
				completeConnectionDeprecatedOptionsModel.SetUserID("testString")
				completeConnectionDeprecatedOptionsModel.SetPassword("testString")
				completeConnectionDeprecatedOptionsModel.SetCertificateRoot("testString")
				completeConnectionDeprecatedOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeConnectionDeprecatedOptionsModel).ToNot(BeNil())
				Expect(completeConnectionDeprecatedOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionDeprecatedOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionDeprecatedOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionDeprecatedOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionDeprecatedOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionDeprecatedOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCompleteConnectionOptions successfully`, func() {
				// Construct an instance of the CompleteConnectionOptions model
				id := "testString"
				userType := "testString"
				userID := "testString"
				endpointType := "public"
				completeConnectionOptionsModel := ibmCloudDatabasesService.NewCompleteConnectionOptions(id, userType, userID, endpointType)
				completeConnectionOptionsModel.SetID("testString")
				completeConnectionOptionsModel.SetUserType("testString")
				completeConnectionOptionsModel.SetUserID("testString")
				completeConnectionOptionsModel.SetEndpointType("public")
				completeConnectionOptionsModel.SetPassword("testString")
				completeConnectionOptionsModel.SetCertificateRoot("testString")
				completeConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(completeConnectionOptionsModel).ToNot(BeNil())
				Expect(completeConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(completeConnectionOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(completeConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDatabaseUserOptions successfully`, func() {
				// Construct an instance of the CreateDatabaseUserRequestUser model
				createDatabaseUserRequestUserModel := new(ibmclouddatabasesv5.CreateDatabaseUserRequestUser)
				Expect(createDatabaseUserRequestUserModel).ToNot(BeNil())
				createDatabaseUserRequestUserModel.UserType = core.StringPtr("database")
				createDatabaseUserRequestUserModel.Username = core.StringPtr("james")
				createDatabaseUserRequestUserModel.Password = core.StringPtr("kickoutthe")
				Expect(createDatabaseUserRequestUserModel.UserType).To(Equal(core.StringPtr("database")))
				Expect(createDatabaseUserRequestUserModel.Username).To(Equal(core.StringPtr("james")))
				Expect(createDatabaseUserRequestUserModel.Password).To(Equal(core.StringPtr("kickoutthe")))

				// Construct an instance of the CreateDatabaseUserOptions model
				id := "testString"
				userType := "testString"
				createDatabaseUserOptionsModel := ibmCloudDatabasesService.NewCreateDatabaseUserOptions(id, userType)
				createDatabaseUserOptionsModel.SetID("testString")
				createDatabaseUserOptionsModel.SetUserType("testString")
				createDatabaseUserOptionsModel.SetUser(createDatabaseUserRequestUserModel)
				createDatabaseUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDatabaseUserOptionsModel).ToNot(BeNil())
				Expect(createDatabaseUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseUserOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(createDatabaseUserOptionsModel.User).To(Equal(createDatabaseUserRequestUserModel))
				Expect(createDatabaseUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLogicalReplicationSlotOptions successfully`, func() {
				// Construct an instance of the LogicalReplicationSlotLogicalReplicationSlot model
				logicalReplicationSlotLogicalReplicationSlotModel := new(ibmclouddatabasesv5.LogicalReplicationSlotLogicalReplicationSlot)
				Expect(logicalReplicationSlotLogicalReplicationSlotModel).ToNot(BeNil())
				logicalReplicationSlotLogicalReplicationSlotModel.Name = core.StringPtr("customer_replication")
				logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName = core.StringPtr("customers")
				logicalReplicationSlotLogicalReplicationSlotModel.PluginType = core.StringPtr("wal2json")
				Expect(logicalReplicationSlotLogicalReplicationSlotModel.Name).To(Equal(core.StringPtr("customer_replication")))
				Expect(logicalReplicationSlotLogicalReplicationSlotModel.DatabaseName).To(Equal(core.StringPtr("customers")))
				Expect(logicalReplicationSlotLogicalReplicationSlotModel.PluginType).To(Equal(core.StringPtr("wal2json")))

				// Construct an instance of the CreateLogicalReplicationSlotOptions model
				id := "testString"
				createLogicalReplicationSlotOptionsModel := ibmCloudDatabasesService.NewCreateLogicalReplicationSlotOptions(id)
				createLogicalReplicationSlotOptionsModel.SetID("testString")
				createLogicalReplicationSlotOptionsModel.SetLogicalReplicationSlot(logicalReplicationSlotLogicalReplicationSlotModel)
				createLogicalReplicationSlotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLogicalReplicationSlotOptionsModel).ToNot(BeNil())
				Expect(createLogicalReplicationSlotOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createLogicalReplicationSlotOptionsModel.LogicalReplicationSlot).To(Equal(logicalReplicationSlotLogicalReplicationSlotModel))
				Expect(createLogicalReplicationSlotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDatabaseUserOptions successfully`, func() {
				// Construct an instance of the DeleteDatabaseUserOptions model
				id := "testString"
				userType := "testString"
				username := "testString"
				deleteDatabaseUserOptionsModel := ibmCloudDatabasesService.NewDeleteDatabaseUserOptions(id, userType, username)
				deleteDatabaseUserOptionsModel.SetID("testString")
				deleteDatabaseUserOptionsModel.SetUserType("testString")
				deleteDatabaseUserOptionsModel.SetUsername("testString")
				deleteDatabaseUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDatabaseUserOptionsModel).ToNot(BeNil())
				Expect(deleteDatabaseUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseUserOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseUserOptionsModel.Username).To(Equal(core.StringPtr("testString")))
				Expect(deleteDatabaseUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLogicalReplicationSlotOptions successfully`, func() {
				// Construct an instance of the DeleteLogicalReplicationSlotOptions model
				id := "testString"
				name := "testString"
				deleteLogicalReplicationSlotOptionsModel := ibmCloudDatabasesService.NewDeleteLogicalReplicationSlotOptions(id, name)
				deleteLogicalReplicationSlotOptionsModel.SetID("testString")
				deleteLogicalReplicationSlotOptionsModel.SetName("testString")
				deleteLogicalReplicationSlotOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLogicalReplicationSlotOptionsModel).ToNot(BeNil())
				Expect(deleteLogicalReplicationSlotOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLogicalReplicationSlotOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(deleteLogicalReplicationSlotOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWhitelistEntryOptions successfully`, func() {
				// Construct an instance of the DeleteWhitelistEntryOptions model
				id := "testString"
				ipaddress := "testString"
				deleteWhitelistEntryOptionsModel := ibmCloudDatabasesService.NewDeleteWhitelistEntryOptions(id, ipaddress)
				deleteWhitelistEntryOptionsModel.SetID("testString")
				deleteWhitelistEntryOptionsModel.SetIpaddress("testString")
				deleteWhitelistEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWhitelistEntryOptionsModel).ToNot(BeNil())
				Expect(deleteWhitelistEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWhitelistEntryOptionsModel.Ipaddress).To(Equal(core.StringPtr("testString")))
				Expect(deleteWhitelistEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFileSyncOptions successfully`, func() {
				// Construct an instance of the FileSyncOptions model
				id := "testString"
				fileSyncOptionsModel := ibmCloudDatabasesService.NewFileSyncOptions(id)
				fileSyncOptionsModel.SetID("testString")
				fileSyncOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(fileSyncOptionsModel).ToNot(BeNil())
				Expect(fileSyncOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(fileSyncOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAutoscalingConditionsOptions successfully`, func() {
				// Construct an instance of the GetAutoscalingConditionsOptions model
				id := "testString"
				groupID := "testString"
				getAutoscalingConditionsOptionsModel := ibmCloudDatabasesService.NewGetAutoscalingConditionsOptions(id, groupID)
				getAutoscalingConditionsOptionsModel.SetID("testString")
				getAutoscalingConditionsOptionsModel.SetGroupID("testString")
				getAutoscalingConditionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAutoscalingConditionsOptionsModel).ToNot(BeNil())
				Expect(getAutoscalingConditionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAutoscalingConditionsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAutoscalingConditionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBackupInfoOptions successfully`, func() {
				// Construct an instance of the GetBackupInfoOptions model
				backupID := "testString"
				getBackupInfoOptionsModel := ibmCloudDatabasesService.NewGetBackupInfoOptions(backupID)
				getBackupInfoOptionsModel.SetBackupID("testString")
				getBackupInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBackupInfoOptionsModel).ToNot(BeNil())
				Expect(getBackupInfoOptionsModel.BackupID).To(Equal(core.StringPtr("testString")))
				Expect(getBackupInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConnectionDeprecatedOptions successfully`, func() {
				// Construct an instance of the GetConnectionDeprecatedOptions model
				id := "testString"
				userType := "testString"
				userID := "testString"
				getConnectionDeprecatedOptionsModel := ibmCloudDatabasesService.NewGetConnectionDeprecatedOptions(id, userType, userID)
				getConnectionDeprecatedOptionsModel.SetID("testString")
				getConnectionDeprecatedOptionsModel.SetUserType("testString")
				getConnectionDeprecatedOptionsModel.SetUserID("testString")
				getConnectionDeprecatedOptionsModel.SetCertificateRoot("testString")
				getConnectionDeprecatedOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConnectionDeprecatedOptionsModel).ToNot(BeNil())
				Expect(getConnectionDeprecatedOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionDeprecatedOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionDeprecatedOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionDeprecatedOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionDeprecatedOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConnectionOptions successfully`, func() {
				// Construct an instance of the GetConnectionOptions model
				id := "testString"
				userType := "testString"
				userID := "testString"
				endpointType := "public"
				getConnectionOptionsModel := ibmCloudDatabasesService.NewGetConnectionOptions(id, userType, userID, endpointType)
				getConnectionOptionsModel.SetID("testString")
				getConnectionOptionsModel.SetUserType("testString")
				getConnectionOptionsModel.SetUserID("testString")
				getConnectionOptionsModel.SetEndpointType("public")
				getConnectionOptionsModel.SetCertificateRoot("testString")
				getConnectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConnectionOptionsModel).ToNot(BeNil())
				Expect(getConnectionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.UserType).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(getConnectionOptionsModel.CertificateRoot).To(Equal(core.StringPtr("testString")))
				Expect(getConnectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDatabaseConfigurationSchemaOptions successfully`, func() {
				// Construct an instance of the GetDatabaseConfigurationSchemaOptions model
				id := "testString"
				getDatabaseConfigurationSchemaOptionsModel := ibmCloudDatabasesService.NewGetDatabaseConfigurationSchemaOptions(id)
				getDatabaseConfigurationSchemaOptionsModel.SetID("testString")
				getDatabaseConfigurationSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDatabaseConfigurationSchemaOptionsModel).ToNot(BeNil())
				Expect(getDatabaseConfigurationSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDatabaseConfigurationSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDefaultScalingGroupsOptions successfully`, func() {
				// Construct an instance of the GetDefaultScalingGroupsOptions model
				typeVar := "postgresql"
				getDefaultScalingGroupsOptionsModel := ibmCloudDatabasesService.NewGetDefaultScalingGroupsOptions(typeVar)
				getDefaultScalingGroupsOptionsModel.SetType("postgresql")
				getDefaultScalingGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDefaultScalingGroupsOptionsModel).ToNot(BeNil())
				Expect(getDefaultScalingGroupsOptionsModel.Type).To(Equal(core.StringPtr("postgresql")))
				Expect(getDefaultScalingGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeployablesOptions successfully`, func() {
				// Construct an instance of the GetDeployablesOptions model
				getDeployablesOptionsModel := ibmCloudDatabasesService.NewGetDeployablesOptions()
				getDeployablesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeployablesOptionsModel).ToNot(BeNil())
				Expect(getDeployablesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentBackupsOptions successfully`, func() {
				// Construct an instance of the GetDeploymentBackupsOptions model
				id := "testString"
				getDeploymentBackupsOptionsModel := ibmCloudDatabasesService.NewGetDeploymentBackupsOptions(id)
				getDeploymentBackupsOptionsModel.SetID("testString")
				getDeploymentBackupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentBackupsOptionsModel).ToNot(BeNil())
				Expect(getDeploymentBackupsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentBackupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentInfoOptions successfully`, func() {
				// Construct an instance of the GetDeploymentInfoOptions model
				id := "testString"
				getDeploymentInfoOptionsModel := ibmCloudDatabasesService.NewGetDeploymentInfoOptions(id)
				getDeploymentInfoOptionsModel.SetID("testString")
				getDeploymentInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentInfoOptionsModel).ToNot(BeNil())
				Expect(getDeploymentInfoOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentScalingGroupsOptions successfully`, func() {
				// Construct an instance of the GetDeploymentScalingGroupsOptions model
				id := "testString"
				getDeploymentScalingGroupsOptionsModel := ibmCloudDatabasesService.NewGetDeploymentScalingGroupsOptions(id)
				getDeploymentScalingGroupsOptionsModel.SetID("testString")
				getDeploymentScalingGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentScalingGroupsOptionsModel).ToNot(BeNil())
				Expect(getDeploymentScalingGroupsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentScalingGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDeploymentTasksOptions successfully`, func() {
				// Construct an instance of the GetDeploymentTasksOptions model
				id := "testString"
				getDeploymentTasksOptionsModel := ibmCloudDatabasesService.NewGetDeploymentTasksOptions(id)
				getDeploymentTasksOptionsModel.SetID("testString")
				getDeploymentTasksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDeploymentTasksOptionsModel).ToNot(BeNil())
				Expect(getDeploymentTasksOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDeploymentTasksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPITRdataOptions successfully`, func() {
				// Construct an instance of the GetPITRdataOptions model
				id := "testString"
				getPitRdataOptionsModel := ibmCloudDatabasesService.NewGetPITRdataOptions(id)
				getPitRdataOptionsModel.SetID("testString")
				getPitRdataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPitRdataOptionsModel).ToNot(BeNil())
				Expect(getPitRdataOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPitRdataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRegionsOptions successfully`, func() {
				// Construct an instance of the GetRegionsOptions model
				getRegionsOptionsModel := ibmCloudDatabasesService.NewGetRegionsOptions()
				getRegionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRegionsOptionsModel).ToNot(BeNil())
				Expect(getRegionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRemotesOptions successfully`, func() {
				// Construct an instance of the GetRemotesOptions model
				id := "testString"
				getRemotesOptionsModel := ibmCloudDatabasesService.NewGetRemotesOptions(id)
				getRemotesOptionsModel.SetID("testString")
				getRemotesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRemotesOptionsModel).ToNot(BeNil())
				Expect(getRemotesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRemotesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRemotesSchemaOptions successfully`, func() {
				// Construct an instance of the GetRemotesSchemaOptions model
				id := "testString"
				getRemotesSchemaOptionsModel := ibmCloudDatabasesService.NewGetRemotesSchemaOptions(id)
				getRemotesSchemaOptionsModel.SetID("testString")
				getRemotesSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRemotesSchemaOptionsModel).ToNot(BeNil())
				Expect(getRemotesSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRemotesSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTasksOptions successfully`, func() {
				// Construct an instance of the GetTasksOptions model
				id := "testString"
				getTasksOptionsModel := ibmCloudDatabasesService.NewGetTasksOptions(id)
				getTasksOptionsModel.SetID("testString")
				getTasksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTasksOptionsModel).ToNot(BeNil())
				Expect(getTasksOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTasksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserOptions successfully`, func() {
				// Construct an instance of the GetUserOptions model
				id := "testString"
				userID := "testString"
				endpointType := "public"
				getUserOptionsModel := ibmCloudDatabasesService.NewGetUserOptions(id, userID, endpointType)
				getUserOptionsModel.SetID("testString")
				getUserOptionsModel.SetUserID("testString")
				getUserOptionsModel.SetEndpointType("public")
				getUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserOptionsModel).ToNot(BeNil())
				Expect(getUserOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getUserOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(getUserOptionsModel.EndpointType).To(Equal(core.StringPtr("public")))
				Expect(getUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWhitelistOptions successfully`, func() {
				// Construct an instance of the GetWhitelistOptions model
				id := "testString"
				getWhitelistOptionsModel := ibmCloudDatabasesService.NewGetWhitelistOptions(id)
				getWhitelistOptionsModel.SetID("testString")
				getWhitelistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWhitelistOptionsModel).ToNot(BeNil())
				Expect(getWhitelistOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getWhitelistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewKillConnectionsOptions successfully`, func() {
				// Construct an instance of the KillConnectionsOptions model
				id := "testString"
				killConnectionsOptionsModel := ibmCloudDatabasesService.NewKillConnectionsOptions(id)
				killConnectionsOptionsModel.SetID("testString")
				killConnectionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(killConnectionsOptionsModel).ToNot(BeNil())
				Expect(killConnectionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(killConnectionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceWhitelistOptions successfully`, func() {
				// Construct an instance of the WhitelistEntry model
				whitelistEntryModel := new(ibmclouddatabasesv5.WhitelistEntry)
				Expect(whitelistEntryModel).ToNot(BeNil())
				whitelistEntryModel.Address = core.StringPtr("testString")
				whitelistEntryModel.Description = core.StringPtr("testString")
				Expect(whitelistEntryModel.Address).To(Equal(core.StringPtr("testString")))
				Expect(whitelistEntryModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceWhitelistOptions model
				id := "testString"
				replaceWhitelistOptionsModel := ibmCloudDatabasesService.NewReplaceWhitelistOptions(id)
				replaceWhitelistOptionsModel.SetID("testString")
				replaceWhitelistOptionsModel.SetIpAddresses([]ibmclouddatabasesv5.WhitelistEntry{*whitelistEntryModel})
				replaceWhitelistOptionsModel.SetIfMatch("testString")
				replaceWhitelistOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceWhitelistOptionsModel).ToNot(BeNil())
				Expect(replaceWhitelistOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceWhitelistOptionsModel.IpAddresses).To(Equal([]ibmclouddatabasesv5.WhitelistEntry{*whitelistEntryModel}))
				Expect(replaceWhitelistOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceWhitelistOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetAutoscalingConditionsOptions successfully`, func() {
				// Construct an instance of the AutoscalingDiskGroupDiskScalersCapacity model
				autoscalingDiskGroupDiskScalersCapacityModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersCapacity)
				Expect(autoscalingDiskGroupDiskScalersCapacityModel).ToNot(BeNil())
				autoscalingDiskGroupDiskScalersCapacityModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent = core.Int64Ptr(int64(10))
				Expect(autoscalingDiskGroupDiskScalersCapacityModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(autoscalingDiskGroupDiskScalersCapacityModel.FreeSpaceLessThanPercent).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the AutoscalingDiskGroupDiskScalersIoUtilization model
				autoscalingDiskGroupDiskScalersIoUtilizationModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalersIoUtilization)
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel).ToNot(BeNil())
				autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled = core.BoolPtr(true)
				autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod = core.StringPtr("30m")
				autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent = core.Int64Ptr(int64(45))
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel.OverPeriod).To(Equal(core.StringPtr("30m")))
				Expect(autoscalingDiskGroupDiskScalersIoUtilizationModel.AbovePercent).To(Equal(core.Int64Ptr(int64(45))))

				// Construct an instance of the AutoscalingDiskGroupDiskScalers model
				autoscalingDiskGroupDiskScalersModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskScalers)
				Expect(autoscalingDiskGroupDiskScalersModel).ToNot(BeNil())
				autoscalingDiskGroupDiskScalersModel.Capacity = autoscalingDiskGroupDiskScalersCapacityModel
				autoscalingDiskGroupDiskScalersModel.IoUtilization = autoscalingDiskGroupDiskScalersIoUtilizationModel
				Expect(autoscalingDiskGroupDiskScalersModel.Capacity).To(Equal(autoscalingDiskGroupDiskScalersCapacityModel))
				Expect(autoscalingDiskGroupDiskScalersModel.IoUtilization).To(Equal(autoscalingDiskGroupDiskScalersIoUtilizationModel))

				// Construct an instance of the AutoscalingDiskGroupDiskRate model
				autoscalingDiskGroupDiskRateModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDiskRate)
				Expect(autoscalingDiskGroupDiskRateModel).ToNot(BeNil())
				autoscalingDiskGroupDiskRateModel.IncreasePercent = core.Float64Ptr(float64(20))
				autoscalingDiskGroupDiskRateModel.PeriodSeconds = core.Int64Ptr(int64(900))
				autoscalingDiskGroupDiskRateModel.LimitMbPerMember = core.Float64Ptr(float64(3670016))
				autoscalingDiskGroupDiskRateModel.Units = core.StringPtr("mb")
				Expect(autoscalingDiskGroupDiskRateModel.IncreasePercent).To(Equal(core.Float64Ptr(float64(20))))
				Expect(autoscalingDiskGroupDiskRateModel.PeriodSeconds).To(Equal(core.Int64Ptr(int64(900))))
				Expect(autoscalingDiskGroupDiskRateModel.LimitMbPerMember).To(Equal(core.Float64Ptr(float64(3670016))))
				Expect(autoscalingDiskGroupDiskRateModel.Units).To(Equal(core.StringPtr("mb")))

				// Construct an instance of the AutoscalingDiskGroupDisk model
				autoscalingDiskGroupDiskModel := new(ibmclouddatabasesv5.AutoscalingDiskGroupDisk)
				Expect(autoscalingDiskGroupDiskModel).ToNot(BeNil())
				autoscalingDiskGroupDiskModel.Scalers = autoscalingDiskGroupDiskScalersModel
				autoscalingDiskGroupDiskModel.Rate = autoscalingDiskGroupDiskRateModel
				Expect(autoscalingDiskGroupDiskModel.Scalers).To(Equal(autoscalingDiskGroupDiskScalersModel))
				Expect(autoscalingDiskGroupDiskModel.Rate).To(Equal(autoscalingDiskGroupDiskRateModel))

				// Construct an instance of the AutoscalingSetGroupAutoscalingDiskGroup model
				autoscalingSetGroupModel := new(ibmclouddatabasesv5.AutoscalingSetGroupAutoscalingDiskGroup)
				Expect(autoscalingSetGroupModel).ToNot(BeNil())
				autoscalingSetGroupModel.Disk = autoscalingDiskGroupDiskModel
				Expect(autoscalingSetGroupModel.Disk).To(Equal(autoscalingDiskGroupDiskModel))

				// Construct an instance of the SetAutoscalingConditionsOptions model
				id := "testString"
				groupID := "testString"
				setAutoscalingConditionsOptionsModel := ibmCloudDatabasesService.NewSetAutoscalingConditionsOptions(id, groupID)
				setAutoscalingConditionsOptionsModel.SetID("testString")
				setAutoscalingConditionsOptionsModel.SetGroupID("testString")
				setAutoscalingConditionsOptionsModel.SetAutoscaling(autoscalingSetGroupModel)
				setAutoscalingConditionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setAutoscalingConditionsOptionsModel).ToNot(BeNil())
				Expect(setAutoscalingConditionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setAutoscalingConditionsOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(setAutoscalingConditionsOptionsModel.Autoscaling).To(Equal(autoscalingSetGroupModel))
				Expect(setAutoscalingConditionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetDatabaseConfigurationOptions successfully`, func() {
				// Construct an instance of the SetConfigurationConfigurationPGConfiguration model
				setConfigurationConfigurationModel := new(ibmclouddatabasesv5.SetConfigurationConfigurationPGConfiguration)
				Expect(setConfigurationConfigurationModel).ToNot(BeNil())
				setConfigurationConfigurationModel.MaxConnections = core.Int64Ptr(int64(115))
				setConfigurationConfigurationModel.MaxPreparedTransactions = core.Int64Ptr(int64(0))
				setConfigurationConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(100))
				setConfigurationConfigurationModel.EffectiveIoConcurrency = core.Int64Ptr(int64(1))
				setConfigurationConfigurationModel.MaxReplicationSlots = core.Int64Ptr(int64(10))
				setConfigurationConfigurationModel.MaxWalSenders = core.Int64Ptr(int64(12))
				setConfigurationConfigurationModel.SharedBuffers = core.Int64Ptr(int64(16))
				setConfigurationConfigurationModel.SynchronousCommit = core.StringPtr("local")
				setConfigurationConfigurationModel.WalLevel = core.StringPtr("hot_standby")
				setConfigurationConfigurationModel.ArchiveTimeout = core.Int64Ptr(int64(300))
				setConfigurationConfigurationModel.LogMinDurationStatement = core.Int64Ptr(int64(100))
				Expect(setConfigurationConfigurationModel.MaxConnections).To(Equal(core.Int64Ptr(int64(115))))
				Expect(setConfigurationConfigurationModel.MaxPreparedTransactions).To(Equal(core.Int64Ptr(int64(0))))
				Expect(setConfigurationConfigurationModel.DeadlockTimeout).To(Equal(core.Int64Ptr(int64(100))))
				Expect(setConfigurationConfigurationModel.EffectiveIoConcurrency).To(Equal(core.Int64Ptr(int64(1))))
				Expect(setConfigurationConfigurationModel.MaxReplicationSlots).To(Equal(core.Int64Ptr(int64(10))))
				Expect(setConfigurationConfigurationModel.MaxWalSenders).To(Equal(core.Int64Ptr(int64(12))))
				Expect(setConfigurationConfigurationModel.SharedBuffers).To(Equal(core.Int64Ptr(int64(16))))
				Expect(setConfigurationConfigurationModel.SynchronousCommit).To(Equal(core.StringPtr("local")))
				Expect(setConfigurationConfigurationModel.WalLevel).To(Equal(core.StringPtr("hot_standby")))
				Expect(setConfigurationConfigurationModel.ArchiveTimeout).To(Equal(core.Int64Ptr(int64(300))))
				Expect(setConfigurationConfigurationModel.LogMinDurationStatement).To(Equal(core.Int64Ptr(int64(100))))

				// Construct an instance of the SetDatabaseConfigurationOptions model
				id := "testString"
				var setDatabaseConfigurationOptionsConfiguration ibmclouddatabasesv5.SetConfigurationConfigurationIntf = nil
				setDatabaseConfigurationOptionsModel := ibmCloudDatabasesService.NewSetDatabaseConfigurationOptions(id, setDatabaseConfigurationOptionsConfiguration)
				setDatabaseConfigurationOptionsModel.SetID("testString")
				setDatabaseConfigurationOptionsModel.SetConfiguration(setConfigurationConfigurationModel)
				setDatabaseConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setDatabaseConfigurationOptionsModel).ToNot(BeNil())
				Expect(setDatabaseConfigurationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setDatabaseConfigurationOptionsModel.Configuration).To(Equal(setConfigurationConfigurationModel))
				Expect(setDatabaseConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetDeploymentScalingGroupOptions successfully`, func() {
				// Construct an instance of the SetMembersGroupMembers model
				setMembersGroupMembersModel := new(ibmclouddatabasesv5.SetMembersGroupMembers)
				Expect(setMembersGroupMembersModel).ToNot(BeNil())
				setMembersGroupMembersModel.AllocationCount = core.Int64Ptr(int64(4))
				Expect(setMembersGroupMembersModel.AllocationCount).To(Equal(core.Int64Ptr(int64(4))))

				// Construct an instance of the SetDeploymentScalingGroupRequestSetMembersGroup model
				setDeploymentScalingGroupRequestModel := new(ibmclouddatabasesv5.SetDeploymentScalingGroupRequestSetMembersGroup)
				Expect(setDeploymentScalingGroupRequestModel).ToNot(BeNil())
				setDeploymentScalingGroupRequestModel.Members = setMembersGroupMembersModel
				Expect(setDeploymentScalingGroupRequestModel.Members).To(Equal(setMembersGroupMembersModel))

				// Construct an instance of the SetDeploymentScalingGroupOptions model
				id := "testString"
				groupID := "testString"
				var setDeploymentScalingGroupRequest ibmclouddatabasesv5.SetDeploymentScalingGroupRequestIntf = nil
				setDeploymentScalingGroupOptionsModel := ibmCloudDatabasesService.NewSetDeploymentScalingGroupOptions(id, groupID, setDeploymentScalingGroupRequest)
				setDeploymentScalingGroupOptionsModel.SetID("testString")
				setDeploymentScalingGroupOptionsModel.SetGroupID("testString")
				setDeploymentScalingGroupOptionsModel.SetSetDeploymentScalingGroupRequest(setDeploymentScalingGroupRequestModel)
				setDeploymentScalingGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setDeploymentScalingGroupOptionsModel).ToNot(BeNil())
				Expect(setDeploymentScalingGroupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setDeploymentScalingGroupOptionsModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(setDeploymentScalingGroupOptionsModel.SetDeploymentScalingGroupRequest).To(Equal(setDeploymentScalingGroupRequestModel))
				Expect(setDeploymentScalingGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetPromotionOptions successfully`, func() {
				// Construct an instance of the SetPromotionPromotionPromote model
				setPromotionPromotionModel := new(ibmclouddatabasesv5.SetPromotionPromotionPromote)
				Expect(setPromotionPromotionModel).ToNot(BeNil())
				setPromotionPromotionModel.Promotion = make(map[string]interface{})
				Expect(setPromotionPromotionModel.Promotion).To(Equal(make(map[string]interface{})))

				// Construct an instance of the SetPromotionOptions model
				id := "testString"
				var setPromotionOptionsPromotion ibmclouddatabasesv5.SetPromotionPromotionIntf = nil
				setPromotionOptionsModel := ibmCloudDatabasesService.NewSetPromotionOptions(id, setPromotionOptionsPromotion)
				setPromotionOptionsModel.SetID("testString")
				setPromotionOptionsModel.SetPromotion(setPromotionPromotionModel)
				setPromotionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setPromotionOptionsModel).ToNot(BeNil())
				Expect(setPromotionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setPromotionOptionsModel.Promotion).To(Equal(setPromotionPromotionModel))
				Expect(setPromotionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetRemotesOptions successfully`, func() {
				// Construct an instance of the SetRemotesRequestRemotes model
				setRemotesRequestRemotesModel := new(ibmclouddatabasesv5.SetRemotesRequestRemotes)
				Expect(setRemotesRequestRemotesModel).ToNot(BeNil())
				setRemotesRequestRemotesModel.Leader = core.StringPtr("testString")
				Expect(setRemotesRequestRemotesModel.Leader).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SetRemotesOptions model
				id := "testString"
				setRemotesOptionsModel := ibmCloudDatabasesService.NewSetRemotesOptions(id)
				setRemotesOptionsModel.SetID("testString")
				setRemotesOptionsModel.SetRemotes(setRemotesRequestRemotesModel)
				setRemotesOptionsModel.SetSkipInitialBackup(true)
				setRemotesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setRemotesOptionsModel).ToNot(BeNil())
				Expect(setRemotesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setRemotesOptionsModel.Remotes).To(Equal(setRemotesRequestRemotesModel))
				Expect(setRemotesOptionsModel.SkipInitialBackup).To(Equal(core.BoolPtr(true)))
				Expect(setRemotesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewStartOndemandBackupOptions successfully`, func() {
				// Construct an instance of the StartOndemandBackupOptions model
				id := "testString"
				startOndemandBackupOptionsModel := ibmCloudDatabasesService.NewStartOndemandBackupOptions(id)
				startOndemandBackupOptionsModel.SetID("testString")
				startOndemandBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(startOndemandBackupOptionsModel).ToNot(BeNil())
				Expect(startOndemandBackupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(startOndemandBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
