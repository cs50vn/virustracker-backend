echo "Process artifacts"
echo "======================="

export BUILD_TYPE=all

python3 scripts/process-artifacts.py $PWD linux $BUILD_TYPE


