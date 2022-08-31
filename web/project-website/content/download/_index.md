---
title: Download
---

Download Flamenco for your platform here. Each download contains both Flamenco Manager and Worker.

The latest version is: **{{< flamenco/latestVersion >}}**

{{< hint type=caution >}}
This is a **beta** version of the software! This means that the features for the
3.0 version are there, but are likely to be some bugs.

When **upgrading** from a previous beta, it is recommended to start afresh with the following steps:

1. Cancel any running or queued job.
2. Shut down Flamenco Manager and all Workers.
3. Remove `flamenco-manager.yaml` and `flamenco-manager.sqlite`.
4. Download the new version and replace your old Flamenco files with the new ones.
5. Start `flamenco-manager` and go through the setup setup assistant again.
6. Don't forget to re-download the Blender add-on from the Manager's web
   interface, and install it. It has seen development as well, and will need to
   be upgraded.

Read the announcement at the [Blender Studio blog][blog]. <br>
Please report any issue at [developer.blender.org][bugs]. A stable release is
planned by the end of September 2022.

[blog]: (https://studio.blender.org/blog/announcing-flamenco-3-beta/)
[bugs]: (https://developer.blender.org/project/profile/58/)
{{< /hint >}}

| Platform | File                                                 |
|----------|------------------------------------------------------|
| Windows  | {{< flamenco/downloadLink os="windows" ext="zip" >}} |
| Linux    | {{< flamenco/downloadLink os="linux" >}}             |
| macOS    | {{< flamenco/downloadLink os="macos" >}}             |


## License

Flamenco is Free and Open Source software, available under the
[GNU General Public License](https://developer.blender.org/diffusion/F/browse/main/LICENSE).<br>
Download the source code at [developer.blender.org](https://developer.blender.org/diffusion/F/).
