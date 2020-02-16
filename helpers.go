package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//UserDetails lets us quickly access User info
func UserDetails(s *discordgo.Session, userID string) (user *discordgo.User, err error) {
	user, err = s.User(userID)
	if err != nil {
		fmt.Println("error:", err)
	}
	return
}

//GetAvatarURL will return the URL to the user's avatar
func GetAvatarURL() {
	return
}

//GetAvatarImage will return a type image.Image of the user's avatar
func GetAvatarImage(s *discordgo.Session, userID string) (img image.Image, err error) {
	user, err := s.User(userID)
	if err != nil {
		fmt.Println("error:", err)
	}

	return s.UserAvatarDecode(user)
}

//SaveImage ...
func SaveImage(img image.Image, pname, fname string) (err error) {
	fpath := path.Join(pname, fname)

	f, err := os.Create(fpath)
	if err != nil {
		fmt.Println("failed to create path:", err)
		return
	}
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println("failed to encode as jpeg:", err)
		return
	}
	f.Close()
	return nil
}

//StrArrayToInt will iterate over an array of type string, and convert to type int
// this could probably be replaced with smarter code wherever im using this
func StrArrayToInt(a []string) (ia []int, err error) {
	for _, v := range a {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error converting string array to int:", err)
		}
		ia = append(ia, i)
	}

	return ia, nil
}
