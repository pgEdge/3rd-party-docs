<a id="functions-trigger"></a>

## Trigger Functions


 While many uses of triggers involve user-written trigger functions, PostgreSQL provides a few built-in trigger functions that can be used directly in user-defined triggers. These are summarized in [Built-In Trigger Functions](#builtin-triggers-table). (Additional built-in trigger functions exist, which implement foreign key constraints and deferred index constraints. Those are not documented here since users need not use them directly.)


 For more information about creating triggers, see [sql-createtrigger](../../reference/sql-commands/create-trigger.md#sql-createtrigger).
 <a id="builtin-triggers-table"></a>

**Table: Built-In Trigger Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example Usage</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>suppress_redundant_updates_trigger</code> ( ) <code>trigger</code></td>
<td>Suppresses do-nothing update operations. See below for details.</td>
<td><code>CREATE TRIGGER ... suppress_redundant_updates_trigger()</code></td>
</tr>
<tr>
<td><code>tsvector_update_trigger</code> ( ) <code>trigger</code></td>
<td>Automatically updates a <code>tsvector</code> column from associated plain-text document column(s). The text search configuration to use is specified by name as a trigger argument. See <a href="../full-text-search/additional-features.md#textsearch-update-triggers">Triggers for Automatic Updates</a> for details.</td>
<td><code>CREATE TRIGGER ... tsvector_update_trigger(tsvcol, 'pg_catalog.swedish', title, body)</code></td>
</tr>
<tr>
<td><code>tsvector_update_trigger_column</code> ( ) <code>trigger</code></td>
<td>Automatically updates a <code>tsvector</code> column from associated plain-text document column(s). The text search configuration to use is taken from a <code>regconfig</code> column of the table. See <a href="../full-text-search/additional-features.md#textsearch-update-triggers">Triggers for Automatic Updates</a> for details.</td>
<td><code>CREATE TRIGGER ... tsvector_update_trigger_column(tsvcol, tsconfigcol, title, body)</code></td>
</tr>
</tbody>
</table>


 The `suppress_redundant_updates_trigger` function, when applied as a row-level `BEFORE UPDATE` trigger, will prevent any update that does not actually change the data in the row from taking place. This overrides the normal behavior which always performs a physical row update regardless of whether or not the data has changed. (This normal behavior makes updates run faster, since no checking is required, and is also useful in certain cases.)


 Ideally, you should avoid running updates that don't actually change the data in the record. Redundant updates can cost considerable unnecessary time, especially if there are lots of indexes to alter, and space in dead rows that will eventually have to be vacuumed. However, detecting such situations in client code is not always easy, or even possible, and writing expressions to detect them can be error-prone. An alternative is to use `suppress_redundant_updates_trigger`, which will skip updates that don't change the data. You should use this with care, however. The trigger takes a small but non-trivial time for each record, so if most of the records affected by updates do actually change, use of this trigger will make updates run slower on average.


 The `suppress_redundant_updates_trigger` function can be added to a table like this:

```sql

CREATE TRIGGER z_min_update
BEFORE UPDATE ON tablename
FOR EACH ROW EXECUTE FUNCTION suppress_redundant_updates_trigger();
```
 In most cases, you need to fire this trigger last for each row, so that it does not override other triggers that might wish to alter the row. Bearing in mind that triggers fire in name order, you would therefore choose a trigger name that comes after the name of any other trigger you might have on the table. (Hence the “z” prefix in the example.)
