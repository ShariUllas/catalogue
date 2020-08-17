# catalogue

## Problem Definition
Build a set of Micro Services (REST APIs with JSON format responses). These REST services will be used to maintain the product catalogue of an ecommerce company.
Following are the entities to be managed through the API:
Category (products are organized into categories for better discovery on ecommerce websites. For example, you will find shampoos in the following category hierarchy: Beauty > Hair car > Hair)
Product 
Variant (variations of the product you select before making a purchase. Ex: a shirt might have variants like size 30, 32 etc… a shampoo might have variants like 300ml, 600 ml etc…)

## Getting Started

### Launch
Below are steps to help launch catelogue

```sh
Please clone the project
git clone git@github.com:ShariUllas/catalogue.git catalogue
cd catalogue/api
```
You can spin up the catalogue environment, including the database and application server by running:

```sh
cp .env.sample .env
make up
```

Application logs can be viewed by running:

```sh
make log
```