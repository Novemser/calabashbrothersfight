
label: 初级教程3
Title: 装甲再厚也能打破
Description: 葫芦兄弟的火娃被白骨精打伤，落单了，这时候蜘蛛精穿着她的蛛丝做的新盔甲要抓他，尽管她的盔甲用上了互斥锁，但是还是有弱点，抓住时机，利用两个线程，一举击败她。


for {
	mutex.Lock()
	i = i + 2
	if i == 5 {
     critical_section()
  }
	mutex.Unlock()
  
}

for {
	critical_section()
	mutex.Lock()
	i = i - 1
	mutex.Unlock()
 }


////////////////////////////////////////////////
                 分割线
////////////////////////////////////////////////

label: 初级教程4
Title: 武功再高也怕菜刀
Description: 葫芦兄弟正在和白骨精正面战斗，老爷爷这时却遇到了麻烦，蛇精化身为3头巨蛇怪来吃他，没有葫芦兄弟那样本领的爷爷只好手握菜刀，伺机而动。他发现这三头巨蛇虽然强大，但是并不灵活，爷爷想到用死锁的方法，然这3个头互相攻击，循环等待，击败巨蛇。


for {
    mutexA.lock()
    mutexB.lock()
    mutexA.unlock()
    mutexB.unlock()
}

for {
    mutexB.lock()
    mutexC.lock()
    mutexB.unlock()
    mutexC.unlock()
}

for {
    mutexA.lock()
    mutexC.lock()
    mutexA.unlock()
    mutexC.unlock()
}



