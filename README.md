# rfind

A command-line tool that searches for files in reverse order (i.e. to ancestor direction).

## Usage

```
Usage:
  rfind [OPTIONS] ORIGIN_PATH TARGETS...
Options:
  -dir-only
        Find only directory
  -file-only
        Find only file
  -limit uint
        Limit the number of found items (default: unlimited)
  -max-depth-from-root uint
        Max number of path depth to search from root (default: unlimited)
  -max-upper-depth uint
        Max number of path depth to search from the ORIGIN_PATH to ancestor direction (default: unlimited)
```

### Examples

When the directory structure is as follows:

```
/
├── bar
├── buz
│    ├── qux
│    │    └── target.txt
│    └── target.txt
├── foo
│    └── target.txt
└── target.txt
```

And the current directory is `/buz/qux`, the following command will find `target.txt` in the ancestor directories:

```
$ rfind . target.txt # `rfind /buz/qux target.txt` also works as well
/buz/qux/target.txt
/buz/target.txt
/target.txt
```

And the `targets` can be accepted the multi leaved paths:

```
$ rfind /buz/qux foo/target.txt
/foo/target.txt
```

## Author

moznion (<moznion@mail.moznion.net>)

## License

MIT

