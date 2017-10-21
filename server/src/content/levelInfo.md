while (true) {
  Monitor.Enter(mutex);
  i = i + 2;
  if (i == 5) {
     critical_section();
  }
  Monitor.Exit(mutex);
}


while (true) {
  critical_section();
  Monitor.Enter(mutex);
  i = i - 1;
  Monitor.Exit(mutex);
}