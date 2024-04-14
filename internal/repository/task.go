package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

const TaskCollection = "tasks"

type taskRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewTaskRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) TaskRepository {
	return &taskRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *taskRepository) GetTaskByID(ctx context.Context, friendsGroupID, challengeID, id string) (*models.Task, error) {
	doc, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(challengeID).Collection(TaskCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get task by ID: %v", err)
		return nil, err
	}

	var task models.Task
	if err := doc.DataTo(&task); err != nil {
		r.logger.Errorf("Failed to convert task data to struct: %v", err)
		return nil, err
	}

	task.ID = doc.Ref.ID
	return &task, nil
}

func (r *taskRepository) GetTasksByChallengeID(ctx context.Context, friendsGroupID, challengeID string) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)

	snapshot, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(challengeID).Collection(TaskCollection).Documents(ctx).GetAll()
	if err != nil {
		r.logger.Errorf("Failed to get tasks by challenge ID: %v", err)
		return nil, err
	}

	for _, doc := range snapshot {
		var task models.Task
		if err := doc.DataTo(&task); err != nil {
			r.logger.Errorf("Failed to convert task data to struct: %v", err)
			return nil, err
		}
		task.ID = doc.Ref.ID
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *taskRepository) CreateTask(ctx context.Context, friendsGroupID, challengeID string, task *models.Task) error {
	_, _, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(challengeID).Collection(TaskCollection).Add(ctx, task)
	if err != nil {
		r.logger.Errorf("Failed to create task: %v", err)
		return err
	}

	return nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, friendsGroupID, challengeID, id string, task *models.Task) error {
	_, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(challengeID).Collection(TaskCollection).Doc(id).Set(ctx, task)
	if err != nil {
		r.logger.Errorf("Failed to update task: %v", err)
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, friendsGroupID, challengeID, id string) error {
	_, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(challengeID).Collection(TaskCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete task: %v", err)
		return err
	}

	return nil
}
