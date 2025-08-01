// Code generated by Vandor. DO NOT EDIT.

package enum

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Description for OrderStatus enum
type OrderStatus string

const (
	// Order has been cancelled
	OrderStatusCancelled OrderStatus = "CANCELLED"
	// Order has been delivered
	OrderStatusDelivered OrderStatus = "DELIVERED"
	// Order has failed
	OrderStatusFailed OrderStatus = "FAILED"
	// Order is on hold
	OrderStatusOnHold OrderStatus = "ON_HOLD"
	// Order is pending
	OrderStatusPending OrderStatus = "PENDING"
	// Order is being processed
	OrderStatusProcessing OrderStatus = "PROCESSING"
	// Order has been refunded
	OrderStatusRefunded OrderStatus = "REFUNDED"
	// Order has been returned
	OrderStatusReturned OrderStatus = "RETURNED"
	// Order has been shipped
	OrderStatusShipped OrderStatus = "SHIPPED"
)

var AllOrderStatuss = []OrderStatus{
	OrderStatusCancelled,
	OrderStatusDelivered,
	OrderStatusFailed,
	OrderStatusOnHold,
	OrderStatusPending,
	OrderStatusProcessing,
	OrderStatusRefunded,
	OrderStatusReturned,
	OrderStatusShipped,
}

var OrderStatusEnum = struct {
	Cancelled    OrderStatus
	Delivered    OrderStatus
	Failed    OrderStatus
	OnHold    OrderStatus
	Pending    OrderStatus
	Processing    OrderStatus
	Refunded    OrderStatus
	Returned    OrderStatus
	Shipped    OrderStatus
}{
	Cancelled:    OrderStatusCancelled,
	Delivered:    OrderStatusDelivered,
	Failed:    OrderStatusFailed,
	OnHold:    OrderStatusOnHold,
	Pending:    OrderStatusPending,
	Processing:    OrderStatusProcessing,
	Refunded:    OrderStatusRefunded,
	Returned:    OrderStatusReturned,
	Shipped:    OrderStatusShipped,
}

var orderStatusLabels = map[OrderStatus]string{
	OrderStatusCancelled:    "Order has been cancelled",
	OrderStatusDelivered:    "Order has been delivered",
	OrderStatusFailed:    "Order has failed",
	OrderStatusOnHold:    "Order is on hold",
	OrderStatusPending:    "Order is pending",
	OrderStatusProcessing:    "Order is being processed",
	OrderStatusRefunded:    "Order has been refunded",
	OrderStatusReturned:    "Order has been returned",
	OrderStatusShipped:    "Order has been shipped",
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusCancelled,
		OrderStatusDelivered,
		OrderStatusFailed,
		OrderStatusOnHold,
		OrderStatusPending,
		OrderStatusProcessing,
		OrderStatusRefunded,
		OrderStatusReturned,
		OrderStatusShipped:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e OrderStatus) Label() string {
	if label, ok := orderStatusLabels[e]; ok {
		return label
	}
	return string(e)
}

func ParseOrderStatus(s string) (OrderStatus, error) {
	switch strings.ToUpper(s) {
	case "CANCELLED":
		return OrderStatusCancelled, nil
	case "DELIVERED":
		return OrderStatusDelivered, nil
	case "FAILED":
		return OrderStatusFailed, nil
	case "ON_HOLD":
		return OrderStatusOnHold, nil
	case "PENDING":
		return OrderStatusPending, nil
	case "PROCESSING":
		return OrderStatusProcessing, nil
	case "REFUNDED":
		return OrderStatusRefunded, nil
	case "RETURNED":
		return OrderStatusReturned, nil
	case "SHIPPED":
		return OrderStatusShipped, nil
	}
	return "", fmt.Errorf("invalid OrderStatus: %s", s)
}

func (e OrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(e))
}

func (e *OrderStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	val, err := ParseOrderStatus(s)
	if err != nil {
		return err
	}
	*e = val
	return nil
}
