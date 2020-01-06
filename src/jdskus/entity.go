package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func NewProduct() Product {
	uid := uuid.Must(uuid.NewV4())
	return Product{
		createdDate:        time.Now(),
		lastModifiedDate:   time.Now(),
		version:            1,
		name:               "",
		hits:               1,
		isActive:           false,
		isDelivery:         false,
		isList:             false,
		isMarketable:       false,
		isTop:              false,
		price:              0,
		marketPrice:        0,
		maxCommission:      0,
		weekHits:           0,
		weekHitsDate:       time.Now(),
		weekSales:          0,
		weekSalesDate:      time.Now(),
		monthHits:          0,
		monthHitsDate:      time.Now(),
		monthSales:         0,
		monthSalesDate:     time.Now(),
		sales:              0,
		score:              0,
		scoreCount:         0,
		sn:                 uid.String(),
		totalScore:         0,
		_type:              0,
		productCategory_id: 1, // 默认产品类别
		dimension:          "20*20*20",
	}
}

type Product struct {
	id                 int
	createdDate        time.Time
	lastModifiedDate   time.Time
	version            int
	name               string
	hits               int
	isActive           bool
	isDelivery         bool
	isList             bool
	isMarketable       bool
	isTop              bool
	price              float32
	marketPrice        float32
	maxCommission      float32
	weekHits           int
	weekHitsDate       time.Time
	weekSales          int
	weekSalesDate      time.Time
	monthHits          int
	monthHitsDate      time.Time
	monthSales         int
	monthSalesDate     time.Time
	sales              int
	score              float32
	scoreCount         int
	sn                 string
	totalScore         int
	_type              int
	productCategory_id int
	dimension          string
}

func NewSku() Sku {
	uid := uuid.Must(uuid.NewV4())
	return Sku{
		id:               0,
		createdDate:      time.Now(),
		lastModifiedDate: time.Now(),
		version:          1,
		allocatedStock:   1,
		exchangePoint:    1,
		isDefault:        false,
		marketPrice:      1.0,
		maxCommission:    1.0,
		price:            0,
		rewardPoint:      0,
		sn:               uid.String(),
		stock:            0,
		product_id:       0,
	}
}

type Sku struct {
	id               int
	createdDate      time.Time
	lastModifiedDate time.Time
	version          int
	allocatedStock   int
	exchangePoint    int
	isDefault        bool
	marketPrice      float32
	maxCommission    float32
	price            float32
	rewardPoint      int
	sn               string
	stock            int
	product_id       int
}

func NewJDSku() JDSku {
	return JDSku{
		id:         0,
		quantity:   0,
		product_id: 0,
	}
}

type JDSku struct {
	id         int
	quantity   int
	product_id int
	jdPrice    float32
	skuName    string
}
