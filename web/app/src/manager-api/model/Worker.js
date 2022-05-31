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
import WorkerStatus from './WorkerStatus';

/**
 * The Worker model module.
 * @module model/Worker
 * @version 0.0.0
 */
class Worker {
    /**
     * Constructs a new <code>Worker</code>.
     * All information about a Worker
     * @alias module:model/Worker
     * @param id {String} 
     * @param nickname {String} 
     * @param status {module:model/WorkerStatus} 
     * @param ipAddress {String} IP address of the Worker
     * @param platform {String} Operating system of the Worker
     * @param version {String} Version of Flamenco this Worker is running
     * @param supportedTaskTypes {Array.<String>} 
     */
    constructor(id, nickname, status, ipAddress, platform, version, supportedTaskTypes) { 
        
        Worker.initialize(this, id, nickname, status, ipAddress, platform, version, supportedTaskTypes);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, nickname, status, ipAddress, platform, version, supportedTaskTypes) { 
        obj['id'] = id;
        obj['nickname'] = nickname;
        obj['status'] = status;
        obj['ip_address'] = ipAddress;
        obj['platform'] = platform;
        obj['version'] = version;
        obj['supported_task_types'] = supportedTaskTypes;
    }

    /**
     * Constructs a <code>Worker</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Worker} obj Optional instance to populate.
     * @return {module:model/Worker} The populated <code>Worker</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Worker();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('nickname')) {
                obj['nickname'] = ApiClient.convertToType(data['nickname'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = WorkerStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('status_requested')) {
                obj['status_requested'] = WorkerStatus.constructFromObject(data['status_requested']);
            }
            if (data.hasOwnProperty('ip_address')) {
                obj['ip_address'] = ApiClient.convertToType(data['ip_address'], 'String');
            }
            if (data.hasOwnProperty('platform')) {
                obj['platform'] = ApiClient.convertToType(data['platform'], 'String');
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
            if (data.hasOwnProperty('supported_task_types')) {
                obj['supported_task_types'] = ApiClient.convertToType(data['supported_task_types'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
Worker.prototype['id'] = undefined;

/**
 * @member {String} nickname
 */
Worker.prototype['nickname'] = undefined;

/**
 * @member {module:model/WorkerStatus} status
 */
Worker.prototype['status'] = undefined;

/**
 * @member {module:model/WorkerStatus} status_requested
 */
Worker.prototype['status_requested'] = undefined;

/**
 * IP address of the Worker
 * @member {String} ip_address
 */
Worker.prototype['ip_address'] = undefined;

/**
 * Operating system of the Worker
 * @member {String} platform
 */
Worker.prototype['platform'] = undefined;

/**
 * Version of Flamenco this Worker is running
 * @member {String} version
 */
Worker.prototype['version'] = undefined;

/**
 * @member {Array.<String>} supported_task_types
 */
Worker.prototype['supported_task_types'] = undefined;






export default Worker;
