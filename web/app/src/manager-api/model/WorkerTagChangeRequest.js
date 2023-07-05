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
 * The WorkerTagChangeRequest model module.
 * @module model/WorkerTagChangeRequest
 * @version 0.0.0
 */
class WorkerTagChangeRequest {
    /**
     * Constructs a new <code>WorkerTagChangeRequest</code>.
     * Request to change which tags this Worker is assigned to.
     * @alias module:model/WorkerTagChangeRequest
     * @param tagIds {Array.<String>} 
     */
    constructor(tagIds) { 
        
        WorkerTagChangeRequest.initialize(this, tagIds);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, tagIds) { 
        obj['tag_ids'] = tagIds;
    }

    /**
     * Constructs a <code>WorkerTagChangeRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WorkerTagChangeRequest} obj Optional instance to populate.
     * @return {module:model/WorkerTagChangeRequest} The populated <code>WorkerTagChangeRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new WorkerTagChangeRequest();

            if (data.hasOwnProperty('tag_ids')) {
                obj['tag_ids'] = ApiClient.convertToType(data['tag_ids'], ['String']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<String>} tag_ids
 */
WorkerTagChangeRequest.prototype['tag_ids'] = undefined;






export default WorkerTagChangeRequest;
