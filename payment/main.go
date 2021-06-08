// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "Hello Payment!")
// }

// func main() {
//     http.HandleFunc("/", handler)
//     fmt.Println("Server started at 8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/checkout/session"
)


func main() {
  stripe.Key = "sk_test_51IzocuKQk9KU3Mqk49siloOnnwJaE4tkAvxfIlFm0v8ZmhddI3VWyg2YoXmj65NeIHfQQPhjplPN7byLuGIhEc1O00QfLBzq7w"

  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.POST("/create-checkout-session", createCheckoutSession)

  e.Logger.Fatal(e.Start("http://localhost:4242"))
}

type CreateCheckoutSessionResponse struct {
  SessionID string `json:"id"`
}

func createCheckoutSession(c echo.Context) (err error) {
  params := &stripe.CheckoutSessionParams{
    PaymentMethodTypes: stripe.StringSlice([]string{
      "card",
    }),
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
          Currency: stripe.String("usd"),
          ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
            Name: stripe.String("Rent"),
          },
          UnitAmount: stripe.Int64(3000),
        },
        Quantity: stripe.Int64(1),
      },
    },
    SuccessURL: stripe.String("https://example.com/success"),
    CancelURL:  stripe.String("https://example.com/cancel"),
  }

  session, _ := session.New(params)

  if err != nil {
    return err
  }

  data := CreateCheckoutSessionResponse{
    SessionID: session.ID,
  }

  return c.JSON(http.StatusOK, data)
}