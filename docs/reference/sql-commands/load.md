<a id="sql-load"></a>

# LOAD

load a shared library file

## Synopsis


```

LOAD 'FILENAME'
```
 <a id="sql-load-description"></a>

## Description


 This command loads a shared library file into the PostgreSQL server's address space. If the file has been loaded already, the command does nothing. Shared library files that contain C functions are automatically loaded whenever one of their functions is called. Therefore, an explicit `LOAD` is usually only needed to load a library that modifies the server's behavior through “hooks” rather than providing a set of functions.


 The library file name is typically given as just a bare file name, which is sought in the server's library search path (set by [dynamic_library_path](../../server-administration/server-configuration/client-connection-defaults.md#guc-dynamic-library-path)). Alternatively it can be given as a full path name. In either case the platform's standard shared library file name extension may be omitted. See [Dynamic Loading](../../server-programming/extending-sql/c-language-functions.md#xfunc-c-dynload) for more information on this topic.


 Non-superusers can only apply `LOAD` to library files located in `$libdir/plugins/` — the specified *filename* must begin with exactly that string. (It is the database administrator's responsibility to ensure that only “safe” libraries are installed there.)
 <a id="sql-load-compat"></a>

## Compatibility


 `LOAD` is a PostgreSQL extension.


## See Also


 [sql-createfunction](create-function.md#sql-createfunction)
