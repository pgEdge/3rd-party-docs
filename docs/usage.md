# Usage

Connect pgAdmin to the database containing the functions you wish to debug.
Right-click the function to debug, and select Debugging->Debug to execute and
debug the function immediately, or select Debugging->Set Global Breakpoint to
set a breakpoint on the function. This will cause the debugger to wait for
another session (such as a backend servicing a web app) to execute the function
and allow you to debug in-context.

For further information, please see the pgAdmin documentation.

