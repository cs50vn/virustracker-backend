import sys, os, time, config, pandas, sqlite3


def buildTestData():
    sourceFile = config.dataDir + os.sep + config.schemaOutputFile
    desFile = config.genDataDir + os.sep + config.appName + ".db"
    db = sqlite3.connect(desFile)
    dfs = pandas.read_excel(sourceFile, sheet_name=None)
    for table, df in dfs.items():
        df.to_sql(table, db, index=False, if_exists="replace")


def main(argv):
    start = time.time()
    print("===========================================================")
    print("                      \033[1;32;40mBUILD TEST DATA\033[0;37;40m")
    print("===========================================================")

    print(str(argv))
    config.buildProjectPath(argv[0], argv[1], argv[2])

    buildTestData()

    elapsedTime = time.time() - start
    print("Running time: %s s" % str(elapsedTime))


if __name__ == '__main__':
    main(sys.argv[1:])
