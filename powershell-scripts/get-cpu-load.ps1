# gets a flat CPU load percentage

Get-WmiObject Win32_Processor | Select LoadPercentage | Format-List