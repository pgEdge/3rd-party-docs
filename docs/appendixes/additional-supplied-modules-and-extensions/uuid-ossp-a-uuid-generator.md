<a id="uuid-ossp"></a>

## uuid-ossp — a UUID generator


 The `uuid-ossp` module provides functions to generate universally unique identifiers (UUIDs) using one of several standard algorithms. There are also functions to produce certain special UUID constants. This module is only necessary for special requirements beyond what is available in core PostgreSQL. See [UUID Functions](../../the-sql-language/functions-and-operators/uuid-functions.md#functions-uuid) for built-in ways to generate UUIDs.


 This module is considered “trusted”, that is, it can be installed by non-superusers who have `CREATE` privilege on the current database.
 <a id="uuid-ossp-functions-sect"></a>

### `uuid-ossp` Functions


 [Functions for UUID Generation](#uuid-ossp-functions) shows the functions available to generate UUIDs. The relevant standards ITU-T Rec. X.667, ISO/IEC 9834-8:2005, and [RFC 4122](https://datatracker.ietf.org/doc/html/rfc4122) specify four algorithms for generating UUIDs, identified by the version numbers 1, 3, 4, and 5. (There is no version 2 algorithm.) Each of these algorithms could be suitable for a different set of applications.
 <a id="uuid-ossp-functions"></a>

**Table: Functions for UUID Generation**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>uuid_generate_v1</code> () <code>uuid</code></td>
<td>Generates a version 1 UUID. This involves the MAC address of the computer and a time stamp. Note that UUIDs of this kind reveal the identity of the computer that created the identifier and the time at which it did so, which might make it unsuitable for certain security-sensitive applications.</td>
<td></td>
</tr>
<tr>
<td><code>uuid_generate_v1mc</code> () <code>uuid</code></td>
<td>Generates a version 1 UUID, but uses a random multicast MAC address instead of the real MAC address of the computer.</td>
<td></td>
</tr>
<tr>
<td><code>uuid_generate_v3</code> ( <code>namespace</code> <code>uuid</code>, <code>name</code> <code>text</code> ) <code>uuid</code></td>
<td>Generates a version 3 UUID in the given namespace using the specified input name. The namespace should be one of the special constants produced by the <code>uuid_ns_*()</code> functions shown in <a href="#uuid-ossp-constants">Functions Returning UUID Constants</a>. (It could be any UUID in theory.) The name is an identifier in the selected namespace.</td>
<td><p>For example:</p>
<pre><code class="language-sql">
SELECT uuid_generate_v3(uuid_ns_url(), 'http://www.postgresql.org');</code></pre></td>
</tr>
<tr>
<td><code>uuid_generate_v4</code> () <code>uuid</code></td>
<td>Generates a version 4 UUID, which is derived entirely from random numbers.</td>
<td></td>
</tr>
<tr>
<td><code>uuid_generate_v5</code> ( <code>namespace</code> <code>uuid</code>, <code>name</code> <code>text</code> ) <code>uuid</code></td>
<td>Generates a version 5 UUID, which works like a version 3 UUID except that SHA-1 is used as a hashing method. Version 5 should be preferred over version 3 because SHA-1 is thought to be more secure than MD5.</td>
<td></td>
</tr>
</tbody>
</table>
 <a id="uuid-ossp-constants"></a>

**Table: Functions Returning UUID Constants**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>uuid_nil</code> () <code>uuid</code></td>
<td>Returns a “nil” UUID constant, which does not occur as a real UUID.</td>
<td></td>
</tr>
<tr>
<td><code>uuid_ns_dns</code> () <code>uuid</code></td>
<td>Returns a constant designating the DNS namespace for UUIDs.</td>
<td></td>
</tr>
<tr>
<td><code>uuid_ns_url</code> () <code>uuid</code></td>
<td>Returns a constant designating the URL namespace for UUIDs.</td>
<td></td>
</tr>
<tr>
<td><code>uuid_ns_oid</code> () <code>uuid</code></td>
<td>Returns a constant designating the ISO object identifier (OID) namespace for UUIDs. (This pertains to ASN.1 OIDs, which are unrelated to the OIDs used in PostgreSQL.)</td>
<td></td>
</tr>
<tr>
<td><code>uuid_ns_x500</code> () <code>uuid</code></td>
<td>Returns a constant designating the X.500 distinguished name (DN) namespace for UUIDs.</td>
<td></td>
</tr>
</tbody>
</table>
  <a id="uuid-ossp-building"></a>

### Building `uuid-ossp`


 Historically this module depended on the OSSP UUID library, which accounts for the module's name. While the OSSP UUID library can still be found at [http://www.ossp.org/pkg/lib/uuid/](http://www.ossp.org/pkg/lib/uuid/), it is not well maintained, and is becoming increasingly difficult to port to newer platforms. `uuid-ossp` can now be built without the OSSP library on some platforms. On FreeBSD and some other BSD-derived platforms, suitable UUID creation functions are included in the core `libc` library. On Linux, macOS, and some other platforms, suitable functions are provided in the `libuuid` library, which originally came from the `e2fsprogs` project (though on modern Linux it is considered part of `util-linux-ng`). When invoking `configure`, specify `--with-uuid=bsd` to use the BSD functions, or `--with-uuid=e2fs` to use `e2fsprogs`' `libuuid`, or `--with-uuid=ossp` to use the OSSP UUID library. More than one of these libraries might be available on a particular machine, so `configure` does not automatically choose one.
  <a id="uuid-ossp-author"></a>

### Author


 Peter Eisentraut [peter_e@gmx.net](mailto:peter_e@gmx.net)
