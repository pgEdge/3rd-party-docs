# Architecture

The debugger consists of three parts:

1. The client. This is typically a GUI displays the source code, current
   stack frame, variables etc, and allows the user to set breakpoints and
   step throught the code. The client can reside on a different host than
   the database server.

2. The target backend. This is the backend that runs the code being debugged.
   The plugin_debugger.so library must be loaded into the target backend.

3. Debugging proxy. This is another backend process that the client is
   connected to. The API functions, pldbg_* in pldbgapi.so library, are
   run in this backend.

The client is to connected to the debugging proxy using a regular libpq
connection. When a debugging session is active, the proxy is connected
to the target via a socket. The protocol between the proxy and the target
backend is not visible to others, and is subject to change. The pldbg_*
API functions form the public interface to the debugging facility.

```
debugger client  *------ libpq --------* Proxy backend
  (pgAdmin)                                 *
                                            |
                                  pldebugger socket connection
                                            |
                                            *
application client *----- libpq -------* Target backend
```

