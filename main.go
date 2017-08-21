package main

import(
  "bufio"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"

  "github.com/pote/philote-go"
  "github.com/dgrijalva/jwt-go"
)

func main() {
  secret := flag.String("s", "", "Philote server's JWT secret used for authentication.")
  server := flag.String("p", "localhost:6380", "Philote server")
  channel := flag.String("c", "test", "channel you want to connect to")
  token := flag.String("t", "", "Auth token, if you want to use a specific one")
  info := flag.Bool("info", false, "If passed, queries Philote for running info and exits")

  flag.Parse()

  auth := *token; if auth == "" {
    var err error
    auth, err = philote.NewToken(*secret, []string{*channel}, []string{*channel}); if err != nil {
      log.Fatal(err)
    }
  }

  if *info {
    request, err := http.NewRequest("GET", "http://" + *server + "/api/info", nil); if err != nil {
      log.Fatal(err)
    }

    auth := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"api": true})
    token, err := auth.SignedString([]byte(*secret)); if err != nil {
    log.Fatal(err)
    }

    request.Header.Add("Authorization", "Bearer " + token)
    res, err := http.DefaultClient.Do(request); if err != nil {
      log.Fatal(err)
    }

    body, err := ioutil.ReadAll(res.Body); if err != nil {
      log.Fatal(err)
    }

    fmt.Print(string(body))
    os.Exit(0)
  }

  client, err := philote.NewClient("ws://" + *server, auth); if err != nil {
    log.Fatal(err)
  }

  log.Printf("Successful connection to %v on the #%v channel\n", *server, *channel)
  log.Printf("Token: %v\n", client.Token)
  messages := client.NewPhilote()

  go func() {
    for {
      m := <- messages
      log.Println(m.Data)
    }
  }()

  reader := bufio.NewReader(os.Stdin)
  for {
    text, err := reader.ReadString('\n'); if err != nil {
      log.Fatal(err)
    }

    err = client.Publish(&philote.Message{Channel: *channel, Data: text}); if err != nil {
      log.Println(err)
    }
  }
}
