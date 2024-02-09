package cmd

import (
	//"fmt"
	x "mywabot/system"
	//"os"
	//"os/exec"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name: "(s|stiker)",
		Cmd:  []string{"sticker"},
		Tags: "convert",
		IsMedia: true,
		Prefix: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			// quoted sticker
			if m.IsQuotedSticker {
				//conjp := "./tmp/" + m.ID + ".webp"
				//conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        
        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
        /*
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved webp")
					return
				}
				createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", conwp, conwp)
				cmd := exec.Command("bash", "-c", createExif)
				err = cmd.Run()
				if err != nil {
					fmt.Println("Failed to set webp metadata", err)
				}
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
				m.React("✅")
			}

			// quoted image
			if m.IsQuotedImage {
				conjp := "./tmp/" + m.ID + ".jpg"
				conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved jpg")
					return
				}
				x.ImgToWebp(conjp, conwp)
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
        */
				m.React("✅")
			}

			// quoted video
			if m.IsQuotedVideo {
				//conjp := "./tmp/" + m.ID + ".mp4"
				//conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.VideoMessage)

        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)

        /*
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved mp4")
					return
				}
				x.VideoToWebp(conjp, conwp)
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
				m.React("✅")
			}

			// from video
			if m.IsVideo {
				conjp := "./tmp/" + m.ID + ".mp4"
				conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Media)
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved mp4")
					return
				}
				x.VideoToWebp(conjp, conwp)
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
        */
				m.React("✅")
			}

			// from image
			if m.IsImage {
				//conjp := "./tmp/" + m.ID + ".jpg"
				//conwp := "./tmp/" + m.ID + ".webp"
				byte, _ := sock.WA.Download(m.Media)
        
        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
        /*
				err := os.WriteFile(conjp, byte, 0644)
				if err != nil {
					fmt.Println("Failed saved jpg")
					return
				}
				x.ImgToWebp(conjp, conwp)
				sock.StickerPath(m.From, conwp, *m)
				os.Remove(conwp)
				os.Remove(conjp)
        */
				m.React("✅")
			}

		},
	})
}
