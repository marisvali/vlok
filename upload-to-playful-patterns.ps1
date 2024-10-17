$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'
go build -o vlok.wasm github.com/marisvali/vlok
Remove-Item Env:GOOS
Remove-Item Env:GOARCH

$client = New-Object System.Net.WebClient
$client.Credentials = New-Object System.Net.NetworkCredential($Env:VLOK_FTP_USER, $Env:VLOK_FTP_PASSWORD)
$client.UploadFile("ftp://ftp.playful-patterns.com/public_html/vlok.wasm", "vlok.wasm")