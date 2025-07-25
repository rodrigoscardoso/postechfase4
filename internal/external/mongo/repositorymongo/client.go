package repositorymongo

import (
	"context"
	dto "post-tech-challenge-10soat/internal/dto/client"
	"post-tech-challenge-10soat/internal/external/mongo/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientMongoRepositoryImpl struct {
	collection *mongo.Collection
}

func NewClientMongoRepositoryImpl(db *mongo.Database) ClientMongoRepositoryImpl {
	return ClientMongoRepositoryImpl{
		collection: db.Collection("client"),
	}
}

func (repository ClientMongoRepositoryImpl) CreateClient(ctx context.Context, client dto.CreateClientDTO) (dto.ClientDTO, error) {
	clientModel := model.ClientModel{
		Cpf:   client.Cpf,
		Name:  client.Name,
		Email: client.Email,
	}
	res, err := repository.collection.InsertOne(ctx, clientModel)
	if err != nil {
		return dto.ClientDTO{}, err
	}
	clientModel.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return clientModel.ToDTO(), nil
}

func (repository ClientMongoRepositoryImpl) GetClientByCpf(ctx context.Context, cpf string) (dto.ClientDTO, error) {
	var clientModel model.ClientModel
	err := repository.collection.FindOne(ctx, bson.M{"cpf": cpf}).Decode(&clientModel)
	if err != nil {
		return dto.ClientDTO{}, err
	}
	return clientModel.ToDTO(), nil
}

func (repository ClientMongoRepositoryImpl) GetClientById(ctx context.Context, id string) (dto.ClientDTO, error) {
	var clientModel model.ClientModel
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ClientDTO{}, err
	}
	err = repository.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&clientModel)
	if err != nil {
		return dto.ClientDTO{}, err
	}
	return clientModel.ToDTO(), nil
}
