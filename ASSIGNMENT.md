# Assignment: PBM Image Rotation

## Problem statement

Grandmother has an old cell phone that can only take pictures in PBM format. She is wanting you to write an application
in Go for her that will rotate a given arbitrary PBM image 90 degrees clockwise and output the result. Grandmother also
prefers that we work from first principles and do not use any third party libraries to help. If possible, Grandmother
has a feature request, and would like to be able to rotate her PBM images clockwise or counterclockwise by an arbitrary
number of degrees.

## PBM

[PBM](https://en.wikipedia.org/wiki/Netpbm) is an image format that can be represented in ascii and is easily
manipulated.

An example pbm file could look like this:

```
P1
# This is an example bitmap of the letter "J"
6 10
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0
0 0 0 0 0 0
```

We can assume that Grandmother's phone uses pbm files with the magic number P1 only. Alas, grandmother's phone doesn't
always generate square images, it can create rectangular images as well.

## Deliverables

An application in Golang that will accept a filename, and the number of degrees to rotate, and will write the resultant
image in P1 PBM format to disk. As we do not want to startle Grandmother, the application should not panic in any
circumstances.

Please submit your solution to me via email consisting of the zipped (tar.gz) folder containing the source code of the
solution, including the .git directory.
