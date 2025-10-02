package auth

import (
    "errors"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID   int    `json:"user_id"`
    Email    string `json:"email"`
    UserType string `json:"user_type"`
    jwt.RegisteredClaims
}

// Generate JWT token
func GenerateToken(userID int, email string, userType string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "your-default-secret-key" // Fallback for development
    }

    expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

    claims := &Claims{
        UserID:   userID,
        Email:    email,
        UserType: userType,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// Verify JWT token
func VerifyToken(tokenString string) (*Claims, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "your-default-secret-key"
    }

    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
