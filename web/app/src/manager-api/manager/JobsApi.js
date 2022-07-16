/**
 * Flamenco manager
 * Render Farm manager API
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import AvailableJobType from '../model/AvailableJobType';
import AvailableJobTypes from '../model/AvailableJobTypes';
import Error from '../model/Error';
import Job from '../model/Job';
import JobBlocklistEntry from '../model/JobBlocklistEntry';
import JobLastRenderedImageInfo from '../model/JobLastRenderedImageInfo';
import JobStatusChange from '../model/JobStatusChange';
import JobTasksSummary from '../model/JobTasksSummary';
import JobsQuery from '../model/JobsQuery';
import JobsQueryResult from '../model/JobsQueryResult';
import SubmittedJob from '../model/SubmittedJob';
import Task from '../model/Task';
import TaskLogInfo from '../model/TaskLogInfo';
import TaskStatusChange from '../model/TaskStatusChange';

/**
* Jobs service.
* @module manager/JobsApi
* @version 0.0.0
*/
export default class JobsApi {

    /**
    * Constructs a new JobsApi. 
    * @alias module:manager/JobsApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }



    /**
     * Get the URL that serves the last-rendered images.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/JobLastRenderedImageInfo} and HTTP response
     */
    fetchGlobalLastRenderedInfoWithHttpInfo() {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = JobLastRenderedImageInfo;
      return this.apiClient.callApi(
        '/api/v3/jobs/last-rendered', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get the URL that serves the last-rendered images.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/JobLastRenderedImageInfo}
     */
    fetchGlobalLastRenderedInfo() {
      return this.fetchGlobalLastRenderedInfoWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Fetch info about the job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/Job} and HTTP response
     */
    fetchJobWithHttpInfo(jobId) {
      let postBody = null;
      // verify the required parameter 'jobId' is set
      if (jobId === undefined || jobId === null) {
        throw new Error("Missing the required parameter 'jobId' when calling fetchJob");
      }

      let pathParams = {
        'job_id': jobId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Job;
      return this.apiClient.callApi(
        '/api/v3/jobs/{job_id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Fetch info about the job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/Job}
     */
    fetchJob(jobId) {
      return this.fetchJobWithHttpInfo(jobId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Fetch the list of workers that are blocked from doing certain task types on this job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link Array.<module:model/JobBlocklistEntry>} and HTTP response
     */
    fetchJobBlocklistWithHttpInfo(jobId) {
      let postBody = null;
      // verify the required parameter 'jobId' is set
      if (jobId === undefined || jobId === null) {
        throw new Error("Missing the required parameter 'jobId' when calling fetchJobBlocklist");
      }

      let pathParams = {
        'job_id': jobId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = [JobBlocklistEntry];
      return this.apiClient.callApi(
        '/api/v3/jobs/{job_id}/blocklist', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Fetch the list of workers that are blocked from doing certain task types on this job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link Array.<module:model/JobBlocklistEntry>}
     */
    fetchJobBlocklist(jobId) {
      return this.fetchJobBlocklistWithHttpInfo(jobId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Get the URL that serves the last-rendered images of this job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/JobLastRenderedImageInfo} and HTTP response
     */
    fetchJobLastRenderedInfoWithHttpInfo(jobId) {
      let postBody = null;
      // verify the required parameter 'jobId' is set
      if (jobId === undefined || jobId === null) {
        throw new Error("Missing the required parameter 'jobId' when calling fetchJobLastRenderedInfo");
      }

      let pathParams = {
        'job_id': jobId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = JobLastRenderedImageInfo;
      return this.apiClient.callApi(
        '/api/v3/jobs/{job_id}/last-rendered', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get the URL that serves the last-rendered images of this job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/JobLastRenderedImageInfo}
     */
    fetchJobLastRenderedInfo(jobId) {
      return this.fetchJobLastRenderedInfoWithHttpInfo(jobId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Fetch a summary of all tasks of the given job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/JobTasksSummary} and HTTP response
     */
    fetchJobTasksWithHttpInfo(jobId) {
      let postBody = null;
      // verify the required parameter 'jobId' is set
      if (jobId === undefined || jobId === null) {
        throw new Error("Missing the required parameter 'jobId' when calling fetchJobTasks");
      }

      let pathParams = {
        'job_id': jobId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = JobTasksSummary;
      return this.apiClient.callApi(
        '/api/v3/jobs/{job_id}/tasks', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Fetch a summary of all tasks of the given job.
     * @param {String} jobId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/JobTasksSummary}
     */
    fetchJobTasks(jobId) {
      return this.fetchJobTasksWithHttpInfo(jobId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Fetch a single task.
     * @param {String} taskId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/Task} and HTTP response
     */
    fetchTaskWithHttpInfo(taskId) {
      let postBody = null;
      // verify the required parameter 'taskId' is set
      if (taskId === undefined || taskId === null) {
        throw new Error("Missing the required parameter 'taskId' when calling fetchTask");
      }

      let pathParams = {
        'task_id': taskId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Task;
      return this.apiClient.callApi(
        '/api/v3/tasks/{task_id}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Fetch a single task.
     * @param {String} taskId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/Task}
     */
    fetchTask(taskId) {
      return this.fetchTaskWithHttpInfo(taskId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Get the URL of the task log, and some more info.
     * @param {String} taskId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/TaskLogInfo} and HTTP response
     */
    fetchTaskLogInfoWithHttpInfo(taskId) {
      let postBody = null;
      // verify the required parameter 'taskId' is set
      if (taskId === undefined || taskId === null) {
        throw new Error("Missing the required parameter 'taskId' when calling fetchTaskLogInfo");
      }

      let pathParams = {
        'task_id': taskId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = TaskLogInfo;
      return this.apiClient.callApi(
        '/api/v3/tasks/{task_id}/log', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get the URL of the task log, and some more info.
     * @param {String} taskId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/TaskLogInfo}
     */
    fetchTaskLogInfo(taskId) {
      return this.fetchTaskLogInfoWithHttpInfo(taskId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Fetch the last few lines of the task's log.
     * @param {String} taskId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link String} and HTTP response
     */
    fetchTaskLogTailWithHttpInfo(taskId) {
      let postBody = null;
      // verify the required parameter 'taskId' is set
      if (taskId === undefined || taskId === null) {
        throw new Error("Missing the required parameter 'taskId' when calling fetchTaskLogTail");
      }

      let pathParams = {
        'task_id': taskId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['text/plain', 'application/json'];
      let returnType = 'String';
      return this.apiClient.callApi(
        '/api/v3/tasks/{task_id}/logtail', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Fetch the last few lines of the task's log.
     * @param {String} taskId 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link String}
     */
    fetchTaskLogTail(taskId) {
      return this.fetchTaskLogTailWithHttpInfo(taskId)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Get single job type and its parameters.
     * @param {String} typeName 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/AvailableJobType} and HTTP response
     */
    getJobTypeWithHttpInfo(typeName) {
      let postBody = null;
      // verify the required parameter 'typeName' is set
      if (typeName === undefined || typeName === null) {
        throw new Error("Missing the required parameter 'typeName' when calling getJobType");
      }

      let pathParams = {
        'typeName': typeName
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = AvailableJobType;
      return this.apiClient.callApi(
        '/api/v3/jobs/type/{typeName}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get single job type and its parameters.
     * @param {String} typeName 
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/AvailableJobType}
     */
    getJobType(typeName) {
      return this.getJobTypeWithHttpInfo(typeName)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Get list of job types and their parameters.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/AvailableJobTypes} and HTTP response
     */
    getJobTypesWithHttpInfo() {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = AvailableJobTypes;
      return this.apiClient.callApi(
        '/api/v3/jobs/types', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Get list of job types and their parameters.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/AvailableJobTypes}
     */
    getJobTypes() {
      return this.getJobTypesWithHttpInfo()
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Fetch list of jobs.
     * @param {module:model/JobsQuery} jobsQuery Specification of which jobs to get.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/JobsQueryResult} and HTTP response
     */
    queryJobsWithHttpInfo(jobsQuery) {
      let postBody = jobsQuery;
      // verify the required parameter 'jobsQuery' is set
      if (jobsQuery === undefined || jobsQuery === null) {
        throw new Error("Missing the required parameter 'jobsQuery' when calling queryJobs");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = JobsQueryResult;
      return this.apiClient.callApi(
        '/api/v3/jobs/query', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Fetch list of jobs.
     * @param {module:model/JobsQuery} jobsQuery Specification of which jobs to get.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/JobsQueryResult}
     */
    queryJobs(jobsQuery) {
      return this.queryJobsWithHttpInfo(jobsQuery)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Remove entries from a job blocklist.
     * @param {String} jobId 
     * @param {Object} opts Optional parameters
     * @param {Array.<module:model/JobBlocklistEntry>} opts.jobBlocklistEntry Tuples (worker, task type) to be removed from the blocklist.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing HTTP response
     */
    removeJobBlocklistWithHttpInfo(jobId, opts) {
      opts = opts || {};
      let postBody = opts['jobBlocklistEntry'];
      // verify the required parameter 'jobId' is set
      if (jobId === undefined || jobId === null) {
        throw new Error("Missing the required parameter 'jobId' when calling removeJobBlocklist");
      }

      let pathParams = {
        'job_id': jobId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v3/jobs/{job_id}/blocklist', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Remove entries from a job blocklist.
     * @param {String} jobId 
     * @param {Object} opts Optional parameters
     * @param {Array.<module:model/JobBlocklistEntry>} opts.jobBlocklistEntry Tuples (worker, task type) to be removed from the blocklist.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}
     */
    removeJobBlocklist(jobId, opts) {
      return this.removeJobBlocklistWithHttpInfo(jobId, opts)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @param {String} jobId 
     * @param {module:model/JobStatusChange} jobStatusChange The status change to request.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing HTTP response
     */
    setJobStatusWithHttpInfo(jobId, jobStatusChange) {
      let postBody = jobStatusChange;
      // verify the required parameter 'jobId' is set
      if (jobId === undefined || jobId === null) {
        throw new Error("Missing the required parameter 'jobId' when calling setJobStatus");
      }
      // verify the required parameter 'jobStatusChange' is set
      if (jobStatusChange === undefined || jobStatusChange === null) {
        throw new Error("Missing the required parameter 'jobStatusChange' when calling setJobStatus");
      }

      let pathParams = {
        'job_id': jobId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v3/jobs/{job_id}/setstatus', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * @param {String} jobId 
     * @param {module:model/JobStatusChange} jobStatusChange The status change to request.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}
     */
    setJobStatus(jobId, jobStatusChange) {
      return this.setJobStatusWithHttpInfo(jobId, jobStatusChange)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * @param {String} taskId 
     * @param {module:model/TaskStatusChange} taskStatusChange The status change to request.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing HTTP response
     */
    setTaskStatusWithHttpInfo(taskId, taskStatusChange) {
      let postBody = taskStatusChange;
      // verify the required parameter 'taskId' is set
      if (taskId === undefined || taskId === null) {
        throw new Error("Missing the required parameter 'taskId' when calling setTaskStatus");
      }
      // verify the required parameter 'taskStatusChange' is set
      if (taskStatusChange === undefined || taskStatusChange === null) {
        throw new Error("Missing the required parameter 'taskStatusChange' when calling setTaskStatus");
      }

      let pathParams = {
        'task_id': taskId
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v3/tasks/{task_id}/setstatus', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * @param {String} taskId 
     * @param {module:model/TaskStatusChange} taskStatusChange The status change to request.
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}
     */
    setTaskStatus(taskId, taskStatusChange) {
      return this.setTaskStatusWithHttpInfo(taskId, taskStatusChange)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


    /**
     * Submit a new job for Flamenco Manager to execute.
     * @param {module:model/SubmittedJob} submittedJob Job to submit
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/Job} and HTTP response
     */
    submitJobWithHttpInfo(submittedJob) {
      let postBody = submittedJob;
      // verify the required parameter 'submittedJob' is set
      if (submittedJob === undefined || submittedJob === null) {
        throw new Error("Missing the required parameter 'submittedJob' when calling submitJob");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = Job;
      return this.apiClient.callApi(
        '/api/v3/jobs', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null
      );
    }

    /**
     * Submit a new job for Flamenco Manager to execute.
     * @param {module:model/SubmittedJob} submittedJob Job to submit
     * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/Job}
     */
    submitJob(submittedJob) {
      return this.submitJobWithHttpInfo(submittedJob)
        .then(function(response_and_data) {
          return response_and_data.data;
        });
    }


}
