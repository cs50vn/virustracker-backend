import sys, os, shutil, subprocess, time, config

# App config
appPort = ""
registryUrl = ""
registryAppName = ""
registryTagName = ""

def buildPath(rootPath, host, build, port, url, app, tag):
    config.buildProjectPath(rootPath, host, build)
    global appPort
    appPort = port
    global registryUrl
    registryUrl = url
    global registryAppName
    registryAppName = app
    global registryTagName
    registryTagName = tag

def buildImage():
    cmd = ''' docker build -f .ci/docker/Dockerfile-dev --tag "%s/%s:%s" $PWD --build-arg SRC_DIR=%s --build-arg APP_PORT=%s;
            docker images
        ''' % (registryUrl, registryAppName, registryTagName, "_generated" + os.sep + config.genAppDirPath , appPort)
    print(cmd)
    subprocess.call(cmd, shell=True)

def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mBUILD IMAGE\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    buildPath(argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6])

    buildImage()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))

if __name__ == '__main__':
    main(sys.argv[1:])
