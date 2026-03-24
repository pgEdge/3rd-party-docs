# Detailed output of each function

## pg_sys_os_info
- Name
- Version
- Host name
- Domain name
- Handle count
- Process count
- Thread count
- Architecture
- Last bootup time
- Uptime in seconds
    
## pg_sys_cpu_info
- Vendor
- Description
- Model name
- Processor type
- Logical processor
- Physical processor
- Number of cores
- Architecture
- Clock speed in hz
- CPU type
- CPU family
- Byte order
- L1d cache size
- L1i cache size
- L2 cache size
- L3 cache size

## pg_sys_cpu_usage_info
- Percent time spent in processing usermode normal process
- Percent time spent in processing usermode niced process
- Percent time spent in kernel mode process
- Percent time spent in idle mode
- Percent time spent in io completion
- Percent time spent in servicing interrupt
- Percent time spent in servicing software interrupt
- Percent user time spent
- Percent processor time spent
- Percent privileged time spent
- Percent interrupt time spent
    
## pg_sys_memory_info
- Total memory
- Used memory
- Free memory
- Total swap memory
- Used swap memory
- Free swap memory
- Total cache memory
- Total kernel memory
- Kernel paged memory
- Kernel non paged memory
- Total page file
- Available page file

## pg_sys_io_analysis_info
- Block device name
- Total number of reads
- Total number of writes
- Read bytes
- Written bytes
- Time spent in milliseconds for reading
- Time spent in milliseconds for writing

## pg_sys_disk_info
- File system of the disk
- File system type
- Mount point for the file system
- Drive letter
- Drive type
- Total space in bytes
- Used space in bytes
- Available space in bytes
- Number of total inodes
- Number of used inodes
- Number of free inodes

## pg_sys_load_avg_info
- 1 minute load average
- 5 minute load average
- 10 minute load average
- 15 minute load average

## pg_sys_process_info
- Number of total processes
- Number of running processes
- Number of sleeping processes
- Number of stopped processes
- Number of zombie processes

## pg_sys_network_info
- Name of the interface_name
- ipv4 address of the interface
- Number of total bytes transmitted
- Number of total packets transmitted
- Number of transmit errors by this network device
- Number of packets dropped during transmission
- Number of total bytes received
- Number of total packets received
- Number of receive errors by this network device
- Number of packets dropped by this network device
- Interface speed in mbps

## pg_sys_cpu_memory_by_process
- PID of the process
- Process name
- CPU usage in percent time spent on CPU
- Memory usage in bytes
- Total memory used in bytes

