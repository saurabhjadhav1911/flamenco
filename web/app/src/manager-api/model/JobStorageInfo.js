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
 * The JobStorageInfo model module.
 * @module model/JobStorageInfo
 * @version 0.0.0
 */
class JobStorageInfo {
    /**
     * Constructs a new <code>JobStorageInfo</code>.
     * Storage info of a job, which Flamenco can use to remove job-related files when necessary. 
     * @alias module:model/JobStorageInfo
     */
    constructor() { 
        
        JobStorageInfo.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>JobStorageInfo</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/JobStorageInfo} obj Optional instance to populate.
     * @return {module:model/JobStorageInfo} The populated <code>JobStorageInfo</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new JobStorageInfo();

            if (data.hasOwnProperty('shaman_checkout_id')) {
                obj['shaman_checkout_id'] = ApiClient.convertToType(data['shaman_checkout_id'], 'String');
            }
        }
        return obj;
    }


}

/**
 * 'Checkout ID' used when creating the Shaman checkout for this job. Aids in removing the checkout directory when the job is removed from Flamenco. 
 * @member {String} shaman_checkout_id
 */
JobStorageInfo.prototype['shaman_checkout_id'] = undefined;






export default JobStorageInfo;

