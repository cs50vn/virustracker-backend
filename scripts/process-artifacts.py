import sys, os, shutil, subprocess, time, config

def processArtifacts():
    print("")

def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mPROCESS ARTIFACTS\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    config.buildProjectPath(argv[0], argv[1], argv[2])

    processArtifacts()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))

if __name__ == '__main__':
    main(sys.argv[1:])
