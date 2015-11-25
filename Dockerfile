FROM scratch
EXPOSE 80
COPY gohello /
ENTRYPOINT ["/gohello"]
