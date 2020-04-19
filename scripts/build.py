import sys, os, shutil, subprocess, time, config


def buildPackage():
    print("===========================================================")
    print("                      \033[1;32;40mBUILD PACKAGE\033[0;37;40m")
    print("===========================================================")

    src = config.srcDir
    des = config.genAppDir

    print("\033[1;34;40mFrom:\n\033[0;37;40m" + src)
    print("\033[1;34;40mTo\n\033[0;37;40m" + des)

    if os.path.exists(des):
        shutil.rmtree(des)
    shutil.copytree(src, des)

    src = config.dataDir + os.sep + "app_template.db"
    shutil.copyfile(src, des + os.sep + "app.db")

    cmd = "7z a %s.zip %s" % (des, des)
    subprocess.call(cmd, shell=True)
    
    print("\n")


def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mBUILD APPLICATION\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    config.buildProjectPath(argv[0], argv[1], argv[2])

    buildPackage()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))


if __name__ == '__main__':
    main(sys.argv[1:])
