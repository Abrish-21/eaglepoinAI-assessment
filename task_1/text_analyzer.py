import string


class TextAnalyser:
    def __init__(self, text: str):
        self.raw_text = text
        self.words = self.clean_text(text)

    def clean_text(self, text: str): # Cleans and splits the text into words
        
        if not text:
            return []
        lower_text = text.lower()
        translator = str.maketrans('', '', string.punctuation) # Remove punctuation
        clean_text = lower_text.translate(translator)
        
        return clean_text.split()

    def get_total_word_count(self):
        return len(self.words)

    def get_average_word_length(self): # finds average word length
        if not self.words:
            return 0.00
        total_characters = sum(len(word) for word in self.words)
        average = total_characters / len(self.words)
        return round(average, 2)

    def get_longest_word(self) : # finds the longest word(s)
        if not self.words:
            return []
        max_len = 0
        for word in self.words:
            if len(word) > max_len:
                max_len = len(word)
        longest_words = sorted(list(set(word for word in self.words if len(word) == max_len)))
        return longest_words

    def get_word_frequency(self):# gets the frequency of each word
        frequency_map = {}
        for word in self.words:
            if word in frequency_map:
                frequency_map[word] += 1
            else:
                frequency_map[word] = 1
        return frequency_map

    def analyze(self): # compiles all analysis results into a dictionary
        return {
            "word_count": self.get_total_word_count(),
            "average_word_length": self.get_average_word_length(),
            "longest_words": self.get_longest_word(),
            "word_frequency": self.get_word_frequency()
        }


input_text = "The quick brown fox jumps over the lazy dog the fox"
analyzer = TextAnalyser(input_text)
result = analyzer.analyze()
print(result)