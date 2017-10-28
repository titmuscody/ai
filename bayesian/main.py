
import glob
import re

def filter_lines(line):
	''' takes a line and returns the words on that line in lowercase '''
	words = re.sub(r'[\'-;:!,.?"]', '', line).split()
	return [w.lower() for w in words]


def get_stats(filename):
	stats = {}
	with open(filename) as f:
		for line in f:
			for word in filter_lines(line):
				stats[word] = stats.get(word, 0) + 1
	return stats

if __name__ == '__main__':
	test = {}
	reg = re.compile('training/(.*)_(.*).txt')
	matches = [reg.match(val) for val in glob.glob('training/*.txt')]
	authors = {cur.group(2) for cur in matches}
	# dictionary of author to bookname, filepath
	books = {auth: [(match.group(1), match.group(0)) for match in matches if auth == match.group(2)] for auth in authors}

	stats = {}
	for auth in authors:
		tot_stats = {}
		for book_name, filename in books[auth]:
			cur_stat = get_stats(filename)
			for word in cur_stat.keys():
				tot_stats[word] = tot_stats.get(word, 0) + cur_stat[word]
		stats[auth] = tot_stats
	print(stats)
	



