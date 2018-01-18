
set source=%~dp0
set dst=%~dp0\..\bin\snaky-bin

set dirs=html\addon

set files=html\css\global.css html\css\jmdkina.css 
set files=%files% html\css\resume.css html\css\shici.css

set files=%files% html\jmdkina\add.html html\jmdkina\jmdkina.html
set files=%files% html\shici\add.html html\shici\shici.html
set files=%files% html\resume\resume.html
set files=%files% html\404.html

set files=%files% html\js\jmdkina\add.js html\js\jmdkina\jmdkina.js
set files=%files% html\js\shici\add.js html\js\shici\shici.js
set files=%files% html\js\global.js

for %%i in (%dirs% ) do xcopy %source%\%%i %dst%\%%i /S /E /Y

for %%i in (%files%) do xcopy %source%\%%i %dst%\%%i /Y

xcopy %source%\bin\simpleserver.exe %dst%\bin\snaky.exe /Y

pause