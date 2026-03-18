<a id="tree_control"></a>

# Tree Control

The left pane of the main window displays a tree control (Object explorer) that provides access to the objects that reside on a server.

![object explorer panel](images/main_left_pane.png)

You can expand nodes in the tree control to view the database objects that reside on a selected server. The tree control expands to display a hierarchical view:

- Use the plus sign (+) to the left of a node to expand a segment of the tree
    control.

- Click the minus sign (-) to the left of a node to close that node.

You can also **drag and drop** certain objects to the Query Tool which can save time in typing long object names. Text containing the object name will be fully qualified with schema. Double quotes will be added if required. For functions and procedures, the function name along with parameter names will be pasted in the Query Tool.

Access context-sensitive menus by right-clicking on a node of the tree control to perform common tasks. Menus display options that include one or more of the following selections (options appear in alphabetical order):

| Option | Action |
|---|---|
| *Add named restore point* | Click to create and enter the name of a restore point. |
| *Backup...* | Click to open the [Backup...](../backup-and-restore/backup_dialog.md#backup_dialog) dialog to backup database objects. |
| *Backup Globals...* | Click to open the [Backup Globals...](../backup-and-restore/backup_globals_dialog.md#backup_globals_dialog) dialog to backup cluster objects. |
| *Backup Server...* | Click to open the [Backup Server...](../backup-and-restore/backup_server_dialog.md#backup_server_dialog) dialog to backup a server. |
| *Connect Server...* | Click to open the [Connect to Server](../connecting-to-a-server/connect_to_server.md#connect_to_server) dialog to establish a connection with a server. |
| *Create* | Click to access a context menu that provides context-sensitive selections. Your selection opens a *Create* dialog for creating a new object. |
| *CREATE Script* | Click to open the [Query tool](../developer-tools/query_tool.md#query_tool) to edit or view the CREATE script. |
| *Debugging* | Click through to open the [Debug](../developer-tools/debugger.md#debugger) tool or to select *Set breakpoint* to stop or pause a script execution. |
| *Drop* | Click to drop the currently selected object from the server. |
| *Drop (Cascade)* | Click to drop the currently selected object and all dependent objects from the server. |
| *Drop (Force)* | Click to drop the currently selected database with force option. |
| *Disconnect Database...* | Click to terminate a database connection. |
| *Disconnect from server* | Click to disconnect from the currently selected server. |
| *Debugging* | Click to access the [Debugger](../developer-tools/debugger.md#debugger) tool. |
| *Grant Wizard* | Click to access the [Grant Wizard](../management-basics/grant_wizard.md#grant_wizard) tool. |
| *Maintenance...* | Click to open the [Maintenance...](../management-basics/maintenance_dialog.md#maintenance_dialog) dialog to VACUUM, ANALYZE, REINDEX, or CLUSTER. |
| *Properties...* | Click to review or modify the currently selected object's properties. |
| *Refresh...* | Click to refresh the currently selected object. |
| *Reload Configuration...* | Click to update configuration files without restarting the server. |
| *Restore...* | Click to access the [Restore](../backup-and-restore/restore_dialog.md#restore_dialog) dialog to restore database files from a backup. |
| *View Data* | Use the *View Data* option to access the data stored in a selected table with the *Data Output* tab of the *Query Tool*. |

The context-sensitive menus associated with *Tables* and nested *Table* nodes provides additional display options (options appear in alphabetical order):

| Option | Action |
|---|---|
| *Import/Export Data...* | Click open the [Import/Export...](../management-basics/import_export_data.md#import_export_data) dialog to import data to or export data from the selected table. |
| *Reset Statistics* | Click to reset statistics for the selected table. |
| *Scripts* | Click to open the [Query tool](../developer-tools/query_tool.md#query_tool) to edit or view the selected script from the flyout menu. |
| *Truncate* | Click to remove all rows from a table. |
| *Truncate Cascade* | Click to remove all rows from a table and its child tables. |
| *View First 100 Rows* | Click to access a data grid that displays the first 100 rows of the selected table. |
| *View Last 100 Rows* | Click to access a data grid that displays the last 100 rows of the selected table. |
| *View All Rows* | Click to access a a data grid that displays all rows of the selected table. |
| *View Filtered Rows...* | Click to access the *Data Filter* popup to apply a filter to a set of data. |
