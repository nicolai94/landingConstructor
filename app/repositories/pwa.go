package repositories

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"landingConstructor/app/domain/dao"
)

type PwaRepository interface {
	CreatePwa(pwa *dao.Pwa) error
	GetPwaById(id string) (dao.Pwa, error)
	CreatePreLanding(preLanding *dao.PreLanding, pwa *dao.Pwa) error
}

type PwaRepositoryImpl struct {
	db *gorm.DB
}

func (p PwaRepositoryImpl) CreatePwa(pwa *dao.Pwa) error {
	if err := p.db.Create(&pwa).Error; err != nil {
		log.Error("Got an error when create PWA. Error: ", err)
		return err
	}
	return nil
}

func (p PwaRepositoryImpl) CreatePreLanding(preLanding *dao.PreLanding, pwa *dao.Pwa) error {
	if err := p.db.Create(&preLanding).Error; err != nil {
		log.Error("Got an error when create PWA. Error: ", err)
		return err
	}

	if err := p.db.Model(&pwa).Update("pre_landing_id", preLanding.ID).Error; err != nil {
		log.Error("Got an error when updating Pwa. Error: ", err)
		return err
	}

	return nil
}

func (p PwaRepositoryImpl) GetPwaById(id string) (dao.Pwa, error) {
	pwa := dao.Pwa{
		ID: uuid.UUID{},
	}
	err := p.db.First(&pwa).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.Pwa{}, err
	}
	return pwa, nil
}

func PwaRepositoryInit(db *gorm.DB) *PwaRepositoryImpl {
	//if err := db.AutoMigrate(&dao.Pwa{}, &dao.PreLanding{}); err != nil {
	//	log.Error("Got an error when connect to database. Error: ", err)
	//}
	//db.Debug().AutoMigrate(&dao.Pwa{}, &dao.PreLanding{})
	return &PwaRepositoryImpl{
		db: db,
	}
}
