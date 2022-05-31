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


import ApiClient from './ApiClient';
import AssignedTask from './model/AssignedTask';
import AvailableJobSetting from './model/AvailableJobSetting';
import AvailableJobSettingSubtype from './model/AvailableJobSettingSubtype';
import AvailableJobSettingType from './model/AvailableJobSettingType';
import AvailableJobSettingVisibility from './model/AvailableJobSettingVisibility';
import AvailableJobType from './model/AvailableJobType';
import AvailableJobTypes from './model/AvailableJobTypes';
import Command from './model/Command';
import Error from './model/Error';
import FlamencoVersion from './model/FlamencoVersion';
import Job from './model/Job';
import JobAllOf from './model/JobAllOf';
import JobStatus from './model/JobStatus';
import JobStatusChange from './model/JobStatusChange';
import JobTasksSummary from './model/JobTasksSummary';
import JobsQuery from './model/JobsQuery';
import JobsQueryResult from './model/JobsQueryResult';
import ManagerConfiguration from './model/ManagerConfiguration';
import MayKeepRunning from './model/MayKeepRunning';
import RegisteredWorker from './model/RegisteredWorker';
import SecurityError from './model/SecurityError';
import ShamanCheckout from './model/ShamanCheckout';
import ShamanCheckoutResult from './model/ShamanCheckoutResult';
import ShamanFileSpec from './model/ShamanFileSpec';
import ShamanFileSpecWithStatus from './model/ShamanFileSpecWithStatus';
import ShamanFileStatus from './model/ShamanFileStatus';
import ShamanRequirementsRequest from './model/ShamanRequirementsRequest';
import ShamanRequirementsResponse from './model/ShamanRequirementsResponse';
import ShamanSingleFileStatus from './model/ShamanSingleFileStatus';
import SocketIOJobUpdate from './model/SocketIOJobUpdate';
import SocketIOSubscription from './model/SocketIOSubscription';
import SocketIOSubscriptionOperation from './model/SocketIOSubscriptionOperation';
import SocketIOSubscriptionType from './model/SocketIOSubscriptionType';
import SocketIOTaskLogUpdate from './model/SocketIOTaskLogUpdate';
import SocketIOTaskUpdate from './model/SocketIOTaskUpdate';
import SocketIOWorkerUpdate from './model/SocketIOWorkerUpdate';
import SubmittedJob from './model/SubmittedJob';
import Task from './model/Task';
import TaskStatus from './model/TaskStatus';
import TaskStatusChange from './model/TaskStatusChange';
import TaskSummary from './model/TaskSummary';
import TaskUpdate from './model/TaskUpdate';
import TaskWorker from './model/TaskWorker';
import Worker from './model/Worker';
import WorkerList from './model/WorkerList';
import WorkerRegistration from './model/WorkerRegistration';
import WorkerSignOn from './model/WorkerSignOn';
import WorkerStateChange from './model/WorkerStateChange';
import WorkerStateChanged from './model/WorkerStateChanged';
import WorkerStatus from './model/WorkerStatus';
import WorkerSummary from './model/WorkerSummary';
import JobsApi from './manager/JobsApi';
import MetaApi from './manager/MetaApi';
import ShamanApi from './manager/ShamanApi';
import WorkerApi from './manager/WorkerApi';
import WorkerMgtApi from './manager/WorkerMgtApi';


