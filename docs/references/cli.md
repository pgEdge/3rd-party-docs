<a id="cli"></a>

# CLI

PostgREST provides a CLI with the options listed below:

```text
Usage: postgrest [-v|--version] [-e|--example] [--dump-config | --dump-schema | --ready]
               [FILENAME]

  PostgREST / create a REST API to an existing Postgres
  database

Available options:
  -h,--help                Show this help text
  -v,--version             Show the version information
  -e,--example             Show an example configuration file
  --dump-config            Dump loaded configuration and exit
  --dump-schema            Dump loaded schema as JSON and exit (for debugging,
                           output structure is unstable)
  --ready                  Checks the health of PostgREST by doing a request on
                           the admin server /ready endpoint
  FILENAME                 Path to configuration file
```

## FILENAME

Runs PostgREST with the given [Config File](configuration.md#file_config).

## Help

```bash
$ postgrest --help
```

Shows all the options available.

## Version

```bash
$ postgrest --version
```

Prints the PostgREST version.

## Example

```bash
$ postgrest --example
```

Shows example configuration settings.

## Dump Config

```bash
$ postgrest --dump-config
```

Dumps the loaded [Configuration](configuration.md#configuration) values, considering the configuration file, environment variables and [In-Database Configuration](configuration.md#in_db_config).

## Dump Schema

```bash
$ postgrest --dump-schema
```

Dumps the schema cache in JSON format.

## Ready Flag

Makes a request to the `/ready` endpoint of the [Admin Server](admin_server.md#admin_server). It exits with a return code of `0` on success and `1` on failure.

```bash
$ postgrest --ready
OK: http://localhost:3001/ready
```

!!! note

    The `--ready` flag cannot be used when [server-host](configuration.md#server-host) is configured with special hostnames. We suggest to change it to `localhost`.
