package utils

import (
	"cs50vn/virustracker/apprepository"
	"fmt"
	"sort"
)

func LoadFakeData() {
	LoadFakeVersions()
	LoadFakeStories()
	LoadFakeTopCollections()
	LoadFakeTopGenres()
	LoadFakeSearch()
	LoadFakeCDN()
}

func LoadFakeVersions() {
	//Load app versions
	apprepository.ListOfVersions[1]= MakeVersion(1, "force_update", "b2.7perldata.tech/app-v1.apk")
	apprepository.ListOfVersions[2]= MakeVersion(2, "recommend", "b2.7perldata.tech/app-v2.apk")
	apprepository.ListOfVersions[3]= MakeVersion(3, "recommend", "b2.7perldata.tech/app-v3.apk")
	apprepository.ListOfVersions[4]= MakeVersion(4, "force_update", "b2.7perldata.tech/app-v4.apk")
	apprepository.ListOfVersions[5]= MakeVersion(5, "force_update", "b2.7perldata.tech/app-v5.apk")

	apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, 1)
	apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, 2)
	apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, 3)
	apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, 4)
	apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, 5)

	sort.Ints(apprepository.ListOfVersionsKeys)

	for _ , value := range apprepository.ListOfVersionsKeys {
		fmt.Println(value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(apprepository.ListOfVersionsKeys)))

	fmt.Println(apprepository.ListOfVersionsKeys)
}

func LoadFakeTopCollections() {

}

func LoadFakeTopGenres() {

}

func LoadFakeStories() {

}

func LoadFakeSearch() {

}

func LoadFakeCDN() {

}