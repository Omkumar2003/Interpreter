@echo off
setlocal enabledelayedexpansion
set "file_extension=%~x1"

if /i "%file_extension%"==".om" (
    go run main.go %1
) else (
    echo Error: Unsupported file type. Only .om files are supported.
)
