<a id="sql-droprule"></a>

# DROP RULE

remove a rewrite rule

## Synopsis


```

DROP RULE [ IF EXISTS ] NAME ON TABLE_NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP RULE` drops a rewrite rule.


## Parameters


`IF EXISTS`
:   Do not throw an error if the rule does not exist. A notice is issued in this case.

*name*
:   The name of the rule to drop.

*table_name*
:   The name (optionally schema-qualified) of the table or view that the rule applies to.

`CASCADE`
:   Automatically drop objects that depend on the rule, and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the rule if any objects depend on it. This is the default.


## Examples


 To drop the rewrite rule `newrule`:

```sql

DROP RULE newrule ON mytable;
```


## Compatibility


 `DROP RULE` is a PostgreSQL language extension, as is the entire query rewrite system.


## See Also
  [sql-createrule](create-rule.md#sql-createrule), [sql-alterrule](alter-rule.md#sql-alterrule)
