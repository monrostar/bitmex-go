package restful

import (
    "encoding/base64"
    "errors"
    "github.com/qct/bitmex-go/swagger"
    "github.com/satori/go.uuid"
    "golang.org/x/net/context"
    "strings"
    "net/http"
)

type OrderApi struct {
    swaggerOrderApi *swagger.OrderApiService
    ctx             context.Context
}

func NewOrderApi(swaggerOrderApi *swagger.OrderApiService, ctx context.Context) *OrderApi {
    return &OrderApi{swaggerOrderApi: swaggerOrderApi, ctx: ctx}
}

func (o *OrderApi) LimitBuy(symbol string, orderQty float64, price float64, clientOrderIDPrefix string) (resp *http.Response, orderId string, err error) {
    if symbol == "" {
        return nil, "", errors.New("symbol can NOT be empty")
    }
    if price <= 0 {
        return nil, "", errors.New("price must be positive")
    }
    clOrdID := ""
    if clientOrderIDPrefix != "" {
        s := strings.Replace(base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes()), "=", "", -1)
        clOrdID = clientOrderIDPrefix + s
    }

    params := map[string]interface{}{
        "symbol":   symbol,
        "ordType":  "Limit",
        "orderQty": float32(orderQty),
        "price":    price,
        "clOrdID":  clOrdID,
    }
    order, response, err := o.swaggerOrderApi.OrderNew(o.ctx, symbol, params)
    if err != nil || response.StatusCode != 200 {
        return response, order.OrderID, err
    }
    return response, order.OrderID, nil
}

func (o *OrderApi) LimitSell(symbol string, orderQty float64, price float64, clientOrderIDPrefix string) (resp *http.Response, orderId string, err error) {
    if symbol == "" {
        return nil, "", errors.New("symbol can NOT be empty")
    }
    if price <= 0 {
        return nil, "", errors.New("price must be positive")
    }
    clOrdID := ""
    if clientOrderIDPrefix != "" {
        s := strings.Replace(base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes()), "=", "", -1)
        clOrdID = clientOrderIDPrefix + s
    }

    params := map[string]interface{}{
        "symbol":   symbol,
        "orderQty": float32(-orderQty),
        "price":    price,
        "clOrdID":  clOrdID,
    }
    order, response, err := o.swaggerOrderApi.OrderNew(o.ctx, symbol, params)
    if err != nil || response.StatusCode != 200 {
        return response, order.OrderID, err
    }
    return response, order.OrderID, nil
}
