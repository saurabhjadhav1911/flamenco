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
import JobAllOf from './JobAllOf';
import JobStatus from './JobStatus';
import JobStorageInfo from './JobStorageInfo';
import SubmittedJob from './SubmittedJob';

/**
 * The Job model module.
 * @module model/Job
 * @version 0.0.0
 */
class Job {
    /**
     * Constructs a new <code>Job</code>.
     * @alias module:model/Job
     * @implements module:model/SubmittedJob
     * @implements module:model/JobAllOf
     * @param name {String} 
     * @param type {String} 
     * @param priority {Number} 
     * @param submitterPlatform {String} Operating system of the submitter. This is used to recognise two-way variables. This should be a lower-case version of the platform, like \"linux\", \"windows\", \"darwin\", \"openbsd\", etc. Should be ompatible with Go's `runtime.GOOS`; run `go tool dist list` to get a list of possible platforms. As a special case, the platform \"manager\" can be given, which will be interpreted as \"the Manager's platform\". This is mostly to make test/debug scripts easier, as they can use a static document on all platforms. 
     * @param id {String} UUID of the Job
     * @param created {Date} Creation timestamp
     * @param updated {Date} Timestamp of last update.
     * @param status {module:model/JobStatus} 
     * @param activity {String} Description of the last activity on this job.
     */
    constructor(name, type, priority, submitterPlatform, id, created, updated, status, activity) { 
        SubmittedJob.initialize(this, name, type, priority, submitterPlatform);JobAllOf.initialize(this, id, created, updated, status, activity);
        Job.initialize(this, name, type, priority, submitterPlatform, id, created, updated, status, activity);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, priority, submitterPlatform, id, created, updated, status, activity) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['priority'] = priority || 50;
        obj['submitter_platform'] = submitterPlatform;
        obj['id'] = id;
        obj['created'] = created;
        obj['updated'] = updated;
        obj['status'] = status;
        obj['activity'] = activity;
    }

    /**
     * Constructs a <code>Job</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Job} obj Optional instance to populate.
     * @return {module:model/Job} The populated <code>Job</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Job();
            SubmittedJob.constructFromObject(data, obj);
            JobAllOf.constructFromObject(data, obj);

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('type_etag')) {
                obj['type_etag'] = ApiClient.convertToType(data['type_etag'], 'String');
            }
            if (data.hasOwnProperty('priority')) {
                obj['priority'] = ApiClient.convertToType(data['priority'], 'Number');
            }
            if (data.hasOwnProperty('settings')) {
                obj['settings'] = ApiClient.convertToType(data['settings'], {'String': Object});
            }
            if (data.hasOwnProperty('metadata')) {
                obj['metadata'] = ApiClient.convertToType(data['metadata'], {'String': 'String'});
            }
            if (data.hasOwnProperty('submitter_platform')) {
                obj['submitter_platform'] = ApiClient.convertToType(data['submitter_platform'], 'String');
            }
            if (data.hasOwnProperty('storage')) {
                obj['storage'] = JobStorageInfo.constructFromObject(data['storage']);
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('created')) {
                obj['created'] = ApiClient.convertToType(data['created'], 'Date');
            }
            if (data.hasOwnProperty('updated')) {
                obj['updated'] = ApiClient.convertToType(data['updated'], 'Date');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = JobStatus.constructFromObject(data['status']);
            }
            if (data.hasOwnProperty('activity')) {
                obj['activity'] = ApiClient.convertToType(data['activity'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
Job.prototype['name'] = undefined;

/**
 * @member {String} type
 */
Job.prototype['type'] = undefined;

/**
 * Hash of the job type, copied from the `AvailableJobType.etag` property of the job type. The job will be rejected if this field doesn't match the actual job type on the Manager. This prevents job submission with old settings, after the job compiler script has been updated. If this field is ommitted, the check is bypassed. 
 * @member {String} type_etag
 */
Job.prototype['type_etag'] = undefined;

