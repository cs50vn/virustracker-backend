import sys, os, shutil, subprocess, time, config, sqlite3
from model import repo

def generateTestDB():        
    url = "https://f000.backblazeb2.com/file/projecta-build/virustracker-backend/v1/dev/app.db"
    cmd = "curl %s -o %sapp.db" % ( url, (config.genDataDir + os.sep))
    print(cmd)
    subprocess.call(cmd, shell=True)

    cmd = "ls -l %s" % config.genDataDir
    print(cmd)
    subprocess.call(cmd, shell=True)

def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mGENERATE TEST DB\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    config.buildProjectPath(argv[0], argv[1], argv[2])

    generateTestDB()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))

if __name__ == '__main__':
    main(sys.argv[1:])


