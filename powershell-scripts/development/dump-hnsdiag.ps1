param(
    [Parameter(Mandatory=$false)][String]$Extension="dump"
)

Write-Host "Checking endpoints"
hnsdiag list endpoints -d | out-file hns-endpoints-$Extension.txt

Write-Host "Checking Load Balancers"
hnsdiag list loadbalancers -d | out-file hns-loadbalancers-$Extension.txt

Write-Host "Checking Namespaces"
hnsdiag list namespaces -d | out-file hns-namespaces-$Extension.txt

Write-Host "Checking Networks"
hnsdiag list networks -d | out-file hns-networks-$Extension.txt
