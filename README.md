# Kafka Logger with Go

This project is an example of logging IP addresses using Kafka with Go. It consists of four services: `grafana`, `mariadb`, `notifier`, and `processor`.

## Getting Started

To get started, make sure you have Docker installed on your machine. Then, run the following command to start the services:

```bash
docker-compose up -d
```

### Services

The project consists of the following services:

- Grafana: A visualization tool for time series data.
- Mariadb: A relational database management system.
- Notifier: A service that receives IP addresses from Kafka and sends notifications.
- Processor: A service that consumes messages from Kafka and stores them in the database.

The ports are defined in the `.env` file.