/**
 * @member {Number} priority
 * @default 50
 */
Job.prototype['priority'] = 50;

/**
 * @member {Object.<String, Object>} settings
 */
Job.prototype['settings'] = undefined;

/**
 * Arbitrary metadata strings. More complex structures can be modeled by using `a.b.c` notation for the key.
 * @member {Object.<String, String>} metadata
 */
Job.prototype['metadata'] = undefined;

/**
 * Operating system of the submitter. This is used to recognise two-way variables. This should be a lower-case version of the platform, like \"linux\", \"windows\", \"darwin\", \"openbsd\", etc. Should be ompatible with Go's `runtime.GOOS`; run `go tool dist list` to get a list of possible platforms. As a special case, the platform \"manager\" can be given, which will be interpreted as \"the Manager's platform\". This is mostly to make test/debug scripts easier, as they can use a static document on all platforms. 
 * @member {String} submitter_platform
 */
Job.prototype['submitter_platform'] = undefined;

/**
 * @member {module:model/JobStorageInfo} storage
 */
Job.prototype['storage'] = undefined;

/**
 * UUID of the Job
 * @member {String} id
 */
Job.prototype['id'] = undefined;

/**
 * Creation timestamp
 * @member {Date} created
 */
Job.prototype['created'] = undefined;

/**
 * Timestamp of last update.
 * @member {Date} updated
 */
Job.prototype['updated'] = undefined;

/**
 * @member {module:model/JobStatus} status
 */
Job.prototype['status'] = undefined;

/**
 * Description of the last activity on this job.
 * @member {String} activity
 */
Job.prototype['activity'] = undefined;


// Implement SubmittedJob interface:
/**
 * @member {String} name
 */
SubmittedJob.prototype['name'] = undefined;
/**
 * @member {String} type
 */
SubmittedJob.prototype['type'] = undefined;
/**
 * Hash of the job type, copied from the `AvailableJobType.etag` property of the job type. The job will be rejected if this field doesn't match the actual job type on the Manager. This prevents job submission with old settings, after the job compiler script has been updated. If this field is ommitted, the check is bypassed. 
 * @member {String} type_etag
 */
SubmittedJob.prototype['type_etag'] = undefined;
/**
 * @member {Number} priority
 * @default 50
 */
SubmittedJob.prototype['priority'] = 50;
/**
 * @member {Object.<String, Object>} settings
 */
SubmittedJob.prototype['settings'] = undefined;
/**
 * Arbitrary metadata strings. More complex structures can be modeled by using `a.b.c` notation for the key.
 * @member {Object.<String, String>} metadata
 */
SubmittedJob.prototype['metadata'] = undefined;
/**
 * Operating system of the submitter. This is used to recognise two-way variables. This should be a lower-case version of the platform, like \"linux\", \"windows\", \"darwin\", \"openbsd\", etc. Should be ompatible with Go's `runtime.GOOS`; run `go tool dist list` to get a list of possible platforms. As a special case, the platform \"manager\" can be given, which will be interpreted as \"the Manager's platform\". This is mostly to make test/debug scripts easier, as they can use a static document on all platforms. 
 * @member {String} submitter_platform
 */
SubmittedJob.prototype['submitter_platform'] = undefined;
/**
 * @member {module:model/JobStorageInfo} storage
 */
SubmittedJob.prototype['storage'] = undefined;
// Implement JobAllOf interface:
/**
 * UUID of the Job
 * @member {String} id
 */
JobAllOf.prototype['id'] = undefined;
/**
 * Creation timestamp
 * @member {Date} created
 */
JobAllOf.prototype['created'] = undefined;
/**
 * Timestamp of last update.
 * @member {Date} updated
 */
JobAllOf.prototype['updated'] = undefined;
/**
 * @member {module:model/JobStatus} status
 */
JobAllOf.prototype['status'] = undefined;
/**
 * Description of the last activity on this job.
 * @member {String} activity
 */
JobAllOf.prototype['activity'] = undefined;




export default Job;

