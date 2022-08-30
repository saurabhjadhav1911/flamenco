---
title: Getting Started
weight: 1
---

To start, get a **Git checkout** with either of these commands. The 1st one is for
public, read-only access. The 2nd one can be used if you have commit rights to
the project.

```
git clone git://git.blender.org/flamenco.git
git clone git@git.blender.org:flamenco.git
```

Then follow the steps below to get everything up & running.

## 1. Installing Go

Most of Flamenco is made in Go.

1. Install [Go 1.18 or newer](https://go.dev/).
2. Optional: set the environment variable `GOPATH` to where you want Go to put its packages. Go will use `$HOME/go` by default.
3. Ensure `$GOPATH/bin` is included in your `$PATH` environment variable. Run `go env GOPATH` if you're not sure what path to use.

## 2. Installing NodeJS

The web UI is built with Vue, Bootstrap, and Socket.IO for communication with
the backend. **NodeJS+Yarn** is used to collect all of those and build the
frontend files.

{{< tabs "installing-nodejs" >}}
{{< tab "Linux" >}}
It's recommended to install Node via Snap:

```
sudo snap install node --classic --channel=16
```

If you install NodeJS in a different way, it may not be bundled with Yarn. In that case, run:

```
sudo npm install --global yarn
```

{{< /tab >}}
{{< tab "Windows" >}}
Install [Node v16 LTS](https://nodejs.org/en/download/). Be sure to enable the "Automatically install the necessary tools" checkbox.

Then install Yarn via:

```
npm install --global yarn
```

{{< /tab >}}
{{< tab "macOS" >}}
**Option 1** (Native install)

Install [Node v16 LTS](https://nodejs.org/en/download/) and then install Yarn via:

```
npm install --global yarn
```

<br />

**Option 2** (Homebrew)

Install Node 16 via homebrew:

```
brew install node@16
```

Then install yarn:

```
brew install yarn
```

{{< /tab >}}
{{< /tabs >}}

## 3. Utilities

Building Flamenco requires only a few tools to be installed on your system.


{{< tabs "installing-utils" >}}
{{< tab "Linux" >}}
On Linux only `make` is necessary, which can be installed via your package manager.

On Debian, and relatives like Ubuntu, run:

```
sudo apt install make
```
{{< /tab >}}
{{< tab "Windows" >}}
Install [MingW W64][mingw]. If in doubt which version to get, grab the `x86_64...seh` one.
You'll need [7Zip][7zip] to extract it.

[mingw]: https://github.com/niXman/mingw-builds-binaries/releases
[7zip]: https://www.7-zip.org/download.html
{{< /tab >}}
{{< tab "macOS" >}}
TODO: write this documentation.
{{< /tab >}}
{{< /tabs >}}

## 4. Your First Build

Run `make with-deps` to install build-time dependencies and build the application.
Subsequent builds can just run `make` without arguments.

You should now have two executables: `flamenco-manager` and `flamenco-worker`.
Both can be run with the `-help` CLI argument to see the available options.

See [building][building] for more `make` targets, for example to run unit tests,
enable the race condition checker, and all other kinds of useful things.

[building]: {{< relref "../building/" >}}

## 5. Get Involved

If you're interested in helping out with Flamenco development, please read [Get Involved!][get-involved]

[get-involved]: {{ <ref "development/get-involved" >}}
