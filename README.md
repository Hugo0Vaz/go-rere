# go-rere - Universal Behavior Testing Tool in Golang

This software is a reimplementation of [Tsoding's Python Record and Replay](https://github.com/tsoding/rere.py), however I plan to add a few bells and whistles.

## Record

1. Create a file that contains one shell command per line, can be named whatever you want but let's call it `test.list`.
2. To record the behaviour of each command for the list run:

```shell
$ grr record test.list
```

The command should generate a `test.list.bi` file in the [bi format](https://github.com/tsoding/bi-format). The contents of each shell command behaviour can be tweaked from command line options, but the standart recording is:

- stdout
- stderr
- returncode

The output file `test.list.bi` should be comitted to version control because the next time the `replay` command is used, the changes and differences between behaviours can be recorded.

## Replay

1. To check the behaviour against the recorded one:

```
$ grr replay test.list
```

The above command will run each command and compare via text the differences between different runs.
