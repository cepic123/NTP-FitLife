<h1>Projekat iz predmeta Napredne tehnike programiranja</h1>

<h2>FitLife - Aplikacija za praćenje svog fitness napretka uz pregled planova ishrane i treninga postavljenih od strane trenera.</h2>

<h3>Kratak pregled funkcionalnosti</h3>

Neregistrovani korisnik:
  - pregled treninga
  - pregled planova ishrane
  - registracija kao korisnik
  - podnošenje zahteva za registraciju kao trener

Korisnik:
  - prijava na sistem
  - odjava sa sistema
  - pregled i izmena profila
  - pretraga, pregled, ocenjivanje i komentarijasnje treninga
  - pretraga, pregled, ocenjivanje i komentarijasnje planova ishrane
  - pretraga, pregled, ocenjivanje i komentarijasnje trenera
  * pod pregledom i pretragom podrazumeva se prikaz u vidu liste kao i detaljan prikaz
  - praćenje istorije ishrane i treninga na osnovu kalendara sa mogućnostima unošenja treninga/ishrane na dnevnom nivou
  - podnošenje žalbi i blokiranje drugih korisnika/trenera
  - mogućnost pretplate na premium korisnika
  
Trener:
  - sve funkcionalnosti korisnika
  - kreiranje pojedinačnih vežbi uz mogućnost uploada slike/videa
  - kreiranje treninga kombinovanjem vežbi
  - kreiranje planova ishrane
  - mogućnost deljenja treninga i planova ishrane na besplatne i samo dostupne premium korisnicima
  
Administrator sistema:
  - prihvatanje/odbijanje prijava zahteva za registraciju trenera
  - pregled žalbi i blokiranih korisnika uz dalju mogućnost brisanje profila korisnika i trenera

*predlozi za dodatno proširenje: 
  - mogućnost dopisivanja trenera i premium korisnika
 
<h3>Arhitektura sistema</h3>

Servisi:
  - Gateway servis (Go/PostgrateSQL)
  - korisnički servis (Go/PostgrateSQL)
  - servis za komentare (Go/PostgrateSQL)
  - servis za ocenjivanje (Go/PostgrateSQL)
  - servis za treninge (Go/PostgrateSQL)
  - servis za planove ishrane (Go/PostgrateSQL)
  - servis za žalbe (Go/PostgrateSQL)
  - servis za blokiranje (Rust/PostgrateSQL)
  
Klijentska veb aplikacija:
  - Monolitna angular aplikacija
