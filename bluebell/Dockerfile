
FROM debian:stretch-slim


COPY ./templates /templates
COPY ./static /static
COPY ./setting /setting
COPY ./bin/bluebell /
# 从builder镜像中把/dist/app 拷贝到当前目录


# 声明服务端口
EXPOSE 9084

# 需要运行的命令
#ENTRYPOINT ["/bluebell", "conf/config.yaml"]
CMD ["./bluebell"]