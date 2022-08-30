---
title: Blender
---

The location of Blender *on the Worker*, as well as its default arguments, can be
configured via the `blender` and `blenderArgs` variables.

- If the Blender location is just plain `blender`, the worker will try and find
  those by itself. How this is done is different for the two programs, and
  explained below.
- In other cases, it is assumed to be a path and the worker will just use it as
  configured.

Here is an example configuration, which is part of `flamenco-manager.yaml`:

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
  blenderArgs:
    values:
    - platform: all
      value: -b -y
```

## Just `blender`

If the `{blender}` variable is configured to be just `blender` some "smartness"
will kick in. It will pick the first Blender it finds in this order:

1. On Windows, the worker will figure out which Blender is associated with blend
   files. In other words, it will run whatever Blender runs when you
   double-click a `.blend` file. On other platforms this step is skipped.
2. The locations listed in the `PATH` environment variable are searched to find
   Blender. This should run whatever Blender starts when you enter the `blender`
   command in a shell.
3. If none of the above result in a usable Blender, the worker will fail its task.

## An explicit path

If the command is configured to anything other than `blender`, it is assumed to
be a path to the Blender executable.

## Setting Arguments

The `{blenderArgs}` variable can be used to provide arguments to Blender that
are used on every invocation. Flamenco uses these by default:

- `-b`: run Blender in the background, so that it doesn't pop up a window.
- `-y`: allow executing Python code without any confirmation, which is often
  necessary for production rigs to work.

More can be found in Blender's [Command Line Arguments documentation][blendcli].

[blendcli]: https://docs.blender.org/manual/en/latest/advanced/command_line/arguments.html
