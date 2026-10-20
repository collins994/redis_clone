@echo off

if "%1" == "build" (
	pushd code
		go build -o ../build/redis.exe
	popd
	exit /b
)

if "%1" == "run" (
	pushd build
		redis.exe
	popd
	exit /b
)

if "%1" == "test" (
	pushd code
		go test .
	popd
	exit /b
)

if "%1" == "tags" (
	echo generating tags...
	ctags -R  --exclude=build
	EXIT /B
)
