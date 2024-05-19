# Currency Rate Notifier

This project is a currency rate notifier that allows users to subscribe and receive daily updates on the exchange rate between USD and UAH.

## Getting Started

To get started with this project, you'll need to set up an `.env` file with your configuration and sensitive data. This file will be used by Docker to set environment variables when running the Dockerfile.

### Prerequisites

- Docker
- Docker Compose

### Setting Up the .env File

Create a `.env` file in the root directory of the project with the following content:

```sh
MYSQL_ROOT_PASSWORD=your_root_password
MYSQL_DATABASE=your_database_name
MYSQL_USER=your_username
MYSQL_PASSWORD=your_password
MYSQL_HOST=db

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
GMAIL_SMTP_PASSWORD=your_generated_app_password
GMAIL_SMTP_EMAIL=your_gmail_account@gmail.com

EXCHANGERATESAPI_KEY=your_exchangeratesapi_key
EXCHANGERATESAPI_BASE_URL=http://api.exchangeratesapi.io/v1/latest
```

Replace the placeholder values with your actual data.

### Exchange Rates API

We use the ExchangeRatesAPI which offers a free plan with 250 requests per month. The free plan provides EUR to various currencies rates. To calculate USD to UAH rates, we fetch EUR, USD, and UAH rates and then compute the approximate rate for USD to UAH.

### Email Notifications

Email notifications are sent using the Gmail SMTP server. To secure your account and not use your real password, especially if you have 2-step verification enabled, generate an app-specific password:

Go to your Google Account settings.
Search for “App passwords” and follow the link.
Create a password for this application and use it in the `.env` file.

### MySQL Database Credentials

Provide the credentials for your MySQL database in the `.env` file as shown above. Choose the credentials you wish to use for your database.

### Running the Project

With the `.env` file in place, you can start the project using Docker Compose:

```bash
docker-compose up --build
```

This command will build the images and start the containers as defined in your docker-compose.yml file.

## Testing the Project

The project comes with two primary endpoints that can be tested to ensure functionality:

### GET /rate

This endpoint returns the current exchange rate from USD to UAH in a JSON format. The response will look like this:

```json
{
  "rate": 39.25142962569967
}
```

To test this endpoint, you can use tools like curl or Postman. Here’s an example using curl:

```sh
curl -X GET http://localhost:8080/rate
```

### POST /subscribe

This endpoint allows users to subscribe to daily exchange rate updates by sending their email address in a JSON format. The request body should look like this:

```json
{
  "email": "yourEmail@gmail.com"
}
```

Upon successful subscription, the server will return a 200 OK status with the message:
`Subscription successful`

If the email address is already subscribed, the server will return a `409 Conflict` status with the message:

`The email address you have entered is already subscribed.`

To test this endpoint, you can use the following curl command:

```sh
curl -X POST http://localhost:8080/subscribe -H "Content-Type: application/json" -d '{"email":"yourEmail@gmail.com"}'
```

## Receiving Emails

Subscribed users should expect to receive emails with the USD-UAH exchange rate at 8 a.m. UTC+3 daily. Ensure that your SMTP settings are correctly configured in the .env file to receive these emails.

## Unit Testing

Unit tests have been created to ensure the reliability and correctness of the application's core functionalities:

### FetchExchangeRates Function

- **Purpose**: To simulate a third-party API call and verify that the `FetchExchangeRates` function correctly handles the response.
- **Method**: The test mocks the API response, providing predefined data that the function should process.
- **Validation**: It checks if the function correctly interprets the response and handles various scenarios such as successful data retrieval and error conditions.

### Subscription Package

- **Purpose**: To test the database operations related to subscribing users, specifically the insertion of emails and retrieval of all subscribed emails.
- **Method**: Utilizes mock SQL to simulate database interactions, allowing the tests to run without a real database connection.
- **Validation**: Ensures that the `InsertEmail` function can handle new and duplicate email entries appropriately and that the `GetAllEmails` function accurately retrieves the list of subscribers.

These initial tests lay the groundwork for a robust testing suite. Future enhancements will include more comprehensive test coverage, ensuring that all critical logic paths are thoroughly validated for correctness and stability.

## Tools and Libraries

This project utilizes a variety of tools and libraries to achieve its functionality. Below is a brief overview of each and the reasons for their selection:

### MySQL

MySQL is used as the database for this project due to its widespread popularity and reliability as a relational database management system. It's well-supported, scalable, and provides the robustness needed for handling the application's data.

### Goose

For database migrations, we use Goose, a database migration tool written in Go. Goose allows us to apply version control to our database schema changes, making it easier to manage and deploy updates to the database structure.

### Go Libraries

Several Go libraries are employed in this project to facilitate various functionalities:

- **go-sql-driver**: This is a lightweight and fast MySQL driver for Go's (golang) database/sql package. It allows the application to interact with the MySQL database efficiently.

- **godotenv**: A Go port of the Ruby dotenv project, which loads environment variables from a .env file. This library is used to manage configuration and sensitive data without hardcoding them into the source code.

- **robfig/cron**: A cron library for Go that enables the scheduling of jobs to run at specified intervals. It’s used in this project to trigger daily email notifications to subscribed users.

These tools and libraries were chosen for their reliability, ease of use, and strong community support, ensuring that the project remains maintainable and extensible.

## Docker Compose Workflow

The `docker-compose.yml` file defines the workflow for running the services necessary for the Currency Rate Notifier. Here's a brief overview:

1. **Database Initialization**: The MySQL database container is started first, using the environment variables defined in the `.env` file to set up the database credentials.

2. **Application Startup**: Once the database is up and running, the application container is started. It waits for the database to be fully operational before initiating the connection.

3. **Database Migration**: In parallel with the application startup, the migrator service runs the Goose tool to apply any pending database migrations. This ensures that the database schema is always up-to-date.

By orchestrating the services in this manner, Docker Compose ensures a smooth startup sequence and prepares the application environment for use.
