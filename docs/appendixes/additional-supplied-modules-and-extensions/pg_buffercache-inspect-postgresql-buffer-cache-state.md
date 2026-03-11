<a id="pgbuffercache"></a>

## pg_buffercache — inspect PostgreSQL buffer cache state


 The `pg_buffercache` module provides a means for examining what's happening in the shared buffer cache in real time. It also offers a low-level way to evict data from it, for testing purposes.


 This module provides the `pg_buffercache_pages()` function (wrapped in the `pg_buffercache` view), the `pg_buffercache_numa_pages()` function (wrapped in the `pg_buffercache_numa` view), the `pg_buffercache_summary()` function, the `pg_buffercache_usage_counts()` function, the `pg_buffercache_evict()` function, the `pg_buffercache_evict_relation()` function and the `pg_buffercache_evict_all()` function.


 The `pg_buffercache_pages()` function returns a set of records, each row describing the state of one shared buffer entry. The `pg_buffercache` view wraps the function for convenient use.


 The `pg_buffercache_numa_pages()` function provides NUMA node mappings for shared buffer entries. This information is not part of `pg_buffercache_pages()` itself, as it is much slower to retrieve. The `pg_buffercache_numa` view wraps the function for convenient use.


 The `pg_buffercache_summary()` function returns a single row summarizing the state of the shared buffer cache.


 The `pg_buffercache_usage_counts()` function returns a set of records, each row describing the number of buffers with a given usage count.


 By default, use of the above functions is restricted to superusers and roles with privileges of the `pg_monitor` role. Access may be granted to others using `GRANT`.


 The `pg_buffercache_evict()` function allows a block to be evicted from the buffer pool given a buffer identifier. Use of this function is restricted to superusers only.


 The `pg_buffercache_evict_relation()` function allows all unpinned shared buffers in the relation to be evicted from the buffer pool given a relation identifier. Use of this function is restricted to superusers only.


 The `pg_buffercache_evict_all()` function allows all unpinned shared buffers to be evicted in the buffer pool. Use of this function is restricted to superusers only.
 <a id="pgbuffercache-pg-buffercache"></a>

