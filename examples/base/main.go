package main

import (
    "log"
    "net/http"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
    "github.com/labstack/echo/v5"
)

// Aapka HTML Code yahan string ke andar aayega (Backticks ` ` ke beech mein)
const myHtml = `
<!DOCTYPE html>
<html>
<head>
    <title>DarkBase</title>
    <h1 style="color:white; text-align:center;">DarkBase Loaded Successfully!</h1>
    <p style="color:gray; text-align:center;">Agar ye dikh raha hai, to server sahi hai.</p>
    </head>
<body style="background-color:black;"></body>
</html>
`

func main() {
    app := pocketbase.New()

    app.OnServe().BindFunc(func(e *core.ServeEvent) error {
        
        // Superuser (Admin) Check
        superusers, _ := app.FindCollectionByNameOrId("_superusers")
        if superusers != nil {
            email := "anshkumarchan@gmail.com"
            _, err := app.FindAuthRecordByEmail("_superusers", email)
            if err != nil {
                record := core.NewRecord(superusers)
                record.SetEmail(email)
                record.SetPassword("1234567890")
                app.Save(record)
                log.Println("✅ Admin Created: anshkumarchan@gmail.com")
            }
        }

        // --- DIRECT HTML SERVING (Folder ki zarurat nahi) ---
        e.Router.GET("/*", func(c echo.Context) error {
            return c.HTML(http.StatusOK, myHtml)
        })
        // ---------------------------------------------------
        
        return e.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
