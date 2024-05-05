package connection

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"google.golang.org/api/option"
)

type Connection struct {
	Firebase    *firebase.App
	Firestore   *firestore.Client
	AIAssistant *http.Client
}

func New(cfg *config.Configs) (*Connection, error) {
	ctx := context.Background()
	var firebaseApp *firebase.App
	var err error

	if cfg.App.Env != "local" {
		firebaseApp, err = firebase.NewApp(ctx, nil)
	} else {
		sa := option.WithCredentialsFile("quitsmoke-20141-firebase-adminsdk-ugo14-c5730ea21d.json")
		firebaseApp, err = firebase.NewApp(ctx, nil, sa)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %v", err)
	}

	firestoreClient, err := firebaseApp.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Firestore client: %v", err)
	}

	aiAssistantHTTPClient := http.DefaultClient

	return &Connection{
		Firebase:    firebaseApp,
		Firestore:   firestoreClient,
		AIAssistant: aiAssistantHTTPClient,
	}, nil
}

func (c *Connection) Close() {
	_ = c.Firestore.Close()
}
