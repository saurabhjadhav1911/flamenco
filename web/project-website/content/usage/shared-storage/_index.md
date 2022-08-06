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

## Work Directly on the Shared Storage

Working directly in the shared storage is the simplest way to work with Flamenco.

## Creating a Copy for Each Render Job

The "work on shared storage" approach has the downside that render jobs are not
fully separated from each other. For example, when you change a texture while a
render job is running, the subsequently rendered frames will be using that
altered texture. If this is an issue for you, and you cannot use the [Shaman
Storage System][shaman], the approach described in this section is for you.


## Shaman Storage System

See [Shaman Storage System][shaman].

[shaman]: {{< relref "shaman" >}}
