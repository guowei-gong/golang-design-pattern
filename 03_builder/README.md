## 生成器模式
分步骤创建复杂对象。使用多个构造方法比仅仅使用一个复杂可怕的构造函数更简单。 

分为多个步骤进行构建的潜在问题是， 构建不完整的和不稳定的产品可能会被暴露给客户端。 生成器模式能够在产品完成构建之前使其处于私密状态。

### 实例相关（选读）
`iglooBuilder` 冰屋生成器与 `normalBuilder` 普通房屋生成器可建造不同类型房屋， 即 `igloo` 冰屋和 `normalHouse` 普通房屋。每种房屋类型的建造步骤都是相同的。