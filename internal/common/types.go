package common

import (
	"encoding/json"
	"time"
)

type KeyPair struct {
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
	PublicKey           string `json:"publicKey"`
}

type Account struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Email              string  `json:"email"`
	MasterPasswordHash string  `json:"masterPasswordHash"`
	MasterPasswordHint string  `json:"masterPasswordHint"`
	Key                string  `json:"key"`
	KeyPair            KeyPair `json:"keys"`
	RefreshToken       string  `json:"-"`
}

func (acc Account) getProfile() Profile {
	return Profile{
		Id:                 acc.Id,
		Name:               nil,
		Email:              acc.Email,
		EmailVerified:      false,
		Premium:            false,
		Culture:            "en-US",
		TwoFactorEnabled:   false,
		Key:                acc.Key,
		SecurityStamp:      nil,
		Organizations:      make([]string, 0),
		MasterPasswordHint: nil,
		PrivateKey:         acc.KeyPair.EncryptedPrivateKey,
		Object:             "profile",
	}
}

// The data we store and send to the client
type Cipher struct {
	Type                int
	FolderId            *string // Must be pointer to output null in json. Android app will crash if not null
	OrganizationId      *string
	Favorite            bool
	Edit                bool
	Id                  string
	Data                CipherData
	Attachments         []string
	OrganizationUseTotp bool
	RevisionDate        time.Time
	Object              string
	CollectionIds       []string
}

type CipherData struct {
	Uri      string
	Username string
	Password string
	Totp     *string // Must be pointer to output null in json. Android app will crash if not null
	Name     string
	Notes    *string // Must be pointer to output null in json. Android app will crash if not null
	Fields   []string
}

func (data *CipherData) bytes() ([]byte, error) {
	b, err := json.Marshal(&data)
	return b, err
}

type Profile struct {
	Id                 string
	Name               *string
	Email              string
	EmailVerified      bool
	Premium            bool
	MasterPasswordHint *string
	Culture            string
	TwoFactorEnabled   bool
	Key                string
	PrivateKey         string
	SecurityStamp      *string
	Organizations      []string
	Object             string
}

type SyncData struct {
	Profile Profile
	Folders []Folder
	Ciphers []Cipher
	Domains Domains
	Object  string
}

type Domains struct {
	EquivalentDomains       []string
	GlobalEquivalentDomains []GlobalEquivalentDomains
	Object                  string
}

type GlobalEquivalentDomains struct {
	Type     int
	Domains  []string
	Excluded bool
}

type Folder struct {
	Id           string
	Name         string
	Object       string
	RevisionDate time.Time
}

type Data struct {
	Object string
	Data   interface{}
}

// Interface to make testing easier
type Database interface {
	Init() error
	addAccount(acc Account) error
	getAccount(username string, refreshtoken string) (Account, error)
	updateAccountInfo(acc Account) error
	getCipher(owner string, ciphID string) (Cipher, error)
	getCiphers(owner string) ([]Cipher, error)
	newCipher(ciph Cipher, owner string) (Cipher, error)
	updateCipher(newData Cipher, owner string, ciphID string) error
	deleteCipher(owner string, ciphID string) error
	Open() error
	Close()
	addFolder(name string, owner string) (Folder, error)
	getFolders(owner string) ([]Folder, error)
}