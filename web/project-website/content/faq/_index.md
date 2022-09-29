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
[bug]: https://developer.blender.org/maniphest/task/edit/form/14/?tags=Flamenco


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
