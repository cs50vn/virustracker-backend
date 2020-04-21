package utils

import "cs50vn/virustracker/apprepository/model"

///////////////////////////////////////////////////////////////////////////////
//Helper

func MakeVersion(versionCode int, status string, downloadLink string) *model.Version {
	return &model.Version{versionCode, status, downloadLink}
}