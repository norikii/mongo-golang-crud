package handlers

import (
	"context"
	"fmt"
	"github.com/tatrasoft/mongo-golang-crud/database"
	"github.com/tatrasoft/mongo-golang-crud/models"
	blogpb "github.com/tatrasoft/mongo-golang-crud/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const name  =

type BlogServiceServer struct {}

func getCollection() (*mongo.Collection, error) {
	collection, err := database.SpecifyCollection(database.DBCli, "mydb", "blog")
	if err != nil {
		return nil, fmt.Errorf("could to get collection: %v", err)
	}

	return collection, nil
}

func (bss *BlogServiceServer) CreateBlog(
	ctx context.Context,
	req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {
		// to access the struct with nil check
		blog := req.GetBlog()
		// converting BlogItem type to BSON
		data := models.BlogItem{
			ID:       primitive.ObjectID{},
			AuthorID: blog.GetAuthorId(),
			Content:  blog.GetContent(),
			Title:    blog.GetTitle(),
		}



}
