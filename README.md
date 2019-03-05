# Instagraph
> Social graph network of your instagram account.

<p align="center"><img width=100% src="https://github.com/ahmdrz/instagraph/blob/master/resources/screenshot.png"></p>

### Docker

```
$ docker pull ahmdrz/instagraph:latest
$ docker run -e INSTA_USERNAME="" -e INSTA_PASSWORD="" -p 8080:8080 ahmdrz/instagraph:latest
```

### Build

```
$ go build -i -o instagraph
$ ./instagraph -username="" -password=""
```

---

Powered with :heart: by `sigma.js` and `ahmdrz/goinsta`.
