# ALTER OPERATOR { #sql-alteroperator }

change the definition of an operator

## Synopsis


```

ALTER OPERATOR NAME ( { LEFT_TYPE | NONE } , RIGHT_TYPE )
    OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }

ALTER OPERATOR NAME ( { LEFT_TYPE | NONE } , RIGHT_TYPE )
    SET SCHEMA NEW_SCHEMA

ALTER OPERATOR NAME ( { LEFT_TYPE | NONE } , RIGHT_TYPE )
    SET ( {  RESTRICT = { RES_PROC | NONE }
           | JOIN = { JOIN_PROC | NONE }
         } [, ... ] )
```


## Description


 `ALTER OPERATOR` changes the definition of an operator.


 You must own the operator to use `ALTER OPERATOR`. To alter the owner, you must be able to `SET ROLE` to the new owning role, and that role must have `CREATE` privilege on the operator's schema. (These restrictions enforce that altering the owner doesn't do anything you couldn't do by dropping and recreating the operator. However, a superuser can alter ownership of any operator anyway.)


## Parameters


*name*
:   The name (optionally schema-qualified) of an existing operator.

*left_type*
:   The data type of the operator's left operand; write `NONE` if the operator has no left operand.

*right_type*
:   The data type of the operator's right operand.

*new_owner*
:   The new owner of the operator.

*new_schema*
:   The new schema for the operator.

*res_proc*
:   The restriction selectivity estimator function for this operator; write NONE to remove existing selectivity estimator.

*join_proc*
:   The join selectivity estimator function for this operator; write NONE to remove existing selectivity estimator.


## Examples


 Change the owner of a custom operator `a @@ b` for type `text`:

```sql

ALTER OPERATOR @@ (text, text) OWNER TO joe;
```


 Change the restriction and join selectivity estimator functions of a custom operator `a && b` for type `int[]`:

```sql

ALTER OPERATOR && (_int4, _int4) SET (RESTRICT = _int_contsel, JOIN = _int_contjoinsel);
```


## Compatibility


 There is no `ALTER OPERATOR` statement in the SQL standard.


## See Also
  [sql-createoperator](create-operator.md#sql-createoperator), [sql-dropoperator](drop-operator.md#sql-dropoperator)
