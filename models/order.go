package models

import (
	"fmt"
	"funding/enums"
	"funding/forms"
	"funding/objects"
	"funding/resultModels"
	"time"
	"unicode/utf8"
)

// 订单
type Order struct {
	BaseModel
	UserId           uint64            // 买家 Id
	Name             string            // 收件人姓名
	Address          string            // 收件人地址
	Phone            string            // 收件人电话
	SellerId         uint64            // 卖家 Id
	ProductId        uint64            // 产品 Id
	ProductPackageId uint64            // 套餐 Id
	Nums             int               // 购买数量
	UnitPrice        float64           // 单价
	TotalPrice       float64           // 总价
	Status           enums.OrderStatus // 订单状态
	CheckingNumber   string            // 物流单号
	PaidAt           *time.Time        // 支付时间
	CloseAt          *time.Time        // 关闭时间
	FinishedAt       *time.Time        // 交易成功时间
}

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
	err := db.Where("user_id = ?", userId).Find(&rets).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}

// 根据卖家 ID 来获取订单列表
func FindOrdersBySellerId(seller uint64) ([]*Order, error) {
	var rets []*Order
	err := db.Where("seller_id = ?", seller).Find(&rets).Error
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
	err = db.Model(&rec).Updates(order).Error
	return err
}

// 从表单信息新增订单
func NewOrderFromForm(userId uint64, form *forms.NewOrderForm) ([]uint64, error) {
	// 开始事务
	tx := db.Begin()
	// 地址信息不能为空
	if form.Name == "" || form.Address == "" || form.Phone == "" {
		return nil, &resultError.AddressInfoErr
	}
	orders := []Order{}
	orderIdList := []uint64{}
	// 多个订单循环插入
	for _, item := range form.OrderPkgList {
		// 检查产品 ID 是否存在
		if tx.Find(&Product{BaseModel: BaseModel{ID: item.ProductID}}).RecordNotFound() {
			tx.Rollback()
			return nil, &resultError.ProductNotFound
		}
		// 检查套餐 ID 是否存在
		if tx.Find(&ProductPackage{BaseModel: BaseModel{ID: item.ProductPackageID}}).RecordNotFound() {
			tx.Rollback()
			return nil, &resultError.ProductPkgNotFound
		}
		// 根据表单创建新订单
		newOrder := Order{
			UserId:           userId,
			Name:             form.Name,
			Address:          form.Address,
			Phone:            form.Phone,
			SellerId:         item.SellerID,
			ProductId:        item.ProductID,
			ProductPackageId: item.ProductPackageID,
			UnitPrice:        item.Price,
			Nums:             item.Nums,
			TotalPrice:       item.Price * float64(item.Nums),
			Status:           enums.OrderStatus_Ordered,
		}
		// 向数据库中插入新订单
		err := tx.Create(&newOrder).Error
		if err != nil {
			tx.Rollback()
			return nil, &resultError.OrderCreateErr
		}
		// 将添加的订单信息添加到 slice 中
		orders = append(orders, newOrder)
		orderIdList = append(orderIdList, newOrder.ID)
	}
	// 提交事务
	tx.Commit()

	// 这里返回订单列表是为了之后付款时修改状态用
	//return orders, nil

	// 还是返回订单 Id 列表，方便后续结算
	return orderIdList, nil
}

const sqlSelectOrderListField = `
SELECT
	o.id,o.user_id,p.user_id AS seller_id,su.nickname AS seller_nickname,
	o.product_package_id,o.nums,o.unit_price,pkg.product_id,pkg.freight,p.end_time,
	p.name AS product_name,pkg.price,pkg.stock,pkg.image_url,pkg.description,pkg.stock,
	p.current_price,p.target_price,
	o.created_at,o.status AS order_status,o.total_price,o.checking_number,
	o.name,o.address,o.phone,o.paid_at,o.finished_at,o.close_at
`

// 查询订单列表的 SQL 字段
const sqlOrderListTable = `
FROM
	orders o
JOIN
	product_packages pkg ON o.product_package_id = pkg.id
JOIN
	products p ON p.id = o.product_id
JOIN 
	users su ON su.id = p.user_id
`

// 根据页码和用户信息来获取订单列表
//const sqlGetOrderList = sqlSelectOrderListField + sqlOrderListTable + `
//WHERE
//	o.deleted_at IS NULL  AND
//	p.deleted_at IS NULL  AND
//	pkg.deleted_at IS NULL AND
//	o.user_id = (?)
//ORDER BY
//	o.created_at DESC
//LIMIT ? OFFSET ?
//`
const sqlCountAsTotal = `SELECT COUNT(1) as total`

