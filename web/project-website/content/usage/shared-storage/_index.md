---
title: Shared Storage
---

Flamenco needs some form of *shared storage*: a place for files to be stored
that can be accessed by all the computers in the farm.

Basically there are three approaches to this:

| Approach                            | Simple | Efficient | Render jobs are isolated |
|-------------------------------------|--------|-----------|--------------------------|
| Work directly on the shared storage | ✅      | ✅         | ❌                        |
| Create a copy for each render job   | ✅      | ❌         | ✅                        |
| Shaman Storage System               | ❌      | ✅         | ✅                        |

Each is explained below.

{{< hint type=Warning >}}
On Windows, Flamenco **only supports drive letters** to indicate locations.
Flamenco does **not** support UNC notation like `\\SERVER\share`. Mount the
network share to a drive letter. The examples below use `S:` for this.
{{< /hint >}}

## Work Directly on the Shared Storage

Working directly in the shared storage is the simplest way to work with
Flamenco. You can enable this mode by pointing Flamenco at the location of your
blend files.

As an example, if `S:\WorkArea` is where your blend files live (or in a
subdirectory thereof), you can update your `flamenco-manager.yaml` like this:

```yaml
shared_storage_path: S:\WorkArea
shaman:
  enabled: false
```

When you submit a file from the shared storage, say
`S:\WorkArea\project\scene\shot\anim.blend`, Flamenco will detect this and
assume the Workers can reach the file there. No copy will be made.

## Creating a Copy for Each Render Job

The "work on shared storage" approach has the downside that render jobs are not
fully separated from each other. For example, when you change a texture while a
render job is running, the subsequently rendered frames will be using that
altered texture. If this is an issue for you, and you cannot use the [Shaman
Storage System][shaman], the approach described in this section is for you.

As an example, if `C:\WorkArea` is where you work on your blend files, and
`S:\Flamenco` is the shared storage for Flamenco, you will automatically use
this approach. You can update your `flamenco-manager.yaml` like this:

```yaml
shared_storage_path: S:\Flamenco
shaman:
  enabled: false
```

As you can see, you do not have to tell Flamenco about `C:\WorkArea`, it'll
automatically detect which storage approach to use from the path of the blend
file you're submitting.

## Shaman Storage System

This requires a bit more to explain. See [Shaman Storage System][shaman].

[shaman]: {{< relref "shaman" >}}
