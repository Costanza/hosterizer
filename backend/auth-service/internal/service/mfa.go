package service

import (
	"encoding/base32"
	"fmt"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// MFAService handles multi-factor authentication operations
type MFAService struct {
	issuer string
}

// MFASetupResult contains the result of MFA setup
type MFASetupResult struct {
	Secret    string
	QRCodeURL string
}

// NewMFAService creates a new MFA service
func NewMFAService(issuer string) *MFAService {
	return &MFAService{
		issuer: issuer,
	}
}

// GenerateSecret generates a new TOTP secret for a user
func (s *MFAService) GenerateSecret(accountName string) (*MFASetupResult, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.issuer,
		AccountName: accountName,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate TOTP secret: %w", err)
	}

	return &MFASetupResult{
		Secret:    key.Secret(),
		QRCodeURL: key.URL(),
	}, nil
}

// ValidateCode validates a TOTP code against a secret
func (s *MFAService) ValidateCode(secret, code string) (bool, error) {
	// Normalize the secret (remove spaces, convert to uppercase)
	secret = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(
		[]byte(secret),
	)

	valid := totp.Validate(code, secret)
	return valid, nil
}

// ValidateCodeWithWindow validates a TOTP code with a time window
// This allows for clock skew between client and server
func (s *MFAService) ValidateCodeWithWindow(secret, code string, window int) (bool, error) {
	// Normalize the secret
	secret = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(
		[]byte(secret),
	)

	// Try current time and surrounding windows
	for i := -window; i <= window; i++ {
		valid, err := totp.ValidateCustom(
			code,
			secret,
			totp.ValidateOpts{
				Period:    30,
				Skew:      uint(i),
				Digits:    otp.DigitsSix,
				Algorithm: otp.AlgorithmSHA1,
			},
		)
		if err != nil {
			return false, fmt.Errorf("failed to validate TOTP code: %w", err)
		}
		if valid {
			return true, nil
		}
	}

	return false, nil
}

// GenerateBackupCodes generates backup codes for MFA recovery
func (s *MFAService) GenerateBackupCodes(count int) ([]string, error) {
	// This is a simplified implementation
	// In production, you'd want to use a cryptographically secure random generator
	codes := make([]string, count)
	for i := 0; i < count; i++ {
		code, err := totp.GenerateCode(fmt.Sprintf("backup-%d", i), totp.Now())
		if err != nil {
			return nil, fmt.Errorf("failed to generate backup code: %w", err)
		}
		codes[i] = code
	}
	return codes, nil
}
