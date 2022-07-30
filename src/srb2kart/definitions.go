package srb2kart

import (
  "GoBun/srb2kart/addons"
)

type srb2kart struct {
  Name string
  Port int
  Volumes []Volume
  Addons addons.AddonCollection
};

type Volume string;
