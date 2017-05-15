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
  port := flag.String("p", "ws://localhost:6380", "Philote server")
  channel := flag.String("c", "test", "channel you want to connect to")

  flag.Parse()

  client, err := philote.NewClient(*port, *secret, []string{*channel}, []string{*channel}); if err != nil {
    log.Fatal(err)
  }

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
