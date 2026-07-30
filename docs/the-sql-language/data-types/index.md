<a id="datatype"></a>

# Data Types

 PostgreSQL has a rich set of native data types available to users. Users can add new types to PostgreSQL using the [sql-createtype](../../reference/sql-commands/create-type.md#sql-createtype) command.

 [Data Types](#datatype-table) shows all the built-in general-purpose data types. Most of the alternative names listed in the “Aliases” column are the names used internally by PostgreSQL for historical reasons. In addition, some internally used or deprecated types are available, but are not listed here.
<a id="datatype-table"></a>

**Table: Data Types**

| Name | Aliases | Description |
| --- | --- | --- |
| `bigint` | `int8` | signed eight-byte integer |
| `bigserial` | `serial8` | autoincrementing eight-byte integer |
| <code>bit [ (</code><em>n</em><code>) ]</code> |  | fixed-length bit string |
| <code>bit varying [ (</code><em>n</em><code>) ]</code> | <code>varbit [ (</code><em>n</em><code>) ]</code> | variable-length bit string |
| `boolean` | `bool` | logical Boolean (true/false) |
| `box` |  | rectangular box on a plane |
| `bytea` |  | binary data (“byte array”) |
| <code>character [ (</code><em>n</em><code>) ]</code> | <code>char [ (</code><em>n</em><code>) ]</code> | fixed-length character string |
| <code>character varying [ (</code><em>n</em><code>) ]</code> | <code>varchar [ (</code><em>n</em><code>) ]</code> | variable-length character string |
| `cidr` |  | IPv4 or IPv6 network address |
| `circle` |  | circle on a plane |
| `date` |  | calendar date (year, month, day) |
| `double precision` | `float`, `float8` | double precision floating-point number (8 bytes) |
| `inet` |  | IPv4 or IPv6 host address |
| `integer` | `int`, `int4` | signed four-byte integer |
| <code>interval [ </code><em>fields</em><code> ] [ (</code><em>p</em><code>) ]</code> |  | time span |
| `json` |  | textual JSON data |
| `jsonb` |  | binary JSON data, decomposed |
| `line` |  | infinite line on a plane |
| `lseg` |  | line segment on a plane |
| `macaddr` |  | MAC (Media Access Control) address |
| `macaddr8` |  | MAC (Media Access Control) address (EUI-64 format) |
| `money` |  | currency amount |
| <code>numeric [ (</code><em>p</em><code>,          </code><em>s</em><code>) ]</code> | <code>decimal [ (</code><em>p</em><code>,          </code><em>s</em><code>) ]</code> | exact numeric of selectable precision |
| `path` |  | geometric path on a plane |
| `pg_lsn` |  | PostgreSQL Log Sequence Number |
| `pg_snapshot` |  | user-level transaction ID snapshot |
| `point` |  | geometric point on a plane |
| `polygon` |  | closed geometric path on a plane |
| `real` | `float4` | single precision floating-point number (4 bytes) |
| `smallint` | `int2` | signed two-byte integer |
| `smallserial` | `serial2` | autoincrementing two-byte integer |
| `serial` | `serial4` | autoincrementing four-byte integer |
| `text` |  | variable-length character string |
| <code>time [ (</code><em>p</em><code>) ] [ without time zone ]</code> |  | time of day (no time zone) |
| <code>time [ (</code><em>p</em><code>) ] with time zone</code> | `timetz` | time of day, including time zone |
| <code>timestamp [ (</code><em>p</em><code>) ] [ without time zone ]</code> |  | date and time (no time zone) |
| <code>timestamp [ (</code><em>p</em><code>) ] with time zone</code> | `timestamptz` | date and time, including time zone |
| `tsquery` |  | text search query |
| `tsvector` |  | text search document |
| `txid_snapshot` |  | user-level transaction ID snapshot (deprecated; see `pg_snapshot`) |
| `uuid` |  | universally unique identifier |
| `xml` |  | XML data |

!!! note "Compatibility"

    The following types (or spellings thereof) are specified by SQL: `bigint`, `bit`, `bit varying`, `boolean`, `char`, `character varying`, `character`, `varchar`, `date`, `double precision`, `integer`, `interval`, `json`, `numeric`, `decimal`, `real`, `smallint`, `time` (with or without time zone), `timestamp` (with or without time zone), `xml`.

 Each data type has an external representation determined by its input and output functions. Many of the built-in types have obvious external formats. However, several types are either unique to PostgreSQL, such as geometric paths, or have several possible formats, such as the date and time types. Some of the input and output functions are not invertible, i.e., the result of an output function might lose accuracy when compared to the original input.

- [Numeric Types](numeric-types.md#datatype-numeric)
- [Monetary Types](monetary-types.md#datatype-money)
- [Character Types](character-types.md#datatype-character)
- [Binary Data Types](binary-data-types.md#datatype-binary)
- [Date/Time Types](date-time-types.md#datatype-datetime)
- [Boolean Type](boolean-type.md#datatype-boolean)
- [Enumerated Types](enumerated-types.md#datatype-enum)
- [Geometric Types](geometric-types.md#datatype-geometric)
- [Network Address Types](network-address-types.md#datatype-net-types)
- [Bit String Types](bit-string-types.md#datatype-bit)
- [Text Search Types](text-search-types.md#datatype-textsearch)
- [UUID Type](uuid-type.md#datatype-uuid)
- [XML Type](xml-type.md#datatype-xml)
- [JSON Types](json-types.md#datatype-json)
- [Arrays](arrays.md#arrays)
- [Composite Types](composite-types.md#rowtypes)
- [Range Types](range-types.md#rangetypes)
- [Domain Types](domain-types.md#domains)
- [Object Identifier Types](object-identifier-types.md#datatype-oid)
- [`pg_lsn` Type](pg_lsn-type.md#datatype-pg-lsn)
- [Pseudo-Types](pseudo-types.md#datatype-pseudo)
