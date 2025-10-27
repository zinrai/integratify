# integratify

A command-line tool for validating JSON data against CUE schemas.

## Overview

integratify validates JSON configuration files and API responses using [CUE (Configure Unify Execute)](https://cuelang.org/) schemas. It reads JSON data and verifies it conforms to the structure and constraints defined in a CUE schema file.

This tool is designed for validating API responses during system migrations and infrastructure changes, where you need to ensure services continue functioning correctly after upgrades.

## Installation

```bash
$ go install github.com/zinrai/integratify@latest
```

## Usage

```bash
$ integratify -schema=<schema.cue> -config=<config.json>
```

Options:
- `-schema`: Path to CUE schema file (required)
- `-config`: Path to JSON file to validate (required)
- `-version`: Show version information

## Example

See the `example` directory for sample schema and configuration files.

Valid configuration - exits with status 0

```bash
$ integratify -schema=example/schema.cue -config=example/valid.json
```

Invalid configuration - exits with status 1

```bash
$ integratify -schema=example/schema.cue -config=example/invalid.json
```

## Exit Codes

- `0`: Validation successful
- `1`: Validation failed or error occurred

## Use Cases

Validate that API responses match expected schemas during system migrations:

```bash
$ curl -o api.json https://api.example.com/config 
$ integratify -schema schema.cue -config api.json
```

## License

This project is licensed under the [MIT License](./LICENSE).
