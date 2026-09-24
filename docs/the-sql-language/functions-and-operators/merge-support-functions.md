<a id="functions-merge-support"></a>

## Merge Support Functions


 PostgreSQL includes one merge support function that may be used in the `RETURNING` list of a [sql-merge](../../reference/sql-commands/merge.md#sql-merge) command to identify the action taken for each row; see [Merge Support Functions](#functions-merge-support-table).
 <a id="functions-merge-support-table"></a>

**Table: Merge Support Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr id="merge-action">
<td><code>merge_action</code> ( ) <code>text</code></td>
<td>Returns the merge action command executed for the current row. This will be <code>'INSERT'</code>, <code>'UPDATE'</code>, or <code>'DELETE'</code>.</td>
<td></td>
</tr>
</tbody>
</table>


 Example:

```

MERGE INTO products p
  USING stock s ON p.product_id = s.product_id
  WHEN MATCHED AND s.quantity > 0 THEN
    UPDATE SET in_stock = true, quantity = s.quantity
  WHEN MATCHED THEN
    UPDATE SET in_stock = false, quantity = 0
  WHEN NOT MATCHED THEN
    INSERT (product_id, in_stock, quantity)
      VALUES (s.product_id, true, s.quantity)
  RETURNING merge_action(), p.*;

 merge_action | product_id | in_stock | quantity
--------------+------------+----------+----------
 UPDATE       |       1001 | t        |       50
 UPDATE       |       1002 | f        |        0
 INSERT       |       1003 | t        |       10
```


 Note that this function can only be used in the `RETURNING` list of a `MERGE` command. It is an error to use it in any other part of a query.
