package crypto

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"database/sql"
)

func TestEncryptionRoundtrip(t *testing.T) {
	db := setupTestDB(t)

	testCases := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"short string", "hello"},
		{"long string", "This is a very long plaintext string for testing encryption functionality and ensuring data security integrity"},
		{"unicode characters", "Hello 世界 🌍"},
		{"numbers and symbols", "12345!@#$%"},
		{"json-like content", `{"name":"test","value":123}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(db)
			ciphertext, err := service.Encrypt(tc.input)
			require.NoError(t, err)
			require.NotEmpty(t, ciphertext)

			plaintext, err := service.Decrypt(ciphertext)
			require.NoError(t, err)
			require.Equal(t, tc.input, plaintext)
		})
	}
}

func TestKeyAutoGeneration(t *testing.T) {
	db := setupTestDB(t)
	input := "test data"

	service := NewService(db)

	ciphertext1, err1 := service.Encrypt(input)
	require.NoError(t, err1)
	ciphertext2, err2 := service.Encrypt(input)
	require.NoError(t, err2)

	require.NotEmpty(t, ciphertext1)
	require.NotEmpty(t, ciphertext2)
	require.NotEqual(t, ciphertext1, ciphertext2)

	plaintext1, err3 := service.Decrypt(ciphertext1)
	require.NoError(t, err3)
	require.Equal(t, input, plaintext1)

	plaintext2, err4 := service.Decrypt(ciphertext2)
	require.NoError(t, err4)
	require.Equal(t, input, plaintext2)
}

func TestMultipleEncryptDecryptCycles(t *testing.T) {
	db := setupTestDB(t)
	input := "cycling test"

	for i := 0; i < 10; i++ {
		service := NewService(db)
		ciphertext, err := service.Encrypt(input)
		require.NoError(t, err)

		plaintext, err := service.Decrypt(ciphertext)
		require.NoError(t, err)
		require.Equal(t, input, plaintext)
	}
}

func TestDecryptWithoutInitialization(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	_, err := service.Decrypt("some ciphertext")
	require.Error(t, err)
}

func TestDecryptWithInvalidBase64(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	_, err := service.Decrypt("not valid base64!")
	require.Error(t, err)
}

func TestDecryptWithInvalidNonce(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	ciphertext, err := service.Encrypt("test")
	require.NoError(t, err)

	ciphertextBytes := []byte(ciphertext)
	ciphertextBytes[5] = 0xFF
	corrupted := string(ciphertextBytes)

	plaintext, err := service.Decrypt(corrupted)
	require.Error(t, err)
	require.NotEqual(t, "test", plaintext)
}

func setupTestDB(t *testing.T) *bun.DB {
	t.Helper()

	sqldb, err := sql.Open(sqliteshim.ShimName, ":memory:")
	require.NoError(t, err)

	_, err = sqldb.ExecContext(context.Background(), "PRAGMA foreign_keys = ON")
	require.NoError(t, err)

	_, err = sqldb.ExecContext(context.Background(), `
		CREATE TABLE IF NOT EXISTS encryption_keys (
			id TEXT PRIMARY KEY,
			key_data TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	require.NoError(t, err)

	return bun.NewDB(sqldb, sqlitedialect.New())
}
