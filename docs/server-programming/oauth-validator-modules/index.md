<a id="oauth-validators"></a>

# OAuth Validator Modules

 PostgreSQL provides infrastructure for creating custom modules to perform server-side validation of OAuth bearer tokens. Because OAuth implementations vary so wildly, and bearer token validation is heavily dependent on the issuing party, the server cannot check the token itself; validator modules provide the integration layer between the server and the OAuth provider in use.

 OAuth validator modules must at least consist of an initialization function (see [Initialization Functions](initialization-functions.md#oauth-validator-init)) and the required callback for performing validation (see [Validate Callback](oauth-validator-callbacks.md#oauth-validator-callback-validate)).

!!! warning

    Since a misbehaving validator might let unauthorized users into the database, correct implementation is crucial for server safety. See [Safely Designing a Validator Module](safely-designing-a-validator-module.md#oauth-validator-design) for design considerations.

- [Safely Designing a Validator Module](safely-designing-a-validator-module.md#oauth-validator-design)
- [Initialization Functions](initialization-functions.md#oauth-validator-init)
- [OAuth Validator Callbacks](oauth-validator-callbacks.md#oauth-validator-callbacks)
