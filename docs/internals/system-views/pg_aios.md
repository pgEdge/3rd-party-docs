<a id="view-pg-aios"></a>

## `pg_aios`


 The `pg_aios` view lists all [glossary-aio](../../appendixes/glossary.md#glossary-aio) handles that are currently in-use. An I/O handle is used to reference an I/O operation that is being prepared, executed or is in the process of completing. `pg_aios` contains one row for each I/O handle.


 This view is mainly useful for developers of PostgreSQL, but may also be useful when tuning PostgreSQL.


**Table: `pg_aios` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>pid</code> <code>int4</code></p>
<p>Process ID of the server process that is issuing this I/O.</p></td>
</tr>
<tr>
<td><p><code>io_id</code> <code>int4</code></p>
<p>Identifier of the I/O handle. Handles are reused once the I/O completed (or if the handle is released before I/O is started). On reuse <a href="#view-pg-aios-io-generation"><code>pg_aios</code>.<code>io_generation</code></a> is incremented.</p></td>
</tr>
<tr>
<td id="view-pg-aios-io-generation"><p><code>io_generation</code> <code>int8</code></p>
<p>Generation of the I/O handle.</p></td>
</tr>
<tr>
<td><p><code>state</code> <code>text</code></p>
<p>State of the I/O handle:</p>
<p>-  <code>HANDED_OUT</code>, referenced by code but not yet used <br>
-  <code>DEFINED</code>, information necessary for execution is known <br>
-  <code>STAGED</code>, ready for execution <br>
-  <code>SUBMITTED</code>, submitted for execution <br>
-  <code>COMPLETED_IO</code>, finished, but result has not yet been processed <br>
-  <code>COMPLETED_SHARED</code>, shared completion processing completed <br>
-  <code>COMPLETED_LOCAL</code>, backend local completion processing completed</p></td>
</tr>
<tr>
<td><p><code>operation</code> <code>text</code></p>
<p>Operation performed using the I/O handle:</p>
<p>-  <code>invalid</code>, not yet known <br>
-  <code>readv</code>, a vectored read <br>
-  <code>writev</code>, a vectored write</p></td>
</tr>
<tr>
<td><p><code>off</code> <code>int8</code></p>
<p>Offset of the I/O operation.</p></td>
</tr>
<tr>
<td><p><code>length</code> <code>int8</code></p>
<p>Length of the I/O operation.</p></td>
</tr>
<tr>
<td><p><code>target</code> <code>text</code></p>
<p>What kind of object is the I/O targeting:</p>
<p>-  <code>smgr</code>, I/O on relations</p></td>
</tr>
<tr>
<td><p><code>handle_data_len</code> <code>int2</code></p>
<p>Length of the data associated with the I/O operation. For I/O to/from <a href="../../server-administration/server-configuration/resource-consumption.md#guc-shared-buffers">shared_buffers</a> and <a href="../../server-administration/server-configuration/resource-consumption.md#guc-temp-buffers">temp_buffers</a>, this indicates the number of buffers the I/O is operating on.</p></td>
</tr>
<tr>
<td><p><code>raw_result</code> <code>int4</code></p>
<p>Low-level result of the I/O operation, or NULL if the operation has not yet completed.</p></td>
</tr>
<tr>
<td><p><code>result</code> <code>text</code></p>
<p>High-level result of the I/O operation:</p>
<p>-  <code>UNKNOWN</code> means that the result of the operation is not yet known. <br>
-  <code>OK</code> means the I/O completed successfully. <br>
-  <code>PARTIAL</code> means that the I/O completed without error, but did not process all data. Commonly callers will need to retry and perform the remainder of the work in a separate I/O. <br>
-  <code>WARNING</code> means that the I/O completed without error, but that execution of the IO triggered a warning. E.g. when encountering a corrupted buffer with <a href="../../server-administration/server-configuration/developer-options.md#guc-zero-damaged-pages">zero_damaged_pages</a> enabled. <br>
-  <code>ERROR</code> means the I/O failed with an error.</p></td>
</tr>
<tr>
<td><p><code>target_desc</code> <code>text</code></p>
<p>Description of what the I/O operation is targeting.</p></td>
</tr>
<tr>
<td><p><code>f_sync</code> <code>bool</code></p>
<p>Flag indicating whether the I/O is executed synchronously.</p></td>
</tr>
<tr>
<td><p><code>f_localmem</code> <code>bool</code></p>
<p>Flag indicating whether the I/O references process local memory.</p></td>
</tr>
<tr>
<td><p><code>f_buffered</code> <code>bool</code></p>
<p>Flag indicating whether the I/O is buffered I/O.</p></td>
</tr>
</tbody>
</table>


 The `pg_aios` view is read-only.


 By default, the `pg_aios` view can be read only by superusers or roles with privileges of the `pg_read_all_stats` role.
