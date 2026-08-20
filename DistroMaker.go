package main

import (
   "fmt"
   "math/rand"
)

func main() {
    distros := []string{
        "Void",
	"CRUX",
	"Alpine",
	"Gentoo",
	"Plan 9",
	"OpenBSD",
	"Source",
	"Open",

     }

     animals := []string{
          "hamster",
	  "Crab",
	  "Penguin",
	  "Tux",
	  "Tuz",
	  "Pigeon",
	  "Seagull",
	  "Seal",
	  "Frog",
	  "Pneumonoultramicroscopicsilicovolcanoconiosis",
  
      }

      des := []string{
          "None",
	  "LXDE",
	  "XFCE",
	  "MATE",
	  "Trinity",
 	  
      }	
	
      wms := []string{
          "None",
	  "dwm",
	  "i3",
	  "OpenBox",
	  "Fluxbox",
	  "Window Maker",
	  "IceWM",
	  
      }
      
      inits := []string{
          "runit",
	  "OpenRC",
	  "s6",
	  "dinit",

      }

      kernels := []string{
              "Linux",
	      "Linux Libre",

      }

      gnus := []string{
          "GNU free",
          "GNU Unfree",

      }  

      distro := distros[rand.Intn(len(distros))]
      animal := animals[rand.Intn(len(animals))]
      de := des[rand.Intn(len(des))]
      wm := wms[rand.Intn(len(des))]
      init := inits[rand.Intn(len(inits))]
      kernel := kernels[rand.Intn(len(kernels))]
      gnu := gnus[rand.Intn(len(gnus))]
      fmt.Println("Your New Distro Is..", distro, animal, "Linux")
      fmt.Println("Desktop Environment:", de)
      fmt.Println("Window Manager:", wm)
      fmt.Println("Init System:", init)
      fmt.Println("Kernel:", kernel)
      fmt.Println("GNU:", gnu)

}