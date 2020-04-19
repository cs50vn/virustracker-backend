@echo off

echo "Process artifacts"
echo "======================="


set BUILD_TYPE=all

python scripts/process-artifacts.py %CD% windows %BUILD_TYPE%
