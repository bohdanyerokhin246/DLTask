package main

import (
	"fmt"
	"sort"
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

func (ob *OrderBook) AddOrder(order Order, tl *TransactionsList) {
	if order.IsBuy {
		ob.BuyOrders = append(ob.BuyOrders, order)
		ob.SortOrders()
		ob.MatchOrders(tl)
	} else {
		ob.SellOrders = append(ob.SellOrders, order)
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

	orderBookUAH := OrderBook{
		BuyOrders: []Order{
			{UserID: 1, Price: 40, Amount: 115, TotalPrice: 4600, Currency: "UAH", IsBuy: true},
			{UserID: 2, Price: 41, Amount: 111, TotalPrice: 4551, Currency: "UAH", IsBuy: true},
			{UserID: 3, Price: 42, Amount: 20, TotalPrice: 840, Currency: "UAH", IsBuy: true},
			{UserID: 4, Price: 43, Amount: 256, TotalPrice: 11008, Currency: "UAH", IsBuy: true},
			{UserID: 5, Price: 44, Amount: 189, TotalPrice: 8316, Currency: "UAH", IsBuy: true}},
		SellOrders: []Order{
			{UserID: 6, Price: 44, Amount: 78, TotalPrice: 3432, Currency: "UAH", IsBuy: false},
			{UserID: 7, Price: 43, Amount: 14, TotalPrice: 602, Currency: "UAH", IsBuy: false},
			{UserID: 8, Price: 42, Amount: 128, TotalPrice: 5376, Currency: "UAH", IsBuy: false},
			{UserID: 9, Price: 41, Amount: 99, TotalPrice: 4059, Currency: "UAH", IsBuy: false},
			{UserID: 10, Price: 40, Amount: 64, TotalPrice: 2560, Currency: "UAH", IsBuy: false}},
	}
	orderUAH := Order{Currency: "UAH"}
	orderBookUSD := OrderBook{
		BuyOrders: []Order{
			{UserID: 11, Price: 40, Amount: 115, TotalPrice: 4600, Currency: "USD", IsBuy: true},
			{UserID: 12, Price: 41, Amount: 111, TotalPrice: 4551, Currency: "USD", IsBuy: true},
			{UserID: 13, Price: 42, Amount: 20, TotalPrice: 840, Currency: "USD", IsBuy: true},
			{UserID: 14, Price: 43, Amount: 256, TotalPrice: 11008, Currency: "USD", IsBuy: true},
			{UserID: 15, Price: 44, Amount: 189, TotalPrice: 8316, Currency: "USD", IsBuy: true}},
		SellOrders: []Order{
			{UserID: 16, Price: 44, Amount: 78, TotalPrice: 3432, Currency: "USD", IsBuy: false},
			{UserID: 17, Price: 43, Amount: 14, TotalPrice: 602, Currency: "USD", IsBuy: false},
			{UserID: 18, Price: 42, Amount: 128, TotalPrice: 5376, Currency: "USD", IsBuy: false},
			{UserID: 19, Price: 41, Amount: 99, TotalPrice: 4059, Currency: "USD", IsBuy: false},
			{UserID: 20, Price: 40, Amount: 64, TotalPrice: 2560, Currency: "USD", IsBuy: false}},
	}
	orderUSD := Order{Currency: "USD"}

	transactionsListUAH := TransactionsList{}
	transactionsListUSD := TransactionsList{}

	orderBookUAH.SortOrders()
	orderBookUSD.SortOrders()

	for {
		currencyChoice := 0
		fmt.Println("Choose currency(Input number of currency):\n" +
			"1. UAH\n" +
			"2. USD")

		fmt.Scanln(&currencyChoice)

		for i := 1; i > 0; {
			switch currencyChoice {
			//UAH currency
			case 1:
				userChoice := 0
				showMenu()
				fmt.Scanln(&userChoice)

				switch userChoice {
				//Add UAH buy order
				case 1:
					fmt.Println("Input your user ID")
					fmt.Scanln(&orderUAH.UserID)
					fmt.Println("Input price")
					fmt.Scanln(&orderUAH.Price)
					fmt.Println("Input amount")
					fmt.Scanln(&orderUAH.Amount)
					orderUAH.TotalPrice = orderUAH.Price * orderUAH.Amount
					orderUAH.IsBuy = true
					orderBookUAH.AddOrder(orderUAH, &transactionsListUAH)

				//Add UAH sell order
				case 2:
					fmt.Println("Input your user ID")
					fmt.Scanln(&orderUAH.UserID)
					fmt.Println("Input price")
					fmt.Scanln(&orderUAH.Price)
					fmt.Println("Input amount")
					fmt.Scanln(&orderUAH.Amount)
					orderUAH.TotalPrice = orderUAH.Price * orderUAH.Amount
					orderUAH.IsBuy = false
					orderBookUAH.AddOrder(orderUAH, &transactionsListUAH)

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
				fmt.Scanln(&userChoice)

				switch userChoice {

				//Add USD buy order
				case 1:
					fmt.Println("Input your user ID")
					fmt.Scanln(&orderUSD.UserID)
					fmt.Println("Input price")
					fmt.Scanln(&orderUSD.Price)
					fmt.Println("Input amount")
					fmt.Scanln(&orderUSD.Amount)
					orderUSD.TotalPrice = orderUSD.Price * orderUSD.Amount
					orderUSD.IsBuy = true
					orderBookUSD.AddOrder(orderUSD, &transactionsListUSD)

				//Add USD sell order
				case 2:
					fmt.Println("Input your user ID")
					fmt.Scanln(&orderUSD.UserID)
					fmt.Println("Input price")
					fmt.Scanln(&orderUSD.Price)
					fmt.Println("Input amount")
					fmt.Scanln(&orderUSD.Amount)
					orderUSD.TotalPrice = orderUSD.Price * orderUSD.Amount
					orderUSD.IsBuy = false
					orderBookUSD.AddOrder(orderUSD, &transactionsListUSD)

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
