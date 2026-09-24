<a id="basic-archive"></a>

## basic_archive — an example WAL archive module


 `basic_archive` is an example of an archive module. This module copies completed WAL segment files to the specified directory. This may not be especially useful, but it can serve as a starting point for developing your own archive module. For more information about archive modules, see [Archive Modules](../../server-programming/archive-modules/index.md#archive-modules).


 In order to function, this module must be loaded via [archive_library](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-library), and [archive_mode](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-mode) must be enabled.
 <a id="basic-archive-configuration-parameters"></a>

### Configuration Parameters


`basic_archive.archive_directory` (`string`)
:   The directory where the server should copy WAL segment files. This directory must already exist. The default is an empty string, which effectively halts WAL archiving, but if [archive_mode](../../server-administration/server-configuration/write-ahead-log.md#guc-archive-mode) is enabled, the server will accumulate WAL segment files in the expectation that a value will soon be provided.


 These parameters must be set in `postgresql.conf`. Typical usage might be:


```

# postgresql.conf
archive_mode = 'on'
archive_library = 'basic_archive'
basic_archive.archive_directory = '/path/to/archive/directory'
```
  <a id="basic-archive-notes"></a>

### Notes


 Server crashes may leave temporary files with the prefix `archtemp` in the archive directory. It is recommended to delete such files before restarting the server after a crash. It is safe to remove such files while the server is running as long as they are unrelated to any archiving still in progress, but users should use extra caution when doing so.
  <a id="basic-archive-author"></a>

### Author


 Nathan Bossart
