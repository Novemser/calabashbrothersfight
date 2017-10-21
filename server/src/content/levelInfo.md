


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

