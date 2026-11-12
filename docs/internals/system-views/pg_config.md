<a id="view-pg-config"></a>

## `pg_config`


 The view `pg_config` describes the compile-time configuration parameters of the currently installed version of PostgreSQL. It is intended, for example, to be used by software packages that want to interface to PostgreSQL to facilitate finding the required header files and libraries. It provides the same basic information as the [app-pgconfig](../../reference/postgresql-client-applications/pg_config.md#app-pgconfig) PostgreSQL client application.


 By default, the `pg_config` view can be read only by superusers.


**Table: `pg_config` Columns**

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
<p>The parameter name</p></td>
</tr>
<tr>
<td><p><code>setting</code> <code>text</code></p>
<p>The parameter value</p></td>
</tr>
</tbody>
</table>
