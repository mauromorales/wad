# WAD

Words a Day is a CLI tool to help you write. If you use Vim make sure to check [vim-wad](https://github.com/mauromorales/vim-wad) plugin.

## Installation

```
go get -u github.com/mauromorales/wad/cmd/wad
```

## Usage

Initialize wad. Creates the `~/.wad/progress.json` file and folder if they don't exist and sets the goal to 500 words.

```
$ wad --config-dir=~/.wad init --goal=500
```

Set a 750 words a day goal

```
$ wad goal 750
```

Track files. Compares the files to the files.json database.

```
$ wad track path/to/file.txt [path/to/other/file.txt]
```

Move a file

```
$ wad mv path/to/file.txt path/to/new/location
```

Current progress

```
$ wad progress
| Day        | Words | Target | Status |
|------------|-------|--------|--------|
| 2017-02-01 | 300   | 500    | -      |
| 2017-02-02 | 500   | 500    | =      |
| 2017-02-03 | 600   | 500    | +      |
| 2017-02-04 | 750   | 750    | =      |
```

## Contributing

You need to have ginkgo and gomega installed in order to run the tests

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
```

Then just run the following command at the root directory to run all tests

```
$ ginkgo -r
```

## License

WAD is released under the [MIT License](./LICENSE)
