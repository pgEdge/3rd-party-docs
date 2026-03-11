<a id="sql-droptransform"></a>

# DROP TRANSFORM

remove a transform

## Synopsis


```

DROP TRANSFORM [ IF EXISTS ] FOR TYPE_NAME LANGUAGE LANG_NAME [ CASCADE | RESTRICT ]
```
 <a id="sql-droptransform-description"></a>

## Description


 `DROP TRANSFORM` removes a previously defined transform.


 To be able to drop a transform, you must own the type and the language. These are the same privileges that are required to create a transform.


## Parameters


`IF EXISTS`
:   Do not throw an error if the transform does not exist. A notice is issued in this case.

*type_name*
:   The name of the data type of the transform.

*lang_name*
:   The name of the language of the transform.

`CASCADE`
:   Automatically drop objects that depend on the transform, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the transform if any objects depend on it. This is the default.
 <a id="sql-droptransform-examples"></a>

## Examples


 To drop the transform for type `hstore` and language `plpython3u`:

```sql

DROP TRANSFORM FOR hstore LANGUAGE plpython3u;
```
 <a id="sql-droptransform-compat"></a>

## Compatibility


 This form of `DROP TRANSFORM` is a PostgreSQL extension. See [sql-createtransform](create-transform.md#sql-createtransform) for details.


## See Also
  [sql-createtransform](create-transform.md#sql-createtransform)
