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
