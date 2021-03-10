# Quay Installer

This application will allow user to install Quay and its required components using a simple CLI tool.

## Pre-Requisites

- RHEL 8 machine with Podman installed
- `sudo` access on desired host (rootless install tbd)
- make (only if compiling using Makefile)

### Compile

To compile the quay-installer.tar.gz for distribution, run the following command:

```
$ git clone https://github.com/jonathankingfc/quay-aioi.git
$ cd quay-aioi
$ make build-offline-installer # OR make build-online-installer
```

This will generate a `quay-installer.tar.gz` which contains this README.md, the `quay-installer` binary, and the `image-archive.tar` (if using offline installer) which contains the images required to set up Quay.

Once generated, you may untar this file on your desired host machine for installation. You may use the following command:

```
tar -xzvf quay-installer.tar.gz
```

NOTE - This may take some time.

### Installation

To install Quay on your desired host machine, run the following command:

```
$ sudo ./quay-installer install -v
```

This command will make the following changes to your machine

- Pulls Quay, Redis, and Postgres containers from registry.redhat.io
- Sets up systemd files on host machine to ensure that container runtimes are persistent
- Creates `~/quay-install` in `$HOME` which contains install files, local storage, and config bundle. This will generally be in `/root/quay-install`.

### Uninstall

To uninstall Quay, run the following command:

```
$ sudo ./quay-installer uninstall -v
```

This command will delete the `~/quay-install` directory and disable all systemd services set up by Quay.

### To Do

- Switch from --net=host to a bridge network (this is safer)
- Figure out SELinux issues (currently not working with SELinux)
- Better config generation with secure passwords
