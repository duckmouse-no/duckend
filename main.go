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
  domain := "https://localhost:8000"
  params := &stripe.CheckoutSessionParams{
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        Price: stripe.String(priceId),
        AdjustableQuantity: &stripe.CheckoutSessionLineItemAdjustableQuantityParams{
          Enabled: stripe.Bool(true),
          Minimum: stripe.Int64(1),
          Maximum: stripe.Int64(50),
        },
        Quantity: stripe.Int64(1),
      },
    },
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    SuccessURL: stripe.String("https://duckmouse.no" + "/takk"),
    CancelURL: stripe.String("https://duckmouse.no"),
    AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(false)},
      PhoneNumberCollection: &stripe.CheckoutSessionPhoneNumberCollectionParams{
      Enabled: stripe.Bool(true),
    },
    ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
      AllowedCountries: stripe.StringSlice([]string{
        "NO",
      }),
    },
    ShippingOptions: []*stripe.CheckoutSessionShippingOptionParams{
      &stripe.CheckoutSessionShippingOptionParams{
        ShippingRate: stripe.String("shr_1Lk4hHC8GFiqryiro7dDaiv8"),
      },
      &stripe.CheckoutSessionShippingOptionParams{
        ShippingRate: stripe.String("shr_1Lp5OSC8GFiqryirjRWIYOBQ"),
      },
    },
  }

  s, err := session.New(params)

  if err != nil {
    log.Printf("session.New: %v", err)
  }

  http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
