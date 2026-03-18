# Azure-Compatible Object Store Support
<a name="azure-support"></a>

pgBackRest supports locating repositories in Azure-compatible object stores. The container used to store the repository must be created in advance &mdash; pgBackRest will not do it automatically. The repository can be located in the container root (`/`) but it's usually best to place it in a subpath so object store logs or other data can also be stored in the container without conflicts.

!!! warning

    Do not enable "hierarchical namespace" as this will cause errors during expire.

**Configure Azure**

```ini
[global]
repo{[azure-setup-repo-id]}-type=azure
[global]
repo{[azure-setup-repo-id]}-path=/demo-repo
[global]
repo{[azure-setup-repo-id]}-azure-account=pgbackrest
[global]
repo{[azure-setup-repo-id]}-azure-key-type=shared
[global]
repo{[azure-setup-repo-id]}-azure-key=YXpLZXk=
[global]
repo{[azure-setup-repo-id]}-azure-container=demo-container
[global]
repo{[azure-setup-repo-id]}-retention-full=4
```

Shared access signatures may be used by setting the `repo<azure-setup-repo-id>-azure-key-type` option to `sas` and the `repo<azure-setup-repo-id>-azure-key` option to the shared access signature token.

Commands are run exactly as if the repository were stored on a local disk.

**Create the stanza**

```bash
pgbackrest --stanza=demo --log-level-console=info stanza-create
```

File creation time in Azure is relatively slow so `backup`/`restore` performance is improved by enabling [file bundling](#/backup/bundle).

**Backup the demo cluster**

```bash
pgbackrest --stanza=demo --repo=2
                    --log-level-console=info backup
```
