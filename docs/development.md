# Development

Nix with flakes is required.

## Set up Nix

First, [Install Nix](https://nixos.org/download.html).

### Enable flakes

Edit `/etc/nix/nix.conf`, add:

```nix.conf
experimental-features = nix-command flakes
```

and restart the daemon, on macOS:

```
sudo launchctl kickstart -k system/org.nixos.nix-daemon
```

### Use binary cache

If you are a repo member, you can skip building for hours and download packages instead by configuring the `nxpg` binary cache.

1. Go to https://app.cachix.org, log in with Github and create a **personal auth token**.

2. Use auth token:
  ```sh
  nix run nixpkgs#cachix authtoken <token>
  ```

3. Use the cache:
  ```sh
  sudo nix run nixpkgs#cachix use nxpg
  ```


!!! danger
    DO NOT use `trusted-users` in `/etc/nix/nix.conf` as it [grants root without password](https://nix.dev/manual/nix/stable/command-ref/conf-file.html#conf-trusted-users). Instead, add the binary cache to `extra-substituters` and `extra-trusted-public-keys`.


## Testing

For testing the module locally, execute:

```bash
# Open a development shell with all deps and tools present
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

## Running a single test

`xpg test` always runs the full test suite. To narrow it down to one test, set
its name in the `REGRESS` variable via `MAKEFLAGS`. For example, to run just the
`test/sql/permission_hints.sql` test:

```bash
$ MAKEFLAGS="REGRESS=permission_hints" xpg -v 15 test
```

## Regress testing against PostgreSQL core

Since supautils modifies default postgres behavior with hooks, we need to test exactly what it changes and see if we don't break existing functionality.
For this you can use:

```bash
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

