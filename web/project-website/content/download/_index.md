---
title: Download
weight: 2
---

Download Flamenco for your platform here. Each download contains both Flamenco
Manager and Worker. The Blender add-on can be downloaded from the Flamenco
Manager web-interface after installation.

The latest version is: **{{< flamenco/latestVersion >}}**

| Platform  | File                                                 |
|-----------|------------------------------------------------------|
| Windows   | {{< flamenco/downloadLink os="windows" ext="zip" >}} |
| Linux     | {{< flamenco/downloadLink os="linux" >}}             |
| macOS     | {{< flamenco/downloadLink os="macos" >}}             |
| checksums | {{< flamenco/sha256link >}}                          |


{{< hint type=caution >}}
When **upgrading** from a previous v3 beta version, it is recommended to start
afresh with the following steps:

1. Cancel any running or queued job.
2. Shut down Flamenco Manager and all Workers.
3. Remove `flamenco-manager.yaml` and `flamenco-manager.sqlite`.
4. Download the new version and replace your old Flamenco files with the new ones.
5. Start `flamenco-manager` and go through the setup setup assistant again.
6. Don't forget to re-download the Blender add-on from the Manager's web
   interface, and install it. It has seen development as well, and will need to
   be upgraded.

Read the beta announcement at the [Blender Studio blog][blog].<br>
Please report any issue at [project.blender.org][bugs].

[blog]: https://studio.blender.org/blog/announcing-flamenco-3-beta/
[bugs]: https://projects.blender.org/studio/flamenco/issues/new?template=.gitea%2fissue_template%2fbug.yaml
{{< /hint >}}


## License

Flamenco is Free and Open Source software, available under the
[GNU General Public License](https://projects.blender.org/studio/flamenco/src/branch/main/LICENSE).<br>
Download the source code at [projects.blender.org](https://projects.blender.org/studio/flamenco).
