// file: data/data.go
package data

import (
	"log"
	"math/rand"
)

const (
	fieldSize = 10
)

type UserInfo struct {
	FieldA string
	FieldB string
	FieldC string
	FieldD string
	FieldE string
	FieldF string
	FieldG int
	FieldH int
	Text   string
}

func GeneratePayload(userCount int, compressionRatio int) []UserInfo {
	if compressionRatio <= 0 || compressionRatio >= 100 {
		log.Fatalf("Compression ratio must be between (0-100), got %d\n", compressionRatio)
	}
	if userCount < 0 {
		log.Fatalf("User count negative, got %d\n", userCount)
	}

	uniqueCount := userCount * (100 - compressionRatio) / 100
	if uniqueCount <= 0 {
		uniqueCount = 1
	}

	users := make([]UserInfo, 0, uniqueCount)
	for i := 0; i < uniqueCount; i++ {
		users = append(users, UserInfo{
			FieldA: compressableString(100, compressionRatio),
			FieldB: compressableString(100, compressionRatio),
			FieldC: compressableString(100, compressionRatio),
			FieldD: compressableString(100, compressionRatio),
			FieldE: compressableString(100, compressionRatio),
			FieldF: compressableString(100, compressionRatio),
			FieldG: rand.Int(),
			FieldH: rand.Int(),
			Text:   compressableString(300, compressionRatio),
		})
	}
	println("Generate payload of count ", userCount, " with ", uniqueCount, " unique data points and ~ compression ratio ", compressionRatio)

	userPayload := make([]UserInfo, 0, userCount)
	for i := 0; i < userCount; i++ {
		user := users[i%len(users)]
		userPayload = append(userPayload, user)
	}
	return userPayload
}

func compressableString(size int, ratio int) string {
	duplicateCount := size * (100 - ratio) / 100
	if duplicateCount <= 0 {
		duplicateCount = 1
	}
	pattern := randomString(duplicateCount)
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = pattern[i%len(pattern)]
	}
	return string(bytes)
}

func randomString(size int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = len(letters)
	result := make([]byte, size)
	for i := range result {
		result[i] = letters[rand.Intn(length)]
	}
	return string(result)
}
