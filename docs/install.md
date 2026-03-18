<a id="installation"></a>

# Installation

Psycopg is a [PostgreSQL](https://www.postgresql.org/) adapter for the [Python](https://www.python.org/) programming language. It is a wrapper for the [libpq](https://www.postgresql.org/docs/current/static/libpq.html), the official PostgreSQL client library.
<a id="PostgreSQL"></a>
<a id="Python"></a>
<a id="binary-packages"></a>

## Quick Install

For most operating systems, the quickest way to install Psycopg is using the [wheel](https://pythonwheels.com/) package available on [PyPI](https://pypi.org/project/psycopg2-binary/):

```console
 $ pip install psycopg2-binary
```

This will install a pre-compiled binary version of the module which does not require the build or runtime prerequisites described below. Make sure to use an up-to-date version of `pip` (you can upgrade it using something like `pip install -U pip`).

You may then import the `psycopg2` package, as usual:

```python
 import psycopg2

 # Connect to your postgres DB
 conn = psycopg2.connect("dbname=test user=postgres")

 # Open a cursor to perform database operations
 cur = conn.cursor()

 # Execute a query
 cur.execute("SELECT * FROM my_data")

 # Retrieve query results
 records = cur.fetchall()
```
<a id="PyPI"></a>
<a id="wheel"></a>

### psycopg vs psycopg-binary

The `psycopg2-binary` package is meant for beginners to start playing with Python and PostgreSQL without the need to meet the build requirements.

If you are the maintainer of a published package depending on `psycopg2` you shouldn't use `psycopg2-binary` as a module dependency. **For production use you are advised to use the source distribution.**

The binary packages come with their own versions of a few C libraries, among which `libpq` and `libssl`, which will be used regardless of other libraries available on the client: upgrading the system libraries will not upgrade the libraries used by `psycopg2`. Please build `psycopg2` from source if you want to maintain binary upgradeability.

!!! warning

    The `psycopg2` wheel package comes packaged, among the others, with its own `libssl` binary. This may create conflicts with other extension modules binding with `libssl` as well, for instance with the Python `ssl` module: in some cases, under concurrency, the interaction between the two libraries may result in a segfault. In case of doubts you are advised to use a package built from source.
<a id="disable-wheel"></a>

### Change in binary packages between Psycopg 2.7 and 2.8

In version 2.7.x, `pip install psycopg2` would have tried to install automatically the binary package of Psycopg. Because of concurrency problems binary packages have displayed, `psycopg2-binary` has become a separate package, and from 2.8 it has become the only way to install the binary package.

If you are using Psycopg 2.7 and you want to disable the use of wheel binary packages, relying on the system libraries available on your client, you can use the `pip` ``--no-binary`` option_, e.g.:

```console
 $ pip install --no-binary :all: psycopg2
```
<a id="_"></a>

which can be specified in your `requirements.txt` files too, e.g. use:

```none
 psycopg2>=2.7,<2.8 --no-binary psycopg2
```

to use the last bugfix release of the `psycopg2` 2.7 package, specifying to always compile it from source. Of course in this case you will have to meet the [build prerequisites](#build-prerequisites).

## Prerequisites

The current `psycopg2` implementation supports:

- Python versions from 3.9 to 3.14

- PostgreSQL server versions from 7.4 to 18

- PostgreSQL client library version from 9.1

!!! note

    Not all the psycopg2 versions support all the supported Python versions.

    Please see the [release notes](news.md#news) to verify when the support for a new Python version was added and when the support for an old Python version was removed.
<a id="build-prerequisites"></a>

### Build prerequisites

The build prerequisites are to be met in order to install Psycopg from source code, from a source distribution package, [GitHub](https://github.com/psycopg/psycopg2) or from PyPI.
<a id="GitHub"></a>

Psycopg is a C wrapper around the [libpq](https://www.postgresql.org/docs/current/static/libpq.html) PostgreSQL client library. To install it from sources you will need:

- A C compiler.

- The Python header files. They are usually installed in a package such as
    **python-dev** or **python3-dev**. A message such as *error: Python.h: No
    such file or directory* is an indication that the Python headers are
    missing.

- The libpq header files. They are usually installed in a package such as
    **libpq-dev**. If you get an *error: libpq-fe.h: No such file or directory*
    you are missing them.

- The `pg_config` program: it is usually installed by the **libpq-dev** package but sometimes it is not in a `PATH` directory. Having it in the `PATH` greatly streamlines the installation, so try running `pg_config --version`: if it returns an error or an unexpected version number then locate the directory containing the `pg_config` shipped with the right libpq version (usually `/usr/lib/postgresql/X.Y/bin/`) and add it to the `PATH`:

    ```console
    $ export PATH=/usr/lib/postgresql/X.Y/bin/:$PATH
    ```

    You only need `pg_config` to compile `psycopg2`, not for its regular usage.

Once everything is in place it's just a matter of running the standard:

```console
 $ pip install psycopg2
```

or, from the directory containing the source code:

```console
 $ python setup.py build
 $ python setup.py install
```

### Runtime requirements

Unless you compile `psycopg2` as a static library, or you install it from a self-contained wheel package, it will need the [libpq](https://www.postgresql.org/docs/current/static/libpq.html) library at runtime (usually distributed in a `libpq.so` or `libpq.dll` file).  `psycopg2` relies on the host OS to find the library if the library is installed in a standard location there is usually no problem; if the library is in a non-standard location you will have to tell Psycopg how to find it, which is OS-dependent (for instance setting a suitable `LD_LIBRARY_PATH` on Linux).

!!! note

    The libpq header files used to compile `psycopg2` should match the version of the library linked at runtime. If you get errors about missing or mismatching libraries when importing `psycopg2` check (e.g. using `ldd`) if the module `psycopg2/_psycopg.so` is linked to the right `libpq.so`.

!!! note

    Whatever version of libpq `psycopg2` is compiled with, it will be possible to connect to PostgreSQL servers of any supported version: just install the most recent libpq version or the most practical, without trying to match it to the version of the PostgreSQL server you will have to connect to.

## Non-standard builds

If you have less standard requirements such as:

- creating a [debug build](#debug-build),

- using `pg_config` not in the `PATH`,

then take a look at the `setup.cfg` file.

Some of the options available in `setup.cfg` are also available as command line arguments of the `build_ext` sub-command. For instance you can specify an alternate `pg_config` location using:

```console
 $ python setup.py build_ext --pg-config /path/to/pg_config build
```

Use `python setup.py build_ext --help` to get a list of the options supported.
<a id="debug-build"></a>

### Creating a debug build

In case of problems, Psycopg can be configured to emit detailed debug messages, which can be very useful for diagnostics and to report a bug. In order to create a debug package:

- `Download`__ and unpack the Psycopg *source package* (the `.tar.gz`
    package).

- Edit the `setup.cfg` file adding the `PSYCOPG_DEBUG` flag to the
    `define` option.

- [Compile and install](#build-prerequisites) the package.

- Set the `PSYCOPG_DEBUG` environment variable:

```console
 $ export PSYCOPG_DEBUG=1
```

- Run your program (making sure that the `psycopg2` package imported is the
    one you just compiled and not e.g. the system one): you will have a copious
    stream of informations printed on stderr.
<a id="_"></a>

### Non-standard Python Implementation

The `psycopg2` package is the current mature implementation of the adapter: it is a C extension and as such it is only compatible with [CPython](https://en.wikipedia.org/wiki/CPython). If you want to use Psycopg on a different Python implementation (PyPy, Jython, IronPython) there is a couple of alternative:

- a `Ctypes port`__, but it is not as mature as the C implementation yet
    and it is not as feature-complete;

- a `CFFI port`__ which is currently more used and reported more efficient on
    PyPy, but please be careful of its version numbers because they are not
    aligned to the official psycopg2 ones and some features may differ.
<a id="PostgreSQL"></a>
<a id="Python"></a>
<a id="libpq"></a>
<a id="CPython"></a>
<a id="Ctypes"></a>
<a id="_"></a>
<a id="_"></a>
<a id="test-suite"></a>

## Running the test suite

Once `psycopg2` is installed you can run the test suite to verify it is working correctly. From the source directory, you can run:

```console
 $ python -c "import tests; tests.unittest.main(defaultTest='tests.test_suite')" --verbose
```

The tests run against a database called `psycopg2_test` on UNIX socket and the standard port. You can configure a different database to run the test by setting the environment variables:

- `PSYCOPG2_TESTDB`

- `PSYCOPG2_TESTDB_HOST`

- `PSYCOPG2_TESTDB_PORT`

- `PSYCOPG2_TESTDB_USER`

The database should already exist before running the tests.
<a id="other-problems"></a>

## If you still have problems

Try the following. *In order:*

- Read again the [Build prerequisites](#build-prerequisites).

- Read the [FAQ](faq.md#faq-compile).

- Google for `psycopg2` *your error message*. Especially useful the week
    after the release of a new OS X version.

- Write to the Mailing List

- If you think that you have discovered a bug, test failure or missing feature
    please raise a ticket in the [bug tracker](https://github.com/psycopg/psycopg2/issues)

- Complain on your blog or on Twitter that `psycopg2` is the worst package
    ever and about the quality time you have wasted figuring out the correct
    `ARCHFLAGS`. Especially useful from the Starbucks near you.
<a id="mailing list"></a>
<a id="bug tracker"></a>
