in folder service
  we had implemented mocks on our own
  if we observe (both the interface and method are not exported)
  still it works fine

with mockery while generating the code 
    by default it needs the interface and method be to exported
    cmd used (mockery --name=<interface_name>)