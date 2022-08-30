---
title: FFmpeg
---

FFmpeg is bundled with Flamenco, and at the moment it cannot be configured to
use any other install of FFmpeg. This is intended to be added in some future,
but there is no timeline when this will happen (patches welcome!).

The worker looks next to the `flamenco-worker` executable for a `tools`
directory. Given the current OS (`windows`, `linux`, `darwin`, etc.) and
architecture (`amd64`, `x86`, etc.) it will try to find the most-specific one
for your system in that `tools` directory.

The worker tries the following ones; the first one found is used:

1. `ffmpeg-OS-ARCHITECTURE`, for example `ffmpeg-windows-amd64.exe` or `ffmpeg-linux-x86`
2. `ffmpeg-OS`, for example `ffmpeg-windows.exe` or `ffmpeg-linux`
3. `ffmpeg`, so `ffmpeg.exe` on Windows and `ffmpeg` on any other platform.
