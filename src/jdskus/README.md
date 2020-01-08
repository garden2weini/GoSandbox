# 将JDSku转为系统的Product和Sku

## 假设前提
- JDSkus和Product共用ProductCategory定义
- 此场景中，Product与SKU是一一对应关系
- 此转换工具仅减少运营人员日常工作，但不能替代。转换后的Product均未上架/失效状态，待运维人员确认并完善数据后手动进行上架操作

## 主要步骤
1. 获取JDSkus表中指定类型的记录（比如：食品饮料／矿泉水）
1. 解析JDSkus记录
   1. 选取JDSkus.skuName内容作为Product.name。维护List(JDSkus),List(Sku)和Map(Product.name, Product)。如果Product.name已经存在，则使用已有Product。
   2. TODO: 从JDSkus.skuName中获取包装数量，并赋值给JDSkus.quantity.
   3. JDSkus.jdPrice作为Product.price/sku.price
   4. 如果JDSkus.productCategory_id不为空，则Product.productCategory_id; 否则，Product.productCategory_id=1。
   5. 通过IdGenerator获得Product.id和Sku.id, JDSkus.product_id=Product.id, Sku.product_id=Product.id
   6. TODO:Product.productImages处理，及其与JDSkus.imagePath关系
2. 遍历List(JDSkus),List(Sku)和Map(Product.name, Product)，创建Product和Sku记录，更新JDSkus

注意：多个JDSkus记录会对应一个Product记录。