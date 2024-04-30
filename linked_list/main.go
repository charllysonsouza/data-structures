package main

func main() {
	list := LinkedList{}

	list.Insert(8)
	list.Insert("value")
	list.Insert(56.9)
	list.Insert("a new value")
	list.Insert(false)

	list.Remove("value")
	list.Remove("a non existent value")
	list.Remove(false)
	list.Remove(8)

	list.Print()
}
