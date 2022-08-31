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
3. Update the version for the website in
   `web/project-website/data/flamenco.yaml`. This is currently a separate step,
   as it influences the project website and not the built software. This could
   be included in the `update-version` make target at some point.
4. Commit & tag with the commands shown in the 2nd step.
5. Run `make release-package`
6. Check that the files in `dist/` are there and have a non-zero size.
7. Run `make publish-release-packages` to upload the packages to the website.
8. Run `make project-website` to generate and publish the new website.
