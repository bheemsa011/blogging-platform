# Blogging Platform

This application provides a simple API collection for managing blog posts. It offers endpoints for creating, reading, updating, and deleting blog posts.

## Application Details

- **Application URL**: [http://localhost:8080](http://localhost:8080)
- **Port**: 8080

## Endpoints

### Get All Posts

- **Endpoint**: `/posts`
- **Method**: GET
- **Description**: Retrieve a list of all blog posts.
- **Response**: Returns a JSON array of blog posts.
- **Sample Response**: 
{
    "_id": "6523380438210d55f4fa03ba",
    "title": "My blog",
    "content": "I have written a blog"
}

### Get a Post by ID

- **Endpoint**: `/post/:id`
- **Method**: GET
- **Description**: Retrieve a specific blog post by its unique identifier.
- **Parameters**: `id` - The ID of the blog post to retrieve.
- **Response**: Returns a JSON object representing the blog post.

### Create a New Post

- **Endpoint**: `/post`
- **Method**: POST
- **Description**: Create a new blog post.
- **Request Body**: JSON object containing the blog post's title and content.
- **Response**: Returns a JSON object with the ID of the newly created blog post.
- **Sample Payload**: 
{
    "title": "My blog", 
    "content": "I have written a blog"
}

### Update a Post

- **Endpoint**: `/post/:id`
- **Method**: PUT
- **Description**: Update an existing blog post by its unique identifier.
- **Parameters**: `id` - The ID of the blog post to update.
- **Request Body**: JSON object containing the updated title and content.
- **Response**: Returns a success message.
- **Sample Payload**: 
{
    "title": "My blog", 
    "content": "I have written a blog"
}

### Delete a Post

- **Endpoint**: `/post/:id`
- **Method**: DELETE
- **Description**: Delete a blog post by its unique identifier.
- **Parameters**: `id` - The ID of the blog post to delete.
- **Response**: Returns a success message.
