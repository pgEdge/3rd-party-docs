<a id="protocol-overview"></a>

## Overview


 The protocol has separate phases for startup and normal operation. In the startup phase, the frontend opens a connection to the server and authenticates itself to the satisfaction of the server. (This might involve a single message, or multiple messages depending on the authentication method being used.) If all goes well, the server then sends status information to the frontend, and finally enters normal operation. Except for the initial startup-request message, this part of the protocol is driven by the server.


 During normal operation, the frontend sends queries and other commands to the backend, and the backend sends back query results and other responses. There are a few cases (such as `NOTIFY`) wherein the backend will send unsolicited messages, but for the most part this portion of a session is driven by frontend requests.


 Termination of the session is normally by frontend choice, but can be forced by the backend in certain cases. In any case, when the backend closes the connection, it will roll back any open (incomplete) transaction before exiting.


 Within normal operation, SQL commands can be executed through either of two sub-protocols. In the “simple query” protocol, the frontend just sends a textual query string, which is parsed and immediately executed by the backend. In the “extended query” protocol, processing of queries is separated into multiple steps: parsing, binding of parameter values, and execution. This offers flexibility and performance benefits, at the cost of extra complexity.


 Normal operation has additional sub-protocols for special operations such as `COPY`.
 <a id="protocol-message-concepts"></a>

### Messaging Overview


 All communication is through a stream of messages. The first byte of a message identifies the message type, and the next four bytes specify the length of the rest of the message (this length count includes itself, but not the message-type byte). The remaining contents of the message are determined by the message type. For historical reasons, the very first message sent by the client (the startup message) has no initial message-type byte.


 To avoid losing synchronization with the message stream, both servers and clients typically read an entire message into a buffer (using the byte count) before attempting to process its contents. This allows easy recovery if an error is detected while processing the contents. In extreme situations (such as not having enough memory to buffer the message), the receiver can use the byte count to determine how much input to skip before it resumes reading messages.


 Conversely, both servers and clients must take care never to send an incomplete message. This is commonly done by marshaling the entire message in a buffer before beginning to send it. If a communications failure occurs partway through sending or receiving a message, the only sensible response is to abandon the connection, since there is little hope of recovering message-boundary synchronization.
  <a id="protocol-query-concepts"></a>

