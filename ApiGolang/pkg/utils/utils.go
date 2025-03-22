package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func DecodeBody[customType any](body io.ReadCloser) (*customType, error) {
	var data customType
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return &data, err
	}
	return &data, nil
}

func SendJsonResp(w http.ResponseWriter, status int, data any) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(jsonResponse); err != nil {
		log.Printf("!!ATTENTION!! failed to write JSON response: %v", err)
		return
	}
}

func Hashing(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckHashing(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err == nil
}

func ExtractID(urlPath string, index int) (uint, error) {
	parts := strings.Split(urlPath, "/")

	if index < 0 || index >= len(parts) {
		return 0, errors.New("index out of range")
	}

	idStr := parts[index]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid id format")
	}

	return uint(id), nil
}
