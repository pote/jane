package main

import(
  "bufio"
  "flag"
  "log"
  "os"

  "github.com/pote/philote-go"
)

func main() {
  secret := flag.String("s", "", "Philote server's JWT secret used for authentication.")
  server := flag.String("p", "ws://localhost:6380", "Philote server")
  channel := flag.String("c", "test", "channel you want to connect to")
  token := flag.String("t", "", "Auth token, if you want to use a specific one")

  flag.Parse()

  auth := *token; if auth == "" {
    var err error
    auth, err = philote.NewToken(*secret, []string{*channel}, []string{*channel}); if err != nil {
      log.Fatal(err)
    }
  }

  client, err := philote.NewClient(*server, auth); if err != nil {
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