### Extended Query Overview


 In the extended-query protocol, execution of SQL commands is divided into multiple steps. The state retained between steps is represented by two types of objects: *prepared statements* and *portals*. A prepared statement represents the result of parsing and semantic analysis of a textual query string. A prepared statement is not in itself ready to execute, because it might lack specific values for *parameters*. A portal represents a ready-to-execute or already-partially-executed statement, with any missing parameter values filled in. (For `SELECT` statements, a portal is equivalent to an open cursor, but we choose to use a different term since cursors don't handle non-`SELECT` statements.)


 The overall execution cycle consists of a *parse* step, which creates a prepared statement from a textual query string; a *bind* step, which creates a portal given a prepared statement and values for any needed parameters; and an *execute* step that runs a portal's query. In the case of a query that returns rows (`SELECT`, `SHOW`, etc.), the execute step can be told to fetch only a limited number of rows, so that multiple execute steps might be needed to complete the operation.


 The backend can keep track of multiple prepared statements and portals (but note that these exist only within a session, and are never shared across sessions). Existing prepared statements and portals are referenced by names assigned when they were created. In addition, an “unnamed” prepared statement and portal exist. Although these behave largely the same as named objects, operations on them are optimized for the case of executing a query only once and then discarding it, whereas operations on named objects are optimized on the expectation of multiple uses.
  <a id="protocol-format-codes"></a>

### Formats and Format Codes


 Data of a particular data type might be transmitted in any of several different *formats*. As of PostgreSQL 7.4 the only supported formats are “text” and “binary”, but the protocol makes provision for future extensions. The desired format for any value is specified by a *format code*. Clients can specify a format code for each transmitted parameter value and for each column of a query result. Text has format code zero, binary has format code one, and all other format codes are reserved for future definition.


 The text representation of values is whatever strings are produced and accepted by the input/output conversion functions for the particular data type. In the transmitted representation, there is no trailing null character; the frontend must add one to received values if it wants to process them as C strings. (The text format does not allow embedded nulls, by the way.)


 Binary representations for integers use network byte order (most significant byte first). For other data types consult the documentation or source code to learn about the binary representation. Keep in mind that binary representations for complex data types might change across server versions; the text format is usually the more portable choice.
  <a id="protocol-versions"></a>

### Protocol Versions


 The current, latest version of the protocol is version 3.2. However, for backwards compatibility with old server versions and middleware that don't support the version negotiation yet, libpq still uses protocol version 3.0 by default.


!!! note

    During the PostgreSQL 19 beta period, libpq will instead default to requesting protocol version 3.9999, to test that servers and middleware properly implement protocol version negotiation. Servers that support negotiation will automatically downgrade to version 3.2 or 3.0. Users can bypass this beta-only behavior by explicitly setting `max_protocol_version=3.0` in their connection string.


 A single server can support multiple protocol versions. The initial startup-request message tells the server which protocol version the client is attempting to use. If the major version requested by the client is not supported by the server, the connection will be rejected (for example, this would occur if the client requested protocol version 4.0, which does not exist as of this writing). If the minor version requested by the client is not supported by the server (e.g., the client requests version 3.2, but the server supports only 3.0), the server may either reject the connection or may respond with a NegotiateProtocolVersion message containing the highest minor protocol version which it supports. The client may then choose either to continue with the connection using the specified protocol version or to abort the connection.


 The protocol negotiation was introduced in PostgreSQL version 9.3.21. Earlier versions would reject the connection if the client requested a minor version that was not supported by the server.


 [Supported Protocol Versions](#protocol-versions-table) shows the currently supported protocol versions. [Other Protocol Versions](#other-protocol-versions-table) documents protocol versions that are unsupported or otherwise reserved.
 <a id="protocol-versions-table"></a>

**Table: Supported Protocol Versions**

| Version | Supported by | Description |
| --- | --- | --- |
| 3.2 | PostgreSQL 18 and later | Current latest version. The secret key used in query cancellation was enlarged from 4 bytes to a variable length field. The BackendKeyData message was changed to accommodate that, and the CancelRequest message was redefined to have a variable length payload. |
| 3.0 | PostgreSQL 7.4 and later |
 <a id="other-protocol-versions-table"></a>

**Table: Other Protocol Versions**

| Version | Supported by | Description |
| --- | --- | --- |
| 3.9999 | - | Reserved for protocol greasing. libpq may use this version, which is higher than any minor version the project ever expects to use, to test that servers and middleware properly implement protocol version negotiation. Servers *must not* add special-case logic for this version; they should simply compare it to their latest supported version (which will always be smaller) and downgrade via a NegotiateProtocolVersion message. |
| 3.1 | - | Reserved. Version 3.1 has not been used by any PostgreSQL version, but it was skipped because old versions of the popular pgbouncer application had a bug in the protocol negotiation which made it incorrectly claim that it supported version 3.1. |
| 2.0 | up to PostgreSQL 13 | Obsolete. See previous releases of the PostgreSQL documentation for details. |
  <a id="protocol-extensions"></a>

### Protocol Extensions


 Servers and clients may additionally negotiate individual extensions to the protocol version in use. These are offered by the client in the startup message, as specially-named parameters with a `_pq_.` prefix. Servers reject any unknown or unsupported extensions by sending a NegotiateProtocolVersion message containing the list of rejected parameter names, at which point the client may choose whether to continue with the connection. [Supported Protocol Extensions](#protocol-extensions-table) and [Reserved Protocol Extensions](#reserved-protocol-extensions-table) document the supported and reserved protocol extension parameters, respectively.
 <a id="protocol-extensions-table"></a>

**Table: Supported Protocol Extensions**

<table>
<thead>
<tr>
<th>Parameter Name</th>
<th>Values</th>
<th>Supported by</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>(No supported protocol extensions are currently defined.)</em></td>
</tr>
</tbody>
</table>
 <a id="reserved-protocol-extensions-table"></a>

**Table: Reserved Protocol Extensions**

| Parameter Name | Description |
| --- | --- |
| <code>_pq_.</code><em>[name]</em> | Any other parameter names beginning with `_pq_.`, that are not defined above, are reserved for future protocol expansion. Servers *must* reject any that are received from a client, by sending a NegotiateProtocolVersion message during the [startup flow](message-flow.md#protocol-flow-start-up), and should otherwise continue the connection. |
| `_pq_.test_protocol_negotiation` | Reserved for protocol greasing. libpq may send this extension to test that servers and middleware properly implement protocol extension negotiation. Servers *must not* add special-case logic for this parameter; they should simply send the list of all unsupported options (including this one) via a NegotiateProtocolVersion message. |
