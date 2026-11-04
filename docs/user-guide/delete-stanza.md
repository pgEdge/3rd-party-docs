# Delete a Stanza
<a name="delete-stanza"></a>
`stanza-delete`

**Stop  cluster to be removed**

```bash
pg_ctlcluster 17 demo stop
```

**Stop  for the stanza**

```bash
pgbackrest --stanza=demo --log-level-console=info stop
```

**Delete the stanza from one repository**

```bash
pgbackrest --stanza=demo --repo=1
                    --log-level-console=info stanza-delete
```
