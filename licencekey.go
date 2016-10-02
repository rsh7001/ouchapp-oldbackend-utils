package ouchapp

type LicenceKey struct {
	ID               string `json:"id"`
	LicenceKeyString string `json:"licensekeystring"`
	UserID           string `json:"userid"`
	Active           bool   `json:"active"`
}
