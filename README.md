![Hacker News Journal Logo](./logo.png)

# Hacker News Journal

This is a project whose purpose is to store hackernews top articles and be able to view them randomly. This comes from the fact that some articles are timeless and can be stored in a fashion like a web archive.

This project divided into three programs

* Cronjob: This is the scrapper that gets the articles from hackernews into the mongodb database.
* API: This is the way we extern the articles stored to consume with any type of client.
* Client: In this case is a mobile app made with flutter that displays articles given randomly by the API.

## Requirements

* docker-compose 1.2+
* go 1.14+
* flutter 1.20+

## Build project

To have all the programs we must get up our docker container with the database. If you only want to consume the API with the application, just go directly to the client build.

#### Pre build

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

> Only tested in the iPhone SE 

You should have connected your phone to the computer in order to install it

```bash
cd hnclient
flutter create .
flutter pub get
flutter pub run flutter_launcher_icons:main
flutter run
```

## License

This project is under the MIT license
