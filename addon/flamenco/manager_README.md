# Flamenco
Render Farm manager API

The `flamenco.manager` package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: 1.0.0
- Package version: 3.0
- Build package: org.openapitools.codegen.languages.PythonClientCodegen
For more information, please visit [https://flamenco.io/](https://flamenco.io/)

## Requirements.

Python >=3.6

## Installation & Usage

This python library package is generated without supporting files like setup.py or requirements files

To be able to use it, you will need these dependencies in your own package that uses this library:

* urllib3 >= 1.25.3
* python-dateutil

## Getting Started

In your own code, to use this library to connect and interact with Flamenco,
you can run the following:

```python

import time
import flamenco.manager
from pprint import pprint
from flamenco.manager.api import jobs_api
from flamenco.manager.model.available_job_types import AvailableJobTypes
from flamenco.manager.model.error import Error
from flamenco.manager.model.job import Job
from flamenco.manager.model.submitted_job import SubmittedJob
# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = flamenco.manager.Configuration(
    host = "http://localhost"
)



# Enter a context with an instance of the API client
with flamenco.manager.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_id = "job_id_example" # str | 

    try:
        # Fetch info about the job.
        api_response = api_instance.fetch_job(job_id)
        pprint(api_response)
    except flamenco.manager.ApiException as e:
        print("Exception when calling JobsApi->fetch_job: %s\n" % e)
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*JobsApi* | [**fetch_job**](flamenco/manager/docs/JobsApi.md#fetch_job) | **GET** /api/jobs/{job_id} | Fetch info about the job.
*JobsApi* | [**get_job_types**](flamenco/manager/docs/JobsApi.md#get_job_types) | **GET** /api/jobs/types | Get list of job types and their parameters.
*JobsApi* | [**submit_job**](flamenco/manager/docs/JobsApi.md#submit_job) | **POST** /api/jobs | Submit a new job for Flamenco Manager to execute.
*WorkerApi* | [**register_worker**](flamenco/manager/docs/WorkerApi.md#register_worker) | **POST** /api/worker/register-worker | Register a new worker
*WorkerApi* | [**schedule_task**](flamenco/manager/docs/WorkerApi.md#schedule_task) | **POST** /api/worker/task | Obtain a new task to execute
*WorkerApi* | [**sign_off**](flamenco/manager/docs/WorkerApi.md#sign_off) | **POST** /api/worker/sign-off | Mark the worker as offline
*WorkerApi* | [**sign_on**](flamenco/manager/docs/WorkerApi.md#sign_on) | **POST** /api/worker/sign-on | Authenticate &amp; sign in the worker.
*WorkerApi* | [**task_update**](flamenco/manager/docs/WorkerApi.md#task_update) | **POST** /api/worker/task/{task_id} | Update the task, typically to indicate progress, completion, or failure.
*WorkerApi* | [**worker_state**](flamenco/manager/docs/WorkerApi.md#worker_state) | **GET** /api/worker/state | 
*WorkerApi* | [**worker_state_changed**](flamenco/manager/docs/WorkerApi.md#worker_state_changed) | **POST** /api/worker/state-changed | Worker changed state. This could be as acknowledgement of a Manager-requested state change, or in response to worker-local signals.


## Documentation For Models

 - [AssignedTask](flamenco/manager/docs/AssignedTask.md)
 - [AvailableJobSetting](flamenco/manager/docs/AvailableJobSetting.md)
 - [AvailableJobSettingSubtype](flamenco/manager/docs/AvailableJobSettingSubtype.md)
 - [AvailableJobSettingType](flamenco/manager/docs/AvailableJobSettingType.md)
 - [AvailableJobType](flamenco/manager/docs/AvailableJobType.md)
 - [AvailableJobTypes](flamenco/manager/docs/AvailableJobTypes.md)
 - [Command](flamenco/manager/docs/Command.md)
 - [Configuration](flamenco/manager/docs/Configuration.md)
 - [ConfigurationMeta](flamenco/manager/docs/ConfigurationMeta.md)
 - [Error](flamenco/manager/docs/Error.md)
 - [Job](flamenco/manager/docs/Job.md)
 - [JobAllOf](flamenco/manager/docs/JobAllOf.md)
 - [JobMetadata](flamenco/manager/docs/JobMetadata.md)
 - [JobSettings](flamenco/manager/docs/JobSettings.md)
 - [JobStatus](flamenco/manager/docs/JobStatus.md)
 - [RegisteredWorker](flamenco/manager/docs/RegisteredWorker.md)
 - [SecurityError](flamenco/manager/docs/SecurityError.md)
 - [SubmittedJob](flamenco/manager/docs/SubmittedJob.md)
 - [TaskStatus](flamenco/manager/docs/TaskStatus.md)
 - [TaskUpdate](flamenco/manager/docs/TaskUpdate.md)
 - [WorkerRegistration](flamenco/manager/docs/WorkerRegistration.md)
 - [WorkerSignOn](flamenco/manager/docs/WorkerSignOn.md)
 - [WorkerStateChange](flamenco/manager/docs/WorkerStateChange.md)
 - [WorkerStateChanged](flamenco/manager/docs/WorkerStateChanged.md)
 - [WorkerStatus](flamenco/manager/docs/WorkerStatus.md)


## Documentation For Authorization


## worker_auth

- **Type**: HTTP basic authentication


## Author




## Notes for Large OpenAPI documents
If the OpenAPI document is large, imports in flamenco.manager.apis and flamenco.manager.models may fail with a
RecursionError indicating the maximum recursion limit has been exceeded. In that case, there are a couple of solutions:

Solution 1:
Use specific imports for apis and models like:
- `from flamenco.manager.api.default_api import DefaultApi`
- `from flamenco.manager.model.pet import Pet`

Solution 2:
Before importing the package, adjust the maximum recursion limit as shown below:
```
import sys
sys.setrecursionlimit(1500)
import flamenco.manager
from flamenco.manager.apis import *
from flamenco.manager.models import *
```
