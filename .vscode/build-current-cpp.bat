@echo off
setlocal

set "SOURCE=%~1"
if "%SOURCE%"=="" (
    echo Missing source file.
    exit /b 1
)

set "ROOT=%~dp0"
set "OUT_DIR=%ROOT%bin"
if not exist "%OUT_DIR%" mkdir "%OUT_DIR%"

call "C:\Program Files\Microsoft Visual Studio\2022\Community\VC\Auxiliary\Build\vcvars64.bat" >nul
if errorlevel 1 exit /b 1

cl /nologo /EHsc /std:c++17 /Fe:"%OUT_DIR%\%~n1.exe" "%SOURCE%"

endlocal