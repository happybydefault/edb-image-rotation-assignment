# PBM Image Rotation

This repository contains a Golang application that can rotate PBM images. It's in response to an assignment from
EnterpriseDB (EDB) during the interview process for a position as a Golang/Kubernetes Developer.

The details of the assignment are defined in [ASSIGNMENT.md](ASSIGNMENT.md).

## Install

```sh
go install github.com/happybydefault/edb-image-rotation-assignment/cmd/pbmrotate@latest
```

## Run

```sh
pbmrotate -h
```

```
Usage: pbmrotate [OPTIONS] FILE

Options:
    -d      Number of degrees (default 90)
    -r      Counterclockwise
    -h      Print help
```

### Examples

```sh
# Rotate an image 270 degrees clockwise
pbmrotate -d 270 example-image.pbm

# Rotate an image 90 degrees counterclockwise
pbmrotate -d 90 -r example-image.pbm
```

## Docker

```sh
docker pull ghcr.io/happybydefault/pbmrotate
```
