package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

var (
	// timeNowWrapper returns time now, that can be changed for unit test
	timeNowWrapper = func() time.Time {
		return time.Now().UTC()
	}

	defaultHeader = Header{
		Alg: "HS256",
		Typ: "JWT",
	}
)

// GenerateToken creates new HS256 JWT token by given payload
func GenerateToken(secret []byte, claims Claims) (string, error) {
	header, err := json.Marshal(defaultHeader)
	if err != nil {
		return "", err
	}

	payload, err := marshalClaims(claims)
	if err != nil {
		return "", err
	}

	// Base64 URL encode the header and payload
	encodedHeader := base64UrlEncode(header)
	encodedPayload := base64UrlEncode(payload)

	message := append(encodedHeader, []byte(".")...)
	message = append(message, encodedPayload...)

	// Computes signature with HMAC SHA-256
	signature := hmacSHA256Encode(message, secret)

	encodedSignature := base64UrlEncode(signature)

	return fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, encodedSignature), nil
}

// VerifyToken verify token an return claims
func VerifyToken(secret []byte, token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Claims{}, ErrInvalidToken
	}

	// Verify header
	decodedHeader, err := base64UrlDecode([]byte(parts[0]))
	if err != nil {
		return Claims{}, err
	}

	var header Header
	if err := json.Unmarshal(decodedHeader, &header); err != nil {
		return Claims{}, err
	}

	// !Only support JWT HS256 on this application
	if header.Alg != "HS256" || header.Typ != "JWT" {
		return Claims{}, ErrInvalidHeader
	}

	// Verify signature
	message := parts[0] + "." + parts[1]
	signature := hmacSHA256Encode([]byte(message), secret)

	if parts[2] != string(base64UrlEncode(signature)) {
		return Claims{}, ErrInvalidSignature
	}

	// Verify payload
	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Claims{}, err
	}

	claims, err := unmarshalClaims(decodedPayload)
	if err != nil {
		return Claims{}, err
	}

	// Return error if current time after expired time
	if timeNowWrapper().Unix() > claims.Exp {
		return Claims{}, ErrTokenExpired
	}

	return claims, nil
}

func hmacSHA256Encode(message, secret []byte) []byte {
	mac := hmac.New(sha256.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}

func base64UrlEncode(data []byte) []byte {
	encoded := base64.URLEncoding.EncodeToString(data)
	return []byte(strings.TrimRight(encoded, "=")) // Remove any trailing '='
}

func base64UrlDecode(data []byte) ([]byte, error) {
	decoded, err := base64.URLEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	return decoded, nil
}
