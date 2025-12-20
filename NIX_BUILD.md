# Building gnparser with Nix

This document provides instructions for building and installing `gnparser`
using Nix package manager.

<!-- TOC GFM -->

* [Nix Files](#nix-files)
* [Prerequisites](#prerequisites)
* [Building Methods](#building-methods)
  * [Method 1: Using Nix Flakes (Recommended)](#method-1-using-nix-flakes-recommended)
    * [Build with Flakes](#build-with-flakes)
    * [Run without installing](#run-without-installing)
    * [Install with Flakes](#install-with-flakes)
  * [Method 2: Using Traditional Nix](#method-2-using-traditional-nix)
    * [Build with Nix](#build-with-nix)
    * [Install with Nix Env](#install-with-nix-env)
    * [Run the built binary](#run-the-built-binary)
* [Development Shell](#development-shell)
  * [With Flakes](#with-flakes)
  * [Without Flakes](#without-flakes)
* [Build Configuration](#build-configuration)
  * [Supported Platforms](#supported-platforms)
* [Updating vendorHash](#updating-vendorhash)
* [Troubleshooting](#troubleshooting)
  * [Flakes not enabled](#flakes-not-enabled)
  * [Hash mismatch errors](#hash-mismatch-errors)
* [Using in NixOS Configuration](#using-in-nixos-configuration)

<!-- /TOC -->

## Nix Files

The repository contains the following Nix-related files:

* `default.nix` - Main build derivation using buildGoModule
* `flake.nix` - Nix flakes configuration for modern Nix workflows
* `shell.nix` - Development shell with Go and gopls

## Prerequisites

* Nix package manager installed (see <https://nixos.org/download.html>)
* For flake-based builds: Nix 2.4+ with flakes enabled

Make sure that `version`, `date` and `vendorHash` in `default.nix`
are updated for a successful build.

The simplest way to get the vendorHash is to run build and copy corrected hash
from the error message.

## Building Methods

gnparser supports both traditional Nix and modern Nix Flakes build methods.

### Method 1: Using Nix Flakes (Recommended)

#### Build with Flakes

```bash
nix build
```

This will create a `result` symlink in the current directory pointing to the
build output.

#### Run without installing

```bash
nix run
```

#### Install with Flakes

```bash
nix profile install
```

Or install from a specific flake reference:

```bash
nix profile install github:gnames/gnparser
```

### Method 2: Using Traditional Nix

#### Build with Nix

```bash
nix-build
```

This will create a `result` symlink pointing to the built package.

#### Install with Nix Env

```bash
nix-env -f default.nix -i
```

#### Run the built binary

```bash
./result/bin/gnparser
```

## Development Shell

Enter a development environment with Go and necessary tools:

### With Flakes

```bash
nix develop
```

### Without Flakes

```bash
nix-shell
```

This provides:

* Go compiler
* gopls (Go language server)

## Build Configuration

The build is configured in `default.nix` with the following features:

* **Static binary**: Built with `-linkmode external -extldflags -static`
* **Version stamping**: Version and build date are embedded at build time
* **Vendor dependencies**: Go dependencies are managed with a fixed `vendorHash`
* **Optimized**: Stripped symbols (`-s -w`) for smaller binary size

### Supported Platforms

* x86_64-linux
* aarch64-linux
* x86_64-darwin
* aarch64-darwin

## Updating vendorHash

If Go dependencies change, you may need to update the `vendorHash` in `default.nix`:

1. Set `vendorHash = lib.fakeHash;` in `default.nix`
2. Run `nix build` (or `nix-build`)
3. Copy the correct hash from the error message
4. Update `vendorHash` with the correct value

## Troubleshooting

### Flakes not enabled

If you get an error about flakes not being recognized, enable them:

```bash
mkdir -p ~/.config/nix
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf
```

Or use the flag directly:

```bash
nix --experimental-features 'nix-command flakes' build
```

### Hash mismatch errors

If you encounter vendorHash mismatch errors, follow the "Updating
vendorHash" section above.

## Using in NixOS Configuration

You can add gnparser to your NixOS configuration:

```nix
{
  environment.systemPackages = [
    (pkgs.callPackage /path/to/gnparser/default.nix {})
  ];
}
```

Or with flakes in your `flake.nix`:

```nix
{
  inputs.gnparser.url = "github:gnames/gnparser";

  outputs = { self, nixpkgs, gnparser }: {
    nixosConfigurations.yourhostname = nixpkgs.lib.nixosSystem {
      modules = [
        {
          environment.systemPackages = [ gnparser.packages.x86_64-linux.default ];
        }
      ];
    };
  };
}
```
