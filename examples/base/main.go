package main

import (
    "log"
    "os"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/models"
)

func main() {
    app := pocketbase.New()

    // --- MAGIC CODE: AUTOMATIC ADMIN CREATION ---
    app.OnServe().Add(func(e *core.ServeEvent) error {
        // Check if admin already exists
        admin, _ := app.Dao().FindAdminByEmail("anshkumarchan@gmail.com")
        
        // If not, create it automatically
        if admin == nil {
            admin = &models.Admin{}
            admin.Email = "anshkumarchan@gmail.com"
            admin.SetPassword("1234567890") // Password ye rahega
            app.Dao().SaveAdmin(admin)
            log.Println("✅ SYSTEM: Auto-Admin Created Successfully!")
        }
        return nil
    })
    // ---------------------------------------------

    app.OnServe().Add(func(e *core.ServeEvent) error {
        // Serves static files from "pb_public" directory
        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
