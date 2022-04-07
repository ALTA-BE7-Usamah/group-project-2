## About The Project
This is a Backend Project about Ecommerce App using Golang programming language.
There's a jwt token for login authentication. and if user already have an account, user can just login with email and password. if not, user should register first.

User can edit or delete their details and create their own product.
User can buy products and edit the quantity on the cart section.
In the Order section, user that already have a product list to buy on the shopping cart can just checkout the cart by inputing the shipping address and the credit card details.


### Built With
* [Gorm](https://gorm.io/)
* [Echo](https://echo.labstack.com/)
* [Docker](https://www.docker.com/)
* [Database Stored in RDS Cloud by Amazon Web Services](https://aws.amazon.com/id/?nc2=h_lg)

### EndPoint
* `/users` with method `POST` to Create account/Register
* `/auth` with method `POST` to Login to the system
* `/users/:id` with method `GET` to Get one user details
* `/users/:id` with method `PUT` to edit user details
* `/users/:id` with method `DELETE` to delete account
* `/products` with method `GET` to See all the products
* `/products` with method `POST` to create product
* `/products/:id` with method `GET` to get one product by id
* `/products/:id` with method `PUT` to edit one product details by id
* `/products/:id` with method `DELETE` to delete product
* `/cart` with method `POST` to create a cart on the shopping cart
* `/cart` with method `GET` to Read all the user's cart created by user that have logged in
* `/cart/:id` with method `PUT` to edit the cart, either the product or quantity that user want to edit
* `/cart/:id` with method `DELETE` to delete a cart
* `/order` with method `POST` to create an order by inputing carts, shipping address, and credit card details
* `/order` with method `GET` to read all the order history

