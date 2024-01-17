# Written on a single line to make copy pasting into an RDP session a bit easier / cleaner

# n - number of lines that should be displayed for each refresh
param([Int32]$n=15)

while(1){get-eventlog -logname application -source rke2 | Select-Object ReplacementStrings | select -first $n | Format-Table -wrap; sleep 3; clear}
