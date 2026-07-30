<a id="sql-selectinto"></a>

# SELECT INTO

define a new table from the results of a query

## Synopsis


```

[ WITH [ RECURSIVE ] WITH_QUERY [, ...] ]
SELECT [ ALL | DISTINCT [ ON ( EXPRESSION [, ...] ) ] ]
    [ { * | EXPRESSION [ [ AS ] OUTPUT_NAME ] } [, ...] ]
    INTO [ TEMPORARY | TEMP | UNLOGGED ] [ TABLE ] NEW_TABLE
    [ FROM FROM_ITEM [, ...] ]
    [ WHERE CONDITION ]
    [ GROUP BY { ALL | [ ALL | DISTINCT ] GROUPING_ELEMENT [, ...] } ]
    [ HAVING CONDITION ]
    [ WINDOW WINDOW_NAME AS ( WINDOW_DEFINITION ) [, ...] ]
    [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] SELECT ]
    [ ORDER BY EXPRESSION [ ASC | DESC | USING OPERATOR ] [ NULLS { FIRST | LAST } ] [, ...] ]
    [ LIMIT { COUNT | ALL } ]
    [ OFFSET START [ ROW | ROWS ] ]
    [ FETCH { FIRST | NEXT } [ COUNT ] { ROW | ROWS } { ONLY | WITH TIES } ]
    [ FOR { UPDATE | NO KEY UPDATE | SHARE | KEY SHARE } [ OF FROM_REFERENCE [, ...] ] [ NOWAIT | SKIP LOCKED ] [...] ]
```


## Description


 `SELECT INTO` creates a new table and fills it with data computed by a query. The data is not returned to the client, as it is with a normal `SELECT`. The new table's columns have the names and data types associated with the output columns of the `SELECT`.


## Parameters


`TEMPORARY` or `TEMP`
:   If specified, the table is created as a temporary table. Refer to [sql-createtable](create-table.md#sql-createtable) for details.

`UNLOGGED`
:   If specified, the table is created as an unlogged table. Refer to [sql-createtable](create-table.md#sql-createtable) for details.

*new_table*
:   The name (optionally schema-qualified) of the table to be created.


 All other parameters are described in detail under [sql-select](select.md#sql-select).


## Notes


 [`CREATE TABLE AS`](create-table-as.md#sql-createtableas) is functionally similar to `SELECT INTO`. `CREATE TABLE AS` is the recommended syntax, since this form of `SELECT INTO` is not available in ECPG or PL/pgSQL, because they interpret the `INTO` clause differently. Furthermore, `CREATE TABLE AS` offers a superset of the functionality provided by `SELECT INTO`.


 In contrast to `CREATE TABLE AS`, `SELECT INTO` does not allow specifying properties like a table's access method with [USING method](create-table.md#sql-createtable-method) or the table's tablespace with [TABLESPACE tablespace_name](create-table.md#sql-createtable-tablespace). Use `CREATE TABLE AS` if necessary. Therefore, the default table access method is chosen for the new table. See [default_table_access_method](../../server-administration/server-configuration/client-connection-defaults.md#guc-default-table-access-method) for more information.


## Examples


 Create a new table `films_recent` consisting of only recent entries from the table `films`:

```sql

SELECT * INTO films_recent FROM films WHERE date_prod >= '2002-01-01';
```


## Compatibility


 The SQL standard uses `SELECT INTO` to represent selecting values into scalar variables of a host program, rather than creating a new table. This indeed is the usage found in ECPG (see [ECPG — Embedded SQL in C](../../client-interfaces/ecpg-embedded-sql-in-c/index.md#ecpg)) and PL/pgSQL (see [PL/pgSQL — SQL Procedural Language](../../server-programming/pl-pgsql-sql-procedural-language/index.md#plpgsql)). The PostgreSQL usage of `SELECT INTO` to represent table creation is historical. Some other SQL implementations also use `SELECT INTO` in this way (but most SQL implementations support `CREATE TABLE AS` instead). Apart from such compatibility considerations, it is best to use `CREATE TABLE AS` for this purpose in new code.


## See Also
  [sql-createtableas](create-table-as.md#sql-createtableas)
