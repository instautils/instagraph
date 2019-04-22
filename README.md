# Instagraph
> Social graph network of your Instagram account.

<p align="center"><img width=100% src="https://github.com/ahmdrz/instagraph/raw/master/resources/screenshot.png"></p>

<p align="center"><img width=100% src="https://github.com/ahmdrz/instagraph/raw/master/resources/neo4j.png"></p>


Another branch named as 'sigma' is ready to use. Master branch is using `neo4j` as graph database.

### Installation (recommended)

You have to run neo4j on your system. You can use docker to use neo4j with following code:

```
$ docker run -p 7474:7474 -p 7687:7687 neo4j:3.0
$ # connection string of neo4j will be something like this:
$ # bolt://neo4j:neo4j@neo4j:7687
```

```
$ # If you have Golang on your computer.
$ go get -u github.com/ahmdrz/instagraph
```

### Build

```
$ git clone https://github.com/ahmdrz/instagraph
$ cd instagraph
$ go build -i -o instagraph
$ INSTAGRAPH_USERNAME="" INSTAGRAPH_PASSWORD="" INSTAGRAPH_NEO_ADDR="" ./instagraph
```

##### Warning

Please make sure that the `<username>.json` is in the safe place. It's your login information of your Instagram.

### Docker

```
$ docker pull ahmdrz/instagraph:latest
$ docker run -e INSTAGRAPH_USERNAME="" -e INSTAGRAPH_PASSWORD="" -e INSTAGRAPH_NEO_ADDR="" ahmdrz/instagraph:latest
```

---

<p align="center"><img width=100% src="https://raw.githubusercontent.com/ahmdrz/instagraph/master/resources/screenshot2.png"></p>

---

Powered with :heart: by `neo4j` and `ahmdrz/goinsta`.
