package repository

import (
	"context"
	"gPRC/dto"
	"gPRC/models"
	"gPRC/models/mongodb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type policyRepository struct {
	collection *mongo.Collection
}

func NewPolicyRepository() PolicyRepository {
	return &policyRepository{collection: mongodb.GeColl("policy")}
}

type PolicyRepository interface {
	GetPolicies(queryRequest *dto.QueryRequest) []models.Policy
}

func (repo *policyRepository) GetPolicies(queryRequest *dto.QueryRequest) []models.Policy {
	var policies []models.Policy
	// If the queryRequest is nil, log and get all policies
	if queryRequest == nil {
		log.Print("queryRequest is nil, fetching all policies")
		repo.getAllPolicies(&policies) // Pass policies by reference
		return policies
	}
	log.Print(queryRequest.String())
	// Handle filter
	filter := bson.D{}
	if len(queryRequest.Filter) > 0 {
		var filterKey, filterValue string
		for key, value := range queryRequest.Filter {
			filterKey = key
			filterValue = value
		}
		filter = bson.D{{filterKey, bson.D{
			{"$regex", filterValue}, /* Filter with no absolute correction */
			{"$options", "i"},       /* case-insensitive */
		}}}
	}

	opts := options.Find()
	// Handle sort
	opts.SetSort(bson.D{{"name", 1}}) /* Default sort by name */
	opts.SetLimit(0)                  /* Get all */
	opts.SetSkip(0)                   /* No skip */
	if len(queryRequest.Sort) > 0 {
		var sortFields bson.D
		for _, sort := range queryRequest.Sort {
			order := 1
			if sort.Order == "desc" {
				order = -1
			}
			sortFields = append(sortFields, bson.E{sort.Field, order})
		}
		opts.SetSort(sortFields)
	}
	// Handle paging
	if queryRequest.PageSize > 0 {
		opts.SetLimit(int64(queryRequest.PageSize))
	}
	if queryRequest.Page > 0 {
		skip := int64(queryRequest.Page-1) * int64(queryRequest.PageSize)
		opts.SetSkip(skip)
	}
	cursor, err := repo.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &policies); err != nil {
		panic(err)
	}

	return policies
}

func (repo *policyRepository) getAllPolicies(policies *[]models.Policy) {
	// Find all policies
	cursor, err := repo.collection.Find(context.TODO(), bson.D{},
		options.Find().SetSort(bson.D{{"name", 1}}))
	if err != nil {
		panic(err)
	}

	defer cursor.Close(context.TODO())

	// Decode the cursor into the policies slice
	if err = cursor.All(context.TODO(), policies); err != nil {
		panic(err)
	}
}
