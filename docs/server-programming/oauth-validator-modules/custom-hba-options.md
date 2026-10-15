<a id="oauth-validator-hba"></a>

## Custom HBA Options


 Like other preloaded libraries, validator modules may define [custom GUC parameters](../../server-administration/server-configuration/customized-options.md#runtime-config-custom) for user configuration in `postgresql.conf`. However, it may be desirable to configure behavior at a more granular level (say, for a particular issuer or a group of users) instead of globally.


 Beginning in PostgreSQL 19, validator implementations may define custom options for use inside `pg_hba.conf`. These options are then [made available](../../server-administration/client-authentication/oauth-authorization-authentication.md#auth-oauth-validator-option) to the user as <code>validator.</code><em>option</em>. The API for registering and retrieving custom options is described below.
 <a id="oauth-validator-hba-api"></a>

### Options API


 Modules register custom HBA option names during the `startup_cb` callback, using `RegisterOAuthHBAOptions()`:

```

/*
 * Register a list of custom option names for use in pg_hba.conf. For each name
 * "foo" registered here, that option will be provided as "validator.foo" in
 * the HBA.
 *
 * Valid option names consist of alphanumeric ASCII, underscore (_), and hyphen
 * (-). Invalid option names will be ignored with a WARNING logged at
 * connection time.
 *
 * This function may only be called during the startup_cb callback. Multiple
 * calls are permitted, which will append to the existing list of registered
 * options; options cannot be unregistered.
 *
 * Parameters:
 *
 * - state: the state pointer passed to the startup_cb callback
 * - num:   the number of options in the opts array
 * - opts:  an array of null-terminated option names to register
 *
 * The list of option names is copied internally, and the opts array is not
 * required to remain valid after the call.
 */
void RegisterOAuthHBAOptions(ValidatorModuleState *state, int num,
                             const char *opts[]);
```


 Each option's value, if set, may be later retrieved using `GetOAuthHBAOption()`:

```

/*
 * Retrieve the string value of an HBA option which was registered via
 * RegisterOAuthHBAOptions(). Usable only during validate_cb or shutdown_cb.
 *
 * If the user has set the corresponding option in pg_hba.conf, this function
 * returns that value as a null-terminated string, which must not be modified
 * or freed. NULL is returned instead if the user has not set this option, if
 * the option name was not registered, or if this function is incorrectly called
 * during the startup_cb.
 *
 * Parameters:
 *
 * - state:   the state pointer passed to the validate_cb/shutdown_cb callback
 * - optname: the name of the option to retrieve
 */
const char *GetOAuthHBAOption(const ValidatorModuleState *state,
                              const char *optname);
```


 See [Example Usage](#oauth-validator-hba-example-usage) for sample usage.
  <a id="oauth-validator-hba-limitations"></a>

### Limitations


-  Option names are limited to ASCII alphanumeric characters, underscores (`_`), and hyphens (`-`).
-  Option values are always freeform strings (in contrast to custom GUCs, which support numerics, booleans, and enums).
-  Option names and values cannot be checked by the server during a reload of the configuration. Any unregistered options in `pg_hba.conf` will instead result in connection failures. It is the responsibility of each module to document and verify the syntax of option values as needed.  (If a module finds an invalid option value during `validate_cb`, it's recommended to [signal an internal error](oauth-validator-callbacks.md#oauth-validator-callback-validate) by setting `result->error_detail` to a description of the problem and returning `false`.)

  <a id="oauth-validator-hba-example-usage"></a>

### Example Usage


 For a hypothetical module, the options `foo` and `bar` could be registered as follows:

```

static void
validator_startup(ValidatorModuleState *state)
{
    static const char *opts[] = {
        "foo",      /* description of access privileges */
        "bar",      /* magic URL for additional administrator powers */
    };

    RegisterOAuthHBAOptions(state, lengthof(opts), opts);

    /* ...other setup... */
}
```


 The following sample entries in `pg_hba.conf` can then make use of these options:

```

# TYPE   DATABASE   USER   ADDRESS    METHOD
hostssl  postgres   admin  0.0.0.0/0  oauth issuer=https://admin.example.com \
                                            scope="pg-admin openid email" \
                                            map=oauth-email \
                                            validator.foo="admin access" \
                                            validator.bar=https://magic.example.com

hostssl  postgres   all    0.0.0.0/0  oauth issuer=https://www.example.com \
                                            scope="pg-user openid email" \
                                            map=oauth-email \
                                            validator.foo="user access"
```


 The module can retrieve the option settings from the HBA during validation:

```

static bool
validate_token(const ValidatorModuleState *state,
               const char *token, const char *role,
               ValidatorModuleResult *res)
{
    const char *foo = GetOAuthHBAOption(state, "foo"); /* "admin access" or "user access" */
    const char *bar = GetOAuthHBAOption(state, "bar"); /* "https://magic.example.com" or NULL */

    if (bar && !is_valid_url(bar))
    {
        res->error_detail = psprintf("validator.bar (\"%s\") is not a valid URL.", bar);
        return false;
    }

    /* proceed to validate token */
}
```


 When multiple validators are in use, their registered option lists remain independent:

```

in postgresql.conf:
oauth_validator_libraries = 'example_org, my_validator'

in pg_hba.conf:
# TYPE   DATABASE   USER   ADDRESS    METHOD
hostssl  postgres   admin  0.0.0.0/0  oauth issuer=https://admin.example.com \
                                            scope="pg-admin openid email" \
                                            map=oauth-email \
                                            validator=my_validator \
                                            validator.foo="admin access" \
                                            validator.bar=https://magic.example.com

hostssl  postgres   all    0.0.0.0/0  oauth issuer=https://www.example.org \
                                            scope="pg-user openid profile" \
                                            validator=example_org \
                                            delegate_ident_mapping=1 \
                                            validator.magic=on \
                                            validator.more_magic=off
```
