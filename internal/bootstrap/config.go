package bootstrap

import "go-backend-clean-arch/configs"

func ConfigLoad(env configs.Env) *configs.Config {
	return configs.Load(env)
}
