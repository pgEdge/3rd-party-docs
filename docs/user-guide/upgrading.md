# Upgrading pgBackRest
<a name="upgrading"></a>

## Upgrading pgBackRest from v2.x to v2.y
<a name="v2.x"></a>

Upgrading from v2.x to v2.y is straight-forward. The repository format has not changed, so for most installations it is simply a matter of installing binaries for the new version. It is also possible to downgrade if you have not used new features that are unsupported by the older version.

!!! important

    The local and remote pgBackRest versions must match exactly so they should be upgraded together. If there is a mismatch, WAL archiving and backups will not function until the versions match. In such a case, the following error will be reported: `[ProtocolError] expected value '2.x' for greeting key 'version' but got '2.y'`.
