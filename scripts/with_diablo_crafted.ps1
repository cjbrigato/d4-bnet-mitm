$here=pwd
$origin=$here.Path

cd "C:\Program Files (x86)\Diablo IV"
& '.\Diablo IV.exe'
cd $origin
.\d4-bnet-mitm.exe -crafted-auth-response true
