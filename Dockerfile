# Source
FROM golang

# Working directory to previously added app directory
WORKDIR /app/

# Copy files to the container
COPY . /app/

# Provide defaults for an executing container
CMD ["go", "run", "./cmd/app/main.go"]