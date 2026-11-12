<a id="dml-delete"></a>

## Deleting Data


 So far we have explained how to add data to tables and how to change data. What remains is to discuss how to remove data that is no longer needed. Just as adding data is only possible in whole rows, you can only remove entire rows from a table. In the previous section we explained that SQL does not provide a way to directly address individual rows. Therefore, removing rows can only be done by specifying conditions that the rows to be removed have to match. If you have a primary key in the table then you can specify the exact row. But you can also remove groups of rows matching a condition, or you can remove all rows in the table at once.


 You use the [sql-delete](../../reference/sql-commands/delete.md#sql-delete) command to remove rows; the syntax is very similar to the [sql-update](../../reference/sql-commands/update.md#sql-update) command. For instance, to remove all rows from the products table that have a price of 10, use:

```sql

DELETE FROM products WHERE price = 10;
```


 If you simply write:

```sql

DELETE FROM products;
```
 then all rows in the table will be deleted! Caveat programmer.
