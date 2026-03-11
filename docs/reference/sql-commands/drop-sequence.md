<a id="sql-dropsequence"></a>

# DROP SEQUENCE

remove a sequence

## Synopsis


```

DROP SEQUENCE [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP SEQUENCE` removes sequence number generators. A sequence can only be dropped by its owner or a superuser.


## Parameters


`IF EXISTS`
:   Do not throw an error if the sequence does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of a sequence.

`CASCADE`
:   Automatically drop objects that depend on the sequence, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the sequence if any objects depend on it. This is the default.


## Examples


 To remove the sequence `serial`:

```sql

DROP SEQUENCE serial;
```


## Compatibility


 `DROP SEQUENCE` conforms to the SQL standard, except that the standard only allows one sequence to be dropped per command, and apart from the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also
  [sql-createsequence](create-sequence.md#sql-createsequence), [sql-altersequence](alter-sequence.md#sql-altersequence)
