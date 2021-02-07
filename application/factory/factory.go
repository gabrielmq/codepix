package factory

import (
	"github.com/gabrielmq/codepix/application/usecase"
	"github.com/gabrielmq/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}
	return usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixKeyRepository:      pixRepository,
	}
}
