---
title: Variables & Multi-Platform Support
weight: 2
---

For managing default parameters for Blender and FFmpeg, as well as for
mixed-platform render farms, Flamenco uses *variables*.

Each variable consists of:

- Its **name**: just a sequence of letters, like `ffmpeg` or `whateveryouwant`.
- Its **values**: per *platform* and *audience*, described below.

The variables are configured in `flamenco-manager.yaml`.

Here is an example `blender` variable:

```yaml
variables:
  blender:
    values:
    - platform: linux
      value: /home/sybren/Downloads/blenders/blender-3.2.2-release/blender -b -y
    - platform: windows
      value: B:\Downloads\blenders\blender-3.2.2-release\blender.exe -b -y
    - platform: darwin
      value: /home/sybren/Downloads/blenders/blender-3.2.2-release/blender -b -y
```

Whenever a Worker gets a task that contains the string `{blender}`, that'll be
replaced by the appropriate value for that worker.

## Platform

**The goal of the variables system is to cater for different platforms.**
Blender will very likely be installed in different locations on Windows and
Linux. It might even require some different parameters for your farm, depending
on the platform. The variables system allows you to configure this.

The platform can be `windows`, `linux`, or `darwin` for macOS. Other platforms
are also allowed, if you happen to use them in your farm.

{{< hint type=note title="Custom Job Types" >}}
This documentation section focuses on pre-existing variables, `blender` and
`ffmpeg`. There is nothing special about these. Apart from being part of
Flamenco's default configuration, that is. When you go the more advanced route
of creating your own [custom job types][jobtypes] you're free to create your own
set of variables to suit your needs.

[jobtypes]: {{< ref "usage/job-types" >}}
{{< /hint >}}

## Audience

The audience of a value is who that value is for: `workers`, `users`, or `all`
if there is no difference in value for workers and users.

This is an advanced feature, and was introduced for in-the-cloud render farms.
In such situations, the location where the workers store the rendered frames
might be different from where users go to pick them up.

- `all`: values that are used for all audiences. This is the default, and is
  what's used in the above example (because there is no `audience` mentioned).
- `users`: values are used when submitting jobs from Blender and showing them in
  the web interface.
- `workers`: values that are used when sending tasks to workers.

## Blender and FFmpeg

The location of Blender and FFmpeg on the Worker, as well as their default
arguments, can be configured via the `blender` and `ffmpeg` variables.

- If the Blender or FFmpeg location is just plain `blender` resp. `ffmpeg`, the
  worker will try and find those by itself. How this is done is different for
  the two programs, and explained below.
- In other cases, it is assumed to be a path and the worker will just use it as
  configured.

### Blender

The built-in [job types][jobtypes] use `{blender}` instead of hard-coding a
particular command. This makes it possible to configure your own Blender
command. The Worker has a few strategies for finding Blender, described below.

[jobtypes]: {{< ref "usage/job-types" >}}

#### Just `blender`

If the command is configured to be just `blender` with some arguments, for
example `blender -b -y`, some "smartness" will kick in. It will pick the first
Blender it finds in this order:

1. On Windows, the worker will figure out which Blender is associated with blend
   files. In other words, it will run whatever Blender runs when you
   double-click a `.blend` file. On other platforms this step is skipped.
2. The locations listed in the `PATH` environment variable are searched to find
   Blender. This should run whatever Blender starts when you enter the `blender`
   command in a shell.
3. If none of the above result in a usable Blender, the worker will fail its task.

#### An explicit path

If the command is configured to anything other than `blender` (arguments
excluded), it is assumed to be a path to the Blender executable. For an example,
see the top of this page.

### FFmpeg

Similar to `{blender}`, described above, the `{ffmpeg}` variable can be used to
configure both which FFmpeg executable to use, and which additional parameters
to pass.

#### Just `ffmpeg`

If the command is configured to be just `ffmpeg` with some arguments, for
example `ffmpeg -hide_banner`, the worker will try to use the bundled FFmpeg.
This is the default behavior, so if you do not have any `ffmpeg` variable
defined, this is what will happen.

The worker looks next to the `flamenco-worker` executable for a `tools`
directory. Given the current OS (`windows`, `linux`, `darwin`, etc.) and
architecture (`amd64`, `x86`, etc.) it will try to find the most-specific one
for your system in that `tools` directory.

The worker tries the following ones; the first one found is used:

1. `ffmpeg-OS-ARCHITECTURE`, for example `ffmpeg-windows-amd64.exe` or `ffmpeg-linux-x86`
2. `ffmpeg-OS`, for example `ffmpeg-windows.exe` or `ffmpeg-linux`
3. `ffmpeg`, so `ffmpeg.exe` on Windows and `ffmpeg` on any other platform.

#### An explicit path

If the command is configured to anything other than `ffmpeg` (arguments
excluded), it is assumed to be a path to the FFmpeg executable.

Example configuration from `flamenco-manager.yaml`:

```yaml
variables:
  ffmpeg:
    values:
    - platform: linux
      value: /media/shared/tools/ffmpeg-5.0/ffmpeg-linux
    - platform: windows
      value: S:\tools\ffmpeg-5.0\ffmpeg.exe
    - platform: darwin
      value: /Volumes/shared/tools/ffmpeg-5.0/ffmpeg-macos
```

## Two-way Variables

Two-way variables are there to support mixed-platform Flamenco farms. Basically
they perform *path prefix replacement*.

Let's look at an example configuration:

```yaml
variables:
  shared_storage:
    is_twoway: true
    values:
    - platform: linux
      value: /media/shared/flamenco
    - platform: windows
      value: F:\flamenco
    - platform: darwin
      value: /Volumes/shared/flamenco
```

The difference with regular variables is that regular variables are one-way, so
only `{variablename}` is replaced with the value. Two-way variables go both ways, as follows:

- When submitting a job from a Linux workstation, `/media/shared/flamenco` (the
  variable's value for Linux) will be replaced with `{shared_storage}`.
- When sending a task for this job to a Windows worker, `{shared_storage}` will
  be replaced with `F:\flamenco`.

Let's look at a more concrete example, with the same configuration as above.

- Alice runs Blender on **macOS**. She submits a job that has its render output set
  to `/Volumes/shared/flamenco/renders/shot_010_a_anim`.
- Flamenco recognises the path, and stores the job as rendering to
  `{shared_storage}/renders/shot_010_a_anim`.
- Bob's computer is running the Worker on **Windows**, so when it receives a render
  task Flamenco will tell it to render to
  `F:\flamenco\renders\shot_010_a_anim`.
- Carol's computer is also running a worker, but on **Linux**. When it receives a
  render task, Flamenco will tell it to render to
  `/media/shared/flamenco/renders/shot_010_a_anim`.
