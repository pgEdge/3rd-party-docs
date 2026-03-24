# Troubleshooting

The majority of problems we've encountered with the plugin are caused by
failing to add (or incorrectly adding) the debugger plugin library to the
shared_preload_libraries configuration directive in postgresql.conf (following
which, the server *must* be restarted). This will prevent global breakpoints
working on all platforms, and on some (notably Windows) may prevent the 
pldbgapi.sql script from executing correctly.

