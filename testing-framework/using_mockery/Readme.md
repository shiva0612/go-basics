in folder service
1.  we had implemented mocks on our own
2.  if we observe (both the interface and method are not exported)
  still it works fine

with mockery while generating the code 
1.  by default it needs the interface and method be to exported
2.  cmd used (mockery --name=<interface_name>)