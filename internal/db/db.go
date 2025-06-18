package db

import (
	"context"
	"github.com/dana-team/axiom-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InsertClusterInfoToMongo(logger logr.Logger, clusterInfo v1alpha1.ClusterInfo) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("<mongo-uri>"),
	)
	if err != nil {
		logger.Error(err, "Failed to connect to MongoDB")
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database("axiom").Collection("clusterInfo")
	filter := bson.M{"clusterID": clusterInfo.Status.ClusterID}
	update := bson.M{"$set": clusterInfo}
	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logger.Error(err, "Failed to insert cluster info to MongoDB")
		return
	}
	logger.Info("Inserted cluster info to MongoDB")
}
