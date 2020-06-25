![Hacker News Journal Logo](./logo.png)

# Hacker News Journal

This a simple program that scrapes the hackernews top articles and store the titles with the url in a database. It was made to run as a cronjob that is going to insert into a mongodb database.

## Requirements

```bash
docker-compose 1.2+
go 1.14+
```

## Local build

```bash
git clone https://github.com/manoloesparta/hnjournal && cd hnjournal
docker-compose up -d
go run main.go
```

If you want it as a cronjob

```bash
git clone https://github.com/manoloesparta/hnjournal && cd hnjournal
docker-compose up -d
go build
crontab -e # add "0 * * * * /path/to/hnjournal"
```

## License

This project is under the MIT license
