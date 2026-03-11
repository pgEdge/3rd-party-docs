<a id="sql-dropextension"></a>

# DROP EXTENSION

remove an extension

## Synopsis


```

DROP EXTENSION [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP EXTENSION` removes extensions from the database. Dropping an extension causes its member objects, and other explicitly dependent routines (see [sql-alterroutine](alter-routine.md#sql-alterroutine), the <code>DEPENDS ON EXTENSION </code><em>extension_name</em><code>
   </code> action), to be dropped as well.


 You must own the extension to use `DROP EXTENSION`.


## Parameters


`IF EXISTS`
:   Do not throw an error if the extension does not exist. A notice is issued in this case.

*name*
:   The name of an installed extension.

`CASCADE`
:   Automatically drop objects that depend on the extension, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   This option prevents the specified extensions from being dropped if other objects, besides these extensions, their members, and their explicitly dependent routines, depend on them. This is the default.


## Examples


 To remove the extension `hstore` from the current database:

```sql

DROP EXTENSION hstore;
```
 This command will fail if any of `hstore`'s objects are in use in the database, for example if any tables have columns of the `hstore` type. Add the `CASCADE` option to forcibly remove those dependent objects as well.


## Compatibility


 `DROP EXTENSION` is a PostgreSQL extension.


## See Also
  [sql-createextension](create-extension.md#sql-createextension), [sql-alterextension](alter-extension.md#sql-alterextension)
