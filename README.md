# datafolket
Oppgave 1                                                                             
Binære tall	                                                                                                          
1101	             	             
110111101010	                                          	            
1010111100110100	                                               
1111111111111111	                                                        	          
10001011110001010                     

Hexadesimaltall                                                                                                                   
0xD                                                       
0xDEA                                                             
0xAF34                                                            
0xFFFF                                                        
0x1178A	                                                          

Desimaltall                                                                                                                   
13                                                              
3562                                                        
44852                                                                   
65535                                                           
71562                                                                   

1a.

Dele opp tallet i grupper på fire, fylle inn manglende tall med 0 og så starte bakerst i tallet og oversette verdien fra et tallsystem til det andre. F.eks 1110 blir 14, som i heksadesimal er E. Om dette var det første av to tallgrupper på 4 så har jeg oversatt halve tallet til heksadesimal. Om det andre tallet i rekka var 1001, så oversettes det til 9, og nå har hele tallet blitt oversatt fra 1001110 til 9+E, eller 9E. For å gjøre det motsatt så ser man på hvert heksadesimale siffer og finner tilsvarende siffer i binær verdi.                                                                                                                                      For å gå fra binær til desimaltall så må man ta sifferet (0 eller 1) ganger 2 som er grunntallet til det binære tallsystemet og opphøye i posisjonen til tallet, som starter fra 0 helt til høyre i tallet og går oppover tilsvarende hvor langt tallet er. F.eks 101010 så er tallet i posisjon 3 = 1, og må først ganges med 2 (grunntallet) for og så opphøyes i 3 som er posisjonen.  1*2^3=8. Dette må gjøres for alle tallene og blir 42.  For å gå fra desimal til binær må man veldig enkelt dele på 2 til svaret blir 0. Om det blir rest så skriver man ned resten, som alltid er 1, om det ikke blir rest skriver man 0. Dette gjøres helt til man har konvertert hele tallet. Et eksempel med tallet 65 er under.

65 / 2 = 32,5 runde av til 32, fordi det blir rest. =1
32 /2= 16, ingen rest.=0
16 /2= 8, ingen rest =0  
8/2= 4, ingen rest =0
4/2 = 2, ingen rest = 0
2/2 = 0, ingen rest =0
1/2= 0.5, runder av til 0, fordi det blir rest= 1

1000001

1b.

For å gå fra heksadesimal så må man starte med å vite at grunntallet er 16. Så finne verdiene til heksadesimalene, og gange dem med 16 opphøyd i posisjon. F.eks i BEEF så er E i posisjon 1. Da vil utregningen se slik ut. 14*16^1 = 224. Dette må gjøres for alle verdiene i posisjonene. For å regne fra desimal til heksadesimal så følger man samme metode fra desimal til binær, men man deler på 16 istedenfor 2, og rest kan være andre tall enn 1 og 0. Man fyller inn bokstavene tilsvarende verdiene der det trengs. Hvis det blir 13 i rest, så blir heksadesimal resten tilsvarende D.



Oppgave 2

https://github.com/felpus/datafolket/tree/master/Oppgave%202/algorithms

I sorting.go er den modifiserte bubble sort funksjonen som er en raskere bubble sort algorytme.

Sorting_test.go har en samling benchmarks som lar oss teste effektiviteten på de forskjellige algorytmene i programmet. 
https://github.com/felpus/datafolket/blob/master/Oppgave%202/algorithms/benchmarkresult.PNG

Bubble sort er en algorytme som blir svært ineffektiv jo større utregningen er tar, den er plasseffektiv men fungerer svært dårlig ved høyt antall operasjoner.

Vår modifiserte bubble stort algorytme er betydelig raskere, den er nesten 20% raskere på 10000 benchmarken, ca 40% på 1000 og går ned til ca 10% på 100.

Quicksort algorytmen er betydelig raskere, nesten tre ganger så rask i alle tre benchmark testene.

Ifølge Big-O så er quicksort algorytmen bedre i et gjennomsnittlig tilfelle og tilsvarende bubble sort i det dårligste tilfelle. I beste tilfelle så er faktisk Bubble sort mer effektiv.

Oppgave 3                                                                                 
Denne oppgaven er besvart i infinite_main.go og infinite.go i mappen infinite. Programmet bruker ikke nevneverdig mye CPU og RAM når det kjører (1-2%). Når SIGINT-signal mottas skriver programmet ut «You pressed ctrl + C. User interrupted infinite loop.». 
Programmet reagerer også på kommandoene:                                                                                
quit                                                                                                              
term                                                                                                    
Programmet reagerer ved å produsere meldingene «Prosessen er avsluttet.» eller «Prosessen er terminert.».

Oppgave 4


