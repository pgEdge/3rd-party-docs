# GCS-Compatible Object Store Support
<a name="gcs-support"></a>

pgBackRest supports locating repositories in GCS-compatible object stores. The bucket used to store the repository must be created in advance &mdash; pgBackRest will not do it automatically. The repository can be located in the bucket root (`/`) but it's usually best to place it in a subpath so object store logs or other data can also be stored in the bucket without conflicts.

**Configure GCS**

```ini
[global]
repo{[gcs-setup-repo-id]}-type=gcs
[global]
repo{[gcs-setup-repo-id]}-path=/demo-repo
[global]
repo{[gcs-setup-repo-id]}-gcs-key-type=service
[global]
repo{[gcs-setup-repo-id]}-gcs-key=/etc/pgbackrest/gcs-key.json
[global]
repo{[gcs-setup-repo-id]}-gcs-bucket=demo-bucket
```

When running in GCE set `repo<gcs-setup-repo-id>-gcs-key-type=auto` to automatically authenticate using the instance service account.

Commands are run exactly as if the repository were stored on a local disk.

File creation time in GCS is relatively slow so `backup`/`restore` performance is improved by enabling [file bundling](user-guide/backup.md#bundle).
