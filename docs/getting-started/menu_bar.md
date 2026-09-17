<a id="menu_bar"></a>

# Menu Bar

The pgAdmin menu bar provides drop-down menus for access to options, commands, and utilities. Its layout adapts based on your deployment: in Web mode, you will see *File*, *Object*, *Tools*, and *Help* menus. In Desktop mode, the *pgAdmin4* application menu replaces the standard *File* menu. Selections may be grayed out which indicates they are disabled for the object currently selected in the *pgAdmin* tree control.

# The File Menu

![pgAdmin file menu bar](../images/file_menu.png)

Use the *File* menu available in web mode to access the following options:

| Option | Action |
|---|---|
| *Preferences* | Click to open the [Preferences](preferences.md#preferences) dialog to customize your pgAdmin settings. |
| *Reset Layout* | If you have modified the workspace, click to restore the default layout. |

# The pgAdmin4 Menu

![pgAdmin pgadmin4 menu bar](../images/pgadmin4_menu.png)

Use the *pgAdmin4* menu available in desktop mode to access the following options:

| Option | Action |
|---|---|
| *About pgAdmin 4* | Provide pgadmin4 configuration information like version, mode etc. |
| *Preferences/settings* | Click to open the [Preferences](preferences.md#preferences) dialog to customize your pgAdmin settings. |
| *Reset Layout* | If you have modified the workspace, click to restore the default layout. |
| *View Logs* | This will show current pgadmin4 logs. |
| *Configure runtime* | Click to open a  window that allows to configure application port, connection timeout to launch the application. To know more about runtime menu [click here](deployment/desktop_deployment.md#desktop_deployment) |

# The Object Menu

![pgAdmin object menu bar](../images/object_menu.png)

The *Object* menu is context-sensitive. Use the *Object* menu to access the following options (in alphabetical order):

<table>
<thead>
<tr>
  <th>Option</th>
  <th>Action</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>Register</em>1) <em>Server</em>2) <em>Deploy Cloud Instance</em></td>
  <td>Click to open the <a href="../connecting-to-a-server/server_dialog.md#server_dialog">Server</a> dialog to register a server.Click to open the <a href="../connecting-to-a-server/cloud_deployment.md#cloud_deployment">Cloud Deployment</a> dialog to deploy an cloud instance.</td>
</tr>
<tr>
  <td><em>Change Password...</em></td>
  <td>Click to open the <a href="../management-basics/change_password_dialog.md#change_password_dialog">Change Password...</a> dialog to change your password.</td>
</tr>
<tr>
  <td><em>Clear Saved Password</em></td>
  <td>If you have saved the database server password, click to clear the saved password. Enable only when password is already saved.</td>
</tr>
<tr>
  <td><em>Clear SSH Tunnel Password</em></td>
  <td>If you have saved the ssh tunnel password, click to clear the saved password. Enable only when password is already saved.</td>
</tr>
<tr>
  <td><em>Connect Server</em></td>
  <td>Click to open the <a href="../connecting-to-a-server/connect_to_server.md#connect_to_server">Connect to Server</a> dialog to establish a connection with a server.</td>
</tr>
<tr>
  <td><em>Copy Server...</em></td>
  <td>Click to copy the currently selected server.</td>
</tr>
<tr>
  <td><em>Create</em></td>
  <td>Click <em>Create</em> to access a context menu that provides context-sensitive selections. Your selection opens a <em>Create</em> dialog for creating a new object.</td>
</tr>
<tr>
  <td><em>Drop</em></td>
  <td>Click to drop the currently selected object from the server.</td>
</tr>
<tr>
  <td><em>Drop (Cascade)</em></td>
  <td>Click to drop the currently selected object and all dependent objects from the server.</td>
</tr>
<tr>
  <td><em>Drop (Force)</em></td>
  <td>Click to drop the currently selected database with force option.</td>
</tr>
<tr>
  <td><em>Disconnect from server</em></td>
  <td>Click to disconnect from the currently selected server.</td>
</tr>
<tr>
  <td><em>Properties...</em></td>
  <td>Click to review or modify the currently selected object's properties.</td>
</tr>
<tr>
  <td><em>Refresh</em></td>
  <td>Click to refresh the currently selected object.</td>
</tr>
<tr>
  <td><em>Remove Server</em></td>
  <td>Click to remove the currently selected server.</td>
</tr>
<tr>
  <td><em>Scripts</em></td>
  <td>Click to open the <a href="../developer-tools/query_tool.md#query_tool">Query tool</a> to edit or view the selected script from the flyout menu.</td>
</tr>
<tr>
  <td><em>Trigger(s)</em></td>
  <td>Click to <em>Disable</em> or <em>Enable</em> trigger(s) for the currently selected table. Options are displayed on the flyout menu.</td>
</tr>
<tr>
  <td><em>Truncate</em></td>
  <td>Click to remove all rows from a table/foreign tables (<em>Truncate</em>), to remove all rows from a table/foreign tables and its child tables (<em>Truncate Cascade</em>) or to remove all rows from a table/foreign tables and automatically restart sequences owned by columns (<em>Truncate Restart Identity</em>). Options are displayed on the flyout menu.</td>
</tr>
<tr>
  <td><em>View Data</em></td>
  <td>Click to access a context menu that provides several options for viewing data (see below).</td>
</tr>
<tr>
  <td><em>ERD For Database</em></td>
  <td>Click to open the ERD tool with automatically generated diagram for the database selected. This option is available only when a database is selected. Options are displayed on the flyout menu.</td>
</tr>
<tr>
  <td><em>ERD For Table</em></td>
  <td>Click to open the ERD tool with automatically generated diagram for the table selected. This option is available only when a table is selected. Options are displayed on the flyout menu.</td>
</tr>
</tbody>
</table>

# The Tools Menu

![pgAdmin tools menu bar](../images/tool_menu.png)

Use the *Tools* menu to access the following options (in alphabetical order):

<table>
<thead>
<tr>
  <th>Option</th>
  <th>Action</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>ERD Tool</em></td>
  <td>Click to open the <a href="../developer-tools/erd_tool.md#erd_tool">ERD Tool</a> and start designing your database.</td>
</tr>
<tr>
  <td><em>Grant Wizard...</em></td>
  <td>Click to access the <a href="../management-basics/grant_wizard.md#grant_wizard">Grant Wizard</a> tool.</td>
</tr>
<tr>
  <td><em>PSQL Tool</em></td>
  <td>Click to open the <a href="../developer-tools/psql_tool.md#psql_tool">PSQL Tool</a> and start PSQL in the current database context.</td>
</tr>
<tr>
  <td><em>Query tool</em></td>
  <td>Click to open the <a href="../developer-tools/query_tool.md#query_tool">Query tool</a> for the currently selected object.</td>
</tr>
<tr>
  <td><em>Schema Diff</em></td>
  <td>Click to open the <a href="../developer-tools/schema_diff.md#schema_diff_feature">Schema Diff</a> and start comparing two database or two schema.</td>
</tr>
<tr>
  <td><em>Backup Globals...</em></td>
  <td>Click to open the <a href="../backup-and-restore/backup_globals_dialog.md#backup_globals_dialog">Backup Globals...</a> dialog to backup cluster objects.</td>
</tr>
<tr>
  <td><em>Backup Server...</em></td>
  <td>Click to open the <a href="../backup-and-restore/backup_server_dialog.md#backup_server_dialog">Backup Server...</a> dialog to backup a server.</td>
</tr>
<tr>
  <td><em>Backup...</em></td>
  <td>Click to open the <a href="../backup-and-restore/backup_dialog.md#backup_dialog">Backup...</a> dialog to backup database objects.</td>
</tr>
<tr>
  <td><em>Restore...</em></td>
  <td>Click to access the <a href="../backup-and-restore/restore_dialog.md#restore_dialog">Restore</a> dialog to restore database files from a backup.</td>
</tr>
<tr>
  <td><em>Export Data Using Query...</em></td>
  <td>Click to open the <a href="../management-basics/export_data_using_query.md#export_data_using_query">Export Data Using Query...</a> dialog to export data from a table using query.</td>
</tr>
<tr>
  <td><em>Import/Export Data...</em></td>
  <td>Click to open the <a href="../management-basics/import_export_data.md#import_export_data">Import/Export data...</a> dialog to import or export data from a table.</td>
</tr>
<tr>
  <td><em>Maintenance...</em></td>
  <td>Click to open the <a href="../management-basics/maintenance_dialog.md#maintenance_dialog">Maintenance...</a> dialog to VACUUM, ANALYZE, REINDEX, or CLUSTER.</td>
</tr>
<tr>
  <td><em>Search Objects...</em></td>
  <td>Click to open the <a href="search_objects.md#search_objects">Search Objects...</a> and start searching any kind of objects in a database.</td>
</tr>
<tr>
  <td><em>AI Reports</em></td>
  <td>Click to access a submenu with AI-powered analysis options (requires <a href="../developer-tools/ai_tools.md#ai_tools">AI configuration</a>):<ul><li><em>Security Report</em> - Generate an AI-powered security analysis for the selected server, database, or schema.</li><li><em>Performance Report</em> - Generate an AI-powered performance analysis for the selected server or database.</li><li><em>Design Report</em> - Generate an AI-powered design review for the selected database or schema.</li></ul></td>
</tr>
<tr>
  <td><em>Add named restore point</em></td>
  <td>Click to open the <a href="../management-basics/add_restore_point_dialog.md#add_restore_point_dialog">Add named restore point...</a> dialog to take a point-in-time snapshot of the current server state.</td>
</tr>
<tr>
  <td><em>Pause replay of WAL</em></td>
  <td>Click to pause the replay of the WAL log.</td>
</tr>
<tr>
  <td><em>Resume replay of WAL</em></td>
  <td>Click to resume the replay of the WAL log.</td>
</tr>
<tr>
  <td><em>Reload Configuration...</em></td>
  <td>Click to update configuration files without restarting the server.</td>
</tr>
<tr>
  <td><em>Storage Manager</em></td>
  <td>Click to open the <a href="../management-basics/storage_manager.md#storage_manager">Storage Manager</a> to upload, delete, or download the backup files.</td>
</tr>
</tbody>
</table>

# The Help Menu

![pgAdmin help menu bar](../images/help_menu.png)

Use the options on the *Help* menu to access online help documents, or to review information about the pgAdmin installation (in alphabetical order):

<table>
<thead>
<tr>
  <th>Option</th>
  <th>Action</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>Quick Search</em></td>
  <td>Type your keywords in the Quick Search field. Typing at least three characters will display all the matching possibilities under Menu items and the relevant documents under Help articles. Click on the options under Menu items to perform action of particular functionality or object. Click on any of the Help articles to open the help of that topic with highlighted text in a separate window.<strong>Note</strong>:- If any of the option under Menu items is disabled, then it will provide information via info icon.</td>
</tr>
<tr>
  <td><em>About pgAdmin 4</em></td>
  <td>Click to open a window where you will find information about pgAdmin; this includes the current version and the current user.</td>
</tr>
<tr>
  <td><em>Online Help</em></td>
  <td>Click to open documentation support for using pgAdmin utilities, tools and dialogs. Navigate (in the newly opened tab?) help documents in the left browser pane or use the search bar to specify a topic.</td>
</tr>
<tr>
  <td><em>pgAdmin Website</em></td>
  <td>Click to open the <em>pgAdmin.org</em> website in a browser window.</td>
</tr>
<tr>
  <td><em>PostgreSQL Website</em></td>
  <td>Click to access the PostgreSQL core documentation hosted at the PostgreSQL site. The site also offers guides, tutorials, and resources.</td>
</tr>
</tbody>
</table>
