env GOOS=linux GOARCH=amd64 go build -o my-movie-list
docker build -t my-movie-list .
docker tag my-movie-list eatsfulllemons/my-movie-list
docker push eatsfulllemons/my-movie-list