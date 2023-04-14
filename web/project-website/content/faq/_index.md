---
title: Frequently Asked Questions
weight: 10
---

This is a list of frequently asked questions, with their answers. It's by no
means an exhaustive list.

## I use a mix of different operating systems, can I still use Flamenco?

Yes, absolutely. To support multiple platforms, first configure your Manager for
its own platform (so if you run that on Linux, use Linux paths). Then you can
use [Two-way Variables][twovars] to translate those paths to the other
platforms.

Do note that Flamenco was developed on Linux, for the Linux-only [Blender
Studio][studio]. You may find issues that the developers did not run into
themselves. If you do, please [report a bug][bug].

[twovars]: {{< ref "usage/variables/multi-platform" >}}
[studio]: https://studio.blender.org/
[bug]: https://projects.blender.org/studio/flamenco/issues/new?template=.gitea%2fissue_template%2fbug.yaml


## My Worker cannot find my Manager, what do I do?

First check the Manager output on the terminal, to see if it shows any messages
about "auto-discovery" or "UPnP/SSDP". Most of the time it's actually Spotify
getting in the way, so make sure to close that before you start the Manager.

If that doesn't help, you'll have to tell the Worker where it can find the
Manager. This can be done on the commandline, by running it like
`flamenco-worker -manager http://192.168.0.1:8080/` (adjust the address to your
situation) or more permanently by [editing the worker configuration
file][workercfg].

[workercfg]: {{< ref "usage/worker-configuration" >}}


## Can I change the paths/names of the rendered files?

Where Flamenco places the rendered files is determined by the job type. You can
create [your own custom job type][jobtypes] or check the existing
[third-party job types][thirdpartyjobs] to change this. With that, you can
even add your own custom job settings like a sequence identifier and use that to
determine the location of rendered files.


## Can I use the Compositor to output multiple EXR files or Passes?

This is possible with Flamenco, but it takes a bit of work. Although it's not
managed by Flamenco's default job types, you can use a [custom job type][jobtypes]
for this.

With that, you have control over the arguments that get used before and/or after
the filename on the CLI.

There are Flamenco jobs out there that support compositor nodes,
multi-platform, and multiple pass outputs. You can check our [third-party jobs
section][thirdpartyjobs].

If you wish to contribute to the project, you're invited to
[get involved with Flamenco][getinvolved]!

[jobtypes]: {{< ref "usage/job-types" >}}
[thirdpartyjobs]: {{< ref "third-party-jobs" >}}
[getinvolved]: {{< ref "development/get-involved" >}}


## Can I use SyncThing, Dropbox, Google Drive, or other file syncing software?

Flamenco assumes that once a file has been written by one worker, it is
immediately available to any other worker, like what you'd get with a NAS.
Similarly, it assumes that when a job has been submitted, it can be worked on
immediately.

Such assumptions no longer hold true when using an asynchronous service like
SyncThing, Dropbox, etc.

Note that this is not just about the initally submitted files. Also the
rendering of a preview video from individual images assumes that those images
are immediately accessible after they've been rendered.

It might be possible to create a complex [custom job type][jobtypes] for this,
but that's all untested. The hardest part is to know when all necessary files
have arrived on a specific worker, without waiting for *all* syncing to be
completed (as someone may have just submitted another job).

## What do "Error: Cached job type is old" or "job type etag does not match" mean?

This means that you have to click on the little "Refresh" icon next to the job type:

<img src="job-types-refresh.webp" width="396" height="41">


## What does "command exited abnormally with code 1" mean?

It means that the program (probably Blender) exited with an error status. Take a
look at the task log, which you can access by going to the task in Flamenco's
web interface.


## ​What's the difference with OpenCue?

OpenCue is aimed at a different audience than Flamenco. OpenCue is a large and
complex project, and relies on a lot of components
([source](https://www.opencue.io/docs/getting-started/)), whereas Flamenco is
made for simplicity and use in small studios or at home, running on your own
hardware.

## Why do I get an Error Performing BAT Pack Message?

As of yet, we've only encountered the issue below on Windows installations. If
you get this issue, please {{< flamenco/reportBugLink size="small" >}}let us
know{{< /flamenco/reportBugLink >}} so that it can be properly investigated.

```
Error performing BAT pack: [WinError 267] The directory name is invalid:
'C:\\The\\Path\\To\\Your\\Project.blend'
```

This is most likely some sort of incompatibility that occurs in some cases where
you might be using linked assets from an asset library in your project.

To work around this issue, try the following:

 * In Blender, use File → External Data → Make Paths Relative.
 * Submit your job again.
