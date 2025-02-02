package terminal

// Windows command lines commands
var OpenConnections = []string{"cmd", "/c", "netstat"}

// PowerShell command lines commands
var OpenConnectionStatisticsPowerShell = []string{"powershell", "-Command", "Get-NetAdapterStatistics", "|", "Format-Table", "-AutoSize"}
var GetInterfaceNames = []string{"powershell", "-Command", "Get-NetAdapter", "|", "Select-Object", "Name"}

// Ubuntu command to check open connections (equivalent to netstat)
var OpenConnectionsLinux = []string{"ss", "-tuln"}

// Ubuntu command to get network adapter statistics (equivalent to Get-NetAdapterStatistics)
var OpenConnectionStatisticsLinux = []string{"ip", "-s", "link", "show"}

// Ubuntu command to get network interface names (equivalent to Get-NetAdapter)
var GetInterfaceNamesLinux = []string{"ip", "link", "show"}
