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
    -d      Number of degrees. Possible values are only 90, 180, and 270 (default 90)
    -c      Counterclockwise
    -o      Write the result to file instead of stdout
    -h      Print this help text
```

### Examples

```sh
# Rotate an image 270 degrees clockwise and write the result to a file
pbmrotate -d=270 -o="example-image-rotated.pbm" example-image.pbm

# Rotate an image 90 degrees counterclockwise and write the result to stdout
pbmrotate -d=90 -c example-image.pbm

# Rotate an image 180 degrees from stdin and write the result to a file
curl "https://example.com/internet-image.pbm" | pbmrotate -d=180 -o="internet-image-rotated.pbm"
```

## Docker

```sh
docker pull ghcr.io/happybydefault/pbmrotate
```
