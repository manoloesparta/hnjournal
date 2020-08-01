![Hacker News Journal Logo](./logo.png)

# Hacker News Journal

This is a project whose purpose is to store hackernews top articles and be able to view them randomly. This comes from the fact that some articles are timeless and can be stored in a fashion like a web archive.

This is divided into three programs

* Cronjob: This is the scrapper that gets the articles from hackernews into the mongodb database.
* API: This is the way we extern the articles stored to consume with any type of client.
* Client: In this case is a mobile app made with flutter that displays articles given randomly by the API.

## Requirements

* docker-compose 1.2+
* go 1.14+
* flutter 1.20+

## Build project

To have all the programs we must get up our docker-compose.yml. This was written in a fashion where it is installed on a server except for the client build.

#### General setup

```bash
git clone https://github.com/manoloesparta/hnjournal && cd hnjournal
docker-compose up -d
```

#### Cronjob

```bash
cd cronjob
go build
crontab -e 
# add "0 * * * * /path/to/cronjob/executable"
```

#### API

```bash
screen
cd api
go run main.go 
# ctrl + a followed by ctrl + d
```

#### Client

> This is only tested in the iPhone SE 

```bash
cd hnclient
flutter create .
flutter pub get
flutter pub run flutter_launcher_icons:main
flutter run
```

## License

This project is under the MIT license
