package authentication

func HashPassword(password string) (string, error) {
 hash, err := argon2id.
}

func VerifyPassword() {

}
// VerifyPassword checks if the provided password matches the stored hash.