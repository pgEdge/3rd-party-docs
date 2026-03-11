<a id="client-authentication"></a>

# Client Authentication

 When a client application connects to the database server, it specifies which PostgreSQL database user name it wants to connect as, much the same way one logs into a Unix computer as a particular user. Within the SQL environment the active database user name determines access privileges to database objects — see [Database Roles](../database-roles/index.md#user-manag) for more information. Therefore, it is essential to restrict which database users can connect.

!!! note

    As explained in [Database Roles](../database-roles/index.md#user-manag), PostgreSQL actually does privilege management in terms of “roles”. In this chapter, we consistently use *database user* to mean “role with the `LOGIN` privilege”.

 *Authentication* is the process by which the database server establishes the identity of the client, and by extension determines whether the client application (or the user who runs the client application) is permitted to connect with the database user name that was requested.

 PostgreSQL offers a number of different client authentication methods. The method used to authenticate a particular client connection can be selected on the basis of (client) host address, database, and user.

 PostgreSQL database user names are logically separate from user names of the operating system in which the server runs. If all the users of a particular server also have accounts on the server's machine, it makes sense to assign database user names that match their operating system user names. However, a server that accepts remote connections might have many database users who have no local operating system account, and in such cases there need be no connection between database user names and OS user names.

- [The `pg_hba.conf` File](the-pg_hba-conf-file.md#auth-pg-hba-conf)
- [User Name Maps](user-name-maps.md#auth-username-maps)
- [Authentication Methods](authentication-methods.md#auth-methods)
- [Trust Authentication](trust-authentication.md#auth-trust)
- [Password Authentication](password-authentication.md#auth-password)
- [GSSAPI Authentication](gssapi-authentication.md#gssapi-auth)
- [SSPI Authentication](sspi-authentication.md#sspi-auth)
- [Ident Authentication](ident-authentication.md#auth-ident)
- [Peer Authentication](peer-authentication.md#auth-peer)
- [LDAP Authentication](ldap-authentication.md#auth-ldap)
- [RADIUS Authentication](radius-authentication.md#auth-radius)
- [Certificate Authentication](certificate-authentication.md#auth-cert)
- [PAM Authentication](pam-authentication.md#auth-pam)
- [BSD Authentication](bsd-authentication.md#auth-bsd)
- [OAuth Authorization/Authentication](oauth-authorization-authentication.md#auth-oauth)
- [Authentication Problems](authentication-problems.md#client-authentication-problems)
