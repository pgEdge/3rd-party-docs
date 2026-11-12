<a id="auth-cert"></a>

## Certificate Authentication


 This authentication method uses SSL client certificates to perform authentication. It is therefore only available for SSL connections; see [OpenSSL Configuration](../server-setup-and-operation/secure-tcp-ip-connections-with-ssl.md#ssl-openssl-config) for SSL configuration instructions. When using this authentication method, the server will require that the client provide a valid, trusted certificate. No password prompt will be sent to the client. The `cn` (Common Name) attribute of the certificate will be compared to the requested database user name, and if they match the login will be allowed. User name mapping can be used to allow `cn` to be different from the database user name.


 The following configuration options are supported for SSL certificate authentication:

`map`
:   Allows for mapping between system and database user names. See [User Name Maps](user-name-maps.md#auth-username-maps) for details.


 It is redundant to use the `clientcert` option with `cert` authentication because `cert` authentication is effectively `trust` authentication with `clientcert=verify-full`.
