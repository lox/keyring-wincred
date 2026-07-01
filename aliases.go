package wincred

import "github.com/lox/keyring/v2"

type Item = keyring.Item
type Metadata = keyring.Metadata

const WinCredBackend = keyring.WinCredBackend

var (
	ErrCredentialTooLarge   = keyring.ErrCredentialTooLarge
	ErrKeyNotFound          = keyring.ErrKeyNotFound
	ErrMetadataNotSupported = keyring.ErrMetadataNotSupported
)
