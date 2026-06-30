# Development

[Nix](https://nixos.org/download.html) is required to set up the environment.

## Testing

For testing the module locally, execute:

```bash
# This will download all the dependencies from the cache (when prompted for trusting the nxpg cache answer yes)
$ nix develop

# test on pg 13
$ xpg -v 13 test

# test on pg 14
$ xpg -v 14 test

# you can also test manually with
$ xpg -v 13 psql -U rolecreator
```

To test with a postgres built with assertions enabled:

```bash
$ xpg -v 17 --cassert test
```

## Regress testing against PostgreSQL core

Since supautils modifies default postgres behavior with hooks, we need to test exactly what it changes and see if we don't break existing functionality.
For this you can use:

```
xpg -v 15 test-core
```

Works on pg 15, 16, 17 and 18.

## Coverage

For coverage, execute:

```bash
$ xpg -v 17 coverage
```

## Style

For automatic formatting of source and header files use:

```bash
$ supautils-style
```

