@echo off
set BUILD_DIR=bin
set TOOLS_DIR=tools

if "%1"=="" goto help
if "%1"=="all" goto all
if "%1"=="clean" goto clean
if "%1"=="lint" goto lint
if "%1"=="build" goto build
if "%1"=="test" goto test
if "%1"=="run" goto run
if "%1"=="help" goto help
echo ERROR: unknown command
goto help

:all
call :clean
call :lint
call :test
call :build
call :run
goto :eof

:build
if not exist %BUILD_DIR% mkdir %BUILD_DIR%
set CGO_ENABLED=0
go build -ldflags="-s -w" -o %BUILD_DIR%\server.exe .\cmd\server
goto :eof

:clean
if exist %BUILD_DIR% rd /s /q %BUILD_DIR%
if exist %TOOLS_DIR% rd /s /q %TOOLS_DIR%
go mod tidy
goto :eof

:run
call :build
%BUILD_DIR%\server.exe
goto :eof

:lint
if not exist %TOOLS_DIR% mkdir %TOOLS_DIR%
if not exist %TOOLS_DIR%\golangci-lint\golangci-lint.exe (
  curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b %TOOLS_DIR%\golangci-lint latest
)
%TOOLS_DIR%\golangci-lint\golangci-lint.exe run .\...
goto :eof

:test
go test -race -cover -coverprofile=coverage.txt -covermode=atomic .\...
goto :eof

:help
echo Available commands:
echo   all       Run all tests, then build and run
echo   build     Build the binary
echo   clean     Clean up build artifacts
echo   lint      Run linters
echo   run       Run the binary
echo   test      Run tests
goto :eof