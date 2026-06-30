# Development

[Nix](https://nixos.org/download.html) is required to set up the environment.

## Testing

For testing the module locally, execute:

```bash
# might take a while in downloading all the dependencies
$ nix-shell

# test on pg 13
$ xpg -v 13 test

# test on pg 14
$ xpg -v 14 test

# you can also test manually with
$ xpg -v 13 psql -U rolecreator
```

To test with a postgres built with assertions enabled:

```bash
$ nix-shell

# test on pg 13
$ xpg -v 17 --cassert test
```

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

