@echo off

set BUILD_TYPE=all

python scripts/make-test-data.py %CD% windows %BUILD_TYPE%
