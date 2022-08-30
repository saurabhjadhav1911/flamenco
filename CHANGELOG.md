# Flamenco Change Log

This file contains the history of changes to Flamenco. Only changes that might
be interesting for users are listed here, such as new features and fixes for
bugs in actually-released versions.

## 3.0-beta2 - in development

- Manager & Add-on: avoid error that could occur when submitting jobs with UDIM files
  ([44ccc6c3ca70](https://developer.blender.org/rF44ccc6c3ca706fdd268bf310f3e8965d58482449)).
- Manager: don't stop when the Flamenco Setup Assistant cannot start a webbrowser
  ([7d3d3d1d6078](https://developer.blender.org/rF7d3d3d1d6078828122b4b2d1376b1aaf2ba03b8b)).
- Change path inside the Linux and macOS tarballs, so that they contain an
  embedded `flamenco-3.x.y-xxxx/` directory with all the files (instead of
  putting all the files in the root of the tarball).
- Two-way variable replacement now also changes the path separators to the target platform.
- Allow setting priority when submitting a job.


## 3.0-beta1 - released 2022-08-03

This was the first version of Flamenco to be released to the public, and thus it
serves as the starting point for this change log.
