package p181920substringsearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type IndexTest struct {
	s   string
	sep string
	out int
}

var indexTests = []IndexTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "foo", -1},
	{"fo", "foo", -1},
	{"foo", "foo", 0},
	{"oofofoofooo", "f", 2},
	{"oofofoofooo", "foo", 4},
	{"barfoobarfoo", "foo", 3},
	{"foo", "", 0},
	{"foo", "o", 1},
	{"abcABCabc", "A", 3},
	{"jrzm6jjhorimglljrea4w3rlgosts0w2gia17hno2td4qd1jz", "jz", 47},
	{"ekkuk5oft4eq0ocpacknhwouic1uua46unx12l37nioq9wbpnocqks6", "ks6", 52},
	{"999f2xmimunbuyew5vrkla9cpwhmxan8o98ec", "98ec", 33},
	{"9lpt9r98i04k8bz6c6dsrthb96bhi", "96bhi", 24},
	{"55u558eqfaod2r2gu42xxsu631xf0zobs5840vl", "5840vl", 33},
	// cases with one byte strings - test special case in Index()
	{"", "a", -1},
	{"x", "a", -1},
	{"x", "x", 0},
	{"abc", "a", 0},
	{"abc", "b", 1},
	{"abc", "c", 2},
	{"abc", "x", -1},
	// test special cases in Index() for short strings
	{"", "ab", -1},
	{"bc", "ab", -1},
	{"ab", "ab", 0},
	{"xab", "ab", 1},
	{"xab"[:2], "ab", -1},
	{"", "abc", -1},
	{"xbc", "abc", -1},
	{"abc", "abc", 0},
	{"xabc", "abc", 1},
	{"xabc"[:3], "abc", -1},
	{"xabxc", "abc", -1},
	{"", "abcd", -1},
	{"xbcd", "abcd", -1},
	{"abcd", "abcd", 0},
	{"xabcd", "abcd", 1},
	{"xyabcd"[:5], "abcd", -1},
	{"xbcqq", "abcqq", -1},
	{"abcqq", "abcqq", 0},
	{"xabcqq", "abcqq", 1},
	{"xyabcqq"[:6], "abcqq", -1},
	{"xabxcqq", "abcqq", -1},
	{"xabcqxq", "abcqq", -1},
	{"", "01234567", -1},
	{"32145678", "01234567", -1},
	{"01234567", "01234567", 0},
	{"x01234567", "01234567", 1},
	{"x0123456x01234567", "01234567", 9},
	{"xx01234567"[:9], "01234567", -1},
	{"", "0123456789", -1},
	{"3214567844", "0123456789", -1},
	{"0123456789", "0123456789", 0},
	{"x0123456789", "0123456789", 1},
	{"x012345678x0123456789", "0123456789", 11},
	{"xyz0123456789"[:12], "0123456789", -1},
	{"x01234567x89", "0123456789", -1},
	{"", "0123456789012345", -1},
	{"3214567889012345", "0123456789012345", -1},
	{"0123456789012345", "0123456789012345", 0},
	{"x0123456789012345", "0123456789012345", 1},
	{"x012345678901234x0123456789012345", "0123456789012345", 17},
	{"", "01234567890123456789", -1},
	{"32145678890123456789", "01234567890123456789", -1},
	{"01234567890123456789", "01234567890123456789", 0},
	{"x01234567890123456789", "01234567890123456789", 1},
	{"x0123456789012345678x01234567890123456789", "01234567890123456789", 21},
	{"xyz01234567890123456789"[:22], "01234567890123456789", -1},
	{"", "0123456789012345678901234567890", -1},
	{"321456788901234567890123456789012345678911", "0123456789012345678901234567890", -1},
	{"0123456789012345678901234567890", "0123456789012345678901234567890", 0},
	{"x0123456789012345678901234567890", "0123456789012345678901234567890", 1},
	{"x012345678901234567890123456789x0123456789012345678901234567890", "0123456789012345678901234567890", 32},
	{"xyz0123456789012345678901234567890"[:33], "0123456789012345678901234567890", -1},
	{"", "01234567890123456789012345678901", -1},
	{"32145678890123456789012345678901234567890211", "01234567890123456789012345678901", -1},
	{"01234567890123456789012345678901", "01234567890123456789012345678901", 0},
	{"x01234567890123456789012345678901", "01234567890123456789012345678901", 1},
	{"x0123456789012345678901234567890x01234567890123456789012345678901", "01234567890123456789012345678901", 33},
	{"xyz01234567890123456789012345678901"[:34], "01234567890123456789012345678901", -1},
	{"xxxxxx012345678901234567890123456789012345678901234567890123456789012", "012345678901234567890123456789012345678901234567890123456789012", 6},
	{"", "0123456789012345678901234567890123456789", -1},
	{"xx012345678901234567890123456789012345678901234567890123456789012", "0123456789012345678901234567890123456789", 2},
	{"xx012345678901234567890123456789012345678901234567890123456789012"[:41], "0123456789012345678901234567890123456789", -1},
	{"xx012345678901234567890123456789012345678901234567890123456789012", "0123456789012345678901234567890123456xxx", -1},
	{"xx0123456789012345678901234567890123456789012345678901234567890120123456789012345678901234567890123456xxx", "0123456789012345678901234567890123456xxx", 65},
	// test fallback to Rabin-Karp.
	{"oxoxoxoxoxoxoxoxoxoxoxoy", "oy", 22},
	{"oxoxoxoxoxoxoxoxoxoxoxox", "oy", -1},
}

