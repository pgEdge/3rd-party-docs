<a id="catalog-pg-language"></a>

## `pg_language`


 The catalog `pg_language` registers languages in which you can write functions or stored procedures. See [sql-createlanguage](../../reference/sql-commands/create-language.md#sql-createlanguage) and [Procedural Languages](../../server-programming/procedural-languages/index.md#xplang) for more information about language handlers.


**Table: `pg_language` Columns**

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
<td><p><code>lanname</code> <code>name</code></p>
<p>Name of the language</p></td>
</tr>
<tr>
<td><p><code>lanowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the language</p></td>
</tr>
<tr>
<td><p><code>lanispl</code> <code>bool</code></p>
<p>This is false for internal languages (such as SQL) and true for user-defined languages. Currently, pg_dump still uses this to determine which languages need to be dumped, but this might be replaced by a different mechanism in the future.</p></td>
</tr>
<tr>
<td><p><code>lanpltrusted</code> <code>bool</code></p>
<p>True if this is a trusted language, which means that it is believed not to grant access to anything outside the normal SQL execution environment. Only superusers can create functions in untrusted languages.</p></td>
</tr>
<tr>
<td><p><code>lanplcallfoid</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>For noninternal languages this references the language handler, which is a special function that is responsible for executing all functions that are written in the particular language. Zero for internal languages.</p></td>
</tr>
<tr>
<td><p><code>laninline</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>This references a function that is responsible for executing “inline” anonymous code blocks (<a href="../../reference/sql-commands/do.md#sql-do">sql-do</a> blocks). Zero if inline blocks are not supported.</p></td>
</tr>
<tr>
<td><p><code>lanvalidator</code> <code>oid</code> (references <a href="pg_proc.md#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>This references a language validator function that is responsible for checking the syntax and validity of new functions when they are created. Zero if no validator is provided.</p></td>
</tr>
<tr>
<td><p><code>lanacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>
