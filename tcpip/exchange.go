//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//| ------------------------------------------------------------------
//| Cellgo Framework tcpip/exchange file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-08-06

package tcpip

//Exchange Operation type
type Exchange struct {
	ExchangeName   string            //Exchange's name
	ExchangeNumber string            //Exchange's number
	Queue          map[string]*Queue //Exchange's Queue
	PushedNum      int64             //Exchange's total push
	PulledNum      int64             //Exchange's total pull
}

//Create a Exchange
func (e *Exchange) CreateExchange(name string, number string) (bool, error) {
	return true, nil
}

//Renew Exchange data
func (e *Exchange) RenewExchange(style int, value map[string]interface{}) (bool, error) {
	return true, nil
}

//Destroy a Exchange
func (e *Exchange) DestroyExchange(number string) (bool, error) {
	return true, nil
}

//Allow an Queue to enter the Exchange
func (e *Exchange) IncreaseQueue(queue *Queue, carryInfo string) (bool, error) {
	return true, nil
}
