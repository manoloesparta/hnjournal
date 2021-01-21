<p align="center">
  <a>
    <img width="200px" src="logo.png" alt="Wikigraph" />
    <h1 align="center">Hacker News Journal</h1>
  </a>
</p>

<p align="center">
  <a><img src="https://img.shields.io/badge/Version-1.0.0-red.svg" alt="Version"></a>
  <a><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="License"></a>
  <a><img src="https://img.shields.io/badge/Made%20with-Go-%2300ACD7" alt="Go"></a>
  <a><img src="https://img.shields.io/badge/Made%20with-JavaScript-yellow" alt="JavaScript"></a>
</p>

This is a project whose purpose is to store hackernews top articles and be able to view them randomly. This comes from the fact that some articles are timeless and can be stored in a fashion like a web archive.

This project divided into three programs

* Cronjob: This is the scrapper that gets the articles from hackernews into the mongodb database.
* API: This is the way we extern the articles stored to consume with any type of client.
* Client: In this case is a mobile app made with flutter that displays articles given randomly by the API.

## Requirements

* docker-compose 1.2+
* go 1.14+
* node 14.5+

## Build project

To have all the programs we must get up our docker container with the database. If you only want to consume the API with the application, just go directly to the client build.

### API

```bash
git clone https://github.com/manoloesparta/hnjournal && cd hnjournal
docker-compose up -d
```

### Cronjob

```bash
cd cronjob
go build
crontab -e 
echo "0 * * * * /path/to/executable" >> /etc/crontab
```

### Client

```bash
cd client
npm install && npm run build
open dist/index.html
```

## License

This project is under the MIT license
