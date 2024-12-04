package main

const day04Input = `MXMXMAMXAXMSMMMSSSXSSSMXXMSSXSMSASXSMXSAMXSMMMSMSMMSSXMSMSMSMMMSSMAAXSAMXSMSSMASAMMMSAMXSSSSSSSSXSAXMMMAXMMSSXMASXXSAXSXXAMXSSMSXMSAMXMAMAMX
SMMAMXMSAMXAAAAAXMAMXAMAASAMSXXXAXAXAASAMAAAAAAXMAAAMAMXAAMAAAAAAMMXMSXMASAAAMXMAMAMSAMXXAMAAAAAXMAMSASMXSAASASAXMAMXMASMASASAMXASXMSMSMSSSM
AMSMSMXXAAXXSMSSMMAMSXMASMASAMXMMMMMMMXAMASMMSSSSMMSSMMMSMSSSMMSSMXAASXMASMSMMASMMXXXMASMAMSMMMMMMXMSASMAMMSSXMSXMXMAMAMSAMXSAMSXMAMAAAAAAAA
MMMAAAMMSSSMXAXMAXAMXMXSAMXMAMSAAAAASMSSMMAAAAAAAMAMAMAAAAXMAMAMXMSMXSASMXXMASASAMXMASASMSMMAXXXMASMMMMMAASAXXXXMSAMASAXMXSXMAXMASXXASMSMMMS
XMMSMMAXAAAASAMASMMSMSMMMSSSXMAMXMMMSAAXAMXMMMSMMMASAMMMSSMXXMXXAXAMXSAMXXASAMXSAMAMMMMSAMASMMSMXMAXAAXMASMMMMSSMSASASMSMAMMAMMMAAMSMMAMXXXX
MXMXMASMMSMMMAMAMAAXAAAMAAMXMSSSSXSAXMASXMAXXAXXXXXMMSSXMAXAASMSMSASAMAMAMXMMMXSXSAMMMAMAMAMMAAASMMSSSSMXMASMMAAXSAMXSAXMAMAAXSAMXAAXMAMXXMS
AAMXMMXAXMAMSAMXSMMMSMSMMMSAMMXAAASAMXMMASXSMMSXASAMXAMASAMXMMAAXXMMXSAMXSAAASMMMSMMAMASMMASMSSSXAAAAAXMASAMAMSSMMXMXMXMMSSSSSMXSMSMMSAMMXMX
SXSAMSSSMSXMSASAXAXXAXXXMASASXMSMXMAXXAXXMMSAASMMMSSMSXMMXMMMMSMSMSSMSASXSMSMSASMSMMMSAMAXAAXAMMXMMMMMMSAMXSXMMAXSAMXMASXMAXXXXAXXAMXSSXSAMA
MAMMSAAMAXSAMXMASMXSAMXMMASMMMAMXSMSMSMMSAASMMMAXMAMXAMSASMSAAMAAMAAXMAMAMAMXSAMASAAXMXSAMMMMMSSSXXMASAXSSMMAMSAMXAMXSAMAMXMMSMSSSMMXMAMSASM
XAMXMMSMSMXMASXAXXAMASXAXAXASMSMSAAAAXMASMMMXSSSMMASXMASAMASMSMXMMSSMMXMAMMXMMMMAMXMSAMXXSAXMAAAXMSXMAASASAXAMMASMMMMMMXXSAAAAAAMXSAAXSMSXMM
XMXXXXAAXXXMAMMMSMXSAMASMMSASAMXSMXMXMSMMASXAMAMXSASMXAMAMXMXAXXMAMAXAXSSMSXAAAMXSXMMMMMASASAMMSMMMAXXMMASXMSXSAMAAXMASAAAMSSMSMMAMMSMXASAMX
XSASMSMSMSMMAMAMXAAMASAMSXAXMAMAMXMAXXXMMAMMMSAMXMASMMXMMSSXSSSXMAMAMMMMAAXMMXXSAMASAMXSASAMMAXMXSAMMSXMAMAXXAMXSSMMSAMMSMXXMXMXMASAMAMXMAXX
XAAMAAAAAXXMASMSMMMSAMASMMSXSAMAMAASXMAMMAMAXSMSXMXMMMMSAMXAXAMXSSXSASASMMMMSSMMASAMSXXMAMMMXSMSASASAXAMXSMMMXMAMASAMXSXXXMMAAMAXAMMXXAMSSMS
SMSMSSSMSMSMASMAAAAMASXMMAMASASXXXMXASXMSAMSAMASMMMAAAXMAMMAXXMAXXAMXMASXAXMAAXSAMXXMSXMXMAMAMAMASXMXSXSMMXAAAMMSAMXSXSAMAMAMSSXMMSSSMSXMMAA
XXXAAAMXMAAMSMMXMMXSMMAAMXSASMMMSSSSXMMXMAMMAMXXAASMMXSSXMAMAXMXSMMMAMXMMASMSSMXXMXAMXSAMSMSSMSMMMMMAMMMAXSMSXSAMMMXSMSAMAXSMAXAXSAAAAXAAMSM
SSMMMMMAMMMMXAASAMXXXSMMMMMMSXXXAAXMAXMAXSMXSMXMMMSAAAMXMMSSMMXMAXXMAXMAMXAXMAMMSMMASAMXASXAMAMMAAAXAAAMSMSXMAXMXMSXMASAMXMXMASMMMMSMMMSXMMX
SAMMSMSMSSSMMSMXMAAMXMSMAAMAMXMMMSMMMMSSSXSAXMMMAAMMMMXAXAAAAXMSAMMSASAXSSMXSMSAAXAMMMMXSXMXXAAXXXMSMSMMAAXAMMMSAMXAMXMXMXAXMASXSXXXAMAXAXMM
SMMXAAAAXAASAMXASMSMMAASMSSSXAASAAMAAAXXAAMXMAASMMSAMSSMMMSSMMMMAMAMAXMAAAMMSAMMMSSXAMSMMSSXSSSSMMXAMXXXMXMXAAASASXXMXSAMXSAMXSAMXXSAMASMMSS
MXSSMMMSMSMMXSSMSAAXMSMSXAAXMSSMMSSSSSMMMMMAMSMSAASAMXAXMXAAAAXSAMMSSMXMSXSAMMMMMSASXMAAMAMMAAAAASXMSXMASXMSSMMSAMXAAAXASAMASAMAMMMSAMMXMAAA
XAXAAASXMXAMSAMXMMMMMAAXMMMMAMAXAMMAAAMMXAXAXXAMMMSAMXMMMMMSMMMSASAAAMXAMXMMMXAMXSASMSXSMASAMMMMMMAXSAMMMAAXXSAMXMXAMXXXMASAMMSSMMAXSMAXMMSS
MASMMMSAMXAMMMSAMXXAMMMMXSAMXSAMSSMMSMMASMSSSMSMXXXMASXAMAAXASASXMMXSMASXSSSMSMMAMAMXSAAXXSMXXSSSSSMMSSMSMMMSMXXAXSSSMSSSMMMSXAAXMMSMMMMXMAX
SASAMXSAMSSMAASAMSASXSMXAMXMXMMAAXMMMMMXMAAXMAXASXMMXXSSSMMSAMAMMASAXXAAASAAMASMMMMMAMAMAXSXSXSAMAMAMAXMASASAMMSSMAASMAAASXSSMXSXXXAXAAMXMMM
MASXMASAMAMSMMSXMXAMASAMMMSMMMMMMMXAAAMAMMMMMSMAMAASXMAXAAAMXMAMMAMMSMXMXMSMMASXXAXMASASMMMAMSMAMXMAMSSSXXXMAXMAAMMMMMMSMMSAMXMMXMSXXSSXASAM
MAMAMMSSMAMSMMXMMMSMAMMXMASAMXAASXSMSSSMXMXMMAMASMMMAAMSMMMMASMSMXSMXMAMSMXSMXSXSXMSXSAMAAMXMASXMXSXMXAMAMSSXMMSMMXXAXXAXMMXMAAMAXASMMMMAXAS
MSXSMXMXSSMSAMSAXSAMXMXXMASAMSSXSAXAAMAXXAMXXXSXSAMXXSMAAMXSASMAMSXXAXAMXAAMMMMXXAMSAMXSSMSMXAXMAMAMXMAMXMMAXXAXMSSMSMSMSMMMSMSSMMASXAAAMSMM
XMAMXSMAMXASAMSMMSASMXSXMASMMMMAMMMMMSXMASXXSASXMASMMMXSSMMAASXSMSASMSSMAMMSAASMMMMSAMXAMAMMMSSSMSMMXSMMXMXAMMSMAAAAAAAMAMAXAMMAAMSMXMSMMAAA
XMAXAXMASMMMSMSSXSAMXXMAXAMMAXMXAXMSMAMAXMAAMAMAMAAAAAAMAMAMXSAMAMAMMAXXAMAXXMXAAXASAMXMXMXAMMAXASAMXSAXAMAXMXAMMMSSMSMXAMMMAXSSMMAAAXAAMSSM
XMXMSMSMSAMAMXMXMSASMSSMMSSSMSMSMMMAXMASMMMMMMSAMSSSMSXXAXMXAMAMXMAXMAXSSSXSSSMXSMXSAMMXAMXSMMAMXMAXASASASMSSXXSAAXAMXXMSMMSMMXMMMMXMSSSMAAA
XAMSAAAXSXMASAMAAMAMAMXXAAMAAAXAASXXSMMXAAXXMAMXXXMAXAMSMSMMASAMXMSSMMMAAMASAAMAMXMXAMAMASMAXMASXXXMXSAAAXMAXASXMMSMMASXMAAAASXMMMSAAXAMXMMM
MSMXMMMXSMSASASAXMAMAMMMMMSMSMSSSMMSMAMSXMXAMXXSAMXXXMXAAAXAMMXSXAAAAXMMMMSMSMMAMASXMMMXAMMMMMSMMSMMMMMMMMMSSMMSAMMAMASASAMSSMASAAAMXMMMMXXX
AMAMSASMXAMXXAMASXASASASMXAXAAXXAAMMMAMXAMXSSSMMASXMMXSMSMMMASAMMMSSMXXAXXXAMXMASMMAASXSMMSXMXAAMAAAXSAMXAAXXXAXMMXAMAMMMMAMXXAMMSMMAAAXXXMS
SMMASASMMMMSMAMAMMMSXMXAXMAXMSMSSMMSSSSSSMAXAAASAMAAMXXAAAMMMMASAAMAMMSMSMMMMXSAXXMXMMXAAAMASMSSMSSMSMMSSMSASMMSASMMMXMASXMASMSAAMASXSXMASXS
XAXMMXMASAAAAAMMSMAMMMMMMSSMXAAXAAXMAMAAXMMMXMMMASXSMMMMMXSAXSSMMSSSMAMMAAASAXMASXMSASXSMMMAMXAXXMAMXXAXXXMAAMAMMSAXAXMASXAMAMXMXSAAAMASXMAS
SMMXXAMASXSSSMXAAMMSAXAXMAAXSXMSMMMMAMMMMAAAASASAMAAXAAXAAXMXXMSXXAXMASMSSMSXSMAMAASAMXMAMMSSSMSXSAXSMMSXXMMMAAXXSXMSMSMSAXSXSMSAMMSASMSAMAS
XMAMMSMASMMAXAMSSSXSXSXMMSSMMMXSAMASXSAMSSXSASASMMSMMSSMMMSXMASMXXMSMMSAMMAMXMMSXMMMAMAMAMXAAAAAASMMMAMMMMMMXSMSAMXMAAAAMAMAAMASMSAXXSXSAMAS
MMASAXMXSAMAMMXMAMAMASMXAXAAXMAMAAASXSAXXAAMAMXMAAXAAAXMAMXAXSAMMSXAMSMMMMSAMXAXAXXXXSMSSSMMSMXMXMMXXAMAAMMSMAMSASAMMMMMMAAXXMAMXMASAMASMMAS
ASASMXXAMXMAMXXMXMSMAMAMXSXMMMXSMMMXMSMMSMXMAMAMMSSMMSMSSSSMMXMXAMMMXXAXAAAAMMASMMSMMAXXMAXXAMSMMAAASMSSXSAAXSAXAMXXMXSASMSSXMXXAMXMAMAMAMAS
MMAXXXMMSMSMSMMXXXAMAXMAMMAXAAXAXXXAAMAAXXSMMXXXAMMAMAASXAAMXMXMMSAXMSSMMSSMMMAAAXAASMMMSMMSASMASMMMSAAAAMMSMMXMSMSMSASASAAAXXSSMSSSXMASXMAM
XMASMMSXAXAMAAAMMMMSASMXAXAMXSAMXMSASXMMSMSAMASMMSXAMMSSMXMMXMASASMSXAAAXMAMXMXSAMSXMAAAAAAMXMMMMXXXMMMMMMXMAMXSAAAAMAMAMMMSMXAAAXXAMSXMXMAS
XMMSAAAAMSMSSMMAAAXMXSASMSAMXAMSAAMAXAXXAAXAMAMAAXXMSXXMAXSAXSAMXSAMXXSMMXAMXMAXAMMMSSMSSSXMASAMMXAMXASMMSASMMXAMMMSMAMXMMAAMASMMMSSMSAMASAS
XMAMMMMSXMMAMASXMSXMXMAMAAXMASAXAXSAXAMSXSMMMSMSMMASAXMMSMXAMMSSMMMMMXMASMASXMASAMXAAXMAXMAMAMAXMMMMSMSMASXXAMXXXAXAMXSAXMXMSAMXSAAAXSAMXXMS
AMASAAXXAMAXMAMAAXMMSMSMMMSXXMXMSMMMXAAXXXASMMAMXAXSASMAMAMXMAXXASAAXAXAMSXMXMASAMMMSSMSMSXMSSSMMSAAAASMMMMXAMMMSAMXMMMMSMSMMASAMMSMMMAMXAXS
MSXSMSXSAMMMMSSMMMAAAMAXXXMMXMXAAAAMSSSMMSAMAMAMXSXXAMMASAMMMXSAMSSXMAXSMMXSXMASMMMAMMAXMMXXXAAAASMSMXSAXAMXXSAAXMSXSAXXAAAXSAMXSXMASXSMSMMX
MMAMXMASXMXAAAAAASMMXSMMXXSAAXAMSXMMAAAXAMMMAMAMMMAMMMXAXXXAMAXXXXAMMSMXAMMMAXAMXMMXSMAMASMMMSMMMSMXXSMXMMSAASMSMAAAXMXSMMMMXXMASAXXMAAMAMSM
MMAMMMAMMXSMSSMMMXXAASXSAASMSXMMMXMMMMMMSSMSMSXSAMXXAAMXSSXSMSSMSMAMAMASMMSSMMXMAMSAMMXXAMAAXMSAMXMMMSMASAMMMMAXAMMMMXMXAAMAMXMMSAMSMSMMAAMA
XMASAMAMMMMMMMMMMMMMMSAMMXMMMMSAAAMASAXMMAASAAAXAXSMSSSXAMAXAXAMASAMXSMXMAXAXMASMXASXASMMSSSMMXMXXMAAXXAMXMMASXXSXSMSSMSSXSAMXMMMMMSAMXSMSXS
SSMSXSAXXAAXAAAMASAAAMAMMXXMAAMMSMSASXSXMMMMMMSSSMSAXAMMAMMMMMMXAMMSMMMAMSMXMSASXSMAMAXAAAAAASMSMASMSSMMXSASAXMAAMMMAAAAAMXAMXAAAAAMAMAXAMMM
XAMMXSASMSMSMSMSASXMXSMSSMAMMMSAXXMMSAMXMSXSAXMAMAMMMAMSXSMXMASMMSAAAAMAMMXSXMAXAAXMAMXMMXMMMMAAAXXMAAAAAMXMXSMSMSAMSMMMSMSSMMMSSMSMSMSMMMMM
SMMMAXMAMMMAXMAMMSXSASAAASXMAXMXMMMAMAMXMXAMAMMAMAXMSXMMMMMASXSAXMXSSMSSSMAMAMAMXMMXSAASMSSXXSSMMMMMMXMMSMXMAMAAXMXMAMXMAMAMXAAXXAXMMAMAXSAS
AXXMAMAMASMMSMXMXMAAAXMSMMMMMMMXMASMSASAMMSMAMSXSXSMXMXMAAXMXMMMSMAMAAAAAMAMSMMSMASAAMMSAAXXXMAMXSXAMXSXXMAMASMSSSSSMMASMMSSSMXSMMMXMAMMASAS
SMSMSMSMASAXMMMSAMXMMMXXAAAASASASXSXSASMSAMMAMXAXXSXAAASMSSSSXMAMMAMSMASXMASAAAAMAMXSXAMMMSXMSAMASMXSAMAMXAMXSXAAAAMMXAXAAAMXSASAMASMXSMMMAM
XAXAAAAMSSMMXAAXAMXAXXAXSMSMSASASASXMAMXMASXMXXXAMMASMMXAXAMMAMSXSAXXXXAAXSAXMMSMMMAXAMXAXMAXSXMASMMMMSAXSXSXMMMMMMMMMSSMMMSAMXSXMASMAAXXMAM
MAMMMSMSAMAMSMMSAMSSMMMXAMXMMMMMMAMAXASXMXMAMSAMXAAAMXMMSMSMSAMAAXMASXXSMMMMXXXXXMSSMMXXMMMSMMASMXXXAXMAMSAMXXAXAAMMMAMAAMAMASXSMSSXMSMMXMAS
AAXAAMXMAMAMAAAMXXAAAAXXSMXAMXSMSSSMMAMXXMSAMXMMASMXSAMXMAMXSSSMMSXMMAAXMAAMXSMMMXXMASMASAMAAMAMXAMMSMMSSXAMMSXXSMSAMXSSMMASXMASAXMAXMXAXSMX
SMXSSSMSAMSSMSMSMMSSMSXSASAMXAXMAAAAMAXSSMMAMMSMXMXASMXMMAMXMAXXASMSSMMMSSSMASAAMSASXMXAXASXMMASMAMAMXAXAMAMXSAAXASMSAMXMSMSAMAMAMSMMXMMMASX
ASXXAXAMXMXAMXMAAAMXAAMSAMXXMXMMMXMXMMMXAMSSMAAXXMMASMSMMAMAMXMMMSAXAXAMAXAMASASXSMMAXMASMMMXSASMSMASMMMMSMXAMMMMAMXMASXMAASMMMSMXAMMAAXSAMS
SXSMAMMXXMSMMMXMMSAMMMXMMMMMMAASAMXXMXSSMMAMMSMXAXMXMAAMSSSMSAXXXMMMSSXSSSMMMSXMXSASMMSMXXASXSAXAAMXMAXAXAXMSSMXMAMAXAMMSMXMMAAAXSASXXSXMAMX
XAAMAMXSMMMMAMXSAMAXAXAAAAAAMXMMASMAXSMAXMASMAASMXMXMSMMAAAASMMAXAMXMSAXAXAAXMASASXMAMASAMXMMMAMSMSAMXAMSMSXMAMSSSSXSAMXSMSASMSAXSMSAAMASAMM
MXMAXSAMAAAMASXMASMMMSMXSSSSMAXSAMMSMAXAMMAXXMAXAAMXMXXMMSMMXSSSSXSAAXMMMSSMSMMMMMMSMSMMXXSMXMSMMAXXXAMXAXAMSAMXAAAMSAMXSASXMAMMMMASMAMAMASA
MAXXXMAXSXSSMSAMXMAXMAMAAXMAXMMXXMAMSMMMSXSSXSXXSSMMSSMXMAXAXMMAAASMMMXMXMAAAXXMAAAAXMXSMMMSAMXMMMMSXSXSMSMMXSMMSMXMSSMAMAMXSMXAXMAMMXMASXMM
SSSMMMSMMAAMMMMMAXMXSASMSMSASMMXSASXSMAXXAMMAMMXMAAXAMMAXMMSMSAMMMMMAMAMASMMMMMSSMSXSMAAAMMAMXXAMXXXAMAAXAAMAMMAXXXXMAMXSXMAMXXMMMMSMSMMXXAA
XAXMSAMAMMMSXMASMSXXMASAMAMASASMSAMMMMSSMAMMAMMAMSMMSSSMSSXMAXAMXXASMSMSASXMAXXAMXAASMSSSMMSSSMAMXAMAMSMMSSMAXMASXMAXMMMMAMXMAAXAMXAMAAMSMMM
MXMAMASMMSAAASMMMSMXMAMXMAMASMAAMAMAMAAAMAMSAMSAXAXXAAAMAMMMSMAAXXXSXAAMASMSSSMMSMMSMAAAMAMMAMXAMMMSAMXAXAXXMMMASAAXXSASMSMMSSXSASMSSSXMAAAM
MSMMXMMAAMMMXMMAAMXAMAMAMXMAXXMMSXMSMMSSMXXMAXSMMASXMSMMASAAAMMMSMSMMMMMMMXMAAXAAASXMMMMMAMMAMSASXAMXXAXMMSXMAMSSMSAMSASAXAXAAAMAMAXMAXSSSSS
AAAXSXMMMMMMSASMMXMSMXMAMAMSSMXMSMAAMXAMMXMMSMMMSAMMMMXSASMMSSMAAXAASXXSASAMXMMSSMMAXXMSMMXSAMXXAMXMSAMXAXMASXMMSMMAAMAMAMSMMSMSSMXMMAMMAMAM
SSSMAXXSXSXASAMXSSMMASXMMAMAAAXAMMSMSXSSXSAAMAAMMSMAXMASAXASXAMSSXXAMXASASAXSAMXMASMMMAMMMMMMSXASXSXAAAXAMXMXSAAXXXMMMXMAMXAXXAAXXXXMASMAMAM
XAAXMXAXAAMMSMSAAAAMAMASXSMSSMMXSXMXAMXXAXMSSMMSAMMSMAMMSMXAMMMMAMXSAMMMXMAMMXMAMMASAMAXSXMAAXXMMAMXSMMXSMSXAXMMMMXMASMSMSXMMMMMSSMMMXXMXMAS
SSSMXMSAMXMXMAMXSMMMAMAMXAAAAXAXSASXXMSMMMXXXAXMMMAAAXAAXXXMSXMSMSMXASXMSMAMXSMSMSASXMASXASMMSMMMSMMXMXAMXAMMSMSASXSMSAASXXAAXMSAXAAMSMXMAXX
XMAXXAMASAAAMMMMXMMMSMSSSMMSMMXXMAMXMAAAXXSXMMSMSMSSSXMMMSXXAAMAAAAMXMAAASASXMAMAMAMAMXAXMMAASMAAMXSASMAXASAMAAMAMMAMXXXSXMSASXMASXMMSAXSAMX
XSAMXMSAMXSASXSSXSAAMAAAXXAXXSXXMXMAAMSXSXMAMSAAAMAMXXMAMMMSSSMMSMSMXSMMMXXSAMXMAMMMSAMXSXSMMMMMMSASASMSAMXAXMSMMMSMMAMMMMAMAMXMXMMMXMAMAAMS
MMXAMAMMMMXXXAMXAXMMMMMMMMXSAMASMMXMSXMAXMXAXXMMMSXMSMSASMAMAXMXMAXMXXMXXMXMXMAMXMAXMAMAMXSAMSMSAMASAXAMXMMSMXAAXAMXMAMAAAMMAMXXMAXSAMMMSMMX
AAASXXXAAMMSMAMMMMSASXMSXSAMAMAAAXSXXAMAMMMMXXXSMMXAAAMASMSMAMMSMAMXSAMSSMMSMSSSMXMMSAMXSASAMAAMXSAMXMXMASAXXSMSMAXAMSSSSSXSSSXSXMXMXSXAXXXM
XSXXAAMSMSAAMSMSXAMAMAAAXMASXMSSSMAAMMSMMXASAXAMAASXMSMSXMAMSSXMMSMXMASAAAAAMAMXAASXSSMAMASXMMSMAAAXMMXXXMASAXAXMMMMMMAXAMAXAMMXAMXSXXMSMMMS
AXMMMMMAAMMSMMAMMSSSSMMMMXAMXXMAXAMXMAAAAXAMXXMSMMMASXXMAMXMAMXMAXXXSMMMSMMSMSXMXXSAMXXAMAMAXMAMSSMMASMSMSMSAMXMASASASAMMMMMAMAMAMXAAMAMASAS
AMAAMASMSMAMMMAMXMAXAXXXAMSSMMMMXSXMMSSSMMMSMMXAMASMMXSSMMAMMSXMASMXMAAXXAXXAXMSXXMASMSXSASMMXMXMAMSXXAAXXASXAXSXMASAXXXXAAXSMSAXXXXMMASXMAS
XSSMSASXAMASASAMAMXMSMMMMSXAXMAMMMSAAMAXXAAAAXSASXSXMAMAASXXXSXSXXAASMMMMSXMAMAAXMMMAMSXSMAMAASMSAMSMMMMSXAMMSMMAMMMMMMMMSMSAAMMSMSMXMXXXMAM
XMXAMASMMMMSASMSAXAAASXAXMMMMSSMSAAMXMASXMSXXMMMMAMMMMXSMMXSASMSAXSAMXAXAMMAAMMMXSASAXMAXMAMSASASXMXMSAAAMMMXAASAMSAXXAAAMMAMXMAAAXAXMMMXAXM
XSAMMMMMXSMMMMASASXSMAXSMMAXMXMAMXSSSMASXMXMSSSXMAMXAMXASAXMAMAXAMXSSXMMASXSXSMXSAASASXSSSSMMAMAMMSAASMSXMASXMXSMSMAMSSMSXSXXAMSMXSXMXAMMMSM
MAMXASAMXAAAMMMMAMXAXMAMASASAAMXMAXAAMXMASXAAAMASXXSAMMAMMSMAMXMSMXXMXMSXMAMAMXXAMXMXMMAAAMAXXMXMXSMMMXAAXASASAXXXXAMXMAXXAAMAMAMXMASXMSAAAX
MASXXSASXSMMSASMMMSXMXXSAMASXXMAMXMXMMMSMMMMMSMMXAMMMSMXMAMMMSMMXAMXMMXAMMMMAMSXMXSXAXMMMMMXMMSMSAXMAAMSSMASMMMXMMSASXSASXMSSSSSSXSAMAASMSMS
SAMXXSAMAXAASXSAMASAMXMMAMXMAXMAMMSMAMAAASMMAAMXXMXAAXMSMMMAXSXAMAMAMXAMSAXXAMXAMASXMMAAMASXMXAAMXSSSSMAAMMMMAMXSASMMXMASAXAAXXAAAMXMMMXAMMS
MMSMXMMMAMMMMMSXMASMSAMXSMMSAXMAMAAXMAXXXXAMSSSMSASMSSXMAASXMSMASAXXSXMMSXXSMMSAMASAXASMSASMSAMXMXXXMAMMSMAASXSAMXSAXSMASXMMSMMMMMMMXSMMXMAX
MMAAAXAMAXXXAXXXMXSXSAXXAMXMAXSAMXMAXAMXSSMMAAAAAMXAAXAXMASAAXAAMXSSMMXMMSMAAXMMMSSMMXMXMASAMASAMMMMSSMMAMMXMMMAXMSAMXMASAXAAAMXAXSMAMAMMMXM
SSSSMSMMSMMMMMMMMXMASAMSSMXMAMMXXMASMASAAAXMASMMXXMMMSSMSAMMMMSASMMAAAXAAAXXMMAXMXMAMSSSMMMXMASASXXAAAASMMSAAXSAMXMSSXMXSAMSMSSXMSMMASAMXXXX
AXAAAXAAXASXSMSASAMXMAMAAMAMXSXMSMAMMAMMSMSAMXMSXMMSMAMAMASXMAMAAAMSXMSMSSMSASXMAASAMAXAXAAXMASAMMMMMSAMAASMSMMXSAAXMASMXXMAMMAMMAMSAMASAMMM
MMSMMMMMSAMASASXSSSXSSMMMSAMAXAMAMAMMASXMMMMXAXMAMAAMAXMMXMXMASMSSMXAMMMAAASAMMMSMSXSMSMMMXSMXMAMXAAAMAMMMSXMXXASMSMSMMXMMSMSMAMSAMMMSMMAXMA
XAMXXAAAMAMXMAMMMAMXMAAXASAXSSSSMSASXMSAAXAXXMSSMMSSSSSMASMMMAXXAAXSXMAMMSMMSMXAAXXMSXSMAXMMMSMMSSMSSSSMMXSAMXMMSAXAAXMAAXAXXXAMMXMMAAXSASAM
MSSSMSMSSXMAMXMAMAMAMMMMMSAXMAAAAXXXAASXMMMSMAAAAAAMMAAXMSAAXASMSMMSMXXSAMMAMXMSSMSXSAMXXXAAAAAAXAAAXMAAAAMAMMASMAMXMSSMSSMSMMSSXSMMSSXMASMS
AXAMAAAMXASXSASMMASMMSASMMMMMAMMMXSXMMMXSAAAMMMSMMSSMSMMXSXMSXSAMXSXSXXMASAXXMAMAASMMSMSSMSSSSMMSMMMSSSMMXSAMAMMMAMASXXXAAMAAXAMXMAAAXXMAMXA
MMAMSMSMAMMMXASXSMSAASXSAASAMXMAXSMSMXMASMSSXSAMXMAMAAXMXMAXSXMAMMMAMMMMMMMXMXSMMMMAAAMAAMAMAXMXAAXSMMAAXXSMXSAMSSSXSASMSMSSSMASXSMMSMMMAXSX
XMAMXXXMAMAXMXMAAXMMMMASMMMMXMXXSMAAXAXAXAAAMMMMSAMXSMSXASMMSMSAMAMXMASAXXMAMAMAXXSMMSSSMMMXSAMSAMXSASXMMXMMAXAAAMMMMXMAXAXAAMSMMXSAAXAMXXSM
MSMSSMMSASXXXAMXMMMAAMAMASMSMSMMMMXMSMSSMMMMMAAXXSAMXAMXXXXXXXXASMSMSXSMSMSASMSMMMSAMXAXMAMXMAMAXXASAMMMSAMMMSAMXSASMAMSMSSSMSXSXMASMSXSSMSA
MAMXAAASASXASMXXXASXMMASXMAAAAMSASAXAAAXAMXMSMMMXMXXAAXSXSSXMAMXMASASASAAASAXAAMAAXAMMXXSASXSXMMSMMMSMAASASAAAASAMMASASAAXAXMXAMMSMMMXMAMAXS
MASXSMMSXXMASAASXMSAMSAMAMSMXMMSASMSMMMXSAMXSAXSAMSSSMAXSAMXMASAMXMAMAMXMXMAMSSSMSSMMSAASASAMSAMXMMAMSMXXAMMXSAMAXMMSXSMSMSMXMAXAAAMAMXMMMMM
SMSMXXMXMXXMMMMMAAMAMMASMMMXSXAMAMXAMAAMMMAXMAXXAMAAAMAMMXMXMASMSMMSMXMAMAMXMAAAMAAMAMXMMAMAMXSAMXMAMMMSXAXXAXMSSMSXXAMMMAMXXXMXSXXMAXXXAXAS
AXMMMASAMSAMXMASMMMSMSMMMAMMAMMXMAXXSMXSAMSXMSMSMMMSMMAXMASMMAMMAXMMASMMSAMSSMXMMMMMSSMXMXMXMAXXSASASAAMXMXMASAAAMXASXSASXSMSAAXAASXMMMSMSAS
MAMAXMMASXASMMMXSXAAMAAAMSXMMAMAXAXXAMMSMSMXXAAAXAMAASXMSASMMSMSASMSAMAAXAMAXAASAMSAMXAAXMMAMXSXSASASMXSASXSAMMMSAMXMAAXSAAASXSXMXMAAXXAAMMM
XASMSAMXMXXMAAMAMMSSMSSMXMAAXXSASMSMASAXMAMAMMSMSMXSXMMMMASAAAXMASASAMXMSXMXSMMXAAMXMASMSASAMMXAMXMAXMASASAMXSMSAASMMXMSMMMMMMMAXASMMMMMSMMS
SMSASXSMMMSSSMMAXXAAAXXXAXAMXXMXSAAMAMMMMAXMAXAMXMAXMMAAMXXMASMMMMMMAMXXAASMSASMSMSAMMXXAMXMSSSMMAMXMMAMXMXMAMXMAMMASAXAXAMSAAXMMMSAXMAMXMAS
AAMAMMXMAAAXAASXXMMXMMAMMSSMMAMAMMMMMAAAMASXMXAMXMASASMSXMASMXXSXAASAMMSMMSASAMAMXSASASMSMXSAMAASMSAXMAMSXMMSSMSSSSXMASXXSSSMSXSAAMMMSSSXSAS
MAMMMMASXSMSSMMSSSMMSMXMAAAAXXMAMAXMSSSXSASMMSSMMMXSAMAMMMXMSMMMSXMMAXAMXXMXMSMXMASAMAMAAASMASXMMASMMSSMMAAAMAAAAXAMXXMMAXAMXMASMSXXXSAMAMXS
XASMAMASXMXXMAMAXAAAMASMMSSMMMSMSMSMAAAAMXMAXAMAMSMMMMMMAMMAAXAMXMSSSMASXXSAXASAMXXMMAMXMSXSMMMAMAMMXMMASMMXSMMMSMMAMMAMXMAMAMMMAMAMXMAMSMAS
MAMXAMXXAMAXXAMMSSMMSAMAAAAAXAAMAASMMSMSMMMMMXSAMAAAXAXXAMMSMSSSSXAAMXMXMASMSMSXSMSMSMSXMXMXMASAMXSXAASAMAXAMASXMAXXAMXMASXSASAMMMMMASAMXMAM
MSMXMSMXXMASXMSXAXMXMMSXMSXMMSMSMAMAXMXXAASASASASMSXSAMMAXMXAXAAXMMSMASAMXMAAXXMMAAXAAMMMAAAMXMAAAMMSMMMSAMXSAMASMMMSSMSXSASASXSXMSSXSAMXMAS
AAAXSAMXAMAXAAXMXMAMSASXMMASMMAMXMXSMXMMSMSASXSAMXMMSASXSMXMMMMMMMAAMASXSAMSMXSAMSMSSSMASASXSASXMMSAMMSMSMMMMAMMMAAAAAAMAMMMXMAXAXMAMXAMMMMA
SMSMSAMMMMSMMMMMXSMAMASXMSASAMXMASAMXAXXXAMAMAMMMMMAMAMAMXMAMASAAMAMXXMAXAXXAASMMXAAAAXMMAMXXAMAMAMXSAAAMAXSSSMSSSMMSMMMMSXMAAAXMMMMAXAMXAAM
XAAMMMMXSAMAXSXSAMAMMAMAXMASAMASAMXSSMMMMMMAMAMMAMXSMMMAMASMMAXMMMSSMSAXSXMMMMSMSMMMMMMSMXMMXMSSMMSASXMSMMMXAMAXAXMAMXAAAXASMMMXMAAXSASMSMSX
MMMMASMXMASMMMAMMSSXMASMMMMMMMXMASAMXMASASMMSMMSASMXASXSSXSAMSSMSAMAAMAXSAXXSXSAXASXMSASMASAMMAXXXMMXAMXAXMMSMMMAMMSXMMMSMMMMAAAMSMMXAXAXAAA
SAAMMSAAMAMXAMXMXXMASXXAXMAAAMMSMMMSAMMSASXMSAASASXSXMAAMXSAMXAXMAMMMMMMXAMAXAMAMXSMASAMSMSAXMAMSAMXMSMSMMSMMASMXMAAAMSMXMAASMSMXAMXMAMMMMMS
SSMSXXMXMXSSXMXXMAMXMSSMSSSSMSAAAXAMAXAMAMAMMMMSAMAMXMSMSASAMMMMSMMASAXXXAMMMXMXMXXMAMXMAMMMMXSSXXXSXMAXMASASAMMAMMSSMAASAMXSXAMSMMAMAMMAMAM
MXMSASXXXMXMAMXMMSAMXASXAAXAXMMMMMMSMMSSMSXMASXMMMAMAXAAMASAMXMAAASASAAMSXMAMMMSMXXMXSXSMSMAAXXMMXMXAMAMXXSXMAXSASXAXXMMMMXAMXMMXSSXSSSSSSSS
AAXMASXMXAAMXSSMAXAAMASXMSMSMSSSMAXAMXXAXSAMASMMMSMSSXMXMXMAMXXSSMMXSMMXAAMMMSAAXXMMXMMSAAXASXSAAMAMXMSMXASXSSMSXSMMXXMMAXMSSSMAAMXAMAAXAAAM
SSMMAMAMASMSAAAMAXMSAMMMXAAAAAAAAXSASMSMMSAMAXAAAAAMAMSXSXSXSMXXXMMMXAAXMMMSAMSSMSAMSAAMSMMMXAXXMASXMAXXXXMASXAXMXMAASMSXSXAAAMMSMXXMMMMMMMM
XXASMSAMMMAMMMMMSMMAMAASMMSMMMSMMMMMXAAXASMMSMXMSSSMMMSAMAAAAXXSSMASXMMMXSXMAMXMAMAMMMMXAXXSMXMMMAMAMAMSXMMMMMSMMAMAMMAAAMMMAMXMAMSXMAMAAXXX
MSMMAMMMAMSMSXSAAAMASXMSXAAAXAAMXAXAMSMMASAAMXMXAMAAXAMAMXMSMAMMASAMAAAAASXMXMAXSSSMMAMMXMSMAMSSMXSXMASMXMAMAAXASXSSSMSMSMSXAMAMAMSAMAMSMSAM
AAMMSMASMSMAXAMMXMMXSXASMSMSMSSSSXSXXAAMAXMMSAMMASXMMMSAMXXAXMXSAMMSSMMMMSAASMSMAAAXMAMSAMAMMMAAXAXASXSAMMSMMMSAMAAAAAXAMASXASASXMSXMAXXAMAS
SASAAMXMXAMSMMMSSMSMSMXSAXAMXXMAXASAMXXMMXSAMAMSMMAXAXSMMMMASXMMAMXAAXXMASMMXAMAMSMXSAMSASASAMSSMXSMMMSASAXASAMAMAMSMMMAMAMMAXXMAXMMSSSMASAM
XXMXSXXXSMMXASAXMAAAXAMMAMXAAMMSMMMSAMXXSAMASMMXASXMSMMMAAMAMXMMMXMASASMAXAMMMMXMAXXMMMSAMXMMXMXMASXAASAMAXAMMSSMXAXAMXAMASXMMASXMAAXAAXXMXX
XSSMMASMXMAMMMMSMSMSMAAMAMXMAMAXASAAMSAMMASAMXXSAMXMAMASMSMSXXAASAMMAMXMAMAMSXSXXXSMXMAMMXSXMAMAMXMXMXMXMXMSMMAMMSSSMMSXSAMAASAMASXSMSMMMMMS
XXAAMAMAMSSMXAXMAXAMXXSXMMMSAMASAMSXXMAXSAMXXMAXMAXSASMSAMXAAXSMSASXAAXASXSASAXMSMMMSSSXSAMASXMXSAMXSAMXSAMAAMASAAXAMXSAMASXMMASAMXMAAAAAXAA
XSMMMAMMMAAMSSSMMMMMAAMASAASXSAMXMXXMSMMMXSAASXXSMASASAMXMMMXMMASAMXXMXMMMMAMAMAAAXAAAXAMASMMSAMSASASAMXSASXSMAMMSXXMAMXMAMAASAMXSAMXSSMMMSS
AAAXSMSMMMXMXAAMSAMXSMSAMMXMAMAMXXXMMAAMXAMXMMAASAMMAMXMXMASXMMAMAMXSMMMAXMXMSASMXMMMMMXMAXXAMMASXMASXMXSXMXAMASXMMMMSSMSAMMMMXSASMSAXAXXXAX
SSMMXAAXMASXXSMMSASAMXMAMSAMXMAMXMSAMAMMMXMXSMMMMAMXXXAMXSASAAMXMXMAAAAMSXSAAXMXMXSMMMSAMXXMSMMMMXMXMMSAMXMXMSXSAMASAAAMSSMSSSXMMSXMXMAMMAMX
XAMXMSMSMMXMAXAXSXMAMASAMXMXAAASAAASMSSSXSSXMASASXMAMSASAMASXMMSSMMSSSMMXASMSMAMXAAAAXSASMMMMASXMAXAAXAXXAMAMMMSMMMMMSMMXAXAAXAAXSXSSMAMXSXM
XSMSAAMAXMAMXMMMSXMAMAMXXAXMASASMSMXXMAMAASASASMSAMXMAMAAMAMAXAXAXAXMAXAMAMMMAASMSSSMXMMAAXAXAMASAMSSMMMMXMASAAMASXAMMXSMMMMSMMSMMAAXMXMAMAM
SMASMXSAXMASAMMXMASAMSSSSXSAMXAXAAAXAAAMMMSXMASMSMMMMASXMMSSSMSSMMSMSSMMSASXMSASAMAMXXSSMXMMSXSAMXMMAMAAXMSASXMMAMXXXAMMAAAAAAAAAMMMXMAMASAM
AMAMMMMMXSASAMAASMMMSXAAXMAXMSMMAMMXMMXSMAXAMXMAXXAAAAXASXAAMAXAXAAMMMMXSASAMXMMXMASMXMASXAMAXSMSMXSASMXSAAXMMXMXSMSMMAMSMMSSXMMXXMSASASMSXS
MMASAXAAMMXMXXSMMAAXXMMMMSMAMAAASXSXXSMMMMSSMSMMSSXSMMSAMXMMSSSSMSXSAAXAMXMMMAAXXMAMMMMAAMSMXXMASAMSAMASMXMASXMSXAAXAMSAMAXMAMSXSMMSASASASMS
XSASMSMSSMMMSXMASXMXXXAAMAXMSSMMMAXMAMAASMAMMXAXXAMMAMMASXXXAXMAXAMSXMXAMXMXSXSMSMMMXXMXMSAMXXMAMAXMAMXASAXMXMAMMMMMXMAASAMMAMMAASAMXMXMAMAM
XMAXAMAMAMAAAAMAMMAMSSSMSAMMAMAAMMMSSSMMXMAMSSXMMMASAXXAMAXMASMMMSMSAMMAMSAMXXMAAAAAMXSMMSASMMMSSMMSMMXMSMSMSMMSASMSSMMMMXXXXMMMMMMXAAXSAMMM
SMSMSMAXMSMSSMMMSMAMXAMXMMMMASXMMSAAAMXSMMSSMMMMAMMSMSMASXMASAAMXMAMAMXSAXSAMXSSSMMSAMXAAXAMXAMMAMMASAMXSXAAAAMSMSAAAMXSASMSXMXSASXSXSASMSXX
XAAAXXMMXXAAAAAAAXAMXAMXMMAMAMASMMMSSMXXAMMAAXAXASASAAXXMMMSMXXMAMAMXSMAAXAXSAMAXAAXMXSMMSMXSSSMAMMASXSAMSSSMSMMXMMMXSAMMSAAAMASXSAAAMXMAMMM
MSMSMASMSMMMSSMSXSSSSXMSMSASASAMXAAAAXMAMSSSMSSSMSMMSMSXMASASASXMSXSAMXMMMMMMASAMMXMSMAMAAMXXMAMXSMMMMMAXAXAAXXMASAAAMASXXMSSMASAMXMXMSMSXAA
AXSXXMAAMXSXXMAXASAMXAAXASASXSAXXMMSSSXSXMAMAAAAASXXAMAMSMSASMAAAXAMXXXMAXAAXAMXXXMMXSASMXSASMSMAMXMASXMMXSMXMASASMSSSSMMAMAXMXMAMAXAXAAXXSM
XSAMXXMXMAAAMAAMXMASXXMMMMAMASXMSMAAAAAMMMAMXMXMMMXSAMXMAMMAMXSMMMXMSAMSASXMSAMAMSMMASXSAAXXAAAMASMSMSAAXMXXASMMAMXXAMXASXMASMSAMXMSXSAMXXMM
MMMMMXXAMSSXAMXMSSSMMAAXXMXMAMXMAMMSSMXMASMSXXXSXSXSXMAXAXMXMAMMXSXAMMAAXMAMMXMAMAAMAMXMMMMSMXMXAMXAASMMMMMSASXSMSMMSASMMSMAXAMAMMMAXSAMSAMX
SASASXMMXAXXMXSSMXMAXSAMXMXMAMMSAMXXXXXSAMAAAMXMASAXAXSSSXSAMXSXAMMMMXSXXMXMMSSSSXSMSSXXAAAXAASMSMSMXMMMXAAMXMAXAAAAAMXAASMMSSMMMAMMXSAMMAMM
SASASMSSSMMSMMMAMMSAMAMSAMXSASAAASMMAXAMMSXMASAMAMMSSMMAAAMASAXMASAMSAMMSMMAAXAAAAXAMXMMSSSMXMSAAAXMXAAMSSXSSXAMSMMMSSSMMMMAAMASASMMAMSMXAMX
MAMMMAAAAXAAAASAMAMSSXMAXXMSASXMSAASAMXMASAXXAAMXSAAAAMMMMSAMXSAXSASMAMSAASMMSMMMMMAMXAAXAAAXXMMMSMSSMMMAMXXAMMXAAMAAXAXMAMMMSAMAXAMXSAMSXSX
MXMSMSMSMMSSSMSASAMXMASMXSXMAMXXMXXMXXMAMSAMXSMMMMMSSXMXXMAMXXSAMMMMMXMMSXMAMSXXMMSSMXSSMSMMXXSSXXAXAMXMASXMASXSSSMMSSSMSASAXMASXSSMMSXMSAMX`

