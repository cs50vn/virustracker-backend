import sys, os, shutil, subprocess, time, config


def buildGoProgram():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD GO\033[0;37;40m")
    print("===========================================================")
    
    os.chdir(config.srcDir)

    if config.buildType == "release":
        if config.hostType == "windows":
            os.environ["GOOS"] = "windows"
            cmd = "go build -ldflags \"-s -w\" -o %s.exe main.go" % (config.genRootDir + os.sep + config.outputFile)
            print(cmd)
            subprocess.call(cmd, shell=True)
            cmd = "upx %s.exe" % (config.genRootDir + os.sep + config.outputFile)
            subprocess.call(cmd, shell=True)

        else:
            os.environ["GOOS"] = "linux"
            cmd = "go build -ldflags \"-s -w\" -o %s main.go" % (config.genRootDir + os.sep + config.outputFile)
            print(cmd)
            subprocess.call(cmd, shell=True)
            cmd = "upx %s" % (config.genRootDir + os.sep + config.outputFile)
            subprocess.call(cmd, shell=True)
            cmd = "chmod 740 %s" % (config.genRootDir + os.sep + config.outputFile)
            subprocess.call(cmd, shell=True)
    else:
        #Debug build
        if config.hostType == "windows":
            os.environ["GOOS"] = "windows"
            cmd = "go build -ldflags \"-s -w\" -o %s.exe main.go" % (config.genRootDir + os.sep + config.outputFile)
            print(cmd)
            subprocess.call(cmd, shell=True)
        else:
            os.environ["GOOS"] = "linux"
            cmd = "go build -ldflags \"-s -w\" -o %s main.go" % (config.genRootDir + os.sep + config.outputFile)
            print(cmd)
            subprocess.call(cmd, shell=True)
            cmd = "chmod 740 %s" % (config.genRootDir + os.sep + config.outputFile)
            subprocess.call(cmd, shell=True)

    os.chdir(config.rootDir)


def buildPackage():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD PACKAGE\033[0;37;40m")
    print("===========================================================")

    if config.buildType == "release":
        print("")
    else:
        if config.hostType == "windows":
            src = config.genRootDir + os.sep + config.outputFile + ".exe"
            des = config.genAppDir

            print("\033[1;34;40mFrom:\n\033[0;37;40m" + src)
            print("\033[1;34;40mTo\n\033[0;37;40m" + des)

            if not os.path.exists(des):
                os.makedirs(des, exist_ok=True)
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "config.json"
            shutil.copy(src, des)

            src = config.genDataDir + os.sep + "%s.db" % config.appName
            shutil.copyfile(src, config.genAppDir + os.sep + "%s.db" % config.appName)

            cmd = "7z a %s.zip %s" % (des, des)
            subprocess.call(cmd, shell=True)

        else:
            src = config.genRootDir + os.sep + config.outputFile
            des = config.genAppDir

            print("\033[1;34;40mFrom:\n\033[0;37;40m" + src)
            print("\033[1;34;40mTo\n\033[0;37;40m" + des)

            if not os.path.exists(des):
                os.makedirs(des, exist_ok=True)
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "install.sh"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "uninstall.sh"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "config.json"
            shutil.copy(src, des)

            src = config.templateDir + os.sep + "%s.service" % config.appName
            shutil.copy(src, des)

            src = config.genDataDir + os.sep + "%s.db" % config.appName
            shutil.copyfile(src, config.genAppDir + os.sep + "%s.db" % config.appName)

            cmd = "7z a %s.zip %s" % (des, des)
            subprocess.call(cmd, shell=True)
        
            print("")

    print("\n")


def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mBUILD APPLICATION\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    config.buildProjectPath(argv[0], argv[1], argv[2])

    buildGoProgram()

    buildPackage()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))


if __name__ == '__main__':
    main(sys.argv[1:])
