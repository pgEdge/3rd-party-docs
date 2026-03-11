<a id="catalog-pg-proc"></a>

## `pg_proc`


 The catalog `pg_proc` stores information about functions, procedures, aggregate functions, and window functions (collectively also known as routines). See [sql-createfunction](../../reference/sql-commands/create-function.md#sql-createfunction), [sql-createprocedure](../../reference/sql-commands/create-procedure.md#sql-createprocedure), and [User-Defined Functions](../../server-programming/extending-sql/user-defined-functions.md#xfunc) for more information.


 If `prokind` indicates that the entry is for an aggregate function, there should be a matching row in [`pg_aggregate`](pg_aggregate.md#catalog-pg-aggregate).


**Table: `pg_proc` Columns**

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
<td><p><code>proname</code> <code>name</code></p>
<p>Name of the function</p></td>
</tr>
<tr>
<td><p><code>pronamespace</code> <code>oid</code> (references <a href="pg_namespace.md#catalog-pg-namespace"><code>pg_namespace</code></a>.<code>oid</code>)</p>
<p>The OID of the namespace that contains this function</p></td>
</tr>
<tr>
<td><p><code>proowner</code> <code>oid</code> (references <a href="pg_authid.md#catalog-pg-authid"><code>pg_authid</code></a>.<code>oid</code>)</p>
<p>Owner of the function</p></td>
</tr>
<tr>
<td><p><code>prolang</code> <code>oid</code> (references <a href="pg_language.md#catalog-pg-language"><code>pg_language</code></a>.<code>oid</code>)</p>
<p>Implementation language or call interface of this function</p></td>
</tr>
<tr>
<td><p><code>procost</code> <code>float4</code></p>
<p>Estimated execution cost (in units of <a href="../../server-administration/server-configuration/query-planning.md#guc-cpu-operator-cost">cpu_operator_cost</a>); if <code>proretset</code>, this is cost per row returned</p></td>
</tr>
<tr>
<td><p><code>prorows</code> <code>float4</code></p>
<p>Estimated number of result rows (zero if not <code>proretset</code>)</p></td>
</tr>
<tr>
<td><p><code>provariadic</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type of the variadic array parameter's elements, or zero if the function does not have a variadic parameter</p></td>
</tr>
<tr>
<td><p><code>prosupport</code> <code>regproc</code> (references <a href="#catalog-pg-proc"><code>pg_proc</code></a>.<code>oid</code>)</p>
<p>Planner support function for this function (see <a href="../../server-programming/extending-sql/function-optimization-information.md#xfunc-optimization">Function Optimization Information</a>), or zero if none</p></td>
</tr>
<tr>
<td><p><code>prokind</code> <code>char</code></p>
<p><code>f</code> for a normal function, <code>p</code> for a procedure, <code>a</code> for an aggregate function, or <code>w</code> for a window function</p></td>
</tr>
<tr>
<td><p><code>prosecdef</code> <code>bool</code></p>
<p>Function is a security definer (i.e., a “setuid” function)</p></td>
</tr>
<tr>
<td><p><code>proleakproof</code> <code>bool</code></p>
<p>The function has no side effects. No information about the arguments is conveyed except via the return value. Any function that might throw an error depending on the values of its arguments is not leak-proof.</p></td>
</tr>
<tr>
<td><p><code>proisstrict</code> <code>bool</code></p>
<p>Function returns null if any call argument is null. In that case the function won't actually be called at all. Functions that are not “strict” must be prepared to handle null inputs.</p></td>
</tr>
<tr>
<td><p><code>proretset</code> <code>bool</code></p>
<p>Function returns a set (i.e., multiple values of the specified data type)</p></td>
</tr>
<tr>
<td><p><code>provolatile</code> <code>char</code></p>
<p><code>provolatile</code> tells whether the function's result depends only on its input arguments, or is affected by outside factors. It is <code>i</code> for “immutable” functions, which always deliver the same result for the same inputs. It is <code>s</code> for “stable” functions, whose results (for fixed inputs) do not change within a scan. It is <code>v</code> for “volatile” functions, whose results might change at any time. (Use <code>v</code> also for functions with side-effects, so that calls to them cannot get optimized away.)</p></td>
</tr>
<tr>
<td><p><code>proparallel</code> <code>char</code></p>
<p><code>proparallel</code> tells whether the function can be safely run in parallel mode. It is <code>s</code> for functions which are safe to run in parallel mode without restriction. It is <code>r</code> for functions which can be run in parallel mode, but their execution is restricted to the parallel group leader; parallel worker processes cannot invoke these functions. It is <code>u</code> for functions which are unsafe in parallel mode; the presence of such a function forces a serial execution plan.</p></td>
</tr>
<tr>
<td><p><code>pronargs</code> <code>int2</code></p>
<p>Number of input arguments</p></td>
</tr>
<tr>
<td><p><code>pronargdefaults</code> <code>int2</code></p>
<p>Number of arguments that have defaults</p></td>
</tr>
<tr>
<td><p><code>prorettype</code> <code>oid</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>Data type of the return value</p></td>
</tr>
<tr>
<td><p><code>proargtypes</code> <code>oidvector</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>An array of the data types of the function arguments. This includes only input arguments (including <code>INOUT</code> and <code>VARIADIC</code> arguments), and thus represents the call signature of the function.</p></td>
</tr>
<tr>
<td><p><code>proallargtypes</code> <code>oid[]</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>An array of the data types of the function arguments. This includes all arguments (including <code>OUT</code> and <code>INOUT</code> arguments); however, if all the arguments are <code>IN</code> arguments, this field will be null. Note that subscripting is 1-based, whereas for historical reasons <code>proargtypes</code> is subscripted from 0.</p></td>
</tr>
<tr>
<td><p><code>proargmodes</code> <code>char[]</code></p>
<p>An array of the modes of the function arguments, encoded as <code>i</code> for <code>IN</code> arguments, <code>o</code> for <code>OUT</code> arguments, <code>b</code> for <code>INOUT</code> arguments, <code>v</code> for <code>VARIADIC</code> arguments, <code>t</code> for <code>TABLE</code> arguments. If all the arguments are <code>IN</code> arguments, this field will be null. Note that subscripts correspond to positions of <code>proallargtypes</code> not <code>proargtypes</code>.</p></td>
</tr>
<tr>
<td><p><code>proargnames</code> <code>text[]</code></p>
<p>An array of the names of the function arguments. Arguments without a name are set to empty strings in the array. If none of the arguments have a name, this field will be null. Note that subscripts correspond to positions of <code>proallargtypes</code> not <code>proargtypes</code>.</p></td>
</tr>
<tr>
<td><p><code>proargdefaults</code> <code>pg_node_tree</code></p>
<p>Expression trees (in <code>nodeToString()</code> representation) for default values. This is a list with <code>pronargdefaults</code> elements, corresponding to the last <em>N</em> <em>input</em> arguments (i.e., the last <em>N</em> <code>proargtypes</code> positions). If none of the arguments have defaults, this field will be null.</p></td>
</tr>
<tr>
<td><p><code>protrftypes</code> <code>oid[]</code> (references <a href="pg_type.md#catalog-pg-type"><code>pg_type</code></a>.<code>oid</code>)</p>
<p>An array of the argument/result data type(s) for which to apply transforms (from the function's <code>TRANSFORM</code> clause). Null if none.</p></td>
</tr>
<tr>
<td><p><code>prosrc</code> <code>text</code></p>
<p>This tells the function handler how to invoke the function. It might be the actual source code of the function for interpreted languages, a link symbol, a file name, or just about anything else, depending on the implementation language/call convention.</p></td>
</tr>
<tr>
<td><p><code>probin</code> <code>text</code></p>
<p>Additional information about how to invoke the function. Again, the interpretation is language-specific.</p></td>
</tr>
<tr>
<td><p><code>prosqlbody</code> <code>pg_node_tree</code></p>
<p>Pre-parsed SQL function body. This is used for SQL-language functions when the body is given in SQL-standard notation rather than as a string literal. It's null in other cases.</p></td>
</tr>
<tr>
<td><p><code>proconfig</code> <code>text[]</code></p>
<p>Function's local settings for run-time configuration variables</p></td>
</tr>
<tr>
<td><p><code>proacl</code> <code>aclitem[]</code></p>
<p>Access privileges; see <a href="../../the-sql-language/data-definition/privileges.md#ddl-priv">Privileges</a> for details</p></td>
</tr>
</tbody>
</table>


 For compiled functions, both built-in and dynamically loaded, `prosrc` contains the function's C-language name (link symbol). For SQL-language functions, `prosrc` contains the function's source text if that is specified as a string literal; but if the function body is specified in SQL-standard style, `prosrc` is unused (typically it's an empty string) and `prosqlbody` contains the pre-parsed definition. For all other currently-known language types, `prosrc` contains the function's source text. `probin` is null except for dynamically-loaded C functions, for which it gives the name of the shared library file containing the function.
