# flamenco.manager.WorkerMgtApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_worker_cluster**](WorkerMgtApi.md#create_worker_cluster) | **POST** /api/v3/worker-mgt/clusters | Create a new worker cluster.
[**delete_worker**](WorkerMgtApi.md#delete_worker) | **DELETE** /api/v3/worker-mgt/workers/{worker_id} | Remove the given worker. It is recommended to only call this function when the worker is in &#x60;offline&#x60; state. If the worker is still running, stop it first. Any task still assigned to the worker will be requeued. 
[**delete_worker_cluster**](WorkerMgtApi.md#delete_worker_cluster) | **DELETE** /api/v3/worker-mgt/cluster/{cluster_id} | Remove this worker cluster. This unassigns all workers from the cluster and removes it.
[**fetch_worker**](WorkerMgtApi.md#fetch_worker) | **GET** /api/v3/worker-mgt/workers/{worker_id} | Fetch info about the worker.
[**fetch_worker_cluster**](WorkerMgtApi.md#fetch_worker_cluster) | **GET** /api/v3/worker-mgt/cluster/{cluster_id} | Get a single worker cluster.
[**fetch_worker_clusters**](WorkerMgtApi.md#fetch_worker_clusters) | **GET** /api/v3/worker-mgt/clusters | Get list of worker clusters.
[**fetch_worker_sleep_schedule**](WorkerMgtApi.md#fetch_worker_sleep_schedule) | **GET** /api/v3/worker-mgt/workers/{worker_id}/sleep-schedule | 
[**fetch_workers**](WorkerMgtApi.md#fetch_workers) | **GET** /api/v3/worker-mgt/workers | Get list of workers.
[**request_worker_status_change**](WorkerMgtApi.md#request_worker_status_change) | **POST** /api/v3/worker-mgt/workers/{worker_id}/setstatus | 
[**set_worker_clusters**](WorkerMgtApi.md#set_worker_clusters) | **POST** /api/v3/worker-mgt/workers/{worker_id}/setclusters | 
[**set_worker_sleep_schedule**](WorkerMgtApi.md#set_worker_sleep_schedule) | **POST** /api/v3/worker-mgt/workers/{worker_id}/sleep-schedule | 
[**update_worker_cluster**](WorkerMgtApi.md#update_worker_cluster) | **PUT** /api/v3/worker-mgt/cluster/{cluster_id} | Update an existing worker cluster.


# **create_worker_cluster**
> WorkerCluster create_worker_cluster(worker_cluster)

Create a new worker cluster.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
from flamenco.manager.model.worker_cluster import WorkerCluster
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
    worker_cluster = WorkerCluster(
        id="id_example",
        name="name_example",
        description="description_example",
    ) # WorkerCluster | The worker cluster.

    # example passing only required values which don't have defaults set
    try:
        # Create a new worker cluster.
        api_response = api_instance.create_worker_cluster(worker_cluster)
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->create_worker_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_cluster** | [**WorkerCluster**](WorkerCluster.md)| The worker cluster. |

### Return type

[**WorkerCluster**](WorkerCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The cluster was created. The created cluster is returned, so that the caller can know its UUID. |  -  |
**0** | Error message |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_worker**
> delete_worker(worker_id)

Remove the given worker. It is recommended to only call this function when the worker is in `offline` state. If the worker is still running, stop it first. Any task still assigned to the worker will be requeued. 

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
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
        # Remove the given worker. It is recommended to only call this function when the worker is in `offline` state. If the worker is still running, stop it first. Any task still assigned to the worker will be requeued. 
        api_instance.delete_worker(worker_id)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->delete_worker: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_id** | **str**|  |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Normal response, worker has been deleted |  -  |
**0** | Unexpected error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_worker_cluster**
> delete_worker_cluster(cluster_id)

Remove this worker cluster. This unassigns all workers from the cluster and removes it.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
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
    cluster_id = "cluster_id_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Remove this worker cluster. This unassigns all workers from the cluster and removes it.
        api_instance.delete_worker_cluster(cluster_id)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->delete_worker_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **str**|  |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | The cluster has been removed. |  -  |
**0** | Unexpected error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **fetch_worker_cluster**
> WorkerCluster fetch_worker_cluster(cluster_id)

Get a single worker cluster.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.worker_cluster import WorkerCluster
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
    cluster_id = "cluster_id_example" # str | 

    # example passing only required values which don't have defaults set
    try:
        # Get a single worker cluster.
        api_response = api_instance.fetch_worker_cluster(cluster_id)
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->fetch_worker_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **str**|  |

### Return type

[**WorkerCluster**](WorkerCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The worker cluster. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fetch_worker_clusters**
> WorkerClusterList fetch_worker_clusters()

Get list of worker clusters.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.worker_cluster_list import WorkerClusterList
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
        # Get list of worker clusters.
        api_response = api_instance.fetch_worker_clusters()
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->fetch_worker_clusters: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

[**WorkerClusterList**](WorkerClusterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Worker clusters. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fetch_worker_sleep_schedule**
> WorkerSleepSchedule fetch_worker_sleep_schedule(worker_id)



### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
from flamenco.manager.model.worker_sleep_schedule import WorkerSleepSchedule
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
        api_response = api_instance.fetch_worker_sleep_schedule(worker_id)
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->fetch_worker_sleep_schedule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_id** | **str**|  |

### Return type

[**WorkerSleepSchedule**](WorkerSleepSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Normal response, the sleep schedule. |  -  |
**204** | The worker has no sleep schedule. |  -  |
**0** | Unexpected error. |  -  |

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

# **request_worker_status_change**
> request_worker_status_change(worker_id, worker_status_change_request)



### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
from flamenco.manager.model.worker_status_change_request import WorkerStatusChangeRequest
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
    worker_status_change_request = WorkerStatusChangeRequest(
        status=WorkerStatus("starting"),
        is_lazy=True,
    ) # WorkerStatusChangeRequest | The status change to request.

    # example passing only required values which don't have defaults set
    try:
        api_instance.request_worker_status_change(worker_id, worker_status_change_request)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->request_worker_status_change: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_id** | **str**|  |
 **worker_status_change_request** | [**WorkerStatusChangeRequest**](WorkerStatusChangeRequest.md)| The status change to request. |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Status change was accepted. |  -  |
**0** | Unexpected error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_worker_clusters**
> set_worker_clusters(worker_id, worker_cluster_change_request)



### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
from flamenco.manager.model.worker_cluster_change_request import WorkerClusterChangeRequest
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
    worker_cluster_change_request = WorkerClusterChangeRequest(
        cluster_ids=[
            "cluster_ids_example",
        ],
    ) # WorkerClusterChangeRequest | The list of cluster IDs this worker should be a member of.

    # example passing only required values which don't have defaults set
    try:
        api_instance.set_worker_clusters(worker_id, worker_cluster_change_request)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->set_worker_clusters: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_id** | **str**|  |
 **worker_cluster_change_request** | [**WorkerClusterChangeRequest**](WorkerClusterChangeRequest.md)| The list of cluster IDs this worker should be a member of. |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | Status change was accepted. |  -  |
**0** | Unexpected error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **set_worker_sleep_schedule**
> set_worker_sleep_schedule(worker_id, worker_sleep_schedule)



### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
from flamenco.manager.model.worker_sleep_schedule import WorkerSleepSchedule
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
    worker_sleep_schedule = WorkerSleepSchedule(
        is_active=True,
        days_of_week="days_of_week_example",
        start_time="start_time_example",
        end_time="end_time_example",
    ) # WorkerSleepSchedule | The new sleep schedule.

    # example passing only required values which don't have defaults set
    try:
        api_instance.set_worker_sleep_schedule(worker_id, worker_sleep_schedule)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->set_worker_sleep_schedule: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worker_id** | **str**|  |
 **worker_sleep_schedule** | [**WorkerSleepSchedule**](WorkerSleepSchedule.md)| The new sleep schedule. |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | The schedule has been stored. |  -  |
**0** | Unexpected error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_worker_cluster**
> update_worker_cluster(cluster_id, worker_cluster)

Update an existing worker cluster.

### Example


```python
import time
import flamenco.manager
from flamenco.manager.api import worker_mgt_api
from flamenco.manager.model.error import Error
from flamenco.manager.model.worker_cluster import WorkerCluster
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
    cluster_id = "cluster_id_example" # str | 
    worker_cluster = WorkerCluster(
        id="id_example",
        name="name_example",
        description="description_example",
    ) # WorkerCluster | The updated worker cluster.

    # example passing only required values which don't have defaults set
    try:
        # Update an existing worker cluster.
        api_instance.update_worker_cluster(cluster_id, worker_cluster)
    except flamenco.manager.ApiException as e:
        print("Exception when calling WorkerMgtApi->update_worker_cluster: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster_id** | **str**|  |
 **worker_cluster** | [**WorkerCluster**](WorkerCluster.md)| The updated worker cluster. |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | The cluster update has been stored. |  -  |
**0** | Error message |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

