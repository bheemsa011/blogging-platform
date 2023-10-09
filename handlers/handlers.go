package handlers

import (
	"blogging-platform/model"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var db *mongo.Database

func SetDB(database *mongo.Database) {
	db = database
}

func GetPosts(c *gin.Context) {
	// Retrieve the list of all blog posts
	posts := []model.Post{}

	cursor, err := db.Collection("posts").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		post := model.Post{}
		if err := cursor.Decode(&post); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Parse the post ID into an ObjectID (assuming you are using ObjectID as the ID type)
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Search for the post in the database by its ID
	post := model.Post{}
	err = db.Collection("posts").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	// Parse the request body to get the new post data
	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get a reference to the MongoDB collection
	collection := db.Collection("posts")

	// Insert the new post into the collection
	result, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created post with its new ObjectID
	c.JSON(http.StatusCreated, gin.H{"insertID": result.InsertedID})
}

func UpdatePost(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")
	fmt.Println("Received postID:", postID) // Debug print
	// Parse the post ID into an ObjectID
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Define a structure to hold the updated post data
	updatedPost := model.Post{}
	if err := c.BindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create a filter for the specific post
	filter := bson.M{"_id": objID}

	// Define an update query to set the new values
	update := bson.M{
		"$set": bson.M{
			"title":   updatedPost.Title,
			"content": updatedPost.Content,
		},
	}

	// Update the post in the database
	_, err = db.Collection("posts").UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func DeletePostByID(c *gin.Context) {
	// Get the post ID from the URL parameter
	postID := c.Param("id")

	// Parse the post ID into an ObjectID
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Delete the post from the database
	filter := bson.M{"_id": objID}
	_, err = db.Collection("posts").DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
