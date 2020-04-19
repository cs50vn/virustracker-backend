echo "Upload an app image"
echo "======================="

export REGISTRY_URL=$1
export REGISTRY_USER=$2
export REGISTRY_PASS=$3
export REGISTRY_APPNAME=$4
export REGISTRY_TAGNAME=$5

#docker login $REGISTRY_URL -u $REGISTRY_USER -p $REGISTRY_PASS
docker push ${REGISTRY_URL}/${REGISTRY_APPNAME}:${REGISTRY_TAGNAME}