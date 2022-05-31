# flamenco.manager.WorkerMgtApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**fetch_worker**](WorkerMgtApi.md#fetch_worker) | **GET** /api/worker-mgt/workers/{worker_id} | Fetch info about the worker.
[**fetch_workers**](WorkerMgtApi.md#fetch_workers) | **GET** /api/worker-mgt/workers | Get list of workers.


# **fetch_worker**
> Worker fetch_worker(worker_id)

Fetch info about the worker.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.worker import Worker
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = flamenco.manager.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with flamenco.manager.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = worker_mgt_api.WorkerMgtApi(api_client)
    worker_id = "worker_id_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Fetch info about the worker.
        api_response = api_instance.fetch_worker(worker_id)
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->fetch_worker: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_id** | **str**|  |

### Return type

[**Worker**](Worker.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Worker info |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fetch_workers**
> WorkerList fetch_workers()

Get list of workers.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.worker_list import WorkerList
from pprint import pprint
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = flamenco.manager.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with flamenco.manager.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = worker_mgt_api.WorkerMgtApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        # Get list of workers.
        api_response = api_instance.fetch_workers()
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->fetch_workers: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**WorkerList**](WorkerList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Known workers |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
