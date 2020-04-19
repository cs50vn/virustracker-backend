echo "Prepare artifacts"
echo "======================="

export BUILD_TYPE=release

python3 scripts/prepare-artifacts.py $PWD linux $BUILD_TYPE


