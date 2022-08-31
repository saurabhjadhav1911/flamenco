# SPDX-License-Identifier: GPL-3.0-or-later

# <pep8 compliant>

import logging
from typing import TYPE_CHECKING, Optional

from urllib3.exceptions import HTTPError, MaxRetryError
import bpy

_flamenco_client = None
_log = logging.getLogger(__name__)

if TYPE_CHECKING:
    from flamenco.manager import ApiClient as _ApiClient
    from flamenco.manager.models import (
        FlamencoVersion as _FlamencoVersion,
        ManagerConfiguration as _ManagerConfiguration,
    )
    from .preferences import FlamencoPreferences as _FlamencoPreferences
else:
    _ApiClient = object
    _FlamencoPreferences = object
    _FlamencoVersion = object
    _ManagerConfiguration = object


def flamenco_api_client(manager_url: str) -> _ApiClient:
    """Returns an API client for communicating with a Manager."""
    global _flamenco_client

    if _flamenco_client is not None:
        return _flamenco_client

    from . import dependencies

    dependencies.preload_modules()

    from . import manager

    configuration = manager.Configuration(host=manager_url.rstrip("/"))
    _flamenco_client = manager.ApiClient(configuration)
    _log.info("created API client for Manager at %s", manager_url)

    return _flamenco_client


def discard_flamenco_data():
    global _flamenco_client

    if _flamenco_client is None:
        return

    _log.info("closing Flamenco client")
    _flamenco_client.close()
    _flamenco_client = None


def ping_manager_with_report(
    context: bpy.types.Context, api_client: _ApiClient, prefs: _FlamencoPreferences
) -> tuple[str, str]:
    """Ping the Manager, update preferences, and return a report as string.

    :returns: tuple (report, level). The report will be something like "<name>
        version <version> found", or an error message. The level will be
        'ERROR', 'WARNING', or 'INFO', suitable for reporting via
        `Operator.report()`.
    """

    context.window_manager.flamenco_status_ping = "..."

    version, _, err = ping_manager(api_client, prefs)
    if err:
        context.window_manager.flamenco_status_ping = err
        return err, "ERROR"

    assert version is not None
    report = "%s version %s found" % (version.name, version.version)
    context.window_manager.flamenco_status_ping = report
    return report, "INFO"


def ping_manager(
    api_client: _ApiClient, prefs: _FlamencoPreferences
) -> tuple[Optional[_FlamencoVersion], Optional[_ManagerConfiguration], str]:
    """Fetch Manager config & version, and update preferences.

    :returns: tuple (version, config, error).
    """

    # Do a late import, so that the API is only imported when actually used.
    from flamenco.manager import ApiException
    from flamenco.manager.apis import MetaApi
    from flamenco.manager.models import FlamencoVersion, ManagerConfiguration

    meta_api = MetaApi(api_client)
    try:
        version: FlamencoVersion = meta_api.get_version()
        config: ManagerConfiguration = meta_api.get_configuration()
    except ApiException as ex:
        return (None, None, "Manager cannot be reached: %s" % ex)
    except MaxRetryError as ex:
        # This is the common error, when for example the port number is
        # incorrect and nothing is listening. The exception text is not included
        # because it's very long and confusing.
        return (None, None, "Manager cannot be reached")
    except HTTPError as ex:
        return (None, None, "Manager cannot be reached: %s" % ex)

    # Store whether this Manager supports the Shaman API.
    prefs.is_shaman_enabled = config.shaman_enabled
    prefs.job_storage = config.storage_location

    return version, config, ""