// 根据页码和用户信息获取订单
func GetOrderListByUserId(form *forms.SellerGetOrderListForm, userId uint64, roleId int) (*resultModels.OrderList, error) {
	result := resultModels.OrderList{}
	var list []*resultModels.OrderListItem

	page, pageSize := 1, 10
	// 如果页码和每页数量大于 0
	if form.Page > 0 && form.PageSize > 0 {
		page = form.Page
		pageSize = form.PageSize
	}
	sql := `WHERE
	o.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL
	`
	// 判断角色，如果是商家就查找 seller_id
	if roleId == 2 {
		sql = sql + ` AND o.seller_id = (?) `
	} else {
		sql = sql + ` AND o.user_id = (?) `
	}

	// 如果条件有产品 ID
	if form.ProductId > 0 {
		sql = sql + `AND o.product_id = ` + fmt.Sprintf("(%d)", form.ProductId)
	}
	// 如果条件有订单状态
	if form.OrderStatus > 0 {
		sql = sql + `AND o.status = ` + fmt.Sprintf("(%d)", form.OrderStatus)
	}
	// 如果条件有众筹状态
	if form.FundingStatus > 0 {
		switch form.FundingStatus {
		case enums.FundingStatus_Success:
			success := `p.end_time < CURRENT_TIMESTAMP() AND p.current_price >= p.target_price`
			sql = sql + `AND ` + success
		case enums.FundingStatus_Fail:
			fail := `p.end_time < CURRENT_TIMESTAMP() AND p.current_price < p.target_price`
			sql = sql + `AND ` + fail
		case enums.FundingStatus_Ing:
			ing := `p.end_time > CURRENT_TIMESTAMP()`
			sql = sql + `AND ` + ing
		}
	}

	// 统计数量
	err := db.Raw(sqlCountAsTotal+sqlOrderListTable+sql, userId).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	// 分页限制，这里可不能漏掉了 order by 前面的空格
	sql = sql + ` ORDER BY o.created_at DESC LIMIT ? OFFSET ?`

	// 根据 SQL 字符串拼接查询订单相关信息列表
	err = db.Raw(sqlSelectOrderListField+sqlOrderListTable+sql, userId, pageSize, (page-1)*pageSize).
		Scan(&list).Error
	if err != nil {
		return nil, err
	}

	// 返回众筹状态 TODO 这里做一个每天凌晨的定时器去更新状态到数据库里会省事一点
	if form.FundingStatus > 0 {
		for i := range list {
			list[i].FundingStatus = form.FundingStatus
			// 如果为已付款状态而且众筹成功了则为备货状态
			if list[i].FundingStatus == enums.FundingStatus_Success &&
				list[i].OrderStatus == enums.OrderStatus_Paid {
				list[i].OrderStatus = enums.OrderStatus_Prepare
			}
		}
	} else {
		timeNow := time.Now()
		for i := range list {
			list[i].FundingStatus = CalcFundingStatus(timeNow, list[i].EndTime,
				list[i].CurrentPrice, list[i].TargetPrice)
			// 如果为已付款状态而且众筹成功了则为备货状态
			if list[i].FundingStatus == enums.FundingStatus_Success &&
				list[i].OrderStatus == enums.OrderStatus_Paid {
				list[i].OrderStatus = enums.OrderStatus_Prepare
			}
		}
	}

	result.Page = page
	result.OrderList = list
	return &result, err
}

const sqlGetOrderListInOrderIds = sqlSelectOrderListField + sqlOrderListTable + `
WHERE
	o.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	o.user_id = (?) AND
    o.id IN (?)
ORDER BY
	o.created_at DESC
`

const sqlSellerGetOrderList = sqlSelectOrderListField + sqlOrderListTable + `
WHERE
	o.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	o.seller_id = (?) AND
    o.id IN (?)
ORDER BY
	o.created_at DESC`

// 根据订单id列表和用户 Id 查询订单信息
func GetOrderListByOrderIds(orderIds []uint64, userId uint64, roleId int) ([]*resultModels.OrderListItem, error) {
	var list []*resultModels.OrderListItem
	var sql string
	switch roleId {
	case 0:
		sql = sqlGetOrderListInOrderIds
	case 2:
		sql = sqlSellerGetOrderList
	}
	// 根据 SQL 字符串拼接查询订单相关信息列表
	err := db.Raw(sql, userId, orderIds).Scan(&list).Error
	timeNow := time.Now()
	for i := range list {
		list[i].FundingStatus = CalcFundingStatus(timeNow, list[i].EndTime,
			list[i].CurrentPrice, list[i].TargetPrice)
		// 如果为已付款状态而且众筹成功了则为备货状态
		if list[i].FundingStatus == enums.FundingStatus_Success &&
			list[i].OrderStatus == enums.OrderStatus_Paid {
			list[i].OrderStatus = enums.OrderStatus_Prepare
		}
	}
	return list, err
}

