package svc

import "add/internal/config"

type ServiceContext struct {
        Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
        return &ServiceContext{
                Config: c,
        }
}
