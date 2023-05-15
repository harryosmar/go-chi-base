package factories

import (
	"github.com/harryosmar/go-chi-base/app/user/handlers"
	"github.com/harryosmar/go-chi-base/app/user/repositories"
	"github.com/harryosmar/go-chi-base/app/user/services"
)

func MakeUserRepository() repositories.CredentialRepository {
	return repositories.NewCredentialRepository(nil)
}

func MakeProfileRepository() repositories.ProfileRepository {
	return repositories.NewProfileRepository(nil)
}

func MakeUserService() services.UserService {
	return services.NewUserService(
		MakeUserRepository(),
		MakeProfileRepository(),
	)
}

func MakeUserHandler() *handlers.UserHandler {
	return handlers.NewUserHandler(MakeUserService())
}
