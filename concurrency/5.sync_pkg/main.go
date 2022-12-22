package main

func main() {
	sync_map()
	sync_once()
}

/*

readers use rlock and runlock
writer use lock and unlock

readers(r1 r2 r3 ...) rlock - read - runlock
they fight each other on rlock

but when writer locks (only one writer - obviously)
all readers should wait

*/
