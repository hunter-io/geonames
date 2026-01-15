package geonames

// AdminToISO maps GeoNames administrative division codes to ISO 3166-2 subdivision codes.
var AdminToISO = map[string]string{
	"AD.02":       "AD-02",  // Canillo
	"AD.03":       "AD-03",  // Encamp
	"AD.04":       "AD-04",  // La Massana
	"AD.05":       "AD-05",  // Ordino
	"AD.06":       "AD-06",  // Sant Julià de Loria
	"AD.07":       "AD-07",  // Andorra la Vella
	"AD.08":       "AD-08",  // Escaldes-Engordany
	"AE.01":       "AE-AZ",  // Abu Dhabi
	"AE.02":       "AE-AJ",  // Ajman
	"AE.03":       "AE-DU",  // Dubai
	"AE.04":       "AE-FU",  // Fujairah
	"AE.05":       "AE-RK",  // Ras Al Khaimah
	"AE.06":       "AE-SH",  // Sharjah
	"AE.07":       "AE-UQ",  // Umm al Qaywayn
	"AF.01":       "AF-BDS", // Badakhshan
	"AF.02":       "AF-BDG", // Badghis
	"AF.03":       "AF-BGL", // Baghlan
	"AF.05":       "AF-BAM", // Bamyan
	"AF.06":       "AF-FRA", // Farah
	"AF.07":       "AF-FYB", // Faryab
	"AF.08":       "AF-GHA", // Ghazni
	"AF.09":       "AF-GHO", // Ghowr
	"AF.10":       "AF-HEL", // Helmand
	"AF.11":       "AF-HER", // Herat
	"AF.13":       "AF-KAB", // Kabul
	"AF.14":       "AF-KAP", // Kapisa
	"AF.17":       "AF-LOG", // Logar
	"AF.18":       "AF-NAN", // Nangarhar
	"AF.19":       "AF-NIM", // Nimroz
	"AF.23":       "AF-KAN", // Kandahar
	"AF.24":       "AF-KDZ", // Kunduz
	"AF.26":       "AF-TAK", // Takhar
	"AF.27":       "AF-WAR", // Maidan Wardak Province
	"AF.28":       "AF-ZAB", // Zabul
	"AF.29":       "AF-PKA", // Paktika
	"AF.30":       "AF-BAL", // Balkh
	"AF.31":       "AF-JOW", // Jowzjan
	"AF.32":       "AF-SAM", // Samangan
	"AF.33":       "AF-SAR", // Sar-e Pol Province
	"AF.34":       "AF-KNR", // Kunar
	"AF.35":       "AF-LAG", // Laghman
	"AF.36":       "AF-PIA", // Paktia
	"AF.37":       "AF-KHO", // Khowst
	"AF.38":       "AF-NUR", // Nuristan
	"AF.39":       "AF-URU", // Oruzgan
	"AF.40":       "AF-PAR", // Parwan
	"AF.41":       "AF-DAY", // Daykundi
	"AF.42":       "AF-PAN", // Panjshir
	"AG.01":       "AG-10",  // Barbuda
	"AG.03":       "AG-03",  // Saint George Parish
	"AG.04":       "AG-04",  // Saint John Parish
	"AG.05":       "AG-05",  // Saint Mary Parish
	"AG.06":       "AG-06",  // Saint Paul Parish
	"AG.07":       "AG-07",  // Saint Peter Parish
	"AG.08":       "AG-08",  // Saint Philip Parish
	"AG.09":       "AG-11",  // Redonda
	"AL.40":       "AL-01",  // Berat County
	"AL.41":       "AL-09",  // Dibër County
	"AL.42":       "AL-02",  // Durrës County
	"AL.43":       "AL-03",  // Elbasan County
	"AL.44":       "AL-04",  // Fier County
	"AL.45":       "AL-05",  // Gjirokastër County
	"AL.46":       "AL-06",  // Korçë County
	"AL.47":       "AL-07",  // Kukës County
	"AL.48":       "AL-08",  // Lezhë County
	"AL.49":       "AL-10",  // Shkodër County
	"AL.50":       "AL-11",  // Tirana
	"AL.51":       "AL-12",  // Vlorë County
	"AM.01":       "AM-AG",  // Aragatsotn
	"AM.02":       "AM-AR",  // Ararat
	"AM.03":       "AM-AV",  // Armavir
	"AM.04":       "AM-GR",  // Gegharkunik
	"AM.05":       "AM-KT",  // Kotayk
	"AM.06":       "AM-LO",  // Lori
	"AM.07":       "AM-SH",  // Shirak
	"AM.08":       "AM-SU",  // Syunik
	"AM.09":       "AM-TV",  // Tavush
	"AM.10":       "AM-VD",  // Vayots Dzor
	"AM.11":       "AM-ER",  // Yerevan
	"AO.01":       "AO-BGU", // Benguela
	"AO.02":       "AO-BIE", // Bíe
	"AO.03":       "AO-CAB", // Cabinda
	"AO.04":       "AO-CCU", // Cuando Cobango
	"AO.05":       "AO-CNO", // Cuanza Norte
	"AO.06":       "AO-CUS", // Kwanza Sul
	"AO.07":       "AO-CNN", // Cunene
	"AO.08":       "AO-HUA", // Huambo
	"AO.09":       "AO-HUI", // Huíla
	"AO.12":       "AO-MAL", // Malanje
	"AO.13":       "AO-NAM", // Namibe
	"AO.14":       "AO-MOX", // Moxico
	"AO.15":       "AO-UIG", // Uíge
	"AO.16":       "AO-ZAI", // Zaire
	"AO.17":       "AO-LNO", // Luanda Norte
	"AO.18":       "AO-LSU", // Lunda Sul
	"AO.19":       "AO-BGO", // Bengo
	"AO.20":       "AO-LUA", // Luanda
	"AR.01":       "AR-C",   // Buenos Aires
	"AR.02":       "AR-K",   // Catamarca
	"AR.03":       "AR-H",   // Chaco
	"AR.04":       "AR-U",   // Chubut
	"AR.05":       "AR-X",   // Cordoba
	"AR.06":       "AR-W",   // Corrientes
	"AR.08":       "AR-E",   // Entre Rios
	"AR.09":       "AR-P",   // Formosa
	"AR.10":       "AR-Y",   // Jujuy
	"AR.11":       "AR-L",   // La Pampa
	"AR.12":       "AR-F",   // La Rioja
	"AR.13":       "AR-M",   // Mendoza
	"AR.14":       "AR-N",   // Misiones
	"AR.15":       "AR-Q",   // Neuquen
	"AR.16":       "AR-R",   // Rio Negro
	"AR.17":       "AR-A",   // Salta
	"AR.18":       "AR-J",   // San Juan
	"AR.19":       "AR-D",   // San Luis
	"AR.20":       "AR-Z",   // Santa Cruz
	"AR.21":       "AR-S",   // Santa Fe
	"AR.22":       "AR-G",   // Santiago del Estero
	"AR.23":       "AR-V",   // Tierra del Fuego
	"AR.24":       "AR-T",   // Tucuman
	"AT.01":       "AT-1",   // Burgenland
	"AT.02":       "AT-2",   // Carinthia
	"AT.03":       "AT-3",   // Lower Austria
	"AT.04":       "AT-4",   // Upper Austria
	"AT.05":       "AT-5",   // Salzburg
	"AT.06":       "AT-6",   // Styria
	"AT.07":       "AT-7",   // Tyrol
	"AT.08":       "AT-8",   // Vorarlberg
	"AT.09":       "AT-9",   // Vienna
	"AU.01":       "AU-ACT", // Australian Capital Territory
	"AU.02":       "AU-NSW", // New South Wales
	"AU.03":       "AU-NT",  // Northern Territory
	"AU.04":       "AU-QLD", // Queensland
	"AU.05":       "AU-SA",  // South Australia
	"AU.06":       "AU-TAS", // Tasmania
	"AU.07":       "AU-VIC", // Victoria
	"AU.08":       "AU-WA",  // Western Australia
	"AZ.01":       "AZ-ABS", // Abşeron
	"AZ.02":       "AZ-AGC", // Ağcabǝdi
	"AZ.03":       "AZ-AGM", // Ağdam
	"AZ.04":       "AZ-AGS", // Ağdaş
	"AZ.05":       "AZ-AGA", // Ağstafa
	"AZ.06":       "AZ-AGU", // Agsu District
	"AZ.07":       "AZ-SR",  // Shirvan
	"AZ.08":       "AZ-AST", // Astara
	"AZ.09":       "AZ-BA",  // Baki
	"AZ.10":       "AZ-BAL", // Balakan District
	"AZ.11":       "AZ-BAR", // Barda
	"AZ.12":       "AZ-BEY", // Beyləqan
	"AZ.13":       "AZ-BIL", // Bilasuvar District
	"AZ.14":       "AZ-CAB", // Jabrayil
	"AZ.15":       "AZ-CAL", // Jalilabad
	"AZ.16":       "AZ-DAS", // Daşkǝsǝn
	"AZ.17":       "AZ-SBN", // Shabran
	"AZ.18":       "AZ-FUZ", // Fuzuli District
	"AZ.19":       "AZ-GAD", // Gadabay District
	"AZ.20":       "AZ-GA",  // Gǝncǝ
	"AZ.21":       "AZ-GOR", // Goranboy District
	"AZ.22":       "AZ-GOY", // Göyçay
	"AZ.23":       "AZ-HAC", // Hacıqabul
	"AZ.24":       "AZ-IMI", // İmişli
	"AZ.25":       "AZ-ISM", // Ismayilli District
	"AZ.26":       "AZ-KAL", // Kalbajar
	"AZ.27":       "AZ-KUR", // Kurdamir District
	"AZ.28":       "AZ-LAC", // Lachin Rayon
	"AZ.29":       "AZ-LA",  // Lənkəran
	"AZ.30":       "AZ-LA",  // Lankaran Sahari
	"AZ.31":       "AZ-LER", // Lerik District
	"AZ.32":       "AZ-MAS", // Masally
	"AZ.33":       "AZ-MI",  // Mingǝcevir
	"AZ.34":       "AZ-NA",  // Naftalan
	"AZ.35":       "AZ-NX",  // Nakhchivan
	"AZ.36":       "AZ-NEF", // Neftçala
	"AZ.37":       "AZ-OGU", // Oğuz
	"AZ.38":       "AZ-QAB", // Qabala District
	"AZ.39":       "AZ-QAX", // Qax
	"AZ.40":       "AZ-QAZ", // Qazax
	"AZ.41":       "AZ-QOB", // Qobustan
	"AZ.42":       "AZ-QBA", // Quba
	"AZ.43":       "AZ-QBI", // Qubadlı
	"AZ.44":       "AZ-QUS", // Qusar District
	"AZ.45":       "AZ-SAT", // Saatlı
	"AZ.46":       "AZ-SAB", // Sabirabad District
	"AZ.47":       "AZ-SA",  // Shaki
	"AZ.48":       "AZ-SA",  // Shaki City
	"AZ.49":       "AZ-SAL", // Salyan District
	"AZ.50":       "AZ-SMI", // Şamaxı
	"AZ.51":       "AZ-SKR", // Şǝmkir
	"AZ.52":       "AZ-SMX", // Samux
	"AZ.53":       "AZ-SIY", // Siyǝzǝn
	"AZ.54":       "AZ-SM",  // Sumqayit
	"AZ.55":       "AZ-SUS", // Shusha
	"AZ.56":       "AZ-SUS", // Shusha City
	"AZ.57":       "AZ-TAR", // Tartar District
	"AZ.58":       "AZ-TOV", // Tovuz District
	"AZ.60":       "AZ-XAC", // Xaçmaz
	"AZ.61":       "AZ-XA",  // Xankǝndi
	"AZ.62":       "AZ-GYG", // Goygol Rayon
	"AZ.63":       "AZ-XIZ", // Xızı
	"AZ.64":       "AZ-XCI", // Khojaly
	"AZ.65":       "AZ-XVD", // Khojavend
	"AZ.66":       "AZ-YAR", // Yardimli District
	"AZ.67":       "AZ-YE",  // Yevlax
	"AZ.68":       "AZ-YE",  // Yevlax City
	"AZ.69":       "AZ-ZAN", // Zangilan District
	"AZ.70":       "AZ-ZAQ", // Zaqatala District
	"AZ.71":       "AZ-ZAR", // Zərdab
	"BA.01":       "BA-BIH", // Federation of B&H
	"BA.02":       "BA-SRP", // Srpska
	"BA.BRC":      "BA-BRC", // Brčko
	"BB.01":       "BB-01",  // Christ Church
	"BB.02":       "BB-02",  // Saint Andrew
	"BB.03":       "BB-03",  // Saint George
	"BB.04":       "BB-04",  // Saint James
	"BB.05":       "BB-05",  // Saint John
	"BB.06":       "BB-06",  // Saint Joseph
	"BB.07":       "BB-07",  // Saint Lucy
	"BB.08":       "BB-08",  // Saint Michael
	"BB.09":       "BB-09",  // Saint Peter
	"BB.10":       "BB-10",  // Saint Philip
	"BB.11":       "BB-11",  // Saint Thomas
	"BD.81":       "BD-C",   // Dhaka Division
	"BD.82":       "BD-D",   // Khulna Division
	"BD.83":       "BD-E",   // Rajshahi Division
	"BD.84":       "BD-B",   // Chittagong
	"BD.85":       "BD-A",   // Barisal Division
	"BD.86":       "BD-G",   // Sylhet Division
	"BD.87":       "BD-F",   // Rangpur Division
	"BD.H":        "BD-H",   // Mymensingh Division
	"BE.BRU":      "BE-BRU", // Brussels Capital
	"BE.VLG":      "BE-VLG", // Flanders
	"BE.WAL":      "BE-WAL", // Wallonia
	"BF.01":       "BF-01",  // Boucle du Mouhoun
	"BF.02":       "BF-02",  // Cascades Region
	"BF.03":       "BF-03",  // Centre
	"BF.04":       "BF-04",  // Centre-Est
	"BF.05":       "BF-05",  // Centre-Nord
	"BF.06":       "BF-06",  // Centre-Ouest
	"BF.07":       "BF-07",  // Centre-Sud
	"BF.08":       "BF-08",  // Est
	"BF.09":       "BF-09",  // Hauts-Bassins
	"BF.10":       "BF-10",  // Nord
	"BF.11":       "BF-11",  // Plateau-Central
	"BF.12":       "BF-12",  // Sahel
	"BF.13":       "BF-13",  // Sud-Ouest
	"BG.38":       "BG-01",  // Blagoevgrad
	"BG.39":       "BG-02",  // Burgas
	"BG.40":       "BG-08",  // Dobrich
	"BG.41":       "BG-07",  // Gabrovo
	"BG.42":       "BG-22",  // Sofia-Capital
	"BG.43":       "BG-26",  // Haskovo
	"BG.44":       "BG-09",  // Kardzhali
	"BG.45":       "BG-10",  // Kyustendil
	"BG.46":       "BG-11",  // Lovech
	"BG.47":       "BG-12",  // Montana
	"BG.48":       "BG-13",  // Pazardzhik
	"BG.49":       "BG-14",  // Pernik
	"BG.50":       "BG-15",  // Pleven
	"BG.51":       "BG-16",  // Plovdiv
	"BG.52":       "BG-17",  // Razgrad
	"BG.53":       "BG-18",  // Ruse
	"BG.54":       "BG-27",  // Shumen
	"BG.55":       "BG-19",  // Silistra
	"BG.56":       "BG-20",  // Sliven
	"BG.57":       "BG-21",  // Smolyan
	"BG.58":       "BG-22",  // Sofia
	"BG.59":       "BG-24",  // Stara Zagora
	"BG.60":       "BG-25",  // Targovishte
	"BG.61":       "BG-03",  // Varna
	"BG.62":       "BG-04",  // Veliko Tarnovo
	"BG.63":       "BG-05",  // Vidin
	"BG.64":       "BG-06",  // Vratsa
	"BG.65":       "BG-28",  // Yambol
	"BH.15":       "BH-15",  // Muharraq
	"BH.16":       "BH-13",  // Manama
	"BH.17":       "BH-14",  // Southern Governorate
	"BH.19":       "BH-17",  // Northern
	"BI.09":       "BI-BB",  // Bubanza
	"BI.10":       "BI-BR",  // Bururi
	"BI.11":       "BI-CA",  // Cankuzo
	"BI.12":       "BI-CI",  // Cibitoke
	"BI.13":       "BI-GI",  // Gitega
	"BI.14":       "BI-KR",  // Karuzi
	"BI.15":       "BI-KY",  // Kayanza
	"BI.16":       "BI-KI",  // Kirundo
	"BI.17":       "BI-MA",  // Makamba
	"BI.18":       "BI-MY",  // Muyinga
	"BI.19":       "BI-NG",  // Ngozi
	"BI.20":       "BI-RT",  // Rutana
	"BI.21":       "BI-RY",  // Ruyigi
	"BI.22":       "BI-MU",  // Muramvya
	"BI.23":       "BI-MW",  // Mwaro
	"BI.24":       "BI-BM",  // Bujumbura Mairie
	"BI.25":       "BI-BL",  // Bujumbura Rural
	"BI.26":       "BI-RM",  // Rumonge
	"BJ.07":       "BJ-AL",  // Alibori
	"BJ.08":       "BJ-AK",  // Atakora
	"BJ.09":       "BJ-AQ",  // Atlantique
	"BJ.10":       "BJ-BO",  // Borgou
	"BJ.11":       "BJ-CO",  // Collines
	"BJ.12":       "BJ-KO",  // Kouffo
	"BJ.13":       "BJ-DO",  // Donga
	"BJ.14":       "BJ-LI",  // Littoral
	"BJ.15":       "BJ-MO",  // Mono
	"BJ.16":       "BJ-OU",  // Ouémé
	"BJ.17":       "BJ-PL",  // Plateau
	"BJ.18":       "BJ-ZO",  // Zou
	"BN.01":       "BN-BE",  // Belait
	"BN.02":       "BN-BM",  // Brunei-Muara District
	"BN.03":       "BN-TE",  // Temburong
	"BN.04":       "BN-TU",  // Tutong
	"BO.01":       "BO-H",   // Chuquisaca Department
	"BO.02":       "BO-C",   // Cochabamba
	"BO.03":       "BO-B",   // Beni Department
	"BO.04":       "BO-L",   // La Paz Department
	"BO.05":       "BO-O",   // Oruro
	"BO.06":       "BO-N",   // Pando
	"BO.07":       "BO-P",   // Potosí Department
	"BO.08":       "BO-S",   // Santa Cruz Department
	"BO.09":       "BO-T",   // Tarija Department
	"BQ.BO":       "BQ-BO",  // Bonaire
	"BQ.SB":       "BQ-SA",  // Saba
	"BQ.SE":       "BQ-SE",  // Sint Eustatius
	"BR.01":       "BR-AC",  // Acre
	"BR.02":       "BR-AL",  // Alagoas
	"BR.03":       "BR-AP",  // Amapá
	"BR.04":       "BR-AM",  // Amazonas
	"BR.05":       "BR-BA",  // Bahia
	"BR.06":       "BR-CE",  // Ceará
	"BR.07":       "BR-DF",  // Federal District
	"BR.08":       "BR-ES",  // Espírito Santo
	"BR.11":       "BR-MS",  // Mato Grosso do Sul
	"BR.13":       "BR-MA",  // Maranhão
	"BR.14":       "BR-MT",  // Mato Grosso
	"BR.15":       "BR-MG",  // Minas Gerais
	"BR.16":       "BR-PA",  // Pará
	"BR.17":       "BR-PB",  // Paraíba
	"BR.18":       "BR-PR",  // Paraná
	"BR.20":       "BR-PI",  // Piauí
	"BR.21":       "BR-RJ",  // Rio de Janeiro
	"BR.22":       "BR-RN",  // Rio Grande do Norte
	"BR.23":       "BR-RS",  // Rio Grande do Sul
	"BR.24":       "BR-RO",  // Rondônia
	"BR.25":       "BR-RR",  // Roraima
	"BR.26":       "BR-SC",  // Santa Catarina
	"BR.27":       "BR-SP",  // São Paulo
	"BR.28":       "BR-SE",  // Sergipe
	"BR.29":       "BR-GO",  // Goiás
	"BR.30":       "BR-PE",  // Pernambuco
	"BR.31":       "BR-TO",  // Tocantins
	"BS.05":       "BS-BI",  // Bimini
	"BS.06":       "BS-CI",  // Cat Island
	"BS.10":       "BS-EX",  // Exuma
	"BS.13":       "BS-IN",  // Inagua
	"BS.15":       "BS-LI",  // Long Island
	"BS.16":       "BS-MG",  // Mayaguana
	"BS.18":       "BS-RI",  // Ragged Island
	"BS.22":       "BS-HI",  // Harbour Island
	"BS.23":       "BS-NP",  // New Providence
	"BS.24":       "BS-AK",  // Acklins
	"BS.25":       "BS-FP",  // Freeport
	"BS.32":       "BS-BY",  // Berry Islands
	"BS.35":       "BS-SS",  // San Salvador
	"BS.36":       "BS-BP",  // Black Point
	"BS.37":       "BS-CO",  // Central Abaco
	"BS.38":       "BS-CS",  // Central Andros
	"BS.39":       "BS-CE",  // Central Eleuthera
	"BS.40":       "BS-CK",  // Crooked Island and Long Cay
	"BS.41":       "BS-EG",  // East Grand Bahama
	"BS.42":       "BS-GC",  // Grand Cay
	"BS.43":       "BS-HT",  // Hope Town
	"BS.44":       "BS-MC",  // Mangrove Cay
	"BS.45":       "BS-MI",  // Moore’s Island
	"BS.46":       "BS-NO",  // North Abaco
	"BS.47":       "BS-NS",  // North Andros
	"BS.48":       "BS-NE",  // North Eleuthera
	"BS.49":       "BS-RC",  // Rum Cay
	"BS.50":       "BS-SO",  // South Abaco
	"BS.51":       "BS-SA",  // South Andros
	"BS.52":       "BS-SE",  // South Eleuthera
	"BS.53":       "BS-SW",  // Spanish Wells
	"BS.54":       "BS-WG",  // West Grand Bahama
	"BT.05":       "BT-33",  // Bumthang District
	"BT.06":       "BT-12",  // Chukha
	"BT.07":       "BT-21",  // Tsirang District
	"BT.08":       "BT-22",  // Dagana
	"BT.09":       "BT-31",  // Sarpang District
	"BT.10":       "BT-13",  // Haa
	"BT.11":       "BT-44",  // Lhuntse
	"BT.12":       "BT-42",  // Mongar
	"BT.13":       "BT-11",  // Paro
	"BT.14":       "BT-43",  // Pemagatshel
	"BT.15":       "BT-23",  // Punakha
	"BT.16":       "BT-14",  // Samtse District
	"BT.17":       "BT-45",  // Samdrup Jongkhar
	"BT.18":       "BT-34",  // Zhemgang District
	"BT.19":       "BT-41",  // Trashigang District
	"BT.20":       "BT-15",  // Thimphu District
	"BT.21":       "BT-32",  // Tongsa
	"BT.22":       "BT-24",  // Wangdi Phodrang
	"BT.23":       "BT-GA",  // Gasa
	"BT.24":       "BT-TY",  // Trashi Yangste
	"BW.01":       "BW-CE",  // Central
	"BW.03":       "BW-GH",  // Ghanzi
	"BW.04":       "BW-KG",  // Kgalagadi
	"BW.05":       "BW-KL",  // Kgatleng
	"BW.06":       "BW-KW",  // Kweneng
	"BW.08":       "BW-NE",  // North-East
	"BW.09":       "BW-SE",  // South-East
	"BW.11":       "BW-NW",  // North-West
	"BW.12":       "BW-CH",  // Chobe
	"BW.13":       "BW-FR",  // City of Francistown
	"BW.14":       "BW-GA",  // Gaborone
	"BW.15":       "BW-JW",  // Jwaneng
	"BW.16":       "BW-LO",  // Lobatse
	"BW.17":       "BW-SP",  // Selibe Phikwe
	"BW.18":       "BW-ST",  // Sowa Town
	"BY.01":       "BY-BR",  // Brest
	"BY.02":       "BY-HO",  // Homyel’ Voblasc’
	"BY.03":       "BY-HR",  // Grodnenskaya
	"BY.04":       "BY-HM",  // Minsk City
	"BY.05":       "BY-HM",  // Minsk
	"BY.06":       "BY-MA",  // Mogilev
	"BY.07":       "BY-VI",  // Vitebsk
	"BZ.02":       "BZ-CY",  // Cayo District
	"BZ.03":       "BZ-CZL", // Corozal District
	"BZ.04":       "BZ-OW",  // Orange Walk District
	"BZ.06":       "BZ-TOL", // Toledo District
	"CA.01":       "CA-AB",  // Alberta
	"CA.02":       "CA-BC",  // British Columbia
	"CA.03":       "CA-MB",  // Manitoba
	"CA.04":       "CA-NB",  // New Brunswick
	"CA.05":       "CA-NL",  // Newfoundland and Labrador
	"CA.07":       "CA-NS",  // Nova Scotia
	"CA.08":       "CA-ON",  // Ontario
	"CA.09":       "CA-PE",  // Prince Edward Island
	"CA.10":       "CA-QC",  // Quebec
	"CA.11":       "CA-SK",  // Saskatchewan
	"CA.12":       "CA-YT",  // Yukon
	"CA.13":       "CA-NT",  // Northwest Territories
	"CA.14":       "CA-NU",  // Nunavut
	"CD.02":       "CD-EQ",  // Équateur
	"CD.06":       "CD-KN",  // Kinshasa
	"CD.08":       "CD-BC",  // Bas-Congo
	"CD.10":       "CD-MA",  // Maniema
	"CD.11":       "CD-NK",  // North Kivu
	"CD.12":       "CD-SK",  // South Kivu
	"CD.13":       "CD-BU",  // Bas-Uele
	"CD.14":       "CD-HK",  // Haut-Katanga
	"CD.15":       "CD-HL",  // Haut-Lomami
	"CD.16":       "CD-HU",  // Haut-Uele
	"CD.17":       "CD-IT",  // Ituri
	"CD.18":       "CD-KS",  // Kasai
	"CD.19":       "CD-KG",  // Kwango
	"CD.20":       "CD-KL",  // Kwilu
	"CD.21":       "CD-LO",  // Lomami
	"CD.22":       "CD-LU",  // Lualaba
	"CD.23":       "CD-KC",  // Kasai-Central
	"CD.24":       "CD-MN",  // Mai-Ndombe
	"CD.25":       "CD-MO",  // Mongala
	"CD.26":       "CD-NU",  // Nord-Ubangi
	"CD.27":       "CD-SA",  // Sankuru
	"CD.28":       "CD-SU",  // Sud-Ubangi
	"CD.29":       "CD-TA",  // Tanganyika
	"CD.30":       "CD-TO",  // Tshopo
	"CD.31":       "CD-TU",  // Tshuapa
	"CF.01":       "CF-BB",  // Bamingui-Bangoran
	"CF.02":       "CF-BK",  // Basse-Kotto
	"CF.03":       "CF-HK",  // Haute-Kotto
	"CF.04":       "CF-HS",  // Mambéré-Kadéï
	"CF.05":       "CF-HM",  // Haut-Mbomou
	"CF.06":       "CF-KG",  // Kémo
	"CF.07":       "CF-LB",  // Lobaye
	"CF.08":       "CF-MB",  // Mbomou
	"CF.09":       "CF-NM",  // Nana-Mambéré
	"CF.11":       "CF-UK",  // Ouaka
	"CF.12":       "CF-AC",  // Ouham
	"CF.13":       "CF-OP",  // Ouham-Pendé
	"CF.14":       "CF-VK",  // Vakaga
	"CF.15":       "CF-KB",  // Nana-Grébizi
	"CF.16":       "CF-SE",  // Sangha-Mbaéré
	"CF.17":       "CF-MP",  // Ombella-M'Poko
	"CF.18":       "CF-BGF", // Bangui
	"CF.19":       "CF-HS",  // Mambéré
	"CF.21":       "CF-OP",  // Lim-Pendé
	"CG.01":       "CG-11",  // Bouenza
	"CG.04":       "CG-5",   // Kouilou
	"CG.05":       "CG-2",   // Lékoumou
	"CG.06":       "CG-7",   // Likouala
	"CG.07":       "CG-9",   // Niari
	"CG.08":       "CG-14",  // Plateaux
	"CG.10":       "CG-13",  // Sangha
	"CG.11":       "CG-12",  // Pool
	"CG.12":       "CG-BZV", // Brazzaville
	"CG.13":       "CG-8",   // Cuvette
	"CG.14":       "CG-15",  // Cuvette-Ouest
	"CG.15":       "CG-16",  // Pointe-Noire
	"CH.AG":       "CH-AG",  // Aargau
	"CH.AI":       "CH-AI",  // Appenzell Innerrhoden
	"CH.AR":       "CH-AR",  // Appenzell Ausserrhoden
	"CH.BE":       "CH-BE",  // Bern
	"CH.BL":       "CH-BL",  // Basel-Landschaft
	"CH.BS":       "CH-BS",  // Basel-City
	"CH.FR":       "CH-FR",  // Fribourg
	"CH.GE":       "CH-GE",  // Geneva
	"CH.GL":       "CH-GL",  // Glarus
	"CH.GR":       "CH-GR",  // Grisons
	"CH.JU":       "CH-JU",  // Jura
	"CH.LU":       "CH-LU",  // Lucerne
	"CH.NE":       "CH-NE",  // Neuchâtel
	"CH.NW":       "CH-NW",  // Nidwalden
	"CH.OW":       "CH-OW",  // Obwalden
	"CH.SG":       "CH-SG",  // Saint Gallen
	"CH.SH":       "CH-SH",  // Schaffhausen
	"CH.SO":       "CH-SO",  // Solothurn
	"CH.SZ":       "CH-SZ",  // Schwyz
	"CH.TG":       "CH-TG",  // Thurgau
	"CH.TI":       "CH-TI",  // Ticino
	"CH.UR":       "CH-UR",  // Uri
	"CH.VD":       "CH-VD",  // Vaud
	"CH.VS":       "CH-VS",  // Valais
	"CH.ZG":       "CH-ZG",  // Zug
	"CH.ZH":       "CH-ZH",  // Zurich
	"CI.76":       "CI-BS",  // Bas-Sassandra District
	"CI.77":       "CI-DN",  // Denguélé District
	"CI.78":       "CI-MG",  // Montagnes
	"CI.81":       "CI-LC",  // Lacs District
	"CI.82":       "CI-LG",  // Lagunes District
	"CI.87":       "CI-SV",  // Savanes District
	"CI.90":       "CI-VB",  // Vallée du Bandama District
	"CI.92":       "CI-ZZ",  // Zanzan District
	"CI.93":       "CI-AB",  // Abidjan Autonomous District
	"CI.94":       "CI-CM",  // Comoé District
	"CI.95":       "CI-GD",  // Gôh-Djiboua
	"CI.96":       "CI-SM",  // Sassandra-Marahoué
	"CI.97":       "CI-WR",  // Woroba
	"CI.98":       "CI-YM",  // Yamoussoukro
	"CL.01":       "CL-VS",  // Valparaíso
	"CL.02":       "CL-AI",  // Aysén
	"CL.03":       "CL-AN",  // Antofagasta
	"CL.04":       "CL-AR",  // Araucanía
	"CL.05":       "CL-AT",  // Atacama
	"CL.06":       "CL-BI",  // Biobío
	"CL.07":       "CL-CO",  // Coquimbo Region
	"CL.08":       "CL-LI",  // O'Higgins Region
	"CL.10":       "CL-MA",  // Region of Magallanes
	"CL.11":       "CL-ML",  // Maule Region
	"CL.12":       "CL-RM",  // Santiago Metropolitan
	"CL.14":       "CL-LL",  // Los Lagos Region
	"CL.15":       "CL-TA",  // Tarapacá
	"CL.16":       "CL-AP",  // Arica y Parinacota Region
	"CL.17":       "CL-LR",  // Los Ríos Region
	"CL.18":       "CL-NB",  // Ñuble
	"CM.04":       "CM-ES",  // East
	"CM.05":       "CM-LT",  // Littoral
	"CM.07":       "CM-NW",  // North-West
	"CM.08":       "CM-OU",  // West
	"CM.09":       "CM-SW",  // South-West
	"CM.10":       "CM-AD",  // Adamaoua
	"CM.11":       "CM-CE",  // Centre
	"CM.12":       "CM-EN",  // Far North
	"CM.13":       "CM-NO",  // North
	"CM.14":       "CM-SU",  // South
	"CN.01":       "CN-AH",  // Anhui
	"CN.02":       "CN-ZJ",  // Zhejiang
	"CN.03":       "CN-JX",  // Jiangxi
	"CN.04":       "CN-JS",  // Jiangsu
	"CN.05":       "CN-JL",  // Jilin
	"CN.06":       "CN-QH",  // Qinghai
	"CN.07":       "CN-FJ",  // Fujian
	"CN.08":       "CN-HL",  // Heilongjiang
	"CN.09":       "CN-HA",  // Henan
	"CN.10":       "CN-HE",  // Hebei
	"CN.11":       "CN-HN",  // Hunan
	"CN.12":       "CN-HB",  // Hubei
	"CN.13":       "CN-XJ",  // Xinjiang
	"CN.14":       "CN-XZ",  // Tibet
	"CN.15":       "CN-GS",  // Gansu
	"CN.16":       "CN-GX",  // Guangxi
	"CN.18":       "CN-GZ",  // Guizhou
	"CN.19":       "CN-LN",  // Liaoning
	"CN.20":       "CN-NM",  // Inner Mongolia
	"CN.21":       "CN-NX",  // Ningxia
	"CN.22":       "CN-BJ",  // Beijing
	"CN.23":       "CN-SH",  // Shanghai
	"CN.24":       "CN-SX",  // Shanxi
	"CN.25":       "CN-SD",  // Shandong
	"CN.26":       "CN-SN",  // Shaanxi
	"CN.28":       "CN-TJ",  // Tianjin
	"CN.29":       "CN-YN",  // Yunnan
	"CN.30":       "CN-GD",  // Guangdong
	"CN.31":       "CN-HI",  // Hainan
	"CN.32":       "CN-SC",  // Sichuan
	"CN.33":       "CN-CQ",  // Chongqing
	"CO.01":       "CO-AMA", // Amazonas Department
	"CO.02":       "CO-ANT", // Antioquia
	"CO.03":       "CO-ARA", // Arauca Department
	"CO.04":       "CO-ATL", // Atlántico
	"CO.08":       "CO-CAQ", // Caquetá
	"CO.09":       "CO-CAU", // Cauca Department
	"CO.10":       "CO-CES", // Cesar Department
	"CO.11":       "CO-CHO", // Chocó
	"CO.12":       "CO-COR", // Córdoba
	"CO.14":       "CO-GUV", // Guaviare Department
	"CO.15":       "CO-GUA", // Guainía Department
	"CO.16":       "CO-HUI", // Huila Department
	"CO.17":       "CO-LAG", // La Guajira Department
	"CO.19":       "CO-MET", // Meta Department
	"CO.20":       "CO-NAR", // Nariño
	"CO.21":       "CO-NSA", // Norte de Santander Department
	"CO.22":       "CO-PUT", // Putumayo Department
	"CO.23":       "CO-QUI", // Quindío Department
	"CO.24":       "CO-RIS", // Risaralda Department
	"CO.25":       "CO-SAP", // San Andres y Providencia
	"CO.26":       "CO-SAN", // Santander Department
	"CO.27":       "CO-SUC", // Sucre Department
	"CO.28":       "CO-TOL", // Tolima Department
	"CO.29":       "CO-VAC", // Valle del Cauca Department
	"CO.30":       "CO-VAU", // Vaupés
	"CO.31":       "CO-VID", // Vichada Department
	"CO.32":       "CO-CAS", // Casanare Department
	"CO.33":       "CO-CUN", // Cundinamarca
	"CO.34":       "CO-DC",  // Bogota D.C.
	"CO.35":       "CO-BOL", // Bolívar
	"CO.36":       "CO-BOY", // Boyacá
	"CO.37":       "CO-CAL", // Caldas Department
	"CO.38":       "CO-MAG", // Magdalena Department
	"CR.01":       "CR-A",   // Alajuela Province
	"CR.02":       "CR-C",   // Cartago Province
	"CR.03":       "CR-G",   // Guanacaste Province
	"CR.04":       "CR-H",   // Heredia Province
	"CR.06":       "CR-L",   // Limón Province
	"CR.07":       "CR-P",   // Puntarenas Province
	"CR.08":       "CR-SJ",  // San José
	"CU.01":       "CU-01",  // Pinar del Río
	"CU.02":       "CU-03",  // Havana
	"CU.03":       "CU-04",  // Matanzas Province
	"CU.04":       "CU-99",  // Isla de la Juventud
	"CU.05":       "CU-09",  // Camagüey
	"CU.07":       "CU-08",  // Ciego de Ávila Province
	"CU.08":       "CU-06",  // Cienfuegos Province
	"CU.09":       "CU-12",  // Granma Province
	"CU.10":       "CU-14",  // Guantánamo Province
	"CU.12":       "CU-11",  // Holguín Province
	"CU.13":       "CU-10",  // Las Tunas Province
	"CU.14":       "CU-07",  // Sancti Spíritus Province
	"CU.15":       "CU-13",  // Santiago de Cuba Province
	"CU.16":       "CU-05",  // Villa Clara Province
	"CU.AR":       "CU-15",  // Artemisa
	"CU.MA":       "CU-16",  // Mayabeque
	"CV.01":       "CV-B",   // Boa Vista
	"CV.02":       "CV-S",   // Brava
	"CV.04":       "CV-S",   // Maio
	"CV.05":       "CV-B",   // Paul
	"CV.07":       "CV-B",   // Ribeira Grande
	"CV.08":       "CV-B",   // Sal
	"CV.11":       "CV-B",   // São Vicente
	"CV.13":       "CV-S",   // Mosteiros
	"CV.14":       "CV-S",   // Praia
	"CV.15":       "CV-S",   // Santa Catarina
	"CV.16":       "CV-S",   // Santa Cruz
	"CV.17":       "CV-S",   // São Domingos
	"CV.18":       "CV-S",   // São Filipe
	"CV.19":       "CV-S",   // São Miguel
	"CV.20":       "CV-S",   // Tarrafal
	"CV.21":       "CV-B",   // Porto Novo
	"CV.22":       "CV-B",   // Ribeira Brava
	"CV.23":       "CV-S",   // Ribeira Grande de Santiago
	"CV.24":       "CV-S",   // Santa Catarina do Fogo
	"CV.25":       "CV-S",   // São Lourenço dos Órgãos
	"CV.26":       "CV-S",   // São Salvador do Mundo
	"CV.27":       "CV-B",   // Tarrafal de São Nicolau
	"CY.02":       "CY-06",  // Keryneia
	"CY.03":       "CY-03",  // Larnaka
	"CY.04":       "CY-01",  // Nicosia
	"CY.05":       "CY-02",  // Limassol
	"CY.06":       "CY-05",  // Pafos
	"CZ.52":       "CZ-10",  // Prague
	"CZ.78":       "CZ-64",  // South Moravian
	"CZ.79":       "CZ-31",  // Jihočeský kraj
	"CZ.80":       "CZ-63",  // Vysočina
	"CZ.81":       "CZ-41",  // Karlovarský kraj
	"CZ.82":       "CZ-52",  // Královéhradecký kraj
	"CZ.83":       "CZ-51",  // Liberecký kraj
	"CZ.84":       "CZ-71",  // Olomoucký
	"CZ.85":       "CZ-80",  // Moravskoslezský
	"CZ.86":       "CZ-53",  // Pardubický
	"CZ.87":       "CZ-32",  // Plzeň Region
	"CZ.88":       "CZ-20",  // Central Bohemia
	"CZ.89":       "CZ-42",  // Ústecký kraj
	"CZ.90":       "CZ-72",  // Zlín
	"DE.01":       "DE-BW",  // Baden-Wurttemberg
	"DE.02":       "DE-BY",  // Bavaria
	"DE.03":       "DE-HB",  // City state Bremen
	"DE.04":       "DE-HH",  // Hamburg
	"DE.05":       "DE-HE",  // Hesse
	"DE.06":       "DE-NI",  // Lower Saxony
	"DE.07":       "DE-NW",  // North Rhine-Westphalia
	"DE.08":       "DE-RP",  // Rheinland-Pfalz
	"DE.09":       "DE-SL",  // Saarland
	"DE.10":       "DE-SH",  // Schleswig-Holstein
	"DE.11":       "DE-BB",  // Brandenburg
	"DE.12":       "DE-MV",  // Mecklenburg-Vorpommern
	"DE.13":       "DE-SN",  // Saxony
	"DE.14":       "DE-ST",  // Saxony-Anhalt
	"DE.15":       "DE-TH",  // Thuringia
	"DE.16":       "DE-BE",  // State of Berlin
	"DJ.01":       "DJ-AS",  // Ali Sabieh
	"DJ.04":       "DJ-OB",  // Obock
	"DJ.05":       "DJ-TA",  // Tadjourah
	"DJ.06":       "DJ-DI",  // Dikhil
	"DJ.07":       "DJ-DJ",  // Djibouti
	"DJ.08":       "DJ-AR",  // Arta
	"DK.17":       "DK-84",  // Capital Region
	"DK.18":       "DK-82",  // Central Jutland
	"DK.19":       "DK-81",  // North Denmark
	"DK.20":       "DK-85",  // Zealand
	"DK.21":       "DK-83",  // South Denmark
	"DM.02":       "DM-02",  // Saint Andrew Parish
	"DM.03":       "DM-03",  // Saint David Parish
	"DM.04":       "DM-04",  // Saint George Parish
	"DM.05":       "DM-05",  // Saint John Parish
	"DM.06":       "DM-06",  // Saint Joseph Parish
	"DM.07":       "DM-07",  // Saint Luke Parish
	"DM.08":       "DM-08",  // Saint Mark Parish
	"DM.09":       "DM-09",  // Saint Patrick Parish
	"DM.10":       "DM-10",  // Saint Paul Parish
	"DM.11":       "DM-11",  // Saint Peter Parish
	"DO.01":       "DO-41",  // Azua Province
	"DO.02":       "DO-38",  // Baoruco Province
	"DO.03":       "DO-38",  // Barahona Province
	"DO.04":       "DO-34",  // Dajabón
	"DO.06":       "DO-33",  // Duarte Province
	"DO.08":       "DO-35",  // Espaillat Province
	"DO.09":       "DO-38",  // Independencia
	"DO.10":       "DO-42",  // La Altagracia Province
	"DO.11":       "DO-37",  // Elías Piña
	"DO.12":       "DO-42",  // La Romana
	"DO.14":       "DO-33",  // María Trinidad Sánchez
	"DO.15":       "DO-34",  // Monte Cristi Province
	"DO.16":       "DO-38",  // Pedernales Province
	"DO.18":       "DO-35",  // Puerto Plata
	"DO.19":       "DO-33",  // Hermanas Mirabal
	"DO.20":       "DO-33",  // Samaná
	"DO.21":       "DO-36",  // Sánchez Ramírez
	"DO.23":       "DO-37",  // San Juan Province
	"DO.24":       "DO-39",  // San Pedro de Macorís
	"DO.25":       "DO-35",  // Santiago Province
	"DO.26":       "DO-34",  // Santiago Rodríguez
	"DO.27":       "DO-34",  // Valverde Province
	"DO.28":       "DO-42",  // El Seíbo
	"DO.29":       "DO-39",  // Hato Mayor Province
	"DO.30":       "DO-36",  // La Vega
	"DO.31":       "DO-36",  // Monseñor Nouel
	"DO.32":       "DO-39",  // Monte Plata Province
	"DO.33":       "DO-41",  // San Cristóbal
	"DO.34":       "DO-40",  // Nacional
	"DO.35":       "DO-41",  // Peravia
	"DO.36":       "DO-41",  // San José de Ocoa
	"DO.37":       "DO-40",  // Santo Domingo Province
	"DZ.01":       "DZ-16",  // Algiers
	"DZ.03":       "DZ-05",  // Batna
	"DZ.04":       "DZ-25",  // Constantine
	"DZ.06":       "DZ-26",  // Medea
	"DZ.07":       "DZ-27",  // Mostaganem
	"DZ.09":       "DZ-31",  // Oran
	"DZ.10":       "DZ-20",  // Saida
	"DZ.12":       "DZ-19",  // Sétif
	"DZ.13":       "DZ-14",  // Tiaret
	"DZ.14":       "DZ-15",  // Tizi Ouzou
	"DZ.15":       "DZ-13",  // Tlemcen
	"DZ.18":       "DZ-06",  // Béjaïa
	"DZ.19":       "DZ-07",  // Biskra
	"DZ.20":       "DZ-09",  // Blida
	"DZ.21":       "DZ-10",  // Bouira
	"DZ.22":       "DZ-17",  // Djelfa
	"DZ.23":       "DZ-24",  // Guelma
	"DZ.24":       "DZ-18",  // Jijel
	"DZ.25":       "DZ-03",  // Laghouat
	"DZ.26":       "DZ-29",  // Mascara
	"DZ.27":       "DZ-28",  // M'Sila
	"DZ.29":       "DZ-04",  // Oum el Bouaghi
	"DZ.30":       "DZ-22",  // Sidi Bel Abbès
	"DZ.31":       "DZ-21",  // Skikda
	"DZ.33":       "DZ-12",  // Tébessa
	"DZ.34":       "DZ-01",  // Adrar
	"DZ.35":       "DZ-44",  // Aïn Defla
	"DZ.36":       "DZ-46",  // Aïn Témouchent
	"DZ.37":       "DZ-23",  // Annaba
	"DZ.38":       "DZ-08",  // Béchar
	"DZ.39":       "DZ-34",  // Bordj Bou Arréridj
	"DZ.40":       "DZ-35",  // Boumerdes
	"DZ.41":       "DZ-02",  // Chlef
	"DZ.42":       "DZ-32",  // El Bayadh
	"DZ.43":       "DZ-39",  // El Oued
	"DZ.44":       "DZ-36",  // El Tarf
	"DZ.45":       "DZ-47",  // Ghardaia
	"DZ.46":       "DZ-33",  // Illizi
	"DZ.47":       "DZ-40",  // Khenchela
	"DZ.48":       "DZ-43",  // Mila
	"DZ.49":       "DZ-45",  // Naama
	"DZ.50":       "DZ-30",  // Ouargla
	"DZ.51":       "DZ-48",  // Relizane
	"DZ.52":       "DZ-41",  // Souk Ahras
	"DZ.53":       "DZ-11",  // Tamanrasset
	"DZ.54":       "DZ-37",  // Tindouf
	"DZ.55":       "DZ-42",  // Tipaza
	"DZ.56":       "DZ-38",  // Tissemsilt
	"EC.01":       "EC-W",   // Galápagos
	"EC.02":       "EC-A",   // Azuay
	"EC.03":       "EC-B",   // Bolívar
	"EC.04":       "EC-F",   // Cañar
	"EC.05":       "EC-C",   // Carchi
	"EC.06":       "EC-H",   // Chimborazo Province
	"EC.07":       "EC-X",   // Cotopaxi
	"EC.08":       "EC-O",   // El Oro
	"EC.09":       "EC-E",   // Esmeraldas
	"EC.10":       "EC-G",   // Guayas
	"EC.11":       "EC-I",   // Imbabura
	"EC.12":       "EC-L",   // Loja
	"EC.13":       "EC-R",   // Los Ríos
	"EC.14":       "EC-M",   // Manabí
	"EC.15":       "EC-S",   // Morona-Santiago Province
	"EC.17":       "EC-Y",   // Pastaza Province
	"EC.18":       "EC-P",   // Pichincha
	"EC.19":       "EC-T",   // Tungurahua Province
	"EC.20":       "EC-Z",   // Zamora Chinchipe
	"EC.22":       "EC-U",   // Sucumbíos
	"EC.23":       "EC-N",   // Napo
	"EC.24":       "EC-D",   // Orellana Province
	"EC.25":       "EC-SE",  // Santa Elena
	"EC.26":       "EC-SD",  // Santo Domingo de los Tsáchilas
	"EE.01":       "EE-37",  // Harjumaa
	"EE.02":       "EE-205", // Hiiumaa
	"EE.03":       "EE-45",  // Ida-Virumaa
	"EE.04":       "EE-52",  // Järvamaa
	"EE.05":       "EE-50",  // Jõgevamaa
	"EE.07":       "EE-56",  // Lääne
	"EE.08":       "EE-60",  // Lääne-Virumaa
	"EE.11":       "EE-68",  // Pärnumaa
	"EE.12":       "EE-64",  // Põlvamaa
	"EE.13":       "EE-71",  // Raplamaa
	"EE.14":       "EE-74",  // Saare
	"EE.18":       "EE-79",  // Tartu
	"EE.19":       "EE-81",  // Valgamaa
	"EE.20":       "EE-84",  // Viljandimaa
	"EE.21":       "EE-87",  // Võrumaa
	"EG.01":       "EG-DK",  // Dakahlia
	"EG.02":       "EG-BA",  // Red Sea
	"EG.03":       "EG-BH",  // Beheira
	"EG.04":       "EG-FYM", // Faiyum
	"EG.05":       "EG-GH",  // Gharbia
	"EG.06":       "EG-ALX", // Alexandria
	"EG.07":       "EG-IS",  // Ismailia
	"EG.08":       "EG-GZ",  // Giza
	"EG.09":       "EG-MNF", // Monufia
	"EG.10":       "EG-MN",  // Minya
	"EG.11":       "EG-C",   // Cairo
	"EG.12":       "EG-KB",  // Qalyubia
	"EG.13":       "EG-WAD", // New Valley
	"EG.14":       "EG-SHR", // Sharqia
	"EG.15":       "EG-SUZ", // Suez
	"EG.16":       "EG-ASN", // Aswan
	"EG.17":       "EG-AST", // Asyut
	"EG.18":       "EG-BNS", // Beni Suweif
	"EG.19":       "EG-PTS", // Port Said
	"EG.20":       "EG-DT",  // Damietta
	"EG.21":       "EG-KFS", // Kafr el-Sheikh
	"EG.22":       "EG-MT",  // Matruh
	"EG.23":       "EG-KN",  // Qena
	"EG.24":       "EG-SHG", // Sohag
	"EG.26":       "EG-JS",  // South Sinai
	"EG.27":       "EG-SIN", // North Sinai
	"EG.28":       "EG-LX",  // Luxor
	"ER.01":       "ER-AN",  // Anseba
	"ER.02":       "ER-DU",  // Debub
	"ER.03":       "ER-DK",  // Southern Red Sea
	"ER.04":       "ER-GB",  // Gash-Barka
	"ER.05":       "ER-MA",  // Maekel
	"ER.06":       "ER-SK",  // Northern Red Sea
	"ES.07":       "ES-IB",  // Balearic Islands
	"ES.27":       "ES-RI",  // La Rioja
	"ES.29":       "ES-MD",  // Madrid
	"ES.31":       "ES-MU",  // Murcia
	"ES.32":       "ES-NA",  // Navarre
	"ES.34":       "ES-AS",  // Asturias
	"ES.39":       "ES-CB",  // Cantabria
	"ES.51":       "ES-AN",  // Andalusia
	"ES.52":       "ES-AR",  // Aragon
	"ES.53":       "ES-CN",  // Canary Islands
	"ES.54":       "ES-CM",  // Castille-La Mancha
	"ES.55":       "ES-CL",  // Castille and León
	"ES.56":       "ES-CT",  // Catalonia
	"ES.57":       "ES-EX",  // Extremadura
	"ES.58":       "ES-GA",  // Galicia
	"ES.59":       "ES-PV",  // Basque Country
	"ES.60":       "ES-VC",  // Valencia
	"ES.CE":       "ES-CE",  // Ceuta
	"ES.ML":       "ES-ML",  // Melilla
	"ET.44":       "ET-AA",  // Addis Ababa
	"ET.45":       "ET-AF",  // Āfar
	"ET.46":       "ET-AM",  // Amhara
	"ET.47":       "ET-BE",  // Bīnshangul Gumuz
	"ET.48":       "ET-DD",  // Dire Dawa
	"ET.49":       "ET-GA",  // Gambela
	"ET.50":       "ET-HA",  // Harari
	"ET.51":       "ET-OR",  // Oromiya
	"ET.52":       "ET-SO",  // Somali
	"ET.53":       "ET-TI",  // Tigray
	"ET.56":       "ET-SN",  // South Ethiopia Regional State
	"ET.SI":       "ET-SI",  // Sidama Region
	"FI.01":       "FI-18",  // Uusimaa
	"FI.02":       "FI-19",  // Southwest Finland
	"FI.04":       "FI-17",  // Satakunta
	"FI.06":       "FI-11",  // Pirkanmaa
	"FI.07":       "FI-16",  // Paijat-Hame
	"FI.08":       "FI-09",  // Kymenlaakso
	"FI.09":       "FI-02",  // South Karelia
	"FI.10":       "FI-04",  // South Savo
	"FI.11":       "FI-15",  // North Savo
	"FI.12":       "FI-13",  // North Karelia
	"FI.13":       "FI-08",  // Central Finland
	"FI.14":       "FI-03",  // South Ostrobothnia
	"FI.15":       "FI-12",  // Ostrobothnia
	"FI.16":       "FI-07",  // Central Ostrobothnia
	"FI.17":       "FI-14",  // North Ostrobothnia
	"FI.18":       "FI-05",  // Kainuu
	"FI.19":       "FI-10",  // Lapland
	"FJ.01":       "FJ-C",   // Central
	"FJ.02":       "FJ-E",   // Eastern
	"FJ.03":       "FJ-N",   // Northern
	"FJ.04":       "FJ-R",   // Rotuma
	"FJ.05":       "FJ-W",   // Western
	"FM.01":       "FM-KSA", // Kosrae
	"FM.02":       "FM-PNI", // Pohnpei State
	"FM.03":       "FM-TRK", // Chuuk
	"FM.04":       "FM-YAP", // Yap State
	"FR.11":       "FR-IDF", // Île-de-France
	"FR.24":       "FR-CVL", // Centre-Val de Loire
	"FR.27":       "FR-BFC", // Bourgogne
	"FR.28":       "FR-NOR", // Normandy
	"FR.32":       "FR-HDF", // Hauts-de-France
	"FR.44":       "FR-GES", // Grand Est
	"FR.52":       "FR-PDL", // Pays de la Loire
	"FR.53":       "FR-BRE", // Brittany
	"FR.75":       "FR-NAQ", // Nouvelle-Aquitaine
	"FR.76":       "FR-OCC", // Occitanie
	"FR.84":       "FR-ARA", // Rhône-Alpes
	"FR.93":       "FR-PAC", // Provence-Alpes-Côte d'Azur
	"FR.94":       "FR-20R", // Corsica
	"GA.01":       "GA-1",   // Estuaire
	"GA.02":       "GA-2",   // Haut-Ogooué
	"GA.03":       "GA-3",   // Moyen-Ogooué
	"GA.04":       "GA-4",   // Ngouni
	"GA.05":       "GA-5",   // Nyanga
	"GA.06":       "GA-6",   // Ogooué-Ivindo
	"GA.07":       "GA-7",   // Ogooué-Lolo
	"GA.08":       "GA-8",   // Ogooué-Maritime
	"GA.09":       "GA-9",   // Woleu-Ntem
	"GB.ENG":      "GB-ENG", // England
	"GB.NIR":      "GB-NIR", // Northern Ireland
	"GB.SCT":      "GB-SCT", // Scotland
	"GB.WLS":      "GB-WLS", // Wales
	"GD.01":       "GD-01",  // Saint Andrew Parish
	"GD.02":       "GD-02",  // Saint David Parish
	"GD.03":       "GD-03",  // Saint George Parish
	"GD.04":       "GD-04",  // Saint John Parish
	"GD.05":       "GD-05",  // Saint Mark Parish
	"GD.06":       "GD-06",  // Saint Patrick Parish
	"GD.10":       "GD-10",  // Carriacou and Petite Martinique
	"GE.02":       "GE-AB",  // Abkhazia
	"GE.04":       "GE-AJ",  // Adjara
	"GE.51":       "GE-TB",  // Tbilisi
	"GE.65":       "GE-GU",  // Guria
	"GE.66":       "GE-IM",  // Imereti
	"GE.67":       "GE-KA",  // Kakheti
	"GE.68":       "GE-KK",  // Kvemo Kartli
	"GE.69":       "GE-MM",  // Mtskheta-Mtianeti
	"GE.70":       "GE-RL",  // Racha-Lechkhumi and Kvemo Svaneti
	"GE.71":       "GE-SZ",  // Samegrelo and Zemo Svaneti
	"GE.72":       "GE-SJ",  // Samtskhe-Javakheti
	"GE.73":       "GE-SK",  // Shida Kartli
	"GH.01":       "GH-AA",  // Greater Accra
	"GH.02":       "GH-AH",  // Ashanti
	"GH.04":       "GH-CP",  // Central
	"GH.05":       "GH-EP",  // Eastern
	"GH.06":       "GH-NP",  // Northern
	"GH.08":       "GH-TV",  // Volta
	"GH.09":       "GH-WP",  // Western
	"GH.10":       "GH-UE",  // Upper East
	"GH.11":       "GH-UW",  // Upper West
	"GH.12":       "GH-AF",  // Ahafo
	"GH.13":       "GH-BO",  // Bono
	"GH.14":       "GH-BE",  // Bono East
	"GH.15":       "GH-NE",  // North East
	"GH.16":       "GH-OT",  // Oti
	"GH.17":       "GH-SV",  // Savannah
	"GH.18":       "GH-WN",  // Western North
	"GL.04":       "GL-KU",  // Kujalleq
	"GL.06":       "GL-QE",  // Qeqqata
	"GL.07":       "GL-SM",  // Sermersooq
	"GL.11839534": "GL-QT",  // Qeqertalik
	"GL.11839537": "GL-AV",  // Avannaata
	"GM.01":       "GM-B",   // Banjul
	"GM.02":       "GM-L",   // Lower River
	"GM.03":       "GM-M",   // Central River
	"GM.04":       "GM-U",   // Upper River
	"GM.07":       "GM-N",   // North Bank
	"GN.04":       "GN-C",   // Conakry
	"GN.B":        "GN-B",   // Boké Region
	"GN.D":        "GN-D",   // Kindia
	"GN.F":        "GN-F",   // Faranah
	"GN.K":        "GN-K",   // Kankan Region
	"GN.L":        "GN-L",   // Labé Region
	"GN.M":        "GN-M",   // Mamou Region
	"GN.N":        "GN-N",   // Nzérékoré Region
	"GQ.03":       "GQ-I",   // Annobon
	"GQ.04":       "GQ-I",   // Bioko Norte
	"GQ.05":       "GQ-I",   // Bioko Sur
	"GQ.06":       "GQ-C",   // Centro Sur
	"GQ.07":       "GQ-C",   // Kié-Ntem
	"GQ.08":       "GQ-C",   // Litoral
	"GQ.09":       "GQ-C",   // Wele-Nzas
	"GQ.10":       "GQ-C",   // Djibloho
	"GR.736572":   "GR-69",  // Mount Athos
	"GR.ESYE11":   "GR-A",   // East Macedonia and Thrace
	"GR.ESYE12":   "GR-B",   // Central Macedonia
	"GR.ESYE13":   "GR-C",   // West Macedonia
	"GR.ESYE14":   "GR-E",   // Thessaly
	"GR.ESYE21":   "GR-D",   // Epirus
	"GR.ESYE22":   "GR-F",   // Ionian Islands
	"GR.ESYE23":   "GR-G",   // West Greece
	"GR.ESYE24":   "GR-H",   // Central Greece
	"GR.ESYE25":   "GR-J",   // Peloponnese
	"GR.ESYE31":   "GR-I",   // Attica
	"GR.ESYE41":   "GR-K",   // North Aegean
	"GR.ESYE42":   "GR-L",   // South Aegean
	"GR.ESYE43":   "GR-M",   // Crete
	"GT.01":       "GT-16",  // Alta Verapaz
	"GT.02":       "GT-15",  // Baja Verapaz
	"GT.03":       "GT-04",  // Chimaltenango
	"GT.04":       "GT-20",  // Chiquimula
	"GT.05":       "GT-02",  // El Progreso
	"GT.06":       "GT-05",  // Escuintla
	"GT.07":       "GT-01",  // Guatemala
	"GT.08":       "GT-13",  // Huehuetenango
	"GT.09":       "GT-18",  // Izabal Department
	"GT.10":       "GT-21",  // Jalapa
	"GT.11":       "GT-22",  // Jutiapa
	"GT.12":       "GT-17",  // Petén
	"GT.13":       "GT-09",  // Quetzaltenango
	"GT.14":       "GT-14",  // Quiché
	"GT.15":       "GT-11",  // Retalhuleu
	"GT.16":       "GT-03",  // Sacatepéquez
	"GT.17":       "GT-12",  // San Marcos
	"GT.18":       "GT-06",  // Santa Rosa Department
	"GT.19":       "GT-07",  // Sololá
	"GT.20":       "GT-10",  // Suchitepeque
	"GT.21":       "GT-08",  // Totonicapán
	"GT.22":       "GT-19",  // Zacapa
	"GW.01":       "GW-L",   // Bafatá
	"GW.02":       "GW-S",   // Quinara
	"GW.04":       "GW-N",   // Oio
	"GW.05":       "GW-S",   // Bolama
	"GW.06":       "GW-N",   // Cacheu
	"GW.07":       "GW-S",   // Tombali
	"GW.10":       "GW-L",   // Gabú
	"GW.11":       "GW-BS",  // Bissau
	"GW.12":       "GW-N",   // Biombo
	"GY.10":       "GY-BA",  // Barima-Waini
	"GY.11":       "GY-CU",  // Cuyuni-Mazaruni
	"GY.12":       "GY-DE",  // Demerara-Mahaica
	"GY.13":       "GY-EB",  // East Berbice-Corentyne
	"GY.14":       "GY-ES",  // Essequibo Islands-West Demerara
	"GY.15":       "GY-MA",  // Mahaica-Berbice
	"GY.16":       "GY-PM",  // Pomeroon-Supenaam
	"GY.17":       "GY-PT",  // Potaro-Siparuni
	"GY.18":       "GY-UD",  // Upper Demerara-Berbice
	"GY.19":       "GY-UT",  // Upper Takutu-Upper Essequibo
	"HN.01":       "HN-AT",  // Atlántida Department
	"HN.02":       "HN-CH",  // Choluteca Department
	"HN.03":       "HN-CL",  // Colón Department
	"HN.04":       "HN-CM",  // Comayagua Department
	"HN.05":       "HN-CP",  // Copán Department
	"HN.06":       "HN-CR",  // Cortés Department
	"HN.07":       "HN-EP",  // El Paraíso Department
	"HN.08":       "HN-FM",  // Francisco Morazán Department
	"HN.09":       "HN-GD",  // Gracias a Dios Department
	"HN.10":       "HN-IN",  // Intibucá Department
	"HN.11":       "HN-IB",  // Bay Islands
	"HN.12":       "HN-LP",  // La Paz Department
	"HN.13":       "HN-LE",  // Lempira Department
	"HN.14":       "HN-OC",  // Ocotepeque Department
	"HN.15":       "HN-OL",  // Olancho Department
	"HN.16":       "HN-SB",  // Santa Bárbara Department
	"HN.17":       "HN-VA",  // Valle Department
	"HN.18":       "HN-YO",  // Yoro Department
	"HR.01":       "HR-07",  // Bjelovar-Bilogora
	"HR.02":       "HR-12",  // Brod-Posavina
	"HR.03":       "HR-19",  // Dubrovnik-Neretva
	"HR.04":       "HR-18",  // Istria
	"HR.05":       "HR-04",  // Karlovac
	"HR.06":       "HR-06",  // Koprivnica-Križevci
	"HR.07":       "HR-02",  // Krapina-Zagorje
	"HR.08":       "HR-09",  // Lika-Senj
	"HR.09":       "HR-20",  // Međimurje
	"HR.10":       "HR-14",  // County of Osijek-Baranja
	"HR.11":       "HR-11",  // Požega-Slavonia
	"HR.12":       "HR-08",  // Primorje-Gorski Kotar
	"HR.13":       "HR-15",  // Šibenik-Knin
	"HR.14":       "HR-03",  // Sisak-Moslavina
	"HR.15":       "HR-17",  // Split-Dalmatia
	"HR.16":       "HR-05",  // Varaždin
	"HR.17":       "HR-10",  // Virovitica-Podravina
	"HR.18":       "HR-16",  // Vukovar-Srijem
	"HR.19":       "HR-13",  // Zadar
	"HR.20":       "HR-01",  // Zagreb County
	"HR.21":       "HR-21",  // Zagreb
	"HT.03":       "HT-NO",  // Nord-Ouest
	"HT.06":       "HT-AR",  // Artibonite
	"HT.07":       "HT-CE",  // Centre
	"HT.09":       "HT-ND",  // Nord
	"HT.10":       "HT-NE",  // Nord-Est
	"HT.11":       "HT-OU",  // Ouest
	"HT.12":       "HT-SD",  // Sud
	"HT.13":       "HT-SE",  // Sud-Est
	"HT.14":       "HT-GA",  // Grand'Anse
	"HT.15":       "HT-NI",  // Nippes
	"HU.01":       "HU-BK",  // Bács-Kiskun
	"HU.02":       "HU-BA",  // Baranya
	"HU.03":       "HU-BE",  // Bekes County
	"HU.04":       "HU-BZ",  // Borsod-Abaúj-Zemplén
	"HU.05":       "HU-BU",  // Budapest
	"HU.06":       "HU-CS",  // Csongrád
	"HU.08":       "HU-FE",  // Fejér
	"HU.09":       "HU-GS",  // Győr-Moson-Sopron
	"HU.10":       "HU-HB",  // Hajdú-Bihar
	"HU.11":       "HU-HE",  // Heves County
	"HU.12":       "HU-KE",  // Komárom-Esztergom
	"HU.14":       "HU-NO",  // Nógrád
	"HU.16":       "HU-PE",  // Pest County
	"HU.17":       "HU-SO",  // Somogy County
	"HU.18":       "HU-SZ",  // Szabolcs-Szatmár-Bereg
	"HU.20":       "HU-JN",  // Jász-Nagykun-Szolnok
	"HU.21":       "HU-TO",  // Tolna County
	"HU.22":       "HU-VA",  // Vas County
	"HU.23":       "HU-VM",  // Veszprém
	"HU.24":       "HU-ZA",  // Zala County
	"ID.01":       "ID-SM",  // Aceh
	"ID.02":       "ID-NU",  // Bali
	"ID.03":       "ID-SM",  // Bengkulu
	"ID.04":       "ID-JW",  // Jakarta
	"ID.05":       "ID-SM",  // Jambi
	"ID.07":       "ID-JW",  // Central Java
	"ID.08":       "ID-JW",  // East Java
	"ID.10":       "ID-JW",  // Yogyakarta
	"ID.11":       "ID-KA",  // West Kalimantan
	"ID.12":       "ID-KA",  // South Kalimantan
	"ID.13":       "ID-KA",  // Central Kalimantan
	"ID.14":       "ID-KA",  // East Kalimantan
	"ID.15":       "ID-SM",  // Lampung
	"ID.17":       "ID-NU",  // West Nusa Tenggara
	"ID.18":       "ID-NU",  // East Nusa Tenggara
	"ID.21":       "ID-SL",  // Central Sulawesi
	"ID.22":       "ID-SL",  // Southeast Sulawesi
	"ID.24":       "ID-SM",  // West Sumatra
	"ID.26":       "ID-SM",  // North Sumatra
	"ID.28":       "ID-ML",  // Maluku
	"ID.29":       "ID-ML",  // North Maluku
	"ID.30":       "ID-JW",  // West Java
	"ID.31":       "ID-SL",  // North Sulawesi
	"ID.32":       "ID-SM",  // South Sumatra
	"ID.33":       "ID-JW",  // Banten
	"ID.34":       "ID-SL",  // Gorontalo
	"ID.35":       "ID-SM",  // Bangka–Belitung Islands
	"ID.36":       "ID-PP",  // Papua
	"ID.37":       "ID-SM",  // Riau
	"ID.38":       "ID-SL",  // South Sulawesi
	"ID.39":       "ID-PP",  // West Papua
	"ID.40":       "ID-SM",  // Riau Islands
	"ID.41":       "ID-SL",  // West Sulawesi
	"ID.42":       "ID-KA",  // North Kalimantan
	"ID.PE":       "ID-PP",  // Highland Papua
	"IE.C":        "IE-C",   // Connacht
	"IE.L":        "IE-L",   // Leinster
	"IE.M":        "IE-M",   // Munster
	"IE.U":        "IE-U",   // Ulster
	"IL.01":       "IL-D",   // Southern District
	"IL.02":       "IL-M",   // Central District
	"IL.03":       "IL-Z",   // Northern District
	"IL.04":       "IL-HA",  // Haifa
	"IL.05":       "IL-TA",  // Tel Aviv
	"IL.06":       "IL-JM",  // Jerusalem
	"IN.01":       "IN-AN",  // Andaman and Nicobar
	"IN.02":       "IN-AP",  // Andhra Pradesh
	"IN.03":       "IN-AS",  // Assam
	"IN.05":       "IN-CH",  // Chandigarh
	"IN.07":       "IN-DL",  // Delhi
	"IN.09":       "IN-GJ",  // Gujarat
	"IN.10":       "IN-HR",  // Haryana
	"IN.11":       "IN-HP",  // Himachal Pradesh
	"IN.12":       "IN-JK",  // Jammu and Kashmir
	"IN.13":       "IN-KL",  // Kerala
	"IN.14":       "IN-LD",  // Lakshadweep
	"IN.16":       "IN-MH",  // Maharashtra
	"IN.17":       "IN-MN",  // Manipur
	"IN.18":       "IN-ML",  // Meghalaya
	"IN.19":       "IN-KA",  // Karnataka
	"IN.20":       "IN-NL",  // Nagaland
	"IN.21":       "IN-OR",  // Odisha
	"IN.22":       "IN-PY",  // Puducherry
	"IN.23":       "IN-PB",  // Punjab
	"IN.24":       "IN-RJ",  // Rajasthan
	"IN.25":       "IN-TN",  // Tamil Nadu
	"IN.26":       "IN-TR",  // Tripura
	"IN.28":       "IN-WB",  // West Bengal
	"IN.29":       "IN-SK",  // Sikkim
	"IN.30":       "IN-AR",  // Arunachal Pradesh
	"IN.31":       "IN-MZ",  // Mizoram
	"IN.33":       "IN-GA",  // Goa
	"IN.34":       "IN-BR",  // Bihar
	"IN.35":       "IN-MP",  // Madhya Pradesh
	"IN.36":       "IN-UP",  // Uttar Pradesh
	"IN.37":       "IN-CT",  // Chhattisgarh
	"IN.38":       "IN-JH",  // Jharkhand
	"IN.39":       "IN-UT",  // Uttarakhand
	"IN.40":       "IN-TG",  // Telangana
	"IN.41":       "IN-LA",  // Ladakh
	"IN.52":       "IN-DH",  // Dadra and Nagar Haveli and Daman and Diu
	"IQ.01":       "IQ-AN",  // Al Anbar
	"IQ.02":       "IQ-BA",  // Basra
	"IQ.03":       "IQ-MU",  // Al Muthanna Governorate
	"IQ.04":       "IQ-QA",  // Al Qādisīyah
	"IQ.05":       "IQ-KR",  // Sulaymaniyah
	"IQ.06":       "IQ-BB",  // Bābil
	"IQ.07":       "IQ-BG",  // Baghdad
	"IQ.08":       "IQ-KR",  // Duhok
	"IQ.09":       "IQ-DQ",  // Dhi Qar
	"IQ.10":       "IQ-DI",  // Diyālá
	"IQ.11":       "IQ-KR",  // Erbil
	"IQ.12":       "IQ-KA",  // Karbalāʼ
	"IQ.13":       "IQ-KI",  // Kirkuk
	"IQ.14":       "IQ-MA",  // Maysan
	"IQ.15":       "IQ-NI",  // Nineveh
	"IQ.16":       "IQ-WA",  // Wāsiţ
	"IQ.17":       "IQ-NA",  // An Najaf
	"IQ.18":       "IQ-SD",  // Salah ad Din
	"IR.01":       "IR-02",  // West Azerbaijan
	"IR.03":       "IR-08",  // Chaharmahal and Bakhtiari
	"IR.04":       "IR-13",  // Sistan and Baluchestan
	"IR.05":       "IR-18",  // Kohgiluyeh and Boyer-Ahmad
	"IR.07":       "IR-14",  // Fars
	"IR.08":       "IR-19",  // Gilan Province
	"IR.09":       "IR-24",  // Hamadan Province
	"IR.10":       "IR-05",  // Ilam Province
	"IR.11":       "IR-23",  // Hormozgan
	"IR.13":       "IR-17",  // Kermanshah Province
	"IR.15":       "IR-10",  // Khuzestan
	"IR.16":       "IR-16",  // Kurdistan Province
	"IR.22":       "IR-06",  // Bushehr
	"IR.23":       "IR-20",  // Lorestan Province
	"IR.25":       "IR-12",  // Semnan
	"IR.26":       "IR-07",  // Tehran
	"IR.28":       "IR-04",  // Isfahan
	"IR.29":       "IR-15",  // Kerman
	"IR.32":       "IR-03",  // Ardabil Province
	"IR.33":       "IR-01",  // East Azerbaijan
	"IR.34":       "IR-22",  // Markazi
	"IR.35":       "IR-21",  // Māzandarān
	"IR.36":       "IR-11",  // Zanjan
	"IR.37":       "IR-27",  // Golestan
	"IR.38":       "IR-28",  // Qazvin Province
	"IR.39":       "IR-26",  // Qom Province
	"IR.40":       "IR-25",  // Yazd Province
	"IR.41":       "IR-29",  // South Khorasan Province
	"IR.42":       "IR-30",  // Razavi Khorasan
	"IS.38":       "IS-7",   // East
	"IS.39":       "IS-1",   // Capital Region
	"IS.40":       "IS-6",   // Northeast
	"IS.41":       "IS-5",   // Northwest
	"IS.42":       "IS-8",   // South
	"IS.43":       "IS-2",   // Southern Peninsula
	"IS.44":       "IS-4",   // Westfjords
	"IS.45":       "IS-3",   // West
	"IT.01":       "IT-65",  // Abruzzo
	"IT.02":       "IT-77",  // Basilicate
	"IT.03":       "IT-78",  // Calabria
	"IT.04":       "IT-72",  // Campania
	"IT.05":       "IT-45",  // Emilia-Romagna
	"IT.06":       "IT-36",  // Friuli Venezia Giulia
	"IT.07":       "IT-62",  // Lazio
	"IT.08":       "IT-42",  // Liguria
	"IT.09":       "IT-25",  // Lombardy
	"IT.10":       "IT-57",  // The Marches
	"IT.11":       "IT-67",  // Molise
	"IT.12":       "IT-21",  // Piedmont
	"IT.13":       "IT-75",  // Apulia
	"IT.14":       "IT-88",  // Sardinia
	"IT.15":       "IT-82",  // Sicily
	"IT.16":       "IT-52",  // Tuscany
	"IT.17":       "IT-32",  // Trentino-Alto Adige
	"IT.18":       "IT-55",  // Umbria
	"IT.19":       "IT-23",  // Aosta Valley
	"IT.20":       "IT-34",  // Veneto
	"JM.01":       "JM-13",  // Clarendon
	"JM.02":       "JM-09",  // Hanover
	"JM.04":       "JM-12",  // Manchester
	"JM.07":       "JM-04",  // Portland Parish
	"JM.08":       "JM-02",  // Saint Andrew Parish
	"JM.09":       "JM-06",  // Saint Ann Parish
	"JM.10":       "JM-14",  // Saint Catherine Parish
	"JM.11":       "JM-11",  // St. Elizabeth
	"JM.12":       "JM-08",  // Saint James Parish
	"JM.13":       "JM-05",  // Saint Mary Parish
	"JM.14":       "JM-03",  // Saint Thomas Parish
	"JM.15":       "JM-07",  // Trelawny Parish
	"JM.16":       "JM-10",  // Westmoreland
	"JM.17":       "JM-01",  // Kingston
	"JO.02":       "JO-BA",  // Balqa
	"JO.09":       "JO-KA",  // Karak
	"JO.12":       "JO-AT",  // Tafielah
	"JO.15":       "JO-MA",  // Mafraq
	"JO.16":       "JO-AM",  // Amman
	"JO.17":       "JO-AZ",  // Zarqa
	"JO.18":       "JO-IR",  // Irbid
	"JO.19":       "JO-MN",  // Ma’an
	"JO.20":       "JO-AJ",  // Ajloun
	"JO.21":       "JO-AQ",  // Aqaba
	"JO.22":       "JO-JA",  // Jerash
	"JO.23":       "JO-MD",  // Madaba
	"JP.01":       "JP-23",  // Aichi
	"JP.02":       "JP-05",  // Akita
	"JP.03":       "JP-02",  // Aomori
	"JP.04":       "JP-12",  // Chiba
	"JP.05":       "JP-38",  // Ehime
	"JP.06":       "JP-18",  // Fukui
	"JP.07":       "JP-40",  // Fukuoka
	"JP.08":       "JP-07",  // Fukushima
	"JP.09":       "JP-21",  // Gifu
	"JP.10":       "JP-10",  // Gunma
	"JP.11":       "JP-34",  // Hiroshima
	"JP.12":       "JP-01",  // Hokkaido
	"JP.13":       "JP-28",  // Hyōgo
	"JP.14":       "JP-08",  // Ibaraki
	"JP.15":       "JP-17",  // Ishikawa
	"JP.16":       "JP-03",  // Iwate
	"JP.17":       "JP-37",  // Kagawa
	"JP.18":       "JP-46",  // Kagoshima
	"JP.19":       "JP-14",  // Kanagawa
	"JP.20":       "JP-39",  // Kochi
	"JP.21":       "JP-43",  // Kumamoto
	"JP.22":       "JP-26",  // Kyoto
	"JP.23":       "JP-24",  // Mie
	"JP.24":       "JP-04",  // Miyagi
	"JP.25":       "JP-45",  // Miyazaki
	"JP.26":       "JP-20",  // Nagano
	"JP.27":       "JP-42",  // Nagasaki
	"JP.28":       "JP-29",  // Nara
	"JP.29":       "JP-15",  // Niigata
	"JP.30":       "JP-44",  // Oita
	"JP.31":       "JP-33",  // Okayama
	"JP.32":       "JP-27",  // Osaka
	"JP.33":       "JP-41",  // Saga
	"JP.34":       "JP-11",  // Saitama
	"JP.35":       "JP-25",  // Shiga
	"JP.36":       "JP-32",  // Shimane
	"JP.37":       "JP-22",  // Shizuoka
	"JP.38":       "JP-09",  // Tochigi
	"JP.39":       "JP-36",  // Tokushima
	"JP.40":       "JP-13",  // Tokyo
	"JP.41":       "JP-31",  // Tottori
	"JP.42":       "JP-16",  // Toyama
	"JP.43":       "JP-30",  // Wakayama
	"JP.44":       "JP-06",  // Yamagata
	"JP.45":       "JP-35",  // Yamaguchi
	"JP.46":       "JP-19",  // Yamanashi
	"JP.47":       "JP-47",  // Okinawa
	"KE.05":       "KE-30",  // Nairobi County
	"KE.10":       "KE-01",  // Baringo
	"KE.11":       "KE-02",  // Bomet County
	"KE.12":       "KE-03",  // Bungoma County
	"KE.13":       "KE-04",  // Busia County
	"KE.14":       "KE-05",  // Elegeyo-Marakwet
	"KE.15":       "KE-06",  // Embu County
	"KE.16":       "KE-07",  // Garissa County
	"KE.17":       "KE-08",  // Homa Bay County
	"KE.18":       "KE-09",  // Isiolo County
	"KE.19":       "KE-10",  // Kajiado County
	"KE.20":       "KE-11",  // Kakamega County
	"KE.21":       "KE-12",  // Kericho County
	"KE.22":       "KE-13",  // Kiambu County
	"KE.23":       "KE-14",  // Kilifi County
	"KE.24":       "KE-15",  // Kirinyaga County
	"KE.25":       "KE-16",  // Kisii County
	"KE.26":       "KE-17",  // Kisumu County
	"KE.27":       "KE-18",  // Kitui County
	"KE.28":       "KE-19",  // Kwale County
	"KE.29":       "KE-20",  // Laikipia
	"KE.30":       "KE-21",  // Lamu
	"KE.31":       "KE-22",  // Machakos County
	"KE.32":       "KE-23",  // Makueni County
	"KE.33":       "KE-24",  // Mandera County
	"KE.34":       "KE-25",  // Marsabit County
	"KE.35":       "KE-26",  // Meru County
	"KE.36":       "KE-27",  // Migori County
	"KE.37":       "KE-28",  // Mombasa County
	"KE.38":       "KE-29",  // Murang'A
	"KE.39":       "KE-31",  // Nakuru County
	"KE.40":       "KE-32",  // Nandi
	"KE.41":       "KE-33",  // Narok County
	"KE.42":       "KE-34",  // Nyamira county
	"KE.43":       "KE-35",  // Nyandarua County
	"KE.44":       "KE-36",  // Nyeri County
	"KE.45":       "KE-37",  // Samburu County
	"KE.46":       "KE-38",  // Siaya County
	"KE.47":       "KE-39",  // Taita Taveta
	"KE.48":       "KE-40",  // Tana River County
	"KE.49":       "KE-41",  // Tharaka - Nithi
	"KE.50":       "KE-42",  // Trans Nzoia
	"KE.51":       "KE-43",  // Turkana County
	"KE.52":       "KE-44",  // Uasin Gishu County
	"KE.53":       "KE-45",  // Vihiga County
	"KE.54":       "KE-46",  // Wajir County
	"KE.55":       "KE-47",  // West Pokot County
	"KG.01":       "KG-GB",  // Bishkek
	"KG.02":       "KG-C",   // Chuy Region
	"KG.03":       "KG-J",   // Jalal-Abad Region
	"KG.04":       "KG-N",   // Naryn Region
	"KG.06":       "KG-T",   // Talas Region
	"KG.07":       "KG-Y",   // Issyk-Kul
	"KG.08":       "KG-O",   // Osh Region
	"KG.09":       "KG-B",   // Batken
	"KG.10":       "KG-GO",  // Osh City
	"KH.02":       "KH-3",   // Kampong Cham
	"KH.03":       "KH-4",   // Kampong Chhnang
	"KH.04":       "KH-5",   // Kampong Speu Province
	"KH.05":       "KH-6",   // Kampong Thom
	"KH.07":       "KH-8",   // Kandal
	"KH.08":       "KH-9",   // Koh Kong
	"KH.09":       "KH-10",  // Kratie
	"KH.10":       "KH-11",  // Mondolkiri
	"KH.12":       "KH-15",  // Pursat
	"KH.13":       "KH-13",  // Preah Vihear
	"KH.14":       "KH-14",  // Prey Veng
	"KH.17":       "KH-19",  // Stung Treng
	"KH.18":       "KH-20",  // Svay Rieng
	"KH.19":       "KH-21",  // Takeo
	"KH.21":       "KH-7",   // Kampot
	"KH.22":       "KH-12",  // Phnom Penh
	"KH.23":       "KH-16",  // Ratanakiri
	"KH.24":       "KH-17",  // Siem Reap
	"KH.25":       "KH-1",   // Banteay Meanchey
	"KH.26":       "KH-23",  // Kep
	"KH.27":       "KH-22",  // Ŏtâr Méanchey
	"KH.29":       "KH-2",   // Battambang
	"KH.30":       "KH-24",  // Pailin
	"KH.31":       "KH-25",  // Tboung Khmum
	"KI.01":       "KI-G",   // Gilbert Islands
	"KI.02":       "KI-L",   // Line Islands
	"KI.03":       "KI-P",   // Phoenix Islands
	"KM.01":       "KM-A",   // Anjouan
	"KM.02":       "KM-G",   // Grande Comore
	"KM.03":       "KM-M",   // Mohéli
	"KN.01":       "KN-K",   // Christ Church Nichola Town
	"KN.02":       "KN-K",   // Saint Anne Sandy Point
	"KN.03":       "KN-K",   // Saint George Basseterre
	"KN.04":       "KN-N",   // Saint George Gingerland
	"KN.05":       "KN-N",   // Saint James Windwa
	"KN.06":       "KN-K",   // Saint John Capesterre
	"KN.07":       "KN-N",   // Saint John Figtree
	"KN.08":       "KN-K",   // Saint Mary Cayon
	"KN.09":       "KN-K",   // Saint Paul Capesterre
	"KN.10":       "KN-N",   // Saint Paul Charlestown
	"KN.11":       "KN-K",   // Saint Peter Basseterre
	"KN.12":       "KN-N",   // Saint Thomas Lowland
	"KN.13":       "KN-K",   // Middle Island
	"KN.15":       "KN-K",   // Trinity Palmetto Point
	"KP.01":       "KP-04",  // Chagang
	"KP.03":       "KP-08",  // South Hamgyong
	"KP.06":       "KP-05",  // South Hwanghae
	"KP.07":       "KP-06",  // North Hwanghae
	"KP.09":       "KP-07",  // Kangwŏn-do
	"KP.11":       "KP-02",  // P'yŏngan-bukto
	"KP.12":       "KP-01",  // Pyongyang
	"KP.13":       "KP-10",  // Yanggang-do
	"KP.14":       "KP-14",  // Nampo
	"KP.15":       "KP-02",  // South Pyongan
	"KP.17":       "KP-09",  // North Hamgyong
	"KP.18":       "KP-13",  // Rason
	"KR.01":       "KR-49",  // Jeju-do
	"KR.03":       "KR-45",  // Jeollabuk-do
	"KR.05":       "KR-43",  // North Chungcheong
	"KR.06":       "KR-42",  // Gangwon-do
	"KR.10":       "KR-26",  // Busan
	"KR.11":       "KR-11",  // Seoul
	"KR.12":       "KR-28",  // Incheon
	"KR.13":       "KR-41",  // Gyeonggi-do
	"KR.14":       "KR-47",  // Gyeongsangbuk-do
	"KR.15":       "KR-27",  // Daegu
	"KR.16":       "KR-46",  // Jeollanam-do
	"KR.17":       "KR-44",  // Chungcheongnam-do
	"KR.18":       "KR-29",  // Gwangju
	"KR.19":       "KR-30",  // Daejeon
	"KR.20":       "KR-48",  // Gyeongsangnam-do
	"KR.21":       "KR-31",  // Ulsan
	"KR.22":       "KR-50",  // Sejong-si
	"KW.02":       "KW-KU",  // Al Asimah
	"KW.04":       "KW-AH",  // Al Aḩmadī
	"KW.05":       "KW-JA",  // Al Jahra Governorate
	"KW.07":       "KW-FA",  // Al Farwaniyah
	"KW.08":       "KW-HA",  // Hawalli
	"KW.09":       "KW-MU",  // Mubārak al Kabīr
	"KZ.01":       "KZ-ALM", // Almaty Region
	"KZ.02":       "KZ-ALA", // Almaty
	"KZ.03":       "KZ-AKM", // Aqmola
	"KZ.04":       "KZ-AKT", // Aktyubinskaya Oblast’
	"KZ.05":       "KZ-AST", // Astana
	"KZ.06":       "KZ-ATY", // Atyrau Oblisi
	"KZ.07":       "KZ-ZAP", // Batys Qazaqstan
	"KZ.09":       "KZ-MAN", // Mangghystaū
	"KZ.11":       "KZ-PAV", // Pavlodar Region
	"KZ.12":       "KZ-KAR", // Karaganda
	"KZ.13":       "KZ-KUS", // Qostanay
	"KZ.14":       "KZ-KZY", // Qyzylorda
	"KZ.15":       "KZ-VOS", // East Kazakhstan
	"KZ.1537272":  "KZ-SHY", // Shymkent
	"KZ.16":       "KZ-SEV", // North Kazakhstan
	"KZ.17":       "KZ-ZHA", // Zhambyl
	"LA.02":       "LA-CH",  // Champasak Province
	"LA.03":       "LA-HO",  // Houaphan
	"LA.07":       "LA-OU",  // Oudômxai
	"LA.13":       "LA-XA",  // Xiagnabouli
	"LA.14":       "LA-XI",  // Xiangkhoang
	"LA.15":       "LA-KH",  // Khammouane Province
	"LA.16":       "LA-LM",  // Loungnamtha
	"LA.17":       "LA-LP",  // Luang Prabang Province
	"LA.18":       "LA-PH",  // Phôngsali
	"LA.19":       "LA-SL",  // Salavan Province
	"LA.20":       "LA-SV",  // Savannakhet Province
	"LA.22":       "LA-BK",  // Bokeo
	"LA.23":       "LA-BL",  // Bolikhamsai
	"LA.24":       "LA-VT",  // Vientiane Prefecture
	"LA.26":       "LA-XE",  // Xékong
	"LA.27":       "LA-VT",  // Vientiane
	"LA.28":       "LA-XS",  // Xaisomboun
	"LB.04":       "LB-BA",  // Beyrouth
	"LB.05":       "LB-JL",  // Mont-Liban
	"LB.06":       "LB-JA",  // South Governorate
	"LB.07":       "LB-NA",  // Nabatîyé
	"LB.08":       "LB-BI",  // Béqaa
	"LB.10":       "LB-AK",  // Aakkâr
	"LB.11":       "LB-BH",  // Baalbek-Hermel Governorate
	"LC.01":       "LC-01",  // Anse-la-Raye
	"LC.03":       "LC-02",  // Castries
	"LC.04":       "LC-03",  // Choiseul
	"LC.05":       "LC-05",  // Dennery
	"LC.06":       "LC-06",  // Gros-Islet
	"LC.07":       "LC-07",  // Laborie
	"LC.08":       "LC-08",  // Micoud
	"LC.09":       "LC-10",  // Soufrière
	"LC.10":       "LC-11",  // Vieux-Fort
	"LC.12":       "LC-12",  // Canaries
	"LI.01":       "LI-01",  // Balzers
	"LI.02":       "LI-02",  // Eschen
	"LI.03":       "LI-03",  // Gamprin
	"LI.04":       "LI-04",  // Mauren
	"LI.05":       "LI-05",  // Planken
	"LI.06":       "LI-06",  // Ruggell
	"LI.07":       "LI-07",  // Schaan
	"LI.08":       "LI-08",  // Schellenberg
	"LI.09":       "LI-09",  // Triesen
	"LI.10":       "LI-10",  // Triesenberg
	"LI.11":       "LI-11",  // Vaduz
	"LK.29":       "LK-2",   // Central Province
	"LK.30":       "LK-7",   // North Central Province
	"LK.32":       "LK-6",   // North Western Province
	"LK.33":       "LK-9",   // Sabaragamuwa Province
	"LK.34":       "LK-3",   // Southern Province
	"LK.35":       "LK-8",   // Uva Province
	"LK.36":       "LK-1",   // Western Province
	"LK.37":       "LK-5",   // Eastern Province
	"LK.38":       "LK-4",   // Northern Province
	"LR.01":       "LR-BG",  // Bong County
	"LR.09":       "LR-NI",  // Nimba
	"LR.10":       "LR-SI",  // Sinoe County
	"LR.11":       "LR-GB",  // Grand Bassa County
	"LR.12":       "LR-CM",  // Grand Cape Mount County
	"LR.13":       "LR-MY",  // Maryland County
	"LR.14":       "LR-MO",  // Montserrado County
	"LR.15":       "LR-BM",  // Bomi County
	"LR.16":       "LR-GK",  // Grand Kru County
	"LR.17":       "LR-MG",  // Margibi County
	"LR.18":       "LR-RI",  // River Cess County
	"LR.19":       "LR-GG",  // Grand Gedeh County
	"LR.20":       "LR-LO",  // Lofa County
	"LR.21":       "LR-GP",  // Gbarpolu County
	"LR.22":       "LR-RG",  // River Gee County
	"LS.10":       "LS-D",   // Berea
	"LS.11":       "LS-B",   // Butha-Buthe
	"LS.12":       "LS-C",   // Leribe
	"LS.13":       "LS-E",   // Mafeteng District
	"LS.14":       "LS-A",   // Maseru District
	"LS.15":       "LS-F",   // Mohale's Hoek District
	"LS.16":       "LS-J",   // Mokhotlong District
	"LS.17":       "LS-H",   // Qacha's Nek District
	"LS.18":       "LS-G",   // Quthing
	"LS.19":       "LS-K",   // Thaba-Tseka
	"LT.56":       "LT-AL",  // Alytus
	"LT.57":       "LT-KU",  // Kaunas
	"LT.58":       "LT-KL",  // Klaipėda County
	"LT.59":       "LT-MR",  // Marijampolė County
	"LT.60":       "LT-PN",  // Panevėžys
	"LT.61":       "LT-SA",  // Siauliai
	"LT.62":       "LT-TA",  // Tauragė County
	"LT.63":       "LT-TE",  // Telsiai
	"LT.64":       "LT-UT",  // Utena
	"LT.65":       "LT-VL",  // Vilnius
	"LU.CA":       "LU-CA",  // Capellen
	"LU.CL":       "LU-CL",  // Clervaux
	"LU.DI":       "LU-DI",  // Diekirch
	"LU.EC":       "LU-EC",  // Echternach
	"LU.ES":       "LU-ES",  // Esch-sur-Alzette
	"LU.GR":       "LU-GR",  // Grevenmacher
	"LU.LU":       "LU-LU",  // Luxembourg
	"LU.ME":       "LU-ME",  // Mersch
	"LU.RD":       "LU-RD",  // Redange
	"LU.RM":       "LU-RM",  // Remich
	"LU.VD":       "LU-VD",  // Vianden
	"LU.WI":       "LU-WI",  // Wiltz
	"LV.01":       "LV-002", // Aizkraukle Municipality
	"LV.02":       "LV-007", // Alūksne Municipality
	"LV.03":       "LV-015", // Balvi Municipality
	"LV.04":       "LV-016", // Bauska Municipality
	"LV.05":       "LV-022", // Cēsis Municipality
	"LV.06":       "LV-DGV", // Daugavpils
	"LV.08":       "LV-026", // Dobele Municipality
	"LV.09":       "LV-033", // Gulbene Municipality
	"LV.10":       "LV-042", // Jēkabpils Municipality
	"LV.11":       "LV-JEL", // Jelgava
	"LV.12":       "LV-041", // Jelgava Municipality
	"LV.13":       "LV-JUR", // Jūrmala
	"LV.14":       "LV-047", // Krāslava Municipality
	"LV.15":       "LV-050", // Kuldīga Municipality
	"LV.16":       "LV-LPX", // Liepāja
	"LV.18":       "LV-054", // Limbaži Municipality
	"LV.19":       "LV-058", // Ludza Municipality
	"LV.20":       "LV-059", // Madonas novads
	"LV.21":       "LV-067", // Ogre
	"LV.22":       "LV-073", // Preiļu novads
	"LV.23":       "LV-REZ", // Rēzekne
	"LV.24":       "LV-077", // Rēzekne Municipality
	"LV.25":       "LV-RIX", // Riga
	"LV.27":       "LV-088", // Saldus Rajons
	"LV.28":       "LV-097", // Talsu novads
	"LV.29":       "LV-099", // Tukums Municipality
	"LV.30":       "LV-101", // Valka
	"LV.31":       "LV-113", // Valmiera
	"LV.32":       "LV-VEN", // Ventspils
	"LV.33":       "LV-106", // Ventspils Municipality
	"LV.34":       "LV-011", // Ādaži
	"LV.80":       "LV-052", // Ķekava
	"LV.90":       "LV-056", // Līvāni
	"LV.95":       "LV-062", // Mārupe
	"LV.A2":       "LV-068", // Olaine
	"LV.AN":       "LV-111", // Augšdaugava Municipality
	"LV.B5":       "LV-080", // Ropaži Municipality
	"LV.C3":       "LV-087", // Salaspils Municipality
	"LV.C5":       "LV-089", // Saulkrasti Municipality
	"LV.C7":       "LV-091", // Sigulda Municipality
	"LV.D1":       "LV-094", // Smiltene Municipality
	"LV.DN":       "LV-112", // South Kurzeme Municipality
	"LV.E1":       "LV-102", // Varakļāni Municipality
	"LY.63":       "LY-JA",  // Al Jabal al Akhḑar
	"LY.64":       "LY-JU",  // Al Jufrah
	"LY.65":       "LY-KF",  // Al Kufrah
	"LY.66":       "LY-MJ",  // Al Marj
	"LY.67":       "LY-NQ",  // An Nuqāţ al Khams
	"LY.68":       "LY-ZA",  // Az Zāwiyah
	"LY.69":       "LY-BA",  // Banghāzī
	"LY.70":       "LY-DR",  // Darnah
	"LY.71":       "LY-GT",  // Ghāt
	"LY.72":       "LY-MI",  // Mişrātah
	"LY.73":       "LY-MQ",  // Murzuq District
	"LY.74":       "LY-NL",  // Nālūt
	"LY.75":       "LY-SB",  // Sabha District
	"LY.76":       "LY-SR",  // Surt
	"LY.77":       "LY-TB",  // Tripoli
	"LY.78":       "LY-WS",  // Ash Shāţiʼ
	"LY.79":       "LY-BU",  // Al Buţnān
	"LY.80":       "LY-JG",  // Jabal al Gharbi
	"LY.81":       "LY-JI",  // Al Jafārah
	"LY.82":       "LY-MB",  // Al Marqab
	"LY.83":       "LY-WA",  // Al Wāḩāt
	"LY.84":       "LY-WD",  // Wādī al Ḩayāt
	"MA.01":       "MA-01",  // Tanger-Tetouan-Al Hoceima
	"MA.02":       "MA-04",  // Oriental
	"MA.06":       "MA-06",  // Casablanca-Settat
	"MA.07":       "MA-07",  // Marrakesh-Safi
	"MA.11":       "MA-11",  // Laâyoune-Sakia El Hamra
	"MA.12":       "MA-12",  // Dakhla-Oued Ed-Dahab
	"MD.51":       "MD-GA",  // Gagauzia
	"MD.57":       "MD-CU",  // Chișinău Municipality
	"MD.58":       "MD-SN",  // Transnistria
	"MD.59":       "MD-AN",  // Anenii Noi
	"MD.60":       "MD-BA",  // Bălţi
	"MD.61":       "MD-BS",  // Basarabeasca
	"MD.62":       "MD-BD",  // Bender Municipality
	"MD.63":       "MD-BR",  // Briceni
	"MD.64":       "MD-CA",  // Cahul
	"MD.65":       "MD-CT",  // Cantemir
	"MD.66":       "MD-CL",  // Călăraşi
	"MD.67":       "MD-CS",  // Căuşeni
	"MD.68":       "MD-CM",  // Cimişlia
	"MD.69":       "MD-CR",  // Criuleni
	"MD.70":       "MD-DO",  // Donduşeni
	"MD.71":       "MD-DR",  // Drochia
	"MD.72":       "MD-DU",  // Dubăsari District
	"MD.73":       "MD-ED",  // Raionul Edineţ
	"MD.74":       "MD-FA",  // Fălești
	"MD.75":       "MD-FL",  // Floreşti
	"MD.76":       "MD-GL",  // Glodeni
	"MD.77":       "MD-HI",  // Hînceşti
	"MD.78":       "MD-IA",  // Ialoveni
	"MD.79":       "MD-LE",  // Leova
	"MD.80":       "MD-NI",  // Nisporeni
	"MD.81":       "MD-OC",  // Raionul Ocniţa
	"MD.82":       "MD-OR",  // Orhei
	"MD.83":       "MD-RE",  // Rezina
	"MD.84":       "MD-RI",  // Rîşcani
	"MD.85":       "MD-SI",  // Sîngerei
	"MD.86":       "MD-SD",  // Şoldăneşti
	"MD.87":       "MD-SO",  // Soroca District
	"MD.88":       "MD-SV",  // Ştefan-Vodă
	"MD.89":       "MD-ST",  // Strășeni
	"MD.90":       "MD-TA",  // Taraclia
	"MD.91":       "MD-TE",  // Teleneşti
	"MD.92":       "MD-UN",  // Ungheni
	"ME.01":       "ME-01",  // Andrijevica
	"ME.02":       "ME-02",  // Bar
	"ME.03":       "ME-03",  // Berane
	"ME.04":       "ME-04",  // Bijelo Polje
	"ME.05":       "ME-05",  // Budva
	"ME.06":       "ME-06",  // Cetinje
	"ME.07":       "ME-07",  // Danilovgrad
	"ME.08":       "ME-08",  // Herceg Novi
	"ME.09":       "ME-09",  // Opština Kolašin
	"ME.10":       "ME-10",  // Kotor
	"ME.11":       "ME-11",  // Mojkovac
	"ME.12":       "ME-12",  // Opština Nikšić
	"ME.13":       "ME-13",  // Opština Plav
	"ME.14":       "ME-14",  // Pljevlja
	"ME.15":       "ME-15",  // Opština Plužine
	"ME.16":       "ME-16",  // Podgorica
	"ME.17":       "ME-17",  // Rožaje Municipality
	"ME.18":       "ME-18",  // Opština Šavnik
	"ME.19":       "ME-19",  // Tivat
	"ME.20":       "ME-20",  // Ulcinj
	"ME.21":       "ME-21",  // Opština Žabljak
	"ME.22":       "ME-22",  // Gusinje Municipality
	"ME.23":       "ME-23",  // Petnjica Municipality
	"ME.24":       "ME-24",  // Tuzi Municipality
	"MG.11":       "MG-T",   // Analamanga
	"MG.21":       "MG-F",   // Upper Matsiatra
	"MG.26":       "MG-F",   // Fitovinany Region
	"MG.41":       "MG-M",   // Boeny
	"MG.51":       "MG-U",   // Atsimo-Andrefana
	"MG.71":       "MG-D",   // Diana
	"MH.010":      "MH-L",   // Ailinglaplap Atoll
	"MH.030":      "MH-T",   // Ailuk Atoll
	"MH.040":      "MH-T",   // Arno Atoll
	"MH.050":      "MH-T",   // Aur Atoll
	"MH.080":      "MH-L",   // Ebon Atoll
	"MH.090":      "MH-L",   // Enewetak Atoll
	"MH.110":      "MH-L",   // Jabat Island
	"MH.120":      "MH-L",   // Jaluit Atoll
	"MH.140":      "MH-L",   // Kili Island
	"MH.150":      "MH-L",   // Kwajalein Atoll
	"MH.160":      "MH-L",   // Lae Atoll
	"MH.170":      "MH-L",   // Lib Island
	"MH.180":      "MH-T",   // Likiep Atoll
	"MH.190":      "MH-T",   // Majuro Atoll
	"MH.300":      "MH-T",   // Maloelap Atoll
	"MH.310":      "MH-T",   // Mejit Island
	"MH.320":      "MH-T",   // Mili Atoll
	"MH.330":      "MH-L",   // Namdrik Atoll
	"MH.340":      "MH-L",   // Namu Atoll
	"MH.350":      "MH-L",   // Rongelap Atoll
	"MH.390":      "MH-L",   // Ujae Atoll
	"MH.410":      "MH-T",   // Utrik Atoll
	"MH.420":      "MH-L",   // Wotho Atoll
	"MH.430":      "MH-T",   // Wotje Atoll
	"MK.01":       "MK-802", // Arachinovo
	"MK.04":       "MK-201", // Berovo
	"MK.06":       "MK-501", // Bitola
	"MK.08":       "MK-401", // Bogdanci
	"MK.11":       "MK-402", // Bosilovo
	"MK.12":       "MK-602", // Brvenica
	"MK.18":       "MK-313", // Centar Zhupa
	"MK.19":       "MK-210", // Češinovo-Obleševo
	"MK.20":       "MK-816", // Chucher Sandevo
	"MK.22":       "MK-203", // Delchevo
	"MK.25":       "MK-103", // Demir Kapija
	"MK.28":       "MK-503", // Dolneni
	"MK.33":       "MK-405", // Gevgelija
	"MK.35":       "MK-102", // Gradsko
	"MK.36":       "MK-807", // Ilinden
	"MK.40":       "MK-205", // Karbinci
	"MK.43":       "MK-307", // Kichevo
	"MK.46":       "MK-206", // Kochani
	"MK.47":       "MK-407", // Konche
	"MK.51":       "MK-701", // Kratovo
	"MK.52":       "MK-702", // Kriva Palanka
	"MK.53":       "MK-504", // Krivogashtani
	"MK.54":       "MK-505", // Krushevo
	"MK.59":       "MK-704", // Lipkovo
	"MK.60":       "MK-105", // Lozovo
	"MK.62":       "MK-207", // Makedonska Kamenica
	"MK.69":       "MK-106", // Negotino
	"MK.72":       "MK-408", // Novo Selo
	"MK.78":       "MK-208", // Pehchevo
	"MK.79":       "MK-810", // Petrovec
	"MK.80":       "MK-311", // Plasnica
	"MK.83":       "MK-209", // Probishtip
	"MK.85":       "MK-705", // Rankovce
	"MK.86":       "MK-509", // Resen
	"MK.87":       "MK-107", // Rosoman
	"MK.92":       "MK-812", // Sopište
	"MK.97":       "MK-706", // Staro Nagorichane
	"MK.98":       "MK-211", // Shtip
	"MK.A2":       "MK-813", // Studenichani
	"MK.A4":       "MK-108", // Sveti Nikole
	"MK.A5":       "MK-608", // Tearce
	"MK.A9":       "MK-404", // Vasilevo
	"MK.B3":       "MK-301", // Vevchani
	"MK.B4":       "MK-202", // Vinica
	"MK.B7":       "MK-603", // Vrapchishte
	"MK.C2":       "MK-806", // Zelenikovo
	"MK.C3":       "MK-605", // Zhelino
	"MK.C6":       "MK-204", // Zrnovci
	"MK.C7":       "MK-601", // Bogovinje
	"MK.D2":       "MK-303", // Debar
	"MK.D3":       "MK-502", // Demir Hisar
	"MK.D4":       "MK-604", // Gostivar
	"MK.D5":       "MK-606", // Jegunovce
	"MK.D6":       "MK-104", // Kavadarci
	"MK.D7":       "MK-703", // Kumanovo
	"MK.D8":       "MK-308", // Makedonski Brod
	"MK.D9":       "MK-506", // Mogila
	"MK.E1":       "MK-507", // Novaci
	"MK.E2":       "MK-310", // Ohrid
	"MK.E3":       "MK-508", // Prilep
	"MK.E4":       "MK-607", // Mavrovo and Rostuša
	"MK.E5":       "MK-406", // Dojran
	"MK.E6":       "MK-312", // Struga
	"MK.E7":       "MK-410", // Strumica
	"MK.E8":       "MK-609", // Tetovo
	"MK.E9":       "MK-403", // Valandovo
	"MK.F1":       "MK-101", // Veles
	"ML.01":       "ML-BKO", // Bamako
	"ML.03":       "ML-1",   // Kayes
	"ML.04":       "ML-5",   // Mopti
	"ML.05":       "ML-4",   // Ségou
	"ML.06":       "ML-3",   // Sikasso
	"ML.07":       "ML-2",   // Koulikoro
	"ML.08":       "ML-6",   // Tombouctou
	"ML.09":       "ML-7",   // Gao
	"ML.10":       "ML-8",   // Kidal
	"ML.12070575": "ML-10",  // Taoudénit
	"ML.12070577": "ML-9",   // Ménaka
	"MM.01":       "MM-16",  // Rakhine
	"MM.02":       "MM-14",  // Chin State
	"MM.03":       "MM-07",  // Ayeyarwady
	"MM.04":       "MM-11",  // Kachin State
	"MM.05":       "MM-13",  // Kayin State
	"MM.06":       "MM-12",  // Kayah State
	"MM.08":       "MM-04",  // Mandalay Region
	"MM.10":       "MM-01",  // Sagaing Region
	"MM.11":       "MM-17",  // Shan State
	"MM.12":       "MM-05",  // Tanintharyi Region
	"MM.13":       "MM-15",  // Mon
	"MM.15":       "MM-03",  // Magway
	"MM.16":       "MM-02",  // Bago Region
	"MM.17":       "MM-06",  // Yangon
	"MM.18":       "MM-18",  // Nay Pyi Taw
	"MN.01":       "MN-073", // Arkhangai Province
	"MN.02":       "MN-069", // Bayanhongor
	"MN.03":       "MN-071", // Bayan-Ölgii Province
	"MN.10":       "MN-065", // Govi-Altai Province
	"MN.11":       "MN-039", // Hentiy
	"MN.12":       "MN-043", // Hovd
	"MN.13":       "MN-041", // Khövsgöl Province
	"MN.14":       "MN-053", // Ömnögovĭ
	"MN.15":       "MN-055", // Övörhangay
	"MN.16":       "MN-049", // Selenge Province
	"MN.17":       "MN-051", // Sühbaatar
	"MN.19":       "MN-046", // Uvs Province
	"MN.20":       "MN-1",   // Ulaanbaatar
	"MN.21":       "MN-067", // Bulgan
	"MN.23":       "MN-037", // Darhan Uul
	"MN.24":       "MN-064", // Govĭ-Sumber
	"MN.25":       "MN-035", // Orhon
	"MR.02":       "MR-02",  // Hodh El Gharbi
	"MR.03":       "MR-03",  // Assaba
	"MR.04":       "MR-04",  // Gorgol
	"MR.05":       "MR-05",  // Brakna
	"MR.06":       "MR-06",  // Trarza
	"MR.07":       "MR-07",  // Adrar
	"MR.08":       "MR-08",  // Dakhlet Nouadhibou
	"MR.09":       "MR-09",  // Tagant
	"MR.10":       "MR-10",  // Guidimaka
	"MR.11":       "MR-11",  // Tiris Zemmour
	"MR.12":       "MR-12",  // Inchiri
	"MR.13":       "MR-13",  // Nouakchott Ouest
	"MR.14":       "MR-14",  // Nouakchott Nord
	"MR.15":       "MR-15",  // Nouakchott Sud
	"MT.01":       "MT-01",  // Attard
	"MT.02":       "MT-02",  // Balzan
	"MT.03":       "MT-03",  // Il-Birgu
	"MT.04":       "MT-04",  // Birkirkara
	"MT.05":       "MT-05",  // Birżebbuġa
	"MT.06":       "MT-06",  // Bormla
	"MT.07":       "MT-07",  // Dingli
	"MT.08":       "MT-08",  // Il-Fgura
	"MT.09":       "MT-09",  // Floriana
	"MT.10":       "MT-10",  // Il-Fontana
	"MT.11":       "MT-13",  // Għajnsielem
	"MT.12":       "MT-14",  // L-Għarb
	"MT.13":       "MT-15",  // Ħal Għargħur
	"MT.14":       "MT-16",  // L-Għasri
	"MT.15":       "MT-17",  // Ħal Għaxaq
	"MT.16":       "MT-11",  // Il-Gudja
	"MT.17":       "MT-12",  // Il-Gżira
	"MT.18":       "MT-18",  // Il-Ħamrun
	"MT.19":       "MT-19",  // L-Iklin
	"MT.20":       "MT-29",  // L-Imdina
	"MT.21":       "MT-31",  // L-Imġarr
	"MT.22":       "MT-33",  // L-Imqabba
	"MT.23":       "MT-34",  // L-Imsida
	"MT.24":       "MT-35",  // Mtarfa
	"MT.25":       "MT-20",  // Senglea
	"MT.26":       "MT-21",  // Il-Kalkara
	"MT.27":       "MT-22",  // Ta’ Kerċem
	"MT.28":       "MT-23",  // Kirkop
	"MT.29":       "MT-24",  // Lija
	"MT.30":       "MT-25",  // Luqa
	"MT.31":       "MT-26",  // Il-Marsa
	"MT.32":       "MT-27",  // Marsaskala
	"MT.33":       "MT-28",  // Marsaxlokk
	"MT.34":       "MT-30",  // Il-Mellieħa
	"MT.35":       "MT-32",  // Il-Mosta
	"MT.36":       "MT-36",  // Il-Munxar
	"MT.37":       "MT-37",  // In-Nadur
	"MT.38":       "MT-38",  // In-Naxxar
	"MT.39":       "MT-39",  // Paola
	"MT.40":       "MT-40",  // Pembroke
	"MT.41":       "MT-41",  // Tal-Pietà
	"MT.42":       "MT-42",  // Il-Qala
	"MT.43":       "MT-43",  // Qormi
	"MT.44":       "MT-44",  // Il-Qrendi
	"MT.45":       "MT-46",  // Ir-Rabat
	"MT.46":       "MT-45",  // Victoria
	"MT.47":       "MT-47",  // Safi
	"MT.49":       "MT-48",  // Saint Julian
	"MT.50":       "MT-50",  // Saint Lawrence
	"MT.51":       "MT-53",  // Saint Lucia
	"MT.52":       "MT-51",  // Saint Paul’s Bay
	"MT.53":       "MT-54",  // Saint Venera
	"MT.54":       "MT-52",  // Sannat
	"MT.55":       "MT-55",  // Is-Siġġiewi
	"MT.56":       "MT-56",  // Tas-Sliema
	"MT.57":       "MT-57",  // Is-Swieqi
	"MT.58":       "MT-59",  // Tarxien
	"MT.59":       "MT-58",  // Ta’ Xbiex
	"MT.60":       "MT-60",  // Valletta
	"MT.61":       "MT-61",  // Ix-Xagħra
	"MT.62":       "MT-62",  // Ix-Xewkija
	"MT.63":       "MT-63",  // Ix-Xgħajra
	"MT.64":       "MT-64",  // Ħaż-Żabbar
	"MT.65":       "MT-66",  // Ħaż-Żebbuġ
	"MT.66":       "MT-66",  // Iż-Żebbuġ
	"MT.67":       "MT-67",  // Iż-Żejtun
	"MT.68":       "MT-68",  // Iż-Żurrieq
	"MU.12":       "MU-BL",  // Black River
	"MU.13":       "MU-FL",  // Flacq
	"MU.14":       "MU-GP",  // Grand Port
	"MU.15":       "MU-MO",  // Moka
	"MU.16":       "MU-PA",  // Pamplemousses
	"MU.17":       "MU-PW",  // Plaines Wilhems
	"MU.18":       "MU-PL",  // Port Louis
	"MU.19":       "MU-RR",  // Rivière du Rempart
	"MU.20":       "MU-SA",  // Savanne
	"MU.21":       "MU-AG",  // Agalega Islands
	"MU.22":       "MU-CC",  // Cargados Carajos
	"MU.23":       "MU-RO",  // Rodrigues
	"MV.05":       "MV-05",  // Laamu
	"MV.31":       "MV-20",  // Baa Atholhu
	"MV.32":       "MV-17",  // Dhaalu Atholhu
	"MV.33":       "MV-14",  // Faafu Atholhu
	"MV.34":       "MV-27",  // Gaafu Alif Atoll
	"MV.35":       "MV-28",  // Gaafu Dhaalu Atoll
	"MV.36":       "MV-07",  // Haa Alifu Atholhu
	"MV.37":       "MV-23",  // Haa Dhaalu Atholhu
	"MV.38":       "MV-26",  // Kaafu Atoll
	"MV.40":       "MV-MLE", // Male
	"MV.41":       "MV-12",  // Meemu Atholhu
	"MV.42":       "MV-29",  // Gnyaviyani Atoll
	"MV.43":       "MV-25",  // Noonu Atoll
	"MV.44":       "MV-13",  // Raa Atoll
	"MV.45":       "MV-24",  // Shaviyani Atholhu
	"MV.46":       "MV-08",  // Thaa Atholhu
	"MV.47":       "MV-04",  // Vaavu Atholhu
	"MW.C":        "MW-C",   // Central Region
	"MW.N":        "MW-N",   // Northern Region
	"MW.S":        "MW-S",   // Southern Region
	"MX.01":       "MX-AGU", // Aguascalientes
	"MX.02":       "MX-BCN", // Baja California
	"MX.03":       "MX-BCS", // Baja California Sur
	"MX.04":       "MX-CAM", // Campeche
	"MX.05":       "MX-CHP", // Chiapas
	"MX.06":       "MX-CHH", // Chihuahua
	"MX.07":       "MX-COA", // Coahuila
	"MX.08":       "MX-COL", // Colima
	"MX.09":       "MX-CMX", // Mexico City
	"MX.10":       "MX-DUR", // Durango
	"MX.11":       "MX-GUA", // Guanajuato
	"MX.12":       "MX-GRO", // Guerrero
	"MX.13":       "MX-HID", // Hidalgo
	"MX.14":       "MX-JAL", // Jalisco
	"MX.15":       "MX-MEX", // México
	"MX.16":       "MX-MIC", // Michoacán
	"MX.17":       "MX-MOR", // Morelos
	"MX.18":       "MX-NAY", // Nayarit
	"MX.19":       "MX-NLE", // Nuevo León
	"MX.20":       "MX-OAX", // Oaxaca
	"MX.21":       "MX-PUE", // Puebla
	"MX.22":       "MX-QUE", // Querétaro
	"MX.23":       "MX-ROO", // Quintana Roo
	"MX.24":       "MX-SLP", // San Luis Potosí
	"MX.25":       "MX-SIN", // Sinaloa
	"MX.26":       "MX-SON", // Sonora
	"MX.27":       "MX-TAB", // Tabasco
	"MX.28":       "MX-TAM", // Tamaulipas
	"MX.29":       "MX-TLA", // Tlaxcala
	"MX.30":       "MX-VER", // Veracruz
	"MX.31":       "MX-YUC", // Yucatán
	"MX.32":       "MX-ZAC", // Zacatecas
	"MY.01":       "MY-01",  // Johor
	"MY.02":       "MY-02",  // Kedah
	"MY.03":       "MY-03",  // Kelantan
	"MY.04":       "MY-04",  // Melaka
	"MY.05":       "MY-05",  // Negeri Sembilan
	"MY.06":       "MY-06",  // Pahang
	"MY.07":       "MY-08",  // Perak
	"MY.08":       "MY-09",  // Perlis
	"MY.09":       "MY-07",  // Penang
	"MY.11":       "MY-13",  // Sarawak
	"MY.12":       "MY-10",  // Selangor
	"MY.13":       "MY-11",  // Terengganu
	"MY.14":       "MY-14",  // Kuala Lumpur
	"MY.15":       "MY-15",  // Labuan
	"MY.16":       "MY-12",  // Sabah
	"MY.17":       "MY-16",  // Putrajaya
	"MZ.01":       "MZ-P",   // Cabo Delgado Province
	"MZ.02":       "MZ-G",   // Gaza Province
	"MZ.03":       "MZ-I",   // Inhambane Province
	"MZ.04":       "MZ-L",   // Maputo Province
	"MZ.05":       "MZ-S",   // Sofala
	"MZ.06":       "MZ-N",   // Nampula
	"MZ.07":       "MZ-A",   // Niassa Province
	"MZ.08":       "MZ-T",   // Tete
	"MZ.09":       "MZ-Q",   // Zambezia Province
	"MZ.10":       "MZ-B",   // Manica
	"MZ.11":       "MZ-MPM", // Maputo City
	"NA.21":       "NA-KH",  // Khomas Region
	"NA.28":       "NA-CA",  // Zambezi Region
	"NA.29":       "NA-ER",  // Erongo Region
	"NA.30":       "NA-HA",  // Hardap Region
	"NA.31":       "NA-KA",  // Karas Region
	"NA.32":       "NA-KU",  // Kunene Region
	"NA.33":       "NA-OW",  // Ohangwena Region
	"NA.35":       "NA-OH",  // Omaheke Region
	"NA.36":       "NA-OS",  // Omusati Region
	"NA.37":       "NA-ON",  // Oshana Region
	"NA.38":       "NA-OT",  // Oshikoto Region
	"NA.39":       "NA-OD",  // Otjozondjupa Region
	"NA.40":       "NA-KE",  // Kavango East
	"NA.41":       "NA-KW",  // Kavango West
	"NE.01":       "NE-1",   // Agadez
	"NE.02":       "NE-2",   // Diffa
	"NE.03":       "NE-3",   // Dosso Region
	"NE.04":       "NE-4",   // Maradi Region
	"NE.06":       "NE-5",   // Tahoua Region
	"NE.07":       "NE-7",   // Zinder Region
	"NE.08":       "NE-8",   // Niamey
	"NE.09":       "NE-6",   // Tillabéri Region
	"NG.05":       "NG-LA",  // Lagos
	"NG.11":       "NG-FC",  // FCT
	"NG.16":       "NG-OG",  // Ogun State
	"NG.21":       "NG-AK",  // Akwa Ibom State
	"NG.22":       "NG-CR",  // Cross River State
	"NG.23":       "NG-KD",  // Kaduna State
	"NG.24":       "NG-KT",  // Katsina State
	"NG.25":       "NG-AN",  // Anambra
	"NG.26":       "NG-BE",  // Benue State
	"NG.27":       "NG-BO",  // Borno State
	"NG.28":       "NG-IM",  // Imo State
	"NG.29":       "NG-KN",  // Kano State
	"NG.30":       "NG-KW",  // Kwara State
	"NG.31":       "NG-NI",  // Niger State
	"NG.32":       "NG-OY",  // Oyo State
	"NG.35":       "NG-AD",  // Adamawa
	"NG.36":       "NG-DE",  // Delta
	"NG.37":       "NG-ED",  // Edo State
	"NG.39":       "NG-JI",  // Jigawa State
	"NG.40":       "NG-KE",  // Kebbi
	"NG.41":       "NG-KO",  // Kogi State
	"NG.42":       "NG-OS",  // Osun State
	"NG.43":       "NG-TA",  // Taraba State
	"NG.44":       "NG-YO",  // Yobe State
	"NG.45":       "NG-AB",  // Abia State
	"NG.46":       "NG-BA",  // Bauchi
	"NG.47":       "NG-EN",  // Enugu State
	"NG.48":       "NG-ON",  // Ondo State
	"NG.49":       "NG-PL",  // Plateau State
	"NG.50":       "NG-RI",  // Rivers State
	"NG.51":       "NG-SO",  // Sokoto
	"NG.52":       "NG-BY",  // Bayelsa State
	"NG.53":       "NG-EB",  // Ebonyi State
	"NG.54":       "NG-EK",  // Ekiti State
	"NG.55":       "NG-GO",  // Gombe State
	"NG.56":       "NG-NA",  // Nasarawa State
	"NG.57":       "NG-ZA",  // Zamfara State
	"NI.01":       "NI-BO",  // Boaco Department
	"NI.02":       "NI-CA",  // Carazo Department
	"NI.03":       "NI-CI",  // Chinandega Department
	"NI.04":       "NI-CO",  // Chontales Department
	"NI.05":       "NI-ES",  // Estelí Department
	"NI.06":       "NI-GR",  // Granada Department
	"NI.07":       "NI-JI",  // Jinotega Department
	"NI.08":       "NI-LE",  // León Department
	"NI.09":       "NI-MD",  // Madriz Department
	"NI.10":       "NI-MN",  // Managua Department
	"NI.11":       "NI-MS",  // Masaya Department
	"NI.12":       "NI-MT",  // Matagalpa Department
	"NI.13":       "NI-NS",  // Nueva Segovia Department
	"NI.14":       "NI-SJ",  // Río San Juan Department
	"NI.15":       "NI-RI",  // Rivas Department
	"NI.17":       "NI-AN",  // North Caribbean Coast
	"NI.18":       "NI-AS",  // South Caribbean Coast
	"NL.01":       "NL-DR",  // Drenthe
	"NL.02":       "NL-FR",  // Friesland
	"NL.03":       "NL-GE",  // Gelderland
	"NL.04":       "NL-GR",  // Groningen
	"NL.05":       "NL-LI",  // Limburg
	"NL.06":       "NL-NB",  // North Brabant
	"NL.07":       "NL-NH",  // North Holland
	"NL.09":       "NL-UT",  // Utrecht
	"NL.10":       "NL-ZE",  // Zeeland
	"NL.11":       "NL-ZH",  // South Holland
	"NL.15":       "NL-OV",  // Overijssel
	"NL.16":       "NL-FL",  // Flevoland
	"NO.05":       "NO-54",  // Finnmark
	"NO.08":       "NO-15",  // Møre og Romsdal
	"NO.09":       "NO-18",  // Nordland
	"NO.12":       "NO-03",  // Oslo
	"NO.14":       "NO-11",  // Rogaland
	"NO.17":       "NO-38",  // Telemark
	"NO.20":       "NO-38",  // Vestfold
	"NO.21":       "NO-50",  // Trøndelag
	"NO.34":       "NO-34",  // Innlandet
	"NO.42":       "NO-42",  // Agder
	"NO.46":       "NO-46",  // Vestland
	"NP.1":        "NP-4",   // Koshi
	"NP.2":        "NP-P2",  // Madhesh
	"NP.3":        "NP-1",   // Bagmati Province
	"NP.4":        "NP-P4",  // Gandaki Pradesh
	"NP.5":        "NP-3",   // Lumbini Province
	"NP.6":        "NP-P6",  // Karnali Pradesh
	"NP.7":        "NP-5",   // Sudurpashchim Pradesh
	"NR.01":       "NR-01",  // Aiwo District
	"NR.02":       "NR-02",  // Anabar District
	"NR.03":       "NR-03",  // Anetan District
	"NR.04":       "NR-04",  // Anibare District
	"NR.05":       "NR-05",  // Baiti District
	"NR.06":       "NR-06",  // Boe District
	"NR.07":       "NR-07",  // Buada District
	"NR.08":       "NR-08",  // Denigomodu District
	"NR.09":       "NR-09",  // Ewa District
	"NR.10":       "NR-10",  // Ijuw District
	"NR.11":       "NR-11",  // Meneng District
	"NR.12":       "NR-12",  // Nibok District
	"NR.13":       "NR-13",  // Uaboe District
	"NR.14":       "NR-14",  // Yaren District
	"NZ.10":       "NZ-CIT", // Chatham Islands
	"NZ.E7":       "NZ-AUK", // Auckland
	"NZ.E8":       "NZ-BOP", // Bay of Plenty
	"NZ.E9":       "NZ-CAN", // Canterbury
	"NZ.F1":       "NZ-GIS", // Gisborne
	"NZ.F2":       "NZ-HKB", // Hawke's Bay Region
	"NZ.F3":       "NZ-MWT", // Manawatu-Wanganui
	"NZ.F4":       "NZ-MBH", // Marlborough
	"NZ.F5":       "NZ-NSN", // Nelson Region
	"NZ.F6":       "NZ-NTL", // Northland
	"NZ.F7":       "NZ-OTA", // Otago
	"NZ.F8":       "NZ-STL", // Southland
	"NZ.F9":       "NZ-TKI", // Taranaki Region
	"NZ.G1":       "NZ-WKO", // Waikato Region
	"NZ.G2":       "NZ-WGN", // Wellington Region
	"NZ.G3":       "NZ-WTC", // West Coast
	"NZ.TAS":      "NZ-TAS", // Tasman District
	"OM.01":       "OM-DA",  // Ad Dakhiliyah
	"OM.02":       "OM-BJ",  // Al Batinah South
	"OM.03":       "OM-WU",  // Al Wusta Governorate
	"OM.06":       "OM-MA",  // Muscat
	"OM.07":       "OM-MU",  // Musandam Governorate
	"OM.08":       "OM-ZU",  // Dhofar
	"OM.09":       "OM-ZA",  // Ad Dhahirah
	"OM.10":       "OM-BU",  // Al Buraimi
	"OM.11":       "OM-BS",  // Al Batinah North
	"PA.01":       "PA-1",   // Bocas del Toro Province
	"PA.02":       "PA-4",   // Chiriquí Province
	"PA.03":       "PA-2",   // Coclé
	"PA.04":       "PA-3",   // Colón
	"PA.05":       "PA-5",   // Darién
	"PA.06":       "PA-6",   // Herrera Province
	"PA.07":       "PA-7",   // Los Santos Province
	"PA.08":       "PA-8",   // Panamá
	"PA.09":       "PA-KY",  // Guna Yala
	"PA.10":       "PA-9",   // Veraguas Province
	"PA.11":       "PA-EM",  // Emberá
	"PA.12":       "PA-NB",  // Ngöbe-Buglé Comarca
	"PA.13":       "PA-10",  // Panamá Oeste Province
	"PA.NT":       "PA-NT",  // Naso Tjër Di
	"PE.01":       "PE-AMA", // Amazonas
	"PE.02":       "PE-ANC", // Ancash
	"PE.03":       "PE-APU", // Apurímac Department
	"PE.04":       "PE-ARE", // Arequipa
	"PE.05":       "PE-AYA", // Ayacucho
	"PE.06":       "PE-CAJ", // Cajamarca Department
	"PE.07":       "PE-CAL", // Callao
	"PE.08":       "PE-CUS", // Cuzco Department
	"PE.09":       "PE-HUV", // Huancavelica
	"PE.10":       "PE-HUC", // Huánuco Department
	"PE.11":       "PE-ICA", // Ica
	"PE.12":       "PE-JUN", // Junin
	"PE.13":       "PE-LAL", // La Libertad
	"PE.14":       "PE-LAM", // Lambayeque
	"PE.15":       "PE-LIM", // Lima region
	"PE.16":       "PE-LOR", // Loreto
	"PE.17":       "PE-MDD", // Madre de Dios
	"PE.18":       "PE-MOQ", // Moquegua Department
	"PE.19":       "PE-PAS", // Pasco
	"PE.20":       "PE-PIU", // Piura
	"PE.21":       "PE-PUN", // Puno
	"PE.22":       "PE-SAM", // San Martín Department
	"PE.23":       "PE-TAC", // Tacna
	"PE.24":       "PE-TUM", // Tumbes
	"PE.25":       "PE-UCA", // Ucayali
	"PE.LMA":      "PE-LMA", // Lima Province
	"PG.01":       "PG-CPM", // Central Province
	"PG.02":       "PG-GPK", // Gulf Province
	"PG.03":       "PG-MBA", // Milne Bay Province
	"PG.04":       "PG-NPP", // Oro Province
	"PG.05":       "PG-SHM", // Southern Highlands Province
	"PG.06":       "PG-WPD", // Western Province
	"PG.07":       "PG-NSB", // Bougainville
	"PG.08":       "PG-CPK", // Chimbu Province
	"PG.09":       "PG-EHG", // Eastern Highlands Province
	"PG.10":       "PG-EBR", // East New Britain Province
	"PG.11":       "PG-ESW", // East Sepik Province
	"PG.12":       "PG-MPM", // Madang Province
	"PG.13":       "PG-MRL", // Manus Province
	"PG.14":       "PG-MPL", // Morobe Province
	"PG.15":       "PG-NIK", // New Ireland
	"PG.16":       "PG-WHM", // Western Highlands Province
	"PG.17":       "PG-WBK", // West New Britain Province
	"PG.18":       "PG-SAN", // Sandaun Province
	"PG.19":       "PG-EPW", // Enga Province
	"PG.20":       "PG-NCD", // National Capital
	"PG.21":       "PG-HLA", // Hela Province
	"PG.22":       "PG-JWK", // Jiwaka Province
	"PH.01":       "PH-01",  // Ilocos
	"PH.02":       "PH-02",  // Cagayan Valley
	"PH.03":       "PH-03",  // Central Luzon
	"PH.05":       "PH-05",  // Bicol Region
	"PH.06":       "PH-06",  // Western Visayas
	"PH.07":       "PH-07",  // Central Visayas
	"PH.08":       "PH-08",  // Eastern Visayas
	"PH.09":       "PH-09",  // Zamboanga Peninsula
	"PH.10":       "PH-10",  // Northern Mindanao
	"PH.11":       "PH-11",  // Davao Region
	"PH.12":       "PH-12",  // Soccsksargen
	"PH.13":       "PH-13",  // Caraga
	"PH.14":       "PH-14",  // Autonomous Region in Muslim Mindanao
	"PH.15":       "PH-15",  // Cordillera
	"PH.40":       "PH-40",  // Calabarzon
	"PH.41":       "PH-41",  // Mimaropa
	"PH.NCR":      "PH-00",  // National Capital Region
	"PK.02":       "PK-BA",  // Balochistan
	"PK.03":       "PK-KP",  // Khyber Pakhtunkhwa
	"PK.04":       "PK-PB",  // Punjab
	"PK.05":       "PK-SD",  // Sindh
	"PK.06":       "PK-JK",  // Azad Kashmir
	"PK.07":       "PK-GB",  // Gilgit-Baltistan
	"PK.08":       "PK-IS",  // Islamabad
	"PL.72":       "PL-02",  // Lower Silesia
	"PL.73":       "PL-04",  // Kujawsko-Pomorskie
	"PL.74":       "PL-10",  // Łódź Voivodeship
	"PL.75":       "PL-06",  // Lublin
	"PL.76":       "PL-08",  // Lubusz
	"PL.77":       "PL-12",  // Lesser Poland
	"PL.78":       "PL-14",  // Mazovia
	"PL.79":       "PL-16",  // Opole Voivodeship
	"PL.80":       "PL-18",  // Subcarpathia
	"PL.81":       "PL-20",  // Podlasie
	"PL.82":       "PL-22",  // Pomerania
	"PL.83":       "PL-24",  // Silesia
	"PL.84":       "PL-26",  // Świętokrzyskie
	"PL.85":       "PL-28",  // Warmia-Masuria
	"PL.86":       "PL-30",  // Greater Poland
	"PL.87":       "PL-32",  // West Pomerania
	"PS.GZ":       "PS-GZA", // Gaza Strip
	"PT.02":       "PT-01",  // Aveiro
	"PT.03":       "PT-02",  // Beja
	"PT.04":       "PT-03",  // Braga
	"PT.05":       "PT-04",  // Bragança
	"PT.06":       "PT-05",  // Castelo Branco
	"PT.07":       "PT-06",  // Coimbra
	"PT.08":       "PT-07",  // Évora
	"PT.09":       "PT-08",  // Faro
	"PT.10":       "PT-30",  // Madeira
	"PT.11":       "PT-09",  // Guarda
	"PT.13":       "PT-10",  // Leiria
	"PT.14":       "PT-11",  // Lisbon
	"PT.16":       "PT-12",  // Portalegre
	"PT.17":       "PT-13",  // Porto
	"PT.18":       "PT-14",  // Santarém
	"PT.19":       "PT-15",  // Setúbal
	"PT.20":       "PT-16",  // Viana do Castelo
	"PT.21":       "PT-17",  // Vila Real
	"PT.22":       "PT-18",  // Viseu
	"PT.23":       "PT-20",  // Azores
	"PW.01":       "PW-002", // Aimeliik
	"PW.02":       "PW-004", // Airai
	"PW.03":       "PW-010", // Angaur
	"PW.04":       "PW-050", // Hatohobei
	"PW.05":       "PW-100", // Kayangel
	"PW.06":       "PW-150", // Koror
	"PW.07":       "PW-212", // Melekeok
	"PW.08":       "PW-214", // Ngaraard
	"PW.09":       "PW-218", // Ngarchelong
	"PW.10":       "PW-222", // Ngardmau
	"PW.11":       "PW-224", // Ngatpang
	"PW.12":       "PW-226", // Ngchesar
	"PW.13":       "PW-227", // Ngeremlengui
	"PW.14":       "PW-228", // Ngiwal
	"PW.15":       "PW-350", // Peleliu
	"PW.16":       "PW-370", // Sonsorol
	"PY.01":       "PY-10",  // Alto Paraná Department
	"PY.02":       "PY-13",  // Amambay Department
	"PY.04":       "PY-5",   // Caaguazú Department
	"PY.05":       "PY-6",   // Caazapá
	"PY.06":       "PY-11",  // Central Department
	"PY.07":       "PY-1",   // Concepción
	"PY.08":       "PY-3",   // Cordillera Department
	"PY.10":       "PY-4",   // Guairá
	"PY.11":       "PY-7",   // Itapúa
	"PY.12":       "PY-8",   // Misiones Department
	"PY.13":       "PY-12",  // Ñeembucú Department
	"PY.15":       "PY-9",   // Paraguarí
	"PY.16":       "PY-15",  // Presidente Hayes
	"PY.17":       "PY-2",   // San Pedro Department
	"PY.19":       "PY-14",  // Canindeyú
	"PY.22":       "PY-ASU", // Asunción
	"PY.23":       "PY-16",  // Alto Paraguay
	"PY.24":       "PY-19",  // Boquerón Department
	"QA.01":       "QA-DA",  // Baladīyat ad Dawḩah
	"QA.04":       "QA-KH",  // Al Khor
	"QA.06":       "QA-RA",  // Baladīyat ar Rayyān
	"QA.08":       "QA-MS",  // Madīnat ash Shamāl
	"QA.09":       "QA-US",  // Umm Salal
	"QA.10":       "QA-WA",  // Al Wakrah
	"QA.13":       "QA-ZA",  // Al Daayen
	"QA.14":       "QA-SH",  // Al-Shahaniya
	"RO.01":       "RO-AB",  // Alba County
	"RO.02":       "RO-AR",  // Arad County
	"RO.03":       "RO-AG",  // Arges
	"RO.04":       "RO-BC",  // Bacău County
	"RO.05":       "RO-BH",  // Bihor County
	"RO.06":       "RO-BN",  // Bistrița-Năsăud County
	"RO.07":       "RO-BT",  // Botoșani County
	"RO.08":       "RO-BR",  // Brăila County
	"RO.09":       "RO-BV",  // Brașov County
	"RO.10":       "RO-B",   // București
	"RO.11":       "RO-BZ",  // Buzău County
	"RO.12":       "RO-CS",  // Caraș-Severin County
	"RO.13":       "RO-CJ",  // Cluj County
	"RO.14":       "RO-CT",  // Constanța County
	"RO.15":       "RO-CV",  // Covasna County
	"RO.16":       "RO-DB",  // Dâmbovița County
	"RO.17":       "RO-DJ",  // Dolj
	"RO.18":       "RO-GL",  // Galați County
	"RO.19":       "RO-GJ",  // Gorj County
	"RO.20":       "RO-HR",  // Harghita County
	"RO.21":       "RO-HD",  // Hunedoara County
	"RO.22":       "RO-IL",  // Ialomița County
	"RO.23":       "RO-IS",  // Iași County
	"RO.25":       "RO-MM",  // Maramureş
	"RO.26":       "RO-MH",  // Mehedinți County
	"RO.27":       "RO-MS",  // Mureș County
	"RO.28":       "RO-NT",  // Neamț County
	"RO.29":       "RO-OT",  // Olt
	"RO.30":       "RO-PH",  // Prahova
	"RO.31":       "RO-SJ",  // Sălaj County
	"RO.32":       "RO-SM",  // Satu Mare County
	"RO.33":       "RO-SB",  // Sibiu County
	"RO.34":       "RO-SV",  // Suceava
	"RO.35":       "RO-TR",  // Teleorman County
	"RO.36":       "RO-TM",  // Timiș County
	"RO.37":       "RO-TL",  // Tulcea County
	"RO.38":       "RO-VS",  // Vaslui County
	"RO.39":       "RO-VL",  // Vâlcea County
	"RO.40":       "RO-VN",  // Vrancea
	"RO.41":       "RO-CL",  // Călărași County
	"RO.42":       "RO-GR",  // Giurgiu County
	"RO.43":       "RO-IF",  // Ilfov
	"RS.VO":       "RS-VO",  // Vojvodina
	"RU.01":       "RU-AD",  // Adygeya Republic
	"RU.03":       "RU-AL",  // Altai
	"RU.04":       "RU-ALT", // Altai Krai
	"RU.05":       "RU-AMU", // Amur Oblast
	"RU.06":       "RU-ARK", // Arkhangelskaya
	"RU.07":       "RU-AST", // Astrakhan Oblast
	"RU.08":       "RU-BA",  // Bashkortostan Republic
	"RU.09":       "RU-BEL", // Belgorod Oblast
	"RU.10":       "RU-BRY", // Bryansk Oblast
	"RU.11":       "RU-BU",  // Buryatiya Republic
	"RU.12":       "RU-CE",  // Chechnya
	"RU.13":       "RU-CHE", // Chelyabinsk
	"RU.15":       "RU-CHU", // Chukotka
	"RU.16":       "RU-CU",  // Chuvash Republic
	"RU.17":       "RU-DA",  // Dagestan
	"RU.19":       "RU-IN",  // Ingushetiya Republic
	"RU.20":       "RU-IRK", // Irkutsk Oblast
	"RU.21":       "RU-IVA", // Ivanovo Oblast
	"RU.22":       "RU-KB",  // Kabardino-Balkariya Republic
	"RU.23":       "RU-KGD", // Kaliningrad Oblast
	"RU.24":       "RU-KL",  // Kalmykiya Republic
	"RU.25":       "RU-KLU", // Kaluga Oblast
	"RU.27":       "RU-KC",  // Karachayevo-Cherkesiya Republic
	"RU.28":       "RU-KR",  // Karelia
	"RU.29":       "RU-KEM", // Kuzbass
	"RU.30":       "RU-KHA", // Khabarovsk
	"RU.31":       "RU-KK",  // Khakasiya Republic
	"RU.32":       "RU-KHM", // Khanty-Mansia
	"RU.33":       "RU-KIR", // Kirov Oblast
	"RU.34":       "RU-KO",  // Komi
	"RU.37":       "RU-KOS", // Kostroma Oblast
	"RU.38":       "RU-KDA", // Krasnodar Krai
	"RU.40":       "RU-KGN", // Kurgan Oblast
	"RU.41":       "RU-KRS", // Kursk Oblast
	"RU.42":       "RU-LEN", // Leningradskaya Oblast'
	"RU.43":       "RU-LIP", // Lipetsk Oblast
	"RU.44":       "RU-MAG", // Magadan Oblast
	"RU.45":       "RU-ME",  // Mariy-El Republic
	"RU.46":       "RU-MO",  // Mordoviya Republic
	"RU.47":       "RU-MOS", // Moscow Oblast
	"RU.48":       "RU-MOW", // Moscow
	"RU.49":       "RU-MUR", // Murmansk
	"RU.50":       "RU-NEN", // Nenets
	"RU.51":       "RU-NIZ", // Nizhny Novgorod Oblast
	"RU.52":       "RU-NGR", // Novgorod Oblast
	"RU.53":       "RU-NVS", // Novosibirsk Oblast
	"RU.54":       "RU-OMS", // Omsk Oblast
	"RU.55":       "RU-ORE", // Orenburg Oblast
	"RU.56":       "RU-ORL", // Oryol oblast
	"RU.57":       "RU-PNZ", // Penza Oblast
	"RU.59":       "RU-PRI", // Primorye
	"RU.60":       "RU-PSK", // Pskov Oblast
	"RU.61":       "RU-ROS", // Rostov
	"RU.62":       "RU-RYA", // Ryazan Oblast
	"RU.63":       "RU-SA",  // Sakha
	"RU.64":       "RU-SAK", // Sakhalin Oblast
	"RU.65":       "RU-SAM", // Samara Oblast
	"RU.66":       "RU-SPE", // St.-Petersburg
	"RU.67":       "RU-SAR", // Saratov Oblast
	"RU.68":       "RU-SE",  // North Ossetia–Alania
	"RU.69":       "RU-SMO", // Smolensk Oblast
	"RU.70":       "RU-STA", // Stavropol Kray
	"RU.71":       "RU-SVE", // Sverdlovsk Oblast
	"RU.72":       "RU-TAM", // Tambov Oblast
	"RU.73":       "RU-TA",  // Tatarstan Republic
	"RU.75":       "RU-TOM", // Tomsk Oblast
	"RU.76":       "RU-TUL", // Tula Oblast
	"RU.77":       "RU-TVE", // Tver Oblast
	"RU.78":       "RU-TYU", // Tyumen Oblast
	"RU.79":       "RU-TY",  // Republic of Tyva
	"RU.80":       "RU-UD",  // Udmurtiya Republic
	"RU.81":       "RU-ULY", // Ulyanovsk
	"RU.83":       "RU-VLA", // Vladimir Oblast
	"RU.84":       "RU-VGG", // Volgograd Oblast
	"RU.85":       "RU-VLG", // Vologda Oblast
	"RU.86":       "RU-VOR", // Voronezh Oblast
	"RU.87":       "RU-YAN", // Yamalo-Nenets
	"RU.88":       "RU-YAR", // Yaroslavl Oblast
	"RU.89":       "RU-YEV", // Jewish Autonomous Oblast
	"RU.90":       "RU-PER", // Perm Krai
	"RU.91":       "RU-KYA", // Krasnoyarsk Krai
	"RU.92":       "RU-KAM", // Kamchatka
	"RU.93":       "RU-ZAB", // Zabaykalskiy (Transbaikal) Kray
	"RW.11":       "RW-02",  // Eastern Province
	"RW.12":       "RW-01",  // Kigali
	"RW.13":       "RW-03",  // Northern Province
	"RW.14":       "RW-04",  // Western Province
	"RW.15":       "RW-05",  // Southern Province
	"SA.02":       "SA-11",  // Al Bahah Region
	"SA.05":       "SA-03",  // Medina Region
	"SA.06":       "SA-04",  // Eastern Province
	"SA.08":       "SA-05",  // Al-Qassim Region
	"SA.10":       "SA-01",  // Riyadh Region
	"SA.11":       "SA-14",  // 'Asir Region
	"SA.13":       "SA-06",  // Ha'il Region
	"SA.14":       "SA-02",  // Mecca Region
	"SA.15":       "SA-08",  // Northern Borders Region
	"SA.16":       "SA-10",  // Najran Region
	"SA.17":       "SA-09",  // Jazan Region
	"SA.19":       "SA-07",  // Tabuk Region
	"SA.20":       "SA-12",  // Al Jawf Region
	"SB.03":       "SB-ML",  // Malaita Province
	"SB.06":       "SB-GU",  // Guadalcanal Province
	"SB.07":       "SB-IS",  // Isabel Province
	"SB.08":       "SB-MK",  // Makira-Ulawa Province
	"SB.09":       "SB-TE",  // Temotu Province
	"SB.10":       "SB-CE",  // Central Province
	"SB.11":       "SB-WE",  // Western Province
	"SB.12":       "SB-CH",  // Choiseul Province
	"SB.13":       "SB-RB",  // Rennell and Bellona Province
	"SB.14":       "SB-CT",  // Honiara
	"SC.01":       "SC-01",  // Anse-aux-Pins
	"SC.02":       "SC-02",  // Anse Boileau
	"SC.03":       "SC-03",  // Anse Etoile
	"SC.05":       "SC-05",  // Anse Royale
	"SC.06":       "SC-06",  // Baie Lazare
	"SC.07":       "SC-07",  // Baie Sainte Anne
	"SC.08":       "SC-08",  // Beau Vallon
	"SC.09":       "SC-09",  // Bel Air
	"SC.10":       "SC-10",  // Bel Ombre
	"SC.11":       "SC-11",  // Cascade
	"SC.12":       "SC-12",  // Glacis
	"SC.12200079": "SC-26",  // Ile Perseverance I
	"SC.12200080": "SC-27",  // Ile Perseverance II
	"SC.14":       "SC-14",  // Grand Anse Praslin
	"SC.17":       "SC-17",  // Mont Buxton
	"SC.18":       "SC-18",  // Mont Fleuri
	"SC.19":       "SC-19",  // Plaisance
	"SC.20":       "SC-20",  // Pointe La Rue
	"SC.22":       "SC-22",  // Saint Louis
	"SC.23":       "SC-23",  // Takamaka
	"SC.24":       "SC-13",  // Grand Anse Mahe
	"SC.25":       "SC-15",  // La Digue and Inner Islands
	"SC.26":       "SC-16",  // La Rivière Anglaise
	"SC.27":       "SC-21",  // Port Glaud
	"SC.28":       "SC-04",  // Au Cap
	"SC.29":       "SC-24",  // Les Mamelles
	"SC.30":       "SC-25",  // Roche Caiman
	"SD.29":       "SD-KH",  // Khartoum
	"SD.36":       "SD-RS",  // Red Sea
	"SD.38":       "SD-GZ",  // Al Jazirah
	"SD.39":       "SD-GD",  // Al Qaḑārif
	"SD.41":       "SD-NW",  // White Nile
	"SD.42":       "SD-NB",  // Blue Nile
	"SD.43":       "SD-NO",  // Northern State
	"SD.47":       "SD-DW",  // Western Darfur
	"SD.49":       "SD-DS",  // Southern Darfur
	"SD.50":       "SD-KS",  // Southern Kordofan
	"SD.52":       "SD-KA",  // Kassala
	"SD.53":       "SD-NR",  // River Nile
	"SD.55":       "SD-DN",  // Northern Darfur
	"SD.56":       "SD-KN",  // North Kordofan
	"SD.58":       "SD-SI",  // Sennar
	"SD.60":       "SD-DE",  // Eastern Darfur
	"SD.61":       "SD-DC",  // Central Darfur
	"SD.62":       "SD-GK",  // West Kordofan
	"SE.02":       "SE-K",   // Blekinge
	"SE.03":       "SE-X",   // Gävleborg
	"SE.05":       "SE-I",   // Gotland
	"SE.06":       "SE-N",   // Halland
	"SE.07":       "SE-Z",   // Jämtland
	"SE.08":       "SE-F",   // Jönköping
	"SE.09":       "SE-H",   // Kalmar
	"SE.10":       "SE-W",   // Dalarna
	"SE.12":       "SE-G",   // Kronoberg
	"SE.14":       "SE-BD",  // Norrbotten
	"SE.15":       "SE-T",   // Örebro
	"SE.16":       "SE-E",   // Östergötland
	"SE.18":       "SE-D",   // Södermanland
	"SE.21":       "SE-C",   // Uppsala
	"SE.22":       "SE-S",   // Värmland
	"SE.23":       "SE-AC",  // Västerbotten
	"SE.24":       "SE-Y",   // Västernorrland
	"SE.25":       "SE-U",   // Västmanland
	"SE.26":       "SE-AB",  // Stockholm
	"SE.27":       "SE-M",   // Skåne
	"SE.28":       "SE-O",   // Västra Götaland
	"SH.01":       "SH-AC",  // Ascension
	"SH.02":       "SH-HL",  // Saint Helena
	"SI.01":       "SI-001", // Ajdovščina Municipality
	"SI.02":       "SI-002", // Beltinci
	"SI.03":       "SI-003", // Bled
	"SI.04":       "SI-004", // Bohinj Municipality
	"SI.05":       "SI-005", // Borovnica
	"SI.06":       "SI-006", // Bovec
	"SI.07":       "SI-007", // Brda
	"SI.08":       "SI-009", // Brežice
	"SI.09":       "SI-008", // Brezovica
	"SI.11":       "SI-011", // Celje
	"SI.12":       "SI-012", // Cerklje na Gorenjskem
	"SI.13":       "SI-013", // Cerknica
	"SI.14":       "SI-014", // Cerkno
	"SI.15":       "SI-015", // Črenšovci
	"SI.16":       "SI-016", // Črna na Koroškem Municipality
	"SI.17":       "SI-017", // Črnomelj Municipality
	"SI.19":       "SI-019", // Divača
	"SI.20":       "SI-020", // Dobrepolje
	"SI.22":       "SI-022", // Dol pri Ljubljani
	"SI.24":       "SI-024", // Dornava
	"SI.25":       "SI-025", // Dravograd
	"SI.26":       "SI-026", // Duplek
	"SI.27":       "SI-027", // Občina Gorenja vas-Poljane
	"SI.28":       "SI-028", // Gorišnica
	"SI.29":       "SI-029", // Gornja Radgona
	"SI.30":       "SI-030", // Gornji Grad
	"SI.31":       "SI-031", // Gornji Petrovci
	"SI.32":       "SI-032", // Grosuplje
	"SI.34":       "SI-034", // Hrastnik Municipality
	"SI.35":       "SI-035", // Hrpelje-Kozina
	"SI.36":       "SI-036", // Idrija
	"SI.37":       "SI-037", // Ig Municipality
	"SI.38":       "SI-038", // Ilirska Bistrica
	"SI.39":       "SI-039", // Ivančna Gorica
	"SI.40":       "SI-040", // Izola-Isola
	"SI.42":       "SI-042", // Juršinci
	"SI.44":       "SI-044", // Kanal ob Soči Municipality
	"SI.45":       "SI-045", // Kidričevo
	"SI.46":       "SI-046", // Kobarid Municipality
	"SI.47":       "SI-047", // Kobilje
	"SI.49":       "SI-049", // Komen
	"SI.50":       "SI-050", // Koper-Capodistria
	"SI.51":       "SI-051", // Kozje
	"SI.52":       "SI-052", // Kranj
	"SI.53":       "SI-053", // Kranjska Gora
	"SI.54":       "SI-054", // Krško
	"SI.55":       "SI-055", // Kungota
	"SI.57":       "SI-057", // Laško Municipality
	"SI.61":       "SI-061", // Ljubljana
	"SI.62":       "SI-062", // Ljubno
	"SI.64":       "SI-064", // Logatec
	"SI.66":       "SI-066", // Loški Potok
	"SI.68":       "SI-068", // Lukovica
	"SI.71":       "SI-071", // Medvode
	"SI.72":       "SI-072", // Mengeš
	"SI.73":       "SI-073", // Metlika
	"SI.74":       "SI-074", // Mežica
	"SI.76":       "SI-076", // Mislinja
	"SI.77":       "SI-077", // Moravče
	"SI.78":       "SI-078", // Moravske Toplice
	"SI.79":       "SI-079", // Mozirje
	"SI.80":       "SI-080", // Murska Sobota
	"SI.81":       "SI-081", // Muta
	"SI.82":       "SI-082", // Naklo
	"SI.83":       "SI-083", // Nazarje
	"SI.84":       "SI-084", // Nova Gorica
	"SI.86":       "SI-086", // Odranci
	"SI.87":       "SI-087", // Ormož
	"SI.88":       "SI-088", // Osilnica
	"SI.89":       "SI-089", // Pesnica
	"SI.91":       "SI-091", // Pivka
	"SI.92":       "SI-092", // Podčetrtek
	"SI.94":       "SI-094", // Postojna
	"SI.97":       "SI-097", // Puconci
	"SI.98":       "SI-098", // Rače-Fram
	"SI.99":       "SI-099", // Radeče Municipality
	"SI.A1":       "SI-100", // Radenci
	"SI.A2":       "SI-101", // Radlje ob Dravi
	"SI.A3":       "SI-102", // Radovljica
	"SI.A6":       "SI-105", // Rogašovci
	"SI.A7":       "SI-106", // Rogaška Slatina
	"SI.A8":       "SI-107", // Rogatec
	"SI.B1":       "SI-109", // Semič
	"SI.B2":       "SI-117", // Šenčur
	"SI.B3":       "SI-118", // Šentilj
	"SI.B4":       "SI-119", // Šentjernej
	"SI.B6":       "SI-110", // Sevnica Municipality
	"SI.B7":       "SI-111", // Sežana
	"SI.B8":       "SI-121", // Škocjan
	"SI.B9":       "SI-122", // Škofja Loka
	"SI.C1":       "SI-123", // Škofljica
	"SI.C2":       "SI-112", // Slovenj Gradec
	"SI.C4":       "SI-114", // Slovenska Konjice
	"SI.C5":       "SI-124", // Šmarje pri Jelšah
	"SI.C6":       "SI-125", // Šmartno ob Paki
	"SI.C7":       "SI-126", // Šoštanj
	"SI.C8":       "SI-115", // Starše
	"SI.C9":       "SI-127", // Štore
	"SI.D1":       "SI-116", // Sveti Jurij
	"SI.D2":       "SI-128", // Tolmin
	"SI.D3":       "SI-129", // Trbovlje
	"SI.D4":       "SI-130", // Trebnje
	"SI.D5":       "SI-131", // Tržič
	"SI.D6":       "SI-132", // Turnišče
	"SI.D7":       "SI-133", // Velenje
	"SI.D8":       "SI-134", // Velike Lašče
	"SI.E1":       "SI-136", // Vipava
	"SI.E2":       "SI-137", // Vitanje
	"SI.E3":       "SI-138", // Vodice
	"SI.E5":       "SI-140", // Vrhnika
	"SI.E6":       "SI-141", // Vuzenica
	"SI.E7":       "SI-142", // Zagorje ob Savi Municipality
	"SI.E9":       "SI-143", // Zavrč
	"SI.F1":       "SI-146", // Železniki Municipality
	"SI.F2":       "SI-147", // Žiri
	"SI.F3":       "SI-144", // Zreče Municipality
	"SI.F4":       "SI-148", // Benedikt
	"SI.F5":       "SI-149", // Bistrica ob Sotli
	"SI.F6":       "SI-150", // Bloke
	"SI.F7":       "SI-151", // Braslovče
	"SI.F8":       "SI-152", // Cankova
	"SI.F9":       "SI-153", // Cerkvenjak
	"SI.G1":       "SI-018", // Destrnik
	"SI.G2":       "SI-154", // Dobje
	"SI.G3":       "SI-155", // Dobrna
	"SI.G4":       "SI-021", // Dobrova-Horjul-Polhov Gradec
	"SI.G5":       "SI-156", // Dobrovnik-Dobronak
	"SI.G6":       "SI-157", // Dolenjske Toplice
	"SI.G7":       "SI-023", // Domžale
	"SI.G8":       "SI-158", // Grad
	"SI.G9":       "SI-159", // Hajdina
	"SI.H1":       "SI-160", // Hoče-Slivnica
	"SI.H2":       "SI-161", // Hodoš-Hodos
	"SI.H3":       "SI-162", // Horjul
	"SI.H4":       "SI-041", // Jesenice
	"SI.H5":       "SI-163", // Jezersko
	"SI.H6":       "SI-043", // Kamnik
	"SI.H7":       "SI-048", // Kočevje
	"SI.H8":       "SI-164", // Komenda
	"SI.H9":       "SI-165", // Kostel
	"SI.I1":       "SI-166", // Križevci
	"SI.I2":       "SI-056", // Kuzma
	"SI.I3":       "SI-058", // Lenart
	"SI.I4":       "SI-059", // Lendava-Lendva
	"SI.I5":       "SI-060", // Litija
	"SI.I6":       "SI-063", // Ljutomer
	"SI.I7":       "SI-065", // Loška Dolina Municipality
	"SI.I8":       "SI-167", // Lovrenc na Pohorju
	"SI.I9":       "SI-067", // Luče
	"SI.J1":       "SI-069", // Majšperk
	"SI.J2":       "SI-070", // Maribor City Municipality
	"SI.J3":       "SI-168", // Markovci
	"SI.J4":       "SI-169", // Miklavž na Dravskem Polju
	"SI.J5":       "SI-075", // Miren-Kostanjevica
	"SI.J6":       "SI-170", // Mirna Peč Municipality
	"SI.J7":       "SI-085", // Novo Mesto
	"SI.J8":       "SI-171", // Oplotnica Municipality
	"SI.J9":       "SI-090", // Piran-Pirano
	"SI.K1":       "SI-172", // Podlehnik
	"SI.K2":       "SI-093", // Podvelka
	"SI.K3":       "SI-173", // Polzela
	"SI.K4":       "SI-174", // Prebold
	"SI.K5":       "SI-095", // Preddvor
	"SI.K6":       "SI-175", // Prevalje
	"SI.K7":       "SI-096", // Ptuj
	"SI.K8":       "SI-103", // Ravne na Koroškem
	"SI.K9":       "SI-176", // Razkrižje
	"SI.L1":       "SI-104", // Ribnica
	"SI.L2":       "SI-177", // Ribnica na Pohorju
	"SI.L3":       "SI-108", // Ruše Municipality
	"SI.L4":       "SI-033", // Šalovci
	"SI.L5":       "SI-178", // Selnica ob Dravi
	"SI.L6":       "SI-183", // Šempeter-Vrtojba
	"SI.L7":       "SI-120", // Sentjur
	"SI.L8":       "SI-113", // Slovenska Bistrica
	"SI.L9":       "SI-194", // Šmartno pri Litiji
	"SI.M1":       "SI-179", // Sodražica
	"SI.M2":       "SI-180", // Solčava Municipality
	"SI.M3":       "SI-181", // Sveta Ana Municipality
	"SI.M4":       "SI-182", // Sveti Andraž v Slovenskih Goricah
	"SI.M5":       "SI-184", // Tabor
	"SI.M6":       "SI-010", // Tišina
	"SI.M7":       "SI-185", // Trnovska Vas
	"SI.M8":       "SI-186", // Trzin
	"SI.M9":       "SI-187", // Velika Polana
	"SI.N1":       "SI-188", // Veržej
	"SI.N2":       "SI-135", // Videm
	"SI.N3":       "SI-139", // Vojnik
	"SI.N4":       "SI-189", // Vransko
	"SI.N5":       "SI-190", // Žalec
	"SI.N6":       "SI-191", // Žetale
	"SI.N7":       "SI-192", // Žirovnica
	"SI.N8":       "SI-193", // Žužemberk
	"SI.N9":       "SI-195", // Apače
	"SI.O1":       "SI-196", // Cirkulane
	"SI.O2":       "SI-207", // Gorje
	"SI.O3":       "SI-197", // Kostanjevica na Krki
	"SI.O4":       "SI-208", // Log–Dragomer
	"SI.O5":       "SI-198", // Makole
	"SI.O6":       "SI-212", // Mirna
	"SI.O7":       "SI-199", // Mokronog-Trebelno Municipality
	"SI.O8":       "SI-200", // Poljčane
	"SI.O9":       "SI-209", // Rečica ob Savinji
	"SI.P1":       "SI-201", // Renče-Vogrsko
	"SI.P2":       "SI-211", // Šentrupert
	"SI.P3":       "SI-206", // Šmarješke Toplice
	"SI.P4":       "SI-202", // Središče ob Dravi
	"SI.P5":       "SI-203", // Straža
	"SI.P6":       "SI-204", // Sv. Trojica v Slov. Goricah
	"SI.P7":       "SI-210", // Sveti Jurij v Slovenskih Goricah
	"SI.P8":       "SI-205", // Sveti Tomaž
	"SI.P9":       "SI-213", // Ankaran Municipality
	"SK.01":       "SK-BC",  // Banská Bystrica Region
	"SK.02":       "SK-BL",  // Bratislava Region
	"SK.03":       "SK-KI",  // Košice Region
	"SK.04":       "SK-NI",  // Nitra Region
	"SK.05":       "SK-PV",  // Prešov Region
	"SK.06":       "SK-TC",  // Trenčín Region
	"SK.07":       "SK-TA",  // Trnava Region
	"SK.08":       "SK-ZI",  // Žilina Region
	"SL.01":       "SL-E",   // Eastern Province
	"SL.02":       "SL-N",   // Northern Province
	"SL.03":       "SL-S",   // Southern Province
	"SL.04":       "SL-W",   // Western Area
	"SL.05":       "SL-NW",  // North West
	"SM.01":       "SM-01",  // Acquaviva
	"SM.02":       "SM-02",  // Chiesanuova
	"SM.03":       "SM-03",  // Domagnano
	"SM.04":       "SM-04",  // Faetano
	"SM.05":       "SM-05",  // Fiorentino
	"SM.06":       "SM-06",  // Borgo Maggiore
	"SM.07":       "SM-07",  // San Marino
	"SM.08":       "SM-08",  // Montegiardino
	"SM.09":       "SM-09",  // Serravalle
	"SN.01":       "SN-DK",  // Dakar
	"SN.03":       "SN-DB",  // Diourbel Region
	"SN.05":       "SN-TC",  // Tambacounda
	"SN.07":       "SN-TH",  // Thiès
	"SN.09":       "SN-FK",  // Fatick
	"SN.10":       "SN-KL",  // Kaolack
	"SN.11":       "SN-KD",  // Kolda
	"SN.12":       "SN-ZG",  // Ziguinchor
	"SN.13":       "SN-LG",  // Louga
	"SN.14":       "SN-SL",  // Saint-Louis
	"SN.15":       "SN-MT",  // Matam
	"SN.16":       "SN-KA",  // Kaffrine
	"SN.17":       "SN-KE",  // Kédougou
	"SN.18":       "SN-SE",  // Sédhiou
	"SO.01":       "SO-BK",  // Bakool
	"SO.02":       "SO-BN",  // Banaadir
	"SO.03":       "SO-BR",  // Bari
	"SO.04":       "SO-BY",  // Bay
	"SO.05":       "SO-GA",  // Galguduud
	"SO.06":       "SO-GE",  // Gedo
	"SO.07":       "SO-HI",  // Hiiraan
	"SO.08":       "SO-JD",  // Middle Juba
	"SO.09":       "SO-JH",  // Lower Juba
	"SO.10":       "SO-MU",  // Mudug
	"SO.12":       "SO-SA",  // Sanaag
	"SO.14":       "SO-SH",  // Lower Shabeelle
	"SO.18":       "SO-NU",  // Nugaal
	"SO.19":       "SO-TO",  // Togdheer
	"SO.20":       "SO-WO",  // Woqooyi Galbeed
	"SO.21":       "SO-AW",  // Awdal
	"SO.22":       "SO-SO",  // Sool
	"SR.10":       "SR-BR",  // Brokopondo District
	"SR.11":       "SR-CM",  // Commewijne District
	"SR.12":       "SR-CR",  // Coronie District
	"SR.13":       "SR-MA",  // Marowijne District
	"SR.14":       "SR-NI",  // Nickerie District
	"SR.15":       "SR-PR",  // Para District
	"SR.16":       "SR-PM",  // Paramaribo District
	"SR.17":       "SR-SA",  // Saramacca District
	"SR.18":       "SR-SI",  // Sipaliwini District
	"SR.19":       "SR-WA",  // Wanica District
	"SS.01":       "SS-EC",  // Central Equatoria
	"SS.02":       "SS-EE",  // Eastern Equatoria
	"SS.03":       "SS-JG",  // Jonglei
	"SS.04":       "SS-LK",  // Lakes
	"SS.05":       "SS-BN",  // Northern Bahr al Ghazal
	"SS.06":       "SS-UY",  // Unity
	"SS.07":       "SS-NU",  // Upper Nile
	"SS.08":       "SS-WR",  // Warrap
	"SS.09":       "SS-BW",  // Western Bahr al Ghazal
	"SS.10":       "SS-EW",  // Western Equatoria
	"ST.01":       "ST-P",   // Príncipe
	"SV.01":       "SV-AH",  // Ahuachapán
	"SV.02":       "SV-CA",  // Cabañas
	"SV.03":       "SV-CH",  // Chalatenango Department
	"SV.04":       "SV-CU",  // Cuscatlán
	"SV.05":       "SV-LI",  // La Libertad Department
	"SV.06":       "SV-PA",  // La Paz Department
	"SV.07":       "SV-UN",  // La Unión
	"SV.08":       "SV-MO",  // Morazán
	"SV.09":       "SV-SM",  // San Miguel Department
	"SV.10":       "SV-SS",  // San Salvador Department
	"SV.11":       "SV-SA",  // Santa Ana Department
	"SV.12":       "SV-SV",  // San Vicente Department
	"SV.13":       "SV-SO",  // Sonsonate Department
	"SV.14":       "SV-US",  // Usulután
	"SY.01":       "SY-HA",  // Al-Hasakah
	"SY.02":       "SY-LA",  // Latakia
	"SY.03":       "SY-QU",  // Quneitra
	"SY.04":       "SY-RA",  // Raqqa
	"SY.05":       "SY-SU",  // Suwayda
	"SY.06":       "SY-DR",  // Daraa
	"SY.07":       "SY-DY",  // Deir ez-Zor
	"SY.08":       "SY-RD",  // Rif-dimashq
	"SY.09":       "SY-HL",  // Aleppo
	"SY.10":       "SY-HM",  // Hama
	"SY.11":       "SY-HI",  // Homs
	"SY.12":       "SY-ID",  // Idlib
	"SY.13":       "SY-DI",  // Damascus Governorate
	"SY.14":       "SY-TA",  // Tartus
	"SZ.01":       "SZ-HH",  // Hhohho Region
	"SZ.02":       "SZ-LU",  // Lubombo Region
	"SZ.03":       "SZ-MA",  // Manzini Region
	"SZ.04":       "SZ-SH",  // Shiselweni
	"TD.01":       "TD-BA",  // Batha
	"TD.02":       "TD-WF",  // Wadi Fira
	"TD.05":       "TD-GR",  // Guéra
	"TD.06":       "TD-KA",  // Kanem
	"TD.07":       "TD-LC",  // Lac
	"TD.08":       "TD-LO",  // Logone Occidental
	"TD.09":       "TD-LR",  // Logone Oriental
	"TD.12":       "TD-OD",  // Ouadaï
	"TD.13":       "TD-SA",  // Salamat
	"TD.14":       "TD-TA",  // Tandjilé
	"TD.15":       "TD-CB",  // Chari-Baguirmi
	"TD.16":       "TD-ME",  // Mayo-Kebbi Est
	"TD.17":       "TD-MC",  // Moyen-Chari
	"TD.18":       "TD-HL",  // Hadjer-Lamis
	"TD.19":       "TD-MA",  // Mandoul
	"TD.20":       "TD-MO",  // Mayo-Kebbi Ouest
	"TD.21":       "TD-ND",  // N’Djaména
	"TD.22":       "TD-BG",  // Barh el Gazel
	"TD.23":       "TD-BO",  // Borkou
	"TD.25":       "TD-SI",  // Sila
	"TD.26":       "TD-TI",  // Tibesti
	"TD.27":       "TD-EE",  // Ennedi-Est
	"TD.28":       "TD-EO",  // Ennedi-Ouest
	"TG.22":       "TG-C",   // Centrale
	"TG.23":       "TG-K",   // Kara
	"TG.24":       "TG-M",   // Maritime
	"TG.25":       "TG-P",   // Plateaux
	"TG.26":       "TG-S",   // Savanes
	"TH.01":       "TH-58",  // Mae Hong Son
	"TH.02":       "TH-50",  // Chiang Mai
	"TH.03":       "TH-57",  // Chiang Rai
	"TH.04":       "TH-55",  // Nan
	"TH.05":       "TH-51",  // Lamphun
	"TH.06":       "TH-52",  // Lampang
	"TH.07":       "TH-54",  // Phrae
	"TH.08":       "TH-63",  // Tak
	"TH.09":       "TH-64",  // Sukhothai
	"TH.10":       "TH-53",  // Uttaradit
	"TH.11":       "TH-62",  // Kamphaeng Phet
	"TH.12":       "TH-65",  // Phitsanulok
	"TH.13":       "TH-66",  // Phichit
	"TH.14":       "TH-67",  // Phetchabun
	"TH.15":       "TH-61",  // Uthai Thani
	"TH.16":       "TH-60",  // Nakhon Sawan
	"TH.17":       "TH-43",  // Nong Khai
	"TH.18":       "TH-42",  // Loei
	"TH.20":       "TH-47",  // Sakon Nakhon
	"TH.22":       "TH-40",  // Khon Kaen
	"TH.23":       "TH-46",  // Kalasin
	"TH.24":       "TH-44",  // Maha Sarakham
	"TH.25":       "TH-45",  // Roi Et
	"TH.26":       "TH-36",  // Chaiyaphum
	"TH.27":       "TH-30",  // Nakhon Ratchasima
	"TH.28":       "TH-31",  // Buriram
	"TH.29":       "TH-32",  // Surin
	"TH.30":       "TH-33",  // Si Sa Ket
	"TH.31":       "TH-96",  // Narathiwat
	"TH.32":       "TH-18",  // Chai Nat
	"TH.33":       "TH-17",  // Sing Buri
	"TH.34":       "TH-16",  // Lopburi
	"TH.35":       "TH-15",  // Ang Thong
	"TH.36":       "TH-14",  // Phra Nakhon Si Ayutthaya
	"TH.37":       "TH-19",  // Saraburi
	"TH.38":       "TH-12",  // Nonthaburi
	"TH.39":       "TH-13",  // Pathum Thani
	"TH.40":       "TH-10",  // Bangkok
	"TH.41":       "TH-56",  // Phayao
	"TH.42":       "TH-11",  // Samut Prakan
	"TH.43":       "TH-26",  // Nakhon Nayok
	"TH.44":       "TH-24",  // Chachoengsao
	"TH.46":       "TH-20",  // Chon Buri
	"TH.47":       "TH-21",  // Rayong
	"TH.48":       "TH-22",  // Chanthaburi
	"TH.49":       "TH-23",  // Trat
	"TH.50":       "TH-71",  // Kanchanaburi
	"TH.51":       "TH-72",  // Suphan Buri
	"TH.52":       "TH-70",  // Ratchaburi
	"TH.53":       "TH-73",  // Nakhon Pathom
	"TH.54":       "TH-75",  // Samut Songkhram
	"TH.55":       "TH-74",  // Samut Sakhon
	"TH.56":       "TH-76",  // Phetchaburi
	"TH.57":       "TH-77",  // Prachuap Khiri Khan
	"TH.58":       "TH-86",  // Chumphon
	"TH.59":       "TH-85",  // Ranong province
	"TH.60":       "TH-84",  // Surat Thani
	"TH.61":       "TH-82",  // Phang Nga
	"TH.62":       "TH-83",  // Phuket
	"TH.63":       "TH-81",  // Krabi
	"TH.64":       "TH-80",  // Nakhon Si Thammarat
	"TH.65":       "TH-92",  // Trang
	"TH.66":       "TH-93",  // Phatthalung
	"TH.67":       "TH-91",  // Satun
	"TH.68":       "TH-90",  // Songkhla
	"TH.69":       "TH-94",  // Pattani
	"TH.70":       "TH-95",  // Yala
	"TH.72":       "TH-35",  // Yasothon
	"TH.73":       "TH-48",  // Nakhon Phanom
	"TH.74":       "TH-25",  // Prachin Buri
	"TH.75":       "TH-34",  // Ubon Ratchathani
	"TH.76":       "TH-41",  // Udon Thani
	"TH.77":       "TH-37",  // Amnat Charoen
	"TH.78":       "TH-49",  // Mukdahan
	"TH.79":       "TH-39",  // Nong Bua Lamphu
	"TH.80":       "TH-27",  // Sa Kaeo
	"TH.81":       "TH-38",  // Bueng Kan
	"TJ.01":       "TJ-GB",  // Gorno-Badakhshan
	"TJ.02":       "TJ-KT",  // Khatlon Province
	"TJ.03":       "TJ-SU",  // Sughd
	"TJ.04":       "TJ-DU",  // Dushanbe
	"TL.AL":       "TL-AL",  // Aileu
	"TL.AN":       "TL-AN",  // Ainaro
	"TL.BA":       "TL-BA",  // Baucau
	"TL.BO":       "TL-BO",  // Bobonaro
	"TL.CO":       "TL-CO",  // Cova Lima
	"TL.DI":       "TL-DI",  // Dili Municipality
	"TL.ER":       "TL-ER",  // Ermera
	"TL.LA":       "TL-LA",  // Lautém
	"TL.LI":       "TL-LI",  // Liquiçá
	"TL.MF":       "TL-MF",  // Manufahi
	"TL.MT":       "TL-MT",  // Manatuto
	"TL.OE":       "TL-OE",  // Oecusse
	"TL.VI":       "TL-VI",  // Viqueque
	"TM.01":       "TM-A",   // Ahal
	"TM.02":       "TM-B",   // Balkan
	"TM.03":       "TM-D",   // Daşoguz
	"TM.04":       "TM-L",   // Lebap
	"TM.05":       "TM-M",   // Mary
	"TM.S":        "TM-S",   // Ashgabat
	"TN.02":       "TN-42",  // Kasserine Governorate
	"TN.03":       "TN-41",  // Kairouan
	"TN.06":       "TN-32",  // Jendouba Governorate
	"TN.14":       "TN-33",  // Kef Governorate
	"TN.15":       "TN-53",  // Mahdia Governorate
	"TN.16":       "TN-52",  // Monastir Governorate
	"TN.17":       "TN-31",  // Béja Governorate
	"TN.18":       "TN-23",  // Bizerte Governorate
	"TN.19":       "TN-21",  // Nabeul Governorate
	"TN.22":       "TN-34",  // Siliana Governorate
	"TN.23":       "TN-51",  // Sousse Governorate
	"TN.27":       "TN-13",  // Ben Arous Governorate
	"TN.28":       "TN-82",  // Medenine Governorate
	"TN.29":       "TN-81",  // Gabès Governorate
	"TN.30":       "TN-71",  // Gafsa
	"TN.31":       "TN-73",  // Kebili Governorate
	"TN.32":       "TN-61",  // Sfax Governorate
	"TN.33":       "TN-43",  // Sidi Bouzid Governorate
	"TN.34":       "TN-83",  // Tataouine
	"TN.35":       "TN-72",  // Tozeur Governorate
	"TN.36":       "TN-11",  // Tunis Governorate
	"TN.37":       "TN-22",  // Zaghouan Governorate
	"TN.38":       "TN-12",  // Ariana Governorate
	"TN.39":       "TN-14",  // Manouba
	"TO.01":       "TO-02",  // Ha‘apai
	"TO.02":       "TO-04",  // Tongatapu
	"TO.03":       "TO-05",  // Vava‘u
	"TO.EU":       "TO-01",  // ʻEua
	"TO.NI":       "TO-03",  // Niuas
	"TR.02":       "TR-02",  // Adıyaman Province
	"TR.03":       "TR-03",  // Afyonkarahisar Province
	"TR.04":       "TR-04",  // Ağrı
	"TR.05":       "TR-05",  // Amasya
	"TR.07":       "TR-07",  // Antalya
	"TR.08":       "TR-08",  // Artvin
	"TR.09":       "TR-09",  // Aydın
	"TR.10":       "TR-10",  // Balıkesir
	"TR.11":       "TR-11",  // Bilecik
	"TR.12":       "TR-12",  // Bingöl
	"TR.13":       "TR-13",  // Bitlis
	"TR.14":       "TR-14",  // Bolu
	"TR.15":       "TR-15",  // Burdur
	"TR.16":       "TR-16",  // Bursa Province
	"TR.17":       "TR-17",  // Canakkale
	"TR.19":       "TR-19",  // Çorum
	"TR.20":       "TR-20",  // Denizli
	"TR.21":       "TR-21",  // Diyarbakır Province
	"TR.22":       "TR-22",  // Edirne
	"TR.23":       "TR-23",  // Elazığ
	"TR.24":       "TR-24",  // Erzincan
	"TR.25":       "TR-25",  // Erzurum
	"TR.26":       "TR-26",  // Eskişehir
	"TR.28":       "TR-28",  // Giresun
	"TR.31":       "TR-31",  // Hatay
	"TR.32":       "TR-33",  // Mersin
	"TR.33":       "TR-32",  // Isparta
	"TR.34":       "TR-34",  // Istanbul
	"TR.35":       "TR-35",  // İzmir Province
	"TR.37":       "TR-37",  // Kastamonu
	"TR.38":       "TR-38",  // Kayseri
	"TR.39":       "TR-39",  // Kırklareli
	"TR.40":       "TR-40",  // Kırşehir
	"TR.41":       "TR-41",  // Kocaeli
	"TR.43":       "TR-43",  // Kütahya
	"TR.44":       "TR-44",  // Malatya
	"TR.45":       "TR-45",  // Manisa
	"TR.46":       "TR-46",  // Kahramanmaraş
	"TR.48":       "TR-48",  // Muğla
	"TR.49":       "TR-49",  // Muş
	"TR.50":       "TR-50",  // Nevşehir Province
	"TR.52":       "TR-52",  // Ordu
	"TR.53":       "TR-53",  // Rize Province
	"TR.54":       "TR-54",  // Sakarya
	"TR.55":       "TR-55",  // Samsun
	"TR.57":       "TR-57",  // Sinop
	"TR.58":       "TR-58",  // Sivas
	"TR.59":       "TR-59",  // Tekirdağ
	"TR.60":       "TR-60",  // Tokat Province
	"TR.61":       "TR-61",  // Trabzon
	"TR.62":       "TR-62",  // Tunceli
	"TR.63":       "TR-63",  // Şanlıurfa
	"TR.64":       "TR-64",  // Uşak
	"TR.65":       "TR-65",  // Van
	"TR.66":       "TR-66",  // Yozgat
	"TR.68":       "TR-06",  // Ankara
	"TR.69":       "TR-29",  // Gümüşhane Province
	"TR.70":       "TR-30",  // Hakkâri
	"TR.71":       "TR-42",  // Konya
	"TR.72":       "TR-47",  // Mardin
	"TR.73":       "TR-51",  // Niğde Province
	"TR.74":       "TR-56",  // Siirt
	"TR.75":       "TR-68",  // Aksaray
	"TR.76":       "TR-72",  // Batman
	"TR.77":       "TR-69",  // Bayburt Province
	"TR.78":       "TR-70",  // Karaman
	"TR.79":       "TR-71",  // Kırıkkale
	"TR.80":       "TR-73",  // Şırnak
	"TR.81":       "TR-01",  // Adana
	"TR.82":       "TR-18",  // Çankırı
	"TR.83":       "TR-27",  // Gaziantep
	"TR.84":       "TR-36",  // Kars Province
	"TR.85":       "TR-67",  // Zonguldak Province
	"TR.86":       "TR-75",  // Ardahan
	"TR.87":       "TR-74",  // Bartın
	"TR.88":       "TR-76",  // Iğdır
	"TR.89":       "TR-78",  // Karabük Province
	"TR.90":       "TR-79",  // Kilis
	"TR.91":       "TR-80",  // Osmaniye
	"TR.92":       "TR-77",  // Yalova
	"TR.93":       "TR-81",  // Düzce
	"TT.01":       "TT-ARI", // Borough of Arima
	"TT.03":       "TT-MRC", // Mayaro
	"TT.05":       "TT-POS", // Port of Spain
	"TT.10":       "TT-SFO", // San Fernando
	"TT.11":       "TT-TOB", // Tobago
	"TT.CHA":      "TT-CHA", // Chaguanas
	"TT.CTT":      "TT-CTT", // Couva-Tabaquite-Talparo
	"TT.DMN":      "TT-DMN", // Diego Martin Regional Corporation
	"TT.PED":      "TT-PED", // Penal/Debe
	"TT.PRT":      "TT-PRT", // Princes Town
	"TT.PTF":      "TT-PTF", // Point Fortin
	"TT.SGE":      "TT-SGE", // Sangre Grande Regional Corporation
	"TT.SIP":      "TT-SIP", // Siparia Regional Corporation
	"TT.SJL":      "TT-SJL", // San Juan/Laventille
	"TT.TUP":      "TT-TUP", // Tunapuna/Piarco
	"TV.FUN":      "TV-FUN", // Funafuti
	"TV.NIT":      "TV-NIT", // Niutao
	"TV.NKF":      "TV-NKF", // Nukufetau
	"TV.NKL":      "TV-NKL", // Nukulaelae
	"TV.NMA":      "TV-NMA", // Nanumea
	"TV.NMG":      "TV-NMG", // Nanumanga
	"TV.NUI":      "TV-NUI", // Nui
	"TV.VAI":      "TV-VAI", // Vaitupu
	"TW.02":       "TW-KHH", // Takao
	"TW.03":       "TW-TPE", // Taipei
	"TZ.02":       "TZ-19",  // Pwani
	"TZ.03":       "TZ-03",  // Dodoma
	"TZ.04":       "TZ-04",  // Iringa
	"TZ.05":       "TZ-08",  // Kigoma
	"TZ.06":       "TZ-09",  // Kilimanjaro
	"TZ.07":       "TZ-12",  // Lindi
	"TZ.08":       "TZ-13",  // Mara
	"TZ.09":       "TZ-14",  // Mbeya
	"TZ.10":       "TZ-16",  // Morogoro
	"TZ.11":       "TZ-17",  // Mtwara
	"TZ.12":       "TZ-18",  // Mwanza
	"TZ.13":       "TZ-06",  // Pemba North
	"TZ.14":       "TZ-21",  // Ruvuma
	"TZ.15":       "TZ-22",  // Shinyanga
	"TZ.16":       "TZ-23",  // Singida
	"TZ.17":       "TZ-24",  // Tabora
	"TZ.18":       "TZ-25",  // Tanga
	"TZ.19":       "TZ-05",  // Kagera
	"TZ.20":       "TZ-10",  // Pemba South
	"TZ.21":       "TZ-11",  // Zanzibar Central/South
	"TZ.22":       "TZ-07",  // Zanzibar North
	"TZ.23":       "TZ-02",  // Dar es Salaam Region
	"TZ.24":       "TZ-20",  // Rukwa
	"TZ.25":       "TZ-15",  // Zanzibar Urban/West
	"TZ.26":       "TZ-01",  // Arusha
	"TZ.27":       "TZ-26",  // Manyara
	"TZ.28":       "TZ-27",  // Geita
	"TZ.29":       "TZ-28",  // Katavi
	"TZ.30":       "TZ-29",  // Njombe
	"TZ.31":       "TZ-30",  // Simiyu
	"TZ.32":       "TZ-31",  // Songwe
	"UA.02":       "UA-74",  // Chernihiv
	"UA.03":       "UA-77",  // Chernivtsi
	"UA.04":       "UA-12",  // Dnipropetrovsk
	"UA.05":       "UA-14",  // Donetsk
	"UA.07":       "UA-63",  // Kharkivs’ka Oblast’
	"UA.08":       "UA-65",  // Kherson
	"UA.09":       "UA-68",  // Khmelnytskyi
	"UA.10":       "UA-35",  // Kirovohrad
	"UA.11":       "UA-43",  // Crimea
	"UA.12":       "UA-30",  // Kyiv City
	"UA.13":       "UA-30",  // Kyiv
	"UA.14":       "UA-09",  // Luhansk
	"UA.15":       "UA-46",  // Lviv
	"UA.16":       "UA-48",  // Mykolaiv
	"UA.17":       "UA-51",  // Odessa
	"UA.18":       "UA-53",  // Poltava
	"UA.19":       "UA-56",  // Rivne
	"UA.20":       "UA-40",  // Sevastopol City
	"UA.21":       "UA-59",  // Sumy
	"UA.22":       "UA-61",  // Ternopil
	"UA.23":       "UA-05",  // Vinnytsia
	"UA.24":       "UA-07",  // Volyn
	"UA.25":       "UA-21",  // Zakarpattia
	"UA.26":       "UA-23",  // Zaporizhzhia
	"UA.27":       "UA-18",  // Zhytomyr
	"UG.C":        "UG-C",   // Central Region
	"UG.E":        "UG-E",   // Eastern Region
	"UG.N":        "UG-N",   // Northern Region
	"UG.W":        "UG-W",   // Western Region
	"UM.050":      "UM-81",  // Baker Island
	"UM.100":      "UM-84",  // Howland Island
	"UM.150":      "UM-86",  // Jarvis Island
	"UM.200":      "UM-67",  // Johnston Atoll
	"UM.250":      "UM-89",  // Kingman Reef
	"UM.300":      "UM-71",  // Midway Islands
	"UM.350":      "UM-76",  // Navassa Island
	"UM.400":      "UM-95",  // Palmyra Atoll
	"UM.450":      "UM-79",  // Wake Island
	"US.AK":       "US-AK",  // Alaska
	"US.AL":       "US-AL",  // Alabama
	"US.AR":       "US-AR",  // Arkansas
	"US.AZ":       "US-AZ",  // Arizona
	"US.CA":       "US-CA",  // California
	"US.CO":       "US-CO",  // Colorado
	"US.CT":       "US-CT",  // Connecticut
	"US.DC":       "US-DC",  // District of Columbia
	"US.DE":       "US-DE",  // Delaware
	"US.FL":       "US-FL",  // Florida
	"US.GA":       "US-GA",  // Georgia
	"US.HI":       "US-HI",  // Hawaii
	"US.IA":       "US-IA",  // Iowa
	"US.ID":       "US-ID",  // Idaho
	"US.IL":       "US-IL",  // Illinois
	"US.IN":       "US-IN",  // Indiana
	"US.KS":       "US-KS",  // Kansas
	"US.KY":       "US-KY",  // Kentucky
	"US.LA":       "US-LA",  // Louisiana
	"US.MA":       "US-MA",  // Massachusetts
	"US.MD":       "US-MD",  // Maryland
	"US.ME":       "US-ME",  // Maine
	"US.MI":       "US-MI",  // Michigan
	"US.MN":       "US-MN",  // Minnesota
	"US.MO":       "US-MO",  // Missouri
	"US.MS":       "US-MS",  // Mississippi
	"US.MT":       "US-MT",  // Montana
	"US.NC":       "US-NC",  // North Carolina
	"US.ND":       "US-ND",  // North Dakota
	"US.NE":       "US-NE",  // Nebraska
	"US.NH":       "US-NH",  // New Hampshire
	"US.NJ":       "US-NJ",  // New Jersey
	"US.NM":       "US-NM",  // New Mexico
	"US.NV":       "US-NV",  // Nevada
	"US.NY":       "US-NY",  // New York
	"US.OH":       "US-OH",  // Ohio
	"US.OK":       "US-OK",  // Oklahoma
	"US.OR":       "US-OR",  // Oregon
	"US.PA":       "US-PA",  // Pennsylvania
	"US.RI":       "US-RI",  // Rhode Island
	"US.SC":       "US-SC",  // South Carolina
	"US.SD":       "US-SD",  // South Dakota
	"US.TN":       "US-TN",  // Tennessee
	"US.TX":       "US-TX",  // Texas
	"US.UT":       "US-UT",  // Utah
	"US.VA":       "US-VA",  // Virginia
	"US.VT":       "US-VT",  // Vermont
	"US.WA":       "US-WA",  // Washington
	"US.WI":       "US-WI",  // Wisconsin
	"US.WV":       "US-WV",  // West Virginia
	"US.WY":       "US-WY",  // Wyoming
	"UY.01":       "UY-AR",  // Artigas
	"UY.02":       "UY-CA",  // Canelones
	"UY.03":       "UY-CL",  // Cerro Largo
	"UY.04":       "UY-CO",  // Colonia
	"UY.05":       "UY-DU",  // Durazno Department
	"UY.06":       "UY-FS",  // Flores Department
	"UY.07":       "UY-FD",  // Florida
	"UY.08":       "UY-LA",  // Lavalleja
	"UY.09":       "UY-MA",  // Maldonado
	"UY.10":       "UY-MO",  // Montevideo Department
	"UY.11":       "UY-PA",  // Paysandú Department
	"UY.12":       "UY-RN",  // Río Negro Department
	"UY.13":       "UY-RV",  // Rivera Department
	"UY.14":       "UY-RO",  // Rocha Department
	"UY.15":       "UY-SA",  // Salto Department
	"UY.16":       "UY-SJ",  // San José Department
	"UY.17":       "UY-SO",  // Soriano
	"UY.18":       "UY-TA",  // Tacuarembó Department
	"UY.19":       "UY-TT",  // Treinta y Tres Department
	"UZ.01":       "UZ-AN",  // Andijan Region
	"UZ.02":       "UZ-BU",  // Bukhara
	"UZ.03":       "UZ-FA",  // Fergana
	"UZ.05":       "UZ-XO",  // Xorazm Region
	"UZ.06":       "UZ-NG",  // Namangan
	"UZ.07":       "UZ-NW",  // Navoiy Region
	"UZ.08":       "UZ-QA",  // Qashqadaryo
	"UZ.09":       "UZ-QR",  // Karakalpakstan
	"UZ.10":       "UZ-SA",  // Samarqand Region
	"UZ.12":       "UZ-SU",  // Surxondaryo Region
	"UZ.13":       "UZ-TK",  // Tashkent
	"UZ.14":       "UZ-TO",  // Tashkent Region
	"UZ.15":       "UZ-JI",  // Jizzakh Region
	"UZ.16":       "UZ-SI",  // Sirdaryo Region
	"VC.01":       "VC-01",  // Charlotte Parish
	"VC.02":       "VC-02",  // Saint Andrew Parish
	"VC.03":       "VC-03",  // Saint David Parish
	"VC.04":       "VC-04",  // Saint George Parish
	"VC.05":       "VC-05",  // Saint Patrick Parish
	"VC.06":       "VC-06",  // Grenadines Parish
	"VE.01":       "VE-Z",   // Amazonas
	"VE.02":       "VE-B",   // Anzoátegui
	"VE.03":       "VE-C",   // Apure
	"VE.04":       "VE-D",   // Aragua
	"VE.05":       "VE-E",   // Barinas
	"VE.06":       "VE-F",   // Bolívar
	"VE.07":       "VE-G",   // Carabobo
	"VE.08":       "VE-H",   // Cojedes
	"VE.09":       "VE-Y",   // Delta Amacuro
	"VE.11":       "VE-I",   // Falcón
	"VE.12":       "VE-J",   // Guárico
	"VE.13":       "VE-K",   // Lara
	"VE.14":       "VE-L",   // Mérida
	"VE.15":       "VE-M",   // Miranda
	"VE.16":       "VE-N",   // Monagas
	"VE.17":       "VE-O",   // Nueva Esparta
	"VE.18":       "VE-P",   // Portuguesa
	"VE.19":       "VE-R",   // Sucre
	"VE.20":       "VE-S",   // Táchira
	"VE.21":       "VE-T",   // Trujillo
	"VE.22":       "VE-U",   // Yaracuy
	"VE.23":       "VE-V",   // Zulia
	"VE.24":       "VE-W",   // Dependencias Federales
	"VE.25":       "VE-A",   // Distrito Federal
	"VE.26":       "VE-X",   // Vargas
	"VN.01":       "VN-HN",  // Hanoi
	"VN.04":       "VN-04",  // Cao Bằng Province
	"VN.08":       "VN-07",  // Tuyen Quang
	"VN.11":       "VN-71",  // Điện Biên Province
	"VN.12":       "VN-01",  // Lai Châu Province
	"VN.14":       "VN-05",  // Sơn La Province
	"VN.15":       "VN-02",  // Lao Cai
	"VN.19":       "VN-69",  // Thai Nguyen
	"VN.20":       "VN-09",  // Lạng Sơn Province
	"VN.22":       "VN-13",  // Quảng Ninh
	"VN.24":       "VN-56",  // Bac Ninh
	"VN.25":       "VN-68",  // Phu Tho
	"VN.31":       "VN-HP",  // Hai Phong
	"VN.33":       "VN-66",  // Hưng Yên Province
	"VN.37":       "VN-18",  // Ninh Binh
	"VN.38":       "VN-21",  // Thanh Hóa Province
	"VN.40":       "VN-22",  // Nghệ An Province
	"VN.42":       "VN-23",  // Hà Tĩnh Province
	"VN.44":       "VN-25",  // Quang Tri
	"VN.46":       "VN-26",  // Thừa Thiên Huế Province
	"VN.48":       "VN-DN",  // Da Nang City
	"VN.51":       "VN-29",  // Quang Ngai
	"VN.52":       "VN-30",  // Gia Lai
	"VN.56":       "VN-34",  // Khanh Hoa
	"VN.66":       "VN-33",  // Dak Lak
	"VN.68":       "VN-35",  // Lam Dong
	"VN.75":       "VN-39",  // Dong Nai
	"VN.79":       "VN-SG",  // Ho Chi Minh City (HCMC)
	"VN.80":       "VN-37",  // Tay Ninh
	"VN.82":       "VN-45",  // Dong Thap
	"VN.86":       "VN-49",  // Vinh Long
	"VN.91":       "VN-44",  // An Giang
	"VN.92":       "VN-CT",  // Can Tho City
	"VN.96":       "VN-59",  // Ca Mau
	"VU.07":       "VU-TOB", // Torba
	"VU.13":       "VU-SAM", // Sanma
	"VU.15":       "VU-TAE", // Tafea
	"VU.16":       "VU-MAP", // Malampa
	"VU.17":       "VU-PAM", // Penama
	"VU.18":       "VU-SEE", // Shefa
	"WF.98611":    "WF-AL",  // Alo
	"WF.98612":    "WF-SG",  // Sigave
	"WF.98613":    "WF-UV",  // Uvea
	"WS.01":       "WS-AA",  // A'ana
	"WS.02":       "WS-AL",  // Aiga-i-le-Tai
	"WS.03":       "WS-AT",  // Atua
	"WS.04":       "WS-FA",  // Fa'asaleleaga
	"WS.05":       "WS-GE",  // Gaga'emauga
	"WS.06":       "WS-VF",  // Va'a-o-Fonoti
	"WS.07":       "WS-GI",  // Gaga'ifomauga
	"WS.08":       "WS-PA",  // Palauli
	"WS.09":       "WS-SA",  // Satupa'itea
	"WS.10":       "WS-TU",  // Tuamasaga
	"WS.11":       "WS-VS",  // Vaisigano
	"YE.01":       "YE-AB",  // Abyan Governorate
	"YE.02":       "YE-AD",  // Aden
	"YE.03":       "YE-MR",  // Al Mahrah Governorate
	"YE.04":       "YE-HD",  // Muhafazat Hadramaout
	"YE.05":       "YE-SH",  // Shabwah
	"YE.08":       "YE-HU",  // Al Hudaydah
	"YE.10":       "YE-MW",  // Al Mahwit Governorate
	"YE.11":       "YE-DH",  // Dhamār
	"YE.14":       "YE-MA",  // Ma’rib
	"YE.15":       "YE-SD",  // Şa‘dah
	"YE.16":       "YE-SN",  // Sanaa Governorate
	"YE.19":       "YE-AM",  // Omran
	"YE.20":       "YE-BA",  // Al Bayda
	"YE.21":       "YE-JA",  // Al Jawf
	"YE.22":       "YE-HJ",  // Ḩajjah
	"YE.23":       "YE-IB",  // Ibb Governorate
	"YE.24":       "YE-LA",  // Laḩij
	"YE.25":       "YE-TA",  // Ta‘izz
	"YE.26":       "YE-SA",  // Amanat Alasimah
	"YE.27":       "YE-RA",  // Raymah
	"YE.28":       "YE-SU",  // Soqatra
	"ZA.02":       "ZA-KZN", // KwaZulu-Natal
	"ZA.03":       "ZA-FS",  // Free State
	"ZA.05":       "ZA-EC",  // Eastern Cape
	"ZA.06":       "ZA-GP",  // Gauteng
	"ZA.07":       "ZA-MP",  // Mpumalanga
	"ZA.08":       "ZA-NC",  // Northern Cape
	"ZA.09":       "ZA-LP",  // Limpopo
	"ZA.10":       "ZA-NW",  // North West
	"ZA.11":       "ZA-WC",  // Western Cape
	"ZM.01":       "ZM-01",  // Western Province
	"ZM.02":       "ZM-02",  // Central Province
	"ZM.03":       "ZM-03",  // Eastern Province
	"ZM.04":       "ZM-04",  // Luapula Province
	"ZM.05":       "ZM-05",  // Northern Province
	"ZM.06":       "ZM-06",  // North-Western
	"ZM.07":       "ZM-07",  // Southern Province
	"ZM.08":       "ZM-08",  // Copperbelt
	"ZM.09":       "ZM-09",  // Lusaka Province
	"ZM.10":       "ZM-10",  // Muchinga
	"ZW.01":       "ZW-MA",  // Manicaland
	"ZW.02":       "ZW-MI",  // Midlands Province
	"ZW.03":       "ZW-MC",  // Mashonaland Central
	"ZW.04":       "ZW-ME",  // Mashonaland East Province
	"ZW.05":       "ZW-MW",  // Mashonaland West
	"ZW.06":       "ZW-MN",  // Matabeleland North
	"ZW.07":       "ZW-MS",  // Matabeleland South Province
	"ZW.08":       "ZW-MV",  // Masvingo Province
	"ZW.09":       "ZW-BU",  // Bulawayo
	"ZW.10":       "ZW-HA",  // Harare
}
