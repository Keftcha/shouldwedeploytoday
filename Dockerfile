FROM alpine

WORKDIR /bin/shouldwedeploytoday

COPY ./images/ ./images/
COPY ./out/server .
COPY ./static/ ./static/

EXPOSE 5461
CMD ["./server"]
