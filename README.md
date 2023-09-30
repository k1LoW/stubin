# stubin

This is stub binary.

## Usage

By setting environment variables prefixed with the binary name, STDOUT, STDERR, and Exit status can be controlled.

```console
$ export STUBIN_STDOUT="Hello, stubin!
"
$ stubin
Hello, stubin!
$
```

| Binary name | Environment Variables for STDOUT | Environment Variables for STDERR | Environment Variables for Exit status |
| --- | --- | --- | --- |
| `stubin` | `STUBIN_STDOUT` | `STUBIN_STDERR` | `STUBIN_STATUS` |
| Rename to `stubin2` | `STUBIN2_STDOUT` | `STUBIN2_STDERR` | `STUBIN2_STATUS` |

Also, it is possible to fine-tune control using the `[Binary name]_EXPECT_` environment variables.

```console
$ export STUBIN_EXPECT_COND_HELP="'-h' in args"
$ export STUBIN_EXPECT_STDOUT_HELP="Usage: stubin [-h]
"
$ export STUBIN_EXPECT_COND_TEST="'test' in args"
$ export STUBIN_EXPECT_STDOUT_TEST="This is test.
"
$ stubin
$ stubin -h
Usage: stubin [-h]
$ stubin test
This is test.
```

## Install

**go install:**

```console
$ go install github.com/k1LoW/stubin@latest
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/stubin
```

**deb:**

``` console
$ export STUBIN_VERSION=X.X.X
$ curl -o stubin.deb -L https://github.com/k1LoW/stubin/releases/download/v$STUBIN_VERSION/stubin_$STUBIN_VERSION-1_amd64.deb
$ dpkg -i stubin.deb
```

**RPM:**

``` console
$ export STUBIN_VERSION=X.X.X
$ yum install https://github.com/k1LoW/stubin/releases/download/v$STUBIN_VERSION/stubin_$STUBIN_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export STUBIN_VERSION=X.X.X
$ curl -o stubin.apk -L https://github.com/k1LoW/stubin/releases/download/v$STUBIN_VERSION/stubin_$STUBIN_VERSION-1_amd64.apk
$ apk add stubin.apk
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/runn/releases)


