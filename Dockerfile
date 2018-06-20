FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ADD my-movie-list /

WORKDIR /

ENV MY_MOVIE_LIST_PORT $MY_MOVIE_LIST_PORT
ENV TMDB_API_KEY $TMDB_API_KEY

EXPOSE $MY_MOVIE_LIST_PORT

CMD [ "./my-movie-list" ]