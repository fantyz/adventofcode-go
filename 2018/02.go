package main

import (
	"fmt"
	"strings"
)

/*

--- Day 2: Inventory Management System ---
You stop falling through time, catch your breath, and check the screen on the device. "Destination reached. Current Year: 1518. Current Location: North Pole Utility Closet 83N10." You made it! Now, to find those anomalies.

Outside the utility closet, you hear footsteps and a voice. "...I'm not sure either. But now that so many people have chimneys, maybe he could sneak in that way?" Another voice responds, "Actually, we've been working on a new kind of suit that would let him fit through tight spaces like that. But, I heard that a few days ago, they lost the prototype fabric, the design plans, everything! Nobody on the team can even seem to remember important details of the project!"

"Wouldn't they have had enough fabric to fill several boxes in the warehouse? They'd be stored together, so the box IDs should be similar. Too bad it would take forever to search the warehouse for two similar box IDs..." They walk too far away to hear any more.

Late at night, you sneak to the warehouse - who knows what kinds of paradoxes you could cause if you were discovered - and use your fancy wrist device to quickly scan every box and produce a list of the likely candidates (your puzzle input).

To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing exactly two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary checksum and compare it to what your device predicts.

For example, if you see the following box IDs:

abcdef contains no letters that appear exactly two or three times.
bababc contains two a and three b, so it counts for both.
abbcde contains two b, but no letter appears exactly three times.
abcccd contains three c, but no letter appears exactly two times.
aabcdd contains two a and two d, but it only counts once.
abcdee contains two e.
ababab contains three a and three b, but it only counts once.
Of these box IDs, four of them contain a letter which appears exactly twice, and three of them contain a letter which appears exactly three times. Multiplying these together produces a checksum of 4 * 3 = 12.

What is the checksum for your list of box IDs?

Your puzzle answer was 8820.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Confident that your list of box IDs is complete, you're ready to find the boxes full of prototype fabric.

The boxes will have IDs which differ by exactly one character at the same position in both strings. For example, given the following box IDs:

abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
The IDs abcde and axcye are close, but they differ by two characters (the second and fourth). However, the IDs fghij and fguij differ by exactly one character, the third (h and u). Those must be the correct boxes.

What letters are common between the two correct box IDs? (In the example above, this is found by removing the differing character from either ID, producing fgij.)

*/

func main() {
	fmt.Println("Day 2")
	fmt.Println(" > Checksum:", Checksum(strings.Split(puzzleInput, "\n")))
	fmt.Println(" > Common letters:", CommonLetters(strings.Split(puzzleInput, "\n")))
}

func CommonLetters(in []string) string {
	longest := 0
	shortestCommon := ""

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			diff := 0
			common := ""
			for idx := 0; idx < len(in[i]); idx++ {
				if in[i][idx] == in[j][idx] {
					diff++
					common += string(in[i][idx])
				}
			}
			if diff > longest {
				shortestCommon = common
				longest = diff
			}
		}
	}

	return shortestCommon
}

func Checksum(in []string) int {
	twos, threes := 0, 0
	for _, id := range in {
		counts := map[byte]int{}
		for i := 0; i < len(id); i++ {
			counts[id[i]]++
		}
		foundTwos := false
		foundThrees := false
		for _, count := range counts {
			switch count {
			case 2:
				if !foundTwos {
					twos++
					foundTwos = true
				}
			case 3:
				if !foundThrees {
					threes++
					foundThrees = true
				}
			}
		}
	}
	return twos * threes
}

