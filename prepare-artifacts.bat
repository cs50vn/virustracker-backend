echo off

set BUILD_TYPE=release

python scripts/prepare-artifacts.py %CD% windows %BUILD_TYPE%

