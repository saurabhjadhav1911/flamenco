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
 * The SocketIOWorkerUpdate model module.
 * @module model/SocketIOWorkerUpdate
 * @version 0.0.0
 */
class SocketIOWorkerUpdate {
    /**
     * Constructs a new <code>SocketIOWorkerUpdate</code>.
     * Subset of a Worker, sent over SocketIO when a worker changes. For new workers, &#x60;previous_status&#x60; will be excluded. 
     * @alias module:model/SocketIOWorkerUpdate
     * @param id {String} UUID of the Worker
     * @param nickname {String} Name of the worker
     * @param updated {Date} Timestamp of last update
     * @param status {module:model/WorkerStatus} 
     * @param version {String} 
     */
    constructor(id, nickname, updated, status, version) { 
        
        SocketIOWorkerUpdate.initialize(this, id, nickname, updated, status, version);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, nickname, updated, status, version) { 
        obj['id'] = id;
        obj['nickname'] = nickname;
        obj['updated'] = updated;
        obj['status'] = status;
        obj['version'] = version;
    }

    /**
     * Constructs a <code>SocketIOWorkerUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SocketIOWorkerUpdate} obj Optional instance to populate.
     * @return {module:model/SocketIOWorkerUpdate} The populated <code>SocketIOWorkerUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SocketIOWorkerUpdate();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('nickname')) {
                obj['nickname'] = ApiClient.convertToType(data['nickname'], 'String');
            }
            if (data.hasOwnProperty('updated')) {
                obj['updated'] = ApiClient.convertToType(data['updated'], 'Date');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = WorkerStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('previous_status')) {
                obj['previous_status'] = WorkerStatus.constructFromObject(data['previous_status']);
            }
            if (data.hasOwnProperty('status_requested')) {
                obj['status_requested'] = WorkerStatus.constructFromObject(data['status_requested']);
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
        }
        return obj;
    }


}

/**
 * UUID of the Worker
 * @member {String} id
 */
SocketIOWorkerUpdate.prototype['id'] = undefined;

/**
 * Name of the worker
 * @member {String} nickname
 */
SocketIOWorkerUpdate.prototype['nickname'] = undefined;

/**
 * Timestamp of last update
 * @member {Date} updated
 */
SocketIOWorkerUpdate.prototype['updated'] = undefined;

/**
 * @member {module:model/WorkerStatus} status
 */
SocketIOWorkerUpdate.prototype['status'] = undefined;

/**
 * @member {module:model/WorkerStatus} previous_status
 */
SocketIOWorkerUpdate.prototype['previous_status'] = undefined;

/**
 * @member {module:model/WorkerStatus} status_requested
 */
SocketIOWorkerUpdate.prototype['status_requested'] = undefined;

/**
 * @member {String} version
 */
SocketIOWorkerUpdate.prototype['version'] = undefined;






export default SocketIOWorkerUpdate;
