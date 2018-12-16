import textract
import sys

if __name__ == '__main__':
    file = sys.argv[1]
    text = textract.process(file)
    text = text.decode('utf-8')
    text = text.replace('\x0c', ' ').replace('\n', ' ').replace('\t', '  ').replace(
        '\r', ' ').replace('\\', '').replace('*', '').replace('  ', ' ')
    to = sys.argv[2]
    with open(to, 'w+') as written:
        written.write(text)
