---
title: Variables
weight: 2
---

For managing default parameters for Blender and FFmpeg, as well as for
mixed-platform render farms, Flamenco uses *variables*.

Each variable consists of:

- Its **name**: just a sequence of letters, like `blender` or `whateveryouwant`.
- Its **values**: per *platform* and *audience*, described below.

The variables are configured in `flamenco-manager.yaml`.

Here is an example `blender` variable:

```yaml
variables:
  blender:
    values:
    - platform: linux
      value: /home/sybren/Downloads/blenders/blender-3.2.2-release/blender
    - platform: windows
      value: B:\Downloads\blenders\blender-3.2.2-release\blender.exe
    - platform: darwin
      value: /home/sybren/Downloads/blenders/blender-3.2.2-release/blender
```

Whenever a Worker gets a task that contains `{blender}`, that'll be replaced by
the appropriate value for that worker.

{{< expand "Custom Job Types" >}}
This documentation section focuses on pre-existing variables, `blender` and
`blenderArgs`. There is nothing special about these. Apart from being part of
Flamenco's default configuration, that is. When you go the more advanced route
of creating your own [custom job types][jobtypes] you're free to create your own
set of variables to suit your needs.

[jobtypes]: {{< ref "usage/job-types" >}}
{{< /expand >}}

## Platform

**The goal of the variables system is to cater for different platforms.**
Blender will very likely be installed in different locations on Windows and
Linux. It might even require some different parameters for your farm, depending
on the platform. The variables system allows you to configure this.

The platform can be `windows`, `linux`, or `darwin` for macOS. Other platforms
are also allowed, if you happen to use them in your farm.

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
