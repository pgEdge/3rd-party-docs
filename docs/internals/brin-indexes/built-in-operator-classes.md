## Built-in Operator Classes { #brin-builtin-opclasses }


 The core PostgreSQL distribution includes the BRIN operator classes shown in [Built-in BRIN Operator Classes](#brin-builtin-opclasses-table).


 The *minmax* operator classes store the minimum and the maximum values appearing in the indexed column within the range. The *inclusion* operator classes store a value which includes the values in the indexed column within the range. The *bloom* operator classes build a Bloom filter for all values in the range. The *minmax-multi* operator classes store multiple minimum and maximum values, representing values appearing in the indexed column within the range.


<a id="brin-builtin-opclasses-table"></a>
**Table: Built-in BRIN Operator Classes**

<table>
<thead>
<tr>
<th>Name</th>
<th>Indexable Operators</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan="4"><code>bit_minmax_ops</code></td>
<td><code>= (bit,bit)</code></td>
</tr>
<tr>
<td><code>&lt; (bit,bit)</code></td>
</tr>
<tr>
<td><code>&gt; (bit,bit)</code></td>
</tr>
<tr>
<td><code>&lt;= (bit,bit)</code></td>
</tr>
<tr>
<td><code>&gt;= (bit,bit)</code></td>
</tr>
<tr>
<td rowspan="12"><code>box_inclusion_ops</code></td>
<td><code>@&gt; (box,point)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>&lt;@ (box,box)</code></td>
</tr>
<tr>
<td><code>@&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>~= (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (box,box)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;| (box,box)</code></td>
</tr>
<tr>
<td><code>&amp;&lt;| (box,box)</code></td>
</tr>
<tr>
<td><code>|&amp;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>|&gt;&gt; (box,box)</code></td>
</tr>
<tr>
<td><code>bpchar_bloom_ops</code></td>
<td><code>= (character,character)</code></td>
</tr>
<tr>
<td rowspan="4"><code>bpchar_minmax_ops</code></td>
<td><code>= (character,character)</code></td>
</tr>
<tr>
<td><code>&lt; (character,character)</code></td>
</tr>
<tr>
<td><code>&lt;= (character,character)</code></td>
</tr>
<tr>
<td><code>&gt; (character,character)</code></td>
</tr>
<tr>
<td><code>&gt;= (character,character)</code></td>
</tr>
<tr>
<td><code>bytea_bloom_ops</code></td>
<td><code>= (bytea,bytea)</code></td>
</tr>
<tr>
<td rowspan="4"><code>bytea_minmax_ops</code></td>
<td><code>= (bytea,bytea)</code></td>
</tr>
<tr>
<td><code>&lt; (bytea,bytea)</code></td>
</tr>
<tr>
<td><code>&lt;= (bytea,bytea)</code></td>
</tr>
<tr>
<td><code>&gt; (bytea,bytea)</code></td>
</tr>
<tr>
<td><code>&gt;= (bytea,bytea)</code></td>
</tr>
<tr>
<td><code>char_bloom_ops</code></td>
<td><code>= ("char","char")</code></td>
</tr>
<tr>
<td rowspan="4"><code>char_minmax_ops</code></td>
<td><code>= ("char","char")</code></td>
</tr>
<tr>
<td><code>&lt; ("char","char")</code></td>
</tr>
<tr>
<td><code>&lt;= ("char","char")</code></td>
</tr>
<tr>
<td><code>&gt; ("char","char")</code></td>
</tr>
<tr>
<td><code>&gt;= ("char","char")</code></td>
</tr>
<tr>
<td><code>date_bloom_ops</code></td>
<td><code>= (date,date)</code></td>
</tr>
<tr>
<td rowspan="4"><code>date_minmax_ops</code></td>
<td><code>= (date,date)</code></td>
</tr>
<tr>
<td><code>&lt; (date,date)</code></td>
</tr>
<tr>
<td><code>&lt;= (date,date)</code></td>
</tr>
<tr>
<td><code>&gt; (date,date)</code></td>
</tr>
<tr>
<td><code>&gt;= (date,date)</code></td>
</tr>
<tr>
<td rowspan="4"><code>date_minmax_multi_ops</code></td>
<td><code>= (date,date)</code></td>
</tr>
<tr>
<td><code>&lt; (date,date)</code></td>
</tr>
<tr>
<td><code>&lt;= (date,date)</code></td>
</tr>
<tr>
<td><code>&gt; (date,date)</code></td>
</tr>
<tr>
<td><code>&gt;= (date,date)</code></td>
</tr>
<tr>
<td><code>float4_bloom_ops</code></td>
<td><code>= (float4,float4)</code></td>
</tr>
<tr>
<td rowspan="4"><code>float4_minmax_ops</code></td>
<td><code>= (float4,float4)</code></td>
</tr>
<tr>
<td><code>&lt; (float4,float4)</code></td>
</tr>
<tr>
<td><code>&gt; (float4,float4)</code></td>
</tr>
<tr>
<td><code>&lt;= (float4,float4)</code></td>
</tr>
<tr>
<td><code>&gt;= (float4,float4)</code></td>
</tr>
<tr>
<td rowspan="4"><code>float4_minmax_multi_ops</code></td>
<td><code>= (float4,float4)</code></td>
</tr>
<tr>
<td><code>&lt; (float4,float4)</code></td>
</tr>
<tr>
<td><code>&gt; (float4,float4)</code></td>
</tr>
<tr>
<td><code>&lt;= (float4,float4)</code></td>
</tr>
<tr>
<td><code>&gt;= (float4,float4)</code></td>
</tr>
<tr>
<td><code>float8_bloom_ops</code></td>
<td><code>= (float8,float8)</code></td>
</tr>
<tr>
<td rowspan="4"><code>float8_minmax_ops</code></td>
<td><code>= (float8,float8)</code></td>
</tr>
<tr>
<td><code>&lt; (float8,float8)</code></td>
</tr>
<tr>
<td><code>&lt;= (float8,float8)</code></td>
</tr>
<tr>
<td><code>&gt; (float8,float8)</code></td>
</tr>
<tr>
<td><code>&gt;= (float8,float8)</code></td>
</tr>
<tr>
<td rowspan="4"><code>float8_minmax_multi_ops</code></td>
<td><code>= (float8,float8)</code></td>
</tr>
<tr>
<td><code>&lt; (float8,float8)</code></td>
</tr>
<tr>
<td><code>&lt;= (float8,float8)</code></td>
</tr>
<tr>
<td><code>&gt; (float8,float8)</code></td>
</tr>
<tr>
<td><code>&gt;= (float8,float8)</code></td>
</tr>
<tr>
<td rowspan="5"><code>inet_inclusion_ops</code></td>
<td><code>&lt;&lt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt;&lt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;&gt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (inet,inet)</code></td>
</tr>
<tr>
<td><code>inet_bloom_ops</code></td>
<td><code>= (inet,inet)</code></td>
</tr>
<tr>
<td rowspan="4"><code>inet_minmax_ops</code></td>
<td><code>= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;= (inet,inet)</code></td>
</tr>
<tr>
<td rowspan="4"><code>inet_minmax_multi_ops</code></td>
<td><code>= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&lt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt; (inet,inet)</code></td>
</tr>
<tr>
<td><code>&gt;= (inet,inet)</code></td>
</tr>
<tr>
<td><code>int2_bloom_ops</code></td>
<td><code>= (int2,int2)</code></td>
</tr>
<tr>
<td rowspan="4"><code>int2_minmax_ops</code></td>
<td><code>= (int2,int2)</code></td>
</tr>
<tr>
<td><code>&lt; (int2,int2)</code></td>
</tr>
<tr>
<td><code>&gt; (int2,int2)</code></td>
</tr>
<tr>
<td><code>&lt;= (int2,int2)</code></td>
</tr>
<tr>
<td><code>&gt;= (int2,int2)</code></td>
</tr>
<tr>
<td rowspan="4"><code>int2_minmax_multi_ops</code></td>
<td><code>= (int2,int2)</code></td>
</tr>
<tr>
<td><code>&lt; (int2,int2)</code></td>
</tr>
<tr>
<td><code>&gt; (int2,int2)</code></td>
</tr>
<tr>
<td><code>&lt;= (int2,int2)</code></td>
</tr>
<tr>
<td><code>&gt;= (int2,int2)</code></td>
</tr>
<tr>
<td><code>int4_bloom_ops</code></td>
<td><code>= (int4,int4)</code></td>
</tr>
<tr>
<td rowspan="4"><code>int4_minmax_ops</code></td>
<td><code>= (int4,int4)</code></td>
</tr>
<tr>
<td><code>&lt; (int4,int4)</code></td>
</tr>
<tr>
<td><code>&gt; (int4,int4)</code></td>
</tr>
<tr>
<td><code>&lt;= (int4,int4)</code></td>
</tr>
<tr>
<td><code>&gt;= (int4,int4)</code></td>
</tr>
<tr>
<td rowspan="4"><code>int4_minmax_multi_ops</code></td>
<td><code>= (int4,int4)</code></td>
</tr>
<tr>
<td><code>&lt; (int4,int4)</code></td>
</tr>
<tr>
<td><code>&gt; (int4,int4)</code></td>
</tr>
<tr>
<td><code>&lt;= (int4,int4)</code></td>
</tr>
<tr>
<td><code>&gt;= (int4,int4)</code></td>
</tr>
<tr>
<td><code>int8_bloom_ops</code></td>
<td><code>= (bigint,bigint)</code></td>
</tr>
<tr>
<td rowspan="4"><code>int8_minmax_ops</code></td>
<td><code>= (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&lt; (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&gt; (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&lt;= (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&gt;= (bigint,bigint)</code></td>
</tr>
<tr>
<td rowspan="4"><code>int8_minmax_multi_ops</code></td>
<td><code>= (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&lt; (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&gt; (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&lt;= (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>&gt;= (bigint,bigint)</code></td>
</tr>
<tr>
<td><code>interval_bloom_ops</code></td>
<td><code>= (interval,interval)</code></td>
</tr>
<tr>
<td rowspan="4"><code>interval_minmax_ops</code></td>
<td><code>= (interval,interval)</code></td>
</tr>
<tr>
<td><code>&lt; (interval,interval)</code></td>
</tr>
<tr>
<td><code>&lt;= (interval,interval)</code></td>
</tr>
<tr>
<td><code>&gt; (interval,interval)</code></td>
</tr>
<tr>
<td><code>&gt;= (interval,interval)</code></td>
</tr>
<tr>
<td rowspan="4"><code>interval_minmax_multi_ops</code></td>
<td><code>= (interval,interval)</code></td>
</tr>
<tr>
<td><code>&lt; (interval,interval)</code></td>
</tr>
<tr>
<td><code>&lt;= (interval,interval)</code></td>
</tr>
<tr>
<td><code>&gt; (interval,interval)</code></td>
</tr>
<tr>
<td><code>&gt;= (interval,interval)</code></td>
</tr>
<tr>
<td><code>macaddr_bloom_ops</code></td>
<td><code>= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td rowspan="4"><code>macaddr_minmax_ops</code></td>
<td><code>= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&lt; (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&lt;= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&gt; (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&gt;= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td rowspan="4"><code>macaddr_minmax_multi_ops</code></td>
<td><code>= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&lt; (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&lt;= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&gt; (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>&gt;= (macaddr,macaddr)</code></td>
</tr>
<tr>
<td><code>macaddr8_bloom_ops</code></td>
<td><code>= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td rowspan="4"><code>macaddr8_minmax_ops</code></td>
<td><code>= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&lt; (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&lt;= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&gt; (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&gt;= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td rowspan="4"><code>macaddr8_minmax_multi_ops</code></td>
<td><code>= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&lt; (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&lt;= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&gt; (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>&gt;= (macaddr8,macaddr8)</code></td>
</tr>
<tr>
<td><code>name_bloom_ops</code></td>
<td><code>= (name,name)</code></td>
</tr>
<tr>
<td rowspan="4"><code>name_minmax_ops</code></td>
<td><code>= (name,name)</code></td>
</tr>
<tr>
<td><code>&lt; (name,name)</code></td>
</tr>
<tr>
<td><code>&lt;= (name,name)</code></td>
</tr>
<tr>
<td><code>&gt; (name,name)</code></td>
</tr>
<tr>
<td><code>&gt;= (name,name)</code></td>
</tr>
<tr>
<td><code>numeric_bloom_ops</code></td>
<td><code>= (numeric,numeric)</code></td>
</tr>
<tr>
<td rowspan="4"><code>numeric_minmax_ops</code></td>
<td><code>= (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&lt; (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&lt;= (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&gt; (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&gt;= (numeric,numeric)</code></td>
</tr>
<tr>
<td rowspan="4"><code>numeric_minmax_multi_ops</code></td>
<td><code>= (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&lt; (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&lt;= (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&gt; (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>&gt;= (numeric,numeric)</code></td>
</tr>
<tr>
<td><code>oid_bloom_ops</code></td>
<td><code>= (oid,oid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>oid_minmax_ops</code></td>
<td><code>= (oid,oid)</code></td>
</tr>
<tr>
<td><code>&lt; (oid,oid)</code></td>
</tr>
<tr>
<td><code>&gt; (oid,oid)</code></td>
</tr>
<tr>
<td><code>&lt;= (oid,oid)</code></td>
</tr>
<tr>
<td><code>&gt;= (oid,oid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>oid_minmax_multi_ops</code></td>
<td><code>= (oid,oid)</code></td>
</tr>
<tr>
<td><code>&lt; (oid,oid)</code></td>
</tr>
<tr>
<td><code>&gt; (oid,oid)</code></td>
</tr>
<tr>
<td><code>&lt;= (oid,oid)</code></td>
</tr>
<tr>
<td><code>&gt;= (oid,oid)</code></td>
</tr>
<tr>
<td><code>pg_lsn_bloom_ops</code></td>
<td><code>= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td rowspan="4"><code>pg_lsn_minmax_ops</code></td>
<td><code>= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&lt; (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&gt; (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&lt;= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&gt;= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td rowspan="4"><code>pg_lsn_minmax_multi_ops</code></td>
<td><code>= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&lt; (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&gt; (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&lt;= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td><code>&gt;= (pg_lsn,pg_lsn)</code></td>
</tr>
<tr>
<td rowspan="13"><code>range_inclusion_ops</code></td>
<td><code>= (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&lt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;= (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&gt;= (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&amp; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange,anyelement)</code></td>
</tr>
<tr>
<td><code>@&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;@ (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&lt;&lt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&gt;&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&lt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>&amp;&gt; (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>-|- (anyrange,anyrange)</code></td>
</tr>
<tr>
<td><code>text_bloom_ops</code></td>
<td><code>= (text,text)</code></td>
</tr>
<tr>
<td rowspan="4"><code>text_minmax_ops</code></td>
<td><code>= (text,text)</code></td>
</tr>
<tr>
<td><code>&lt; (text,text)</code></td>
</tr>
<tr>
<td><code>&lt;= (text,text)</code></td>
</tr>
<tr>
<td><code>&gt; (text,text)</code></td>
</tr>
<tr>
<td><code>&gt;= (text,text)</code></td>
</tr>
<tr>
<td><code>tid_bloom_ops</code></td>
<td><code>= (tid,tid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>tid_minmax_ops</code></td>
<td><code>= (tid,tid)</code></td>
</tr>
<tr>
<td><code>&lt; (tid,tid)</code></td>
</tr>
<tr>
<td><code>&gt; (tid,tid)</code></td>
</tr>
<tr>
<td><code>&lt;= (tid,tid)</code></td>
</tr>
<tr>
<td><code>&gt;= (tid,tid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>tid_minmax_multi_ops</code></td>
<td><code>= (tid,tid)</code></td>
</tr>
<tr>
<td><code>&lt; (tid,tid)</code></td>
</tr>
<tr>
<td><code>&gt; (tid,tid)</code></td>
</tr>
<tr>
<td><code>&lt;= (tid,tid)</code></td>
</tr>
<tr>
<td><code>&gt;= (tid,tid)</code></td>
</tr>
<tr>
<td><code>timestamp_bloom_ops</code></td>
<td><code>= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td rowspan="4"><code>timestamp_minmax_ops</code></td>
<td><code>= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&lt; (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&lt;= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&gt; (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&gt;= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td rowspan="4"><code>timestamp_minmax_multi_ops</code></td>
<td><code>= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&lt; (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&lt;= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&gt; (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>&gt;= (timestamp,timestamp)</code></td>
</tr>
<tr>
<td><code>timestamptz_bloom_ops</code></td>
<td><code>= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td rowspan="4"><code>timestamptz_minmax_ops</code></td>
<td><code>= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&lt; (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&lt;= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&gt; (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&gt;= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td rowspan="4"><code>timestamptz_minmax_multi_ops</code></td>
<td><code>= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&lt; (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&lt;= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&gt; (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>&gt;= (timestamptz,timestamptz)</code></td>
</tr>
<tr>
<td><code>time_bloom_ops</code></td>
<td><code>= (time,time)</code></td>
</tr>
<tr>
<td rowspan="4"><code>time_minmax_ops</code></td>
<td><code>= (time,time)</code></td>
</tr>
<tr>
<td><code>&lt; (time,time)</code></td>
</tr>
<tr>
<td><code>&lt;= (time,time)</code></td>
</tr>
<tr>
<td><code>&gt; (time,time)</code></td>
</tr>
<tr>
<td><code>&gt;= (time,time)</code></td>
</tr>
<tr>
<td rowspan="4"><code>time_minmax_multi_ops</code></td>
<td><code>= (time,time)</code></td>
</tr>
<tr>
<td><code>&lt; (time,time)</code></td>
</tr>
<tr>
<td><code>&lt;= (time,time)</code></td>
</tr>
<tr>
<td><code>&gt; (time,time)</code></td>
</tr>
<tr>
<td><code>&gt;= (time,time)</code></td>
</tr>
<tr>
<td><code>timetz_bloom_ops</code></td>
<td><code>= (timetz,timetz)</code></td>
</tr>
<tr>
<td rowspan="4"><code>timetz_minmax_ops</code></td>
<td><code>= (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&lt; (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&lt;= (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&gt; (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&gt;= (timetz,timetz)</code></td>
</tr>
<tr>
<td rowspan="4"><code>timetz_minmax_multi_ops</code></td>
<td><code>= (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&lt; (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&lt;= (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&gt; (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>&gt;= (timetz,timetz)</code></td>
</tr>
<tr>
<td><code>uuid_bloom_ops</code></td>
<td><code>= (uuid,uuid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>uuid_minmax_ops</code></td>
<td><code>= (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&lt; (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&gt; (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&lt;= (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&gt;= (uuid,uuid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>uuid_minmax_multi_ops</code></td>
<td><code>= (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&lt; (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&gt; (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&lt;= (uuid,uuid)</code></td>
</tr>
<tr>
<td><code>&gt;= (uuid,uuid)</code></td>
</tr>
<tr>
<td rowspan="4"><code>varbit_minmax_ops</code></td>
<td><code>= (varbit,varbit)</code></td>
</tr>
<tr>
<td><code>&lt; (varbit,varbit)</code></td>
</tr>
<tr>
<td><code>&gt; (varbit,varbit)</code></td>
</tr>
<tr>
<td><code>&lt;= (varbit,varbit)</code></td>
</tr>
<tr>
<td><code>&gt;= (varbit,varbit)</code></td>
</tr>
</tbody>
</table>


### Operator Class Parameters { #brin-builtin-opclasses--parameters }


 Some of the built-in operator classes allow specifying parameters affecting behavior of the operator class. Each operator class has its own set of allowed parameters. Only the `bloom` and `minmax-multi` operator classes allow specifying parameters:


 bloom operator classes accept these parameters:


`n_distinct_per_range`
:   Defines the estimated number of distinct non-null values in the block range, used by BRIN bloom indexes for sizing of the Bloom filter. It behaves similarly to `n_distinct` option for [sql-altertable](../../reference/sql-commands/alter-table.md#sql-altertable). When set to a positive value, each block range is assumed to contain this number of distinct non-null values. When set to a negative value, which must be greater than or equal to -1, the number of distinct non-null values is assumed to grow linearly with the maximum possible number of tuples in the block range (about 290 rows per block). The default value is `-0.1`, and the minimum number of distinct non-null values is `16`.

`false_positive_rate`
:   Defines the desired false positive rate used by BRIN bloom indexes for sizing of the Bloom filter. The values must be between 0.0001 and 0.25. The default value is 0.01, which is 1% false positive rate.


 minmax-multi operator classes accept these parameters:


`values_per_range`
:   Defines the maximum number of values stored by BRIN minmax indexes to summarize a block range. Each value may represent either a point, or a boundary of an interval. Values must be between 8 and 256, and the default value is 32.
