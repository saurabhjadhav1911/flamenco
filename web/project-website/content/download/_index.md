---
title: Download
weight: 2
geekdocNav: false
geekdocHidden: true
---

Download Flamenco for your platform here. Each download contains both Flamenco
Manager and Worker. The Blender add-on can be downloaded from the Flamenco
Manager web-interface after installation.

The latest version is: **{{< flamenco/latestVersion >}}**

| Platform      | File                                                 |
|---------------|------------------------------------------------------|
| Windows       | {{< flamenco/downloadLink os="windows" ext="zip" >}} |
| Linux         | {{< flamenco/downloadLink os="linux" >}}             |
| macOS (Intel) | {{< flamenco/downloadLink os="macos" >}}             |
| checksums     | {{< flamenco/sha256link >}}                          |

Please report any issue at [projects.blender.org][bugs].

## Go Experimental!

The latest experimental version is: **{{< flamenco/latestExperimentalVersion
>}}**. This version is not guaranteed to be stable, so do not run this on
production systems. Or at least make a backup of your `flamenco-manager.yaml`
and `flamenco-manager.sqlite` files before you venture forth.

To see what's new, check [the changelog](https://projects.blender.org/studio/flamenco/src/branch/main/CHANGELOG.md).

| Platform            | File                                                                                                     |
|---------------------|----------------------------------------------------------------------------------------------------------|
| Windows             | {{< flamenco/downloadExperimentalLink os="windows" ext="zip" >}}                                         |
| Linux               | {{< flamenco/downloadExperimentalLink os="linux" >}}                                                     |
| macOS (Intel)       | {{< flamenco/downloadExperimentalLink os="macos" >}}                                                     |
| macOS (Silicon/ARM) | {{< flamenco/downloadExperimentalLink os="macos" arch="arm64" >}} (without FFmpeg, [see below][mac-arm]) |
| checksums           | {{< flamenco/sha256linkExperimental >}}                                                                  |

Please report any issue at [projects.blender.org][bugs].

[bugs]: https://projects.blender.org/studio/flamenco/issues/new?template=.gitea%2fissue_template%2fbug.yaml
[mac-arm]: #macos-silicon-builds

<!--

{{< hint type=caution >}}
When **upgrading** from a previous experimental version, it is recommended to
start afresh with the following steps:

1. Cancel any running or queued job.
2. Shut down Flamenco Manager and all Workers.
3. Remove `flamenco-manager.yaml` and `flamenco-manager.sqlite`.
4. Download the new version and replace your old Flamenco files with the new ones.
5. Start `flamenco-manager` and go through the setup setup assistant again.
6. Don't forget to re-download the Blender add-on from the Manager's web
   interface, and install it. It has seen development as well, and will need to
   be upgraded.

[blog]: https://studio.blender.org/blog/announcing-flamenco-3-beta/
{{< /hint >}}
-->

## macOS "Silicon" builds

The FFmpeg binaries that are bundled with Flamenco are linked directly from
[ffmpeg.org](https://ffmpeg.org/), and thus there is a level of trust. That
website, however, does not link directly to any trusted binaries for macOS
"Silicon" ARM64. This is why Flamenco did not ship macOS/ARM64 builds for a long
time. As of v3.3 this architecture will be included in the official Flamenco
builds, but for now will be shipped without FFmpeg binary.

To install this binary, download it from a source you trust or build it
yourself, then place it into Flamenco's `tools` directory.

## License

Flamenco is Free and Open Source software, available under the
[GNU General Public License](https://projects.blender.org/studio/flamenco/src/branch/main/LICENSE).<br>
Download the source code at [projects.blender.org](https://projects.blender.org/studio/flamenco).
