<a id="view-pg-cursors"></a>

## `pg_cursors`


 The `pg_cursors` view lists the cursors that are currently available. Cursors can be defined in several ways:

-  via the [`DECLARE`](../../reference/sql-commands/declare.md#sql-declare) statement in SQL
-  via the Bind message in the frontend/backend protocol, as described in [Extended Query](../frontend-backend-protocol/message-flow.md#protocol-flow-ext-query)
-  via the Server Programming Interface (SPI), as described in [Interface Functions](../../server-programming/server-programming-interface/interface-functions.md#spi-interface)
 The `pg_cursors` view displays cursors created by any of these means. Cursors only exist for the duration of the transaction that defines them, unless they have been declared `WITH HOLD`. Therefore non-holdable cursors are only present in the view until the end of their creating transaction.

!!! note

    Cursors are used internally to implement some of the components of PostgreSQL, such as procedural languages. Therefore, the `pg_cursors` view might include cursors that have not been explicitly created by the user.


**Table: `pg_cursors` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>name</code> <code>text</code></p>
<p>The name of the cursor</p></td>
</tr>
<tr>
<td><p><code>statement</code> <code>text</code></p>
<p>The verbatim query string submitted to declare this cursor</p></td>
</tr>
<tr>
<td><p><code>is_holdable</code> <code>bool</code></p>
<p><code>true</code> if the cursor is holdable (that is, it can be accessed after the transaction that declared the cursor has committed); <code>false</code> otherwise</p></td>
</tr>
<tr>
<td><p><code>is_binary</code> <code>bool</code></p>
<p><code>true</code> if the cursor was declared <code>BINARY</code>; <code>false</code> otherwise</p></td>
</tr>
<tr>
<td><p><code>is_scrollable</code> <code>bool</code></p>
<p><code>true</code> if the cursor is scrollable (that is, it allows rows to be retrieved in a nonsequential manner); <code>false</code> otherwise</p></td>
</tr>
<tr>
<td><p><code>creation_time</code> <code>timestamptz</code></p>
<p>The time at which the cursor was declared</p></td>
</tr>
</tbody>
</table>


 The `pg_cursors` view is read-only.
