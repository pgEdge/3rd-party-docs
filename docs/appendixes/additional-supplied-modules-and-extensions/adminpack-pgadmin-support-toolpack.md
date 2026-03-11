## adminpack — pgAdmin support toolpack { #adminpack }


 `adminpack` provides a number of support functions which pgAdmin and other administration and management tools can use to provide additional functionality, such as remote management of server log files. Use of all these functions is only allowed to database superusers by default, but may be allowed to other users by using the `GRANT` command.


 The functions shown in [`adminpack` Functions](#functions-adminpack-table) provide write access to files on the machine hosting the server. (See also the functions in [Generic File Access Functions](../../the-sql-language/functions-and-operators/system-administration-functions.md#functions-admin-genfile-table), which provide read-only access.) Only files within the database cluster directory can be accessed, unless the user is a superuser or given privileges of one of the `pg_read_server_files` or `pg_write_server_files` roles, as appropriate for the function, but either a relative or absolute path is allowable.


<a id="functions-adminpack-table"></a>
**Table: `adminpack` Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_catalog.pg_file_write</code> ( <code>filename</code> <code>text</code>, <code>data</code> <code>text</code>, <code>append</code> <code>boolean</code> ) <code>bigint</code></td>
<td>Writes, or appends to, a text file.</td>
<td></td>
</tr>
<tr>
<td><code>pg_catalog.pg_file_sync</code> ( <code>filename</code> <code>text</code> ) <code>void</code></td>
<td>Flushes a file or directory to disk.</td>
<td></td>
</tr>
<tr>
<td><code>pg_catalog.pg_file_rename</code> ( <code>oldname</code> <code>text</code>, <code>newname</code> <code>text</code> [, <code>archivename</code> <code>text</code> ] ) <code>boolean</code></td>
<td>Renames a file.</td>
<td></td>
</tr>
<tr>
<td><code>pg_catalog.pg_file_unlink</code> ( <code>filename</code> <code>text</code> ) <code>boolean</code></td>
<td>Removes a file.</td>
<td></td>
</tr>
<tr>
<td><code>pg_catalog.pg_logdir_ls</code> () <code>setof record</code></td>
<td>Lists the log files in the <code>log_directory</code> directory.</td>
<td></td>
</tr>
</tbody>
</table>


 `pg_file_write` writes the specified `data` into the file named by `filename`. If `append` is false, the file must not already exist. If `append` is true, the file can already exist, and will be appended to if so. Returns the number of bytes written.


 `pg_file_sync` fsyncs the specified file or directory named by `filename`. An error is thrown on failure (e.g., the specified file is not present). Note that [data_sync_retry](../../server-administration/server-configuration/error-handling.md#guc-data-sync-retry) has no effect on this function, and therefore a PANIC-level error will not be raised even on failure to flush database files.


 `pg_file_rename` renames a file. If `archivename` is omitted or NULL, it simply renames `oldname` to `newname` (which must not already exist). If `archivename` is provided, it first renames `newname` to `archivename` (which must not already exist), and then renames `oldname` to `newname`. In event of failure of the second rename step, it will try to rename `archivename` back to `newname` before reporting the error. Returns true on success, false if the source file(s) are not present or not writable; other cases throw errors.


 `pg_file_unlink` removes the specified file. Returns true on success, false if the specified file is not present or the `unlink()` call fails; other cases throw errors.


 `pg_logdir_ls` returns the start timestamps and path names of all the log files in the [log_directory](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-directory) directory. The [log_filename](../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-filename) parameter must have its default setting (`postgresql-%Y-%m-%d_%H%M%S.log`) to use this function.
