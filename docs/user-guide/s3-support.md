# S3-Compatible Object Store Support
<a name="s3-support"></a>

pgBackRest supports locating repositories in S3-compatible object stores. The bucket used to store the repository must be created in advance &mdash; pgBackRest will not do it automatically. The repository can be located in the bucket root (`/`) but it's usually best to place it in a subpath so object store logs or other data can also be stored in the bucket without conflicts.

**Configure S3**

```ini
[global]
repo{[s3-setup-repo-id]}-type=s3
[global]
repo{[s3-setup-repo-id]}-path=/demo-repo
[global]
repo{[s3-setup-repo-id]}-s3-key=accessKey1
[global]
repo{[s3-setup-repo-id]}-s3-key-secret=verySecretKey1
[global]
repo{[s3-setup-repo-id]}-s3-bucket=demo-bucket
[global]
repo{[s3-setup-repo-id]}-s3-endpoint=s3.us-east-1.amazonaws.com
[global]
repo{[s3-setup-repo-id]}-s3-region=us-east-1
[global]
repo{[s3-setup-repo-id]}-retention-full=4
```

!!! note

    The region and endpoint will need to be configured to where the bucket is located. The values given here are for the `us-east-1` region.

A role should be created to run pgBackRest and the bucket permissions should be set as restrictively as possible. If the role is associated with an instance in AWS then pgBackRest will automatically retrieve temporary credentials when `repo3-s3-key-type=auto`, which means that keys do not need to be explicitly set in `/etc/pgbackrest/pgbackrest.conf`.

This sample Amazon S3 policy will restrict all reads and writes to the bucket and repository path.

**Sample Amazon S3 Policy**

```

            {
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Effect": "Allow",
                        "Action": [
                            "s3:ListBucket"
                        ],
                        "Resource": [
                            "arn:aws:s3:::demo-bucket"
                        ],
                        "Condition": {
                            "StringEquals": {
                                "s3:prefix": [
                                    "",
                                    "demo-repo"
                                ],
                                "s3:delimiter": [
                                    "/"
                                ]
                            }
                        }
                    },
                    {
                        "Effect": "Allow",
                        "Action": [
                            "s3:ListBucket"
                        ],
                        "Resource": [
                            "arn:aws:s3:::demo-bucket"
                        ],
                        "Condition": {
                            "StringLike": {
                                "s3:prefix": [
                                    "demo-repo/*"
                                ]
                            }
                        }
                    },
                    {
                        "Effect": "Allow",
                        "Action": [
                            "s3:PutObject",
                            "s3:PutObjectTagging",
                            "s3:GetObject",
                            "s3:GetObjectVersion",
                            "s3:DeleteObject"
                        ],
                        "Resource": [
                            "arn:aws:s3:::demo-bucket/demo-repo/*"
                        ]
                    }
                ]
            }
        
```

Commands are run exactly as if the repository were stored on a local disk.

**Create the stanza**

```bash
pgbackrest --stanza=demo --log-level-console=info stanza-create
```

File creation time in S3 is relatively slow so `backup`/`restore` performance is improved by enabling [file bundling](backup.md#bundle).

**Backup the demo cluster**

```bash
pgbackrest --stanza=demo --repo=3
                    --log-level-console=info backup
```
