# Installation

- Copy this directory to contrib/ in your PostgreSQL source tree.

- Run `make && make install`

- Edit your postgresql.conf file, and modify the shared_preload_libraries config
  option to look like:

  `shared_preload_libraries = '$libdir/plugin_debugger'`

- Restart PostgreSQL for the new setting to take effect.

- Run the following command in the database or databases that you wish to
  debug functions in:

  `CREATE EXTENSION pldbgapi;`

  (on server versions older than 9.1, you must instead run the pldbgapi--1.1.sql
  script directly using psql).

