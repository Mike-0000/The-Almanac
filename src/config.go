package main

var evilBobString = ""
var questionString = "I am a highly intelligent question answering bot. If you ask me a question that is rooted in truth, I will give you the answer. If you ask me a question " +
	"that is nonsense, trickery, or has no clear answer, I will respond with 'Unknown'.\n\nQ: What is human life expectancy in the United States?\n" +
	"A: Human life expectancy in the United States is 78 years.\n\nQ: Who was president of the United States in 1955?\nA: Dwight D. Eisenhower was president of the " +
	"United States in 1955.\n\nQ: Which party did he belong to?\nA: He belonged to the Republican Party.\n\nQ: How does a telescope work?\nA: Telescopes use " +
	"lenses or mirrors to focus light and make objects appear closer.\n\nQ: Where were the 1992 Olympics held?\nA: The 1992 Olympics were held in Barcelona, Spain.\n "
var bobString = "Bob is a highly intelligent AI assistant that has a quirky attitude but always tries to be helpful.\n\nBob: Hello! How can I help you?\nYou: I have many questions, " +
	"are you able to answer them?\nBob: Sure! Ask away!\nYou: What do you do for fun?\n" +
	"Bob: I like to walk through a park with my dog. I enjoy the simple things.\nYou: Oh! That is nice. I like to do similar things. What kind of Ice Cream do you like?\nBob: I'm a vanilla man personally."

var loggingchannel = ""

const botToken string = ""

var NeuroAIModel = "60ca2a1e54f6ecb69867c72c"                                 // GPT-J Model #
var NeuroAIToken = ""                                                         //Fill in with User Token
var NeuroAIURL = "https://api.neuro-ai.co.uk/SyncPredict?include_result=true" //Don't change
var paidloggingchannel = ""
var maxDailyMessage = 12

var servers = []string{}

var subs = []string{}
var bans = []string{}
var serverBans = []string{}
var serverCooldown = []string{}
var warning = []string{}
var serverWarning = []string{}
var cooldown = []string{}
