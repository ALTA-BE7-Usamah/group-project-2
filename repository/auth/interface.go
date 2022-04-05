package auth

type AuthRepositoryInterface interface {
	// login menggunakan email dan password user
	Login(email string, password string) (string, error)
}
