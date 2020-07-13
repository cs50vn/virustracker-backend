echo "Deploy to dev server"
echo "======================="

#Deploy to dev enviroment

export REGISTRY_URL=$1
export REGISTRY_USER=$2
export REGISTRY_PASS=$3
export IMAGE_NAME=$4
export IMAGE_TAG=$5
export DEPLOY_SERVER_URL=$6
export DEPLOY_SERVER_USER=$7
export DEPLOY_SERVER_KEY="$8"

##Init process 
echo "${DEPLOY_SERVER_KEY}" > key.pem
ls -l -a
chmod 600 key.pem

##Set up 3 container

export APP_CMD="docker login $REGISTRY_URL -u $REGISTRY_USER -p $REGISTRY_PASS;
        docker stop virustracker-backend-${IMAGE_TAG};
        docker pull ${REGISTRY_URL}/${IMAGE_NAME}:${IMAGE_TAG}; 
        docker run -d --restart always --name virustracker-backend-${IMAGE_TAG} --rm --network host -d ${REGISTRY_URL}/${IMAGE_NAME}:${IMAGE_TAG};
        docker images;
        docker ps -a" 
ssh -i key.pem -o StrictHostKeyChecking=no $DEPLOY_SERVER_USER@$DEPLOY_SERVER_URL $APP_CMD
#