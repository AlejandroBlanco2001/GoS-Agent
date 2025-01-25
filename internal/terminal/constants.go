package terminal

// Windows command lines commands
var OpenConnections = []string{"cmd", "/c", "netstat"}

// PowerShell command lines commands
var OpenConnectionStatisticsPowerShell = []string{"powershell", "-Command", "Get-NetAdapterStatistics", "|", "Format-Table", "-AutoSize"}
var GetInterfaceNames = []string{"powershell", "-Command", "Get-NetAdapter", "|", "Select-Object", "Name"}
