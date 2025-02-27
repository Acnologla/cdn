package port

import "github.com/Acnologla/cdn/internal/core/domain"

type Cache interface {
	Set(key string, value *domain.File)
	Get(key string) (*domain.File, bool)
	Delete(key string)
	Clear()
}
