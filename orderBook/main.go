package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Order struct {
	UserID     int64
	Price      int64
	Amount     int64
	TotalPrice int64
	Currency   string
	IsBuy      bool
}

type OrderBook struct {
	BuyOrders  []Order
	SellOrders []Order
}

type Transaction struct {
	UserID   int64
	Value    int64
	Currency string
}

type TransactionsList struct {
	Transactions []Transaction
}

func (order *Order) InputOrder(isBuy bool, ob *OrderBook, tl *TransactionsList) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input UserID, price and amount of currency. Example: 11 40 115")
	for scanner.Scan() {
		fmt.Println("\nIf you want to stop input press Ctrl+Z")
		fmt.Println("Input UserID, price and amount of currency. Example: 11 40 115")
		str := strings.Split(scanner.Text(), " ")
		if len(str) == 3 {
			order.UserID, _ = strconv.ParseInt(str[0], 10, 64)
			order.Price, _ = strconv.ParseInt(str[1], 10, 64)
			order.Amount, _ = strconv.ParseInt(str[2], 10, 64)
			order.TotalPrice = order.Price * order.Amount
			order.IsBuy = isBuy
			ob.AddOrder(order, tl)
		} else {
			fmt.Println("\nInput order correctly. Example:11 40 115 ")
			continue
		}
	}
}

func (ob *OrderBook) AddOrder(order *Order, tl *TransactionsList) {
	if order.IsBuy {
		ob.BuyOrders = append(ob.BuyOrders, *order)
		ob.SortOrders()
		ob.MatchOrders(tl)
	} else {
		ob.SellOrders = append(ob.SellOrders, *order)
		ob.SortOrders()
		ob.MatchOrders(tl)
	}
}

func (ob *OrderBook) SortOrders() {
	sort.Slice(ob.BuyOrders, func(i, j int) bool {
		return ob.BuyOrders[i].Price > ob.BuyOrders[j].Price
	})

	sort.Slice(ob.SellOrders, func(i, j int) bool {
		return ob.SellOrders[i].Price < ob.SellOrders[j].Price
	})
}

// MatchOrders check if top SellOrder price <= top BuyOrder price.
// If condition is true crete transaction and remove Order from OrderBook
func (ob *OrderBook) MatchOrders(tl *TransactionsList) {
	for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
		buyOrder := ob.BuyOrders[0]
		sellOrder := ob.SellOrders[0]
		if sellOrder.Price <= buyOrder.Price {
			tl.AddTransaction(buyOrder)
			tl.AddTransaction(sellOrder)
			ob.RemoveOrder(buyOrder)
			ob.RemoveOrder(sellOrder)
		} else {
			break
		}
	}
}

// RemoveOrder remove Order from OrderBook
func (ob *OrderBook) RemoveOrder(order Order) {
	if order.IsBuy {
		for i, o := range ob.BuyOrders {
			if o.Price == order.Price && o.Amount == order.Amount {
				ob.BuyOrders = append(ob.BuyOrders[:i], ob.BuyOrders[i+1:]...)
				break
			}
		}
	} else {
		for i, o := range ob.SellOrders {
			if o.Price == order.Price && o.Amount == order.Amount {
				ob.SellOrders = append(ob.SellOrders[:i], ob.SellOrders[i+1:]...)
				break
			}
		}
	}
}

func (ob *OrderBook) GetOrdersList(getType int) {
	switch getType {
	case 1:

		for _, buyOrder := range ob.BuyOrders {
			fmt.Printf("| UserID | %v | Price | %v | Amount | %v | Side | Buy |\n", buyOrder.UserID, buyOrder.Price, buyOrder.Amount)
		}

		for _, sellOrder := range ob.SellOrders {
			fmt.Printf("| UserID | %v | Price | %v | Amount | %v | Side | Sell |\n", sellOrder.UserID, sellOrder.Price, sellOrder.Amount)
		}
	case 2:
		fmt.Println("Buy order book")
		for _, buyOrder := range ob.BuyOrders {
			fmt.Printf("| UserID | %v | Price | %v | Amount | %v | Side | Buy |\n", buyOrder.UserID, buyOrder.Price, buyOrder.Amount)
		}
	case 3:
		fmt.Println("Sell order book")
		for _, sellOrder := range ob.SellOrders {
			fmt.Printf("| UserID | %v | Price | %v | Amount | %v | Side | Sell |\n", sellOrder.UserID, sellOrder.Price, sellOrder.Amount)
		}
	default:
		break
	}
}

