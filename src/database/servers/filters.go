package servers

import (
  "GoBun/srb2kart"
)

type keyFilter struct {
  Name string `bson::"name"`
}

func makeKeyFilter(server srb2kart.Srb2kart) keyFilter {
  return keyFilter{server.Name}
}
