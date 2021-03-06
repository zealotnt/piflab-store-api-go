package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	// . "github.com/o0khoiclub0o/piflab-store-api-go/models/form"
	// . "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

	"net/http"
)

func GetCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func UpdateCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func UpdateCartItemHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func DeleteCartItemHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var cart Order

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, &cart)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, cart)
	}
}

func CheckoutCartHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.CART_SERVICE, nil)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		WriteBody(w, body)
	}
}

func GetCheckoutHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var order_page OrderPage

		// Forward it to carts service
		resp, body, err := RequestForwarder(r, app.ORDER_SERVICE, &order_page)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, order_page)
	}
}

func GetCheckoutDetailHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var checkout_detail CheckoutReturn

		// Forward it to orders service
		resp, body, err := RequestForwarder(r, app.ORDER_SERVICE, &checkout_detail)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, checkout_detail)
	}
}

func UpdateCheckoutStatusHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var checkout_detail CheckoutReturn

		// Forward it to orders service
		resp, body, err := RequestForwarder(r, app.ORDER_SERVICE, &checkout_detail)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, checkout_detail)
	}
}
