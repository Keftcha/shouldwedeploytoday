FROM alpine

COPY ./images/ ./images/
COPY ./out/server .

EXPOSE 5461
CMD ["./server"]
