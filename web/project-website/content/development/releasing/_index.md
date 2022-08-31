---
title: Releasing
weight: 100
---

This page describes how to release a new version of Flamenco. This requires not
only access to the sources, but also to the `flamenco.blender.org`
infrastructure. As such, it's only practically helpful for a small number of
people.

## Steps to release a new version of Flamenco

1. Update `Makefile` and change the `VERSION` and `RELEASE_CYCLE` variables.
2. Run `make update-version`
3. Commit & tag with the commands shown in the 2nd step.
4. Run `make release-package`
5. Check that the files in `dist/` are there and have a non-zero size.
6. Run `make publish-release-packages` to upload the packages to the website.
7. Run `make project-website` to generate and publish the new website.