const puzzleInput = `bpacnmelhhzpygfsjoxtvkwuor
biacnmelnizqygfsjoctvkwudr
bpaccmllhizyygfsjoxtvkwudr
rpacnmelhizqsufsjoxtvkwudr
bfacnmelhizqygfsjoxtvwwudp
bpacnmelhizqynfsjodtvkyudr
bpafnmelhizqpgfsjjxtvkwudr
bpackmelhizcygfsjoxtvkwudo
bmacnmilhizqygfsjoltvkwudr
bpafnmelhizuygfsjoxtvkwsdr
boacnmylhizqygfsjoxtvxwudr
bpbcjmelhizqygfsjoxtgkwudr
bpacnmglhizqygfsjixtlkwudr
bpacnmclhizqygfsjoxtvkwtqr
bpacnmelhczqygtsjoptvkwudr
bpacnmelhizqywfsaoxtvkbudr
apacnmelhizqygcsjoxtvkwhdr
bpacnmelrizqygfsbpxtvkwudr
tpkcnmelpizqygfsjoxtvkwudr
bpacnmelhizqlgfsjobtmkwudr
npacnmelhizqygffjoxtvkwudf
bpacnmeehqzqygqsjoxtvkwudr
bpecnmelhizqigfsjvxtvkwudr
bpacnmelhizqysfsjoxtvkdfdr
bpacnfelhkzqygfsjoxtvkwfdr
bpacnbelvizqygfsjoxthkwudr
bpacnoelhizqygfejoxtvkwudn
bpacnmelhizqygfzpkxtvkwudr
bpahnmelhizqyufsjoxmvkwudr
bpacnmelhizqygfsnoxtvkwmmr
bpacnmelhizqygfsjoatvkludf
bpacnmylhizqygfsjlxtvksudr
bpacnmekhpzqygysjoxtvkwudr
bpacnselhizqogfswoxtvkwudr
bpacnmelhizqprfsjoxwvkwudr
bpatnmelhinqygfsjoctvkwudr
bpacnqelhqzqygfsxoxtvkwudr
bpabnmelhiyqygfsjoxtykwudr
bpacnivlhizqygfsjoxtviwudr
bpkcnmylhizqygfsjoxtvkwcdr
bpafnmflhizqygtsjoxtvkwudr
bpachmelhizqygfsjixtvkwudg
bpacymelhizqygfsjoxtykwuar
bpacnkelhizqdgfsjoxtskwudr
bpacnmezhizqggbsjoxtvkwudr
bpacnmqlhizqygrsjoxzvkwudr
bpaczmelhizqyhfsjoxfvkwudr
bdacnmelhyzqygusjoxtvkwudr
bpacbmelhizqywfsjostvkwudr
bpacnmelhihzygfstoxtvkwudr
bpactmelhizqygfsjcxtvkwydr
bkacnmethizqytfsjoxtvkwudr
bpacnmalhizqydfskoxtvkwudr
spacnmelbizqygfsjoxdvkwudr
lpalnmelhizoygfsjoxtvkwudr
bpacjmeghizqygfsjoxtviwudr
bpacnmeqhizxygfsjoxgvkwudr
bpacnmelhizqygosjoxtvkkuhr
bpacnmelhiznbxfsjoxtvkwudr
bgacnmelhizqygfsjbxivkwudr
bpacnmelhizqygfjjowtvswudr
bpacnmelhizqygfsjovtgkmudr
bpacnmelcmzqygfspoxtvkwudr
bpvcnmelhizqyvfcjoxtvkwudr
bpacnmeahizqjgfsjoxtvkwukr
bpacnoelwizqygfsjoxtvkaudr
xpacnmelhizqygfsjoxdvkwedr
mpacnmelqizqygfsjoxtvkwudx
bppcnmelhizqygfsjfxtvkhudr
bpacnmclhizqyhfsjaxtvkwudr
opacsmelhizqygfsjmxtvkwudr
bpafnmelhizqjgfsjoxtvkrudr
bpdcnmilhizqygfsjoxtvkludr
bpainmelhizqygfsjtntvkwudr
bradnmelhizqygfsjextvkwudr
bpacnmelhizqygfmsoxtvkwudg
bpacneelhizqygvrjoxtvkwudr
bpacnpelhizqygfsjoxyvkwudf
bpacnmelhizqygfsqoqtvkwodr
bpacnmelhizjyghsjoxcvkwudr
bpacnmelmibqygfsjoxtvnwudr
jpacnmelaizqygfwjoxtvkwudr
zpachmelhizqygfsjsxtvkwudr
bpacnmelfizqykfsjomtvkwudr
bpacnmllwizqygfsjoxtvkwusr
bpaynmelhizqygfsjoxtvowcdr
jpacnmqlhizqygfsjoxtvknudr
bpacxmelhizqyffsjoxtvkwugr
apawnmelhizqygfsjtxtvkwudr
mpacnmelhitqigfsjoxtvkwudr
bpacnmelhhzqygfsjoxtvkyzdr
gpacnmelhizqynfsjoxtvkwudm
bnacnkelhizqygfsjoxtpkwudr
bpacnmelfizqygfsumxtvkwudr
bpacnmelhisqygfsjohtvowudr
bpacnmelhimqygxsjoxtvkwudn
bpscnmeliizqygfsjoxtvkwunr
qpacnmelhizqycfsjoxtvkwndr
bpacnmelhijqygfsjohtvkyudr
bpacnmelhizqykfsjkxtvknudr
bpacnqilhizqygfsjoxtvkoudr
bpacnmelhizqzgmsjoxtvkwurr
bpdcnmelhizqygfsjoutukwudr
bpecnmeghizqygfsjoxgvkwudr
bpicnmelhizqygfrjoxtvlwudr
bpacnmelhizfygfsroxtvkwodr
buacnmelhizqygjsjoxtvkvudr
bpacnmelhixqykfsjoxtvrwudr
bpacnmelhizqygvejcxtvkwudr
bpacnmjlhizqylfsjoxtvkwuor
qpacnmelhizqygfsjoxfdkwudr
bpfcnmemhizqygfsjoxtvknudr
bpacnmelhizqoffsjqxtvkwudr
hpacnielhiqqygfsjoxtvkwudr
gpacnmelhizqygfsewxtvkwudr
bpacnmellizqylxsjoxtvkwudr
bpacnmenhizqymfsjoxtvkmudr
bpacnfelhizqygcsjoltvkwudr
bpacnmelhqqqygfsjoxtvkuudr
bplgnmelhiqqygfsjoxtvkwudr
bpacnzelhizqygfgjoxtvnwudr
bpacnmelhizqygfsjoktvknunr
bpacnmdlhioqygfnjoxtvkwudr
epacnmelwizqyjfsjoxtvkwudr
bpacxmelhazfygfsjoxtvkwudr
bpacnmejhezqygfsjoxtskwudr
bpacnqelhihqyzfsjoxtvkwudr
bpacnbelhizqyrfsjoxtvkmudr
bpacnmelhizqygfsjoxtylwzdr
bpacnmelwizqygfsjodtvkhudr
bpacnnelhizqygfsjoxtwkwadr
bpacimelhizqygfsnoxtvkwuor
bpacnmelhizqyaasjoxtlkwudr
bpacnmelhizqyeffjoxtvkwuds
bpacnmenhizqygxscoxtvkwudr
bpacnmelhidqygfsjowtskwudr
bpacnmeliizqygfsjoxhvkwucr
bpacimelhizqygfsjoxtvktuwr
bpainmelhhzqygfsjzxtvkwudr
bpacamelhizqygfsjogtvkwbdr
bpccnmelgizqygfsjoxtykwudr
bpacnmelhizwegfsjoxtvkwadr
bpackmelhbzqygqsjoxtvkwudr
bpacymeihizqyffsjoxtvkwudr
bpacnielhczqygfsjoxtvkwudk
bpacnmejhizqygffjoxjvkwudr
ppacnmelhizqygfsjoxtigwudr
bpjcnmolhizqygfsjoxtvkwndr
bpacnmelcizqygrsjoxtakwudr
cpawnmelhizqygfsjoxmvkwudr
bwacnmelhizqygesjoxtakwudr
bpacnmelhizqygfsjexsvkwddr
bpaunmelhiuqygfsjoxtvkwtdr
bpacnmellimqygfsjextvkwudr
bpacnmerhizqygfsaoxvvkwudr
bpacnmglhizqygfsjixtukwudr
ppacnmelhizqygfsjoxtvkdudp
bpacnmedhizqygukjoxtvkwudr
bpccnmelhizqngfsjoxtvkwadr
bgacnmeldizqygfscoxtvkwudr
bpacngelhizsygfsjoxtvkwkdr
bpacnpelhizqygfsjoxctkwudr
bpacnmylhizqygfcjoxtvkwmdr
npacnmelhizqygfsjoxtwkwuds
bpaxnmelhizqydfsjoxyvkwudr
bpacnhelhizjygfsjoxtvkmudr
bpacnkelhczqygfnjoxtvkwudr
bfacnmelhizrygfsjoxtvkwodr
bpycnmelhizqygfofoxtvkwudr
qpacpselhizqygfsjoxtvkwudr
bpvcnmelhezqygfsjoxttkwudr
bpacnmwlhizqygfijoxtmkwudr
bsacnmelhikqygfsjoxttkwudr
bpccnxelhizqyafsjoxtvkwudr
bpacnmelhizqygfswhxtvewudr
vpacnmzlhizqygfsvoxtvkwudr
bpacnmelhihqygfsjoxtvkqurr
bpacnmelhixqygazjoxtvkwudr
bpavnmelhizqygfsjozpvkwudr
bpacnmclhizuygfsjoxmvkwudr
bpacnmelhizryufsjoxtkkwudr
bpacnmelhtzqygfsjobtvkwufr
bpacnmelhizqmlfsjoxtvkwudq
bpaaneelhizqygfsjlxtvkwudr
bpacnmelhxzqygfsjoxthkwuhr
bpacnmeshizqygfcjoxtvkwude
bpacnzqlhizqygfsxoxtvkwudr
bgaanmelhizqycfsjoxtvkwudr
bpacnmexhizqygfsroxtvkwudn
bpmmnmelhizqygfajoxtvkwudr
bpacnmelhizqylfsjoxtckwhdr
bpicnmelhizqyrfsjoxtvkwudi
zpacnmelhizvycfsjoxtvkwudr
bpamnmkllizqygfsjoxtvkwudr
bpacnmelhrzqyrfsjoxgvkwudr
bpadnmelhczqygfsjoxtlkwudr
bpacrmelhizqygrsjoxtvkiudr
lpacnmelhizqygfsjoxtgkwxdr
fpacnmalhiuqygfsjoxtvkwudr
bpacnmelhizqygfsjixtvfwcdr
bpccnmelhxzqygfkjoxtvkwudr
bpacnmepaizqygfsjoctvkwudr
tpacnmelhivqygfsxoxtvkwudr
kpacnfelhitqygfsjoxtvkwudr
baacnzelhizqygfsjoxtvkwudx
bcycnmeghizqygfsjoxtvkwudr
wpacotelhizqygfsjoxtvkwudr
bpacnmsshizqygrsjoxtvkwudr
blacnmelhizqygfsjoxtykwvdr
bkacnmelhizqygfsjoxuvkludr
bpacnmelhizaugfsjoxtvhwudr
fpavnmelhizqygfsgoxtvkwudr
bpachmelnizqygfsjextvkwudr
bpacnmelhizqpgfsjoxtvkwldu
bpacnmelhizqygfsloftvywudr
bpacntelhvzqygfejoxtvkwudr
bpacnmeldizqygfsjmxtvkdudr
byacnmelhizqygfsjsxtvkwudh
bpacnmellizqygssxoxtvkwudr
bpacnmelhizqygfsjootvknuir
bpacnmelhitqjgfsjoxivkwudr
bpacnmelhazaygfsjoxtvfwudr
bpacnzenhizqygfsjzxtvkwudr
bpacnmelhizqypfsdoxtvkwuar
bpannmelhizqygnsjoxtvkwndr
bracnmeldizsygfsjoxtvkwudr
bpacnmelhizwygfsjugtvkwudr
bpatnmelhizqygfsjoytvkwulr
upacnmelhizqygfsjurtvkwudr
bpaenmezhizqygfsjostvkwudr
bpacnmelhizpygfsjodhvkwudr
bpacnmelhizqygfsjogtvkguwr
bpacnmelhisqygfsjoxtpkuudr
bxacnmelhizqygfsjdxtvkfudr
bpacnmelhizqygfsjohqvkwudu
bzacnmtlhizqygfsjoxsvkwudr
bpacnmplhixrygfsjoxtvkwudr
bpacnmelhizqhgfsjomtvkwudg
bpacnmezhizqygfsjxxtykwudr
bpacnmwlhizqygfujoxtzkwudr
tpacnmelhizqygfsjoxkvpwudr
bpawsmenhizqygfsjoxtvkwudr
bpacnmelhizqtgfsjoxttkwuqr
bpkcbmelhizqygfsjoxtvkwucr
bpacfmekhizqygfsjoxtvkwuds
bpacnmethizqynfajoxtvkwudr
bpocnmclhizqygfsjoxtvkwukr
zpacnmwlhizqygfsjoxzvkwudr
bpacpoelhqzqygfsjoxtvkwudr
bpacnlelhizqyzfsjoxtvkwukr`
