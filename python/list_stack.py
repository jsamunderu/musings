import random

class Empty(Exception):
	pass

class ArrayStack:
	def __init__(_self):
		_self.data = []

	def __len__(_self):
		return len(_self.data)

	def is_empty(_self):
		return len(_self.data) == 0

	def push(_self, element):
		_self.data.append(element)

	def top(_self):
		if _self.is_empty():
			raise Empty('Stack is empty')
		return _self.data[-1]

	def pop(_self):
		if _self.is_empty():
			raise Empty('Stack is empty')
		return _self.data.pop()

def match(str):
	stack = ArrayStack() 
	for c in str:
		if c == '{':
			stack.push('{')
		if c == '}':
			try:
				stack.pop()
			except Empty:
				return False
	return stack.is_empty()

def reverse(str):
	stack = ArrayStack()
	for c in str:
		stack.push(c)	
	tmp = []
	size = len(stack)
	for i in range(size):
		tmp.append(stack.pop())
	return ''.join(tmp)

if __name__ == '__main__':
	if match('{{}}{}}'):
		print ('{{}}{}} match')
	else:
		print ('{{}}{}} DONT match')

	if match('{{}}{}'):
		print ('{{}}{} match')
	else:
		print ('{{}}{} {value}'.format(value='DONT match'))

	str = 'james'
	print ('[{}] reversed [{}]'.format(str, reverse(str)))