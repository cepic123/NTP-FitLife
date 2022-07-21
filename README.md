<h1>Projekat iz predmeta Napredne tehnike programiranja</h1>

<h2>Aplikacija za pronalazak/postavljanje oglasa za posao</h2>

Neregistrovani korisnik:
  - pregled poslova
  - pregled poslodavaca
  - registracija kao radnik
  - podnošenje zahteva za registraciju kao poslodavac

Radnik i poslodavac:
  - prijava na sistem
  - odjava sa sistema
  - izmena profila
  - podnošenje žalbe na druge korisnike i mogućnost blokiranja
  - pregled istorije i izveštaja na osnovu prethodnih poslova koje je izdavao/radio
  
Radnik:
  - pregled, filtriranje i sortiranje poslova po određenim parametrima:ocena poslodavca, plata od/do, tip posla(fizički posao, programiranje, uslužna delatnost...), vrsta ugovora(jednokratni, part-time, full-time..), kompanija
  - pregled, filtriranje i sortiranje poslodavaca po određenim parametrima: prosečna plata, broj zaposlenih, prosečna ocena..
  - detaljan pregled oglasa za posao
  - detaljan pregled profila poslodavaca
  - prijava na oglas 
  - ocenjivanje i komentarisanje poslodavaca u koliko su imali saradnju u prošlosti
  
Poslodavac:
  - crud ponuda za posao (definisanje plate, lokacije, tipa posla, vrste ugovora..)
  - pregled/prihvatanje/odbijanje prijava ranika na oglase koje je postavio
  - detaljan pregled profila radnika
  - ocenjivanje i komentarisanje radnika u koliko su imali saradnju u prošlosti
 
Administrator sistema:
  - prihvatanje/odbijanje prijava zahteva za registraciju
  - pregled žalbi i blokiranih korisnika uz dalju mogućnost brisanje profila radnika i poslodavaca

*predlozi za proširenje - mogućnost pregleda/postavke odredjene lokacije posla na mapi pri čemu bi korisnik mogao da bira poslove do određene granice udaljenosti/dužine putovanja
  - mogućnost dopisivanja radnika i poslodavca
 
Servisi:
  - korisnički servis (Go/PostgrateSQL)
  - servis za komentare (Go/PostgrateSQL)
  - servis za ocenjivanje (Go/PostgrateSQL)
  - servis za oglase (Go/PostgrateSQL)

Klijentska veb aplikacija:
  - Monolitna angular aplikacija
