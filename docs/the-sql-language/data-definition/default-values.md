<a id="ddl-default"></a>

## Default Values


 A column can be assigned a default value. When a new row is created and no values are specified for some of the columns, those columns will be filled with their respective default values. A data manipulation command can also request explicitly that a column be set to its default value, without having to know what that value is. (Details about data manipulation commands are in [Data Manipulation](../data-manipulation/index.md#dml).)


  If no default value is declared explicitly, the default value is the null value. This usually makes sense because a null value can be considered to represent unknown data.


 In a table definition, default values are listed after the column data type. For example:

```sql

CREATE TABLE products (
    product_no integer,
    name text,
    price numeric DEFAULT 9.99
);
```


 The default value can be an expression, which will be evaluated whenever the default value is inserted (*not* when the table is created). A common example is for a `timestamp` column to have a default of `CURRENT_TIMESTAMP`, so that it gets set to the time of row insertion. Another common example is generating a “serial number” for each row. In PostgreSQL this is typically done by something like:

```sql

CREATE TABLE products (
    product_no integer DEFAULT nextval('products_product_no_seq'),
    ...
);
```
 where the `nextval()` function supplies successive values from a *sequence object* (see [Sequence Manipulation Functions](../functions-and-operators/sequence-manipulation-functions.md#functions-sequence)). This arrangement is sufficiently common that there's a special shorthand for it:

```sql

CREATE TABLE products (
    product_no SERIAL,
    ...
);
```
 The `SERIAL` shorthand is discussed further in [Serial Types](../data-types/numeric-types.md#datatype-serial).
