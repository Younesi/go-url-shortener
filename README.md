

#### Run the Application
Here is the steps to run it with `docker-compose`


### Clone 
```bash
$ git clone https://github.com/Younesi/go-url-shortener/tree/master
```
### move to project
```bash
$ cd go-url-shortener
```
### build it up
```bash
docker-compose up --build -d
```
### check if the containers are running
```bash
$ docker ps
```
#### Create short URL API
```bash
$ curl --request POST \
--data '{
    "long_url": "https://www.linkedin.com/in/mahdi-younesi/"
}' \
  http://localhost:3000/api/v1/create-short-url
```
#### Get short URL API
```bash
$ curl --request GET \
  http://localhost:3000/api/v1/Rx8hdYnM
```

#### Stop
```bash
docker-compose down
```
