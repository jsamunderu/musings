import random

class _Node:
	__slot__ = 'element', 'next'

	def __init__(_self, element, next):
		_self.element = element
		_self.next = next

def add_node(listHead, element):
	if listHead is None:
		listHead = _Node(element, None)
	else:
		tmp = _Node(element, listHead)
		listHead = tmp
	return listHead

def remove_greater(listHead, x):
	while listHead is not None:
		if listHead.element > x:
			tmp = listHead
			listHead = listHead.next
			tmp.next = None
			tmp = None
		else:
			break
	if listHead is None:
		return listHead
	prev = listHead
	current = listHead.next
	while current is not None:
		if current.element > x:
			tmp = current
			prev.next = current.next
			current = prev.next
			tmp.next = None
			tmp = None
		else:
			prev = current
			current = current.next
	return listHead

if __name__ == "__main__":
	listHead = None
	for i in range(0, 10):
		listHead = add_node(listHead, random.randrange(1, 10))

	tmp = listHead
	while tmp is not None:
		print (tmp.element)
		tmp = tmp.next

	listHead = remove_greater(listHead, 5)
	tmp = listHead
	while tmp is not None:
		print (tmp.element)
		tmp = tmp.next

