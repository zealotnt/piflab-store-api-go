package handlers

import (
	. "github.com/o0khoiclub0o/piflab-store-api-go/lib"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models"
	. "github.com/o0khoiclub0o/piflab-store-api-go/models/repository"

	"net/http"
)

func GetProductsDetailHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// Forward it to service
		resp, body, err := RequestForwarder(r, app.PRODUCT_SERVICE, nil)
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

func GetProductsHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		// Forward it to service
		resp, body, err := RequestForwarder(r, app.PRODUCT_SERVICE, nil)
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

func CreateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var product Product

		// Forward it to service
		resp, body, err := RequestForwarder(r, app.PRODUCT_SERVICE, &product)
		if err != nil {
			if resp == nil {
				JSON(w, err)
				return
			}
			JSON(w, err, resp.StatusCode)
			return
		}
		if resp.Status != "201 Created" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, product, 201)
	}
}

func UpdateProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		var product Product

		// Forward it to service
		resp, body, err := RequestForwarder(r, app.PRODUCT_SERVICE, &product)
		if resp.Status != "200 OK" {
			JSON(w, ParseError(body), resp.StatusCode)
			return
		}
		if err != nil {
			JSON(w, err, resp.StatusCode)
			return
		}

		// Temporary not support field selection
		JSON(w, product, 200)
	}
}

func DeleteProductHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		product, err := (ProductRepository{app}).DeleteProduct(c.ID())
		if err != nil {
			JSON(w, err, 500)
			return
		}
		JSON(w, product)
	}
}

func OptionHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		w.Header().Add("Access-Control-Allow-Origin", `*`)
		w.Header().Add("Access-Control-Allow-Methods", `GET, POST, PUT, DELETE, OPTIONS`)
		w.Header().Add("Access-Control-Allow-Headers", `content-type,accept`)
		w.Header().Add("Access-Control-Max-Age", "10")
		w.WriteHeader(http.StatusNoContent)
	}
}
