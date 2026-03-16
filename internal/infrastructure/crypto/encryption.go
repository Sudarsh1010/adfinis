package crypto

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/nacl/secretbox"

	"github.com/uptrace/bun"
)

// Service implements encryption using nacl/secretbox
type Service struct {
	db *bun.DB
}

// NewService creates a new encryption service
func NewService(db *bun.DB) *Service {
	return &Service{db: db}
}

// Encrypt encrypts plaintext using nacl/secretbox
// Returns base64-encoded ciphertext with nonce prepended
func (s *Service) Encrypt(plaintext string) (string, error) {
	// Get or create the encryption key
	key, err := getOrCreateKey(s.db)
	if err != nil {
		return "", fmt.Errorf("failed to get encryption key: %w", err)
	}

	// Generate a new nonce for each encryption
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encode plaintext to bytes
	plaintextBytes := []byte(plaintext)

	// Encrypt using secretbox (returns a slice)
	ciphertext := secretbox.Seal(nil, plaintextBytes, &nonce, &key)

	// Prepend nonce to ciphertext so it can be decrypted later
	nonceAndCiphertext := append(nonce[:], ciphertext...)

	// Encode to base64
	return base64.StdEncoding.EncodeToString(nonceAndCiphertext), nil
}

// Decrypt decrypts base64-encoded ciphertext
func (s *Service) Decrypt(ciphertext string) (string, error) {
	// Get or create the encryption key
	key, err := getOrCreateKey(s.db)
	if err != nil {
		return "", fmt.Errorf("failed to get encryption key: %w", err)
	}

	// Decode base64 ciphertext
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	if len(ciphertextBytes) < 24+secretbox.Overhead {
		return "", errors.New("ciphertext too short")
	}

	// Extract nonce and actual ciphertext
	var nonce [24]byte
	copy(nonce[:], ciphertextBytes[:24])
	ciphertextData := ciphertextBytes[24:]

	// Decrypt using secretbox
	plaintext, ok := secretbox.Open(nil, ciphertextData, &nonce, &key)
	if !ok {
		return "", errors.New("decryption failed: invalid key or corrupted data")
	}

	return string(plaintext), nil
}

// encryptionKey represents a row in the encryption_keys table
type encryptionKey struct {
	ID      string `bun:"id,pk"`
	KeyData string `bun:"key_data"`
}

// getOrCreateKey retrieves an existing key from the database or creates a new one
func getOrCreateKey(db *bun.DB) ([32]byte, error) {
	// First try to get existing key
	var keyModel encryptionKey
	err := db.NewSelect().
		Model(&keyModel).
		Where("id = ?", "default").
		Scan(context.Background())
	if err == nil {
		// Key exists, decode and return
		keyBytes, err := base64.StdEncoding.DecodeString(keyModel.KeyData)
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to decode key: %w", err)
		}

		if len(keyBytes) != 32 {
			return [32]byte{}, errors.New("invalid key length")
		}

		var key [32]byte
		copy(key[:], keyBytes)
		return key, nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return [32]byte{}, fmt.Errorf("failed to query key: %w", err)
	}

	// Key doesn't exist, create a new one
	return generateNewKey(db)
}

// generateNewKey generates a new 32-byte encryption key and stores it in the database
func generateNewKey(db *bun.DB) ([32]byte, error) {
	// Generate a new random key
	var key [32]byte
	if _, err := rand.Read(key[:]); err != nil {
		return [32]byte{}, fmt.Errorf("failed to generate key: %w", err)
	}

	// Encode key to base64
	keyData := base64.StdEncoding.EncodeToString(key[:])

	// Store key in database
	keyModel := &encryptionKey{
		ID:      "default",
		KeyData: keyData,
	}
	_, err := db.NewInsert().Model(keyModel).Exec(context.Background())
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to store key: %w", err)
	}

	return key, nil
}
