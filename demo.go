package main

import "fmt"

func calcTotal(price int, voucher string) int {
	var discont int = 0
	if voucher == "FAZZFOOD50" {
		if price >= 50000 {
			discont = price * 50 / 100
			if discont > 50000 {
				discont = 50000
			}
		}
	}
	if voucher == "DITRAKTIR60" {
		if price >= 25000 {
			discont = price * 60 / 100
			if discont > 50000 {
				discont = 50000
			}
		}
	}
	return discont
}

func calcTax(tax bool, price int) int {
	var taxValue int = 0
	if tax == true {
		taxValue = price * 5 / 100
	}
	return taxValue
}

func calcDistance(distance int) int {
	var shippingFee int = 5000
	if distance > 2 {
		shippingFee = shippingFee + (distance-2)*3000
	}
	return shippingFee
}

func fazzFoodAndDitraktir(price int, voucher string, distance int, tax bool) {
	var discont int = calcTotal(price, voucher)
	var shippingFee int = calcDistance(distance)
	var taxValue int = calcTax(tax, price)

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