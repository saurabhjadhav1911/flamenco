"""
    Flamenco manager

    Render Farm manager API  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import flamenco.manager
from flamenco.manager.model.job_all_of import JobAllOf
from flamenco.manager.model.job_metadata import JobMetadata
from flamenco.manager.model.job_settings import JobSettings
from flamenco.manager.model.job_status import JobStatus
from flamenco.manager.model.submitted_job import SubmittedJob
globals()['JobAllOf'] = JobAllOf
globals()['JobMetadata'] = JobMetadata
globals()['JobSettings'] = JobSettings
globals()['JobStatus'] = JobStatus
globals()['SubmittedJob'] = SubmittedJob
from flamenco.manager.model.job import Job


class TestJob(unittest.TestCase):
    """Job unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testJob(self):
        """Test Job"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Job()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()