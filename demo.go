package main

import "fmt"

func fazzFoodAndDitraktir(price int, voucher string, distance int, tax bool)  {
	var discont int = 0
	var shippingFee int = 5000
	var taxValue int = 0

	if voucher == "FAZZFOOD50" {
		if(price >= 50000) {
			discont = price * 50 / 100 
			if(discont > 50000) {
				discont = 50000
			}
		} 
	} 
	if voucher == "DITRAKTIR60" {
		if(price >= 25000) {
			discont = price * 60 / 100 
			if(discont > 50000) {
				discont = 50000
			}
		} 
	} 

	if  distance > 2  {
		shippingFee = shippingFee + (distance - 2) * 3000;
	}


	
	if  tax == true  {
		taxValue = price * 5 / 100 
	}

	var subTotal int = price - discont + shippingFee + taxValue
	fmt.Printf("harga: %d\n", price)
	fmt.Printf("Potongan: %d\n", discont)
	fmt.Printf("Biaya Antar: %d\n", shippingFee)
	fmt.Printf("Pajak: %d\n", taxValue)
	fmt.Printf("Sub Total: %d\n", subTotal)


}

func main() {
	fazzFoodAndDitraktir(25000, "DITRAKTIR60", 1, false)
}