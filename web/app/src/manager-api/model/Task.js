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

import ApiClient from '../ApiClient';
import Command from './Command';
import TaskStatus from './TaskStatus';
import TaskWorker from './TaskWorker';

/**
 * The Task model module.
 * @module model/Task
 * @version 0.0.0
 */
class Task {
    /**
     * Constructs a new <code>Task</code>.
     * The task as it exists in the Manager database, i.e. before variable replacement.
     * @alias module:model/Task
     * @param id {String} 
     * @param created {Date} Creation timestamp
     * @param updated {Date} Timestamp of last update.
     * @param jobId {String} 
     * @param name {String} 
     * @param status {module:model/TaskStatus} 
     * @param priority {Number} 
     * @param taskType {String} 
     * @param activity {String} 
     * @param commands {Array.<module:model/Command>} 
     */
    constructor(id, created, updated, jobId, name, status, priority, taskType, activity, commands) { 
        
        Task.initialize(this, id, created, updated, jobId, name, status, priority, taskType, activity, commands);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, created, updated, jobId, name, status, priority, taskType, activity, commands) { 
        obj['id'] = id;
        obj['created'] = created;
        obj['updated'] = updated;
        obj['job_id'] = jobId;
        obj['name'] = name;
        obj['status'] = status;
        obj['priority'] = priority;
        obj['task_type'] = taskType;
        obj['activity'] = activity;
        obj['commands'] = commands;
    }

    /**
     * Constructs a <code>Task</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Task} obj Optional instance to populate.
     * @return {module:model/Task} The populated <code>Task</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Task();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('created')) {
                obj['created'] = ApiClient.convertToType(data['created'], 'Date');
            }
            if (data.hasOwnProperty('updated')) {
                obj['updated'] = ApiClient.convertToType(data['updated'], 'Date');
            }
            if (data.hasOwnProperty('job_id')) {
                obj['job_id'] = ApiClient.convertToType(data['job_id'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = TaskStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('priority')) {
                obj['priority'] = ApiClient.convertToType(data['priority'], 'Number');
            }
            if (data.hasOwnProperty('task_type')) {
                obj['task_type'] = ApiClient.convertToType(data['task_type'], 'String');
            }
            if (data.hasOwnProperty('activity')) {
                obj['activity'] = ApiClient.convertToType(data['activity'], 'String');
            }
            if (data.hasOwnProperty('commands')) {
                obj['commands'] = ApiClient.convertToType(data['commands'], [Command]);
            }
            if (data.hasOwnProperty('worker')) {
                obj['worker'] = TaskWorker.constructFromObject(data['worker']);
            }
            if (data.hasOwnProperty('last_touched')) {
                obj['last_touched'] = ApiClient.convertToType(data['last_touched'], 'Date');
            }
            if (data.hasOwnProperty('failed_by_workers')) {
                obj['failed_by_workers'] = ApiClient.convertToType(data['failed_by_workers'], [TaskWorker]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
Task.prototype['id'] = undefined;

/**
 * Creation timestamp
 * @member {Date} created
 */
Task.prototype['created'] = undefined;

/**
 * Timestamp of last update.
 * @member {Date} updated
 */
Task.prototype['updated'] = undefined;

/**
 * @member {String} job_id
 */
Task.prototype['job_id'] = undefined;

/**
 * @member {String} name
 */
Task.prototype['name'] = undefined;

/**
 * @member {module:model/TaskStatus} status
 */
Task.prototype['status'] = undefined;

/**
 * @member {Number} priority
 */
Task.prototype['priority'] = undefined;

/**
 * @member {String} task_type
 */
Task.prototype['task_type'] = undefined;

/**
 * @member {String} activity
 */
Task.prototype['activity'] = undefined;

/**
 * @member {Array.<module:model/Command>} commands
 */
Task.prototype['commands'] = undefined;

/**
 * @member {module:model/TaskWorker} worker
 */
Task.prototype['worker'] = undefined;

/**
 * Timestamp of when any worker worked on this task.
 * @member {Date} last_touched
 */
Task.prototype['last_touched'] = undefined;

/**
 * @member {Array.<module:model/TaskWorker>} failed_by_workers
 */
Task.prototype['failed_by_workers'] = undefined;






export default Task;

