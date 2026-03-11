<a id="protocol"></a>

# Frontend/Backend Protocol

 PostgreSQL uses a message-based protocol for communication between frontends and backends (clients and servers). The protocol is supported over TCP/IP and also over Unix-domain sockets. Port number 5432 has been registered with IANA as the customary TCP port number for servers supporting this protocol, but in practice any non-privileged port number can be used.

 This document describes version 3.0 of the protocol, implemented in PostgreSQL 7.4 and later. For descriptions of the earlier protocol versions, see previous releases of the PostgreSQL documentation. A single server can support multiple protocol versions. The initial startup-request message tells the server which protocol version the client is attempting to use. If the major version requested by the client is not supported by the server, the connection will be rejected (for example, this would occur if the client requested protocol version 4.0, which does not exist as of this writing). If the minor version requested by the client is not supported by the server (e.g., the client requests version 3.1, but the server supports only 3.0), the server may either reject the connection or may respond with a NegotiateProtocolVersion message containing the highest minor protocol version which it supports. The client may then choose either to continue with the connection using the specified protocol version or to abort the connection.

 In order to serve multiple clients efficiently, the server launches a new “backend” process for each client. In the current implementation, a new child process is created immediately after an incoming connection is detected. This is transparent to the protocol, however. For purposes of the protocol, the terms “backend” and “server” are interchangeable; likewise “frontend” and “client” are interchangeable.

- [Overview](overview.md#protocol-overview)
- [Message Flow](message-flow.md#protocol-flow)
- [SASL Authentication](sasl-authentication.md#sasl-authentication)
- [Streaming Replication Protocol](streaming-replication-protocol.md#protocol-replication)
- [Logical Streaming Replication Protocol](logical-streaming-replication-protocol.md#protocol-logical-replication)
- [Message Data Types](message-data-types.md#protocol-message-types)
- [Message Formats](message-formats.md#protocol-message-formats)
- [Error and Notice Message Fields](error-and-notice-message-fields.md#protocol-error-fields)
- [Logical Replication Message Formats](logical-replication-message-formats.md#protocol-logicalrep-message-formats)
- [Summary of Changes since Protocol 2.0](summary-of-changes-since-protocol-2-0.md#protocol-changes)
