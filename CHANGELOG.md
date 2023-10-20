# Flamenco Change Log

This file contains the history of changes to Flamenco. Only changes that might
be interesting for users are listed here, such as new features and fixes for
bugs in actually-released versions.

## 3.3 - in development

- Add Worker Tag support. Workers can be members of any number of tags. Workers will only work on jobs that are assigned to that tag. Jobs that do not have a tag will be available to all workers, regardless of their tag assignment. As a result, tagless workers will only work on tagless jobs.
- Add support for finding the top-level 'project' directory. When submitting files to Flamenco, the add-on will try to retain the directory structure of your Blender project as precisely as possible. This new feature allows the add-on to find the top-level directory of your project by finding a `.blender_project`, `.git`, or `.subversion` directory. This can be configured in the add-on preferences.
- Worker status is remembered when they sign off, so that workers when they come back online do so to the same state ([#99549](https://projects.blender.org/studio/flamenco/issues/99549)).
- Job settings: make it possible for a setting to be "linked" to its automatic value. For job settings that have this new feature enabled, they will not be editable by default, and the setting will just use its `eval` expression to determine the value. This can be toggled by the user in Blender's submission interface, to still allow manual edits of the value when needed.
- Job settings: add a description for the `eval` field. This is shown in the tooltip of the 'set to automatic value' button, to make it clear what that button will do.
- Database integrity tests. These are always run at startup of Flamenco Manager, and by default run periodically every hour. This can be configured by adding/changing the `database_check_period: 1h` setting in `flamenco-manager.yaml`. Setting it to `0` will disable the periodic check. When a database consistency error is found, Flamenco Manager will immediately shut down.
- Workers can be marked as 'restartable' by using the `-restart-exit-code N` commandline option. More info in the [Worker Actions documentation](https://flamenco.blender.org/usage/worker-actions/).
- Upgrade bundled FFmpeg from 5.0 to 5.1.
- Rename the add-on download to `flamenco-addon.zip` (it used to be `flamenco3-addon.zip`). It still contains the same files as before, and in Blender the name of the add-on has not changed.
- Improve speed of queueing up >100 simultaneous job deletions.
- Improve logging of job deletion.
- Fix limitation where a job could have no more than 1000 tasks ([#104201](https://projects.blender.org/studio/flamenco/issues/104201))
- Nicer version display for non-release builds. Instead of `3.3-alpha0-v3.2-76-gdd34d538`, show `3.3-alpha0 (v3.2-76-gdd34d538)`.
- The webapp automatically reloads after a disconnect, when it reconnects to Flamenco Manager and sees the Manager version changed [#104235](https://projects.blender.org/studio/flamenco/pulls/104235).
- Show the configured Flamenco Manager name in the webapp's browser window title.
- The `{timestamp}` placeholder in the render output path is now replaced with a local timestamp (rather than UTC).
- Log more information about the operating system at startup. On Windows this includes the version & edition (like "Core" or "Professional"), and on Linux this includes the distribution, version, and kernel version.
- Security updates of some dependencies:
  - https://pkg.go.dev/vuln/GO-2023-1989
  - https://pkg.go.dev/vuln/GO-2023-1990
  - https://pkg.go.dev/vuln/GO-2023-2102


## 3.2 - released 2023-02-21

- When rendering EXR files, use Blender's preview JPEG files to generate the preview video ([43bc22f10fae](https://developer.blender.org/rF43bc22f10fae0fcaed6a4a3b3ace1be617193e21)).
- Fix issue where workers would switch immediately on a state change request, even if it was of the "after task is finished" kind.
- Add-on: Do a "pre-submission check" before sending files to the farm. This should provide submission errors earlier in the process, without waiting for files to be collected.
- Worker: better handling of long lines from Blender/FFmpeg, splitting them up at character boundaries.
- Manager: change SQLite parameters to have write-through-log journalling and less filesystem synchronisation. This reduces I/O load on the Manager.
- Add-on: Set Blender's experimental flag `use_all_linked_data_direct` to `True` on submitted files, to work around a shortcoming in BAT. See [Blender commit b8c7e93a6504](https://developer.blender.org/rBb8c7e93a6504833ee1e617523dfe2921c4fd0816) for the introduction of that flag.
- Bump the bundled Blender Asset Tracer (BAT) to version 1.15.
- Increase preview image file size from 10 MB to 25 MB. Even though the Worker can down-scale render output before sending to the Manager as preview, they could still be larger than the limit of 10 MB.
- Fix a crash of the Manager when using an invalid frame range (`1 10` for example, instead of `1-10` or `1,10`)
- Make it possible to delete jobs. The job and its tasks are removed from Flamenco, including last-rendered images and logs. The input files (i.e. the to-be-rendered blend files and their dependencies) will only be removed if [the Shaman system](https://flamenco.blender.org/usage/shared-storage/shaman/) was used AND if the job was submitted with Flamenco 3.2 or newer.
- Security updates of some dependencies:
    - https://pkg.go.dev/vuln/GO-2022-1031
    - https://pkg.go.dev/vuln/GO-2023-1571
    - https://pkg.go.dev/vuln/GO-2023-1572


## 3.1 - released 2022-10-18

- Web interface: make the worker IP address clickable; it will be copied to the clipboard when clicked ([50ec5f4f360c](https://developer.blender.org/rF50ec5f4f360ce7cb467f95de31a34200f4942047)).
- Allow changing the priority of an existing job (API: [07f0b38e8a9f](https://developer.blender.org/rF07f0b38e8a9f0e7ea303adc2608ae5265ec7e075), [85d53de1f99f](https://developer.blender.org/rF85d53de1f99f0ccb904dc7c140a75bf4b96b326b), Web: [4389b60197a0](https://developer.blender.org/rF4389b60197a07c9b64b63f1d111679a3104ab60a), [080a63df6a5b](https://developer.blender.org/rF080a63df6a5b1a95e05eeea3c66d3a41fa431e82)).
- Fix FFmpeg packaging issue, which caused the Worker to not find the bundled FFmpeg executable ([1abeb71f570f](https://developer.blender.org/rF1abeb71f570ff978c2ff81bf6fd9851b86cc7be7)).
- Less dramatic logging when Blender cannot be found by the Worker on startup ([161a7f7cb381](https://developer.blender.org/rF161a7f7cb38190bd34757e74ffc22ac0e068fa5f), [759a94e49b21](https://developer.blender.org/rF759a94e49b21b32405237be978146a826dd53a73)).
  This just means that the Manager has to tell the Worker which Blender to use, which is perfectly fine.
- Fix error in sleep scheduler when shutting down the Manager ([59655ea770f6](https://developer.blender.org/rF59655ea770f667a579e7a85cf3afc7d8b33d239e)).
- Workers can now decode TIFF files to generate previews ([a95e8781cf94](https://developer.blender.org/rFa95e8781cf94663b3d6a41745c102586e066bb85)).
- Fix error submitting to Shaman storage from a Windows machine ([0bc0a7ac9b68](https://developer.blender.org/rF0bc0a7ac9b688d1174862e568f327053d05427b4)).


## 3.0 - released 2022-09-12

- Faster & more accurate progress reporting of file submission.
- Add-on: report which files were missing after submitting a job. This is reported in the terminal (aka System Console).


## 3.0-beta3 - released 2022-08-31

- Clean up how version numbers are reported, so that there are no repeats of the
  version (beta2 was reported as `3.0-beta2-v3.0-beta2`).
- Fix an issue running FFmpeg.
- The "Simple Blender Render" job type no longer accepts files that render to
  video (so FFmpeg or one of the built-in AVI options). This was originally
  intended to work, but had various problems. Now the script actively refuses to
  handle such files, and limits itself to images only. It will still create a
  preview video out of these images.
- The "Simple Blender Render" job type no longer renders to an intermediate
  directory. It simply always renders to the configured path. Not only does this
  simplify the script, but it also makes it possible to allow selective
  rerendering in the future.


## 3.0-beta2 - released 2022-08-31

WARNING: this version is backward incompatible. Any job created with Flamenco
3.0-beta1 will not run with Flamenco 3.0-beta2. Only upgrade after
currently-active jobs have finished, or cancel them.

It is recommended to remove `flamenco-manager.yaml`, restart Flamenco Manager,
and reconfigure via the setup assistant.

- Manager & Add-on: avoid error that could occur when submitting jobs with UDIM files
  ([44ccc6c3ca70](https://developer.blender.org/rF44ccc6c3ca706fdd268bf310f3e8965d58482449)).
- Manager: don't stop when the Flamenco Setup Assistant cannot start a webbrowser
  ([7d3d3d1d6078](https://developer.blender.org/rF7d3d3d1d6078828122b4b2d1376b1aaf2ba03b8b)).
- Change path inside the Linux and macOS tarballs, so that they contain an
  embedded `flamenco-3.x.y-xxxx/` directory with all the files (instead of
  putting all the files in the root of the tarball).
- Two-way variable replacement now also changes the path separators to the target platform.
- Allow setting priority when submitting a job
  ([db9aca4a37e1](https://developer.blender.org/rFdb9aca4a37e1be37f802cb609fddab4308e5e40f)).
- Separate "blender location" and "blender arguments" into two variables
  ([e5a20425c474](https://developer.blender.org/rFe5a20425c474ec93edbe03d2667ec5184f32d3ef)).
  - The variable `blender` now should only point at the Blender executable, for
    example `D:\Blender_3.2_stable\blender.exe`.
  - The variable `blenderArgs` can be used to set the default Blender arguments,
    for example `-b -y`.
- Job storage location can now be made multi-platform by using two-way variables
  ([31cf0a4ecc75](https://developer.blender.org/rF31cf0a4ecc75db127877218af449610ce9d8df1c)).

## 3.0-beta1 - released 2022-08-03

This was the first version of Flamenco to be released to the public, and thus it
serves as the starting point for this change log.