// Execute f on each test case.  funcName should be the name of f; it's used
// in failure reports.
func runIndexTests(t *testing.T, f func(s, sep string) int, funcName string, testCases []IndexTest) {
	for _, test := range testCases {
		actual := f(test.s, test.sep)
		if actual != test.out {
			t.Errorf("%s(%q,%q) = %v; want %v", funcName, test.s, test.sep, actual, test.out)
		}
	}
}

func TestSearchSimple(t *testing.T) {
	runIndexTests(t, SearchSimple, "SearchSimple", indexTests)
}
func TestSearchSimpleWithShiftPrefix(t *testing.T) {
	runIndexTests(t, SearchSimpleWithShiftPrefix, "SearchSimpleWithShiftPrefix", indexTests)
}

func TestSearchSimpleWithShiftSuffix(t *testing.T) {
	runIndexTests(t, SearchSimpleWithShiftSuffix, "SearchSimpleWithShiftSuffix", indexTests)
}
func TestSearchBoyerMoore(t *testing.T) {
	runIndexTests(t, SearchBoyerMoore, "SearchBoyerMoore", indexTests)
}

const benchmarkString = `To Sherlock Holmes she is always the woman.
 I have seldom heard him mention her under any other name.  
 In his eyes she eclipses and predominates the whole of her sex.
  It was not that he felt any emotion akin to love for Irene Adler. 
  All emotions, and that one particularly, were abhorrent to his cold, precise but admirably balanced mind. 
  He was, I take it, the most perfect reasoning and observing machine that the world has seen, but as a lover 
  he would have placed himself in a false position. He never spoke of the softer passions, save with a gibe and a sneer. 
  They were admirable things for the observer--excellent for drawing the veil from men's motives and actions.
  But for the trained reasoner to admit such intrusions into his own delicate and finely adjusted temperament was to introduce 
  a distracting factor which might throw a doubt upon all his mental results. Grit in a sensitive instrument, or 
  a crack in one of his own high-power lenses, would not be more disturbing than a strong emotion in a nature such as his.
 And yet there was but one woman to him, and that woman was the late Irene Adler, of dubious and questionable memory.
 I had seen little of Holmes lately. My marriage had drifted us away from each other. 
 My own complete happiness, and the home-centred interests which rise up around the man who first finds himself master 
 of his own establishment, were sufficient to absorb all my attention, while Holmes, 
 who loathed every form of society with his whole Bohemian soul, remained in our lodgings in Baker Street, 
 buried among his old books, and alternating from week to week between cocaine and ambition, 
 the drowsiness of the drug, and the fierce energy of his own keen nature. He was still, as ever, 
 deeply attracted by the study of crime, and occupied his immense faculties and extraordinary powers 
 of observation in following out those clues, and clearing up those mysteries which had been abandoned 
 as hopeless by the official police. From time to time I heard some vague account of his doings: of his 
 summons to Odessa in the case of the Trepoff murder, of his clearing up of the singular tragedy of the 
 Atkinson brothers at Trincomalee, and finally of the mission which he had accomplished so delicately and 
 successfully for the reigning family of Holland. Beyond these signs of his activity, however, which I merely 
 shared with all the readers of the daily press, I knew little of my former friend and companion.`

const benchmarkSubString = "companion"

func BenchmarkSearchSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchSimple(benchmarkString, benchmarkSubString)
	}
}

func BenchmarkSearchSimpleWithShiftPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchSimpleWithShiftPrefix(benchmarkString, benchmarkSubString)
	}
}

func BenchmarkSearchSimpleWithShiftSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchSimpleWithShiftSuffix(benchmarkString, benchmarkSubString)
	}
}

func BenchmarkSearchBoyerMoore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchBoyerMoore(benchmarkString, benchmarkSubString)
	}
}

func TestCreateStateMachineAndSearch(t *testing.T) {
	delta := CreateStateMachine("AABAABAAABA")
	index := SearchByStateMachine("AABAABAABAAABA", delta)
	require.Equal(t, 3, index)
}

func TestCreatePi(t *testing.T) {
	pi1 := CreatePiSlow("AABAABAAABA")
	pi2 := CreatePiFast("AABAABAAABA")
	require.Equal(t, pi1, pi2)
}

func TestSearchKnuthMorrisPratt(t *testing.T) {
	runIndexTests(t, SearchKnuthMorrisPratt, "SearchKnuthMorrisPratt", indexTests)
}

func BenchmarkSearchKnuthMorrisPratt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchKnuthMorrisPratt(benchmarkString, benchmarkSubString)
	}
}
