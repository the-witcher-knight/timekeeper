package auth

import (
	"encoding/json"
	"fmt"
	"slices"
)

type Claims struct {
	Sub         string `json:"sub"`
	Exp         int64  `json:"exp"`
	Iat         int64  `json:"iat"`
	ExtraClaims map[string]interface{}
}

func marshalClaims(claims Claims) ([]byte, error) {
	valueMap := make(map[string]interface{})
	valueMap["sub"] = claims.Sub
	valueMap["exp"] = claims.Exp
	valueMap["iat"] = claims.Iat

	for k, v := range claims.ExtraClaims {
		valueMap[k] = v
	}

	rs, err := json.Marshal(valueMap)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func unmarshalClaims(bytes []byte) (Claims, error) {
	mapClaim := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &mapClaim); err != nil {
		return Claims{}, err
	}

	claims := Claims{
		Sub:         fmt.Sprintf("%s", mapClaim["sub"]),
		Exp:         int64(mapClaim["exp"].(float64)),
		Iat:         int64(mapClaim["iat"].(float64)),
		ExtraClaims: make(map[string]interface{}),
	}

	for k, v := range mapClaim {
		if !slices.Contains([]string{"sub", "exp", "iat"}, k) {
			claims.ExtraClaims[k] = v
		}
	}

	return claims, nil
}