// 根据订单 ID 列表来支付
// 订单支付，支付后需要 更新订单状态、产品增加众筹金额和人数、相应套餐减少库存增加支持人数
// 需要在事务中处理，错了一步就全部回退并返回错误信息
func PayOrderByOrderIdList(orderIds []uint64) error {
	// 开始事务
	tx := db.Begin()
	timeNow := time.Now()
	for _, orderId := range orderIds {
		// 先查询出对应订单信息
		order := Order{}
		err := tx.Last(&order, orderId).Error
		if err != nil {
			tx.Rollback()
			return resultError.NewFallFundingErr("订单查询出错")
		}
		if order.Status != enums.OrderStatus_Ordered {
			tx.Rollback()
			return resultError.NewFallFundingErr("订单已支付或已关闭")
		}
		// 根据订单 ID 将订单状态更新为已支付
		order.Status = enums.OrderStatus_Paid
		err = tx.Model(&order).
			Updates(map[string]interface{}{"status": order.Status, "paid_at": timeNow}).Error

		//err = tx.Save(&order).Error
		if err != nil {
			tx.Rollback()
			return resultError.NewFallFundingErr("订单状态发生错误")
		}
		// 对应的产品增加支持人数 以及筹集金额
		product := Product{}
		err = tx.Last(&product, order.ProductId).Error
		if err != nil {
			tx.Rollback()
			return resultError.NewFallFundingErr("获取产品时发生错误")
		}

		if time.Now().After(product.EndTime) {
			tx.Rollback()
			return resultError.NewFallFundingErr("众筹已结束")
		}

		// 增加支持者人数
		product.Backers++
		// 增加筹集金额
		product.CurrentPrice += order.TotalPrice
		//err = tx.Model(&product).Update("backers", "current_price").Error
		err = tx.Model(&product).
			Updates(map[string]interface{}{"backers": product.Backers, "current_price": product.CurrentPrice}).Error

		//err = tx.Save(&product).Error
		if err != nil {
			tx.Rollback()
			return resultError.NewFallFundingErr("产品状态更新时发生错误")
		}
		// 相应的套餐减少库存 增加支持人数
		pkg := ProductPackage{}
		err = tx.Last(&pkg, order.ProductPackageId).Error
		if err != nil {
			tx.Rollback()
			return resultError.NewFallFundingErr("获取套餐时发生错误")
		}
		// 增加支持者人数
		pkg.Backers++
		// 减少库存
		pkg.Stock -= int64(order.Nums)
		if pkg.Stock < 0 {
			tx.Rollback()
			return resultError.NewFallFundingErr("库存不足")
		}
		err = tx.Model(&pkg).Updates(map[string]interface{}{"backers": pkg.Backers, "stock": pkg.Stock}).Error
		//err = tx.Save(&pkg).Error
		if err != nil {
			tx.Rollback()
			return resultError.NewFallFundingErr("套餐状态更新时发生错误")
		}
	}
	// 提交事务
	tx.Commit()
	return nil
}

/////////////////// 			商家相关					/////////////////

// 发货
func SendOutOrderById(form *forms.OrderSendOutForm, sellerId uint64) error {
	// 如果物流单号小于 4 则返回错误
	if utf8.RuneCountInString(form.CheckingNumber) < 4 {
		return &resultError.FormParamErr
	}
	order := Order{}

	// 首先查找相关订单记录
	err := db.Last(&order, form.OrderId).Error
	if err != nil {
		return resultError.NewFallFundingErr("没有找到相关订单")
	}
	// 如果订单不是这个卖家的则返回错误
	if order.SellerId != sellerId {
		return resultError.NewFallFundingErr("这不是你的订单")
	}
	// 如果订单状态不是已支付或者待发货，就返回错误
	if order.Status != enums.OrderStatus_Paid && order.Status != enums.OrderStatus_Prepare {
		return resultError.NewFallFundingErr("订单状态有误")
	}
	// 将状态改为已发货，并更新订单号
	err = db.Model(&order).
		Updates(map[string]interface{}{"status": enums.OrderStatus_Deliver,
			"checking_number": form.CheckingNumber}).Error
	if err != nil {
		return err
	}
	return nil
}
