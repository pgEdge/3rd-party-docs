<a id="functions-tid"></a>

## TID Functions


 [TID Functions](#functions-tid-table) lists functions for the `tid` data type (tuple identifier).
 <a id="functions-tid-table"></a>

**Table: TID Functions**

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
<td><code>tid_block</code> ( <code>tid</code> ) <code>bigint</code></td>
<td>Extracts the block number from a tuple identifier.</td>
<td><code>tid_block('(42,7)'::tid)</code> <code>42</code></td>
</tr>
<tr>
<td><code>tid_offset</code> ( <code>tid</code> ) <code>integer</code></td>
<td>Extracts the tuple offset within the block from a tuple identifier.</td>
<td><code>tid_offset('(42,7)'::tid)</code> <code>7</code></td>
</tr>
</tbody>
</table>
