package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	userCollection = "users"
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
		if status.Code(err) == codes.NotFound {
			r.logger.Errorf("User document not found: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.NotFoundError, "user document not found")
		}

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
		if status.Code(err) == codes.NotFound {
			r.logger.Errorf("User document not found: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.NotFoundError, "user document not found")
		}

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

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	iter := r.client.Collection(userCollection).Where("username", "==", username).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		r.logger.Errorf("Failed to get user document: %v", err)
		return nil, apperror.NewErrorInfo(ctx, errcodes.NotFoundError, "user document not found")
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
	docRef := r.client.Collection(userCollection).NewDoc()
	user.ID = docRef.ID

	_, err := docRef.Set(ctx, user)
	if err != nil {
		r.logger.Errorf("Failed to create user: %v", err)
		return "", err
	}

	return docRef.ID, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, user *models.User) error {
	model, err := CreateUpdateMap(user)
	if err != nil {
		return err
	}

	_, err = r.client.Collection(userCollection).Doc(id).Set(ctx, model, firestore.MergeAll)
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

func (r *userRepository) SearchUsersByUsernamePrefix(ctx context.Context, excludeID string, usernamePrefix string, limit int) ([]*models.ShortUser, error) {
	if limit <= 0 {
		limit = 10 // default limit if none specified or invalid value is provided
	}

	// lowerUsername := strings.ToLower(usernamePrefix)
	upperBound := usernamePrefix + "\uf8ff" // A high Unicode character to set upper range for query

	query := r.client.Collection("users").
		Where("status", "==", models.STATUS_COMPLETED).
		Where("username", ">=", usernamePrefix).
		Where("username", "<=", upperBound)

	if excludeID != "" {
		query = query.Where("id", "!=", excludeID)
	}

	query = query.Limit(limit)

	iter := query.Documents(ctx)
	var users []*models.ShortUser
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			r.logger.Errorf("Failed to retrieve user document: %v", err)
			return nil, err // Handle error appropriately
		}

		var user models.ShortUser
		if err := doc.DataTo(&user); err != nil {
			r.logger.Errorf("Failed to convert user document to struct: %v", err)
			return nil, err
		}
		user.ID = doc.Ref.ID
		users = append(users, &user)
	}

	return users, nil
}

func (s *userRepository) GetShortUsersByIDs(ctx context.Context, ids []string) ([]*models.ShortUser, error) {
	var users []*models.ShortUser
	var docRefs []*firestore.DocumentRef

	for _, id := range ids {
		docRef := s.client.Collection(userCollection).Doc(id)
		docRefs = append(docRefs, docRef)
	}

	snapshots, err := s.client.GetAll(ctx, docRefs)
	if err != nil {
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get user documents")
	}

	for _, snap := range snapshots {
		if snap.Exists() {
			var user models.ShortUser
			if err := snap.DataTo(&user); err != nil {
				return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to convert user data to struct")
			}

			user.ID = snap.Ref.ID

			users = append(users, &user)
		}
	}

	return users, nil
}
