<!-- markdownlint-disable MD033 -->
# Dockerfile To Develop

This directory contains a Dockerfile to provide the same environment to develop the app.

It includes most of the necessary packages and tools for developing Golang app.

Mostly suitable for [GitHub Codespaces](https://github.com/features/codespaces) and/or [VS Code + Docker](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) users for development.

## Developing Online

If GitHub detects this directory (`.devcontainer`) in your repo, then you will be able to develop online via [GitHub Codespaces](https://github.com/features/codespaces).

Fork this repo to your GitHub account and open it via GitHub Codespaces.

## Developing Locally

### For VS Code + Docker + Remote Container Users

If you already have installed the "[Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)" extension, then press "<kbd>F1</kbd>" and select "`Remote-Containers: Open in Container`".

After a while, you'll get most of the environment needed to develop and debug.

### For Only Docker Users

For non-remote-container user, simply run the below command in the root dir of this repo.

```bash
/bin/bash ./.devcontainer/run-devcontainer.sh
```

- Tested env
  - macOS Catalina (OSX 10.15.7)
  - Docker version: 20.10.5, build 55c4c88

## File Description

- [aliases.sh](aliases.sh) ... Bash alias file to abbreviate type the path to the test script.
- [cobra.yaml](cobra.yaml) ... Default `cobra` command Settings. Used for `$ cobra add ..`
- [createVSCodeUser.sh](createVSCodeUser.sh) ... Bash script that creates `vscode` user in the container.
- [devcontainer.env](devcontainer.env) ... ENV variables to be loaded once when the container's created.
- [devcontainer.json](devcontainer.json) ... VSCode Extensions to be installed and env settings.
- [Dockerfile](Dockerfile) ... Debian 10 (buster) based Golang development container.
- [installToolsForDev.sh](installToolsForDev.sh) ... Bash script that installs the required commands for tests and other helpful commands for development. (For Debian-like OS)
- [postCreateCommand.sh](postCreateCommand.sh) ... Initialization script that runs after the container and the VSCode server is up. Needs to rebuild container if any change made as well.
- [README.md](README.md) ... This file. ;-)
- [run-devcontainer.sh](run-devcontainer.sh) ... Bash script for non-remote-container user. It will pull, build image and runs the container.
- [welcome.sh](welcome.sh) ... Bash script to display the basic info and TIPs to use in the first shell login.

Note: For VSCode + Remote Containers users, if you make any changes to the above files, you need to rebuild the container.

## Required Storage

- You will need around 1.6GB in size.
