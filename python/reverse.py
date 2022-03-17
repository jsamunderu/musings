def reverse_words(sentence):
	words = sentence.replace('.', '').lower().split()

	for i in range(0, len(words)):
		min = i
		for j in range(i+1, len(words)):
			if len(words[min]) > len(words[j]):
				min = j
		tmp = words[min]
		for x in range(min, i, -1):
			words[x] = words[x-1]
		words[i] = tmp
	words[0] = words[0][0].upper() + words[0][1: len(words[0])]
	return ' '.join(words) + '.'

if __name__ == "__main__":
	sentence = 'Are these word the ones we need to load in reverse.'
	print (sentence)
	print (reverse_words(sentence))
