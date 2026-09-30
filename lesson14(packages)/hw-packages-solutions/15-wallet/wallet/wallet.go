package wallet

// Wallet - кошелёк. Поле balance с маленькой буквы - неэкспортируемое,
// снаружи пакета к нему не добраться напрямую, только через методы ниже.
type Wallet struct {
	balance float64
}

// Deposit пополняет баланс.
func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

// Balance возвращает текущий баланс.
func (w *Wallet) Balance() float64 {
	return w.balance
}
