<a id="sql-droplanguage"></a>

# DROP LANGUAGE

remove a procedural language

## Synopsis


```

DROP [ PROCEDURAL ] LANGUAGE [ IF EXISTS ] NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP LANGUAGE` removes the definition of a previously registered procedural language. You must be a superuser or the owner of the language to use `DROP LANGUAGE`.


!!! note

    As of PostgreSQL 9.1, most procedural languages have been made into “extensions”, and should therefore be removed with [`DROP EXTENSION`](drop-extension.md#sql-dropextension) not `DROP LANGUAGE`.


## Parameters


`IF EXISTS`
:   Do not throw an error if the language does not exist. A notice is issued in this case.

*name*
:   The name of an existing procedural language.

`CASCADE`
:   Automatically drop objects that depend on the language (such as functions in the language), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the language if any objects depend on it. This is the default.


## Examples


 This command removes the procedural language `plsample`:

```sql

DROP LANGUAGE plsample;
```


## Compatibility


 There is no `DROP LANGUAGE` statement in the SQL standard.


## See Also
  [sql-alterlanguage](alter-language.md#sql-alterlanguage), [sql-createlanguage](create-language.md#sql-createlanguage)
