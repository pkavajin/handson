FROM alpine

COPY ./handson /handson

ENTRYPOINT [ "/handson" ]
