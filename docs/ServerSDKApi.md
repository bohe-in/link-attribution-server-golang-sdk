# \ServerSDKApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServerSdkCreateLink**](ServerSDKApi.md#ServerSdkCreateLink) | **Post** /sdkserv/v1/links | 



## ServerSdkCreateLink

> ResponsesSdkLinkResponse ServerSdkCreateLink(ctx).XApiKey(xApiKey).XApiSecret(xApiSecret).ServerSdkCreateLink(serverSdkCreateLink).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xApiKey := "xApiKey_example" // string | 
    xApiSecret := "xApiSecret_example" // string | 
    serverSdkCreateLink := *openapiclient.NewServerSdkCreateLink("Name_example") // ServerSdkCreateLink | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerSDKApi.ServerSdkCreateLink(context.Background()).XApiKey(xApiKey).XApiSecret(xApiSecret).ServerSdkCreateLink(serverSdkCreateLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerSDKApi.ServerSdkCreateLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServerSdkCreateLink`: ResponsesSdkLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerSDKApi.ServerSdkCreateLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServerSdkCreateLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **xApiSecret** | **string** |  | 
 **serverSdkCreateLink** | [**ServerSdkCreateLink**](ServerSdkCreateLink.md) |  | 

### Return type

[**ResponsesSdkLinkResponse**](ResponsesSdkLinkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

