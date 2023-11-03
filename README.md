# Go-Comment CLI Tool

The Go-Comment CLI tool is designed to help you automatically add descriptive comments to your code.

## Features

- Automatically generates descriptive comments for your code.
- Supports various file types, making it versatile.

## Installation

To install the Go-Comment CLI tool, follow these steps:

Clone this repository:
```shell
git clone https://github.com/lucasnevespereira/go-comment
```
Navigate to the cloned repository:
```shell
cd go-comment
```
Build the CLI tool:

```shell
go build -o go-comment
```

Move the binary to a directory included in your system's PATH (replace with your desired location):
```shell
mv go-comment /usr/local/bin/
```

## Usage
Basic Usage
To add comments to your Go code, use the following command:

```shell
go-comment -k your_api_key -i input.go -o output.go
```

- `-k` or `--api-key`: Your OpenAI API Key.
- `-i` or `--input`: The input code file.
- `-o` or `--output`: The output code file with comments.

The Go-Comment CLI tool is versatile and supports various file types. You can specify the file type to adjust the comments format accordingly. For example, if you're working with a TypeScript file:

```shell
go-comment -k your_api_key -i input.ts -o output.ts
```
The tool will adapt the comments format to your file type.

### Examples
Here are some example usages:

To add comments to a Go code file and save it as output.go:
```shell
go-comment -k your_api_key -i input.go -o output.go
```

To add comments to a TypeScript code file and save it as output.ts:
```shell
go-comment -k your_api_key -i input.ts -o output.ts
```

### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.