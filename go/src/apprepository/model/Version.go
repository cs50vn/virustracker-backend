package model

type Version struct {
	VersionCode int
	Status string
	DownloadLink string
}

func MakeVersion(versionCode int, status string, downloadLink string) *Version {
	return &Version{VersionCode: versionCode, Status: status, DownloadLink: downloadLink}
}

