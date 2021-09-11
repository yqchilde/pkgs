package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/yqchilde/gin-skeleton/internal/model"
	"github.com/yqchilde/gin-skeleton/internal/store"
	"github.com/yqchilde/gin-skeleton/pkg/app"
	"github.com/yqchilde/gin-skeleton/pkg/auth"
	"github.com/yqchilde/gin-skeleton/pkg/conf"
)

type ApplicationService struct {
	Ctx *gin.Context
}

func NewApplicationService(ctx *gin.Context) *ApplicationService {
	return &ApplicationService{
		Ctx: ctx,
	}
}

func (a *ApplicationService) Register(email, password, firstName, lastName string) error {
	var (
		developerStore = store.NewDBDeveloper()
	)

	// TODO: Verify password rule
	pwd, err := auth.HashAndSalt(password)
	if err != nil {
		return errors.Wrapf(err, "[service.application] encrypt password err")
	}

	developer := &model.Developer{
		UserID:    uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  pwd,
	}

	err = developerStore.Insert(developer)
	if err != nil {
		return err
	}

	return nil
}

func (a *ApplicationService) Login(email, password string) (tokenStr string, err error) {
	developer, err := a.GetByEmail(email)
	if err != nil {
		return "", errors.Wrapf(err, "[service.application] get developer by email")
	}

	if !auth.ComparePasswords(developer.Password, password) {
		return "", errors.Wrapf(err, "[service.application] invalid password")
	}

	payload := map[string]interface{}{"user_id": developer.UserID, "email": developer.Email}
	tokenStr, err = app.Sign(payload, conf.Conf.App.JwtSecret, conf.Conf.App.JwtExpireTime)
	if err != nil {
		return "", errors.Wrapf(err, "[service.application] generate jwt token sign err")
	}

	return tokenStr, nil
}

func (a *ApplicationService) GetByEmail(email string) (*model.Developer, error) {
	developer, err := store.NewDBDeveloper().GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return developer, nil
}

func (a *ApplicationService) CreateApp(userID, appName string) (*model.Application, error) {
	application := &model.Application{
		AppID:   uuid.New().String(),
		AppName: appName,
		Creator: userID,
	}
	//application.AppKey = auth.GenerateAppKey(application.AppID)
	//application.AppSecret = auth.GenerateAppSecret(application.AppID)

	err := store.NewDBApplication().Insert(application)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *ApplicationService) DeleteApp(appID string) error {
	return store.NewDBApplication().DeleteByAppID(appID)
}
