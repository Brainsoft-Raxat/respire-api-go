package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

const (
	userCollection = "users_test"
)

type userRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewUserRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) UserRepository {
	return &userRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	doc, err := r.client.Collection(userCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get user document: %v", err)
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		r.logger.Errorf("Failed to convert user data to struct: %v", err)
		return nil, err
	}

	user.ID = doc.Ref.ID

	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	iter := r.client.Collection(userCollection).Where("email", "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		r.logger.Errorf("Failed to get user document: %v", err)
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		r.logger.Errorf("Failed to convert user data to struct: %v", err)
		return nil, err
	}

	user.ID = doc.Ref.ID

	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (string, error) {
	docRef, _, err := r.client.Collection(userCollection).Add(ctx, user)
	if err != nil {
		r.logger.Errorf("Failed to create user: %v", err)
		return "", err
	}

	return docRef.ID, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, user *models.User) error {
	_, err := r.client.Collection(userCollection).Doc(id).Set(ctx, user)
	if err != nil {
		r.logger.Error("Failed to update user: %v", err)
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := r.client.Collection(userCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete user: %v", err)
		return err
	}

	return nil
}
