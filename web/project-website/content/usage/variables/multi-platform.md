---
title: Two-way Variables for Multi-Platform Support
---

Two-way variables are there to support mixed-platform Flamenco farms. Basically
they perform *path prefix replacement*.

Let's look at an example configuration:

```yaml
variables:
  my_storage:
    is_twoway: true
    values:
    - platform: linux
      value: /media/shared/flamenco
    - platform: windows
      value: F:\flamenco
    - platform: darwin
      value: /Volumes/shared/flamenco
```

The difference with regular variables is that regular variables are one-way:
`{variablename}` is replaced with the value, and that's it.

Two-way variables go both ways, as follows:

- When submitting a job, values are replaced with variables.
- When sending a task to a worker, variables are replaced with values again.

This may seem like a lot of unnecessary work. After all, why go through the
trouble of replacing in one direction, when later the opposite is done? The
power lies in the fact that each replacement step can target a different
platform. In the first step the value for Linux can be recognised, and in the
second step the value for Windows can be put in its place.

Let's look at a more concrete example, based on the configuration shown above.

- Alice runs Blender on **macOS**. She submits a job that has its render output set
  to `/Volumes/shared/flamenco/renders/shot_010_a_anim`.
- Flamenco recognises the path, and stores the job as rendering to
  `{my_storage}/renders/shot_010_a_anim`.
- Bob's computer is running the Worker on **Windows**, so when it receives a render
  task Flamenco will tell it to render to
  `F:\flamenco\renders\shot_010_a_anim`.
- Carol's computer is also running a worker, but on **Linux**. When it receives a
  render task, Flamenco will tell it to render to
  `/media/shared/flamenco/renders/shot_010_a_anim`.
