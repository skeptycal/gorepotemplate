# %%reponame%%

>Tricky and fun ansi text utilities for Go programs. The obligatory ANSI CLI module predominantly from the Go standard library.

---
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skeptycal/%%reponame%%/Go) ![Codecov](https://img.shields.io/codecov/c/github/skeptycal/%%reponame%%)
 [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](code-of-conduct.md)

![Twitter Follow](https://img.shields.io/twitter/follow/skeptycal.svg?label=%40skeptycal&style=social) ![GitHub followers](https://img.shields.io/github/followers/skeptycal.svg?style=social)

---

## Anansi, A Bit of Lore

Named after Anansi, the trickster, of West African and Caribbean
folklore. Before Anansi, there were no stories in the world. What
a terrible and ignorant place that must have been!

It was Anansi who convinced Nyame, The Sky-God, to share his stories
with the world, but only after capturing the Python, the Hornets,
the Leopard, and the Fairy.

---
---

## Getting Started

### Prerequisites

Developed with Go 1.15.5 darwin/amd64 (macOS Big Sur)
Tested on Go 1.13+

---

### Installation

The easiest way is to clone the repository with Go and install from there.

```bash
# add repo to $GOPATH
go get github.com/skeptycal/anansi

cd ${GOPATH}/src/github.com/skeptycal/anansi

go test -v

go install

```

---

### Basic Usage

>This is a copy of the sample script available in the `sample` folder:

```go
package main

import "github.com/skeptycal/anansi"

func main() {
    anansi.Sample()
}

```

To try it out:

```bash
# change to the sample folder
cd sample

# run the main.go program
go run ./main.go

# You should see the following output demonstrating various color combinations.
```

![sample foreground colors](sample/sample_fg.jpg)
![sample foreground colors](sample/sample_bg.jpg)

---

## Code of Conduct and Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us. Please read the [Code of Conduct](CODE_OF_CONDUCT.md) for details before submitting anything.

---

## Versioning

We use SemVer for versioning. For the versions available, see the tags on this repository.

---

## Contributors
- Michael Treanor ([GitHub][github] / [Twitter][twitter]) - Initial work, updates, maintainer
- [Francesc Campoy][Campoy] - Inspiration and great YouTube videos!

See also the list of contributors who participated in this project.

>Much of the basic types, constants, and terminal checks in 'anansi_const.go' are based on the very popular and well documented [color package][fatih].

---

## License

Licensed under the MIT <https://opensource.org/licenses/MIT> - see the [LICENSE](LICENSE) file for details.


[twitter]: (https://www.twitter.com/skeptycal)
[github]: (https://github.com/skeptycal)
[Campoy]: (https://github.com/campoy)
[fatih]: (https://github.com/fatih/color)
