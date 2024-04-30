FROM golang:1.22.0

WORKDIR /app
COPY . .

RUN go mod download

EXPOSE 7000
CMD ["go", "run", "."]

LABEL "version"="1.0.0"
LABEL "project_name"="DevHub"
LABEL "Authors"="fatima abbas, mohammed alsammak, sameer ahmed, Abdulrahman Idrees"