package port

import entity "github.com/tama-jp/basee_web/internal/domain/entities"

type UserRolePort interface {
	FindId(id uint) (*entity.UserRole, error)
	FindList() (*[]entity.UserRole, error)
}
