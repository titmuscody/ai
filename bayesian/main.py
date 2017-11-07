
import glob
import re
import math

def filter_lines(line):
	''' takes a line and returns the words on that line in lowercase '''
	words = re.sub(r'[;:!,.?"]', '', line).split()
	return [w.lower() for w in words]


def get_stats(filename):
	stats = {}
	with open(filename) as f:
		for line in f:
			for word in filter_lines(line):
				stats[word] = stats.get(word, 0) + 1
	return stats

def get_words_in_docs(text_files):
	out = []
	for filename in text_files:
		with open(filename) as f:
			words = []
			for line in f:
				words.extend(filter_lines(line))
		out.extend(words[:1000])

	return out

def get_vocabulary(words):
	words = []
	for file in text_files:
		for word in get_stats(file).keys():
			words.append(word)
	return {x for x in words}

def get_word_freq(words):
	out = {}
	for word in words:
		out[word] = out.get(word, 0) + 1
	return out


if __name__ == '__main__':
	test = {}
	reg = re.compile('training/(.*)_(.*).txt')
	matches = [reg.match(val) for val in glob.glob('training/*.txt')]
	authors = {cur.group(2) for cur in matches}
	# set of tuple with (filename, author)
	books = {(m.group(0), m.group(2)) for m in matches}	
	# dictionary of author to bookname, filepath
	#auth_to_books = {auth: [(match.group(1), match.group(0)) for match in matches if auth == match.group(2)] for auth in authors}

	tot_correct = 0
	tot_wrong = 0
	for test_book in books:
		# get list of books for testing
		testing_books = {book for book in books if test_book[0] != book[0]}
		# get vocab using list of filenames
		tot_vocab = set(get_words_in_docs([x[0] for x in testing_books]))

		test_words = get_words_in_docs([test_book[0]])

		print('test book', test_book[0])
		cur_best = -1000000
		cur_best_auth = ''
		for auth in authors:
			auth_books = [book[0] for book in testing_books if book[1] == auth]
			total_books = len(testing_books)
			book_prob = len(auth_books) / total_books


			# get the word seq of books written by the same author
			words = get_words_in_docs(auth_books)
			auth_vocab = set(words)
			#n = len(words)

			word_count = get_word_freq(words)

			tot = math.log(book_prob)
			for word in test_words:
				try:
					num = word_count.get(word, 0) + 1
					#num = word_count[word] + 1
					#denom = n + len(tot_vocab)
					denom = len(auth_vocab) + len(tot_vocab)
					val = num / denom
					tot += math.log(val)
				except:
					pass

			#tot = math.exp(tot)


			if tot > cur_best:
				cur_best = tot
				cur_best_auth = auth

			print('author', auth)
			print(tot)
		if cur_best_auth == test_book[1]:
			print('correct choice')
			tot_correct += 1
		else:
			print('incorrect choice')
			tot_wrong += 1
	print('correct', tot_correct)
	print('incorrect', tot_wrong)
	print(tot_correct / (tot_correct + tot_wrong) * 100, 'percent')


