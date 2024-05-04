package repository

import (
	"context"
	"time"

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
	sessionCollection = "sessions"
)

type sessionRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewSessionRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) *sessionRepository {
	return &sessionRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *sessionRepository) CreateSession(ctx context.Context, session *models.SmokeSession) (string, error) {
	docRef := s.client.Collection(sessionCollection).NewDoc()
	session.ID = docRef.ID

	_, err := docRef.Set(ctx, session)
	if err != nil {
		s.logger.Errorf("Failed to create smpke session: %v", err)
		return "", err
	}

	return docRef.ID, nil
}

func (s *sessionRepository) GetSessionByID(ctx context.Context, id string) (*models.SmokeSession, error) {
	doc, err := s.client.Collection(sessionCollection).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			s.logger.Errorf("Smoke Session document not found: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.NotFoundError, "session document not found")
		}

		s.logger.Errorf("Failed to get session document: %v", err)
		return nil, err
	}

	var session models.SmokeSession
	if err := doc.DataTo(&session); err != nil {
		s.logger.Errorf("Failed to convert user data to struct: %v", err)
		return nil, err
	}

	session.ID = doc.Ref.ID

	return &session, nil
}

func (s *sessionRepository) GetSessionsByUserID(ctx context.Context, userID string) (SessionsInfo, error) {
	query := s.client.Collection(sessionCollection).Where("user_id", "==", userID)
	iter := query.Documents(ctx)
	var sessions []*models.SmokeSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Errorf("Failed to iterate smoke session documents: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to iterate smoke session documents")
		}

		var session models.SmokeSession
		if err := doc.DataTo(&session); err != nil {
			s.logger.Errorf("Failed to convert smoke session data to struct: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to convert smoke session data to struct")
		}

		session.ID = doc.Ref.ID
		sessions = append(sessions, &session)
	}

	return sessions, nil
}

type SessionsInfo []*models.SmokeSession

func (s SessionsInfo) Sum() int {
	sum := 0
	for _, si := range s {
		sum += si.Count
	}
	return sum
}

func (s SessionsInfo) Count() int {
	return len(s)
}

func (s SessionsInfo) Timestamps() (times []time.Time) {
	for _, si := range s {
		times = append(times, si.Timestamp)
	}
	return
}

func (s *sessionRepository) GetSessionsByUserIDAndDateRange(ctx context.Context, userID string, dateRange [2]time.Time) (SessionsInfo, error) {
	if dateRange[1].Compare(dateRange[0]) == -1 {
		s.logger.Errorf("Wrong date range format: first date should be earlier than the second")
	}
	query := s.client.Collection(sessionCollection).Where("user_id", "==", userID)
	iter := query.Documents(ctx)
	s.logger.Info(dateRange)
	var sessions SessionsInfo
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Info(dateRange)
			s.logger.Errorf("Failed to iterate smoke session documents: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to iterate smoke session documents")
		}

		var session models.SmokeSession
		if err := doc.DataTo(&session); err != nil {
			s.logger.Errorf("Failed to convert smoke session data to struct: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to convert smoke session data to struct")
		}
		s.logger.Info(session.Timestamp)
		if session.Timestamp.Compare(dateRange[0]) >= 0 && session.Timestamp.Compare(dateRange[1]) <= 0 {
			session.ID = doc.Ref.ID
			sessions = append(sessions, &session)
		}
	}

	return sessions, nil
}

const (
	Week  = 7
	Month = 30
)

func GetWeek(today time.Time) [2]time.Time {
	weekAgo := time.Date(today.Year(), today.Month(), (today.Day() - Week), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), today.Location())
	return [2]time.Time{weekAgo, today}
}

func GetMonth(today time.Time) [2]time.Time {
	monthAgo := time.Date(today.Year(), today.Month(), (today.Day() - Month), today.Hour(), today.Minute(), today.Second(), today.Nanosecond(), today.Location())
	return [2]time.Time{monthAgo, today}
}

func (r *sessionRepository) UpdateSession(ctx context.Context, id string, session *models.SmokeSession) error {
	model, err := CreateUpdateMap(session)
	if err != nil {
		return err
	}
	count, ok := model["count"]

	if ok && count == 0 {
		if err := r.DeleteSession(ctx, id); err != nil {
			r.logger.Errorf("Failed to update empty session")
			return err
		}
		return nil
	}
	_, err = r.client.Collection(sessionCollection).Doc(id).Set(ctx, model, firestore.MergeAll)
	if err != nil {
		r.logger.Error("Failed to update session: %v", err)
		return err
	}

	return nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, id string) error {
	_, err := r.client.Collection(sessionCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete session: %v", err)
		return err
	}

	return nil
}

func (s *sessionRepository) GetUserStat(ctx context.Context, userID string) (int, int, int, error) {

	return 0, 0, 0, nil
}
