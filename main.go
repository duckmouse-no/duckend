package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func main() {

  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  stripeKey := os.Getenv("STRIPE_KEY")
  stripe.Key = stripeKey

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
  })

  http.HandleFunc("/create-checkout-session", createCheckoutSession)

  fmt.Println("Running on port 80!")

  log.Fatal(http.ListenAndServe(":80", nil))

}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
  priceId := os.Getenv("PRICE_ID")
  domain := "http://46.101.120.64"
  params := &stripe.CheckoutSessionParams{
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        Price: stripe.String(priceId),
        Quantity: stripe.Int64(1),
      },
    },
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    SuccessURL: stripe.String(domain + "?success=true"),
    CancelURL: stripe.String(domain + "?canceled=true"),
    AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
  }

  s, err := session.New(params)

  if err != nil {
    log.Printf("session.New: %v", err)
  }

  http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
