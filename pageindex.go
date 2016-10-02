package ouchapp

type PageIndex struct {
	ID                string `json:"id"`
	UserID            string `json:"userid"`
	InitialPage       string `json:"initialpage"`
	LastServerUpdated int64  `json:"lastserverupdated"`
	LastRemoteUpdated int64  `json""lastremoteupdated"`
}
