
echo "Build a dev environment"
echo "======================="

export ROOT_DIR=$1
echo "RootDir ${ROOT_DIR}"

#Install prerequisites
echo "Install prerequisites"
echo "======================="
echo -e "\n*****  0  *****"
pip3 install xlrd
pip3 install pandas
#$ROOT_DIR/install-prerequisites.sh

#Make test data
echo -e "\n*****  1  *****"
$ROOT_DIR/make-test-data.sh

#Compile app
echo -e "\n*****  2  *****"
$ROOT_DIR/make.sh debug

#Build app image
echo -e "\n*****  4  *****"
$ROOT_DIR/build-image.sh $APP_PORT_DEV $REGISTRY_URL $REGISTRY_APPNAME $REGISTRY_TAGNAME

#Upload app image
echo -e "\n*****  5  *****"
$ROOT_DIR/upload-image.sh $REGISTRY_URL $REGISTRY_USER $REGISTRY_PASS $REGISTRY_APPNAME $REGISTRY_TAGNAME

#Deploy all to dev server
echo -e "\n*****  6  *****"
$ROOT_DIR/deploy-dev.sh $REGISTRY_URL $REGISTRY_USER $REGISTRY_PASS $REGISTRY_APPNAME $REGISTRY_TAGNAME $DEV_SERVER_URL $DEV_SERVER_USER "${DEV_SERVER_KEY}"

#
