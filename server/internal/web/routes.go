package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(router fiber.Router) {
	// router.Use(adaptor.HTTPMiddleware(func(h http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// Allow any origin to access the resource.
	// 		w.Header().Set("Access-Control-Allow-Origin", "*")
	// 		// Allow these methods.
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//
	// 		w.Header().
	// 			Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
	// 		// Expose specific headers to the client.
	// 		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	//
	// 		// If it's a preflight OPTIONS request, respond immediately.
	// 		if r.Method == http.MethodOptions {
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}
	//
	// 		h.ServeHTTP(w, r)
	// 	})
	// }))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                                                  // Allow all origins (or specify your frontend URL)
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",                        // Allowed methods
		AllowHeaders:     "Content-Type,Authorization,Accept,X-Requested-With", // Allowed headers
		ExposeHeaders:    "Authorization",                                      // Exposed headers
		AllowCredentials: true,                                                 // Allow credentials (if needed)
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Api!")
	})

	router.Post("/direction", postDirection)
	router.Post("/emoji", postEmoji)

	router.Get("/data", getData)
}
