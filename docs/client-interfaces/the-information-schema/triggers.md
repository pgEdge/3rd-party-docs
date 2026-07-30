<a id="infoschema-triggers"></a>

## `triggers`


 The view `triggers` contains all triggers defined in the current database on tables and views that the current user owns or has some privilege other than `SELECT` on.


**Table: `triggers` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>trigger_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the trigger (always the current database)</p></td>
</tr>
<tr>
<td><p><code>trigger_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the trigger</p></td>
</tr>
<tr>
<td><p><code>trigger_name</code> <code>sql_identifier</code></p>
<p>Name of the trigger</p></td>
</tr>
<tr>
<td><p><code>event_manipulation</code> <code>character_data</code></p>
<p>Event that fires the trigger (<code>INSERT</code>, <code>UPDATE</code>, or <code>DELETE</code>)</p></td>
</tr>
<tr>
<td><p><code>event_object_catalog</code> <code>sql_identifier</code></p>
<p>Name of the database that contains the table that the trigger is defined on (always the current database)</p></td>
</tr>
<tr>
<td><p><code>event_object_schema</code> <code>sql_identifier</code></p>
<p>Name of the schema that contains the table that the trigger is defined on</p></td>
</tr>
<tr>
<td><p><code>event_object_table</code> <code>sql_identifier</code></p>
<p>Name of the table that the trigger is defined on</p></td>
</tr>
<tr>
<td><p><code>action_order</code> <code>cardinal_number</code></p>
<p>Firing order among triggers on the same table having the same <code>event_manipulation</code>, <code>action_timing</code>, and <code>action_orientation</code>. In PostgreSQL, triggers are fired in name order, so this column reflects that.</p></td>
</tr>
<tr>
<td><p><code>action_condition</code> <code>character_data</code></p>
<p><code>WHEN</code> condition of the trigger, null if none (also null if the table is not owned by a currently enabled role)</p></td>
</tr>
<tr>
<td><p><code>action_statement</code> <code>character_data</code></p>
<p>Statement that is executed by the trigger (currently always <code>EXECUTE FUNCTION<br>
       </code><em>function</em><code>(...)</code>)</p></td>
</tr>
<tr>
<td><p><code>action_orientation</code> <code>character_data</code></p>
<p>Identifies whether the trigger fires once for each processed row or once for each statement (<code>ROW</code> or <code>STATEMENT</code>)</p></td>
</tr>
<tr>
<td><p><code>action_timing</code> <code>character_data</code></p>
<p>Time at which the trigger fires (<code>BEFORE</code>, <code>AFTER</code>, or <code>INSTEAD OF</code>)</p></td>
</tr>
<tr>
<td><p><code>action_reference_old_table</code> <code>sql_identifier</code></p>
<p>Name of the “old” transition table, or null if none</p></td>
</tr>
<tr>
<td><p><code>action_reference_new_table</code> <code>sql_identifier</code></p>
<p>Name of the “new” transition table, or null if none</p></td>
</tr>
<tr>
<td><p><code>action_reference_old_row</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>action_reference_new_row</code> <code>sql_identifier</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
<tr>
<td><p><code>created</code> <code>time_stamp</code></p>
<p>Applies to a feature not available in PostgreSQL</p></td>
</tr>
</tbody>
</table>


 Triggers in PostgreSQL have two incompatibilities with the SQL standard that affect the representation in the information schema. First, trigger names are local to each table in PostgreSQL, rather than being independent schema objects. Therefore there can be duplicate trigger names defined in one schema, so long as they belong to different tables. (`trigger_catalog` and `trigger_schema` are really the values pertaining to the table that the trigger is defined on.) Second, triggers can be defined to fire on multiple events in PostgreSQL (e.g., `ON INSERT OR UPDATE`), whereas the SQL standard only allows one. If a trigger is defined to fire on multiple events, it is represented as multiple rows in the information schema, one for each type of event. As a consequence of these two issues, the primary key of the view `triggers` is really `(trigger_catalog, trigger_schema, event_object_table, trigger_name, event_manipulation)` instead of `(trigger_catalog, trigger_schema, trigger_name)`, which is what the SQL standard specifies. Nonetheless, if you define your triggers in a manner that conforms with the SQL standard (trigger names unique in the schema and only one event type per trigger), this will not affect you.


!!! note

    Prior to PostgreSQL 9.1, this view's columns `action_timing`, `action_reference_old_table`, `action_reference_new_table`, `action_reference_old_row`, and `action_reference_new_row` were named `condition_timing`, `condition_reference_old_table`, `condition_reference_new_table`, `condition_reference_old_row`, and `condition_reference_new_row` respectively. That was how they were named in the SQL:1999 standard. The new naming conforms to SQL:2003 and later.
