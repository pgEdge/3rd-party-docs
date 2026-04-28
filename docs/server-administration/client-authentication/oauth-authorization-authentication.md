<a id="auth-oauth"></a>

## OAuth Authorization/Authentication


 OAuth 2.0 is an industry-standard framework, defined in [RFC 6749](https://datatracker.ietf.org/doc/html/rfc6749), to enable third-party applications to obtain limited access to a protected resource. OAuth client support has to be enabled when PostgreSQL is built, see [Installation from Source Code](../installation-from-source-code/index.md#installation) for more information.


 This documentation uses the following terminology when discussing the OAuth ecosystem:

Resource Owner (or End User)
:   The user or system who owns protected resources and can grant access to them. This documentation also uses the term *end user* when the resource owner is a person. When you use psql to connect to the database using OAuth, you are the resource owner/end user.

Client
:   The system which accesses the protected resources using access tokens. Applications using libpq, such as psql, are the OAuth clients when connecting to a PostgreSQL cluster.

Resource Server
:   The system hosting the protected resources which are accessed by the client. The PostgreSQL cluster being connected to is the resource server.

Provider
:   The organization, product vendor, or other entity which develops and/or administers the OAuth authorization servers and clients for a given application. Different providers typically choose different implementation details for their OAuth systems; a client of one provider is not generally guaranteed to have access to the servers of another.


     This use of the term "provider" is not standard, but it seems to be in wide use colloquially. (It should not be confused with OpenID's similar term "Identity Provider". While the implementation of OAuth in PostgreSQL is intended to be interoperable and compatible with OpenID Connect/OIDC, it is not itself an OIDC client and does not require its use.)

Authorization Server
:   The system which receives requests from, and issues access tokens to, the client after the authenticated resource owner has given approval. PostgreSQL does not provide an authorization server; it is the responsibility of the OAuth provider.

<a id="auth-oauth-issuer"></a>
Issuer
:   An identifier for an authorization server, printed as an `https://` URL, which provides a trusted "namespace" for OAuth clients and applications. The issuer identifier allows a single authorization server to talk to the clients of mutually untrusting entities, as long as they maintain separate issuers.


!!! note

    For small deployments, there may not be a meaningful distinction between the "provider", "authorization server", and "issuer". However, for more complicated setups, there may be a one-to-many (or many-to-many) relationship: a provider may rent out multiple issuer identifiers to separate tenants, then provide multiple authorization servers, possibly with different supported feature sets, to interact with their clients.


 PostgreSQL supports bearer tokens, defined in [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750), which are a type of access token used with OAuth 2.0 where the token is an opaque string. The format of the access token is implementation specific and is chosen by each authorization server.


 The following configuration options are supported for OAuth:

`issuer`
:   An HTTPS URL which is either the exact [issuer identifier](#auth-oauth-issuer) of the authorization server, as defined by its discovery document, or a well-known URI that points directly to that discovery document. This parameter is required.


     When an OAuth client connects to the server, a URL for the discovery document will be constructed using the issuer identifier. By default, this URL uses the conventions of OpenID Connect Discovery: the path `/.well-known/openid-configuration` will be appended to the end of the issuer identifier. Alternatively, if the `issuer` contains a `/.well-known/` path segment, that URL will be provided to the client as-is.


    !!! warning

        The OAuth client in libpq requires the server's issuer setting to exactly match the issuer identifier which is provided in the discovery document, which must in turn match the client's [oauth_issuer](../../client-interfaces/libpq-c-library/database-connection-control-functions.md#libpq-connect-oauth-issuer) setting. No variations in case or formatting are permitted.

`scope`
:   A space-separated list of the OAuth scopes needed for the server to both authorize the client and authenticate the user. Appropriate values are determined by the authorization server and the OAuth validation module used (see [OAuth Validator Modules](../../server-programming/oauth-validator-modules/index.md#oauth-validators) for more information on validators). This parameter is required.

`validator`
:   The library to use for validating bearer tokens. If given, the name must exactly match one of the libraries listed in [oauth_validator_libraries](../server-configuration/connections-and-authentication.md#guc-oauth-validator-libraries). This parameter is optional unless `oauth_validator_libraries` contains more than one library, in which case it is required.

<a id="auth-oauth-validator-option"></a>
<code>validator.</code><em>option</em>
:   Validator modules may [define](../../server-programming/oauth-validator-modules/custom-hba-options.md#oauth-validator-hba) additional configuration options for `oauth` HBA entries. These validator-specific options are accessible via the `validator.*` "namespace". For example, a module may register the `validator.foo` and `validator.bar` options and define their effects on authentication.


     The name, syntax, and behavior of each *option* are not determined by PostgreSQL; consult the documentation for the validator module in use.


    !!! warning

        A limitation of the current implementation is that unrecognized *option* names will not be caught until connection time. A `pg_ctl reload` will succeed, but matching connections will fail:

        ```

        LOG:  connection received: host=[local]
        WARNING:  unrecognized authentication option name: "validator.bad"
        DETAIL:  The installed validator module ("my_validator") did not define an option named "bad".
        HINT:  All OAuth connections matching this line will fail. Correct the option and reload the server configuration.
        CONTEXT:  line 2 of configuration file "data/pg_hba.conf"
        ```
         Use caution when making changes to validator-specific HBA options in production systems.

`map`
:   Allows for mapping between OAuth identity provider and database user names. See [User Name Maps](user-name-maps.md#auth-username-maps) for details. If a map is not specified, the user name associated with the token (as determined by the OAuth validator) must exactly match the role name being requested. This parameter is optional.

<a id="auth-oauth-delegate-ident-mapping"></a>
`delegate_ident_mapping`
:   An advanced option which is not intended for common use.


     When set to `1`, standard user mapping with `pg_ident.conf` is skipped, and the OAuth validator takes full responsibility for mapping end user identities to database roles. If the validator authorizes the token, the server trusts that the user is allowed to connect under the requested role, and the connection is allowed to proceed regardless of the authentication status of the user.


     This parameter is incompatible with `map`.


    !!! warning

        `delegate_ident_mapping` provides additional flexibility in the design of the authentication system, but it also requires careful implementation of the OAuth validator, which must determine whether the provided token carries sufficient end-user privileges in addition to the [standard checks](../../server-programming/oauth-validator-modules/index.md#oauth-validators) required of all validators. Use with caution.
