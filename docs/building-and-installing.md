# Building and Installing

## Linux and macOS
The module can be built using the PGXS framework:

- Unpack the file archive in a suitable directory.
- Ensure the PATH environment variable includes the directory containing the
  pg_config binary for the PostgreSQL installation you wish to build against.
- Compile and install the code.

For example:

    tar -zxvf system_stats-1.0.tar.gz
    cd system_stats-1.0
    PATH="/usr/local/pgsql/bin:$PATH" make USE_PGXS=1
    sudo PATH="/usr/local/pgsql/bin:$PATH" make install USE_PGXS=1

## Windows
The module built using the Visual Studio project file:

- Unpack the extensions files in $PGSRC/contrib/system_stats
- Set PG_INCLUDE_DIR and PG_LIB_DIR environment variables to make sure the
  PostgreSQL include and lib directories can be found for compilation. For
  example:

        PG_INCLUDE_DIR=C:\Program Files\PostgreSQL\12\include
        PG_LIB_DIR=C:\Program Files\PostgreSQL\12\lib

- Open the Visual Studio project file "system_stats.vcxproj" and build the
  project.

## Installing the Extension
Once the code has been built and installed, you can install the extension in
a database using the following SQL command:

    CREATE EXTENSION system_stats;

## Security
Due to the nature of the information returned by these functions, access is
restricted to superusers and members of the monitor_system_stats role which
will be created when the extension is installed. The monitor_system_stats 
role will not be removed when you run DROP EXTENSION. This means that any 
users or roles that were granted permissions to the monitor_system_stats role 
will still have those permissions even after the extension has been dropped. 
To allow users to access the functions without granting them superuser access, 
add them to the monitor_system_stats role. For example:

    GRANT monitor_system_stats to nagios;

User can grant execute rights for all the below functions to `pg_monitor` role explicitly

e.g.

    GRANT EXECUTE ON FUNCTION pg_sys_os_info() TO pg_monitor;

