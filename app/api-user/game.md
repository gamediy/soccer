# 游戏设计

参考网站 https://luckyrp.com

## 服务端 + 后台

**框架**

- kratos grcp
- websocket
- 数据库 mongoDB

**功能接口**

- 基础框架 (martin 预计周期 2周) **重点1**
- 游戏网管 （martin） **重点2**

- 用户 （whisper, martin 预计周期 1周到2周）
    - 注册：允许用户创建账户并进行游戏。
    - 登录/注销：提供用户登录和注销功能，确保安全性和用户身份验证。
    - 个人资料：允许用户管理和编辑个人信息，如用户名、密码、头像等。
    - 提现银行：为用户提供绑定和管理提现银行账户的功能，方便用户提现游戏中的奖金。
- 钱包 （abri,whisper, martin)
    - 充值：允许用户通过各种支付方式向游戏钱包充值，如银行卡、支付宝、微信支付等。
    - 提现：允许用户从游戏钱包中提现到绑定的银行账户。
    - 交易记录：展示用户的充值和提现历史记录，包括时间、金额和状态等信息，方便用户查看和管理资金流动。
- 游戏 (martin,whisper,abri,barry,apple, 预计周期 3周) **重点3**
    - 游戏种类：提供多种游戏选项，包括solt、龙虎、红黑、捕鱼等，使用户能够选择自己喜欢的游戏进行娱乐。
    - 游戏配置：管理游戏规则、赔率、限额等配置，确保游戏的公平性和可玩性。
    - Solt：设计并实现具有各种不同主题和特性的老虎机游戏，包括不同的赔率和奖励机制。 2周
    - 龙虎：提供经典的龙虎斗游戏，让用户可以在龙和虎之间下注并比较两者的牌面大小。 1周
    - 红黑：提供红黑游戏，让用户根据牌面的颜色进行下注，判断下一张牌的颜色。
    - 捕鱼：设计捕鱼游戏，允许用户通过操作虚拟的渔网或炮台来捕捉不同种类的鱼，获得奖励。
- 机器人 (abri,barry,apple,whisper,martin, 预计周期 1周) **重点4**
- 游戏压力测试 (abri,apple,martin, 0.5周) **重点5**
- vip (apple,barry,martin 预计周期 1周)
    - VIP等级：设立多个VIP等级，根据用户的游戏活跃度和充值金额等因素来划分。
    - 特权和奖励：为VIP用户提供特殊的游戏奖励、赠送和活动优惠等特权，以增加他们的游戏体验和忠诚度。
- 活动 (barry,martin)
    - 定期活动：举办各种定期活动，如周末特惠、节日活动等，为用户提供额外的奖励和优惠。
    - 比赛和竞赛：组织定期的游戏比赛和竞赛，让用户之间相互竞争，获得额外的奖励和荣誉。
- 通知系统 (barry,apple,martin 预计周期 1周)
    - 系统通知：向用户发送重要的系统通知，如账户变动、活动提醒等。
    - 个人通知：为用户提供个性化的通知设置，允许他们选择接收关于游戏和活动的通知，并根据偏好进行定制。

## 客户端 **重点6**

### 游戏大厅

- 技术：游戏大厅的前端可以使用 Cocos Creator 进行设计和开发，后端可以使用 Golang 和 Kratos 进行游戏大厅的管理和逻辑处理。
- 周期：游戏大厅的开发周期取决于其复杂度和功能要求。预计开发周期可能在2周到3周之间。
- 技术人员: abri,apple,joker,jerry,ale,martin

**模块**

- 多游戏展示：在游戏大厅中展示可供选择的多个游戏: Solt、龙虎、红黑、捕鱼等游戏。
- 游戏筛选和分类：提供游戏筛选和分类功能，让用户可以按照游戏类型、热门程度或其他标准来浏览和选择游戏。
- 游戏推荐：根据用户的游戏偏好和历史记录等信息，提供个性化的游戏推荐，增加用户的游戏体验。
- 在线用户显示：展示当前在线的用户列表，让用户可以看到其他玩家的存在，并与他们进行互动和交流。
- 系统通知：向用户发送重要的系统通知，如账户变动、活动提醒等。
- 公告和活动信息：在游戏大厅中展示重要的公告和活动信息，如新游戏上线、比赛通知、奖励活动等。
- 社交功能：提供社交功能，如聊天室或私聊，让用户可以与其他玩家进行交流和互动。
- 用户统计和排行榜：展示用户的游戏统计信息，如胜率、积分、等级等，并提供排行榜功能，让用户可以比较自己与其他玩家的成绩。

总周期 ： 预计 12周/30=2.5个月左右 出两款游戏