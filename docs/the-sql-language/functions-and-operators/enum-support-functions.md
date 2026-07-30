<a id="functions-enum"></a>

## Enum Support Functions


 For enum types (described in [Enumerated Types](../data-types/enumerated-types.md#datatype-enum)), there are several functions that allow cleaner programming without hard-coding particular values of an enum type. These are listed in [Enum Support Functions](#functions-enum-table). The examples assume an enum type created as:

```sql

CREATE TYPE rainbow AS ENUM ('red', 'orange', 'yellow', 'green', 'blue', 'purple');
```

 <a id="functions-enum-table"></a>

**Table: Enum Support Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>enum_first</code> ( <code>anyenum</code> ) <code>anyenum</code></td>
<td>Returns the first value of the input enum type.</td>
<td><code>enum_first(null::rainbow)</code> <code>red</code></td>
</tr>
<tr>
<td><code>enum_last</code> ( <code>anyenum</code> ) <code>anyenum</code></td>
<td>Returns the last value of the input enum type.</td>
<td><code>enum_last(null::rainbow)</code> <code>purple</code></td>
</tr>
<tr>
<td><code>enum_range</code> ( <code>anyenum</code> ) <code>anyarray</code></td>
<td>Returns all values of the input enum type in an ordered array.</td>
<td><code>enum_range(null::rainbow)</code> <code>{red,orange,yellow,​green,blue,purple}</code></td>
</tr>
<tr>
<td><code>enum_range</code> ( <code>anyenum</code>, <code>anyenum</code> ) <code>anyarray</code></td>
<td>Returns the range between the two given enum values, as an ordered array. The values must be from the same enum type. If the first parameter is null, the result will start with the first value of the enum type. If the second parameter is null, the result will end with the last value of the enum type.</td>
<td><code>enum_range('orange'::rainbow, 'green'::rainbow)</code> <code>{orange,yellow,green}</code><br><code>enum_range(NULL, 'green'::rainbow)</code> <code>{red,orange,​yellow,green}</code><br><code>enum_range('orange'::rainbow, NULL)</code> <code>{orange,yellow,green,​blue,purple}</code></td>
</tr>
</tbody>
</table>


 Notice that except for the two-argument form of `enum_range`, these functions disregard the specific value passed to them; they care only about its declared data type. Either null or a specific value of the type can be passed, with the same result. It is more common to apply these functions to a table column or function argument than to a hardwired type name as used in the examples.
