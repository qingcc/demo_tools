FROM scratch

#RUN go build http.go

WORKDIR /app

EXPOSE 10011
#所有的文件复制均使用COPY,仅在需要自动解压缩时使用ADD
COPY http /app/

CMD ["/app/http"]
