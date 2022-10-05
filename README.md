# cpy3x 0.1.0

## What is cpy3x?

cpy3x is a golang package, which is a wrapper around the [python c API](https://docs.python.org/3/c-api). This library assists the usage of python c API in golang, by exposing the the python c API through CGO.

## What versions of python c API are supported?

Currently cpy3x supports python versions starting from **3.5.x till 3.10.x**. It has multiple packages named as follows:

- ``pycore`` - This package contains the APIs that are either part of [Limited API](https://docs.python.org/3/c-api/stable.html#stable-abi-list) or those APIs are common across versions 3.5-3.10.
- ``py35`` - Contains APIs and unit-tests specific to python version 3.5
- ``py36`` - Contains APIs and unit-tests specific to pythoh version 3.6
- ``py37`` - Contains APIs and unit-tests specific to pythoh version 3.7
- ``py38`` - Contains APIs and unit-tests specific to pythoh version 3.8
- ``py39`` - Contains APIs and unit-tests specific to pythoh version 3.9
- ``py310`` - Contains APIs and unit-tests specific to pythoh version 3.10

The main reason behind this approach of separating version specific APIs in different packages is to make the ``pycore`` package a generic one, which contains all the APIs required to build a regular c-python extension, which is compatible across multiple versions. If you use only ``pycore`` package in your program, it can be compiled any python version from 3.5 to 3.10.

As new APIs are introduced in the new versions and older APIs keep depricating, we have those version specific APIs in the version specific packages. If you are using a version specific package say ``py37`` along with ``pycore``, then make sure you link the version specific python during compilation.

These packages contain ``PyXYZ`` functions and macros of the public C-API of CPython have been
exposed. Theoretically, you should be able use https://docs.python.org/3/c-api
and know what to type in your ``go`` program.

## relation to `go-python/cpy3`

This project is a successor to [`go-python/cpy3`](https://github.com/go-python/cpy3), which only supports python version 3.7. If cpy3x works as expected, I might raise a PR so that it could be merged to cpy3. 

# Install

## Deps

We will need `pkg-config` and a working version specific `python3` environment to build these
bindings. Make sure you have Python libraries and header files installed as
well (`python3.7-dev` on Debian or `python3-devel` on Centos for example)..

By default `pkg-config` will look at the `python3` library so if you want to
choose a specific version just symlink `python-X.Y.pc` to `python3.pc` or use
the `PKG_CONFIG_PATH` environment variable. You can take a look into ``interpreters`` folder to check what python3.pc files look like on ``Ubuntu``. For more info on ``cflags`` and ``ldflags`` check out the [documentation](https://docs.python.org/3/extending/embedding.html#compiling-and-linking-under-unix-like-systems). 

## Go get

Then simply `go get github.com/go-python/cpy3x@main`

# API

Some functions mix go code and call to Python function. Those functions will
return and `int` and `error` type. The `int` represent the Python result code
and the `error` represent any issue from the Go layer.

Example:

`func PyRun_AnyFile(filename string)` open `filename` and then call CPython API
function `int PyRun_AnyFile(FILE *fp, const char *filename)`.

Therefore its signature is `func PyRun_AnyFile(filename string) (int, error)`,
the `int` represent the error code from the CPython `PyRun_AnyFile` function
and error will be set if we failed to open `filename`.

If an error is raise before calling th CPython function `int` default to `-1`.

# Usage 

- TODO (You can get started by looking at some unittests in each package)

# Tutorials and Blogs

- [CGO Basics](https://jshiwam.github.io/go-c-python/2022/08/07/gotoc.html)
- [CPython Basics](https://jshiwam.github.io/go-c-python/2022/08/31/ctopy.html)
- [Integrating Cpython and Cgo](https://jshiwam.github.io/go-c-python/2022/09/11/gotopy.html)
- [DataDog Tutorial](https://www.datadoghq.com/blog/engineering/cgo-and-python/)
- [Christian's Tutorial](https://poweruser.blog/embedding-python-in-go-338c0399f3d5)
- [Mike's Tutorial](https://www.ardanlabs.com/blog/2020/09/using-python-memory.html)

