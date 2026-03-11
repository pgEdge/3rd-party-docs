<a id="sql-alterview"></a>

# ALTER VIEW

change the definition of a view

## Synopsis


```

ALTER VIEW [ IF EXISTS ] NAME ALTER [ COLUMN ] COLUMN_NAME SET DEFAULT EXPRESSION
ALTER VIEW [ IF EXISTS ] NAME ALTER [ COLUMN ] COLUMN_NAME DROP DEFAULT
ALTER VIEW [ IF EXISTS ] NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER VIEW [ IF EXISTS ] NAME RENAME [ COLUMN ] COLUMN_NAME TO NEW_COLUMN_NAME
ALTER VIEW [ IF EXISTS ] NAME RENAME TO NEW_NAME
ALTER VIEW [ IF EXISTS ] NAME SET SCHEMA NEW_SCHEMA
ALTER VIEW [ IF EXISTS ] NAME SET ( VIEW_OPTION_NAME [= VIEW_OPTION_VALUE] [, ... ] )
ALTER VIEW [ IF EXISTS ] NAME RESET ( VIEW_OPTION_NAME [, ... ] )
```


## Description


 `ALTER VIEW` changes various auxiliary properties of a view. (If you want to modify the view's defining query, use `CREATE OR REPLACE VIEW`.)


 You must own the view to use `ALTER VIEW`. To change a view's schema, you must also have `CREATE` privilege on the new schema. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the view's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the view. However, a superuser can alter ownership of any view anyway.)


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing view.

*column_name*
:   Name of an existing column.

*new_column_name*
:   New name for an existing column.

`IF EXISTS`
:   Do not throw an error if the view does not exist. A notice is issued in this case.

`SET`/`DROP DEFAULT`
:   These forms set or remove the default value for a column. A view column's default value is substituted into any `INSERT` or `UPDATE` command whose target is the view, before applying any rules or triggers for the view. The view's default will therefore take precedence over any default values from underlying relations.

*new_owner*
:   The user name of the new owner of the view.

*new_name*
:   The new name for the view.

*new_schema*
:   The new schema for the view.

<code>SET ( </code><em>view_option_name</em><code> [= </code><em>view_option_value</em><code>] [, ... ] )</code>, <code>RESET ( </code><em>view_option_name</em><code> [, ... ] )</code>
:   Sets or resets a view option. Currently supported options are:

    `check_option` (`enum`)
    :   Changes the check option of the view. The value must be `local` or `cascaded`.

    `security_barrier` (`boolean`)
    :   Changes the security-barrier property of the view. The value must be a Boolean value, such as `true` or `false`.

    `security_invoker` (`boolean`)
    :   Changes the security-invoker property of the view. The value must be a Boolean value, such as `true` or `false`.


## Notes


 For historical reasons, `ALTER TABLE` can be used with views too; but the only variants of `ALTER TABLE` that are allowed with views are equivalent to the ones shown above.


## Examples


 To rename the view `foo` to `bar`:

```sql

ALTER VIEW foo RENAME TO bar;
```


 To attach a default column value to an updatable view:

```sql

CREATE TABLE base_table (id int, ts timestamptz);
CREATE VIEW a_view AS SELECT * FROM base_table;
ALTER VIEW a_view ALTER COLUMN ts SET DEFAULT now();
INSERT INTO base_table(id) VALUES(1);  -- ts will receive a NULL
INSERT INTO a_view(id) VALUES(2);  -- ts will receive the current time
```


## Compatibility


 `ALTER VIEW` is a PostgreSQL extension of the SQL standard.


## See Also
  [sql-createview](create-view.md#sql-createview), [sql-dropview](drop-view.md#sql-dropview)
