package configs

import (
	"bottlehub/graph/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("UserBase").Collection(collectionName)
}

func (db *DB) ctxDeferHelper(collectionName string) (*mongo.Collection, context.Context) {
	collection := colHelper(db, collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return collection, ctx
}

func (db *DB) resErrHelper(collectionName string, input any) (*mongo.InsertOneResult, error) {
	collection, ctx := db.ctxDeferHelper(collectionName)

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return res, err
}

func (db *DB) CreateComment(input *model.NewComment) (*model.Comment, error) {
	res, err := db.resErrHelper("comments", input)

	comment := &model.Comment{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Text:      input.Text,
		CommentBy: &model.User{},
		CommentOn: &model.Post{},
	}

	return comment, err
}

func (db *DB) CreatePost(input *model.NewPost) (*model.Post, error) {
	res, err := db.resErrHelper("posts", input)

	post := &model.Post{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		PostedBy:    &model.User{},
		ImageURL:    input.ImageURL,
		Description: input.Description,
		Likes:       input.Likes,
	}

	return post, err
}

func (db *DB) CreateUser(input *model.NewUser) (*model.User, error) {
	res, err := db.resErrHelper("users", input)

	user := &model.User{
		ID:             res.InsertedID.(primitive.ObjectID).Hex(),
		Username:       input.Username,
		Name:           input.Name,
		About:          input.About,
		Email:          input.Email,
		AvatarImageURL: input.AvatarImageURL,
		PublicWallet:   input.PublicWallet,
		PrivateWallet:  input.PrivateWallet,
	}

	return user, err
}

func (db *DB) GetComments() ([]*model.Comment, error) {
	collection, ctx := db.ctxDeferHelper("comments")
	var comments []*model.Comment

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleComment *model.Comment
		if err = res.Decode(&singleComment); err != nil {
			log.Fatal(err)
		}
		comments = append(comments, singleComment)
	}

	return comments, err
}

func (db *DB) GetPosts() ([]*model.Post, error) {
	collection, ctx := db.ctxDeferHelper("posts")
	var posts []*model.Post

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singlePost *model.Post
		if err = res.Decode(&singlePost); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, singlePost)
	}

	return posts, err
}

func (db *DB) GetUsers() ([]*model.User, error) {
	collection, ctx := db.ctxDeferHelper("users")
	var users []*model.User

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleUser *model.User
		if err = res.Decode(&singleUser); err != nil {
			log.Fatal(err)
		}
		users = append(users, singleUser)
	}

	return users, err
}

func (db *DB) SingleComment(ID string) (*model.Comment, error) {
	collection, ctx := db.ctxDeferHelper("comments")
	var comment *model.Comment

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&comment)

	return comment, err
}

func (db *DB) SinglePost(ID string) (*model.Post, error) {
	collection, ctx := db.ctxDeferHelper("posts")
	var post *model.Post

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&post)

	return post, err
}

func (db *DB) SingleUser(ID string) (*model.User, error) {
	collection, ctx := db.ctxDeferHelper("users")
	var user *model.User

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

	return user, err
}
