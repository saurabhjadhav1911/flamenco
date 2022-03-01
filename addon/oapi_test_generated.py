#!/usr/bin/env python3

import flamenco.manager
from pprint import pprint
from flamenco.manager.api import jobs_api
from flamenco.manager.model.available_job_types import AvailableJobTypes
from flamenco.manager.model.available_job_type import AvailableJobType
from flamenco.manager.model.error import Error
from flamenco.manager.model.job import Job
from flamenco.manager.model.submitted_job import SubmittedJob

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = flamenco.manager.Configuration(host="http://localhost:8080")


# Enter a context with an instance of the API client
with flamenco.manager.ApiClient(configuration) as api_client:
    job_api_instance = jobs_api.JobsApi(api_client)

    response: AvailableJobTypes = job_api_instance.get_job_types()
    for job_type in response.job_types:
        job_type: AvailableJobType
        print(f"\033[38;5;214m{job_type.label}\033[0m ({job_type.name})")
        for setting in job_type.settings:
            print(f"  - {setting.key:23}  type: {setting.type!r:10}", end="")
            default = getattr(setting, "default", None)
            if default is not None:
                print(f"  default: {repr(default)}", end="")
            print()

    # job_id = "2f03614f-f529-445f-b8c5-272e3f437b73"
    # try:
    #     # Fetch info about the job.
    #     api_response = job_api_instance.fetch_job(job_id)
    #     pprint(api_response)
    # except flamenco.manager.ApiException as e:
    #     print("Exception when calling JobsApi->fetch_job: %s\n" % e)
