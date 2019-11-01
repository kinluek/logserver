FROM golang:latest

# set the working directory to be root folder
WORKDIR /

# copy everything from local project to the working directory
COPY . /

# build the binary and output the name of the binary as main
RUN go build -o ./main .

# when container starts, execute the binary, this will start the sever.
ENTRYPOINT [ "./main" ]
