package addons

import (
  "GoBun/srb2kart/addons"
)

func contentToCollection(in content) (addons.AddonCollection) {
  if in.Kind == "file" {
    return addons.Addon{File: in.Value}
  } else if in.Kind == "group" {
    return addons.AddonGroup{GroupName: in.Value}
  }

  return addons.Addon{}
}
