@echo off

set BUILD_TYPE=all

python scripts/generate-input-sheet.py %CD% windows %BUILD_TYPE%