### The `pg_buffercache` View


 The definitions of the columns exposed by the view are shown in [`pg_buffercache` Columns](#pgbuffercache-columns).
 <a id="pgbuffercache-columns"></a>

**Table: `pg_buffercache` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>bufferid</code> <code>integer</code></p>
<p>ID, in the range 1..<code>shared_buffers</code></p></td>
</tr>
<tr>
<td><p><code>relfilenode</code> <code>oid</code> (references <a href="../../internals/system-catalogs/pg_class.md#catalog-pg-class"><code>pg_class</code></a>.<code>relfilenode</code>)</p>
<p>Filenode number of the relation</p></td>
</tr>
<tr>
<td><p><code>reltablespace</code> <code>oid</code> (references <a href="../../internals/system-catalogs/pg_tablespace.md#catalog-pg-tablespace"><code>pg_tablespace</code></a>.<code>oid</code>)</p>
<p>Tablespace OID of the relation</p></td>
</tr>
<tr>
<td><p><code>reldatabase</code> <code>oid</code> (references <a href="../../internals/system-catalogs/pg_database.md#catalog-pg-database"><code>pg_database</code></a>.<code>oid</code>)</p>
<p>Database OID of the relation</p></td>
</tr>
<tr>
<td><p><code>relforknumber</code> <code>smallint</code></p>
<p>Fork number within the relation; see <code>common/relpath.h</code></p></td>
</tr>
<tr>
<td><p><code>relblocknumber</code> <code>bigint</code></p>
<p>Page number within the relation</p></td>
</tr>
<tr>
<td><p><code>isdirty</code> <code>boolean</code></p>
<p>Is the page dirty?</p></td>
</tr>
<tr>
<td><p><code>usagecount</code> <code>smallint</code></p>
<p>Clock-sweep access count</p></td>
</tr>
<tr>
<td><p><code>pinning_backends</code> <code>integer</code></p>
<p>Number of backends pinning this buffer</p></td>
</tr>
</tbody>
</table>


 There is one row for each buffer in the shared cache. Unused buffers are shown with all fields null except `bufferid`. Shared system catalogs are shown as belonging to database zero.


 Because the cache is shared by all the databases, there will normally be pages from relations not belonging to the current database. This means that there may not be matching join rows in `pg_class` for some rows, or that there could even be incorrect joins. If you are trying to join against `pg_class`, it's a good idea to restrict the join to rows having `reldatabase` equal to the current database's OID or zero.


 Since buffer manager locks are not taken to copy the buffer state data that the view will display, accessing `pg_buffercache` view has less impact on normal buffer activity but it doesn't provide a consistent set of results across all buffers. However, we ensure that the information of each buffer is self-consistent.
  <a id="pgbuffercache-pg-buffercache-numa"></a>

### The `pg_buffercache_numa` View


 The definitions of the columns exposed by the view are shown in [`pg_buffercache_numa` Columns](#pgbuffercache-numa-columns).
 <a id="pgbuffercache-numa-columns"></a>

**Table: `pg_buffercache_numa` Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>bufferid</code> <code>integer</code></p>
<p>ID, in the range 1..<code>shared_buffers</code></p></td>
</tr>
<tr>
<td><p><code>os_page_num</code> <code>bigint</code></p>
<p>number of OS memory page for this buffer</p></td>
</tr>
<tr>
<td><p><code>numa_node</code> <code>int</code></p>
<p>ID of NUMA node</p></td>
</tr>
</tbody>
</table>


 As NUMA node ID inquiry for each page requires memory pages to be paged-in, the first execution of this function can take a noticeable amount of time. In all the cases (first execution or not), retrieving this information is costly and querying the view at a high frequency is not recommended.


!!! warning

    When determining the NUMA node, the view touches all memory pages for the shared memory segment. This will force allocation of the shared memory, if it wasn't allocated already, and the memory may get allocated in a single NUMA node (depending on system configuration).
  <a id="pgbuffercache-summary"></a>

### The `pg_buffercache_summary()` Function


 The definitions of the columns exposed by the function are shown in [`pg_buffercache_summary()` Output Columns](#pgbuffercache-summary-columns).
 <a id="pgbuffercache-summary-columns"></a>

**Table: `pg_buffercache_summary()` Output Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>buffers_used</code> <code>int4</code></p>
<p>Number of used shared buffers</p></td>
</tr>
<tr>
<td><p><code>buffers_unused</code> <code>int4</code></p>
<p>Number of unused shared buffers</p></td>
</tr>
<tr>
<td><p><code>buffers_dirty</code> <code>int4</code></p>
<p>Number of dirty shared buffers</p></td>
</tr>
<tr>
<td><p><code>buffers_pinned</code> <code>int4</code></p>
<p>Number of pinned shared buffers</p></td>
</tr>
<tr>
<td><p><code>usagecount_avg</code> <code>float8</code></p>
<p>Average usage count of used shared buffers</p></td>
</tr>
</tbody>
</table>


 The `pg_buffercache_summary()` function returns a single row summarizing the state of all shared buffers. Similar and more detailed information is provided by the `pg_buffercache` view, but `pg_buffercache_summary()` is significantly cheaper.


 Like the `pg_buffercache` view, `pg_buffercache_summary()` does not acquire buffer manager locks. Therefore concurrent activity can lead to minor inaccuracies in the result.
  <a id="pgbuffercache-usage-counts"></a>

### The `pg_buffercache_usage_counts()` Function


 The definitions of the columns exposed by the function are shown in [`pg_buffercache_usage_counts()` Output Columns](#pgbuffercache_usage_counts-columns).
 <a id="pgbuffercache_usage_counts-columns"></a>

**Table: `pg_buffercache_usage_counts()` Output Columns**

<table>
<thead>
<tr>
<th><p>Column Type</p>
<p>Description</p></th>
</tr>
</thead>
<tbody>
<tr>
<td><p><code>usage_count</code> <code>int4</code></p>
<p>A possible buffer usage count</p></td>
</tr>
<tr>
<td><p><code>buffers</code> <code>int4</code></p>
<p>Number of buffers with the usage count</p></td>
</tr>
<tr>
<td><p><code>dirty</code> <code>int4</code></p>
<p>Number of dirty buffers with the usage count</p></td>
</tr>
<tr>
<td><p><code>pinned</code> <code>int4</code></p>
<p>Number of pinned buffers with the usage count</p></td>
</tr>
</tbody>
</table>


 The `pg_buffercache_usage_counts()` function returns a set of rows summarizing the states of all shared buffers, aggregated over the possible usage count values. Similar and more detailed information is provided by the `pg_buffercache` view, but `pg_buffercache_usage_counts()` is significantly cheaper.


 Like the `pg_buffercache` view, `pg_buffercache_usage_counts()` does not acquire buffer manager locks. Therefore concurrent activity can lead to minor inaccuracies in the result.
  <a id="pgbuffercache-pg-buffercache-evict"></a>

### The `pg_buffercache_evict()` Function


 The `pg_buffercache_evict()` function takes a buffer identifier, as shown in the `bufferid` column of the `pg_buffercache` view. It returns information about whether the buffer was evicted and flushed. The buffer_evicted column is true on success, and false if the buffer wasn't valid, if it couldn't be evicted because it was pinned, or if it became dirty again after an attempt to write it out. The buffer_flushed column is true if the buffer was flushed. This does not necessarily mean that buffer was flushed by us, it might be flushed by someone else. The result is immediately out of date upon return, as the buffer might become valid again at any time due to concurrent activity. The function is intended for developer testing only.
  <a id="pgbuffercache-pg-buffercache-evict-relation"></a>

### The `pg_buffercache_evict_relation()` Function


 The `pg_buffercache_evict_relation()` function is very similar to the `pg_buffercache_evict()` function. The difference is that the `pg_buffercache_evict_relation()` takes a relation identifier instead of buffer identifier. It tries to evict all buffers for all forks in that relation. It returns the number of evicted buffers, flushed buffers and the number of buffers that could not be evicted. Flushed buffers haven't necessarily been flushed by us, they might have been flushed by someone else. The result is immediately out of date upon return, as buffers might immediately be read back in due to concurrent activity. The function is intended for developer testing only.
  <a id="pgbuffercache-pg-buffercache-evict-all"></a>

### The `pg_buffercache_evict_all()` Function


 The `pg_buffercache_evict_all()` function is very similar to the `pg_buffercache_evict()` function. The difference is, the `pg_buffercache_evict_all()` function does not take an argument; instead it tries to evict all buffers in the buffer pool. It returns the number of evicted buffers, flushed buffers and the number of buffers that could not be evicted. Flushed buffers haven't necessarily been flushed by us, they might have been flushed by someone else. The result is immediately out of date upon return, as buffers might immediately be read back in due to concurrent activity. The function is intended for developer testing only.
  <a id="pgbuffercache-sample-output"></a>

### Sample Output


```

regression=# SELECT n.nspname, c.relname, count(*) AS buffers
             FROM pg_buffercache b JOIN pg_class c
             ON b.relfilenode = pg_relation_filenode(c.oid) AND
                b.reldatabase IN (0, (SELECT oid FROM pg_database
                                      WHERE datname = current_database()))
             JOIN pg_namespace n ON n.oid = c.relnamespace
             GROUP BY n.nspname, c.relname
             ORDER BY 3 DESC
             LIMIT 10;

  nspname   |        relname         | buffers
------------+------------------------+---------
 public     | delete_test_table      |     593
 public     | delete_test_table_pkey |     494
 pg_catalog | pg_attribute           |     472
 public     | quad_poly_tbl          |     353
 public     | tenk2                  |     349
 public     | tenk1                  |     349
 public     | gin_test_idx           |     306
 pg_catalog | pg_largeobject         |     206
 public     | gin_test_tbl           |     188
 public     | spgist_text_tbl        |     182
(10 rows)


regression=# SELECT * FROM pg_buffercache_summary();
 buffers_used | buffers_unused | buffers_dirty | buffers_pinned | usagecount_avg
--------------+----------------+---------------+----------------+----------------
          248 |        2096904 |            39 |              0 |       3.141129
(1 row)


regression=# SELECT * FROM pg_buffercache_usage_counts();
 usage_count | buffers | dirty | pinned
-------------+---------+-------+--------
           0 |   14650 |     0 |      0
           1 |    1436 |   671 |      0
           2 |     102 |    88 |      0
           3 |      23 |    21 |      0
           4 |       9 |     7 |      0
           5 |     164 |   106 |      0
(6 rows)
```
  <a id="pgbuffercache-authors"></a>

### Authors


 Mark Kirkwood [markir@paradise.net.nz](mailto:markir@paradise.net.nz)


 Design suggestions: Neil Conway [neilc@samurai.com](mailto:neilc@samurai.com)


 Debugging advice: Tom Lane [tgl@sss.pgh.pa.us](mailto:tgl@sss.pgh.pa.us)
