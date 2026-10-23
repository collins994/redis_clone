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
		go test ./store
	popd
	exit /b
)

if "%1" == "push" (
	git push https://github.com/collins994/redis_clone main
)

if "%1" == "tags" (
	ctags -R  --exclude=build
	EXIT /B
)
