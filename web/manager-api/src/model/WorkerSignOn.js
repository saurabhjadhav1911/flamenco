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

/**
 * The WorkerSignOn model module.
 * @module model/WorkerSignOn
 * @version 0.0.0
 */
class WorkerSignOn {
    /**
     * Constructs a new <code>WorkerSignOn</code>.
     * @alias module:model/WorkerSignOn
     * @param nickname {String} 
     * @param supported_task_types {Array.<String>} 
     * @param software_version {String} 
     */
    constructor(nickname, supported_task_types, software_version) { 
        
        WorkerSignOn.initialize(this, nickname, supported_task_types, software_version);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, nickname, supported_task_types, software_version) { 
        obj['nickname'] = nickname;
        obj['supported_task_types'] = supported_task_types;
        obj['software_version'] = software_version;
    }

    /**
     * Constructs a <code>WorkerSignOn</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WorkerSignOn} obj Optional instance to populate.
     * @return {module:model/WorkerSignOn} The populated <code>WorkerSignOn</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new WorkerSignOn();

            if (data.hasOwnProperty('nickname')) {
                obj['nickname'] = ApiClient.convertToType(data['nickname'], 'String');
            }
            if (data.hasOwnProperty('supported_task_types')) {
                obj['supported_task_types'] = ApiClient.convertToType(data['supported_task_types'], ['String']);
            }
            if (data.hasOwnProperty('software_version')) {
                obj['software_version'] = ApiClient.convertToType(data['software_version'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} nickname
 */
WorkerSignOn.prototype['nickname'] = undefined;

/**
 * @member {Array.<String>} supported_task_types
 */
WorkerSignOn.prototype['supported_task_types'] = undefined;

/**
 * @member {String} software_version
 */
WorkerSignOn.prototype['software_version'] = undefined;






export default WorkerSignOn;
