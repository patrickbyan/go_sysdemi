package main

import (
	"fmt"
)

const (
	buyerName = "Belinda"
	bornDate  = 1635875790371
)

func UnitEligibilityCheck(userUnitId ...int) (bool, string) {
	unitEligible := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var errMessage string
	var err bool

	var passCheck int = 0

	for _, unitValid := range unitEligible {
		for _, unit := range userUnitId {
			if unit == unitValid-1 {
				passCheck++
			}
		}
	}

	if passCheck == len(userUnitId) {
		err = false
		errMessage = "All unit are eligible"
	} else {
		err = true
		errMessage = "Check your units!"
	}

	return err, errMessage
}

func main() {
	fmt.Println("Begin Transaction \n=====================")
	defer fmt.Println("=====================\nEnd Transaction")

	activeStatus, discrepancy, unit := getDataBuyer(buyerName)

	if activeStatus {
		err, errMessage := UnitEligibilityCheck(unit...)

		if err {
			fmt.Printf("error: %v, message: %s \n", err, errMessage)
		} else {
			newBuyer := dataBuyer{
				name:         buyerName,
				activeStatus: activeStatus,
				discrepancy:  int64(discrepancy),
			}

			unitNames, totalPrice, metaData := newBuyer.PurchaseUnits(unit...)
			fmt.Printf("error: false, message: purchase %v units (total: %v data) with a price of %v success \n", unitNames, metaData, totalPrice)
		}
	} else {
		fmt.Println("error: true, message: You are born literally today")
	}
}

// version change: tambahin validasi unit, buatin receipt, pindahin ke function logger, unit hapus dari parameter function, tambahin coupon
