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

### Update a Post

- **Endpoint**: `/post/:id`
- **Method
