@echo off
go build -ldflags "-H=windowsgui" -o KAS_PDC.exe main/main.go

if errorlevel 1 (
    echo Build failed.
) else (
    echo Built Successfully: KAS_PDC.exe
    start .\KAS_PDC.exe
)
