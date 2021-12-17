FROM golang:1.17.5-alpine3.15 

RUN apk update && apk add gcc git && apk add libc-dev && \
    apk --no-cache add libaio libnsl libc6-compat curl && \
    cd /tmp && \
    curl -o instantclient-basiclite.zip https://download.oracle.com/otn_software/linux/instantclient/instantclient-basiclite-linuxx64.zip -SL && \
    unzip instantclient-basiclite.zip && \
    mv instantclient*/ /usr/lib/instantclient && \
    rm instantclient-basiclite.zip && \
    ln -s /usr/lib/instantclient/libclntsh.so.19.1 /usr/lib/libclntsh.so && \
    ln -s /usr/lib/instantclient/libocci.so.19.1 /usr/lib/libocci.so && \
    ln -s /usr/lib/instantclient/libociicus.so /usr/lib/libociicus.so && \
    ln -s /usr/lib/instantclient/libnnz19.so /usr/lib/libnnz19.so && \
    ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
    ln -s /lib/libc.so.6 /usr/lib/libresolv.so.2 && \
    ln -s /lib64/ld-linux-x86-64.so.2 /usr/lib/ld-linux-x86-64.so.2

ENV ORACLE_BASE=/usr/lib/instantclient
ENV LD_LIBRARY_PATH=/usr/lib/instantclient
ENV TNS_ADMIN=/usr/lib/instantclient
ENV ORACLE_HOME=/usr/lib/instantclient    

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o main .
EXPOSE 8881
CMD [ "/app/main" ]

# docker run --restart=always  -d --name survey -p 8881:8881 survey-reports-api:latest
# docker run --name survey-reports-api-redis -p 6379:6379 redis:latest redis-server --requirepass "admin-survey"

