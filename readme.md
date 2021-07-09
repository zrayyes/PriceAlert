# Price Alert
Price Alert is a service that helps you keep track of Cryptocurrency prices.

## Installation

Requirements:

- Docker

In the root folder of the project, start docker compose

    docker-compose up --build

Price Alert defaults to port **5001**

Kafka’s UI defaults to port **5002**

Mailhog’s UI defaults to port **8025**

These ports can be managed inside docker-compose.yml

**Please Note: Kafka’s container takes a couple of minutes to start.**
This program was tested on Docker v20.10.7 running on Windows 10 build 19043.1083

----------
## Usage

To create an alert, send a POST request to /alerts with the appropriate information

    {
        "email":"BTCLOVE@gmail.com", // The email address to send the alert to
        "coin":"BTC", // The crypto coin you want to monitor
        "currency": "USD", // The currency you want to convert to
        "price_max": 33000.990000, // The maximum price for the alert
        "price_min": 31000.990000 // The minimum price for the alert
    }

Currently only the following crypto coins are supported

    BTC,ETC,BNB,XRP,DOT,USDT,DOGE,AXS,BUSD,ADA

And only the following currencies

    USD,EUR,GBP,JPY

An overview of all the endpoints and their usage is available under docs/price-alert.html

----------
## Notes
1. I was not able to find a way to share code between my Go modules (server, producer, consumer). Unless Go has a way to accomplish this, I would split the modules into separate Git repositories, and share the code via Git Submodules.
2. I only used an SQLite database because this was a coding assessment; For a production service, I would look into using something more appropriate.
3. Given the time restriction, I was not able to set up proper database call mocking for unit testing and decided instead to use an in-memory database with test data instead.
4. The Kafka consumer does not handle failure properly. If an error happens during the consumption of an event, the information is lost. A possible solution is to re-queue the event on failure and contact an administrator if it fails more than once.
5. I also did not include any tests for the Kafka producer/consumer. This is primarily due to lack of time and my inexperience with the Kafka platform.
6. I currently have the list of coins/currencies hardcoded. In a real service, I would have this inside a database to also allow for validation while creating an alert.
7. I decided to disable a price alert after it is sent to the queue. The same alert can be enabled by sending a PATCH request to alerts/:id and body {"active": true}.
8. A feature I didn’t have time to implement was to define what coins/currencies a producer handled via CLI. This would, for example, allow you to start one producer for USD alerts and another for EUR alerts, to reduce the pressure on the one producer.

## Todo

- ~~Add Dockerfile for Go service~~
- ~~Add Docker Compose file~~
- ~~Add RAML document~~
- ~~Finish REST API~~
- ~~Add Kafka~~
- ~~Create separate producer/consumer modules~~
- ~~Add alerts to Queue~~
- ~~Add Mailhog~~
- ~~Send out emails~~
- ~~Add only active alerts to queue~~
- ~~Deactivate alert after adding to queue~~
- ~~Switch to price range~~
- ~~Add currency tag~~
- ~~Fetch prices from Crypto API~~
- ~~Add tests for business logic~~
- ~~Godoc everything~~
- ~~Retry on kafka fail~~
- Requeue on fail
