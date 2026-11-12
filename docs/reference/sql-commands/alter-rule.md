<a id="sql-alterrule"></a>

# ALTER RULE

change the definition of a rule

## Synopsis


```

ALTER RULE NAME ON TABLE_NAME RENAME TO NEW_NAME
```


## Description


 `ALTER RULE` changes properties of an existing rule. Currently, the only available action is to change the rule's name.


 To use `ALTER RULE`, you must own the table or view that the rule applies to.


## Parameters


*name*
:   The name of an existing rule to alter.

*table_name*
:   The name (optionally schema-qualified) of the table or view that the rule applies to.

*new_name*
:   The new name for the rule.


## Examples


 To rename an existing rule:

```sql

ALTER RULE notify_all ON emp RENAME TO notify_me;
```


## Compatibility


 `ALTER RULE` is a PostgreSQL language extension, as is the entire query rewrite system.


## See Also
  [sql-createrule](create-rule.md#sql-createrule), [sql-droprule](drop-rule.md#sql-droprule)
