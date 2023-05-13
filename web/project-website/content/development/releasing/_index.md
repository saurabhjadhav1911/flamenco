---
title: Releasing
weight: 100
---

This page describes how to release a new version of Flamenco. This requires not
only access to the sources, but also to the `flamenco.blender.org`
infrastructure. As such, it's only practically helpful for a small number of
people.

## Steps to release a new version of Flamenco

Replace `${VERSION}` with the actual version number.

1. Ensure that env variables are set in the `.env` file (see `.env.example` for reference)
1. Update `CHANGELOG.md` and mark the to-be-released version as released today.
1. Update `web/project-website/data/flamenco.yaml` for the new version.
1. Update `Makefile` and change the `VERSION` and `RELEASE_CYCLE` variables.
1. `make update-version`
1. `git commit -m "Bump version to ${VERSION}"`
1. `git tag v${VERSION}`
1. `make release-package`
1. Check that the files in `dist/` are there and have a non-zero size.
1. `make publish-release-packages` to upload the packages to the website.
1. `make project-website` to generate and publish the new website.
1. `git push && git push --tags`
