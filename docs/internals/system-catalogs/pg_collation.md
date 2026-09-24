<a id="catalog-pg-collation"></a>

## `pg_collation`


 The catalog `pg_collation` describes the available collations, which are essentially mappings from an SQL name to operating system locale categories. See [Collation Support](../../server-administration/localization/collation-support.md#collation) for more information.


**Table: `pg_collation` Columns**

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
<td><p><code>collname</code> <code>name</code></p>
<p>Collation name (unique per namespace and encoding)</p></td>
</tr>
<tr>
<td><p><code>collnamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this collation</p></td>
</tr>
<tr>
<td><p><code>collowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the collation</p></td>
</tr>
<tr>
<td><p><code>collprovider</code> <code>char</code></p>
<p>Provider of the collation: <code>d</code> = database default, <code>b</code> = builtin, <code>c</code> = libc, <code>i</code> = icu</p></td>
</tr>
<tr>
<td><p><code>collisdeterministic</code> <code>bool</code></p>
<p>Is the collation deterministic?</p></td>
</tr>
<tr>
<td><p><code>collencoding</code> <code>int4</code></p>
<p>Encoding in which the collation is applicable, or -1 if it works for any encoding</p></td>
</tr>
<tr>
<td><p><code>collcollate</code> <code>text</code></p>
<p><code>LC_COLLATE</code> for this collation object. If the provider is not <code>libc</code>, <code>collcollate</code> is <code>NULL</code> and <code>colllocale</code> is used instead.</p></td>
</tr>
<tr>
<td><p><code>collctype</code> <code>text</code></p>
<p><code>LC_CTYPE</code> for this collation object. If the provider is not <code>libc</code>, <code>collctype</code> is <code>NULL</code> and <code>colllocale</code> is used instead.</p></td>
</tr>
<tr>
<td><p><code>colllocale</code> <code>text</code></p>
<p>Collation provider locale name for this collation object. If the provider is <code>libc</code>, <code>colllocale</code> is <code>NULL</code>; <code>collcollate</code> and <code>collctype</code> are used instead.</p></td>
</tr>
<tr>
<td><p><code>collicurules</code> <code>text</code></p>
<p>ICU collation rules for this collation object</p></td>
</tr>
<tr>
<td><p><code>collversion</code> <code>text</code></p>
<p>Provider-specific version of the collation. This is recorded when the collation is created and then checked when it is used, to detect changes in the collation definition that could lead to data corruption.</p></td>
</tr>
</tbody>
</table>


 Note that the unique key on this catalog is (`collname`, `collencoding`, `collnamespace`) not just (`collname`, `collnamespace`). PostgreSQL generally ignores all collations that do not have `collencoding` equal to either the current database's encoding or -1, and creation of new entries with the same name as an entry with `collencoding` = -1 is forbidden. Therefore it is sufficient to use a qualified SQL name (*schema*.*name*) to identify a collation, even though this is not unique according to the catalog definition. The reason for defining the catalog this way is that initdb fills it in at cluster initialization time with entries for all locales available on the system, so it must be able to hold entries for all encodings that might ever be used in the cluster.


 In the `template0` database, it could be useful to create collations whose encoding does not match the database encoding, since they could match the encodings of databases later cloned from `template0`. This would currently have to be done manually.
