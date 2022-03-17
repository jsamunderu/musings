import random

class Stack:
	class Node:
		__slot__ = 'element', 'next'
		def __init__(_self, element, next):
			_self.element = element
			_self.next = next

	def __init__(_self):
		_self.top = None
		_self.size = 0

	def __len__(_self):
		return _self.size

	def push(_self, element):
		node = _self.Node(element, _self.top)
		_self.top = node
		_self.size += 1

	def pop(_self):
		if _self.size == 0:
			return None
		node = _self.top
		_self.top = _self.top.next
		node.next = None
		_self.size -= 1
		return node

if __name__ == '__main__':
	stack = Stack()
	for i in range(0, 10):
		value = random.randrange(1, 10)
		stack.push(value)
		print (value)

	print ('Stack size: {}'.format(len(stack)))

	while True:
		value = stack.pop()
		if value is None:
			break
		print (value.element)
	
	print ('Stack size: {}'.format(len(stack)))