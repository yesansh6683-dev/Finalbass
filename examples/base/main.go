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

    // --- NEW MAGIC CODE (v0.24 FIX) ---
    app.OnServe().BindFunc(func(e *core.ServeEvent) error {
        
        // 1. Static Files (UI) Serve karne ka naya tareeka
        e.Router.GET("/*", apis.Static(os.DirFS("./pb_public"), false))

        // 2. Auto-Superuser Creator (Admin ki jagah ab Superuser hai)
        superusers, err := app.FindCollectionByNameOrId("_superusers")
        if err == nil {
            email := "anshkumarchan@gmail.com"
            
            // Check agar user pehle se hai
            _, err := app.FindAuthRecordByEmail("_superusers", email)
            if err != nil {
                // Agar nahi hai, to naya banao
                record := core.NewRecord(superusers)
                record.SetEmail(email)
                record.SetPassword("1234567890")
                
                if err := app.Save(record); err != nil {
                    log.Println("⚠️ Error creating superuser:", err)
                } else {
                    log.Println("✅ SUCCESS: Superuser Created! Login: anshkumarchan@gmail.com / 1234567890")
                }
            }
        }
        
        return e.Next()
    })
    // ----------------------------------

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
