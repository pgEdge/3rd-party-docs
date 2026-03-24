# Functions
The following functions are provided to fetch system level statistics for all
platforms.

## pg_sys_os_info
This interface allows the user to get operating system statistics.

## pg_sys_cpu_info
This interface allows the user to get CPU information.

## pg_sys_cpu_usage_info
This interface allows the user to get CPU usage information. Values are a
percentage of time spent by CPUs for all operations.

## pg_sys_memory_info
This interface allows the user to get memory usage information. All the values
are in bytes.

## pg_sys_io_analysis_info
This interface allows the user to get an I/O analysis of block devices.

## pg_sys_disk_info
This interface allows the user to get the disk information.

## pg_sys_load_avg_info
This interface allows the user to get the average load of the system over 1, 5,
10 and 15 minute intervals.

## pg_sys_process_info
This interface allows the user to get process information.

## pg_sys_network_info
This interface allows the user to get network interface information.

## pg_sys_cpu_memory_by_process
This interface allows the user to get the CPU and memory information for each
process ID.

NOTE: macOS does not allow access to to process information for other users.
      e.g. If the database server is running as the postgres user, this function
      will fetch information only for processes owned by the postgres user.
      Other processes will be listed and include only the process ID and name;
      other columns will be NULL.


