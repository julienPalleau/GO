// https://www.boot.dev/lessons/a800cb84-5d09-4635-941a-93a0c90f9f06

/*
Billing System

Textio is making some changes to how they bill users for bulk messages. The system should now calculate the bill based on:

    The number of messages sent
    The cost per message
    Thresholds for discounts

Assignment

Complete the calculateFinalBill and the calculateDiscountRate functions.
calculateFinalBill

Use the calculateBaseBill function to get the cost for the messages sent. Then, use the calculateDiscountRate function to get the discount rate. Finally, calculate the final bill by applying the discount to the base bill and return the result.

The discount is a percentage represented by a float e.g. 10% = .1. Remember that any percentage x% is equal to x / 100.
calculateDiscountRate

This function should take the number of messages sent, and return the relevant discount based on the following discount rates:

    10% for more than 1000 messages.
    20% for more than 5000 messages.
    0% for anything less.
*/

package main

import "fmt"

func calculateFinalBill(costPerMessage float64, numMessage int) float64 {
    finalBill := calculateBaseBill(costPerMessage, numMessage) - (calculateBaseBill(costPerMessage, numMessage)) * calculateDiscountRate(numMessage)
    return finalBill
}

func calculateDiscountRate(messageSent int) float64 {
    if messageSent > 1000 && messageSent <= 5000 {
        return 0.1
    }
    if messageSent > 5000 {
        return 0.2
    }
    return 0

}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messageSent int) float64 {
    return costPerMessage * float64(messageSent)
}

func main() {
    fmt.Print(calculateFinalBill(10, 20000))

}