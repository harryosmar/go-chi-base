package factories

import (
	"github.com/harryosmar/go-chi-base/app/user/handlers"
	"github.com/harryosmar/go-chi-base/app/user/repositories"
	"github.com/harryosmar/go-chi-base/app/user/services"
	"github.com/spf13/viper"
)

func MakeAccountRepository() repositories.AccountRepository {
	return repositories.NewProfileRepository(nil)
}

func MakeUserService() services.UserService {
	return services.NewUserService(
		viper.GetString("token.iss"),
		MakeAccountRepository(),
	)
}

func MakeUserHandler() *handlers.UserHandler {
	return handlers.NewUserHandler(MakeUserService())
}
