package models

// 订单
type Order struct {
	BaseModel
	UserId           uint64      // 买家 Id
	SellerId         uint64      // 卖家 Id
	ProductId        uint64      // 产品 Id
	ProductPackageId uint64      // 套餐 Id
	Nums             int         // 购买数量
	UnitPrice        float64     // 单价
	TotalPrice       float64     // 总价
	Status           OrderStatus // 订单状态
	CheckingNumber   string      // 物流单号
}

type OrderStatus int

const (
	Ordered  OrderStatus = OrderStatus(iota) //	下单
	Paid                                     //	支付
	Prepare                                  //	配货
	Deliver                                  //	出货 配送
	Finished                                 //	交易成功
	Canceled                                 //	交易失败
)

// 根据订单的 ID 来获取订单
func FindOrderById(orderId uint64) (*Order, error) {
	var ret Order
	err := db.First(&ret, orderId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取订单列表
func FindOrdersByUserId(userId uint64) ([]*Order, error) {
	var rets []*Order
	err := db.Find(&rets).Where("user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}

// 根据卖家 ID 来获取订单列表
func FindOrdersBySellerId(seller uint64) ([]*Order, error) {
	var rets []*Order
	err := db.Find(&rets).Where("seller_id = ?", seller).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}

// 新增订单
func InsertOrder(order *Order) error {
	err := db.Create(order).Error
	return err
}

//删除订单 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteOrderById(id uint64) error {
	err := db.Delete(Order{}, "id = ?", id).Error
	return err
}

//根据 order.ID 来更新其他相应的字段
func UpdateOrder(order *Order) error {
	var rec Order
	err := db.First(&rec, "id = ?", order.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Update(order).Error
	return err
}
