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
import WorkerTag from './WorkerTag';

/**
 * The WorkerTagList model module.
 * @module model/WorkerTagList
 * @version 0.0.0
 */
class WorkerTagList {
    /**
     * Constructs a new <code>WorkerTagList</code>.
     * @alias module:model/WorkerTagList
     */
    constructor() { 
        
        WorkerTagList.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>WorkerTagList</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WorkerTagList} obj Optional instance to populate.
     * @return {module:model/WorkerTagList} The populated <code>WorkerTagList</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new WorkerTagList();

            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [WorkerTag]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/WorkerTag>} tags
 */
WorkerTagList.prototype['tags'] = undefined;






export default WorkerTagList;

