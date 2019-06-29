#Use golang version 1.12.5 or above only
FROM golang:1.12.5 AS builder

# set path environtment variable
ENV PROJECT_PATH=/github.com/THEMEGANAME/workshop-golang

# work directory for application
WORKDIR $PROJECT_PATH

# copy our app to docker container
# copy config & assets to container
# tl;dr config & assets is not golang file(*.go) the compiler cant compile another extention like *.xlsx, *.ttf, *.json
COPY . .
# COPY ./user.db .


# download dependency and build binary file

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

# After build our application
# After build finish we want it small so we use another image to host our application
# we use Scratch(any version)
FROM scratch
COPY --from=builder ./server ./
# COPY --from=builder user.db ./server


# the specific port that you need to use
EXPOSE 7000

# list of env variable
# dev = developement
# uar = user acceptance testing
# prod = production
 ENTRYPOINT ["./server"]

# CMD [ "ls" ]