# PolyRhythm Agent
![](https://travis-ci.org/FourSee/ShellGame.svg?branch=master)

To cross-platform compile:

```bash
 $ PLATFORMS=windows/amd64,darwin/amd64,linux/amd64 ./build.sh
```

The build platform targets are taken from the `PLATFORMS` env var. It is expected to be a comma-separated list


| OS |  CPU arch |
|----|---|
| android | 	arm |
| darwin | 	386 |
| darwin | 	amd64 |
| darwin | 	arm |
| darwin | 	arm64 |
| dragonfly | 	amd64 |
| freebsd | 	386 |
| freebsd | 	amd64 |
| freebsd | 	arm |
| linux | 	386 |
| linux | 	amd64 |
| linux | 	arm |
| linux | 	arm64 |
| linux | 	ppc64 |
| linux | 	ppc64le |
| linux | 	mips |
| linux | 	mipsle |
| linux | 	mips64 |
| linux | 	mips64le |
| netbsd | 	386 |
| netbsd | 	amd64 |
| netbsd | 	arm |
| openbsd | 	386 |
| openbsd | 	amd64 |
| openbsd | 	arm |
| plan9 | 	386 |
| plan9 | 	amd64 |
| solaris | 	amd64 |
| windows | 	386 |
| windows | 	amd64 |