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
pbmrotate [OPTIONS] FILE
```

```
Options:
    -d      Degrees (counterclockwise if negative)
```

### Examples

```sh
# Rotate an image 90° clockwise
pbmrotate -d 90 example-image.pbm

# Rotate an image 270° counterclockwise
pbmrotate -d -270 example-image.pbm
```

## Docker

```sh
docker pull ghcr.io/happybydefault/pbmrotate
```
