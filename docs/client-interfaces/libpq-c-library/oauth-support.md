<a id="libpq-oauth"></a>

## OAuth Support


 libpq implements support for the OAuth v2 Device Authorization client flow, documented in [RFC 8628](https://datatracker.ietf.org/doc/html/rfc8628), as an optional module. See the [installation documentation](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-option-with-libcurl) for information on how to enable support for Device Authorization as a builtin flow.


 When support is enabled and the optional module installed, libpq will use the builtin flow by default if the server [requests a bearer token](../../server-administration/client-authentication/oauth-authorization-authentication.md#auth-oauth) during authentication. This flow can be utilized even if the system running the client application does not have a usable web browser, for example when running a client via SSH.


 The builtin flow will, by default, print a URL to visit and a user code to enter there:

```

$ psql 'dbname=postgres oauth_issuer=https://example.com oauth_client_id=...'
Visit https://example.com/device and enter the code: ABCD-EFGH
```
 (This prompt may be [customized](#libpq-oauth-authdata-prompt-oauth-device).) The user will then log into their OAuth provider, which will ask whether to allow libpq and the server to perform actions on their behalf. It is always a good idea to carefully review the URL and permissions displayed, to ensure they match expectations, before continuing. Permissions should not be given to untrusted third parties.


 Client applications may implement their own flows to customize interaction and integration with applications. See [Authdata Hooks](#libpq-oauth-authdata-hooks) for more information on how add a custom flow to libpq.


 For an OAuth client flow to be usable, the connection string must at minimum contain [oauth_issuer](database-connection-control-functions.md#libpq-connect-oauth-issuer) and [oauth_client_id](database-connection-control-functions.md#libpq-connect-oauth-client-id). (These settings are determined by your organization's OAuth provider.) The builtin flow additionally requires the OAuth authorization server to publish a device authorization endpoint.


!!! note

    The builtin Device Authorization flow is not currently supported on Windows. Custom client flows may still be implemented.
 <a id="libpq-oauth-authdata-hooks"></a>

### Authdata Hooks


 The behavior of the OAuth flow may be modified or replaced by a client using the following hook API:

<a id="libpq-PQsetAuthDataHook"></a>

`PQsetAuthDataHook`
:   Sets the `PGauthDataHook`, overriding libpq's handling of one or more aspects of its OAuth client flow.

    ```

    void PQsetAuthDataHook(PQauthDataHook_type hook);
    ```
     If *hook* is `NULL`, the default handler will be reinstalled. Otherwise, the application passes a pointer to a callback function with the signature:

    ```

    int hook_fn(PGauthData type, PGconn *conn, void *data);
    ```
     which libpq will call when an action is required of the application. *type* describes the request being made, *conn* is the connection handle being authenticated, and *data* points to request-specific metadata. The contents of this pointer are determined by *type*; see [Hook Types](#libpq-oauth-authdata-hooks-types) for the supported list.


     Hooks can be chained together to allow cooperative and/or fallback behavior. In general, a hook implementation should examine the incoming *type* (and, potentially, the request metadata and/or the settings for the particular *conn* in use) to decide whether or not to handle a specific piece of authdata. If not, it should delegate to the previous hook in the chain (retrievable via `PQgetAuthDataHook`).


     Success is indicated by returning an integer greater than zero. Returning a negative integer signals an error condition and abandons the connection attempt. (A zero value is reserved for the default implementation.)
<a id="libpq-PQgetAuthDataHook"></a>

`PQgetAuthDataHook`
:   Retrieves the current value of `PGauthDataHook`.

    ```

    PQauthDataHook_type PQgetAuthDataHook(void);
    ```
     At initialization time (before the first call to `PQsetAuthDataHook`), this function will return `PQdefaultAuthDataHook`.

 <a id="libpq-oauth-authdata-hooks-types"></a>

#### Hook Types


 The following `PGauthData` types and their corresponding *data* structures are defined:

<a id="libpq-oauth-authdata-prompt-oauth-device"></a>

`PQAUTHDATA_PROMPT_OAUTH_DEVICE`
:   *Available in PostgreSQL 18 and later.*


     Replaces the default user prompt during the builtin device authorization client flow. *data* points to an instance of `PGpromptOAuthDevice`:

    ```

    typedef struct _PGpromptOAuthDevice
    {
        const char *verification_uri;   /* verification URI to visit */
        const char *user_code;          /* user code to enter */
        const char *verification_uri_complete;  /* optional combination of URI and
                                                 * code, or NULL */
        int         expires_in;         /* seconds until user code expires */
    } PGpromptOAuthDevice;
    ```


     The OAuth Device Authorization flow which [can be included](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-option-with-libcurl) in libpq requires the end user to visit a URL with a browser, then enter a code which permits libpq to connect to the server on their behalf. The default prompt simply prints the `verification_uri` and `user_code` on standard error. Replacement implementations may display this information using any preferred method, for example with a GUI.


     This callback is only invoked during the builtin device authorization flow. If the application installs a [custom OAuth flow](#libpq-oauth-authdata-oauth-bearer-token), or libpq was not built with support for the builtin flow, this authdata type will not be used.


     If a non-NULL `verification_uri_complete` is provided, it may optionally be used for non-textual verification (for example, by displaying a QR code). The URL and user code should still be displayed to the end user in this case, because the code will be manually confirmed by the provider, and the URL lets users continue even if they can't use the non-textual method. For more information, see section 3.3.1 in [RFC 8628](https://datatracker.ietf.org/doc/html/rfc8628#section-3.3.1).
<a id="libpq-oauth-authdata-oauth-bearer-token"></a>

`PQAUTHDATA_OAUTH_BEARER_TOKEN`
:   *Available in PostgreSQL 18 and later.*


     Adds a custom implementation of a flow, replacing the builtin flow if it is [installed](../../server-administration/installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-option-with-libcurl). The hook should either directly return a Bearer token for the current user/issuer/scope combination, if one is available without blocking, or else set up an asynchronous callback to retrieve one.


    !!! note

        For PostgreSQL releases 19 and later, applications should prefer [`PQAUTHDATA_OAUTH_BEARER_TOKEN_V2`](#libpq-oauth-authdata-oauth-bearer-token-v2).


     *data* points to an instance of `PGoauthBearerRequest`, which should be filled in by the implementation:

    ```

    typedef struct PGoauthBearerRequest
    {
        /* Hook inputs (constant across all calls) */
        const char *openid_configuration; /* OIDC discovery URL */
        const char *scope;                /* required scope(s), or NULL */

        /* Hook outputs */

        /*
         * Callback implementing a custom asynchronous OAuth flow. The signature is
         * platform-dependent: PQ_SOCKTYPE is SOCKET on Windows, and int everywhere
         * else.
         */
        PostgresPollingStatusType (*async) (PGconn *conn,
                                            struct PGoauthBearerRequest *request,
                                            PQ_SOCKTYPE *altsock);

        /* Callback to clean up custom allocations. */
        void        (*cleanup) (PGconn *conn, struct PGoauthBearerRequest *request);

        char       *token;   /* acquired Bearer token */
        void       *user;    /* hook-defined allocated data */
    } PGoauthBearerRequest;
    ```


     Two pieces of information are provided to the hook by libpq: *openid_configuration* contains the URL of an OAuth discovery document describing the authorization server's supported flows, and *scope* contains a (possibly empty) space-separated list of OAuth scopes which are required to access the server. Either or both may be `NULL` to indicate that the information was not discoverable. (In this case, implementations may be able to establish the requirements using some other preconfigured knowledge, or they may choose to fail.)


     The final output of the hook is *token*, which must point to a valid Bearer token for use on the connection. (This token should be issued by the [oauth_issuer](database-connection-control-functions.md#libpq-connect-oauth-issuer) and hold the requested scopes, or the connection will be rejected by the server's validator module.) The allocated token string must remain valid until libpq is finished connecting; the hook should set a *cleanup* callback which will be called when libpq no longer requires it.


     If an implementation cannot immediately produce a *token* during the initial call to the hook, it should set the *async* callback to handle nonblocking communication with the authorization server.  (Performing blocking operations during the `PQAUTHDATA_OAUTH_BEARER_TOKEN` hook callback will interfere with nonblocking connection APIs such as `PQconnectPoll` and prevent concurrent connections from making progress. Applications which only ever use the synchronous connection primitives, such as `PQconnectdb`, may synchronously retrieve a token during the hook instead of implementing the *async* callback, but they will necessarily be limited to one connection at a time.) This will be called to begin the flow immediately upon return from the hook. When the callback cannot make further progress without blocking, it should return either `PGRES_POLLING_READING` or `PGRES_POLLING_WRITING` after setting `*altsock` to the file descriptor that will be marked ready to read/write when progress can be made again. (This descriptor is then provided to the top-level polling loop via `PQsocket()`.) Return `PGRES_POLLING_OK` after setting *token* when the flow is complete, or `PGRES_POLLING_FAILED` to indicate failure.


     Implementations may wish to store additional data for bookkeeping across calls to the *async* and *cleanup* callbacks. The *user* pointer is provided for this purpose; libpq will not touch its contents and the application may use it at its convenience. (Remember to free any allocations during token cleanup.)
<a id="libpq-oauth-authdata-oauth-bearer-token-v2"></a>

`PQAUTHDATA_OAUTH_BEARER_TOKEN_V2`
:   *Available in PostgreSQL 19 and later.*


     Provides all the functionality of [`PQAUTHDATA_OAUTH_BEARER_TOKEN`](#libpq-oauth-authdata-oauth-bearer-token), as well as the ability to set custom error messages and retrieve the OAuth issuer ID that the client has trusted.


     *data* points to an instance of `PGoauthBearerRequestV2`:

    ```

    typedef struct
    {
        PGoauthBearerRequest v1;    /* see the PGoauthBearerRequest struct, above */

        /* Hook inputs (constant across all calls) */
        const char *issuer;            /* the issuer identifier (RFC 9207) in use */

        /* Hook outputs */
        const char *error;             /* hook-defined error message */
    } PGoauthBearerRequestV2;
    ```


     Applications must first use the *v1* struct member to implement the base API, as described [above](#libpq-oauth-authdata-oauth-bearer-token). libpq additionally guarantees that the `request` pointer passed to the *v1.async* and *v1.cleanup* callbacks may be safely cast to `(PGoauthBearerRequestV2 *)`, to make use of the additional members described below.


    !!! warning

        Casting to `(PGoauthBearerRequestV2 *)` is *only* safe when the hook type is `PQAUTHDATA_OAUTH_BEARER_TOKEN_V2`. Applications may crash or misbehave if a hook implementation attempts to access v2 members when handling a v1 (`PQAUTHDATA_OAUTH_BEARER_TOKEN`) hook request.


     In addition to the functionality of the version 1 API, the v2 struct provides an additional input and output for the hook:


     *issuer* contains the issuer identifier, as defined in [RFC 9207](https://datatracker.ietf.org/doc/html/rfc9207), that is in use for the current connection. This identifier is derived from [oauth_issuer](database-connection-control-functions.md#libpq-connect-oauth-issuer). To avoid mix-up attacks, custom flows should ensure that any discovery metadata provided by the authorization server matches this issuer ID.


     *error* may be set to point to a custom error message when a flow fails. The message will be included as part of [PQerrorMessage](connection-status-functions.md#libpq-PQerrorMessage). Hooks must free any error message allocations during the *v1.cleanup* callback.

   <a id="libpq-oauth-debugging"></a>

### Debugging and Developer Settings


 While developing against a local authorization server, it may be helpful to make use of the [oauth_ca_file](database-connection-control-functions.md#libpq-connect-oauth-ca-file) connection parameter (or the equivalent `PGOAUTHCAFILE` environment variable) in the client application.


 Debug features may be enabled by setting the `PGOAUTHDEBUG` environment variable. This functionality is provided for ease of local development and testing. The variable accepts a comma-separated list of debug options:

```

PGOAUTHDEBUG=option1,option2,...         for safe options only
PGOAUTHDEBUG=UNSAFE:option1,option2,...  when using unsafe options
PGOAUTHDEBUG=UNSAFE                      legacy format; enables all options

```


 Available debug options:

`http` (unsafe)
:   Permits the use of unencrypted HTTP during the OAuth provider exchange. This allows OAuth credentials to be transmitted over unencrypted connections, which is extremely dangerous and should only be used for local testing.

`trace` (unsafe)
:   Prints HTTP traffic to standard error during the OAuth flow. This output contains critical secrets including bearer tokens, client secrets, access tokens, and authorization codes. Never share this output with third parties.

`dos-endpoint` (unsafe)
:   Permits the use of zero-second retry intervals instead of the normal minimum of one second. This speeds up tests, but in normal operation it will cause the client to busy-loop, consuming CPU and network resources.

`call-count` (safe)
:   Prints the total number of calls to the flow plugin to standard error when the OAuth flow completes. This helps developers debug the async callback behavior.

`plugin-errors` (safe)
:   Prints plugin loading errors to standard error. This helps developers and package maintainers debug issues when the OAuth plugin fails to load.


 Unsafe options (`http`, `trace`, `dos-endpoint`) require the `UNSAFE:` prefix. If unsafe options are specified without this prefix, or if an option name is unrecognized, a warning is printed to standard error and that option is ignored. Other valid options in the list continue to work. Safe options (`call-count`, `plugin-errors`) can be used without the prefix.


 Examples:

```

PGOAUTHDEBUG=call-count              safe options only
PGOAUTHDEBUG=UNSAFE:http,trace       enable HTTP and traffic logging
PGOAUTHDEBUG=UNSAFE:http,call-count  mix of unsafe and safe

```


!!! warning

    Never use unsafe debug options in production environments. They expose secrets and behaviors that can be used to attack your clients and servers. Do not share `trace` output with third parties.
