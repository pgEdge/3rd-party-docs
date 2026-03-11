## System Information Functions and Operators { #functions-info }


 The functions described in this section are used to obtain various information about a PostgreSQL installation.


### Session Information Functions { #functions-info-session }


 [Session Information Functions](#functions-info-session-table) shows several functions that extract session and system information.


 In addition to the functions listed in this section, there are a number of functions related to the statistics system that also provide system information. See [Statistics Functions](../../server-administration/monitoring-database-activity/the-cumulative-statistics-system.md#monitoring-stats-functions) for more information.


<a id="functions-info-session-table"></a>
**Table: Session Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>current_catalog</code> <code>name</code></td>
<td><code>current_database</code> () <code>name</code></td>
<td>Returns the name of the current database. (Databases are called “catalogs” in the SQL standard, so <code>current_catalog</code> is the standard's spelling.)</td>
</tr>
<tr>
<td><code>current_query</code> () <code>text</code></td>
<td>Returns the text of the currently executing query, as submitted by the client (which might contain more than one statement).</td>
<td></td>
</tr>
<tr>
<td><code>current_role</code> <code>name</code></td>
<td>This is equivalent to <code>current_user</code>.</td>
<td></td>
</tr>
<tr>
<td><code>current_schema</code> <code>name</code></td>
<td><code>current_schema</code> () <code>name</code></td>
<td>Returns the name of the schema that is first in the search path (or a null value if the search path is empty). This is the schema that will be used for any tables or other named objects that are created without specifying a target schema.</td>
</tr>
<tr>
<td><code>current_schemas</code> ( <code>include_implicit</code> <code>boolean</code> ) <code>name[]</code></td>
<td>Returns an array of the names of all schemas presently in the effective search path, in their priority order. (Items in the current <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-search-path">search_path</a> setting that do not correspond to existing, searchable schemas are omitted.) If the Boolean argument is <code>true</code>, then implicitly-searched system schemas such as <code>pg_catalog</code> are included in the result.</td>
<td></td>
</tr>
<tr>
<td><code>current_user</code> <code>name</code></td>
<td>Returns the user name of the current execution context.</td>
<td></td>
</tr>
<tr>
<td><code>inet_client_addr</code> () <code>inet</code></td>
<td>Returns the IP address of the current client, or <code>NULL</code> if the current connection is via a Unix-domain socket.</td>
<td></td>
</tr>
<tr>
<td><code>inet_client_port</code> () <code>integer</code></td>
<td>Returns the IP port number of the current client, or <code>NULL</code> if the current connection is via a Unix-domain socket.</td>
<td></td>
</tr>
<tr>
<td><code>inet_server_addr</code> () <code>inet</code></td>
<td>Returns the IP address on which the server accepted the current connection, or <code>NULL</code> if the current connection is via a Unix-domain socket.</td>
<td></td>
</tr>
<tr>
<td><code>inet_server_port</code> () <code>integer</code></td>
<td>Returns the IP port number on which the server accepted the current connection, or <code>NULL</code> if the current connection is via a Unix-domain socket.</td>
<td></td>
</tr>
<tr>
<td><code>pg_backend_pid</code> () <code>integer</code></td>
<td>Returns the process ID of the server process attached to the current session.</td>
<td></td>
</tr>
<tr>
<td><code>pg_blocking_pids</code> ( <code>integer</code> ) <code>integer[]</code></td>
<td>Returns an array of the process ID(s) of the sessions that are blocking the server process with the specified process ID from acquiring a lock, or an empty array if there is no such server process or it is not blocked.</td>
<td>One server process blocks another if it either holds a lock that conflicts with the blocked process's lock request (hard block), or is waiting for a lock that would conflict with the blocked process's lock request and is ahead of it in the wait queue (soft block). When using parallel queries the result always lists client-visible process IDs (that is, <code>pg_backend_pid</code> results) even if the actual lock is held or awaited by a child worker process. As a result of that, there may be duplicated PIDs in the result. Also note that when a prepared transaction holds a conflicting lock, it will be represented by a zero process ID.<br>Frequent calls to this function could have some impact on database performance, because it needs exclusive access to the lock manager's shared state for a short time.</td>
</tr>
<tr>
<td><code>pg_conf_load_time</code> () <code>timestamp with time zone</code></td>
<td>Returns the time when the server configuration files were last loaded. If the current session was alive at the time, this will be the time when the session itself re-read the configuration files (so the reading will vary a little in different sessions). Otherwise it is the time when the postmaster process re-read the configuration files.</td>
<td></td>
</tr>
<tr>
<td><code>pg_current_logfile</code> ( [ <code>text</code> ] ) <code>text</code></td>
<td>Returns the path name of the log file currently in use by the logging collector. The path includes the <a href="../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-directory">log_directory</a> directory and the individual log file name. The result is <code>NULL</code> if the logging collector is disabled. When multiple log files exist, each in a different format, <code>pg_current_logfile</code> without an argument returns the path of the file having the first format found in the ordered list: <code>stderr</code>, <code>csvlog</code>, <code>jsonlog</code>. <code>NULL</code> is returned if no log file has any of these formats. To request information about a specific log file format, supply either <code>csvlog</code>, <code>jsonlog</code> or <code>stderr</code> as the value of the optional parameter. The result is <code>NULL</code> if the log format requested is not configured in <a href="../../server-administration/server-configuration/error-reporting-and-logging.md#guc-log-destination">log_destination</a>. The result reflects the contents of the <code>current_logfiles</code> file.</td>
<td>This function is restricted to superusers and roles with privileges of the <code>pg_monitor</code> role by default, but other users can be granted EXECUTE to run the function.</td>
</tr>
<tr>
<td><code>pg_get_loaded_modules</code> () <code>setof record</code> ( <code>module_name</code> <code>text</code>, <code>version</code> <code>text</code>, <code>file_name</code> <code>text</code> )</td>
<td>Returns a list of the loadable modules that are loaded into the current server session. The <code>module_name</code> and <code>version</code> fields are NULL unless the module author supplied values for them using the <code>PG_MODULE_MAGIC_EXT</code> macro. The <code>file_name</code> field gives the file name of the module (shared library).</td>
<td></td>
</tr>
<tr>
<td><code>pg_my_temp_schema</code> () <code>oid</code></td>
<td>Returns the OID of the current session's temporary schema, or zero if it has none (because it has not created any temporary tables).</td>
<td></td>
</tr>
<tr>
<td><code>pg_is_other_temp_schema</code> ( <code>oid</code> ) <code>boolean</code></td>
<td>Returns true if the given OID is the OID of another session's temporary schema. (This can be useful, for example, to exclude other sessions' temporary tables from a catalog display.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_jit_available</code> () <code>boolean</code></td>
<td>Returns true if a JIT compiler extension is available (see <a href="../../server-administration/just-in-time-compilation-jit/index.md#jit">Just-in-Time Compilation (JIT)</a>) and the <a href="../../server-administration/server-configuration/query-planning.md#guc-jit">jit</a> configuration parameter is set to <code>on</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_numa_available</code> () <code>boolean</code></td>
<td>Returns true if the server has been compiled with NUMA support.</td>
<td></td>
</tr>
<tr>
<td><code>pg_listening_channels</code> () <code>setof text</code></td>
<td>Returns the set of names of asynchronous notification channels that the current session is listening to.</td>
<td></td>
</tr>
<tr>
<td><code>pg_notification_queue_usage</code> () <code>double precision</code></td>
<td>Returns the fraction (0–1) of the asynchronous notification queue's maximum size that is currently occupied by notifications that are waiting to be processed. See <a href="../../reference/sql-commands/listen.md#sql-listen">sql-listen</a> and <a href="../../reference/sql-commands/notify.md#sql-notify">sql-notify</a> for more information.</td>
<td></td>
</tr>
<tr>
<td><code>pg_postmaster_start_time</code> () <code>timestamp with time zone</code></td>
<td>Returns the time when the server started.</td>
<td></td>
</tr>
<tr>
<td><code>pg_safe_snapshot_blocking_pids</code> ( <code>integer</code> ) <code>integer[]</code></td>
<td>Returns an array of the process ID(s) of the sessions that are blocking the server process with the specified process ID from acquiring a safe snapshot, or an empty array if there is no such server process or it is not blocked.</td>
<td>A session running a <code>SERIALIZABLE</code> transaction blocks a <code>SERIALIZABLE READ ONLY DEFERRABLE</code> transaction from acquiring a snapshot until the latter determines that it is safe to avoid taking any predicate locks. See <a href="../concurrency-control/transaction-isolation.md#xact-serializable">Serializable Isolation Level</a> for more information about serializable and deferrable transactions.<br>Frequent calls to this function could have some impact on database performance, because it needs access to the predicate lock manager's shared state for a short time.</td>
</tr>
<tr>
<td><code>pg_trigger_depth</code> () <code>integer</code></td>
<td>Returns the current nesting level of PostgreSQL triggers (0 if not called, directly or indirectly, from inside a trigger).</td>
<td></td>
</tr>
<tr>
<td><code>session_user</code> <code>name</code></td>
<td>Returns the session user's name.</td>
<td></td>
</tr>
<tr>
<td><code>system_user</code> <code>text</code></td>
<td>Returns the authentication method and the identity (if any) that the user presented during the authentication cycle before they were assigned a database role. It is represented as <code>auth_method:identity</code> or <code>NULL</code> if the user has not been authenticated (for example if <a href="../../server-administration/client-authentication/trust-authentication.md#auth-trust">Trust authentication</a> has been used).</td>
<td></td>
</tr>
<tr>
<td><code>user</code> <code>name</code></td>
<td>This is equivalent to <code>current_user</code>.</td>
<td></td>
</tr>
</tbody>
</table>


!!! note

    `current_catalog`, `current_role`, `current_schema`, `current_user`, `session_user`, and `user` have special syntactic status in SQL: they must be called without trailing parentheses. In PostgreSQL, parentheses can optionally be used with `current_schema`, but not with the others.


 The `session_user` is normally the user who initiated the current database connection; but superusers can change this setting with [sql-set-session-authorization](../../reference/sql-commands/set-session-authorization.md#sql-set-session-authorization). The `current_user` is the user identifier that is applicable for permission checking. Normally it is equal to the session user, but it can be changed with [sql-set-role](../../reference/sql-commands/set-role.md#sql-set-role). It also changes during the execution of functions with the attribute `SECURITY DEFINER`. In Unix parlance, the session user is the “real user” and the current user is the “effective user”. `current_role` and `user` are synonyms for `current_user`. (The SQL standard draws a distinction between `current_role` and `current_user`, but PostgreSQL does not, since it unifies users and roles into a single kind of entity.)


### Access Privilege Inquiry Functions { #functions-info-access }


 [Access Privilege Inquiry Functions](#functions-info-access-table) lists functions that allow querying object access privileges programmatically. (See [Privileges](../data-definition/privileges.md#ddl-priv) for more information about privileges.) In these functions, the user whose privileges are being inquired about can be specified by name or by OID (`pg_authid`.`oid`), or if the name is given as `public` then the privileges of the PUBLIC pseudo-role are checked. Also, the `user` argument can be omitted entirely, in which case the `current_user` is assumed. The object that is being inquired about can be specified either by name or by OID, too. When specifying by name, a schema name can be included if relevant. The access privilege of interest is specified by a text string, which must evaluate to one of the appropriate privilege keywords for the object's type (e.g., `SELECT`). Optionally, `WITH GRANT OPTION` can be added to a privilege type to test whether the privilege is held with grant option. Also, multiple privilege types can be listed separated by commas, in which case the result will be true if any of the listed privileges is held. (Case of the privilege string is not significant, and extra whitespace is allowed between but not within privilege names.) Some examples:

```sql

SELECT has_table_privilege('myschema.mytable', 'select');
SELECT has_table_privilege('joe', 'mytable', 'INSERT, SELECT WITH GRANT OPTION');
```


<a id="functions-info-access-table"></a>
**Table: Access Privilege Inquiry Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>has_any_column_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>table</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for any column of table? This succeeds either if the privilege is held for the whole table, or if there is a column-level grant of the privilege for at least one column. Allowable privilege types are <code>SELECT</code>, <code>INSERT</code>, <code>UPDATE</code>, and <code>REFERENCES</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_column_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>table</code> <code>text</code> or <code>oid</code>, <code>column</code> <code>text</code> or <code>smallint</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for the specified table column? This succeeds either if the privilege is held for the whole table, or if there is a column-level grant of the privilege for the column. The column can be specified by name or by attribute number (<code>pg_attribute</code>.<code>attnum</code>). Allowable privilege types are <code>SELECT</code>, <code>INSERT</code>, <code>UPDATE</code>, and <code>REFERENCES</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_database_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>database</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for database? Allowable privilege types are <code>CREATE</code>, <code>CONNECT</code>, <code>TEMPORARY</code>, and <code>TEMP</code> (which is equivalent to <code>TEMPORARY</code>).</td>
<td></td>
</tr>
<tr>
<td><code>has_foreign_data_wrapper_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>fdw</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for foreign-data wrapper? The only allowable privilege type is <code>USAGE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_function_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>function</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for function? The only allowable privilege type is <code>EXECUTE</code>.</td>
<td><p>When specifying a function by name rather than by OID, the allowed input is the same as for the <code>regprocedure</code> data type (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>). An example is:</p>
<pre><code class="language-sql">
SELECT has_function_privilege('joeuser', 'myfunc(int, text)', 'execute');</code></pre></td>
</tr>
<tr>
<td><code>has_language_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>language</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for language? The only allowable privilege type is <code>USAGE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_largeobject_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>largeobject</code> <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for large object? Allowable privilege types are <code>SELECT</code> and <code>UPDATE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_parameter_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>parameter</code> <code>text</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for configuration parameter? The parameter name is case-insensitive. Allowable privilege types are <code>SET</code> and <code>ALTER SYSTEM</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_schema_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>schema</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for schema? Allowable privilege types are <code>CREATE</code> and <code>USAGE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_sequence_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>sequence</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for sequence? Allowable privilege types are <code>USAGE</code>, <code>SELECT</code>, and <code>UPDATE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_server_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>server</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for foreign server? The only allowable privilege type is <code>USAGE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_table_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>table</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for table? Allowable privilege types are <code>SELECT</code>, <code>INSERT</code>, <code>UPDATE</code>, <code>DELETE</code>, <code>TRUNCATE</code>, <code>REFERENCES</code>, <code>TRIGGER</code>, and <code>MAINTAIN</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_tablespace_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>tablespace</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for tablespace? The only allowable privilege type is <code>CREATE</code>.</td>
<td></td>
</tr>
<tr>
<td><code>has_type_privilege</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>type</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for data type? The only allowable privilege type is <code>USAGE</code>. When specifying a type by name rather than by OID, the allowed input is the same as for the <code>regtype</code> data type (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>).</td>
<td></td>
</tr>
<tr>
<td><code>pg_has_role</code> ( [ <code>user</code> <code>name</code> or <code>oid</code>, ] <code>role</code> <code>text</code> or <code>oid</code>, <code>privilege</code> <code>text</code> ) <code>boolean</code></td>
<td>Does user have privilege for role? Allowable privilege types are <code>MEMBER</code>, <code>USAGE</code>, and <code>SET</code>. <code>MEMBER</code> denotes direct or indirect membership in the role without regard to what specific privileges may be conferred. <code>USAGE</code> denotes whether the privileges of the role are immediately available without doing <code>SET ROLE</code>, while <code>SET</code> denotes whether it is possible to change to the role using the <code>SET ROLE</code> command. <code>WITH ADMIN OPTION</code> or <code>WITH GRANT OPTION</code> can be added to any of these privilege types to test whether the <code>ADMIN</code> privilege is held (all six spellings test the same thing). This function does not allow the special case of setting <code>user</code> to <code>public</code>, because the PUBLIC pseudo-role can never be a member of real roles.</td>
<td></td>
</tr>
<tr>
<td><code>row_security_active</code> ( <code>table</code> <code>text</code> or <code>oid</code> ) <code>boolean</code></td>
<td>Is row-level security active for the specified table in the context of the current user and current environment?</td>
<td></td>
</tr>
</tbody>
</table>


 [`aclitem` Operators](#functions-aclitem-op-table) shows the operators available for the `aclitem` type, which is the catalog representation of access privileges. See [Privileges](../data-definition/privileges.md#ddl-priv) for information about how to read access privilege values.


<a id="functions-aclitem-op-table"></a>
**Table: `aclitem` Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>aclitem</code> <code>=</code> <code>aclitem</code> <code>boolean</code></td>
<td>Are <code>aclitem</code>s equal? (Notice that type <code>aclitem</code> lacks the usual set of comparison operators; it has only equality. In turn, <code>aclitem</code> arrays can only be compared for equality.)</td>
<td><code>'calvin=r<em>w/hobbes'::aclitem = 'calvin=r</em>w*/hobbes'::aclitem</code> <code>f</code></td>
</tr>
<tr>
<td><code>aclitem[]</code> <code>@&gt;</code> <code>aclitem</code> <code>boolean</code></td>
<td>Does array contain the specified privileges? (This is true if there is an array entry that matches the <code>aclitem</code>'s grantee and grantor, and has at least the specified set of privileges.)</td>
<td><code>'{calvin=r<em>w/hobbes,hobbes=r</em>w<em>/postgres}'::aclitem[] @&gt; 'calvin=r</em>/hobbes'::aclitem</code> <code>t</code></td>
</tr>
<tr>
<td><code>aclitem[]</code> <code>~</code> <code>aclitem</code> <code>boolean</code></td>
<td>This is a deprecated alias for <code>@&gt;</code>.</td>
<td><code>'{calvin=r<em>w/hobbes,hobbes=r</em>w<em>/postgres}'::aclitem[] ~ 'calvin=r</em>/hobbes'::aclitem</code> <code>t</code></td>
</tr>
</tbody>
</table>


 [`aclitem` Functions](#functions-aclitem-fn-table) shows some additional functions to manage the `aclitem` type.


<a id="functions-aclitem-fn-table"></a>
**Table: `aclitem` Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>acldefault</code> ( <code>type</code> <code>"char"</code>, <code>ownerId</code> <code>oid</code> ) <code>aclitem[]</code></td>
<td>Constructs an <code>aclitem</code> array holding the default access privileges for an object of type <code>type</code> belonging to the role with OID <code>ownerId</code>. This represents the access privileges that will be assumed when an object's ACL entry is null. (The default access privileges are described in <a href="../data-definition/privileges.md#ddl-priv">Privileges</a>.) The <code>type</code> parameter must be one of 'c' for <code>COLUMN</code>, 'r' for <code>TABLE</code> and table-like objects, 's' for <code>SEQUENCE</code>, 'd' for <code>DATABASE</code>, 'f' for <code>FUNCTION</code> or <code>PROCEDURE</code>, 'l' for <code>LANGUAGE</code>, 'L' for <code>LARGE OBJECT</code>, 'n' for <code>SCHEMA</code>, 'p' for <code>PARAMETER</code>, 't' for <code>TABLESPACE</code>, 'F' for <code>FOREIGN DATA WRAPPER</code>, 'S' for <code>FOREIGN SERVER</code>, or 'T' for <code>TYPE</code> or <code>DOMAIN</code>.</td>
<td></td>
</tr>
<tr>
<td><code>aclexplode</code> ( <code>aclitem[]</code> ) <code>setof record</code> ( <code>grantor</code> <code>oid</code>, <code>grantee</code> <code>oid</code>, <code>privilege_type</code> <code>text</code>, <code>is_grantable</code> <code>boolean</code> )</td>
<td>Returns the <code>aclitem</code> array as a set of rows. If the grantee is the pseudo-role PUBLIC, it is represented by zero in the <code>grantee</code> column. Each granted privilege is represented as <code>SELECT</code>, <code>INSERT</code>, etc (see <a href="../data-definition/privileges.md#privilege-abbrevs-table">ACL Privilege Abbreviations</a> for a full list). Note that each privilege is broken out as a separate row, so only one keyword appears in the <code>privilege_type</code> column.</td>
<td></td>
</tr>
<tr>
<td><code>makeaclitem</code> ( <code>grantee</code> <code>oid</code>, <code>grantor</code> <code>oid</code>, <code>privileges</code> <code>text</code>, <code>is_grantable</code> <code>boolean</code> ) <code>aclitem</code></td>
<td>Constructs an <code>aclitem</code> with the given properties. <code>privileges</code> is a comma-separated list of privilege names such as <code>SELECT</code>, <code>INSERT</code>, etc, all of which are set in the result. (Case of the privilege string is not significant, and extra whitespace is allowed between but not within privilege names.)</td>
<td></td>
</tr>
</tbody>
</table>


### Schema Visibility Inquiry Functions { #functions-info-schema }


 [Schema Visibility Inquiry Functions](#functions-info-schema-table) shows functions that determine whether a certain object is *visible* in the current schema search path. For example, a table is said to be visible if its containing schema is in the search path and no table of the same name appears earlier in the search path. This is equivalent to the statement that the table can be referenced by name without explicit schema qualification. Thus, to list the names of all visible tables:

```sql

SELECT relname FROM pg_class WHERE pg_table_is_visible(oid);
```
 For functions and operators, an object in the search path is said to be visible if there is no object of the same name *and argument data type(s)* earlier in the path. For operator classes and families, both the name and the associated index access method are considered.


<a id="functions-info-schema-table"></a>
**Table: Schema Visibility Inquiry Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_collation_is_visible</code> ( <code>collation</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is collation visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_conversion_is_visible</code> ( <code>conversion</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is conversion visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_function_is_visible</code> ( <code>function</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is function visible in search path? (This also works for procedures and aggregates.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_opclass_is_visible</code> ( <code>opclass</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is operator class visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_operator_is_visible</code> ( <code>operator</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is operator visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_opfamily_is_visible</code> ( <code>opclass</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is operator family visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_statistics_obj_is_visible</code> ( <code>stat</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is statistics object visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_table_is_visible</code> ( <code>table</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is table visible in search path? (This works for all types of relations, including views, materialized views, indexes, sequences and foreign tables.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_ts_config_is_visible</code> ( <code>config</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is text search configuration visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_ts_dict_is_visible</code> ( <code>dict</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is text search dictionary visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_ts_parser_is_visible</code> ( <code>parser</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is text search parser visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_ts_template_is_visible</code> ( <code>template</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is text search template visible in search path?</td>
<td></td>
</tr>
<tr>
<td><code>pg_type_is_visible</code> ( <code>type</code> <code>oid</code> ) <code>boolean</code></td>
<td>Is type (or domain) visible in search path?</td>
<td></td>
</tr>
</tbody>
</table>


 All these functions require object OIDs to identify the object to be checked. If you want to test an object by name, it is convenient to use the OID alias types (`regclass`, `regtype`, `regprocedure`, `regoperator`, `regconfig`, or `regdictionary`), for example:

```sql

SELECT pg_type_is_visible('myschema.widget'::regtype);
```
 Note that it would not make much sense to test a non-schema-qualified type name in this way — if the name can be recognized at all, it must be visible.


### System Catalog Information Functions { #functions-info-catalog }


 [System Catalog Information Functions](#functions-info-catalog-table) lists functions that extract information from the system catalogs.


<a id="functions-info-catalog-table"></a>
**Table: System Catalog Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>format_type</code> ( <code>type</code> <code>oid</code>, <code>typemod</code> <code>integer</code> ) <code>text</code></td>
<td>Returns the SQL name for a data type that is identified by its type OID and possibly a type modifier. Pass NULL for the type modifier if no specific modifier is known.</td>
<td></td>
</tr>
<tr>
<td><code>pg_basetype</code> ( <code>regtype</code> ) <code>regtype</code></td>
<td>Returns the OID of the base type of a domain identified by its type OID. If the argument is the OID of a non-domain type, returns the argument as-is. Returns NULL if the argument is not a valid type OID. If there's a chain of domain dependencies, it will recurse until finding the base type.</td>
<td>Assuming <code>CREATE DOMAIN mytext AS text</code>:<br><code>pg_basetype('mytext'::regtype)</code> <code>text</code></td>
</tr>
<tr>
<td><code>pg_char_to_encoding</code> ( <code>encoding</code> <code>name</code> ) <code>integer</code></td>
<td>Converts the supplied encoding name into an integer representing the internal identifier used in some system catalog tables. Returns <code>-1</code> if an unknown encoding name is provided.</td>
<td></td>
</tr>
<tr>
<td><code>pg_encoding_to_char</code> ( <code>encoding</code> <code>integer</code> ) <code>name</code></td>
<td>Converts the integer used as the internal identifier of an encoding in some system catalog tables into a human-readable string. Returns an empty string if an invalid encoding number is provided.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_catalog_foreign_keys</code> () <code>setof record</code> ( <code>fktable</code> <code>regclass</code>, <code>fkcols</code> <code>text[]</code>, <code>pktable</code> <code>regclass</code>, <code>pkcols</code> <code>text[]</code>, <code>is_array</code> <code>boolean</code>, <code>is_opt</code> <code>boolean</code> )</td>
<td>Returns a set of records describing the foreign key relationships that exist within the PostgreSQL system catalogs. The <code>fktable</code> column contains the name of the referencing catalog, and the <code>fkcols</code> column contains the name(s) of the referencing column(s). Similarly, the <code>pktable</code> column contains the name of the referenced catalog, and the <code>pkcols</code> column contains the name(s) of the referenced column(s). If <code>is_array</code> is true, the last referencing column is an array, each of whose elements should match some entry in the referenced catalog. If <code>is_opt</code> is true, the referencing column(s) are allowed to contain zeroes instead of a valid reference.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_constraintdef</code> ( <code>constraint</code> <code>oid</code> [, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Reconstructs the creating command for a constraint. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_expr</code> ( <code>expr</code> <code>pg_node_tree</code>, <code>relation</code> <code>oid</code> [, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Decompiles the internal form of an expression stored in the system catalogs, such as the default value for a column. If the expression might contain Vars, specify the OID of the relation they refer to as the second parameter; if no Vars are expected, passing zero is sufficient.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_functiondef</code> ( <code>func</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the creating command for a function or procedure. (This is a decompiled reconstruction, not the original text of the command.) The result is a complete <code>CREATE OR REPLACE FUNCTION</code> or <code>CREATE OR REPLACE PROCEDURE</code> statement.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_function_arguments</code> ( <code>func</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the argument list of a function or procedure, in the form it would need to appear in within <code>CREATE FUNCTION</code> (including default values).</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_function_identity_arguments</code> ( <code>func</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the argument list necessary to identify a function or procedure, in the form it would need to appear in within commands such as <code>ALTER FUNCTION</code>. This form omits default values.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_function_result</code> ( <code>func</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the <code>RETURNS</code> clause of a function, in the form it would need to appear in within <code>CREATE FUNCTION</code>. Returns <code>NULL</code> for a procedure.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_indexdef</code> ( <code>index</code> <code>oid</code> [, <code>column</code> <code>integer</code>, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Reconstructs the creating command for an index. (This is a decompiled reconstruction, not the original text of the command.) If <code>column</code> is supplied and is not zero, only the definition of that column is reconstructed.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_keywords</code> () <code>setof record</code> ( <code>word</code> <code>text</code>, <code>catcode</code> <code>"char"</code>, <code>barelabel</code> <code>boolean</code>, <code>catdesc</code> <code>text</code>, <code>baredesc</code> <code>text</code> )</td>
<td>Returns a set of records describing the SQL keywords recognized by the server. The <code>word</code> column contains the keyword. The <code>catcode</code> column contains a category code: <code>U</code> for an unreserved keyword, <code>C</code> for a keyword that can be a column name, <code>T</code> for a keyword that can be a type or function name, or <code>R</code> for a fully reserved keyword. The <code>barelabel</code> column contains <code>true</code> if the keyword can be used as a “bare” column label in <code>SELECT</code> lists, or <code>false</code> if it can only be used after <code>AS</code>. The <code>catdesc</code> column contains a possibly-localized string describing the keyword's category. The <code>baredesc</code> column contains a possibly-localized string describing the keyword's column label status.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_partition_constraintdef</code> ( <code>table</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the definition of a partition constraint. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_partkeydef</code> ( <code>table</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the definition of a partitioned table's partition key, in the form it would have in the <code>PARTITION BY</code> clause of <code>CREATE TABLE</code>. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_ruledef</code> ( <code>rule</code> <code>oid</code> [, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Reconstructs the creating command for a rule. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_serial_sequence</code> ( <code>table</code> <code>text</code>, <code>column</code> <code>text</code> ) <code>text</code></td>
<td>Returns the name of the sequence associated with a column, or NULL if no sequence is associated with the column. If the column is an identity column, the associated sequence is the sequence internally created for that column. For columns created using one of the serial types (<code>serial</code>, <code>smallserial</code>, <code>bigserial</code>), it is the sequence created for that serial column definition. In the latter case, the association can be modified or removed with <code>ALTER SEQUENCE OWNED BY</code>. (This function probably should have been called <code>pg_get_owned_sequence</code>; its current name reflects the fact that it has historically been used with serial-type columns.) The first parameter is a table name with optional schema, and the second parameter is a column name. Because the first parameter potentially contains both schema and table names, it is parsed per usual SQL rules, meaning it is lower-cased by default. The second parameter, being just a column name, is treated literally and so has its case preserved. The result is suitably formatted for passing to the sequence functions (see <a href="sequence-manipulation-functions.md#functions-sequence">Sequence Manipulation Functions</a>).</td>
<td><p>A typical use is in reading the current value of the sequence for an identity or serial column, for example:</p>
<pre><code class="language-sql">
SELECT currval(pg_get_serial_sequence('sometable', 'id'));</code></pre></td>
</tr>
<tr>
<td><code>pg_get_statisticsobjdef</code> ( <code>statobj</code> <code>oid</code> ) <code>text</code></td>
<td>Reconstructs the creating command for an extended statistics object. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_triggerdef</code> ( <code>trigger</code> <code>oid</code> [, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Reconstructs the creating command for a trigger. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_userbyid</code> ( <code>role</code> <code>oid</code> ) <code>name</code></td>
<td>Returns a role's name given its OID.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_viewdef</code> ( <code>view</code> <code>oid</code> [, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Reconstructs the underlying <code>SELECT</code> command for a view or materialized view. (This is a decompiled reconstruction, not the original text of the command.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_viewdef</code> ( <code>view</code> <code>oid</code>, <code>wrap_column</code> <code>integer</code> ) <code>text</code></td>
<td>Reconstructs the underlying <code>SELECT</code> command for a view or materialized view. (This is a decompiled reconstruction, not the original text of the command.) In this form of the function, pretty-printing is always enabled, and long lines are wrapped to try to keep them shorter than the specified number of columns.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_viewdef</code> ( <code>view</code> <code>text</code> [, <code>pretty</code> <code>boolean</code> ] ) <code>text</code></td>
<td>Reconstructs the underlying <code>SELECT</code> command for a view or materialized view, working from a textual name for the view rather than its OID. (This is deprecated; use the OID variant instead.)</td>
<td></td>
</tr>
<tr>
<td><code>pg_index_column_has_property</code> ( <code>index</code> <code>regclass</code>, <code>column</code> <code>integer</code>, <code>property</code> <code>text</code> ) <code>boolean</code></td>
<td>Tests whether an index column has the named property. Common index column properties are listed in <a href="#functions-info-index-column-props">Index Column Properties</a>. (Note that extension access methods can define additional property names for their indexes.) <code>NULL</code> is returned if the property name is not known or does not apply to the particular object, or if the OID or column number does not identify a valid object.</td>
<td></td>
</tr>
<tr>
<td><code>pg_index_has_property</code> ( <code>index</code> <code>regclass</code>, <code>property</code> <code>text</code> ) <code>boolean</code></td>
<td>Tests whether an index has the named property. Common index properties are listed in <a href="#functions-info-index-props">Index Properties</a>. (Note that extension access methods can define additional property names for their indexes.) <code>NULL</code> is returned if the property name is not known or does not apply to the particular object, or if the OID does not identify a valid object.</td>
<td></td>
</tr>
<tr>
<td><code>pg_indexam_has_property</code> ( <code>am</code> <code>oid</code>, <code>property</code> <code>text</code> ) <code>boolean</code></td>
<td>Tests whether an index access method has the named property. Access method properties are listed in <a href="#functions-info-indexam-props">Index Access Method Properties</a>. <code>NULL</code> is returned if the property name is not known or does not apply to the particular object, or if the OID does not identify a valid object.</td>
<td></td>
</tr>
<tr>
<td><code>pg_options_to_table</code> ( <code>options_array</code> <code>text[]</code> ) <code>setof record</code> ( <code>option_name</code> <code>text</code>, <code>option_value</code> <code>text</code> )</td>
<td>Returns the set of storage options represented by a value from <code>pg_class</code>.<code>reloptions</code> or <code>pg_attribute</code>.<code>attoptions</code>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_settings_get_flags</code> ( <code>guc</code> <code>text</code> ) <code>text[]</code></td>
<td>Returns an array of the flags associated with the given GUC, or <code>NULL</code> if it does not exist. The result is an empty array if the GUC exists but there are no flags to show. Only the most useful flags listed in <a href="#functions-pg-settings-flags">GUC Flags</a> are exposed.</td>
<td></td>
</tr>
<tr>
<td><code>pg_tablespace_databases</code> ( <code>tablespace</code> <code>oid</code> ) <code>setof oid</code></td>
<td>Returns the set of OIDs of databases that have objects stored in the specified tablespace. If this function returns any rows, the tablespace is not empty and cannot be dropped. To identify the specific objects populating the tablespace, you will need to connect to the database(s) identified by <code>pg_tablespace_databases</code> and query their <code>pg_class</code> catalogs.</td>
<td></td>
</tr>
<tr>
<td><code>pg_tablespace_location</code> ( <code>tablespace</code> <code>oid</code> ) <code>text</code></td>
<td>Returns the file system path that this tablespace is located in.</td>
<td></td>
</tr>
<tr>
<td><code>pg_typeof</code> ( <code>"any"</code> ) <code>regtype</code></td>
<td>Returns the OID of the data type of the value that is passed to it. This can be helpful for troubleshooting or dynamically constructing SQL queries. The function is declared as returning <code>regtype</code>, which is an OID alias type (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); this means that it is the same as an OID for comparison purposes but displays as a type name.</td>
<td><code>pg_typeof(33)</code> <code>integer</code></td>
</tr>
<tr>
<td><code>COLLATION FOR</code> ( <code>"any"</code> ) <code>text</code></td>
<td>Returns the name of the collation of the value that is passed to it. The value is quoted and schema-qualified if necessary. If no collation was derived for the argument expression, then <code>NULL</code> is returned. If the argument is not of a collatable data type, then an error is raised.</td>
<td><code>collation for ('foo'::text)</code> <code>"default"</code><br><code>collation for ('foo' COLLATE "de_DE")</code> <code>"de_DE"</code></td>
</tr>
<tr>
<td><code>to_regclass</code> ( <code>text</code> ) <code>regclass</code></td>
<td>Translates a textual relation name to its OID. A similar result is obtained by casting the string to type <code>regclass</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regcollation</code> ( <code>text</code> ) <code>regcollation</code></td>
<td>Translates a textual collation name to its OID. A similar result is obtained by casting the string to type <code>regcollation</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regnamespace</code> ( <code>text</code> ) <code>regnamespace</code></td>
<td>Translates a textual schema name to its OID. A similar result is obtained by casting the string to type <code>regnamespace</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regoper</code> ( <code>text</code> ) <code>regoper</code></td>
<td>Translates a textual operator name to its OID. A similar result is obtained by casting the string to type <code>regoper</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found or is ambiguous.</td>
<td></td>
</tr>
<tr>
<td><code>to_regoperator</code> ( <code>text</code> ) <code>regoperator</code></td>
<td>Translates a textual operator name (with parameter types) to its OID. A similar result is obtained by casting the string to type <code>regoperator</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regproc</code> ( <code>text</code> ) <code>regproc</code></td>
<td>Translates a textual function or procedure name to its OID. A similar result is obtained by casting the string to type <code>regproc</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found or is ambiguous.</td>
<td></td>
</tr>
<tr>
<td><code>to_regprocedure</code> ( <code>text</code> ) <code>regprocedure</code></td>
<td>Translates a textual function or procedure name (with argument types) to its OID. A similar result is obtained by casting the string to type <code>regprocedure</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regrole</code> ( <code>text</code> ) <code>regrole</code></td>
<td>Translates a textual role name to its OID. A similar result is obtained by casting the string to type <code>regrole</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>); however, this function will return <code>NULL</code> rather than throwing an error if the name is not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regtype</code> ( <code>text</code> ) <code>regtype</code></td>
<td>Parses a string of text, extracts a potential type name from it, and translates that name into a type OID. A syntax error in the string will result in an error; but if the string is a syntactically valid type name that happens not to be found in the catalogs, the result is <code>NULL</code>. A similar result is obtained by casting the string to type <code>regtype</code> (see <a href="../data-types/object-identifier-types.md#datatype-oid">Object Identifier Types</a>), except that that will throw error for name not found.</td>
<td></td>
</tr>
<tr>
<td><code>to_regtypemod</code> ( <code>text</code> ) <code>integer</code></td>
<td>Parses a string of text, extracts a potential type name from it, and translates its type modifier, if any. A syntax error in the string will result in an error; but if the string is a syntactically valid type name that happens not to be found in the catalogs, the result is <code>NULL</code>. The result is <code>-1</code> if no type modifier is present.</td>
<td><code>to_regtypemod</code> can be combined with <a href="#to-regtype">to_regtype</a> to produce appropriate inputs for <a href="#format-type">format_type</a>, allowing a string representing a type name to be canonicalized.<br><code>format_type(to_regtype('varchar(32)'), to_regtypemod('varchar(32)'))</code> <code>character varying(32)</code></td>
</tr>
</tbody>
</table>


 Most of the functions that reconstruct (decompile) database objects have an optional `pretty` flag, which if `true` causes the result to be “pretty-printed”. Pretty-printing suppresses unnecessary parentheses and adds whitespace for legibility. The pretty-printed format is more readable, but the default format is more likely to be interpreted the same way by future versions of PostgreSQL; so avoid using pretty-printed output for dump purposes. Passing `false` for the `pretty` parameter yields the same result as omitting the parameter.


<a id="functions-info-index-column-props"></a>
**Table: Index Column Properties**

| Name | Description |
| --- | --- |
| `asc` | Does the column sort in ascending order on a forward scan? |
| `desc` | Does the column sort in descending order on a forward scan? |
| `nulls_first` | Does the column sort with nulls first on a forward scan? |
| `nulls_last` | Does the column sort with nulls last on a forward scan? |
| `orderable` | Does the column possess any defined sort ordering? |
| `distance_orderable` | Can the column be scanned in order by a “distance” operator, for example `ORDER BY col <-> constant` ? |
| `returnable` | Can the column value be returned by an index-only scan? |
| `search_array` | Does the column natively support `col = ANY(array)` searches? |
| `search_nulls` | Does the column support `IS NULL` and `IS NOT NULL` searches? |


<a id="functions-info-index-props"></a>
**Table: Index Properties**

| Name | Description |
| --- | --- |
| `clusterable` | Can the index be used in a `CLUSTER` command? |
| `index_scan` | Does the index support plain (non-bitmap) scans? |
| `bitmap_scan` | Does the index support bitmap scans? |
| `backward_scan` | Can the scan direction be changed in mid-scan (to support `FETCH BACKWARD` on a cursor without needing materialization)? |


<a id="functions-info-indexam-props"></a>
**Table: Index Access Method Properties**

| Name | Description |
| --- | --- |
| `can_order` | Does the access method support `ASC`, `DESC` and related keywords in `CREATE INDEX`? |
| `can_unique` | Does the access method support unique indexes? |
| `can_multi_col` | Does the access method support indexes with multiple columns? |
| `can_exclude` | Does the access method support exclusion constraints? |
| `can_include` | Does the access method support the `INCLUDE` clause of `CREATE INDEX`? |


<a id="functions-pg-settings-flags"></a>
**Table: GUC Flags**

| Flag | Description |
| --- | --- |
| `EXPLAIN` | Parameters with this flag are included in `EXPLAIN (SETTINGS)` commands. |
| `NO_SHOW_ALL` | Parameters with this flag are excluded from `SHOW ALL` commands. |
| `NO_RESET` | Parameters with this flag do not support `RESET` commands. |
| `NO_RESET_ALL` | Parameters with this flag are excluded from `RESET ALL` commands. |
| `NOT_IN_SAMPLE` | Parameters with this flag are not included in `postgresql.conf` by default. |
| `RUNTIME_COMPUTED` | Parameters with this flag are runtime-computed ones. |


### Object Information and Addressing Functions { #functions-info-object }


 [Object Information and Addressing Functions](#functions-info-object-table) lists functions related to database object identification and addressing.


<a id="functions-info-object-table"></a>
**Table: Object Information and Addressing Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_get_acl</code> ( <code>classid</code> <code>oid</code>, <code>objid</code> <code>oid</code>, <code>objsubid</code> <code>integer</code> ) <code>aclitem[]</code></td>
<td>Returns the ACL for a database object, specified by catalog OID, object OID and sub-object ID. This function returns <code>NULL</code> values for undefined objects.</td>
<td></td>
</tr>
<tr>
<td><code>pg_describe_object</code> ( <code>classid</code> <code>oid</code>, <code>objid</code> <code>oid</code>, <code>objsubid</code> <code>integer</code> ) <code>text</code></td>
<td>Returns a textual description of a database object identified by catalog OID, object OID, and sub-object ID (such as a column number within a table; the sub-object ID is zero when referring to a whole object). This description is intended to be human-readable, and might be translated, depending on server configuration. This is especially useful to determine the identity of an object referenced in the <code>pg_depend</code> catalog. This function returns <code>NULL</code> values for undefined objects.</td>
<td></td>
</tr>
<tr>
<td><code>pg_identify_object</code> ( <code>classid</code> <code>oid</code>, <code>objid</code> <code>oid</code>, <code>objsubid</code> <code>integer</code> ) <code>record</code> ( <code>type</code> <code>text</code>, <code>schema</code> <code>text</code>, <code>name</code> <code>text</code>, <code>identity</code> <code>text</code> )</td>
<td>Returns a row containing enough information to uniquely identify the database object specified by catalog OID, object OID and sub-object ID. This information is intended to be machine-readable, and is never translated. <code>type</code> identifies the type of database object; <code>schema</code> is the schema name that the object belongs in, or <code>NULL</code> for object types that do not belong to schemas; <code>name</code> is the name of the object, quoted if necessary, if the name (along with schema name, if pertinent) is sufficient to uniquely identify the object, otherwise <code>NULL</code>; <code>identity</code> is the complete object identity, with the precise format depending on object type, and each name within the format being schema-qualified and quoted as necessary. Undefined objects are identified with <code>NULL</code> values.</td>
<td></td>
</tr>
<tr>
<td><code>pg_identify_object_as_address</code> ( <code>classid</code> <code>oid</code>, <code>objid</code> <code>oid</code>, <code>objsubid</code> <code>integer</code> ) <code>record</code> ( <code>type</code> <code>text</code>, <code>object_names</code> <code>text[]</code>, <code>object_args</code> <code>text[]</code> )</td>
<td>Returns a row containing enough information to uniquely identify the database object specified by catalog OID, object OID and sub-object ID. The returned information is independent of the current server, that is, it could be used to identify an identically named object in another server. <code>type</code> identifies the type of database object; <code>object_names</code> and <code>object_args</code> are text arrays that together form a reference to the object. These three values can be passed to <code>pg_get_object_address</code> to obtain the internal address of the object.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_object_address</code> ( <code>type</code> <code>text</code>, <code>object_names</code> <code>text[]</code>, <code>object_args</code> <code>text[]</code> ) <code>record</code> ( <code>classid</code> <code>oid</code>, <code>objid</code> <code>oid</code>, <code>objsubid</code> <code>integer</code> )</td>
<td>Returns a row containing enough information to uniquely identify the database object specified by a type code and object name and argument arrays. The returned values are the ones that would be used in system catalogs such as <code>pg_depend</code>; they can be passed to other system functions such as <code>pg_describe_object</code> or <code>pg_identify_object</code>. <code>classid</code> is the OID of the system catalog containing the object; <code>objid</code> is the OID of the object itself, and <code>objsubid</code> is the sub-object ID, or zero if none. This function is the inverse of <code>pg_identify_object_as_address</code>. Undefined objects are identified with <code>NULL</code> values.</td>
<td></td>
</tr>
</tbody>
</table>


 `pg_get_acl` is useful for retrieving and inspecting the privileges associated with database objects without looking at specific catalogs. For example, to retrieve all the granted privileges on objects in the current database:

```

postgres=# SELECT
    (pg_identify_object(s.classid,s.objid,s.objsubid)).*,
    pg_catalog.pg_get_acl(s.classid,s.objid,s.objsubid) AS acl
FROM pg_catalog.pg_shdepend AS s
JOIN pg_catalog.pg_database AS d
    ON d.datname = current_database() AND
       d.oid = s.dbid
JOIN pg_catalog.pg_authid AS a
    ON a.oid = s.refobjid AND
       s.refclassid = 'pg_authid'::regclass
WHERE s.deptype = 'a';
-[ RECORD 1 ]-----------------------------------------
type     | table
schema   | public
name     | testtab
identity | public.testtab
acl      | {postgres=arwdDxtm/postgres,foo=r/postgres}
```


### Comment Information Functions { #functions-info-comment }


 The functions shown in [Comment Information Functions](#functions-info-comment-table) extract comments previously stored with the [sql-comment](../../reference/sql-commands/comment.md#sql-comment) command. A null value is returned if no comment could be found for the specified parameters.


<a id="functions-info-comment-table"></a>
**Table: Comment Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>col_description</code> ( <code>table</code> <code>oid</code>, <code>column</code> <code>integer</code> ) <code>text</code></td>
<td>Returns the comment for a table column, which is specified by the OID of its table and its column number. (<code>obj_description</code> cannot be used for table columns, since columns do not have OIDs of their own.)</td>
<td></td>
</tr>
<tr>
<td><code>obj_description</code> ( <code>object</code> <code>oid</code>, <code>catalog</code> <code>name</code> ) <code>text</code></td>
<td>Returns the comment for a database object specified by its OID and the name of the containing system catalog. For example, <code>obj_description(123456, 'pg_class')</code> would retrieve the comment for the table with OID 123456.</td>
<td></td>
</tr>
<tr>
<td><code>obj_description</code> ( <code>object</code> <code>oid</code> ) <code>text</code></td>
<td>Returns the comment for a database object specified by its OID alone. This is <em>deprecated</em> since there is no guarantee that OIDs are unique across different system catalogs; therefore, the wrong comment might be returned.</td>
<td></td>
</tr>
<tr>
<td><code>shobj_description</code> ( <code>object</code> <code>oid</code>, <code>catalog</code> <code>name</code> ) <code>text</code></td>
<td>Returns the comment for a shared database object specified by its OID and the name of the containing system catalog. This is just like <code>obj_description</code> except that it is used for retrieving comments on shared objects (that is, databases, roles, and tablespaces). Some system catalogs are global to all databases within each cluster, and the descriptions for objects in them are stored globally as well.</td>
<td></td>
</tr>
</tbody>
</table>


### Data Validity Checking Functions { #functions-info-validity }


 The functions shown in [Data Validity Checking Functions](#functions-info-validity-table) can be helpful for checking validity of proposed input data.


<a id="functions-info-validity-table"></a>
**Table: Data Validity Checking Functions**

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
<td><code>pg_input_is_valid</code> ( <code>string</code> <code>text</code>, <code>type</code> <code>text</code> ) <code>boolean</code></td>
<td>Tests whether the given <code>string</code> is valid input for the specified data type, returning true or false.</td>
<td>This function will only work as desired if the data type's input function has been updated to report invalid input as a “soft” error. Otherwise, invalid input will abort the transaction, just as if the string had been cast to the type directly.<br><code>pg_input_is_valid('42', 'integer')</code> <code>t</code><br><code>pg_input_is_valid('42000000000', 'integer')</code> <code>f</code><br><code>pg_input_is_valid('1234.567', 'numeric(7,4)')</code> <code>f</code></td>
</tr>
<tr>
<td><code>pg_input_error_info</code> ( <code>string</code> <code>text</code>, <code>type</code> <code>text</code> ) <code>record</code> ( <code>message</code> <code>text</code>, <code>detail</code> <code>text</code>, <code>hint</code> <code>text</code>, <code>sql_error_code</code> <code>text</code> )</td>
<td>Tests whether the given <code>string</code> is valid input for the specified data type; if not, return the details of the error that would have been thrown. If the input is valid, the results are NULL. The inputs are the same as for <code>pg_input_is_valid</code>.</td>
<td>This function will only work as desired if the data type's input function has been updated to report invalid input as a “soft” error. Otherwise, invalid input will abort the transaction, just as if the string had been cast to the type directly.<br><p><code>SELECT * FROM pg_input_error_info('42000000000', 'integer')</code></p>
<pre><code>
                       message                        | detail | hint | sql_error_code
------------------------------------------------------+--------+------+----------------
 value "42000000000" is out of range for type integer |        |      | 22003</code></pre></td>
</tr>
</tbody>
</table>


### Transaction ID and Snapshot Information Functions { #functions-info-snapshot }


 The functions shown in [Transaction ID and Snapshot Information Functions](#functions-pg-snapshot) provide server transaction information in an exportable form. The main use of these functions is to determine which transactions were committed between two snapshots.


<a id="functions-pg-snapshot"></a>
**Table: Transaction ID and Snapshot Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>age</code> ( <code>xid</code> ) <code>integer</code></td>
<td>Returns the number of transactions between the supplied transaction id and the current transaction counter.</td>
<td></td>
</tr>
<tr>
<td><code>mxid_age</code> ( <code>xid</code> ) <code>integer</code></td>
<td>Returns the number of multixacts IDs between the supplied multixact ID and the current multixacts counter.</td>
<td></td>
</tr>
<tr>
<td><code>pg_current_xact_id</code> () <code>xid8</code></td>
<td>Returns the current transaction's ID. It will assign a new one if the current transaction does not have one already (because it has not performed any database updates); see <a href="../../internals/transaction-processing/transactions-and-identifiers.md#transaction-id">Transactions and Identifiers</a> for details. If executed in a subtransaction, this will return the top-level transaction ID; see <a href="../../internals/transaction-processing/subtransactions.md#subxacts">Subtransactions</a> for details.</td>
<td></td>
</tr>
<tr>
<td><code>pg_current_xact_id_if_assigned</code> () <code>xid8</code></td>
<td>Returns the current transaction's ID, or <code>NULL</code> if no ID is assigned yet. (It's best to use this variant if the transaction might otherwise be read-only, to avoid unnecessary consumption of an XID.) If executed in a subtransaction, this will return the top-level transaction ID.</td>
<td></td>
</tr>
<tr>
<td><code>pg_xact_status</code> ( <code>xid8</code> ) <code>text</code></td>
<td>Reports the commit status of a recent transaction. The result is one of <code>in progress</code>, <code>committed</code>, or <code>aborted</code>, provided that the transaction is recent enough that the system retains the commit status of that transaction. If it is old enough that no references to the transaction survive in the system and the commit status information has been discarded, the result is <code>NULL</code>. Applications might use this function, for example, to determine whether their transaction committed or aborted after the application and database server become disconnected while a <code>COMMIT</code> is in progress. Note that prepared transactions are reported as <code>in progress</code>; applications must check <a href="../../internals/system-views/pg_prepared_xacts.md#view-pg-prepared-xacts"><code>pg_prepared_xacts</code></a> if they need to determine whether a transaction ID belongs to a prepared transaction.</td>
<td></td>
</tr>
<tr>
<td><code>pg_current_snapshot</code> () <code>pg_snapshot</code></td>
<td>Returns a current <em>snapshot</em>, a data structure showing which transaction IDs are now in-progress. Only top-level transaction IDs are included in the snapshot; subtransaction IDs are not shown; see <a href="../../internals/transaction-processing/subtransactions.md#subxacts">Subtransactions</a> for details.</td>
<td></td>
</tr>
<tr>
<td><code>pg_snapshot_xip</code> ( <code>pg_snapshot</code> ) <code>setof xid8</code></td>
<td>Returns the set of in-progress transaction IDs contained in a snapshot.</td>
<td></td>
</tr>
<tr>
<td><code>pg_snapshot_xmax</code> ( <code>pg_snapshot</code> ) <code>xid8</code></td>
<td>Returns the <code>xmax</code> of a snapshot.</td>
<td></td>
</tr>
<tr>
<td><code>pg_snapshot_xmin</code> ( <code>pg_snapshot</code> ) <code>xid8</code></td>
<td>Returns the <code>xmin</code> of a snapshot.</td>
<td></td>
</tr>
<tr>
<td><code>pg_visible_in_snapshot</code> ( <code>xid8</code>, <code>pg_snapshot</code> ) <code>boolean</code></td>
<td>Is the given transaction ID <em>visible</em> according to this snapshot (that is, was it completed before the snapshot was taken)? Note that this function will not give the correct answer for a subtransaction ID (subxid); see <a href="../../internals/transaction-processing/subtransactions.md#subxacts">Subtransactions</a> for details.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_multixact_members</code> ( <code>multixid</code> <code>xid</code> ) <code>setof record</code> ( <code>xid</code> <code>xid</code>, <code>mode</code> <code>text</code> )</td>
<td>Returns the transaction ID and lock mode for each member of the specified multixact ID. The lock modes <code>forupd</code>, <code>fornokeyupd</code>, <code>sh</code>, and <code>keysh</code> correspond to the row-level locks <code>FOR UPDATE</code>, <code>FOR NO KEY UPDATE</code>, <code>FOR SHARE</code>, and <code>FOR KEY SHARE</code>, respectively, as described in <a href="../concurrency-control/explicit-locking.md#locking-rows">Row-Level Locks</a>. Two additional modes are specific to multixacts: <code>nokeyupd</code>, used by updates that do not modify key columns, and <code>upd</code>, used by updates or deletes that modify key columns.</td>
<td></td>
</tr>
</tbody>
</table>


 The internal transaction ID type `xid` is 32 bits wide and wraps around every 4 billion transactions. However, the functions shown in [Transaction ID and Snapshot Information Functions](#functions-pg-snapshot), except `age`, `mxid_age`, and `pg_get_multixact_members`, use a 64-bit type `xid8` that does not wrap around during the life of an installation and can be converted to `xid` by casting if required; see [Transactions and Identifiers](../../internals/transaction-processing/transactions-and-identifiers.md#transaction-id) for details. The data type `pg_snapshot` stores information about transaction ID visibility at a particular moment in time. Its components are described in [Snapshot Components](#functions-pg-snapshot-parts). `pg_snapshot`'s textual representation is <em>xmin</em><code>:</code><em>xmax</em><code>:</code><em>xip_list</em>. For example `10:20:10,14,15` means `xmin=10, xmax=20, xip_list=10, 14, 15`.


<a id="functions-pg-snapshot-parts"></a>
**Table: Snapshot Components**

| Name | Description |
| --- | --- |
| `xmin` | Lowest transaction ID that was still active. All transaction IDs less than `xmin` are either committed and visible, or rolled back and dead. |
| `xmax` | One past the highest completed transaction ID. All transaction IDs greater than or equal to `xmax` had not yet completed as of the time of the snapshot, and thus are invisible. |
| `xip_list` | Transactions in progress at the time of the snapshot. A transaction ID that is <code>xmin <= </code><em>X</em><code> <         xmax</code> and not in this list was already completed at the time of the snapshot, and thus is either visible or dead according to its commit status. This list does not include the transaction IDs of subtransactions (subxids). |


 In releases of PostgreSQL before 13 there was no `xid8` type, so variants of these functions were provided that used `bigint` to represent a 64-bit XID, with a correspondingly distinct snapshot data type `txid_snapshot`. These older functions have `txid` in their names. They are still supported for backward compatibility, but may be removed from a future release. See [Deprecated Transaction ID and Snapshot Information Functions](#functions-txid-snapshot).


<a id="functions-txid-snapshot"></a>
**Table: Deprecated Transaction ID and Snapshot Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>txid_current</code> () <code>bigint</code></td>
<td>See <code>pg_current_xact_id()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_current_if_assigned</code> () <code>bigint</code></td>
<td>See <code>pg_current_xact_id_if_assigned()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_current_snapshot</code> () <code>txid_snapshot</code></td>
<td>See <code>pg_current_snapshot()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_snapshot_xip</code> ( <code>txid_snapshot</code> ) <code>setof bigint</code></td>
<td>See <code>pg_snapshot_xip()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_snapshot_xmax</code> ( <code>txid_snapshot</code> ) <code>bigint</code></td>
<td>See <code>pg_snapshot_xmax()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_snapshot_xmin</code> ( <code>txid_snapshot</code> ) <code>bigint</code></td>
<td>See <code>pg_snapshot_xmin()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_visible_in_snapshot</code> ( <code>bigint</code>, <code>txid_snapshot</code> ) <code>boolean</code></td>
<td>See <code>pg_visible_in_snapshot()</code>.</td>
<td></td>
</tr>
<tr>
<td><code>txid_status</code> ( <code>bigint</code> ) <code>text</code></td>
<td>See <code>pg_xact_status()</code>.</td>
<td></td>
</tr>
</tbody>
</table>


### Committed Transaction Information Functions { #functions-info-commit-timestamp }


 The functions shown in [Committed Transaction Information Functions](#functions-commit-timestamp) provide information about when past transactions were committed. They only provide useful data when the [track_commit_timestamp](../../server-administration/server-configuration/replication.md#guc-track-commit-timestamp) configuration option is enabled, and only for transactions that were committed after it was enabled. Commit timestamp information is routinely removed during vacuum.


<a id="functions-commit-timestamp"></a>
**Table: Committed Transaction Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_xact_commit_timestamp</code> ( <code>xid</code> ) <code>timestamp with time zone</code></td>
<td>Returns the commit timestamp of a transaction.</td>
<td></td>
</tr>
<tr>
<td><code>pg_xact_commit_timestamp_origin</code> ( <code>xid</code> ) <code>record</code> ( <code>timestamp</code> <code>timestamp with time zone</code>, <code>roident</code> <code>oid</code>)</td>
<td>Returns the commit timestamp and replication origin of a transaction.</td>
<td></td>
</tr>
<tr>
<td><code>pg_last_committed_xact</code> () <code>record</code> ( <code>xid</code> <code>xid</code>, <code>timestamp</code> <code>timestamp with time zone</code>, <code>roident</code> <code>oid</code> )</td>
<td>Returns the transaction ID, commit timestamp and replication origin of the latest committed transaction.</td>
<td></td>
</tr>
</tbody>
</table>


### Control Data Functions { #functions-info-controldata }


 The functions shown in [Control Data Functions](#functions-controldata) print information initialized during `initdb`, such as the catalog version. They also show information about write-ahead logging and checkpoint processing. This information is cluster-wide, not specific to any one database. These functions provide most of the same information, from the same source, as the [app-pgcontroldata](../../reference/postgresql-server-applications/pg_controldata.md#app-pgcontroldata) application.


<a id="functions-controldata"></a>
**Table: Control Data Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_control_checkpoint</code> () <code>record</code></td>
<td>Returns information about current checkpoint state, as shown in <a href="#functions-pg-control-checkpoint"><code>pg_control_checkpoint</code> Output Columns</a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_control_system</code> () <code>record</code></td>
<td>Returns information about current control file state, as shown in <a href="#functions-pg-control-system"><code>pg_control_system</code> Output Columns</a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_control_init</code> () <code>record</code></td>
<td>Returns information about cluster initialization state, as shown in <a href="#functions-pg-control-init"><code>pg_control_init</code> Output Columns</a>.</td>
<td></td>
</tr>
<tr>
<td><code>pg_control_recovery</code> () <code>record</code></td>
<td>Returns information about recovery state, as shown in <a href="#functions-pg-control-recovery"><code>pg_control_recovery</code> Output Columns</a>.</td>
<td></td>
</tr>
</tbody>
</table>


<a id="functions-pg-control-checkpoint"></a>
**Table: `pg_control_checkpoint` Output Columns**

| Column Name | Data Type |
| --- | --- |
| `checkpoint_lsn` | `pg_lsn` |
| `redo_lsn` | `pg_lsn` |
| `redo_wal_file` | `text` |
| `timeline_id` | `integer` |
| `prev_timeline_id` | `integer` |
| `full_page_writes` | `boolean` |
| `next_xid` | `text` |
| `next_oid` | `oid` |
| `next_multixact_id` | `xid` |
| `next_multi_offset` | `xid` |
| `oldest_xid` | `xid` |
| `oldest_xid_dbid` | `oid` |
| `oldest_active_xid` | `xid` |
| `oldest_multi_xid` | `xid` |
| `oldest_multi_dbid` | `oid` |
| `oldest_commit_ts_xid` | `xid` |
| `newest_commit_ts_xid` | `xid` |
| `checkpoint_time` | `timestamp with time zone` |


<a id="functions-pg-control-system"></a>
**Table: `pg_control_system` Output Columns**

| Column Name | Data Type |
| --- | --- |
| `pg_control_version` | `integer` |
| `catalog_version_no` | `integer` |
| `system_identifier` | `bigint` |
| `pg_control_last_modified` | `timestamp with time zone` |


<a id="functions-pg-control-init"></a>
**Table: `pg_control_init` Output Columns**

| Column Name | Data Type |
| --- | --- |
| `max_data_alignment` | `integer` |
| `database_block_size` | `integer` |
| `blocks_per_segment` | `integer` |
| `wal_block_size` | `integer` |
| `bytes_per_wal_segment` | `integer` |
| `max_identifier_length` | `integer` |
| `max_index_columns` | `integer` |
| `max_toast_chunk_size` | `integer` |
| `large_object_chunk_size` | `integer` |
| `float8_pass_by_value` | `boolean` |
| `data_page_checksum_version` | `integer` |
| `default_char_signedness` | `boolean` |


<a id="functions-pg-control-recovery"></a>
**Table: `pg_control_recovery` Output Columns**

| Column Name | Data Type |
| --- | --- |
| `min_recovery_end_lsn` | `pg_lsn` |
| `min_recovery_end_timeline` | `integer` |
| `backup_start_lsn` | `pg_lsn` |
| `backup_end_lsn` | `pg_lsn` |
| `end_of_backup_record_required` | `boolean` |


### Version Information Functions { #functions-info-version }


 The functions shown in [Version Information Functions](#functions-version) print version information.


<a id="functions-version"></a>
**Table: Version Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>version</code> () <code>text</code></td>
<td>Returns a string describing the PostgreSQL server's version. You can also get this information from <a href="../../server-administration/server-configuration/preset-options.md#guc-server-version">server_version</a>, or for a machine-readable version use <a href="../../server-administration/server-configuration/preset-options.md#guc-server-version-num">server_version_num</a>. Software developers should use <code>server_version_num</code> (available since 8.2) or <a href="../../client-interfaces/libpq-c-library/connection-status-functions.md#libpq-PQserverVersion">PQserverVersion</a> instead of parsing the text version.</td>
<td></td>
</tr>
<tr>
<td><code>unicode_version</code> () <code>text</code></td>
<td>Returns a string representing the version of Unicode used by PostgreSQL.</td>
<td></td>
</tr>
<tr>
<td><code>icu_unicode_version</code> () <code>text</code></td>
<td>Returns a string representing the version of Unicode used by ICU, if the server was built with ICU support; otherwise returns <code>NULL</code></td>
<td></td>
</tr>
</tbody>
</table>


### WAL Summarization Information Functions { #functions-info-wal-summary }


 The functions shown in [WAL Summarization Information Functions](#functions-wal-summary) print information about the status of WAL summarization. See [summarize_wal](../../server-administration/server-configuration/write-ahead-log.md#guc-summarize-wal).


<a id="functions-wal-summary"></a>
**Table: WAL Summarization Information Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>pg_available_wal_summaries</code> () <code>setof record</code> ( <code>tli</code> <code>bigint</code>, <code>start_lsn</code> <code>pg_lsn</code>, <code>end_lsn</code> <code>pg_lsn</code> )</td>
<td>Returns information about the WAL summary files present in the data directory, under <code>pg_wal/summaries</code>. One row will be returned per WAL summary file. Each file summarizes WAL on the indicated TLI within the indicated LSN range. This function might be useful to determine whether enough WAL summaries are present on the server to take an incremental backup based on some prior backup whose start LSN is known.</td>
<td></td>
</tr>
<tr>
<td><code>pg_wal_summary_contents</code> ( <code>tli</code> <code>bigint</code>, <code>start_lsn</code> <code>pg_lsn</code>, <code>end_lsn</code> <code>pg_lsn</code> ) <code>setof record</code> ( <code>relfilenode</code> <code>oid</code>, <code>reltablespace</code> <code>oid</code>, <code>reldatabase</code> <code>oid</code>, <code>relforknumber</code> <code>smallint</code>, <code>relblocknumber</code> <code>bigint</code>, <code>is_limit_block</code> <code>boolean</code> )</td>
<td>Returns one information about the contents of a single WAL summary file identified by TLI and starting and ending LSNs. Each row with <code>is_limit_block</code> false indicates that the block identified by the remaining output columns was modified by at least one WAL record within the range of records summarized by this file. Each row with <code>is_limit_block</code> true indicates either that (a) the relation fork was truncated to the length given by <code>relblocknumber</code> within the relevant range of WAL records or (b) that the relation fork was created or dropped within the relevant range of WAL records; in such cases, <code>relblocknumber</code> will be zero.</td>
<td></td>
</tr>
<tr>
<td><code>pg_get_wal_summarizer_state</code> () <code>record</code> ( <code>summarized_tli</code> <code>bigint</code>, <code>summarized_lsn</code> <code>pg_lsn</code>, <code>pending_lsn</code> <code>pg_lsn</code>, <code>summarizer_pid</code> <code>int</code> )</td>
<td>Returns information about the progress of the WAL summarizer. If the WAL summarizer has never run since the instance was started, then <code>summarized_tli</code> and <code>summarized_lsn</code> will be <code>0</code> and <code>0/0</code> respectively; otherwise, they will be the TLI and ending LSN of the last WAL summary file written to disk. If the WAL summarizer is currently running, <code>pending_lsn</code> will be the ending LSN of the last record that it has consumed, which must always be greater than or equal to <code>summarized_lsn</code>; if the WAL summarizer is not running, it will be equal to <code>summarized_lsn</code>. <code>summarizer_pid</code> is the PID of the WAL summarizer process, if it is running, and otherwise NULL.</td>
<td>As a special exception, the WAL summarizer will refuse to generate WAL summary files if run on WAL generated under <code>wal_level=minimal</code>, since such summaries would be unsafe to use as the basis for an incremental backup. In this case, the fields above will continue to advance as if summaries were being generated, but nothing will be written to disk. Once the summarizer reaches WAL generated while <code>wal_level</code> was set to <code>replica</code> or higher, it will resume writing summaries to disk.</td>
</tr>
</tbody>
</table>
