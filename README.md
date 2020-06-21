# puffer

This tool checks for identical files stored in the target directory under different file names.

## Install

If you are on another platform, you can download a binary from our release page and place it in $PATH directory.

Or you can use dep (you need to use go1.12 or later),

```bash
$ dep ensure
$ go install -v ./...
```

## Usage
```bash
$ puffer [command flags]
command flags:
    -in filepath *
        Specify the name of the target directory.(default nil)
    --out filepath
        Specify the file name of the report.
        Required for search mode.(default nil)
    --mode string
        Specifies the mode in which to search or delete.  [search = search only, remove = remove and search] (default search)
    --num int
        Specify the number of target files to be processed simultaneously.
        Normally this is not necessary and is determined according to the number of CPU cores.(default 1)
```

