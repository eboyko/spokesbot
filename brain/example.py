from gensim import corpora, models
from gensim.models import LdaModel, LdaMulticore, LsiModel
import numpy
import gensim.downloader

text8 = gensim.downloader.load("text8")
data = [datum for datum in text8]

statements = ["I love data science",
              "I love coding in python",
              "I love building NLP tool",
              "This is a good phone",
              "This is a good TV",
              "This is a good laptop"]
statements_tokens = [[word for word in statement.split()] for statement in statements]
dictionary = corpora.Dictionary(statements_tokens)
corpus = [dictionary.doc2bow(item, allow_update=True) for item in statements_tokens]

for item in corpus:
    print([[dictionary[id], frequency] for id, frequency in item])

tfidf_model = models.TfidfModel(corpus, smartirs='ntc')

for item in tfidf_model[corpus]:
    print([[dictionary[id], numpy.around(frequency, decimals=2)] for id, frequency in item])

lda_model = LdaMulticore(corpus=corpus,
                         id2word=dictionary,
                         random_state=100,
                         num_topics=6,
                         passes=10,
                         chunksize=1000,
                         batch=False,
                         alpha='asymmetric',
                         decay=0.5,
                         offset=64,
                         eta=None,
                         eval_every=0,
                         iterations=100,
                         gamma_threshold=0.001,
                         per_word_topics=True)

lda_model.save('lda_model.model')

lsi_model = LsiModel(corpus=corpus, id2word=dictionary, num_topics=6, decay=0.5)
print(lsi_model.print_topics(-1))