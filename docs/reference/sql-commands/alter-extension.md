<a id="sql-alterextension"></a>

# ALTER EXTENSION

change the definition of an extension

## Synopsis


```

ALTER EXTENSION NAME UPDATE [ TO NEW_VERSION ]
ALTER EXTENSION NAME SET SCHEMA NEW_SCHEMA
ALTER EXTENSION NAME ADD MEMBER_OBJECT
ALTER EXTENSION NAME DROP MEMBER_OBJECT

where MEMBER_OBJECT is:

  ACCESS METHOD OBJECT_NAME |
  AGGREGATE AGGREGATE_NAME ( AGGREGATE_SIGNATURE ) |
  CAST (SOURCE_TYPE AS TARGET_TYPE) |
  COLLATION OBJECT_NAME |
  CONVERSION OBJECT_NAME |
  DOMAIN OBJECT_NAME |
  EVENT TRIGGER OBJECT_NAME |
  FOREIGN DATA WRAPPER OBJECT_NAME |
  FOREIGN TABLE OBJECT_NAME |
  FUNCTION FUNCTION_NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ] |
  MATERIALIZED VIEW OBJECT_NAME |
  OPERATOR OPERATOR_NAME (LEFT_TYPE, RIGHT_TYPE) |
  OPERATOR CLASS OBJECT_NAME USING INDEX_METHOD |
  OPERATOR FAMILY OBJECT_NAME USING INDEX_METHOD |
  [ PROCEDURAL ] LANGUAGE OBJECT_NAME |
  PROCEDURE PROCEDURE_NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ] |
  ROUTINE ROUTINE_NAME [ ( [ [ ARGMODE ] [ ARGNAME ] ARGTYPE [, ...] ] ) ] |
  SCHEMA OBJECT_NAME |
  SEQUENCE OBJECT_NAME |
  SERVER OBJECT_NAME |
  TABLE OBJECT_NAME |
  TEXT SEARCH CONFIGURATION OBJECT_NAME |
  TEXT SEARCH DICTIONARY OBJECT_NAME |
  TEXT SEARCH PARSER OBJECT_NAME |
  TEXT SEARCH TEMPLATE OBJECT_NAME |
  TRANSFORM FOR TYPE_NAME LANGUAGE LANG_NAME |
  TYPE OBJECT_NAME |
  VIEW OBJECT_NAME

and AGGREGATE_SIGNATURE is:

* |
[ ARGMODE ] [ ARGNAME ] ARGTYPE [ , ... ] |
[ [ ARGMODE ] [ ARGNAME ] ARGTYPE [ , ... ] ] ORDER BY [ ARGMODE ] [ ARGNAME ] ARGTYPE [ , ... ]
```


## Description


 `ALTER EXTENSION` changes the definition of an installed extension. There are several subforms:

`UPDATE`
:   This form updates the extension to a newer version. The extension must supply a suitable update script (or series of scripts) that can modify the currently-installed version into the requested version.

`SET SCHEMA`
:   This form moves the extension's objects into another schema. The extension has to be *relocatable* for this command to succeed.

<code>ADD </code><em>member_object</em>
:   This form adds an existing object to the extension. This is mainly useful in extension update scripts. The object will subsequently be treated as a member of the extension; notably, it can only be dropped by dropping the extension.

<code>DROP </code><em>member_object</em>
:   This form removes a member object from the extension. This is mainly useful in extension update scripts. The object is not dropped, only disassociated from the extension.
 See [Packaging Related Objects into an Extension](../../server-programming/extending-sql/packaging-related-objects-into-an-extension.md#extend-extensions) for more information about these operations.


 You must own the extension to use `ALTER EXTENSION`. The `ADD`/`DROP` forms require ownership of the added/dropped object as well.


## Parameters


*name*
:   The name of an installed extension.

*new_version*
:   The desired new version of the extension. This can be written as either an identifier or a string literal. If not specified, `ALTER EXTENSION UPDATE` attempts to update to whatever is shown as the default version in the extension's control file.

*new_schema*
:   The new schema for the extension.

*object_name*, *aggregate_name*, *function_name*, *operator_name*, *procedure_name*, *routine_name*
:   The name of an object to be added to or removed from the extension. Names of tables, aggregates, domains, foreign tables, functions, operators, operator classes, operator families, procedures, routines, sequences, text search objects, types, and views can be schema-qualified.

*source_type*
:   The name of the source data type of the cast.

*target_type*
:   The name of the target data type of the cast.

*argmode*
:   The mode of a function, procedure, or aggregate argument: `IN`, `OUT`, `INOUT`, or `VARIADIC`. If omitted, the default is `IN`. Note that `ALTER EXTENSION` does not actually pay any attention to `OUT` arguments, since only the input arguments are needed to determine the function's identity. So it is sufficient to list the `IN`, `INOUT`, and `VARIADIC` arguments.

*argname*
:   The name of a function, procedure, or aggregate argument. Note that `ALTER EXTENSION` does not actually pay any attention to argument names, since only the argument data types are needed to determine the function's identity.

*argtype*
:   The data type of a function, procedure, or aggregate argument.

*left_type*, *right_type*
:   The data type(s) of the operator's arguments (optionally schema-qualified). Write `NONE` for the missing argument of a prefix operator.

`PROCEDURAL`
:   This is a noise word.

*type_name*
:   The name of the data type of the transform.

*lang_name*
:   The name of the language of the transform.


## Examples


 To update the `hstore` extension to version 2.0:

```sql

ALTER EXTENSION hstore UPDATE TO '2.0';
```


 To change the schema of the `hstore` extension to `utils`:

```sql

ALTER EXTENSION hstore SET SCHEMA utils;
```


 To add an existing function to the `hstore` extension:

```sql

ALTER EXTENSION hstore ADD FUNCTION populate_record(anyelement, hstore);
```


## Compatibility


 `ALTER EXTENSION` is a PostgreSQL extension.
 <a id="sql-alterextension-see-also"></a>

## See Also
  [sql-createextension](create-extension.md#sql-createextension), [sql-dropextension](drop-extension.md#sql-dropextension)
