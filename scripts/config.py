import sys, os, shutil, subprocess, time


# Build tool
hostType = ""
buildType = ""


# Project config
rootDir     = ""
templateDir = ""
dataDir     = ""
docsDir     = ""
scriptDir   = ""
srcDir      = ""
genRootDir  = ""
genAppDir       = ""
genAppDirPath   = ""
genDataDir  = ""
goPathDir   = ""

# App config
appName = "virustracker"
outputFile = appName
versionCode = "1"
internalVersionCode = "1.0.0"
versionName = "v" + versionCode
prefixNameLinux = "linux_amd64"
prefixNameWin = "win_x64"

def buildProjectPath(rootPath, host, build):
    global rootDir
    rootDir = rootPath
    global hostType
    hostType = host
    global buildType
    buildType = build

    global templateDir
    templateDir = rootDir + os.sep + "templates"
    global dataDir
    dataDir = rootDir + os.sep + "data"
    global docsDir
    docsDir = rootDir + os.sep + "docs"
    global scriptDir
    scriptDir = rootDir + os.sep + "scripts"
    global srcDir
    srcDir = rootDir + os.sep + "go" + os.sep + "src"
    global toolsDir
    toolsDir = rootDir + os.sep + "tools"
    global genRootDir
    genRootDir = rootDir + os.sep + "_generated"

    if hostType == "linux":
        desGenName = "%s-%s_%s-%s" % (appName, versionName, internalVersionCode, prefixNameLinux)
    else:
        desGenName = "%s-%s_%s-%s" % (appName, versionName, internalVersionCode, prefixNameWin)
    global genAppDirPath
    genAppDirPath = versionName + os.sep + desGenName
    global genAppDir
    genAppDir = genRootDir + os.sep + genAppDirPath
    global genDataDir
    genDataDir = genRootDir + os.sep + versionName + os.sep + "data"

    global goPathDir
    goPathDir = rootDir + os.sep + "go"

    os.environ["GOPATH"] = goPathDir
    print("\033[1;34;40mLoad build config\033[0;37;40m")
    print("Host: \033[1;32;40m%s\033[0;37;40m" % hostType)
    print("Build Type: \033[1;32;40m%s\033[0;37;40m" % buildType)
    print("Version code: \033[1;32;40m%s\033[0;37;40m" % versionCode)
    print("Version name: \033[1;32;40m%s\033[0;37;40m\n" % versionName)

    print("\033[1;34;40mLoad project config\033[0;37;40m")
    print("Root dir: 	\033[1;34;40m%s\033[0;37;40m" % rootDir)
    print("Templates dir: 	\033[1;34;40m%s\033[0;37;40m" % templateDir)
    print("Data dir:	\033[1;34;40m%s\033[0;37;40m" % dataDir)
    print("Docs dir: 	\033[1;34;40m%s\033[0;37;40m" % docsDir)
    print("Script dir:	\033[1;34;40m%s\033[0;37;40m" % scriptDir)
    print("Source dir: 	\033[1;34;40m%s\033[0;37;40m" % srcDir)
    print("Gen root dir: 	\033[1;34;40m%s\033[0;37;40m" % genRootDir)
    print("Gen app dir: 	\033[1;34;40m%s\033[0;37;40m" % genAppDir)
    print("Gen data dir: 	\033[1;34;40m%s\033[0;37;40m" % genDataDir)
    print("GOPATH dir:      \033[1;34;40m%s\033[0;37;40m" % goPathDir)

    print("\n")
