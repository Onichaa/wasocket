package cmd

import (
  x "mywabot/system"

  "net/http"
  "encoding/json"
  "net/url"
  "fmt"
  "strings"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(ig|instagram|igdown|igdl)",
    Cmd:    []string{"instagram"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    if !strings.Contains(m.Query, "instagram") {
        m.Reply("Itu bukan link instagram")
      return
    }

      resp, err := http.Get("https://skizo.tech/api/igdl?url="+url.QueryEscape(m.Query)+"&apikey=batu")

      if strings.Contains(m.Query, "https://www.instagram.com/reel/") {
      type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, url := range media {  
          sock.SendVideo(m.From, url, caption, *m)
        }

       } else if strings.Contains(m.Query, "https://www.instagram.com/p/") {
        type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          m.Reply(err.Error())
          return
        }
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, ur := range media {

          resp, err := http.Get(ur)
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

            defer resp.Body.Close()
          
            mime := resp.Header.Get("Content-Type")

            if mime == "" {
              m.Reply("No Content-Type")
            }

            if strings.Contains(mime, "video") {
              sock.SendVideo(m.From, ur, caption, *m)
            } else {
               sock.SendImage(m.From, ur, caption, *m)
            }
        }   

      } else if strings.Contains(m.Query, "https://www.instagram.com/stories/") {
        type respon struct {
          Caption string   `json:"caption"`
          Media   []string `json:"media"`
        }
        if err != nil {
            fmt.Println("Error:", err)
            return
          }
          defer resp.Body.Close()
          var data respon
          err = json.NewDecoder(resp.Body).Decode(&data)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          // Mengambil media
            caption := data.Caption
          media := data.Media
          for _, ur := range media {

            resp, err := http.Get(ur)
              if err != nil {
                fmt.Println("Error:", err)
                return
              }

                defer resp.Body.Close()

                mime := resp.Header.Get("Content-Type")

                if mime == "" {
                  m.Reply("No Content-Type")
                }

                if strings.Contains(mime, "video") {
                  sock.SendVideo(m.From, ur, caption, *m)
                } else {
                   sock.SendImage(m.From, ur, caption, *m)
                }
            }   
          }   

      m.React("✅")
    },
  })
}
