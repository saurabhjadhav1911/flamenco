# Flamenco Change Log

This file contains the history of changes to Flamenco. Only changes that might
be interesting for users are listed here, such as new features and fixes for
bugs in actually-released versions.

## 3.1 - in development

- Web interface: make the worker IP address clickable; it will be copied to the clipboard when clicked.
- Add API operation to change the priority of an existing job.
- Fix FFmpeg packaging issue, which caused the Worker to not find the bundled FFmpeg executable.
- Less dramatic logging when Blender cannot be found by the Worker on startup.
  This just means that the Manager has to tell the Worker which Blender to use,
  which is perfectly fine.
- Fix error in sleep scheduler when shutting down the Manager.
- Workers can now decode TIFF files to generate previews.
- Fix error submitting to Shaman storage from a Windows machine.


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
