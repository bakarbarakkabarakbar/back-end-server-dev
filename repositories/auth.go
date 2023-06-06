package repositories

import (
	"back-end-server-dev/entities"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(dbCrud *gorm.DB) AuthRepo {
	return AuthRepo{
		db: dbCrud,
	}
}

func (ar AuthRepo) GetActorByUsername(username *string) (entities.Actor, error) {
	var admin entities.Actor
	var err = ar.db.Where("username = ?", username).Find(&admin).Error
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (ar AuthRepo) GetLastActorSessionByToken(token *string) (entities.ActorSession, error) {
	var session entities.ActorSession
	var err = ar.db.Where(&entities.ActorSession{Token: *token}).Last(&session).Error
	if err != nil {
		return session, err
	}
	return session, nil
}

func (ar AuthRepo) CreateActorSession(session *entities.ActorSession) error {
	err := ar.db.Model(&entities.ActorSession{}).Create(session).Error
	return err
}
