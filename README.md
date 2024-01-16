# Url-Shortener
Simple url shortener built with Gin framework.

To use this shortener you can clone the repo and then execute the following commands:<br>
```powershell
Set-Location .\Url-Shortener
go run .
```
Also you need to add the line below to your hosts file.<br>
Run the Powershell as Administrator first then execute:<br>
```powershell
Add-Content -Uri "C:\Windows\System32\drivers\etc\hosts" -Value "127.0.0.1 exmpl.cm"
```
API is simple and straightforward. It's have only two endpoints. Example:<br>
```powershell
Invoke-RestMethod -Method Get -Uri "http://localhost/api/getUrls"
Invoke-RestMethod -Method Get -Url "http://localhost/api/getShortUrl/<your-url>"
```
After executing the commands above you get responce as json if you using a browser. If you are using powershell so you will see the following:<br>
```powershell
originalUrl    shortenUrl
-----------    ----------
www.google.com exmpl.cm/nIxffk

originalUrl                                 shortenUrl
-----------                                 ----------
https://pkg.go.dev/std                      exmpl.cm/bNnNki
ya.ru                                       exmpl.cm/MuCTXt
https://www.drive2.ru/l/465151143983449251/ exmpl.cm/hNlYam
https://ottplayer.tv/support                exmpl.cm/vjEbFH
ya.ru                                       exmpl.cm/vMRZbR
www.example.com                             exmpl.cm/mKddNr
```
