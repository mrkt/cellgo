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

import (
	"log"
	"time"
)

var (
	ExchangeMap map[int]*Exchange = make(map[int]*Exchange)
)

//Create a Exchange
func CreateExchange(tcpType int) {
	exchange := &Exchange{
		Exchanges: make(map[string]*exchanges),
	}
newEX:
	res, _ := exchange.NewExchange(tcpType)
	if !res {
		time.Sleep(time.Second * 1)
		goto newEX
		log.Println("Try to start Tcp Exchage ...")
	}
	ExchangeMap[tcpType] = exchange
	log.Println("Tcp Exchage has been started.")
}

//Exchange Operation type
type exchanges struct {
	ExchangeName   string            //Exchange's name
	ExchangeNumber string            //Exchange's number
	Queue          map[string]*Queue //Exchange's Queue
	PushedNum      int               //Exchange's total push
	PulledNum      int               //Exchange's total pull
}

//Exchange Operation type
type Exchange struct {
	Exchanges map[string]*exchanges //Exchange's child

}

//Create a Exchange
func (e *Exchange) NewExchange(tcpType int) (bool, error) {
	res, err := Bind[tcpType].BindMaps["New"].handler("New", nil)
	if err != nil {
		return false, err
	}
	exchange := res.(map[string]string)
	for k, v := range exchange {
		e.Exchanges[k] = &exchanges{
			ExchangeName:   v,
			ExchangeNumber: k,
			Queue:          make(map[string]*Queue),
			PushedNum:      0,
			PulledNum:      0,
		}
	}
	return true, nil
}

//Renew Exchange data
func (e *Exchange) RenewExchange(eventName string, controllerName string, funcName string) (bool, error) {
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
