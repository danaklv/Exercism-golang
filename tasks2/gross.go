package tasks

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	meas := make(map[string]int)
    meas["quarter_of_a_dozen"] = 3 
    meas["half_of_a_dozen"] = 6
    meas["dozen"] = 12
    meas["small_gross"] = 120
    meas["gross"] = 144
    meas["great_gross"] = 1728
    return meas
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	// Добавляем товар в счет или увеличиваем его количество
	if _, itemExists := bill[item]; itemExists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, itemExists := bill[item]
	if !itemExists {
		return false
	}
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}
	newQuantity := quantity - unitValue
	if newQuantity < 0 {
		return false
	}
	if newQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = newQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, itemExists := bill[item]
    if !itemExists {
        return 0, false
    } else {
        return quantity, true
    }
}