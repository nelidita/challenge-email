FROM busybox

WORKDIR "/emails-indexer"
COPY ./zinc .

ARG ZS_BASE_API_URL=http://localhost:4080/
ARG ZS_INDEX=emails
ARG ZS_USER=admin
ARG ZS_PASSWORD=Complexpass#123

ENV ZS_BASE_API_URL=${ZS_BASE_API_URL} \
    ZS_INDEX=${ZS_INDEX} \
    ZS_USER=${ZS_USER} \
    ZS_PASSWORD=${ZS_PASSWORD}

EXPOSE 4080

CMD ["./zinc"]