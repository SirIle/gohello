# Gohello, a test service

This simple program is meant to be used to demonstrate scaling using Docker containers. When called it returns a simple JSON message with three fields including a timestamp and hostname so replies coming from different containers can be distinguished from each other.

For building the container [golang-builder](https://github.com/CenturyLinkLabs/golang-builder) is used. The container can be built (assuming you have a working Docker environment) with:

```bash
docker run --rm   -v "$(pwd):/src"   -v /var/run/docker.sock:/var/run/docker.sock   centurylink/golang-builder
```

If you want to tag the image straight after the build the command is

```bash
docker run --rm   -v "$(pwd):/src"   -v /var/run/docker.sock:/var/run/docker.sock   centurylink/golang-builder sirile/gohello
```

After that running the container can be done with

```bash
docker run --rm -p 80:80 gohello
```

Testing the service can be done with

```bash
curl http://<ip_of_docker_node>
```

which should return something like

```json
{"hostname": "76274eeb9b2d","time": "2015-11-25T12:11:07.085817804Z","language": "go"}
```
