<a id="query_tool_toolbar"></a>

# Query Tool Toolbar

The *Query Tool* toolbar uses context-sensitive icons that provide shortcuts to frequently performed tasks. If an icon is highlighted, the option is enabled; if the icon is grayed-out, the task is disabled.

!!! note

    The [Query Tool](../query_tool.md#query_tool) and [View/Edit Data](../editgrid.md#editgrid) tools are actually different operating modes of the same tool. Some controls will be disabled in either mode.

![Query tool toolbar](../../images/query_toolbar.png)

Hover over an icon in pgAdmin to display a tooltip that describes the icon's functionality.

# File Options

<table>
<thead>
<tr>
  <th>Icon</th>
  <th>Behavior</th>
  <th>Shortcut</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>Open File</em></td>
  <td>Click the <em>Open File</em> icon to display a previously saved query in the same tab of the SQL Editor. To open the file in a new tab, select <em>Open in a new tab?</em> option from the dropdown.</td>
  <td>Cmd/Ctrl + O</td>
</tr>
<tr>
  <td><em>Save File</em></td>
  <td>Click the <em>Save</em> icon to perform a quick-save of a previously saved query, or to access the <em>Save</em> menu:<ul><li>Select <em>Save</em> to save the selected content of the SQL Editor panel in a  file.</li><li>Select <em>Save As</em> to open a new browser dialog and specify a new location to which to save the selected content of the SQL Editor panel.</li></ul></td>
  <td>Cmd/Ctrl + S</td>
</tr>
</tbody>
</table>

# Filter/Limit Options

<table>
<thead>
<tr>
  <th>Icon</th>
  <th>Behavior</th>
  <th>Shortcut</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>Filter</em></td>
  <td>Click the <em>Filter</em> icon to set filtering and sorting criteria for the data when in <em>View/Edit data mode</em>. Click the down arrow to access other filtering and sorting options:<ul><li>In the <em>SQL Filter</em>, you can enter a SQL query as filtering criteria. In <em>Data Sorting</em>, you can select the column and specify the order for sorting.</li><li>Click <em>Filter by Selection</em> to show only the rows containing the values in the selected cells.</li><li>Click <em>Exclude by Selection</em> to show only the rows that do not contain the values in the selected cells.</li><li>Click <em>Remove Sort/Filter</em> to remove any previously selected sort or filtering options.</li></ul></td>
  <td>Option/Alt + F</td>
</tr>
<tr>
  <td>Limit Selector</td>
  <td>Select a value in the <em>Limit Selector</em> to limit the size of the dataset to a number of rows.</td>
  <td>Option/Alt + R</td>
</tr>
</tbody>
</table>

# Query Editing Options

![Query tool editing options](../../images/query_editing.png)

| Icon | Behavior | Shortcut |
|---|---|---|
| *Edit* | Use the *Edit* menu to search, replace, or navigate the code displayed in the SQL Editor: | Option/Alt + Shift + N |
|  | Select *Find* to provide a search target, and search the SQL Editor contents. | Cmd/Ctrl + F |
|  | Select *Replace* to locate and replace (with prompting) individual occurrences of the target. | Option + Cmd + F (MAC) Ctrl + Shift + F (Others) |
|  | Select *Go to Line/Column* to go to specified line number and column position | Cmd/Ctrl + L |
|  | Select *Indent Selection* to indent the currently selected text. | Tab |
|  | Select *Unindent Selection* to remove indentation from the currently selected text. | Shift + Tab |
|  | Select *Toggle Comment* to comment/uncomment any lines that contain the selection in SQL style. | Cmd/Ctrl + / |
|  | Select *Clear Query* to clear the query editor window. | Option/Alt + Ctrl + L |
|  | Select *Format SQL* to format the selected SQL or all the SQL if none is selected | Cmd/Ctrl + K |

# Query Execution

![Query tool execute options](../../images/query_execution.png)

<table>
<thead>
<tr>
  <th>Icon</th>
  <th>Behavior</th>
  <th>Shortcut</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>Stop</em></td>
  <td>Click the <em>Stop</em> icon to cancel the execution of the currently running query.</td>
  <td>Option + Shift + Q</td>
</tr>
<tr>
  <td><em>Execute script</em></td>
  <td>Click the <em>Execute script</em> icon to either execute or refresh the query highlighted in the SQL editor panel. Click the down arrow to access other execution options:<ul><li>Add a check next to <em>Auto rollback on error?</em> to instruct the server to automatically roll back a transaction if an error occurs during the transaction.</li><li>Add a check next to <em>Auto commit?</em> to instruct the server to automatically commit each transaction.  Any changes made by the transaction will be visible to others, and durable in the event of a crash.</li></ul></td>
  <td>F5</td>
</tr>
<tr>
  <td><em>Execute query</em></td>
  <td>Click the <em>Execute query</em> icon to either execute the query where the cursor is present or refresh the query highlighted in the SQL editor panel.</td>
  <td>Option+F5 (MAC) Alt+F5 (Others)</td>
</tr>
<tr>
  <td><em>Explain</em></td>
  <td>Click the <em>Explain</em> icon to view an explanation plan for the current query. The result of the EXPLAIN is displayed graphically on the <em>Explain</em> tab of the output panel, and in text form on the <em>Data Output</em> tab.</td>
  <td>F7</td>
</tr>
<tr>
  <td><em>Explain analyze</em></td>
  <td>Click the <em>Explain analyze</em> icon to invoke an EXPLAIN ANALYZE command on the current query.Navigate through the <em>Explain Options</em> menu to select options for the EXPLAIN command:<ul><li>Select <em>Buffers</em> to include information on buffer usage.</li><li>Select <em>Costs</em> to include information on the estimated startup and total cost of each plan node, as well as the estimated number of rows and the estimated width of each row.</li><li>Select <em>Generic Plan</em> to include the information on the Generic Plan.</li><li>Select <em>Memory</em> to include the information on memory consumption by the query planning phase.</li><li>Select <em>Serialize</em> to include information on the cost of serializing the query's output data, that is converting it to text or binary format to send to the client.</li><li>Select <em>Settings</em> to include the information on the configuration parameters.</li><li>Select <em>Summary</em> to include the summary information about the query plan.</li><li>Select <em>Timing</em> to include information about the startup time and the amount of time spent in each node of the query.</li><li>Select <em>Verbose</em> to display additional information regarding the query plan.</li><li>Select <em>Wal</em> to include the information on WAL record generation.</li></ul></td>
  <td>Shift + F7</td>
</tr>
<tr>
  <td><em>Commit</em></td>
  <td>Click the <em>Commit</em> icon to commit the transaction.</td>
  <td>Shift + Ctrl + M</td>
</tr>
<tr>
  <td><em>Rollback</em></td>
  <td>Click the <em>Rollback</em> icon to rollback the transaction.</td>
  <td>Shift + Ctrl + R</td>
</tr>
<tr>
  <td><em>Macros</em></td>
  <td>Click the <em>Macros</em> icon to manage the macros. You can create, edit or clear the macros through the <em>Manage Macros</em> option.</td>
  <td></td>
</tr>
</tbody>
</table>

# Data Editing Options

![Query tool data editing options](../../images/query_data_editing.png)

<table>
<thead>
<tr>
  <th>Icon</th>
  <th>Behavior</th>
  <th>Shortcut</th>
</tr>
</thead>
<tbody>
<tr>
  <td><em>Add row</em></td>
  <td>Click the <em>Add row</em> icon to add a new row</td>
  <td></td>
</tr>
<tr>
  <td><em>Copy</em></td>
  <td>Click the <em>Copy</em> icon to copy the content with or without header:<ul><li>Click the <em>Copy</em> icon to copy the content that is currently highlighted in the Data Output panel.</li><li> Click <em>Copy with headers</em> to copy the highlighted content along with the header.</li></ul></td>
  <td>Cmd/Ctrl + C</td>
</tr>
<tr>
  <td><em>Paste</em></td>
  <td>Click the <em>Paste</em> icon to paste a previously copied row with or without serial/identity values:<ul><li>Click the <em>Paste</em> icon to paste a previously copied row into a new row.</li><li>Click the <em>Paste with SERIAL/IDENTITY values?</em> if you want to paste the copied column values in the serial/identity columns.</li></ul>Note that copied row having <em>Bytea</em> datatype cell will be pasted as <em>Null</em>.</td>
  <td>Option/Alt + Shift + P</td>
</tr>
<tr>
  <td><em>Delete</em></td>
  <td>Click the <em>Delete</em> icon to mark the selected rows for deletion. These marked rows get deletedwhen you click the <em>Save Data Changes</em> icon.</td>
  <td>Option/Alt + Shift + D</td>
</tr>
<tr>
  <td><em>Save Data Changes</em></td>
  <td>Click the <em>Save Data Changes</em> icon to save data changes (insert, update, or delete) in the Data Output Panel to the server.</td>
  <td>F6</td>
</tr>
<tr>
  <td><em>Save results to</em> <em>file</em></td>
  <td>Click the Save results to file icon to save the result set of the current query as a delimited text file (CSV, if the field separator is set to a comma). This button will only be enabled when a query has been executed and there are results in the data grid. You can specify the CSV/TXT settings in the Preference Dialogue under SQL Editor -> CSV/TXT output.</td>
  <td>F8</td>
</tr>
<tr>
  <td>Graph Visualiser</td>
  <td>Use the Graph Visualiser button to generate graphs of the query results.</td>
  <td></td>
</tr>
<tr>
  <td>SQL</td>
  <td>Use the SQL button to check the current query that gave the data.</td>
  <td></td>
</tr>
</tbody>
</table>

# Pagination Options

![Query tool data pagination options](../../images/query_data_pagination.png)

| Icon | Behavior | Shortcut |
|---|---|---|
| *Rows Range* | Show the current row numbers visible in the data grid. |  |
| *Edit Range* | Click to open the from and to rows range inputs to allow setting them. |  |
| *Show Entire Range* | Click to get all the rows and set the from and to rows range |  |
| *Page No* | Enter the page no you want to jump to out of total shown next to this input |  |
| *First Page* | Click to go to the first page. |  |
| *Previous Page* | Click to go to the previous page. |  |
| *Next Page* | Click to go to the next page. |  |
| *Last Page* | Click to go to the last page. |  |

![Query tool data pagination options](../../images/query_data_pagination_edit.png)

One can click the edit range button to open rows range editor:

- From and to range should be between 1 and total rows.

- The range can be applied by clicking the *Apply* button or by pressing enter in the range inputs.

- Once the range is applied, pgAdmin will recalculate the rows per page. The pagination will then behave based on the new rows per page.

- It may be possible that on pressing next page button, the new rows range is not next to manually enterred range.

# Status Bar

![Query tool status bar](../../images/query_status_bar.png)

The status bar shows the following information:

- **Total rows**: The total number of rows returned by the query.

- **Query complete**: The time is taken by the query to complete.

- **Rows selected**: The number of rows selected in the data output panel.

- **Changes staged**: This information shows the number of rows added, deleted, and updated.

- **LF/CRLF**: It shows the end of line sequence to be used for the editor. When opening an empty editor, it will be decided based on OS.
    And when opening an existing file, it will be based on file end of lines. One can change the EOL by clicking on any of the options.

- **Ln**: In the Query tab, it is the line number at which the cursor is positioned.

- **Col**: In the Query tab, it is the column number at which the cursor is positioned
