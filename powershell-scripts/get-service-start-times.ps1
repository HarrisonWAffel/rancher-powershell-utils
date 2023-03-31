# Credit: https://mikefrobbins.com/2017/11/30/determine-the-start-time-of-a-windows-service-with-powershell/

Get-CimInstance -ClassName Win32_Service -PipelineVariable Service |
        ForEach-Object {
            Get-Process -Id $_.ProcessId |
                    Select-Object -Property @{label='Status';expression={$Service.State}},
                    @{label='Name';expression={$Service.Name}},
                    @{label='DisplayName';expression={$Service.DisplayName}},
                    StartTime
        }