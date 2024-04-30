FROM golang:1.22.0

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs

WORKDIR /app
COPY . .

RUN go mod download
RUN npm install

EXPOSE 7000

# Use a shell form to run multiple commands
CMD go run . & npm run dev

LABEL "version"="1.0.0"
LABEL "project_name"="DevHub"
LABEL "Authors"="fatabbas, malsamma, sahmed, akhaled"