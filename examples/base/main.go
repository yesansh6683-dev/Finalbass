package main

import (
    "log"
    "net/http"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
)

// --- AAPKA HTML CODE YAHAN PASTE KAREIN (Backticks ` ` ke beech mein) ---
const myHtml = `
<!DOCTYPE html>
<html>
<head>
    <title>DarkBase</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>body { background-color: #0f0f0f; color: white; font-family: sans-serif; display: flex; align-items: center; justify-content: center; height: 100vh; margin: 0; }</style>
</head>
<body>
    <div class="text-center">
        <h1 class="text-4xl font-bold text-green-500 mb-4">DarkBase is LIVE! 🚀</h1>
        <p class="text-gray-400">Connected to PocketBase v0.24</p>
        <p class="mt-4 text-sm text-gray-600">Now go to GitHub and replace this HTML with your real Dashboard code.</p>
        <a href="/_/" class="mt-6 inline-block bg-green-600 px-6 py-2 rounded text-white hover:bg-green-700">Go to Admin Panel</a>
    </div>
</body>
</html>
`

func main() {
    app := pocketbase.New()

    app.OnServe().BindFunc(func(e *core.ServeEvent) error {
        
        // 1. Auto-Superuser (Admin) Creator
        superusers, err := app.FindCollectionByNameOrId("_superusers")
        if err == nil {
            email := "anshkumarchan@gmail.com"
            _, err := app.FindAuthRecordByEmail("_superusers", email)
            if err != nil {
                record := core.NewRecord(superusers)
                record.SetEmail(email)
                record.SetPassword("1234567890")
                if err := app.Save(record); err != nil {
                    log.Println("⚠️ Admin Create Error:", err)
                } else {
                    log.Println("✅ Admin Created: anshkumarchan@gmail.com")
                }
            }
        }

        // 2. Direct HTML Serving ( FIXED: "/*" -> "/" )
        // "/" ka matlab hai Root URL (Homepage)
        e.Router.GET("/", func(e *core.RequestEvent) error {
            return e.HTML(http.StatusOK, myHtml)
        })
        
        return e.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