/**
* Render_Farm_manager_API.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var flamencoManager = require('index'); // See note below*.
* var xxxSvc = new flamencoManager.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new flamencoManager.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new flamencoManager.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new flamencoManager.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 0.0.0
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The AssignedTask model constructor.
     * @property {module:model/AssignedTask}
     */
    AssignedTask,

    /**
     * The AvailableJobSetting model constructor.
     * @property {module:model/AvailableJobSetting}
     */
    AvailableJobSetting,

    /**
     * The AvailableJobSettingSubtype model constructor.
     * @property {module:model/AvailableJobSettingSubtype}
     */
    AvailableJobSettingSubtype,

    /**
     * The AvailableJobSettingType model constructor.
     * @property {module:model/AvailableJobSettingType}
     */
    AvailableJobSettingType,

    /**
     * The AvailableJobSettingVisibility model constructor.
     * @property {module:model/AvailableJobSettingVisibility}
     */
    AvailableJobSettingVisibility,

    /**
     * The AvailableJobType model constructor.
     * @property {module:model/AvailableJobType}
     */
    AvailableJobType,

    /**
     * The AvailableJobTypes model constructor.
     * @property {module:model/AvailableJobTypes}
     */
    AvailableJobTypes,

    /**
     * The Command model constructor.
     * @property {module:model/Command}
     */
    Command,

    /**
     * The Error model constructor.
     * @property {module:model/Error}
     */
    Error,

    /**
     * The FlamencoVersion model constructor.
     * @property {module:model/FlamencoVersion}
     */
    FlamencoVersion,

    /**
     * The Job model constructor.
     * @property {module:model/Job}
     */
    Job,

    /**
     * The JobAllOf model constructor.
     * @property {module:model/JobAllOf}
     */
    JobAllOf,

    /**
     * The JobStatus model constructor.
     * @property {module:model/JobStatus}
     */
    JobStatus,

    /**
     * The JobStatusChange model constructor.
     * @property {module:model/JobStatusChange}
     */
    JobStatusChange,

    /**
     * The JobTasksSummary model constructor.
     * @property {module:model/JobTasksSummary}
     */
    JobTasksSummary,

    /**
     * The JobsQuery model constructor.
     * @property {module:model/JobsQuery}
     */
    JobsQuery,

    /**
     * The JobsQueryResult model constructor.
     * @property {module:model/JobsQueryResult}
     */
    JobsQueryResult,

    /**
     * The ManagerConfiguration model constructor.
     * @property {module:model/ManagerConfiguration}
     */
    ManagerConfiguration,

    /**
     * The MayKeepRunning model constructor.
     * @property {module:model/MayKeepRunning}
     */
    MayKeepRunning,

    /**
     * The RegisteredWorker model constructor.
     * @property {module:model/RegisteredWorker}
     */
    RegisteredWorker,

    /**
     * The SecurityError model constructor.
     * @property {module:model/SecurityError}
     */
    SecurityError,

    /**
     * The ShamanCheckout model constructor.
     * @property {module:model/ShamanCheckout}
     */
    ShamanCheckout,

    /**
     * The ShamanCheckoutResult model constructor.
     * @property {module:model/ShamanCheckoutResult}
     */
    ShamanCheckoutResult,

    /**
     * The ShamanFileSpec model constructor.
     * @property {module:model/ShamanFileSpec}
     */
    ShamanFileSpec,

    /**
     * The ShamanFileSpecWithStatus model constructor.
     * @property {module:model/ShamanFileSpecWithStatus}
     */
    ShamanFileSpecWithStatus,

    /**
     * The ShamanFileStatus model constructor.
     * @property {module:model/ShamanFileStatus}
     */
    ShamanFileStatus,

    /**
     * The ShamanRequirementsRequest model constructor.
     * @property {module:model/ShamanRequirementsRequest}
     */
    ShamanRequirementsRequest,

    /**
     * The ShamanRequirementsResponse model constructor.
     * @property {module:model/ShamanRequirementsResponse}
     */
    ShamanRequirementsResponse,

    /**
     * The ShamanSingleFileStatus model constructor.
     * @property {module:model/ShamanSingleFileStatus}
     */
    ShamanSingleFileStatus,

    /**
     * The SocketIOJobUpdate model constructor.
     * @property {module:model/SocketIOJobUpdate}
     */
    SocketIOJobUpdate,

    /**
     * The SocketIOSubscription model constructor.
     * @property {module:model/SocketIOSubscription}
     */
    SocketIOSubscription,

    /**
     * The SocketIOSubscriptionOperation model constructor.
     * @property {module:model/SocketIOSubscriptionOperation}
     */
    SocketIOSubscriptionOperation,

    /**
     * The SocketIOSubscriptionType model constructor.
     * @property {module:model/SocketIOSubscriptionType}
     */
    SocketIOSubscriptionType,

    /**
     * The SocketIOTaskLogUpdate model constructor.
     * @property {module:model/SocketIOTaskLogUpdate}
     */
    SocketIOTaskLogUpdate,

    /**
     * The SocketIOTaskUpdate model constructor.
     * @property {module:model/SocketIOTaskUpdate}
     */
    SocketIOTaskUpdate,

    /**
     * The SocketIOWorkerUpdate model constructor.
     * @property {module:model/SocketIOWorkerUpdate}
     */
    SocketIOWorkerUpdate,

    /**
     * The SubmittedJob model constructor.
     * @property {module:model/SubmittedJob}
     */
    SubmittedJob,

    /**
     * The Task model constructor.
     * @property {module:model/Task}
     */
    Task,

    /**
     * The TaskStatus model constructor.
     * @property {module:model/TaskStatus}
     */
    TaskStatus,

    /**
     * The TaskStatusChange model constructor.
     * @property {module:model/TaskStatusChange}
     */
    TaskStatusChange,

    /**
     * The TaskSummary model constructor.
     * @property {module:model/TaskSummary}
     */
    TaskSummary,

    /**
     * The TaskUpdate model constructor.
     * @property {module:model/TaskUpdate}
     */
    TaskUpdate,

    /**
     * The TaskWorker model constructor.
     * @property {module:model/TaskWorker}
     */
    TaskWorker,

    /**
     * The Worker model constructor.
     * @property {module:model/Worker}
     */
    Worker,

    /**
     * The WorkerList model constructor.
     * @property {module:model/WorkerList}
     */
    WorkerList,

    /**
     * The WorkerRegistration model constructor.
     * @property {module:model/WorkerRegistration}
     */
    WorkerRegistration,

    /**
     * The WorkerSignOn model constructor.
     * @property {module:model/WorkerSignOn}
     */
    WorkerSignOn,

    /**
     * The WorkerStateChange model constructor.
     * @property {module:model/WorkerStateChange}
     */
    WorkerStateChange,

    /**
     * The WorkerStateChanged model constructor.
     * @property {module:model/WorkerStateChanged}
     */
    WorkerStateChanged,

    /**
     * The WorkerStatus model constructor.
     * @property {module:model/WorkerStatus}
     */
    WorkerStatus,

    /**
     * The WorkerSummary model constructor.
     * @property {module:model/WorkerSummary}
     */
    WorkerSummary,

    /**
    * The JobsApi service constructor.
    * @property {module:manager/JobsApi}
    */
    JobsApi,

    /**
    * The MetaApi service constructor.
    * @property {module:manager/MetaApi}
    */
    MetaApi,

    /**
    * The ShamanApi service constructor.
    * @property {module:manager/ShamanApi}
    */
    ShamanApi,

    /**
    * The WorkerApi service constructor.
    * @property {module:manager/WorkerApi}
    */
    WorkerApi,

    /**
    * The WorkerMgtApi service constructor.
    * @property {module:manager/WorkerMgtApi}
    */
    WorkerMgtApi
};
