package calc

func Total(price int, voucher string) int {
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

func Tax(tax bool, price int) int {
	var taxValue int = 0
	if tax == true {
		taxValue = price * 5 / 100
	}
	return taxValue
}

func Distance(distance int) int {
	var shippingFee int = 5000
	if distance > 2 {
		shippingFee = shippingFee + (distance-2)*3000
	}
	return shippingFee
}
