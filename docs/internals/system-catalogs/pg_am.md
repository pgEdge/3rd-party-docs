<a id="catalog-pg-am"></a>

## `pg_am`


 The catalog `pg_am` stores information about relation access methods. There is one row for each access method supported by the system. Currently, only tables and indexes have access methods. The requirements for table and index access methods are discussed in detail in [Table Access Method Interface Definition](../table-access-method-interface-definition.md#tableam) and [Index Access Method Interface Definition](../index-access-method-interface-definition/index.md#indexam) respectively.


**Table: `pg_am` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>oid</code> <code>oid</code></p>
<p>Row identifier</p></td>
</tr>
<tr>
<td><p><code>amname</code> <code>name</code></p>
<p>Name of the access method</p></td>
</tr>
<tr>
<td><p><code>amhandler</code> <code>regproc</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>OID of a handler function that is responsible for supplying information about the access method</p></td>
</tr>
<tr>
<td><p><code>amtype</code> <code>char</code></p>
<p><code>t</code> = table (including materialized views), <code>i</code> = index.</p></td>
</tr>
</tbody>
</table>


!!! note

    Before PostgreSQL 9.6, `pg_am` contained many additional columns representing properties of index access methods. That data is now only directly visible at the C code level. However, `pg_index_column_has_property()` and related functions have been added to allow SQL queries to inspect index access method properties; see [System Catalog Information Functions](../../the-sql-language/functions-and-operators/system-information-functions-and-operators.md#functions-info-catalog-table).
