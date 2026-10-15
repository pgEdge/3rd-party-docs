<a id="sql-alterroutine"></a>

# ALTER ROUTINE

change the definition of a routine

## Synopsis


```

ALTER ROUTINE NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ]
    ACTION [ ... ] [ RESTRICT ]
ALTER ROUTINE NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ]
    RENAME TO NEW_NAME
ALTER ROUTINE NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ]
    OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER ROUTINE NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ]
    SET SCHEMA NEW_SCHEMA
ALTER ROUTINE NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ]
    [ NO ] DEPENDS ON EXTENSION EXTENSION_NAME

where ACTION is one of:

    IMMUTABLE | STABLE | VOLATILE
    [ NOT ] LEAKPROOF
    [ EXTERNAL ] SECURITY INVOKER | [ EXTERNAL ] SECURITY DEFINER
    PARALLEL { UNSAFE | RESTRICTED | SAFE }
    COST EXECUTION_COST
    ROWS RESULT_ROWS
    SET CONFIGURATION_PARAMETER { TO | = } { VALUE | DEFAULT }
    SET CONFIGURATION_PARAMETER FROM CURRENT
    RESET CONFIGURATION_PARAMETER
    RESET ALL
```


## Description


 `ALTER ROUTINE` changes the definition of a routine, which can be an aggregate function, a normal function, or a procedure. See under [sql-alteraggregate](alter-aggregate.md#sql-alteraggregate), [sql-alterfunction](alter-function.md#sql-alterfunction), and [sql-alterprocedure](alter-procedure.md#sql-alterprocedure) for the description of the parameters, more examples, and further details.


## Examples


 To rename the routine `foo` for type `integer` to `foobar`:

```sql

ALTER ROUTINE foo(integer) RENAME TO foobar;
```
 This command will work independent of whether `foo` is an aggregate, function, or procedure.


## Compatibility


 This statement is partially compatible with the `ALTER ROUTINE` statement in the SQL standard. See under [sql-alterfunction](alter-function.md#sql-alterfunction) and [sql-alterprocedure](alter-procedure.md#sql-alterprocedure) for more details. Allowing routine names to refer to aggregate functions is a PostgreSQL extension.


## See Also
  [sql-alteraggregate](alter-aggregate.md#sql-alteraggregate), [sql-alterfunction](alter-function.md#sql-alterfunction), [sql-alterprocedure](alter-procedure.md#sql-alterprocedure), [sql-droproutine](drop-routine.md#sql-droproutine)

 Note that there is no `CREATE ROUTINE` command.
