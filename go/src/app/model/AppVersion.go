package model

type AppVersion struct {
	Major      int
	Minor      int
	Patch      int
	CodeNumber int

	UpdateStatus string //force_update, recommend, none
}
