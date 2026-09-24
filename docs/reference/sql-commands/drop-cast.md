<a id="sql-dropcast"></a>

# DROP CAST

remove a cast

## Synopsis


```

DROP CAST [ IF EXISTS ] (SOURCE_TYPE AS TARGET_TYPE) [ CASCADE | RESTRICT ]
```
 <a id="sql-dropcast-description"></a>

## Description


 `DROP CAST` removes a previously defined cast.


 To be able to drop a cast, you must own the source or the target data type. These are the same privileges that are required to create a cast.


## Parameters


`IF EXISTS`
:   Do not throw an error if the cast does not exist. A notice is issued in this case.

*source_type*
:   The name of the source data type of the cast.

*target_type*
:   The name of the target data type of the cast.

`CASCADE`, `RESTRICT`
:   These key words do not have any effect, since there are no dependencies on casts.
 <a id="sql-dropcast-examples"></a>

## Examples


 To drop the cast from type `text` to type `int`:

```sql

DROP CAST (text AS int);
```
 <a id="sql-dropcast-compat"></a>

## Compatibility


 The `DROP CAST` command conforms to the SQL standard.


## See Also
  [sql-createcast](create-cast.md#sql-createcast)
