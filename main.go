package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ahmdrz/instagraph/src/graph"
	"github.com/ahmdrz/instagraph/src/instagram"
)

func main() {
	var (
		username   string
		password   string
		delay      = 1
		limit      = 300
		usersLimit = 300
		listenAddr = "0.0.0.0:8080"
		g          = graph.New()
		showLast   = false
	)
	username = os.Getenv("INSTA_USERNAME")
	password = os.Getenv("INSTA_PASSWORD")
	if username == "" && password == "" {
		flag.StringVar(&username, "username", "", "Instagram username")
		flag.StringVar(&password, "password", "", "Instagram password")
		flag.IntVar(&limit, "limit", 300, "How many users should be scan in firsth depth of your followings")
		flag.IntVar(&usersLimit, "users-limit", 300, "Max users in each followings to scan")
		flag.IntVar(&delay, "delay", 1, "Sleep between each following")
		flag.BoolVar(&showLast, "latest", false, "Use the latest genereted json file.")
		flag.Parse()
	}

	if !showLast {
		var instance *instagram.Instagram
		if fileExists(username + ".json") {
			var err error
			log.Printf("Loading instagram as %s ...", username)
			instance, err = instagram.Import(username + ".json")
			if err != nil {
				log.Fatal(err)
				return
			}
		} else {
			var err error
			log.Printf("Connecting to instagram as %s ...", username)
			instance, err = instagram.New(username, password)
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Printf("Connected !")

			instance.Export(username + ".json")
		}

		log.Printf("Fetching followings ...")
		followings := instance.Followings()
		shuffle(followings)

		if limit == -1 {
			limit = len(followings)
		}

		for i, user := range followings {
			if i >= limit {
				log.Println("Reached to limit.")
				break
			}

			log.Printf("Scaning (%04d/%04d) user %s ...", i, limit, user.Username)

			users := user.Followings(instance)
			if len(users) > usersLimit {
				users = users[:usersLimit]
			}
			shuffle(users)

			for _, target := range users {
				g.AddConnection(user.Username, target.Username)
			}

			time.Sleep(time.Duration(delay) * time.Second)
		}

		ioutil.WriteFile("static/data.json", g.Marshall(), 0755)
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	})
	handler.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Listening to %s ...", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, handler))
}

func shuffle(vals []instagram.User) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func getPWD() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
