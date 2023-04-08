# SPDX-License-Identifier: GPL-3.0-or-later

from typing import TYPE_CHECKING, Union

import bpy

from . import preferences

if TYPE_CHECKING:
    from flamenco.manager import ApiClient as _ApiClient
else:
    _ApiClient = object


_enum_items: list[Union[tuple[str, str, str], tuple[str, str, str, int, int]]] = []


def refresh(context: bpy.types.Context, api_client: _ApiClient) -> None:
    """Fetch the available worker clusters from the Manager."""
    from flamenco.manager import ApiClient
    from flamenco.manager.api import worker_mgt_api
    from flamenco.manager.model.worker_cluster_list import WorkerClusterList

    assert isinstance(api_client, ApiClient)

    api = worker_mgt_api.WorkerMgtApi(api_client)
    response: WorkerClusterList = api.fetch_worker_clusters()

    # Store on the preferences, so a cached version persists until the next refresh.
    prefs = preferences.get(context)
    prefs.worker_clusters.clear()

    for cluster in response.clusters:
        rna_cluster = prefs.worker_clusters.add()
        rna_cluster.id = cluster.id
        rna_cluster.name = cluster.name
        rna_cluster.description = getattr(cluster, "description", "")


def _get_enum_items(self, context):
    global _enum_items
    prefs = preferences.get(context)

    _enum_items = [
        ("-", "All", "No specific cluster assigned, any worker can handle this job"),
    ]
    _enum_items.extend(
        (cluster.id, cluster.name, cluster.description)
        for cluster in prefs.worker_clusters
    )
    return _enum_items


def register() -> None:
    bpy.types.Scene.flamenco_worker_cluster = bpy.props.EnumProperty(
        name="Worker Cluster",
        items=_get_enum_items,
        description="The set of Workers that can handle tasks of this job",
    )


def unregister() -> None:
    to_del = ((bpy.types.Scene, "flamenco_worker_cluster"),)
    for ob, attr in to_del:
        try:
            delattr(ob, attr)
        except AttributeError:
            pass


if __name__ == "__main__":
    import doctest

    print(doctest.testmod())
