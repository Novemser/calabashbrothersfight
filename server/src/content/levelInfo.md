
label: 初级教程3
Title: 装甲再厚也能打破
Description: 葫芦兄弟的火娃被蛇精打伤，落单了，这时候蝎子精穿着他的新盔甲要抓他，尽管他的盔甲用上了互斥锁，但是还是有弱点，抓住时机，利用两个线程，一举击败他。


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

