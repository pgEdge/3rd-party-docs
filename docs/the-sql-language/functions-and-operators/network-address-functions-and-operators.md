<a id="functions-net"></a>

## Network Address Functions and Operators


 The IP network address types, `cidr` and `inet`, support the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) as well as the specialized operators and functions shown in [IP Address Operators](#cidr-inet-operators-table) and [IP Address Functions](#cidr-inet-functions-table).


 Any `cidr` value can be cast to `inet` implicitly; therefore, the operators and functions shown below as operating on `inet` also work on `cidr` values. (Where there are separate functions for `inet` and `cidr`, it is because the behavior should be different for the two cases.) Also, it is permitted to cast an `inet` value to `cidr`. When this is done, any bits to the right of the netmask are silently zeroed to create a valid `cidr` value.
 <a id="cidr-inet-operators-table"></a>

**Table: IP Address Operators**

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
<td><code>inet</code> <code>&lt;&lt;</code> <code>inet</code> <code>boolean</code></td>
<td>Is subnet strictly contained by subnet? This operator, and the next four, test for subnet inclusion. They consider only the network parts of the two addresses (ignoring any bits to the right of the netmasks) and determine whether one network is identical to or a subnet of the other.</td>
<td><code>inet '192.168.1.5' &lt;&lt; inet '192.168.1/24'</code> <code>t</code><br><code>inet '192.168.0.5' &lt;&lt; inet '192.168.1/24'</code> <code>f</code><br><code>inet '192.168.1/24' &lt;&lt; inet '192.168.1/24'</code> <code>f</code></td>
</tr>
<tr>
<td><code>inet</code> <code>&lt;&lt;=</code> <code>inet</code> <code>boolean</code></td>
<td>Is subnet contained by or equal to subnet?</td>
<td><code>inet '192.168.1/24' &lt;&lt;= inet '192.168.1/24'</code> <code>t</code></td>
</tr>
<tr>
<td><code>inet</code> <code>&gt;&gt;</code> <code>inet</code> <code>boolean</code></td>
<td>Does subnet strictly contain subnet?</td>
<td><code>inet '192.168.1/24' &gt;&gt; inet '192.168.1.5'</code> <code>t</code></td>
</tr>
<tr>
<td><code>inet</code> <code>&gt;&gt;=</code> <code>inet</code> <code>boolean</code></td>
<td>Does subnet contain or equal subnet?</td>
<td><code>inet '192.168.1/24' &gt;&gt;= inet '192.168.1/24'</code> <code>t</code></td>
</tr>
<tr>
<td><code>inet</code> <code>&amp;&amp;</code> <code>inet</code> <code>boolean</code></td>
<td>Does either subnet contain or equal the other?</td>
<td><code>inet '192.168.1/24' &amp;&amp; inet '192.168.1.80/28'</code> <code>t</code><br><code>inet '192.168.1/24' &amp;&amp; inet '192.168.2.0/28'</code> <code>f</code></td>
</tr>
<tr>
<td><code>~</code> <code>inet</code> <code>inet</code></td>
<td>Computes bitwise NOT.</td>
<td><code>~ inet '192.168.1.6'</code> <code>63.87.254.249</code></td>
</tr>
<tr>
<td><code>inet</code> <code>&amp;</code> <code>inet</code> <code>inet</code></td>
<td>Computes bitwise AND.</td>
<td><code>inet '192.168.1.6' &amp; inet '0.0.0.255'</code> <code>0.0.0.6</code></td>
</tr>
<tr>
<td><code>inet</code> <code>|</code> <code>inet</code> <code>inet</code></td>
<td>Computes bitwise OR.</td>
<td><code>inet '192.168.1.6' | inet '0.0.0.255'</code> <code>192.168.1.255</code></td>
</tr>
<tr>
<td><code>inet</code> <code>+</code> <code>bigint</code> <code>inet</code></td>
<td>Adds an offset to an address.</td>
<td><code>inet '192.168.1.6' + 25</code> <code>192.168.1.31</code></td>
</tr>
<tr>
<td><code>bigint</code> <code>+</code> <code>inet</code> <code>inet</code></td>
<td>Adds an offset to an address.</td>
<td><code>200 + inet '::ffff:fff0:1'</code> <code>::ffff:255.240.0.201</code></td>
</tr>
<tr>
<td><code>inet</code> <code>-</code> <code>bigint</code> <code>inet</code></td>
<td>Subtracts an offset from an address.</td>
<td><code>inet '192.168.1.43' - 36</code> <code>192.168.1.7</code></td>
</tr>
<tr>
<td><code>inet</code> <code>-</code> <code>inet</code> <code>bigint</code></td>
<td>Computes the difference of two addresses.</td>
<td><code>inet '192.168.1.43' - inet '192.168.1.19'</code> <code>24</code><br><code>inet '::1' - inet '::ffff:1'</code> <code>-4294901760</code></td>
</tr>
</tbody>
</table>
 <a id="cidr-inet-functions-table"></a>

**Table: IP Address Functions**

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
<td><code>abbrev</code> ( <code>inet</code> ) <code>text</code></td>
<td>Creates an abbreviated display format as text. (The result is the same as the <code>inet</code> output function produces; it is “abbreviated” only in comparison to the result of an explicit cast to <code>text</code>, which for historical reasons will never suppress the netmask part.)</td>
<td><code>abbrev(inet '10.1.0.0/32')</code> <code>10.1.0.0</code></td>
</tr>
<tr>
<td><code>abbrev</code> ( <code>cidr</code> ) <code>text</code></td>
<td>Creates an abbreviated display format as text. (The abbreviation consists of dropping all-zero octets to the right of the netmask; more examples are in <a href="../data-types/network-address-types.md#datatype-net-cidr-table"><code>cidr</code> Type Input Examples</a>.)</td>
<td><code>abbrev(cidr '10.1.0.0/16')</code> <code>10.1/16</code></td>
</tr>
<tr>
<td><code>broadcast</code> ( <code>inet</code> ) <code>inet</code></td>
<td>Computes the broadcast address for the address's network.</td>
<td><code>broadcast(inet '192.168.1.5/24')</code> <code>192.168.1.255/24</code></td>
</tr>
<tr>
<td><code>family</code> ( <code>inet</code> ) <code>integer</code></td>
<td>Returns the address's family: <code>4</code> for IPv4, <code>6</code> for IPv6.</td>
<td><code>family(inet '::1')</code> <code>6</code></td>
</tr>
<tr>
<td><code>host</code> ( <code>inet</code> ) <code>text</code></td>
<td>Returns the IP address as text, ignoring the netmask.</td>
<td><code>host(inet '192.168.1.0/24')</code> <code>192.168.1.0</code></td>
</tr>
<tr>
<td><code>hostmask</code> ( <code>inet</code> ) <code>inet</code></td>
<td>Computes the host mask for the address's network.</td>
<td><code>hostmask(inet '192.168.23.20/30')</code> <code>0.0.0.3</code></td>
</tr>
<tr>
<td><code>inet_merge</code> ( <code>inet</code>, <code>inet</code> ) <code>cidr</code></td>
<td>Computes the smallest network that includes both of the given networks.</td>
<td><code>inet_merge(inet '192.168.1.5/24', inet '192.168.2.5/24')</code> <code>192.168.0.0/22</code></td>
</tr>
<tr>
<td><code>inet_same_family</code> ( <code>inet</code>, <code>inet</code> ) <code>boolean</code></td>
<td>Tests whether the addresses belong to the same IP family.</td>
<td><code>inet_same_family(inet '192.168.1.5/24', inet '::1')</code> <code>f</code></td>
</tr>
<tr>
<td><code>masklen</code> ( <code>inet</code> ) <code>integer</code></td>
<td>Returns the netmask length in bits.</td>
<td><code>masklen(inet '192.168.1.5/24')</code> <code>24</code></td>
</tr>
<tr>
<td><code>netmask</code> ( <code>inet</code> ) <code>inet</code></td>
<td>Computes the network mask for the address's network.</td>
<td><code>netmask(inet '192.168.1.5/24')</code> <code>255.255.255.0</code></td>
</tr>
<tr>
<td><code>network</code> ( <code>inet</code> ) <code>cidr</code></td>
<td>Returns the network part of the address, zeroing out whatever is to the right of the netmask. (This is equivalent to casting the value to <code>cidr</code>.)</td>
<td><code>network(inet '192.168.1.5/24')</code> <code>192.168.1.0/24</code></td>
</tr>
<tr>
<td><code>set_masklen</code> ( <code>inet</code>, <code>integer</code> ) <code>inet</code></td>
<td>Sets the netmask length for an <code>inet</code> value. The address part does not change.</td>
<td><code>set_masklen(inet '192.168.1.5/24', 16)</code> <code>192.168.1.5/16</code></td>
</tr>
<tr>
<td><code>set_masklen</code> ( <code>cidr</code>, <code>integer</code> ) <code>cidr</code></td>
<td>Sets the netmask length for a <code>cidr</code> value. Address bits to the right of the new netmask are set to zero.</td>
<td><code>set_masklen(cidr '192.168.1.0/24', 16)</code> <code>192.168.0.0/16</code></td>
</tr>
<tr>
<td><code>text</code> ( <code>inet</code> ) <code>text</code></td>
<td>Returns the unabbreviated IP address and netmask length as text. (This has the same result as an explicit cast to <code>text</code>.)</td>
<td><code>text(inet '192.168.1.5')</code> <code>192.168.1.5/32</code></td>
</tr>
</tbody>
</table>


!!! tip

    The `abbrev`, `host`, and `text` functions are primarily intended to offer alternative display formats for IP addresses.


 The MAC address types, `macaddr` and `macaddr8`, support the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) as well as the specialized functions shown in [MAC Address Functions](#macaddr-functions-table). In addition, they support the bitwise logical operators `~`, `&` and `|` (NOT, AND and OR), just as shown above for IP addresses.
 <a id="macaddr-functions-table"></a>

**Table: MAC Address Functions**

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
<td><code>trunc</code> ( <code>macaddr</code> ) <code>macaddr</code></td>
<td>Sets the last 3 bytes of the address to zero. The remaining prefix can be associated with a particular manufacturer (using data not included in PostgreSQL).</td>
<td><code>trunc(macaddr '12:34:56:78:90:ab')</code> <code>12:34:56:00:00:00</code></td>
</tr>
<tr>
<td><code>trunc</code> ( <code>macaddr8</code> ) <code>macaddr8</code></td>
<td>Sets the last 5 bytes of the address to zero. The remaining prefix can be associated with a particular manufacturer (using data not included in PostgreSQL).</td>
<td><code>trunc(macaddr8 '12:34:56:78:90:ab:cd:ef')</code> <code>12:34:56:00:00:00:00:00</code></td>
</tr>
<tr>
<td><code>macaddr8_set7bit</code> ( <code>macaddr8</code> ) <code>macaddr8</code></td>
<td>Sets the 7th bit of the address to one, creating what is known as modified EUI-64, for inclusion in an IPv6 address.</td>
<td><code>macaddr8_set7bit(macaddr8 '00:34:56:ab:cd:ef')</code> <code>02:34:56:ff:fe:ab:cd:ef</code></td>
</tr>
</tbody>
</table>
