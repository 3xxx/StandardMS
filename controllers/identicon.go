package controllers

import (
	"github.com/issue9/identicon"
	"image"
	"image/color/palette"
	"math/rand"
	"time"
)

const (
	RANDOM_AVATAR_SIZE = 800
)

func CreateRandomAvatar(data []byte) image.Image {
	randExtent := len(palette.WebSafe) - 32
	rand.Seed(time.Now().UnixNano())
	colorIndex := rand.Intn(randExtent)
	backColorIndex := colorIndex - 1
	if backColorIndex < 0 {
		backColorIndex = randExtent - 1
	}
	backColor := palette.WebSafe[backColorIndex]
	foreColors := palette.WebSafe[colorIndex : colorIndex+32]
	imgMaker, _ := identicon.New(RANDOM_AVATAR_SIZE, backColor, foreColors...)
	createdImg := imgMaker.Make(data)
	return createdImg
}

// 根据用户访问的IP，为其生成一张头像
// img, _ := identicon.Make(128, color.NRGBA{},color.NRGBA{}, []byte("192.168.1.1"))
// fi, _ := os.Create("/tmp/u1.png")
// png.Encode(fi, img)
// fi.Close()

// 或者
// ii, _ := identicon.New(128, color.NRGBA{}, color.NRGBA{}, color.NRGBA{})
// img := ii.Make([]byte("192.168.1.1"))
// img = ii.Make([]byte("192.168.1.2"))
