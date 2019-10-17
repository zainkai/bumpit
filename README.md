```txt
NAME:
   bumpit - Auto update Major/Minor/Patch of any JavaScript project!

USAGE:
   bumpit [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Kevin Turkington (Zainkai)

COMMANDS:
   Patch, p, patch   Update patch version
   Minor, mi, minor  Update minor version
   Major, ma, major  Update major version
   help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value, --path value, -p value  target file (default: "./package.json")
   --commit, -c                                    commits version to git
   --help, -h                                      show help
   --version, -v                                   print the version
```

## Examples:
### Update patch version
```
  bumpit patch
```
### Update minor version for specific file
```
  bumpit -file ./foo/bar/thing.json minor
```
### Update major version and commit to github
```
  bumpit -c major
```
### Update patch for file and commit to github
```
  bumpit -c -file ./foo/bar/thing.json patch
```

## Installation
TBD