func (tl *TransactionsList) AddTransaction(order Order) {
	var transaction *Transaction
	transaction = new(Transaction)

	transaction.UserID = order.UserID
	transaction.Currency = order.Currency
	if order.IsBuy {
		transaction.Value = order.Amount
	} else {
		transaction.Value = -order.Amount
	}

	tl.Transactions = append(tl.Transactions, *transaction)

}

func (tl *TransactionsList) GetTransactionsList() {
	for _, transaction := range tl.Transactions {
		fmt.Printf("| User | %v | | %v | | %v |\n", transaction.UserID, transaction.Value, transaction.Currency)
	}
}

func showMenu() {
	menu := []string{"\nHello, dear UserID. What you want to do?(Input number of operation):\n",
		"1. Add buy order\n",
		"2. Add sell order\n",
		"3. Watch all orders\n",
		"4. Watch buy orders\n",
		"5. Watch sell orders\n",
		"6. Show transactions list\n",
		"7. Change currency\n"}

	for _, s := range menu {
		fmt.Print(s)
	}
}

func main() {

	orderUAH := Order{Currency: "UAH"}
	orderUSD := Order{Currency: "USD"}

	orderBookUAH := OrderBook{}
	orderBookUSD := OrderBook{}

	transactionsListUAH := TransactionsList{}
	transactionsListUSD := TransactionsList{}

	for {
		currencyChoice := 0
		fmt.Println("Choose currency(Input number of currency):\n" +
			"1. UAH\n" +
			"2. USD")

		_, err := fmt.Scanln(&currencyChoice)
		if err != nil {
			fmt.Printf("Error scaning currencyChoice. Error is %v", err)
		}

		for i := 1; i > 0; {
			switch currencyChoice {
			//UAH currency
			case 1:
				userChoice := 0
				showMenu()
				_, err = fmt.Scanln(&userChoice)
				if err != nil {
					fmt.Printf("Error scaning userChoice. Error is %v", err)
				}

				switch userChoice {
				//Add UAH buy order
				case 1:
					orderUAH.InputOrder(true, &orderBookUAH, &transactionsListUAH)

				//Add UAH sell order
				case 2:
					orderUAH.InputOrder(false, &orderBookUAH, &transactionsListUAH)

				//Get UAH all orders
				case 3:
					orderBookUAH.GetOrdersList(1)

				//Get UAH buy orders
				case 4:
					orderBookUAH.GetOrdersList(2)

				//Get UAH sell orders
				case 5:
					orderBookUAH.GetOrdersList(3)

				//Get UAH transactions list
				case 6:
					transactionsListUAH.GetTransactionsList()

				//Change currency
				case 7:
					i = 0

				default:
					break
				}

			//USD currency
			case 2:
				userChoice := 0
				showMenu()
				_, err = fmt.Scanln(&userChoice)
				if err != nil {
					fmt.Printf("Error scaning userChoice. Error is %v", err)
				}

				switch userChoice {

				//Add USD buy order
				case 1:
					orderUSD.InputOrder(true, &orderBookUSD, &transactionsListUSD)

				//Add USD sell order
				case 2:
					orderUSD.InputOrder(false, &orderBookUSD, &transactionsListUSD)

				//Get USD all orders
				case 3:
					orderBookUSD.GetOrdersList(1)

				//Get USD all orders
				case 4:
					orderBookUSD.GetOrdersList(2)

				//Get USD all orders
				case 5:
					orderBookUSD.GetOrdersList(3)

				//Get transactions list
				case 6:
					transactionsListUSD.GetTransactionsList()

				//Change currency
				case 7:
					i = 0

				default:
					break

				}
			}
		}
	}
}
