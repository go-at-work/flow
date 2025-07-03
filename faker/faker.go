package faker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arisromil/flow/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var Password = "$223232ddsfsfsdfsdffs"

var leeterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = leeterRunes[rand.Intn(len(leeterRunes))]
	}
	return string(b)
}

func RandomStringLower(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = leeterRunes[rand.Intn(len(leeterRunes)/2)]
	}
	return string(b)
}

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Username() string {
	return fmt.Sprintf("user%s", RandomString(RandInt(2, 10)))
}

func UUID() string {
	return uuid.GenerateUUID()
}

func Email() string {
	return fmt.Sprintf("%s@example.com", RandomStringLower(RandInt(2, 10)))
}

func ID() string {
	return fmt.Sprintf("%s-%s-%s-%s", RandomString(4), RandomString(4), RandomString(4), RandomString(4))
}
