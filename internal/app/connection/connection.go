package connection

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"google.golang.org/api/option"
)

type Connection struct {
	Firebase  *firebase.App
	Firestore *firestore.Client
}

func New(cfg *config.Configs) (*Connection, error) {
	ctx := context.Background()

	sa := option.WithCredentialsFile("quitsmoke-20141-firebase-adminsdk-ugo14-c5730ea21d.json")
	firebaseApp, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	firestoreClient, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &Connection{
		Firebase:  firebaseApp,
		Firestore: firestoreClient,
	}, nil
}

func (c *Connection) Close() {
	_ = c.Firestore.Close()
}