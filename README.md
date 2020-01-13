## imgres

Resize image with ease. Supports GIF, JPEG and PNG.

### Installation

Go toolchain required.

```
go get -u github.com/maku693/imgres
```

### Usage

```
Usage of imgres:
  -fit string
        fitting of scaled image (default "contain")
  -height int
        max height of out file
  -in string
        input file (optional, default stdin)
  -out string
        out file (optional, default stdout)
  -width int
        max width of out file
```

#### Examples

##### Simple

| Before                                    | After                                     |
| ----------------------------------------- | ----------------------------------------- |
| ![Imgur](https://i.imgur.com/jNxS9du.png) | ![After](https://i.imgur.com/brO3VE7.png) |

```
imgres -in image.png -width 150 -out image150.png
```

##### Specify all options

| Before                                    | After                                     |
| ----------------------------------------- | ----------------------------------------- |
| ![Imgur](https://i.imgur.com/jNxS9du.png) | ![After](https://i.imgur.com/NgRYQLY.png) |

```
imgres -in image.png -width 150 -height 200 -fit cover -out image150x200cover.png
```

##### Use stdout/stdin

| Before                                    | After                                     |
| ----------------------------------------- | ----------------------------------------- |
| ![Imgur](https://i.imgur.com/jNxS9du.png) | ![After](https://i.imgur.com/9MXfkfn.png) |

```
imgres -width 100 < image.png > image100.png
```
