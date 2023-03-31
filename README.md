# Culineira-Backend
created using golang gin
made by Leonardho R Sitanggang

# How to run this app
> go run main.go

# Config 
> Collection of variable and configuration that will be used in this app

# Helpers
> Set of function that will be use to generate, validation, or manipulating data

# Middlewares
> For data authentication and authorization, so users access can be intercepted before they can make the request without login

# Migrations
> Database structure (DDL)

# Modules
> In this app, modules is a collection of unit of code that will be separated based on its functionality
Example :
- Recipe, Steps, and Ingredient will be grouping into one modules of recipes
- Like and Comment will be grouping into one modules of contents

# Handlers
> Processing incoming HTTP requests from the router. One handler for one API endpoint. In this function, data will be transferred to repositories 
so it can be manipulated into the database

# Models
> Set of data and its type before it can be processed to or from the database

# Repositories
> The function that contains logic or query for a database such as Create, Read, Update, and Delete. In this function, it will receive data from 
handler and send it to the database
