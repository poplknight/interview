# interview

1、	
		已经完成
2、	
		缺少一个根据输入长度改变结构体内位数的函数，时间关系这个就无法写了
3、	
		Test开头的文件即测试代码，10亿理想状态位数设定为42位，无碰撞（现在并不能设定位数，固定为64位）
4、	
		将三个关键字分别存入redis队列，现有123三个队列，先从1队列取第一个形容词获取值后头插入回去，2队和1队相同，3队取公司，拿到公司后，插入到结尾，依次轮询
		如果要求随机，将12队读取时顺序打乱，读取时，读过即尾插，每读过一遍重新打乱一遍123队（有碰撞发生，这时必须用去重方法。）

附：
		这些题可能是做过最难的了，9成的东西都没有接触过，很多东西都是现学现用，感谢能给我这个答题吧，让我学到了很多。
