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

    // --- MAGIC CODE: UI + AUTO ADMIN (Updated for v0.23) ---
    app.OnServe().BindFunc(func(e *core.ServeEvent) error {
        // 1. DarkBase UI (Static Files) show karo
        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

        // 2. Admin Account Check karo
        email := "anshkumarchan@gmail.com"
        
        // Agar Admin nahi milta, to naya banao
        _, err := app.FindAdminByEmail(email)
        if err != nil {
            admin := &core.Admin{}
            admin.Email = email
            admin.SetPassword("1234567890")
            
            if saveErr := app.Save(admin); saveErr != nil {
                log.Println("⚠️ Error creating admin:", saveErr)
            } else {
                log.Println("✅ SUCCESS: Admin Created! Login: anshkumarchan@gmail.com / 1234567890")
            }
        }
        
        return e.Next()
    })
    // -------------------------------------------------------

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
