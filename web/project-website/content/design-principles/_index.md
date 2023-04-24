---
title: Design Principles
weight: 25
---

This page describes some of the design ideas & principles behind Flamenco.

## Target Audience

Flamenco is meant for **smaller animation studios and individuals** at home.
Think of roughly **1-10 artists** using it, and **1-100 computers** attached to
the farm to execute tasks. [Blender Studio][studio] uses a handful of servers,
and combines those with various desktop machines when they're not used by the
artists.

## Design Principles

The following principles guide the design of Flamenco:

Blender.org Project
: Flamenco is a true blender.org project. This means that it's Free and Open
Source, made by the community, lead by Blender HQ. Its development will fall
under the umbrella of the [Pipline, Assets & IO][PAIO] module.

[PAIO]: https://projects.blender.org/blender/blender/wiki/Module:%20Pipeline,%20Assets%20&%20I/O

Minimal Authentication & Organisation
: Because Flamenco is aimed at small studios and individuals, it won't offer
much in terms of user authentication, nor the organisation of users into groups.
[Custom job types][jobtypes] can be used to attach arbitrary metadata to jobs,
such as the submitter's name, a project identifier, etc.

[jobtypes]: {{< ref "/usage/job-types" >}}

Minimize External Components
: Running Flamenco should be extremely simple. This means that it should depend
on as few external packages as possible. Apart from the Flamenco components
themselves, all you need to install is [Blender][blender].
: The downside of this is that development might take longer, as some things
that an external service could solve need to be implemented. This trade-off of
developer time for simplicity of use is considered a good thing, though.

[blender]: https://www.blender.org/

No Errors, Guide Users To Success
: Instead of stopping with a description of what's wrong, like "no database
configured", Flamenco should show something helpful in which you're guided
towards a working system.

Customisable
: Studio pipeline developers / TDs should be able to customise the behaviour of
Flamenco. They should be able to create new [job types][jobtypes], and adjust
existing job types to their needs. For this, Flamenco uses JavaScript to convert
a job definition like "*render this blend file, frames 1-100*" into individual
tasks for computers to execute.

Work offline
: Like Blender itself, Flamenco should be able to fully work offline. That is,
work without internet connection. If any future feature should need such a
connection, that feature should always be optional, and be disabled by default.

Data Storage
: Data should be stored as plain files whenever possible. Where a higher level
of coordination is required, an embedded database can be used; currently
Flamenco uses [SQLite][sqlite] for this.

[sqlite]: https://pkg.go.dev/modernc.org/sqlite

## Infrastructure & Supported Platforms

Setting up a render farm is not as simple as pushing a button, but Flamenco aims
to keep things as simple as possible. What you need to run Flamenco is:

- One or more computers to do the work, i.e. running Flamenco Worker.
- A computer to run the central software, Flamenco Manager. This could be one of
  the above computers, or a dedicated one.
- A local network with file sharing already set up, so that the above computers
  can all reach the same set of files.

Since Blender Studio fully runs on Open Source software, Linux is the main
platform Flamenco is developed for. Windows and macOS will also be supported,
but will need help from the community to get tested & developed well.

## Software Design

The Flamenco software follows an **API-first** approach. All the functionality
of Flamenco Manager is exposed via [the OpenAPI interface][openapi] ([more
info](openapi-info)). The web interface is no exception; anything you can do
with the web interface, you can do with any other OpenAPI client.

- The API can be browsed by following the 'API' link in the top-right corner of
  the Flamenco Manager web interface. That's a link to
  `http://your.manager.address/api/v3/swagger-ui/`
- The web interface, Flamenco Worker, and the Blender add-on are all using that
  same API.

[openapi]: https://projects.blender.org/studio/flamenco/src/branch/main/pkg/api/flamenco-openapi.yaml
[openapi-info]: https://www.openapis.org/

## New Features

To add a new feature to Flamenco, these steps are recommended:

1. Define which changes to the API are necessary, and update the [flamenco-openapi.yaml][openapi] file for this.
1. Run `go generate ./pkg/...` to generate the OpenAPI Go code.
1. Implement any new operations in a minimal way, so that the code compiles (but doesn't do anything else).
1. Run `make generate` to regenerate all the code (so also the JavaScript and Python client, and Go mocks).
1. Write unit tests that test the new functionality.
1. Write the code necessary to make the unit tests pass.
1. Now that you know how it can work, refactor to clean it up.
1. Send in a pull request!
