<a id="oauth-validator-callbacks"></a>

## OAuth Validator Callbacks


 OAuth validator modules implement their functionality by defining a set of callbacks. The server will call them as required to process the authentication request from the user.
 <a id="oauth-validator-callback-startup"></a>

### Startup Callback


 The `startup_cb` callback is executed directly after loading the module. This callback can be used to set up local state and perform additional initialization if required. If the validator module has state it can use `state->private_data` to store it.

```

typedef void (*ValidatorStartupCB) (ValidatorModuleState *state);
```

  <a id="oauth-validator-callback-validate"></a>

### Validate Callback


 The `validate_cb` callback is executed during the OAuth exchange when a user attempts to authenticate using OAuth. Any state set in previous calls will be available in `state->private_data`.

```

typedef bool (*ValidatorValidateCB) (const ValidatorModuleState *state,
                                     const char *token, const char *role,
                                     ValidatorModuleResult *result);
```
 *token* will contain the bearer token to validate. PostgreSQL has ensured that the token is well-formed syntactically, but no other validation has been performed. *role* will contain the role the user has requested to log in as. The callback must set output parameters in the `result` struct, which is defined as below:

```

typedef struct ValidatorModuleResult
{
    bool        authorized;
    char       *authn_id;
} ValidatorModuleResult;
```
 The connection will only proceed if the module sets `result->authorized` to `true`. To authenticate the user, the authenticated user name (as determined using the token) shall be palloc'd and returned in the `result->authn_id` field. Alternatively, `result->authn_id` may be set to NULL if the token is valid but the associated user identity cannot be determined.


 A validator may return `false` to signal an internal error, in which case any result parameters are ignored and the connection fails. Otherwise the validator should return `true` to indicate that it has processed the token and made an authorization decision.


 The behavior after `validate_cb` returns depends on the specific HBA setup. Normally, the `result->authn_id` user name must exactly match the role that the user is logging in as. (This behavior may be modified with a usermap.) But when authenticating against an HBA rule with `delegate_ident_mapping` turned on, PostgreSQL will not perform any checks on the value of `result->authn_id` at all; in this case it is up to the validator to ensure that the token carries enough privileges for the user to log in under the indicated *role*.
  <a id="oauth-validator-callback-shutdown"></a>

### Shutdown Callback


 The `shutdown_cb` callback is executed when the backend process associated with the connection exits. If the validator module has any allocated state, this callback should free it to avoid resource leaks.

```

typedef void (*ValidatorShutdownCB) (ValidatorModuleState *state);
```
