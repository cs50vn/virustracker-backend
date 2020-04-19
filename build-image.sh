echo "Build an app image"
echo "======================="

export APP_PORT=$1
export REGISTRY_URL=$2
export REGISTRY_APPNAME=$3
export REGISTRY_TAGNAME=$4
export BUILD_TYPE=release

python3 scripts/build-image.py $PWD linux $BUILD_TYPE $APP_PORT $REGISTRY_URL $REGISTRY_APPNAME $REGISTRY_TAGNAME
