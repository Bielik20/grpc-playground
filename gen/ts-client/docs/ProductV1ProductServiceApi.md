# ProductV1ProductServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**productV1ProductServiceGetProduct**](ProductV1ProductServiceApi.md#productv1productservicegetproduct) | **GET** /v1/products/{product_id} | GetProduct |



## productV1ProductServiceGetProduct

> ProductV1Product productV1ProductServiceGetProduct(productId)

GetProduct

1. Standard gRPC definition

### Example

```ts
import {
  Configuration,
  ProductV1ProductServiceApi,
} from '';
import type { ProductV1ProductServiceGetProductRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new ProductV1ProductServiceApi();

  const body = {
    // string
    productId: productId_example,
  } satisfies ProductV1ProductServiceGetProductRequest;

  try {
    const data = await api.productV1ProductServiceGetProduct(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **productId** | `string` |  | [Defaults to `undefined`] |

### Return type

[**ProductV1Product**](ProductV1Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

