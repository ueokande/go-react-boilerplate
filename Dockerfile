# Stage 1: build frontend
FROM node:16.13.0 as frontend-builder

COPY frontend /tmp/workspace
WORKDIR /tmp/workspace
RUN yarn install --frozen-lockfile && yarn build

# Stage 2: build server
FROM golang:1.17.3 as backend-builder

COPY . /tmp/workspace
WORKDIR /tmp/workspace
RUN CGO_ENABLED=0 go build -o go-react-boilerplate -trimpath .

# Stage 3: final stage to be deployed
FROM scratch

COPY --from=frontend-builder /tmp/workspace/build/ /usr/local/share/go-react-boilerplate/
COPY --from=backend-builder /tmp/workspace/go-react-boilerplate /usr/local/bin/go-react-boilerplate

USER 10000:10000

EXPOSE 8000
ENTRYPOINT ["/usr/local/bin/go-react-boilerplate", "-http", ":8000", "-webroot", "/usr/local/share/go-react-boilerplate"]
