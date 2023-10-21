<h1 align="center">
  GoLang API
  
  ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
  ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
  <br>
</h1>

<p align="center">
  <a href="#introduction">Introduction</a> •
  <a href="#installation-guide">Installation Guide</a> •
  <a href="#project-structure">Project Structure</a> •
  <a href="#api">API Reference</a> •
  <a href="#license">License</a> •
  <a href="#contributors">Contributors</a> 
</p>

<div align="center">

![GitHub Repo stars](https://img.shields.io/github/stars/Ctere1/golang-api)
![GitHub forks](https://img.shields.io/github/forks/Ctere1/golang-api)
![GitHub watchers](https://img.shields.io/github/watchers/Ctere1/golang-api)

</div>

## ℹ️ Introduction
- Product API for CRUD operations. 
- Written in GoLang and PostgreSQL. Bearer token authentication is used. 
- Database connection is made with GORM. Server is created with gorilla/mux. 


## 💾Installation Guide

- To clone and run this application, you'll need [Git](https://git-scm.com), [Go](https://go.dev/), [PostgreSQL](https://www.postgresql.org/) installed on your computer. From your command line:

    ```bash
    # Clone this repository
    $ git clone https://github.com/Ctere1/golang-api
    # Go into cmd folder
    $ cd cmd 
    # Run the app
    $ go run main.go
    ```

- You can change the database connection string in the `app.toml` file. Also, you can change the API listen address and port.   
    > After these steps, you can use the API with Postman or any other API testing tool.

## 🗺️Project Structure 

- You can take a part of it and edit it to suit your own needs and liking.  
- The project structure is as follows:    

  ```bash
  ├───cmd
  │   └───main.go
  ├───pkgs
  │   ├───api
  │   │   ├───router.go
  │   │   └───RestHandler.go
  │   ├───category
  │   │   ├───categoryApiHandler.go
  │   │   └───categoryController.go
  │   ├───configs
  │   │   └───configs.go
  │   ├───product
  │   │   ├───productApiHandler.go
  │   │   └───productController.go
  │   └───storage
  │       ├───categoryPostgres.go
  │       ├───productPostgres.go
  │       └───storage.go
  ├───app.toml
  ├───go.mod
  ├───go.sum
  ├───LICENSE
  ├───README.md
  └───postman_collection.json
  ``` 

## ⚡API
>**Note**   
See postman collection json for detailed information.


### **Product Endpoints**

| HTTP Verb   | Endpoint                    | Description                         |  
| :---------- | :-----------------------    |:----------------------------------  |    
| `GET`       | `/api/v1/product`           |  Returns All Products               |
| `GET`       | `/api/v1/product/{sku}`     |  Returns the product with {sku}     |
| `DELETE`    | `/api/v1/product/{sku}`     |  Deletes the product with {sku}     |
| `POST`      | `/api/v1/product/`          |  Creates and returns product        |
| `PUT`       | `/api/v1/product/`          |  Updates and returns product        |

>**Warning**   
Product sku is unique and required.

### **Category Endpoints**

| HTTP Verb   | Endpoint                    | Description                         |
| :---------- | :-----------------------    |:----------------------------------  |
| `GET`       | `/api/v1/category`          |  Returns All Categories             |
| `GET`       | `/api/v1/category/{id}`     |  Returns the category with {id}     |
| `DELETE`    | `/api/v1/category/{id}`     |  Deletes the category with {id}     |
| `POST`      | `/api/v1/category/`         |  Creates and returns category       |
| `PUT`       | `/api/v1/category/`         |  Updates and returns category       |

>**Warning**
Category name is unique and required.


### **Product Data Example**

```json
{
    "Name": "test1",
    "Price": "123",
    "Description": "test desc",
    "Sku": "sku1",
    "CategoryId": "1"
}
```

### **Category Data Example**

```json
{
    "Id": "1",
    "Name": "test1",
}
```

## ©License
![GitHub](https://img.shields.io/github/license/Ctere1/golang-api?style=flat-square)


## 📌Contributors

<a href="https://github.com/Ctere1/">
  <img src="https://contrib.rocks/image?repo=Ctere1/Ctere1" />
</a>

