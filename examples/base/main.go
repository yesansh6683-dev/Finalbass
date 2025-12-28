package main

import (
    "log"
    "os"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    app.OnServe().BindFunc(func(e *core.ServeEvent) error {
        
        // --- YE HAI FIX: Folder Mode On ---
        // Ye line ab aapke 'pb_public' folder ke andar 'index.html' ko dhundegi
        e.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

        // --- Auto-Admin Code (Login ke liye) ---
        superusers, _ := app.FindCollectionByNameOrId("_superusers")
        if superusers != nil {
            email := "anshkumarchan@gmail.com"
            _, err := app.FindAuthRecordByEmail("_superusers", email)
            if err != nil {
                record := core.NewRecord(superusers)
                record.SetEmail(email)
                record.SetPassword("1234567890")
                app.Save(record)
            }
        }
        
        return e.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
