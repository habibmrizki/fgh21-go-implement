package main

import (
	"fazztrack/demo/calc"
	"fmt"
)



func fazzFoodAndDitraktir(price int, voucher string, distance int, tax bool) {
	var discont int = calc.Total(price, voucher)
	var shippingFee int = calc.Distance(distance)
	var taxValue int = calc.Tax(tax, price)

	var subTotal int = price - discont + shippingFee + taxValue
	fmt.Printf("harga: %d\n", price)
	fmt.Printf("Potongan: %d\n", discont)
	fmt.Printf("Biaya Antar: %d\n", shippingFee)
	fmt.Printf("Pajak: %d\n", taxValue)
	fmt.Printf("Sub Total: %d\n", subTotal)
}
func main() {
	fazzFoodAndDitraktir(25000, "DITRAKTIR60", 1, true)
}