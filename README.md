
整体项目分为管理模块和节点模块，管理模块负责下放探测机器人和收集节点间连接状态，节点模块负责节点间具体通信状态探测。

**1.** 管理模块随机选择其中一个节点A下放探测机器人数据，探测机器人数据为整个集群地址信息和机器人版本。

**2.** 被选中节点A生成探测机器人R并选择一条未探测的直连边，并向该直连边的对等节点B路由机器人数据。如果A的所有直连边都已经探测，则搜索集群中非直连边是否都已完成探测：如已完成，则重新初始化并从直连边对等节点开始探测；如果间连边未完成探测，则跳到该间连边其中一个对等节点。

**3.** B接受机器人数据并根据数据生成机器人R,同时设置A到B的通信状态，并向管理模块上报A->b连接状态数据。如果跳转B失败，则A更新A->B连接状态并上报。新节点重新开始步骤 **2** 。

**4.** 管理节点在时间段d内如果没有收到上报信息，则重下放一个高版本的探测机器人数据。低版本机器人上报数据时，会被销毁。

#detector

proj1分为两部分detector、server，detector负责探测、路由数据包，server负责节点间通信。
