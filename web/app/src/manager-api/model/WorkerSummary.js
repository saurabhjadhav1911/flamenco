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
import WorkerStatusChangeRequest from './WorkerStatusChangeRequest';

/**
 * The WorkerSummary model module.
 * @module model/WorkerSummary
 * @version 0.0.0
 */
class WorkerSummary {
    /**
     * Constructs a new <code>WorkerSummary</code>.
     * Basic information about a Worker.
     * @alias module:model/WorkerSummary
     * @param id {String} 
     * @param name {String} 
     * @param status {module:model/WorkerStatus} 
     * @param version {String} Version of Flamenco this Worker is running
     */
    constructor(id, name, status, version) { 
        
        WorkerSummary.initialize(this, id, name, status, version);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, name, status, version) { 
        obj['id'] = id;
        obj['name'] = name;
        obj['status'] = status;
        obj['version'] = version;
    }

    /**
     * Constructs a <code>WorkerSummary</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WorkerSummary} obj Optional instance to populate.
     * @return {module:model/WorkerSummary} The populated <code>WorkerSummary</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new WorkerSummary();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = WorkerStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('status_change')) {
                obj['status_change'] = WorkerStatusChangeRequest.constructFromObject(data['status_change']);
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} id
 */
WorkerSummary.prototype['id'] = undefined;

/**
 * @member {String} name
 */
WorkerSummary.prototype['name'] = undefined;

/**
 * @member {module:model/WorkerStatus} status
 */
WorkerSummary.prototype['status'] = undefined;

/**
 * @member {module:model/WorkerStatusChangeRequest} status_change
 */
WorkerSummary.prototype['status_change'] = undefined;

/**
 * Version of Flamenco this Worker is running
 * @member {String} version
 */
WorkerSummary.prototype['version'] = undefined;






export default WorkerSummary;

